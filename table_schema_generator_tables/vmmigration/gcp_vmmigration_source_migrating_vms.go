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

type TableGcpVmmigrationSourceMigratingVmsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVmmigrationSourceMigratingVmsGenerator{}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetTableName() string {
	return "gcp_vmmigration_source_migrating_vms"
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.Source)

			gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListMigratingVms(ctx, &pb.ListMigratingVmsRequest{
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

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_sync").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastSync")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recent_clone_jobs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RecentCloneJobs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_vm_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceVmId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vmmigration_sources_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vmmigration_sources.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Policy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_sync_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CurrentSyncInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Group")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("recent_cutover_jobs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RecentCutoverJobs")).Build(),
	}
}

func (x *TableGcpVmmigrationSourceMigratingVmsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpVmmigrationSourceMigratingVmCloneJobsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpVmmigrationSourceMigratingVmCutoverJobsGenerator{}),
	}
}
