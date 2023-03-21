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

type TableGcpAiplatformDataLabelingJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformDataLabelingJobsGenerator{}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetTableName() string {
	return "gcp_aiplatform_data_labeling_jobs"
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListDataLabelingJobsRequest{
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
			it := gcpClient.ListDataLabelingJobs(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labeler_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LabelerCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inputs_schema_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InputsSchemaUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instruction_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstructionUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_spend").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CurrentSpend")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("specialist_pools").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SpecialistPools")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_learning_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ActiveLearningConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("datasets").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Datasets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inputs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Inputs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotation_labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AnnotationLabels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labeling_progress").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LabelingProgress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_job_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_job_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
	}
}

func (x *TableGcpAiplatformDataLabelingJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
