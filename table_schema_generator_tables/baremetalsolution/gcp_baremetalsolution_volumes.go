package baremetalsolution

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBaremetalsolutionVolumesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBaremetalsolutionVolumesGenerator{}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetTableName() string {
	return "gcp_baremetalsolution_volumes"
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListVolumesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListVolumes(ctx, req, c.CallOptions...)
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

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("emergency_size_gib").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EmergencySizeGib")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_grown_size_gib").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AutoGrownSizeGib")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_reservation_detail").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SnapshotReservationDetail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SnapshotEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_size_gib").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CurrentSizeGib")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requested_size_gib").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RequestedSizeGib")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StorageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pod").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Pod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("remaining_space_gib").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RemainingSpaceGib")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_auto_delete_behavior").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SnapshotAutoDeleteBehavior")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
	}
}

func (x *TableGcpBaremetalsolutionVolumesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpBaremetalsolutionVolumeLunsGenerator{}),
	}
}
