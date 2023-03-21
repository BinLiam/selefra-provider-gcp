package cloudiot

import (
	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpCloudiotDevicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudiotDevicesGenerator{}

func (x *TableGcpCloudiotDevicesGenerator) GetTableName() string {
	return "gcp_cloudiot_devices"
}

func (x *TableGcpCloudiotDevicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudiotDevicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudiotDevicesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCloudiotDevicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			req := &pb.ListDevicesRequest{
				Parent: task.ParentRawResult.(*pb.DeviceRegistry).Name,
			}
			gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListDevices(ctx, req, c.CallOptions...)
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

func (x *TableGcpCloudiotDevicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudiotDevicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("blocked").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Blocked")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_error_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastErrorTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LogLevel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_state_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastStateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_config_send_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastConfigSendTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Config")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_cloudiot_device_registries_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_cloudiot_device_registries.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Credentials")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_event_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastEventTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_config_ack_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastConfigAckTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_error_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastErrorStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_heartbeat_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastHeartbeatTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GatewayConfig")).Build(),
	}
}

func (x *TableGcpCloudiotDevicesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpCloudiotDeviceConfigsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpCloudiotDeviceStatesGenerator{}),
	}
}
