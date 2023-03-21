package websecurityscanner

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	websecurityscanner "cloud.google.com/go/websecurityscanner/apiv1"
	pb "cloud.google.com/go/websecurityscanner/apiv1/websecurityscannerpb"
	gax "github.com/googleapis/gax-go/v2"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator{}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetTableName() string {
	return "gcp_websecurityscanner_scan_config_scan_run_findings"
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*pb.ScanRun)

			gcpClient, err := websecurityscanner.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			findingTypes, err := getFindingTypes(ctx, gcpClient, parentItem.Name, c.CallOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for _, findingType := range findingTypes {
				it := gcpClient.ListFindings(ctx, &pb.ListFindingsRequest{
					Parent: parentItem.Name,
					Filter: fmt.Sprintf(`"finding_type="%s"`, findingType),
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
			}

			return nil
		},
	}
}

func getFindingTypes(ctx context.Context, gcpClient *websecurityscanner.Client, parentId string, callOptions ...gax.CallOption) ([]string, error) {
	response, err := gcpClient.ListFindingTypeStats(ctx, &pb.ListFindingTypeStatsRequest{
		Parent: parentId,
	}, callOptions...)

	if err != nil {
		return nil, err
	}

	findingTypes := make([]string, len(response.FindingTypeStats))

	for _, typestat := range response.FindingTypeStats {
		findingTypes = append(findingTypes, typestat.FindingType)
	}

	return findingTypes, nil
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("body").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Body")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("frame_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FrameUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("violating_resource").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ViolatingResource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vulnerable_headers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VulnerableHeaders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xss").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Xss")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reproduction_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ReproductionUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("final_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FinalUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("form").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Form")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("outdated_library").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("OutdatedLibrary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tracking_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TrackingId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xxe").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Xxe")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_websecurityscanner_scan_config_scan_runs_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_websecurityscanner_scan_config_scan_runs.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("finding_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FindingType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_method").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("HttpMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fuzzed_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FuzzedUrl")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("vulnerable_parameters").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VulnerableParameters")).Build(),
	}
}

func (x *TableGcpWebsecurityscannerScanConfigScanRunFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
