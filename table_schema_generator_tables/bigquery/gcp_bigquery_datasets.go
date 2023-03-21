package bigquery

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/bigquery/v2"
)

type TableGcpBigqueryDatasetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBigqueryDatasetsGenerator{}

func (x *TableGcpBigqueryDatasetsGenerator) GetTableName() string {
	return "gcp_bigquery_datasets"
}

func (x *TableGcpBigqueryDatasetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBigqueryDatasetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBigqueryDatasetsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBigqueryDatasetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for {
				output, err := bigqueryService.Datasets.List(c.ProjectId).PageToken(nextPageToken).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				gcp_client.SendResults(resultChannel, output.Datasets, func(result any) (any, error) {
					c := client.(*gcp_client.Client)
					datasetListDataset := result.(*bigquery.DatasetListDatasets)
					bigqueryService, err := bigquery.NewService(ctx, c.ClientOptions...)
					if err != nil {
						return nil, err
					}
					item, err := bigqueryService.Datasets.Get(c.ProjectId, datasetListDataset.DatasetReference.DatasetId).Do()
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

func (x *TableGcpBigqueryDatasetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBigqueryDatasetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("access").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Access")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dataset_reference").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DatasetReference")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("storage_billing_model").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("StorageBillingModel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_collation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultCollation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_table_expiration_ms").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DefaultTableExpirationMs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_case_insensitive").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IsCaseInsensitive")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_time_travel_hours").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxTimeTravelHours")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_encryption_configuration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultEncryptionConfiguration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_partition_expiration_ms").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DefaultPartitionExpirationMs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("friendly_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FriendlyName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_modified_time").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("LastModifiedTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tags").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBigqueryDatasetsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpBigqueryTablesGenerator{}),
	}
}
