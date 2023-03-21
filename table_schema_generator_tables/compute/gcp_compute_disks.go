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

type TableGcpComputeDisksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeDisksGenerator{}

func (x *TableGcpComputeDisksGenerator) GetTableName() string {
	return "gcp_compute_disks"
}

func (x *TableGcpComputeDisksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeDisksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeDisksGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeDisksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListDisksRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewDisksRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.Disks
			}
			return nil
		},
	}
}

func (x *TableGcpComputeDisksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeDisksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("provisioned_iops").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProvisionedIops")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Architecture")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_attach_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastAttachTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_hint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocationHint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_zones").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ReplicaZones")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_disk_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDiskId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("size_gb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("SizeGb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_disk").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDisk")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceImage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceSnapshotId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("users").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Users")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceImageId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceSnapshot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("physical_block_size_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PhysicalBlockSizeBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceImageEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("licenses").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Licenses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource_policies").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ResourcePolicies")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceSnapshotEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_storage_object").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceStorageObject")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Zone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guest_os_features").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GuestOsFeatures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_detach_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastDetachTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_codes").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("LicenseCodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("options").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Options")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("params").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Params")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DiskEncryptionKey")).Build(),
	}
}

func (x *TableGcpComputeDisksGenerator) GetSubTables() []*schema.Table {
	return nil
}
