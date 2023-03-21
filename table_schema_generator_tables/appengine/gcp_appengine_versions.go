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

type TableGcpAppengineVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineVersionsGenerator{}

func (x *TableGcpAppengineVersionsGenerator) GetTableName() string {
	return "gcp_appengine_versions"
}

func (x *TableGcpAppengineVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineVersionsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListVersionsRequest{
				Parent: task.ParentRawResult.(*pb.Service).Name,
			}
			gcpClient, err := appengine.NewVersionsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListVersions(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("instance_class").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceClass")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("readiness_check").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReadinessCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_access_connector").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VpcAccessConnector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Runtime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vm").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Vm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_api_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RuntimeApiVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("handlers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Handlers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_handlers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorHandlers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("env_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EnvVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_channel").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RuntimeChannel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("beta_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BetaSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_usage_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DiskUsageBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("libraries").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Libraries")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_expiration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultExpiration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_check").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HealthCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entrypoint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Entrypoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_main_executable_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RuntimeMainExecutablePath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployment").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Deployment")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inbound_services").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("InboundServices")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_engine_apis").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AppEngineApis")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("env").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Env")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreatedBy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("api_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ApiConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_env_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BuildEnvVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_appengine_services_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_appengine_services.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threadsafe").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Threadsafe")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serving_status").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ServingStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("liveness_check").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LivenessCheck")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoints_api_service").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndpointsApiService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Zones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nobuild_files_regex").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NobuildFilesRegex")).Build(),
	}
}

func (x *TableGcpAppengineVersionsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpAppengineInstancesGenerator{}),
	}
}
