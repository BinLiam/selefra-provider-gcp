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

type TableGcpComputeBackendServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeBackendServicesGenerator{}

func (x *TableGcpComputeBackendServicesGenerator) GetTableName() string {
	return "gcp_compute_backend_services"
}

func (x *TableGcpComputeBackendServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeBackendServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeBackendServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeBackendServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListBackendServicesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewBackendServicesRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.BackendServices
			}
			return nil
		},
	}
}

func (x *TableGcpComputeBackendServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeBackendServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("circuit_breakers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CircuitBreakers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("consistent_hash").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConsistentHash")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_request_headers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CustomRequestHeaders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("edge_security_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("EdgeSecurityPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iap").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Iap")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_draining").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConnectionDraining")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failover_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FailoverPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("session_affinity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SessionAffinity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backends").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Backends")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_c_d_n").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableCDN")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_stream_duration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaxStreamDuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protocol").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Protocol")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_tracking_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConnectionTrackingPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("custom_response_headers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CustomResponseHeaders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("port_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PortName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timeout_sec").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TimeoutSec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locality_lb_policies").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LocalityLbPolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locality_lb_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocalityLbPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecurityPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecuritySettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("compression_mode").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CompressionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Fingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("health_checks").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("HealthChecks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outlier_detection").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutlierDetection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subsetting").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Subsetting")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_bindings").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ServiceBindings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("affinity_cookie_ttl_sec").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AffinityCookieTtlSec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cdn_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CdnPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("load_balancing_scheme").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LoadBalancingScheme")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
	}
}

func (x *TableGcpComputeBackendServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
