package bigquery

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/bigquery/v2"
)

type TableGcpBigqueryTablesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBigqueryTablesGenerator{}

func (x *TableGcpBigqueryTablesGenerator) GetTableName() string {
	return "gcp_bigquery_tables"
}

func (x *TableGcpBigqueryTablesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBigqueryTablesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBigqueryTablesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBigqueryTablesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for {
				output, err := bigqueryService.Tables.List(c.ProjectId, task.ParentRawResult.(*bigquery.Dataset).DatasetReference.DatasetId).PageToken(nextPageToken).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				gcp_client.SendResults(resultChannel, output.Tables, func(result any) (any, error) {
					c := client.(*gcp_client.Client)
					bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
					if err != nil {
						return nil, err
					}
					item, err := bigqueryService.Tables.Get(c.ProjectId, task.ParentRawResult.(*bigquery.Dataset).DatasetReference.DatasetId, result.(*bigquery.TableListTables).TableReference.TableId).Do()
					if err != nil {
						return nil, err
					}
					return item, nil

				})
				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpBigqueryTablesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBigqueryTablesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("clone_definition").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CloneDefinition")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("table_reference").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TableReference")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_partitioning").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TimePartitioning")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("streaming_buffer").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StreamingBuffer")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_data_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExternalDataConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_staleness").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MaxStaleness")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_long_term_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumLongTermBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_active_physical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumActivePhysicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_long_term_logical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumLongTermLogicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("require_partition_filter").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("RequirePartitionFilter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("view").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("View")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("clustering").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Clustering")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("model").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Model")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_rows").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumRows")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("range_partitioning").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RangePartitioning")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("snapshot_definition").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SnapshotDefinition")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("friendly_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FriendlyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_bigquery_datasets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_bigquery_datasets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_long_term_physical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumLongTermPhysicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_partitions").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumPartitions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_physical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumPhysicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_active_logical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumActiveLogicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_time_travel_physical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumTimeTravelPhysicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schema").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Schema")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_collation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultCollation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("encryption_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expiration_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ExpirationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("materialized_view").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MaterializedView")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_total_logical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumTotalLogicalBytes")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_total_physical_bytes").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumTotalPhysicalBytes")).Build(),
	}
}

func (x *TableGcpBigqueryTablesGenerator) GetSubTables() []*schema.Table {
	return nil
}
