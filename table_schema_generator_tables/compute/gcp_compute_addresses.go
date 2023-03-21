package compute

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpComputeAddressesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeAddressesGenerator{}

func (x *TableGcpComputeAddressesGenerator) GetTableName() string {
	return "gcp_compute_addresses"
}

func (x *TableGcpComputeAddressesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeAddressesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeAddressesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeAddressesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListAddressesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewAddressesRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- resp.Value.Addresses
			}
			return nil
		},
	}
}

func (x *TableGcpComputeAddressesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeAddressesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnetwork").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Subnetwork")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Users")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Purpose")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_tier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkTier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AddressType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Address")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_endpoint_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6EndpointType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("prefix_length").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PrefixLength")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
	}
}

func (x *TableGcpComputeAddressesGenerator) GetSubTables() []*schema.Table {
	return nil
}
