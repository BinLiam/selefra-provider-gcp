package bigtableadmin

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"cloud.google.com/go/bigtable"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBigtableadminBackupsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBigtableadminBackupsGenerator{}

func (x *TableGcpBigtableadminBackupsGenerator) GetTableName() string {
	return "gcp_bigtableadmin_backups"
}

func (x *TableGcpBigtableadminBackupsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBigtableadminBackupsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBigtableadminBackupsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBigtableadminBackupsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			instance := task.ParentTask.ParentRawResult.(*bigtable.InstanceInfo)
			cluster := task.ParentRawResult.(*bigtable.ClusterInfo)
			gcpClient, err := bigtable.NewAdminClient(ctx, c.ProjectId, instance.Name, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.Backups(ctx, cluster.Name)
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

func (x *TableGcpBigtableadminBackupsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBigtableadminBackupsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("source_table").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceTable")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_bigtableadmin_clusters_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_bigtableadmin_clusters.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("backup").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Backup")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
	}
}

func (x *TableGcpBigtableadminBackupsGenerator) GetSubTables() []*schema.Table {
	return nil
}
