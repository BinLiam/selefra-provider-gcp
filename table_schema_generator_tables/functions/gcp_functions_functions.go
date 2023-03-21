package functions

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	functions "cloud.google.com/go/functions/apiv1"
	pb "cloud.google.com/go/functions/apiv1/functionspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpFunctionsFunctionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpFunctionsFunctionsGenerator{}

func (x *TableGcpFunctionsFunctionsGenerator) GetTableName() string {
	return "gcp_functions_functions"
}

func (x *TableGcpFunctionsFunctionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpFunctionsFunctionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpFunctionsFunctionsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpFunctionsFunctionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListFunctionsRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := functions.NewCloudFunctionsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListFunctions(ctx, req, c.CallOptions...)
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

func (x *TableGcpFunctionsFunctionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpFunctionsFunctionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("available_memory_mb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AvailableMemoryMb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BuildId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Runtime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BuildName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("docker_repository").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DockerRepository")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_token").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceToken")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_instances").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinInstances")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_environment_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecretEnvironmentVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Timeout")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VersionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingress_settings").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("IngressSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entry_point").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EntryPoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_connector_egress_settings").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("VpcConnectorEgressSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_instances").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxInstances")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vpc_connector").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VpcConnector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account_email").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccountEmail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("environment_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EnvironmentVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_worker_pool").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BuildWorkerPool")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secret_volumes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecretVolumes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("docker_registry").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DockerRegistry")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_environment_variables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BuildEnvironmentVariables")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyName")).Build(),
	}
}

func (x *TableGcpFunctionsFunctionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
