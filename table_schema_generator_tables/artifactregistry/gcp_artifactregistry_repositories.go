package artifactregistry

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	v1cloud "cloud.google.com/go/artifactregistry/apiv1"
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	v1google "google.golang.org/api/artifactregistry/v1"
	"google.golang.org/api/iterator"
)

type TableGcpArtifactregistryRepositoriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpArtifactregistryRepositoriesGenerator{}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetTableName() string {
	return "gcp_artifactregistry_repositories"
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListRepositoriesRequest{
				Parent: task.ParentRawResult.(*v1google.Location).Name,
			}
			gcpClient, err := v1cloud.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListRepositories(ctx, req, c.CallOptions...)
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

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("format").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Format")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kms_key_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KmsKeyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_artifactregistry_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_artifactregistry_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpArtifactregistryRepositoriesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpArtifactregistryDockerImagesGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpArtifactregistryFilesGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpArtifactregistryPackagesGenerator{}),
	}
}
