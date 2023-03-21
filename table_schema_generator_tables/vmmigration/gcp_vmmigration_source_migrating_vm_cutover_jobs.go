package vmmigration

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator{}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetTableName() string {
	return "gcp_vmmigration_source_migrating_vm_cutover_jobs"
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.MigratingVm)

			gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListCutoverJobs(ctx, &pb.ListCutoverJobsRequest{
				Parent: parentItem.Name,
			}, c.CallOptions...)
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

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vmmigration_source_migrating_vms_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vmmigration_source_migrating_vms.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("progress_percent").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProgressPercent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StateMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
