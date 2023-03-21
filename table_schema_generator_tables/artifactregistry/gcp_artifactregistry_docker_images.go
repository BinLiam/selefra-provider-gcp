package artifactregistry

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpArtifactregistryDockerImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpArtifactregistryDockerImagesGenerator{}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetTableName() string {
	return "gcp_artifactregistry_docker_images"
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListDockerImagesRequest{
				Parent: task.ParentRawResult.(*pb.Repository).Name,
			}
			gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListDockerImages(ctx, req, c.CallOptions...)
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

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_size_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ImageSizeBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BuildTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upload_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UploadTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("media_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MediaType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_artifactregistry_repositories_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_artifactregistry_repositories.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpArtifactregistryDockerImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
