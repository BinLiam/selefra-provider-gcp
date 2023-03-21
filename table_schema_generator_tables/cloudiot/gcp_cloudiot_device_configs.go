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

type TableGcpCloudiotDeviceConfigsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudiotDeviceConfigsGenerator{}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetTableName() string {
	return "gcp_cloudiot_device_configs"
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListDeviceConfigVersionsRequest{
				Name: task.ParentRawResult.(*pb.Device).Name,
			}
			gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resp, err := gcpClient.ListDeviceConfigVersions(ctx, req, c.CallOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp.DeviceConfigs
			return nil
		},
	}
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_cloudiot_devices_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_cloudiot_devices.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CloudUpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device_ack_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeviceAckTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("binary_data").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("BinaryData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("device_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpCloudiotDeviceConfigsGenerator) GetSubTables() []*schema.Table {
	return nil
}
