package aiplatform

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/cloud/location"
)

type TableGcpAiplatformBatchPredictionJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformBatchPredictionJobsGenerator{}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetTableName() string {
	return "gcp_aiplatform_batch_prediction_jobs"
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListBatchPredictionJobsRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}

			if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewJobClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListBatchPredictionJobs(ctx, req, c.CallOptions...)
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

func filterLocation(locationId string) bool {
	return locationId == "us" || locationId == "eu"
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resources_consumed").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourcesConsumed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Model")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("partial_failures").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PartialFailures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_version_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ModelVersionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelParameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("explanation_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExplanationSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("manual_batch_tuning_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ManualBatchTuningParameters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generate_explanation").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("GenerateExplanation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutputConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dedicated_resources").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DedicatedResources")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("completion_stats").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CompletionStats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InputConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("unmanaged_container_model").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UnmanagedContainerModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutputInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_job_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_job_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpAiplatformBatchPredictionJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
