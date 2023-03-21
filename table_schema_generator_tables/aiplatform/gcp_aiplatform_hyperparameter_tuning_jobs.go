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

type TableGcpAiplatformHyperparameterTuningJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformHyperparameterTuningJobsGenerator{}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetTableName() string {
	return "gcp_aiplatform_hyperparameter_tuning_jobs"
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListHyperparameterTuningJobsRequest{
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
			it := gcpClient.ListHyperparameterTuningJobs(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("trials").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Trials")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("study_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StudySpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_trial_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxTrialCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("trial_job_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TrialJobSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parallel_trial_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ParallelTrialCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_job_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_job_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_failed_trial_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxFailedTrialCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
	}
}

func (x *TableGcpAiplatformHyperparameterTuningJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
