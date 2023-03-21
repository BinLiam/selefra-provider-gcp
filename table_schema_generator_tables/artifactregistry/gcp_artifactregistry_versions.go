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

type TableGcpArtifactregistryVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpArtifactregistryVersionsGenerator{}

func (x *TableGcpArtifactregistryVersionsGenerator) GetTableName() string {
	return "gcp_artifactregistry_versions"
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListVersionsRequest{
				Parent: task.ParentRawResult.(*pb.Package).Name,
			}
			gcpClient, err := artifactregistry.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListVersions(ctx, req, c.CallOptions...)
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

func (x *TableGcpArtifactregistryVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_artifactregistry_packages_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_artifactregistry_packages.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("related_tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RelatedTags")).Build(),
	}
}

func (x *TableGcpArtifactregistryVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
