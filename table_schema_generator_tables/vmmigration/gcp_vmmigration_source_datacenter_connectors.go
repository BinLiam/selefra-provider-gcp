package vmmigration

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	vmmigration "cloud.google.com/go/vmmigration/apiv1"
	pb "cloud.google.com/go/vmmigration/apiv1/vmmigrationpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpVmmigrationSourceDatacenterConnectorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVmmigrationSourceDatacenterConnectorsGenerator{}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetTableName() string {
	return "gcp_vmmigration_source_datacenter_connectors"
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.Source)

			gcpClient, err := vmmigration.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListDatacenterConnectors(ctx, &pb.ListDatacenterConnectorsRequest{
				Parent: parentItem.Name,
			}, c.CallOptions...)
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

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("available_versions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AvailableVersions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("registration_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RegistrationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Bucket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("upgrade_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpgradeStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Error")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("appliance_software_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplianceSoftwareVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vmmigration_sources_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vmmigration_sources.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("appliance_infrastructure_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ApplianceInfrastructureVersion")).Build(),
	}
}

func (x *TableGcpVmmigrationSourceDatacenterConnectorsGenerator) GetSubTables() []*schema.Table {
	return nil
}
