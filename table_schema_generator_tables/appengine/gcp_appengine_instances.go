package appengine

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpAppengineInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineInstancesGenerator{}

func (x *TableGcpAppengineInstancesGenerator) GetTableName() string {
	return "gcp_appengine_instances"
}

func (x *TableGcpAppengineInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineInstancesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListInstancesRequest{
				Parent: task.ParentRawResult.(*pb.Version).Name,
			}
			gcpClient, err := appengine.NewInstancesClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListInstances(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("vm_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VmId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory_usage").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MemoryUsage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_appengine_versions_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_appengine_versions.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_engine_release").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AppEngineRelease")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VmName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("average_latency").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AverageLatency")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_ip").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VmIp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_liveness").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VmLiveness")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("errors").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Errors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("qps").ColumnType(schema.ColumnTypeFloat).
			Extractor(column_value_extractor.StructSelector("Qps")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VmStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_debug_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("VmDebugEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("availability").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Availability")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm_zone_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VmZoneName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requests").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Requests")).Build(),
	}
}

func (x *TableGcpAppengineInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
