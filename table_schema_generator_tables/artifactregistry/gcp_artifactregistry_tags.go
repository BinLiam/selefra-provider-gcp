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

type TableGcpArtifactregistryTagsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpArtifactregistryTagsGenerator{}

func (x *TableGcpArtifactregistryTagsGenerator) GetTableName() string {
	return "gcp_artifactregistry_tags"
}

func (x *TableGcpArtifactregistryTagsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpArtifactregistryTagsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpArtifactregistryTagsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpArtifactregistryTagsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListTagsRequest{
				Parent: task.ParentRawResult.(*pb.Package).Name,
			}
			gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListTags(ctx, req, c.CallOptions...)
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

func (x *TableGcpArtifactregistryTagsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpArtifactregistryTagsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_artifactregistry_packages_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_artifactregistry_packages.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpArtifactregistryTagsGenerator) GetSubTables() []*schema.Table {
	return nil
}
