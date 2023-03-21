package beyondcorp

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appgateways "cloud.google.com/go/beyondcorp/appgateways/apiv1"
	pb "cloud.google.com/go/beyondcorp/appgateways/apiv1/appgatewayspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBeyondcorpAppGatewaysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBeyondcorpAppGatewaysGenerator{}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetTableName() string {
	return "gcp_beyondcorp_app_gateways"
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListAppGatewaysRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := appgateways.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListAppGateways(ctx, req, c.CallOptions...)
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

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allocated_connections").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AllocatedConnections")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("HostType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpBeyondcorpAppGatewaysGenerator) GetSubTables() []*schema.Table {
	return nil
}
