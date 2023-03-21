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

type TableGcpComputeImagesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeImagesGenerator{}

func (x *TableGcpComputeImagesGenerator) GetTableName() string {
	return "gcp_compute_images"
}

func (x *TableGcpComputeImagesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeImagesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeImagesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeImagesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListImagesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewImagesRESTClient(ctx, c.ClientOptions...)
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

func (x *TableGcpComputeImagesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeImagesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("source_disk_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceDiskEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceSnapshotEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deprecated").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Deprecated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("raw_disk").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RawDisk")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("shielded_instance_initial_state").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ShieldedInstanceInitialState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_disk").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDisk")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SourceImageEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LabelFingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("license_codes").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("LicenseCodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_disk_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceDiskId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceImage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("family").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Family")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_locations").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("StorageLocations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceSnapshot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("architecture").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Architecture")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archive_size_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ArchiveSizeBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("guest_os_features").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GuestOsFeatures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_size_gb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DiskSizeGb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_image_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceImageId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_snapshot_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceSnapshotId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("image_encryption_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImageEncryptionKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("licenses").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Licenses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
	}
}

func (x *TableGcpComputeImagesGenerator) GetSubTables() []*schema.Table {
	return nil
}
