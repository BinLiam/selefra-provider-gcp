package securitycenter

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	securitycenter "cloud.google.com/go/securitycenter/apiv1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpSecuritycenterOrganizationFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSecuritycenterOrganizationFindingsGenerator{}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetTableName() string {
	return "gcp_securitycenter_organization_findings"
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := "organizations/" + c.OrgId + "/sources/-"
			maybeError := fetchFindings(p)(ctx, client, resultChannel)
			return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
		},
	}
}

func fetchFindings(parent string) func(ctx context.Context, client any, res chan<- any) error {
	return func(ctx context.Context, client any, res chan<- any) error {
		c := client.(*gcp_client.Client)
		req := &pb.ListFindingsRequest{
			Parent:   parent,
			PageSize: 1000,
		}
		gcpClient, err := securitycenter.NewClient(ctx, c.ClientOptions...)
		if err != nil {
			return err
		}
		it := gcpClient.ListFindings(ctx, req, c.CallOptions...)
		for {
			resp, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				return err
			}

			res <- resp
		}
		return nil
	}
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandOrgMultiplex()
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state_change").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StateChange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorOrganization()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Finding.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finding").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Finding")).Build(),
	}
}

func (x *TableGcpSecuritycenterOrganizationFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
