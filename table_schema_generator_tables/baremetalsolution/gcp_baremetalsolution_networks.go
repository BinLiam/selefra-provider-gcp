package baremetalsolution

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBaremetalsolutionNetworksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBaremetalsolutionNetworksGenerator{}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetTableName() string {
	return "gcp_baremetalsolution_networks"
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListNetworksRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListNetworks(ctx, req, c.CallOptions...)
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

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("vlan_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("VlanId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vrf").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Vrf")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Cidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reservations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Reservations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mac_address").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("MacAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("services_cidr").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServicesCidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBaremetalsolutionNetworksGenerator) GetSubTables() []*schema.Table {
	return nil
}
