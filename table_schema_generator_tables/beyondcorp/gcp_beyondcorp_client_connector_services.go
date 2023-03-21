package beyondcorp

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	clientconnectorservices "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1"
	pb "cloud.google.com/go/beyondcorp/clientconnectorservices/apiv1/clientconnectorservicespb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBeyondcorpClientConnectorServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBeyondcorpClientConnectorServicesGenerator{}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetTableName() string {
	return "gcp_beyondcorp_client_connector_services"
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListClientConnectorServicesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := clientconnectorservices.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListClientConnectorServices(ctx, req, c.CallOptions...)
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

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ingress").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Ingress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("egress").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Egress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBeyondcorpClientConnectorServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
