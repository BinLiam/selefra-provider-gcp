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

type TableGcpAiplatformIndexEndpointsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformIndexEndpointsGenerator{}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetTableName() string {
	return "gcp_aiplatform_index_endpoints"
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListIndexEndpointsRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}
			if filterIndexesLocations(task) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewIndexEndpointClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListIndexEndpoints(ctx, req, c.CallOptions...)
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

func filterIndexesLocations(task *schema.DataSourcePullTask) bool {
	if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
		return true
	}

	locationId := task.ParentRawResult.(*location.Location).LocationId

	return locationId == "europe-west4"
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_indexendpoint_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_indexendpoint_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_private_service_connect").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnablePrivateServiceConnect")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deployed_indexes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeployedIndexes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
	}
}

func (x *TableGcpAiplatformIndexEndpointsGenerator) GetSubTables() []*schema.Table {
	return nil
}
