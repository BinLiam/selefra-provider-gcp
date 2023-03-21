package clouddeploy

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	deploy "cloud.google.com/go/deploy/apiv1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpClouddeployTargetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouddeployTargetsGenerator{}

func (x *TableGcpClouddeployTargetsGenerator) GetTableName() string {
	return "gcp_clouddeploy_targets"
}

func (x *TableGcpClouddeployTargetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouddeployTargetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouddeployTargetsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpClouddeployTargetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListTargetsRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListTargets(ctx, req, c.CallOptions...)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- resp
			}
			return nil
		},
	}
}

func (x *TableGcpClouddeployTargetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouddeployTargetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TargetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_approval").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequireApproval")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExecutionConfigs")).Build(),
	}
}

func (x *TableGcpClouddeployTargetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
