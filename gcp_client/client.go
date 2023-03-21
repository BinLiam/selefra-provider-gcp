package gcp_client

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/googleapis/gax-go/v2"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	crmv1 "google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	"google.golang.org/grpc/codes"
)

const maxProjectIdsToLog int = 100

const serviceAccountEnvKey = "SELEFRA_SERVICE_ACCOUNT_KEY_JSON"

type Client struct {
	projects []string
	orgs     []string

	ClientOptions []option.ClientOption
	CallOptions   []gax.CallOption

	EnabledServices map[string]map[string]any
	// this is set by table client project multiplexer
	ProjectId string
	// this is set by table client Org multiplexer
	OrgId string
	// this is set by table client Location multiplexer
	Location string
	// Logger
	logger zerolog.Logger
}

//revive:disable:modifies-value-receiver

// withProject allows multiplexer to create a new client with given projectId
func (c *Client) withProject(project string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("project_id", project).Logger()
	newClient.ProjectId = project
	return &newClient
}

func (c *Client) withLocation(location string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("location", location).Logger()
	newClient.Location = location
	return &newClient
}

// withOrg allows multiplexer to create a new client with given organizationId
func (c *Client) withOrg(org string) *Client {
	newClient := *c
	newClient.logger = c.logger.With().Str("org_id", org).Logger()
	newClient.OrgId = org
	return &newClient
}

func isValidJson(content []byte) error {
	var v map[string]any
	err := json.Unmarshal(content, &v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ID() string {
	if c.OrgId != "" {
		return "org:" + c.OrgId
	}
	if c.Location != "" {
		return "project:" + c.ProjectId + ":location:" + c.Location
	}
	return "project:" + c.ProjectId
}

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func NewClients(config Config) ([]*Client, error) {
	ctx := context.Background()
	client, err := New(ctx, config)
	if err != nil {
		return nil, err
	}
	return []*Client{client}, nil
}

func New(ctx context.Context, config Config) (*Client, error) {
	var err error

	c := Client{
		EnabledServices: map[string]map[string]any{},
	}

	config.setDefaults()

	projects := config.ProjectIDs

	if config.BackoffRetries > 0 {
		c.CallOptions = append(c.CallOptions, gax.WithRetry(func() gax.Retryer {
			return &Retrier{
				backoff: gax.Backoff{
					Max: time.Duration(config.BackoffDelay) * time.Second,
				},
				maxRetries: config.BackoffRetries,
				codes:      []codes.Code{codes.ResourceExhausted},
			}
		}))
	}

	serviceAccountKeyJSON := []byte(config.ServiceAccountKeyJSON)

	if len(serviceAccountKeyJSON) == 0 && config.ServiceAccountKeyJSONFile != "" {
		data, err := os.ReadFile(config.ServiceAccountKeyJSONFile)
		if err == nil {
			serviceAccountKeyJSON = data
		}
	}

	if len(serviceAccountKeyJSON) == 0 {
		if config.ServiceAccountEnvKey == "" {
			config.ServiceAccountEnvKey = serviceAccountEnvKey
		}

		data, err := os.ReadFile(os.Getenv(config.ServiceAccountEnvKey))
		if err == nil {
			serviceAccountKeyJSON = data
		}
	}

	c.ClientOptions = append(c.ClientOptions, option.WithRequestReason("selefra resource fetch"))

	if len(serviceAccountKeyJSON) != 0 {
		if err := isValidJson(serviceAccountKeyJSON); err != nil {
			return nil, err
		}
		c.ClientOptions = append(c.ClientOptions, option.WithCredentialsJSON(serviceAccountKeyJSON))
	}

	if len(config.ProjectFilter) > 0 && len(config.FolderIDs) > 0 {
		return nil, fmt.Errorf("project_filter and folder_ids are mutually exclusive")
	}

	projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create projects client: %w", err)
	}
	foldersClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create folders client: %w", err)
	}

	switch {
	case len(projects) == 0 && len(config.FolderIDs) == 0 && len(config.ProjectFilter) == 0:
		c.logger.Info().Msg("No project_ids, folder_ids, or project_filter specified - assuming all active projects")
		projects, err = getProjectsV1(ctx, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}

	case len(config.FolderIDs) > 0:
		var folderIds []string

		for _, parentFolder := range config.FolderIDs {
			c.logger.Info().Msg("Listing folders...")
			childFolders, err := listFolders(ctx, foldersClient, parentFolder, *config.FolderRecursionDepth)
			if err != nil {
				return nil, fmt.Errorf("failed to list folders: %w", err)
			}
			folderIds = append(folderIds, childFolders...)
		}

		logFolderIds(&c.logger, folderIds)

		c.logger.Info().Msg("listing folder projects...")
		folderProjects, err := listProjectsInFolders(ctx, projectsClient, folderIds)
		projects = setUnion(projects, folderProjects)
		if err != nil {
			return nil, fmt.Errorf("failed to list projects: %w", err)
		}

	case len(config.ProjectFilter) > 0:
		c.logger.Info().Msg("Listing projects with filter...")
		projectsWithFilter, err := getProjectsV1WithFilter(ctx, config.ProjectFilter, c.ClientOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to get projects with filter: %w", err)
		}

		projects = setUnion(projects, projectsWithFilter)
	}

	c.projects = projects

	c.orgs, err = getOrganizations(ctx, c.ClientOptions...)
	if err != nil {
		c.logger.Err(err).Msg("failed to get organizations")
	}
	c.logger.Info().Interface("orgs", c.orgs).Msg("Retrieved organizations")

	if len(projects) == 1 {
		c.ProjectId = projects[0]
	}
	//if config.EnabledServicesOnly {
	//	if err := c.configureEnabledServices(ctx, s.Concurrency); err != nil {
	//		if status.Code(err) == codes.ResourceExhausted {
	//			c.logger.Err(err).Msg("failed to list enabled services because of rate limiting. Consider setting larger values for `backoff_retries` and `backoff_delay`")
	//		} else {
	//			c.logger.Err(err).Msg("failed to list enabled services")
	//		}
	//		return nil, err
	//	}
	//}

	return &c, nil
}

