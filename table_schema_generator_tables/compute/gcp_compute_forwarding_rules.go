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

type TableGcpComputeForwardingRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeForwardingRulesGenerator{}

func (x *TableGcpComputeForwardingRulesGenerator) GetTableName() string {
	return "gcp_compute_forwarding_rules"
}

func (x *TableGcpComputeForwardingRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeForwardingRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeForwardingRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListForwardingRulesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewForwardingRulesRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.ForwardingRules
			}
			return nil
		},
	}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeForwardingRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("is_mirroring_collector").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsMirroringCollector")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("no_automate_dns_zone").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("NoAutomateDnsZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Fingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("i_p_protocol").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPProtocol")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backend_service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BackendService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("i_p_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IPAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_global_access").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AllowGlobalAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("all_ports").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AllPorts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PortRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("psc_connection_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PscConnectionStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_directory_registrations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceDirectoryRegistrations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_label").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceLabel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnetwork").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Subnetwork")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancing_scheme").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LoadBalancingScheme")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata_filters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MetadataFilters")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ports").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Ports")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("psc_connection_id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PscConnectionId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_tier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkTier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Target")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpVersion")).Build(),
	}
}

func (x *TableGcpComputeForwardingRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
