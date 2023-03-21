package dns

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/dns/v1"
)

type TableGcpDnsManagedZonesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpDnsManagedZonesGenerator{}

func (x *TableGcpDnsManagedZonesGenerator) GetTableName() string {
	return "gcp_dns_managed_zones"
}

func (x *TableGcpDnsManagedZonesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpDnsManagedZonesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpDnsManagedZonesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpDnsManagedZonesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			dnsClient, err := dns.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			for {
				output, err := dnsClient.ManagedZones.List(c.ProjectId).PageToken(nextPageToken).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.ManagedZones

				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpDnsManagedZonesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpDnsManagedZonesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_servers").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("NameServers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peering_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PeeringConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cloud_logging_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CloudLoggingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_visibility_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PrivateVisibilityConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_directory_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceDirectoryConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dnssec_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DnssecConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DnsName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("forwarding_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ForwardingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name_server_set").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NameServerSet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reverse_lookup_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ReverseLookupConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visibility").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Visibility")).Build(),
	}
}

func (x *TableGcpDnsManagedZonesGenerator) GetSubTables() []*schema.Table {
	return nil
}
