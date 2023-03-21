package cloudsupport

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	cloudsupport "google.golang.org/api/cloudsupport/v2beta"
)

type TableGcpCloudsupportCasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudsupportCasesGenerator{}

func (x *TableGcpCloudsupportCasesGenerator) GetTableName() string {
	return "gcp_cloudsupport_cases"
}

func (x *TableGcpCloudsupportCasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudsupportCasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudsupportCasesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCloudsupportCasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			gcpClient, err := cloudsupport.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for {
				output, err := gcpClient.Cases.List("projects/" + c.ProjectId).PageSize(1000).PageToken(nextPageToken).Context(ctx).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Cases
				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpCloudsupportCasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudsupportCasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("test_case").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("TestCase")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creator").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Creator")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subscriber_email_addresses").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubscriberEmailAddresses")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TimeZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("severity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Severity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("escalated").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Escalated")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("language_code").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LanguageCode")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("classification").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Classification")).Build(),
	}
}

func (x *TableGcpCloudsupportCasesGenerator) GetSubTables() []*schema.Table {
	return nil
}
