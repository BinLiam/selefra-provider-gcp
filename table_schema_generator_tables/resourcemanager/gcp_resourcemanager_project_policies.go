package resourcemanager

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	pb "google.golang.org/api/cloudresourcemanager/v3"
)

type TableGcpResourcemanagerProjectPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpResourcemanagerProjectPoliciesGenerator{}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetTableName() string {
	return "gcp_resourcemanager_project_policies"
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			projectsClient, err := pb.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			output, err := projectsClient.Projects.GetIamPolicy("projects/"+c.ProjectId, &pb.GetIamPolicyRequest{}).Context(ctx).Do()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("audit_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AuditConfigs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bindings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Bindings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpResourcemanagerProjectPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
