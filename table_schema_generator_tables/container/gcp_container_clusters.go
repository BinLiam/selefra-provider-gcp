package container

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpContainerClustersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpContainerClustersGenerator{}

func (x *TableGcpContainerClustersGenerator) GetTableName() string {
	return "gcp_container_clusters"
}

func (x *TableGcpContainerClustersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpContainerClustersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpContainerClustersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpContainerClustersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &containerpb.ListClustersRequest{
				Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
			}
			containerClient, err := container.NewClusterManagerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			output, err := containerClient.ListClusters(ctx, req)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.Clusters
			return nil
		},
	}
}

func (x *TableGcpContainerClustersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpContainerClustersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring_service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MonitoringService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("confidential_nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfidentialNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Zone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_tpu").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableTpu")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceLabels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mesh_certificates").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MeshCertificates")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("workload_identity_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("WorkloadIdentityConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("initial_node_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("InitialNodeCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnetwork").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Subnetwork")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authenticator_groups_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AuthenticatorGroupsConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tpu_ipv4_cidr_block").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TpuIpv4CidrBlock")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_allocation_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IpAllocationPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cost_management_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CostManagementConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("autopilot").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Autopilot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaintenancePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("autoscaling").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Autoscaling")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_ipv4_cidr_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NodeIpv4CidrSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LoggingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_node_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CurrentNodeVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_pool_defaults").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodePoolDefaults")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_pool_auto_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodePoolAutoConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_kubernetes_alpha").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableKubernetesAlpha")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_usage_export_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceUsageExportConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Endpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("initial_cluster_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InitialClusterVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logging_service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LoggingService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shielded_nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ShieldedNodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NotificationConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("monitoring_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MonitoringConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("legacy_abac").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LegacyAbac")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vertical_pod_autoscaling").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VerticalPodAutoscaling")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodeConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_service_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IdentityServiceConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_master_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CurrentMasterVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("services_ipv4_cidr").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServicesIpv4Cidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cluster_ipv4_cidr").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ClusterIpv4Cidr")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("node_pools").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NodePools")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_cluster_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateClusterConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_encryption").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DatabaseEncryption")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_group_urls").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("InstanceGroupUrls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_auth").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MasterAuth")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("locations").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Locations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("binary_authorization").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BinaryAuthorization")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addons_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AddonsConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_authorized_networks_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MasterAuthorizedNetworksConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_max_pods_constraint").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultMaxPodsConstraint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("release_channel").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReleaseChannel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_node_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CurrentNodeCount")).Build(),
	}
}

func (x *TableGcpContainerClustersGenerator) GetSubTables() []*schema.Table {
	return nil
}
