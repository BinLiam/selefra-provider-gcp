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

type TableGcpBaremetalsolutionVolumeLunsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBaremetalsolutionVolumeLunsGenerator{}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetTableName() string {
	return "gcp_baremetalsolution_volume_luns"
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListLunsRequest{
				Parent: task.ParentRawResult.(*pb.Volume).Name,
			}
			gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListLuns(ctx, req, c.CallOptions...)
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

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shareable").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Shareable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("boot_lun").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("BootLun")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StorageType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("wwid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Wwid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_gb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SizeGb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("multiprotocol_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MultiprotocolType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_volume").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageVolume")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_baremetalsolution_volumes_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_baremetalsolution_volumes.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBaremetalsolutionVolumeLunsGenerator) GetSubTables() []*schema.Table {
	return nil
}
