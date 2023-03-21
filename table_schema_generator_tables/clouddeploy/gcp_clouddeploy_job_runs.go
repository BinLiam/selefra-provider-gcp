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

type TableGcpClouddeployJobRunsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouddeployJobRunsGenerator{}

func (x *TableGcpClouddeployJobRunsGenerator) GetTableName() string {
	return "gcp_clouddeploy_job_runs"
}

func (x *TableGcpClouddeployJobRunsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouddeployJobRunsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouddeployJobRunsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpClouddeployJobRunsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListJobRunsRequest{
				Parent: task.ParentRawResult.(*pb.Rollout).Name,
			}
			gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListJobRuns(ctx, req, c.CallOptions...)
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

func (x *TableGcpClouddeployJobRunsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouddeployJobRunsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phase_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PhaseId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_clouddeploy_rollouts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_clouddeploy_rollouts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("JobId")).Build(),
	}
}

func (x *TableGcpClouddeployJobRunsGenerator) GetSubTables() []*schema.Table {
	return nil
}
