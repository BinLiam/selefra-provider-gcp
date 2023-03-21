package batch

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	batch "cloud.google.com/go/batch/apiv1"
	pb "cloud.google.com/go/batch/apiv1/batchpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBatchTaskGroupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBatchTaskGroupsGenerator{}

func (x *TableGcpBatchTaskGroupsGenerator) GetTableName() string {
	return "gcp_batch_task_groups"
}

func (x *TableGcpBatchTaskGroupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBatchTaskGroupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBatchTaskGroupsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBatchTaskGroupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListJobsRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := batch.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListJobs(ctx, req, c.CallOptions...)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- resp.TaskGroups
			}
			return nil
		},
	}
}

func (x *TableGcpBatchTaskGroupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBatchTaskGroupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("task_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TaskSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TaskCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parallelism").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Parallelism")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_hosts_file").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequireHostsFile")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_environments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TaskEnvironments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("task_count_per_node").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TaskCountPerNode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("permissive_ssh").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PermissiveSsh")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBatchTaskGroupsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpBatchTasksGenerator{}),
	}
}
