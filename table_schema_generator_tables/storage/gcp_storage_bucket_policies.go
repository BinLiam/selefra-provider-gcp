package storage

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"cloud.google.com/go/storage"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpStorageBucketPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpStorageBucketPoliciesGenerator{}

func (x *TableGcpStorageBucketPoliciesGenerator) GetTableName() string {
	return "gcp_storage_bucket_policies"
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			bkt := task.ParentRow.GetOrDefault("name", nil).(*string)
			storageClient, err := storage.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			output, err := storageClient.Bucket(*bkt).IAM().V3().Policy(ctx)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_storage_buckets_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_storage_buckets.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bindings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Bindings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpStorageBucketPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
