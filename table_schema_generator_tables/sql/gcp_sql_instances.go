package sql

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

type TableGcpSqlInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSqlInstancesGenerator{}

func (x *TableGcpSqlInstancesGenerator) GetTableName() string {
	return "gcp_sql_instances"
}

func (x *TableGcpSqlInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSqlInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSqlInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpSqlInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			sqlClient, err := sql.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for {
				output, err := sqlClient.Instances.List(c.ProjectId).MaxResults(1000).PageToken(nextPageToken).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Items
				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpSqlInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpSqlInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("database_installed_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseInstalledVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gce_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GceZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("root_password").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RootPassword")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backend_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BackendType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_addresses").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("IpAddresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ipv6_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Ipv6Address")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaintenanceVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReplicaConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ReplicaNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account_email_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccountEmailAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_encryption_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DiskEncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disk_encryption_status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DiskEncryptionStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_ca_cert").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServerCaCert")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_disk_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CurrentDiskSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DatabaseVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Settings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failover_replica").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FailoverReplica")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("master_instance_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MasterInstanceName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("on_premises_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OnPremisesConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InstanceType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_disk_size").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxDiskSize")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("out_of_disk_report").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutOfDiskReport")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("scheduled_maintenance").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ScheduledMaintenance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("available_maintenance_versions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("AvailableMaintenanceVersions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connection_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ConnectionName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Project")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_gce_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecondaryGceZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspension_reason").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SuspensionReason")).Build(),
	}
}

func (x *TableGcpSqlInstancesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpSqlUsersGenerator{}),
	}
}
