package aiplatform

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	pb "google.golang.org/genproto/googleapis/cloud/location"
)

type TableGcpAiplatformIndexendpointLocationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformIndexendpointLocationsGenerator{}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetTableName() string {
	return "gcp_aiplatform_indexendpoint_locations"
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListLocationsRequest{
				Name: "projects/" + c.ProjectId,
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint("us-central1-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewIndexEndpointClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListLocations(ctx, req, c.CallOptions...)
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

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpAiplatformIndexendpointLocationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpAiplatformIndexEndpointsGenerator{}),
	}
}
