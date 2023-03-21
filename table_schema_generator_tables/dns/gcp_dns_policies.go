package dns

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/dns/v1"
)

type TableGcpDnsPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpDnsPoliciesGenerator{}

func (x *TableGcpDnsPoliciesGenerator) GetTableName() string {
	return "gcp_dns_policies"
}

func (x *TableGcpDnsPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpDnsPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpDnsPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpDnsPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			dnsClient, err := dns.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for {
				output, err := dnsClient.Policies.List(c.ProjectId).PageToken(nextPageToken).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Policies

				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpDnsPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpDnsPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("alternative_name_server_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AlternativeNameServerConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_logging").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableLogging")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enable_inbound_forwarding").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("EnableInboundForwarding")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("networks").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Networks")).Build(),
	}
}

func (x *TableGcpDnsPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
