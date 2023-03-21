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

type TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator{}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetTableName() string {
	return "gcp_websecurityscanner_scan_config_scan_run_crawled_urls"
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.ScanRun)

			gcpClient, err := websecurityscanner.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListCrawledUrls(ctx, &pb.ListCrawledUrlsRequest{
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

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("scan_run_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_websecurityscanner_scan_config_scan_runs_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_websecurityscanner_scan_config_scan_runs.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_method").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("url").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Body")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunCrawledUrlsGenerator) GetSubTables() []*schema.Table {
	return nil
}
