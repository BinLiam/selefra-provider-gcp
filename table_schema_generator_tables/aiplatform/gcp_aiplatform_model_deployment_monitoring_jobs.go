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

type TableGcpAiplatformModelDeploymentMonitoringJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformModelDeploymentMonitoringJobsGenerator{}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetTableName() string {
	return "gcp_aiplatform_model_deployment_monitoring_jobs"
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListModelDeploymentMonitoringJobsRequest{
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
			it := gcpClient.ListModelDeploymentMonitoringJobs(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("enable_monitoring_pipeline_logs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableMonitoringPipelineLogs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_job_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_job_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_deployment_monitoring_objective_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelDeploymentMonitoringObjectiveConfigs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predict_instance_schema_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PredictInstanceSchemaUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stats_anomalies_base_directory").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StatsAnomaliesBaseDirectory")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ScheduleState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_schedule_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NextScheduleTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("latest_monitoring_pipeline_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LatestMonitoringPipelineMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_sampling_strategy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoggingSamplingStrategy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sample_predict_instance").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SamplePredictInstance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("analysis_instance_schema_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AnalysisInstanceSchemaUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_ttl").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogTtl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_deployment_monitoring_schedule_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelDeploymentMonitoringScheduleConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_monitoring_alert_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelMonitoringAlertConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bigquery_tables").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BigqueryTables")).Build(),
	}
}

func (x *TableGcpAiplatformModelDeploymentMonitoringJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
