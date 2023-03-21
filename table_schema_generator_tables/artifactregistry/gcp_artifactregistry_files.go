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

type TableGcpArtifactregistryFilesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpArtifactregistryFilesGenerator{}

func (x *TableGcpArtifactregistryFilesGenerator) GetTableName() string {
	return "gcp_artifactregistry_files"
}

func (x *TableGcpArtifactregistryFilesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpArtifactregistryFilesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpArtifactregistryFilesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpArtifactregistryFilesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListFilesRequest{
				Parent: task.ParentRawResult.(*pb.Repository).Name,
			}
			gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListFiles(ctx, req, c.CallOptions...)
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

func (x *TableGcpArtifactregistryFilesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpArtifactregistryFilesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("size_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SizeBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Owner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_artifactregistry_repositories_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_artifactregistry_repositories.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hashes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Hashes")).Build(),
	}
}

func (x *TableGcpArtifactregistryFilesGenerator) GetSubTables() []*schema.Table {
	return nil
}
