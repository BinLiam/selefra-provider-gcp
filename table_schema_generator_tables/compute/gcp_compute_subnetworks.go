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

type TableGcpComputeSubnetworksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeSubnetworksGenerator{}

func (x *TableGcpComputeSubnetworksGenerator) GetTableName() string {
	return "gcp_compute_subnetworks"
}

func (x *TableGcpComputeSubnetworksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeSubnetworksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeSubnetworksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeSubnetworksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListSubnetworksRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewSubnetworksRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.Subnetworks
			}
			return nil
		},
	}
}

func (x *TableGcpComputeSubnetworksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeSubnetworksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_flow_logs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableFlowLogs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_access_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6AccessType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_cidr_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6CidrRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Role")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_ip_ranges").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecondaryIpRanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Fingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_ipv6_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InternalIpv6Prefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_cidr_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpCidrRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StackType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GatewayAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ip_google_access").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("PrivateIpGoogleAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ipv6_google_access").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateIpv6GoogleAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Purpose")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_ipv6_prefix").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExternalIpv6Prefix")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableGcpComputeSubnetworksGenerator) GetSubTables() []*schema.Table {
	return nil
}
