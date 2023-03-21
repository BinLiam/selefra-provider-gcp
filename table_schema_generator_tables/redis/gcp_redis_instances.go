package redis

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	redis "cloud.google.com/go/redis/apiv1"
	pb "cloud.google.com/go/redis/apiv1/redispb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpRedisInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpRedisInstancesGenerator{}

func (x *TableGcpRedisInstancesGenerator) GetTableName() string {
	return "gcp_redis_instances"
}

func (x *TableGcpRedisInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpRedisInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpRedisInstancesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpRedisInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListInstancesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := redis.NewCloudRedisClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListInstances(ctx, req, c.CallOptions...)
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

func (x *TableGcpRedisInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpRedisInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Port")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("server_ca_certs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServerCaCerts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("transit_encryption_mode").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("TransitEncryptionMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_endpoint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReadEndpoint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Tier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AuthEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_replicas_mode").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ReadReplicasMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replica_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ReplicaCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current_location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CurrentLocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reserved_ip_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReservedIpRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Host")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connect_mode").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ConnectMode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redis_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("RedisVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaintenancePolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("secondary_ip_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SecondaryIpRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("redis_configs").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RedisConfigs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("memory_size_gb").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MemorySizeGb")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nodes").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Nodes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("read_endpoint_port").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ReadEndpointPort")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status_message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StatusMessage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("authorized_network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthorizedNetwork")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("persistence_iam_identity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PersistenceIamIdentity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance_schedule").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaintenanceSchedule")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alternative_location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AlternativeLocationId")).Build(),
	}
}

func (x *TableGcpRedisInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
