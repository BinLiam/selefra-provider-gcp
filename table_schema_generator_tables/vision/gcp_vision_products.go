package vision

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	vision "cloud.google.com/go/vision/v2/apiv1"
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpVisionProductsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVisionProductsGenerator{}

func (x *TableGcpVisionProductsGenerator) GetTableName() string {
	return "gcp_vision_products"
}

func (x *TableGcpVisionProductsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVisionProductsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVisionProductsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVisionProductsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := vision.NewProductSearchClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListProducts(ctx, &pb.ListProductsRequest{
				Parent:   "projects/" + c.ProjectId + "/locations/" + c.Location,
				PageSize: 100,
			}, c.CallOptions...)
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

func (x *TableGcpVisionProductsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVisionProductsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_category").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProductCategory")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("product_labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ProductLabels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
	}
}

func (x *TableGcpVisionProductsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpVisionProductReferenceImagesGenerator{}),
	}
}
