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

type TableGcpAiplatformModelsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformModelsGenerator{}

func (x *TableGcpAiplatformModelsGenerator) GetTableName() string {
	return "gcp_aiplatform_models"
}

func (x *TableGcpAiplatformModelsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformModelsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformModelsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformModelsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListModelsRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}
			if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewModelClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListModels(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformModelsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformModelsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VersionCreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("artifact_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ArtifactUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_output_storage_formats").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedOutputStorageFormats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("predict_schemata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PredictSchemata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("training_pipeline").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrainingPipeline")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_input_storage_formats").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SupportedInputStorageFormats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VersionDescription")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployed_models").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeployedModels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("explanation_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExplanationSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model_source_info").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ModelSourceInfo")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_aliases").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VersionAliases")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_schema_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MetadataSchemaUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_export_formats").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SupportedExportFormats")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VersionUpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("container_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ContainerSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_deployment_resources_types").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("SupportedDeploymentResourcesTypes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_artifact").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MetadataArtifact")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_model_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_model_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableGcpAiplatformModelsGenerator) GetSubTables() []*schema.Table {
	return nil
}
