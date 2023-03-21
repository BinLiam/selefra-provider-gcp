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

type TableGcpVmmigrationSourceUtilizationReportsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVmmigrationSourceUtilizationReportsGenerator{}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetTableName() string {
	return "gcp_vmmigration_source_utilization_reports"
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.Source)

			gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListUtilizationReports(ctx, &pb.ListUtilizationReportsRequest{
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

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("frame_end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FrameEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vms").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Vms")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_frame").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TimeFrame")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VmCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vmmigration_sources_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vmmigration_sources.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableGcpVmmigrationSourceUtilizationReportsGenerator) GetSubTables() []*schema.Table {
	return nil
}
