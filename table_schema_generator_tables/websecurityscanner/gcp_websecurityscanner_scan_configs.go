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

type TableGcpWebsecurityscannerScanConfigsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpWebsecurityscannerScanConfigsGenerator{}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetTableName() string {
	return "gcp_websecurityscanner_scan_configs"
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := websecurityscanner.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListScanConfigs(ctx, &pb.ListScanConfigsRequest{
				Parent: "projects/" + c.ProjectId,
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

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("authentication").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Authentication")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_agent").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("UserAgent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("export_to_security_command_center").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ExportToSecurityCommandCenter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("starting_urls").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("StartingUrls")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("risk_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RiskLevel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("static_ip_scan").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("StaticIpScan")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_qps").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxQps")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Schedule")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_scan").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ManagedScan")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ignore_http_status_errors").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IgnoreHttpStatusErrors")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("blacklist_patterns").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("BlacklistPatterns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpWebsecurityscannerScanConfigsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpWebsecurityscannerScanConfigScanRunsGenerator{}),
	}
}
