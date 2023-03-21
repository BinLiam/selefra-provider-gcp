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

type TableGcpComputeInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeInstancesGenerator{}

func (x *TableGcpComputeInstancesGenerator) GetTableName() string {
	return "gcp_compute_instances"
}

func (x *TableGcpComputeInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListInstancesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewInstancesRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.Instances
			}
			return nil
		},
	}
}

func (x *TableGcpComputeInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("start_restricted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StartRestricted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("advanced_machine_features").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AdvancedMachineFeatures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_suspended_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastSuspendedTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("machine_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MachineType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_machine_image").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceMachineImage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_machine_image_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceMachineImageEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disks").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Disks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_stop_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastStopTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_performance_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkPerformanceConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("params").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Params")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scheduling").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Scheduling")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shielded_instance_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ShieldedInstanceConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deletion_protection").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("DeletionProtection")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reservation_affinity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReservationAffinity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_device").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DisplayDevice")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_start_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastStartTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_cpu_platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MinCpuPlatform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_interfaces").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NetworkInterfaces")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("confidential_instance_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ConfidentialInstanceConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Fingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Zone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cpu_platform").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CpuPlatform")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_revocation_action_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("KeyRevocationActionType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guest_accelerators").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GuestAccelerators")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_policies").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ResourcePolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shielded_instance_integrity_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ShieldedInstanceIntegrityPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("can_ip_forward").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("CanIpForward")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hostname").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Hostname")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_ipv6_google_access").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateIpv6GoogleAccess")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_accounts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceAccounts")).Build(),
	}
}

func (x *TableGcpComputeInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
