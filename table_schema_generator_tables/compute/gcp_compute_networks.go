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

type TableGcpComputeNetworksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeNetworksGenerator{}

func (x *TableGcpComputeNetworksGenerator) GetTableName() string {
	return "gcp_compute_networks"
}

func (x *TableGcpComputeNetworksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeNetworksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeNetworksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeNetworksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListNetworksRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewNetworksRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.List(ctx, req, c.CallOptions...)
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

func (x *TableGcpComputeNetworksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeNetworksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("enable_ula_internal_ipv6").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableUlaInternalIpv6")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peerings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Peerings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gateway_i_pv4").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GatewayIPv4")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link_with_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLinkWithId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnetworks").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Subnetworks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auto_create_subnetworks").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AutoCreateSubnetworks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_firewall_policy_enforcement_order").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkFirewallPolicyEnforcementOrder")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("routing_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RoutingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("i_pv4_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPv4Range")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("firewall_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FirewallPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_ipv6_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InternalIpv6Range")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mtu").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Mtu")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpComputeNetworksGenerator) GetSubTables() []*schema.Table {
	return nil
}
