package cloudiot

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpCloudiotDeviceStatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudiotDeviceStatesGenerator{}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetTableName() string {
	return "gcp_cloudiot_device_states"
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListDeviceStatesRequest{
				Name: task.ParentRawResult.(*pb.Device).Name,
			}
			gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resp, err := gcpClient.ListDeviceStates(ctx, req, c.CallOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.DeviceStates
			return nil
		},
	}
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("binary_data").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("BinaryData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_cloudiot_devices_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_cloudiot_devices.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpCloudiotDeviceStatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
