package cloudiot

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	iot "cloud.google.com/go/iot/apiv1"
	pb "cloud.google.com/go/iot/apiv1/iotpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"golang.org/x/sync/errgroup"
	"google.golang.org/api/iterator"
)

type TableGcpCloudiotDeviceRegistriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudiotDeviceRegistriesGenerator{}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetTableName() string {
	return "gcp_cloudiot_device_registries"
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := iot.NewDeviceManagerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			eg, gctx := errgroup.WithContext(ctx)
			for i := range locations {
				location := locations[i]
				eg.Go(func() error {
					req := &pb.ListDeviceRegistriesRequest{
						Parent: "projects/" + c.ProjectId + "/locations/" + location,
					}
					it := gcpClient.ListDeviceRegistries(gctx, req, c.CallOptions...)
					for {
						resp, err := it.Next()
						if err == iterator.Done {
							break
						}
						if err != nil {
							return err
						}

						resultChannel <- resp
					}
					return nil
				})
			}
			maybeError := eg.Wait()
			return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
		},
	}
}

var locations = []string{"us-central1", "europe-west1", "asia-east1"}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("mqtt_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MqttConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LogLevel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credentials").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Credentials")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_notification_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EventNotificationConfigs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_notification_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateNotificationConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HttpConfig")).Build(),
	}
}

func (x *TableGcpCloudiotDeviceRegistriesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpCloudiotDevicesGenerator{}),
	}
}
