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

type TableGcpAiplatformTensorboardsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformTensorboardsGenerator{}

func (x *TableGcpAiplatformTensorboardsGenerator) GetTableName() string {
	return "gcp_aiplatform_tensorboards"
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListTensorboardsRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}
			if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewTensorboardClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListTensorboards(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformTensorboardsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_tensorboard_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_tensorboard_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionSpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blob_storage_path_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BlobStoragePathPrefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("run_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RunCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
	}
}

func (x *TableGcpAiplatformTensorboardsGenerator) GetSubTables() []*schema.Table {
	return nil
}
