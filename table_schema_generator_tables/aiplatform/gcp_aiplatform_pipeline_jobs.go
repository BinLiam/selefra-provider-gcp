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

type TableGcpAiplatformPipelineJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformPipelineJobsGenerator{}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetTableName() string {
	return "gcp_aiplatform_pipeline_jobs"
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListPipelineJobsRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}
			if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewPipelineClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListPipelineJobs(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformPipelineJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("job_detail").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("JobDetail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TemplateMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pipeline_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PipelineSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("runtime_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RuntimeConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("template_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TemplateUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_pipeline_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_pipeline_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpAiplatformPipelineJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
