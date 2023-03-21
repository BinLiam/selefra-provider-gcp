package logging

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	logging "cloud.google.com/go/logging/apiv2"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpLoggingMetricsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpLoggingMetricsGenerator{}

func (x *TableGcpLoggingMetricsGenerator) GetTableName() string {
	return "gcp_logging_metrics"
}

func (x *TableGcpLoggingMetricsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpLoggingMetricsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpLoggingMetricsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpLoggingMetricsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListLogMetricsRequest{
				Parent: "projects/" + c.ProjectId,
			}
			gcpClient, err := logging.NewMetricsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListLogMetrics(ctx, req, c.CallOptions...)
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

func (x *TableGcpLoggingMetricsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpLoggingMetricsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Version")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Filter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("label_extractors").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LabelExtractors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("bucket_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BucketOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metric_descriptor").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MetricDescriptor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("value_extractor").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ValueExtractor")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpLoggingMetricsGenerator) GetSubTables() []*schema.Table {
	return nil
}
