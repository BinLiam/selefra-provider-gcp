package websecurityscanner

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	websecurityscanner "cloud.google.com/go/websecurityscanner/apiv1"
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpWebsecurityscannerScanConfigScanRunsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpWebsecurityscannerScanConfigScanRunsGenerator{}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetTableName() string {
	return "gcp_websecurityscanner_scan_config_scan_runs"
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.ScanConfig)

			gcpClient, err := websecurityscanner.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListScanRuns(ctx, &pb.ListScanRunsRequest{
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

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("urls_tested_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("UrlsTestedCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("progress_percent").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProgressPercent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("warning_traces").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("WarningTraces")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_websecurityscanner_scan_configs_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_websecurityscanner_scan_configs.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("has_vulnerabilities").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("HasVulnerabilities")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("error_trace").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ErrorTrace")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("execution_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ExecutionState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("result_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ResultState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("urls_crawled_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("UrlsCrawledCount")).Build(),
	}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator{}),
	}
}
