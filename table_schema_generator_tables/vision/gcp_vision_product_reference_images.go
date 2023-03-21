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

type TableGcpVisionProductReferenceImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVisionProductReferenceImagesGenerator{}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetTableName() string {
	return "gcp_vision_product_reference_images"
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.Product)

			gcpClient, err := vision.NewProductSearchClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListReferenceImages(ctx, &pb.ListReferenceImagesRequest{
				Parent:   parentItem.Name,
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

func (x *TableGcpVisionProductReferenceImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("bounding_polys").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BoundingPolys")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vision_products_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vision_products.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uri")).Build(),
	}
}

func (x *TableGcpVisionProductReferenceImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