func logFolderIds(logger *zerolog.Logger, folderIds []string) {
	// If there are too many folders, just log the first maxProjectIdsToLog.
	if len(folderIds) > maxProjectIdsToLog {
		logger.Info().Interface("folder_ids", folderIds[:maxProjectIdsToLog]).Msgf("Found %d folders. First %d: ", len(folderIds), maxProjectIdsToLog)
		logger.Debug().Interface("folder_ids", folderIds).Msg("All folders: ")
	} else {
		logger.Info().Interface("folder_ids", folderIds).Msgf("Found %d projects in folders", len(folderIds))
	}
}

// getProjectsV1 requires the `resourcemanager.projects.get` permission to list projects
func getProjectsV1(ctx context.Context, options ...option.ClientOption) ([]string, error) {
	var projects []string

	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	call := service.Projects.List().Filter("lifecycleState=ACTIVE").Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			projects = append(projects, project.ProjectId)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	if len(projects) == 0 {
		return nil, fmt.Errorf("no active projects")
	}

	return projects, nil
}

func getProjectsV1WithFilter(ctx context.Context, filter string, options ...option.ClientOption) ([]string, error) {
	var projects []string

	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	call := service.Projects.List().Filter(filter).Context(ctx)
	for {
		output, err := call.Do()
		if err != nil {
			return nil, err
		}
		for _, project := range output.Projects {
			if project.LifecycleState != "ACTIVE" {
				continue
			}
			projects = append(projects, project.ProjectId)
		}
		if output.NextPageToken == "" {
			break
		}
		call.PageToken(output.NextPageToken)
	}

	return projects, nil
}

// listFolders recursively lists the folders in the 'parent' folder. Includes the 'parent' folder itself.
// recursionDepth is the depth of folders to recurse - where 0 means not to recurse any folders.
func listFolders(ctx context.Context, folderClient *resourcemanager.FoldersClient, parent string, recursionDepth int) ([]string, error) {
	folders := []string{
		parent,
	}
	if recursionDepth <= 0 {
		return folders, nil
	}

	it := folderClient.ListFolders(ctx, &resourcemanagerpb.ListFoldersRequest{
		Parent: parent,
	})

	for {
		child, err := it.Next()

		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		if child.State == resourcemanagerpb.Folder_ACTIVE {
			childFolders, err := listFolders(ctx, folderClient, child.Name, recursionDepth-1)
			if err != nil {
				return nil, err
			}
			folders = append(folders, childFolders...)
		}
	}

	return folders, nil
}

func listProjectsInFolders(ctx context.Context, projectClient *resourcemanager.ProjectsClient, folders []string) ([]string, error) {
	var projects []string
	for _, folder := range folders {
		it := projectClient.ListProjects(ctx, &resourcemanagerpb.ListProjectsRequest{
			Parent: folder,
		})

		for {
			project, err := it.Next()

			if err == iterator.Done {
				break
			}
			if err != nil {
				return nil, err
			}

			if project.State == resourcemanagerpb.Project_ACTIVE {
				projects = append(projects, project.ProjectId)
			}
		}
	}

	return projects, nil
}

func getOrganizations(ctx context.Context, options ...option.ClientOption) ([]string, error) {
	service, err := crmv1.NewService(ctx, options...)
	if err != nil {
		return nil, fmt.Errorf("failed to create cloudresourcemanager service: %w", err)
	}

	var orgs []string
	if err := service.Organizations.Search(&crmv1.SearchOrganizationsRequest{}).Context(ctx).Pages(ctx, func(page *crmv1.SearchOrganizationsResponse) error {
		for _, org := range page.Organizations {
			orgs = append(orgs, strings.TrimPrefix(org.Name, "organizations/"))
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return setUnion(nil, orgs), nil
}

func setUnion(a []string, b []string) []string {
	set := make(map[string]struct{}, len(a)+len(b)) // alloc max
	for _, s := range a {
		set[s] = struct{}{}
	}
	for _, s := range b {
		set[s] = struct{}{}
	}

	union := make([]string, 0, len(set))
	for s := range set {
		union = append(union, s)
	}
	return union
}

func (c *Client) configureEnabledServices(ctx context.Context, concurrency uint64) error {
	var esLock sync.Mutex
	g, ctx := errgroup.WithContext(ctx)
	goroutinesSem := semaphore.NewWeighted(int64(concurrency))
	for _, p := range c.projects {
		project := p
		if err := goroutinesSem.Acquire(ctx, 1); err != nil {
			return err
		}
		g.Go(func() error {
			defer goroutinesSem.Release(1)
			cl := c.withProject(project)
			svc, err := cl.fetchEnabledServices(ctx)
			esLock.Lock()
			c.EnabledServices[project] = svc
			esLock.Unlock()
			return err
		})
	}
	return g.Wait()
}

func (c *Client) fetchEnabledServices(ctx context.Context) (map[string]any, error) {
	enabled := make(map[string]any)
	req := &pb.ListServicesRequest{
		Parent:   "projects/" + c.ProjectId,
		PageSize: 200,
		Filter:   "state:ENABLED",
	}
	gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
	if err != nil {
		return nil, err
	}
	it := gcpClient.ListServices(ctx, req, c.CallOptions...)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		enabled[resp.GetConfig().Name] = resp
	}
	return enabled, nil
}
