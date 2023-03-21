package beyondcorp

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	clientgateways "cloud.google.com/go/beyondcorp/clientgateways/apiv1"
	pb "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBeyondcorpClientGatewaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBeyondcorpClientGatewaysGenerator{}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetTableName() string {
	return "gcp_beyondcorp_client_gateways"
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListClientGatewaysRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := clientgateways.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListClientGateways(ctx, req, c.CallOptions...)
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

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("client_connector_service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClientConnectorService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
	}
}

func (x *TableGcpBeyondcorpClientGatewaysGenerator) GetSubTables() []*schema.Table {
	return nil
}
