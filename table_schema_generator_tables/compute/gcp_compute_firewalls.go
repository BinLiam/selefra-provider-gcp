package compute

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpComputeFirewallsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeFirewallsGenerator{}

func (x *TableGcpComputeFirewallsGenerator) GetTableName() string {
	return "gcp_compute_firewalls"
}

func (x *TableGcpComputeFirewallsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeFirewallsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeFirewallsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeFirewallsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListFirewallsRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewFirewallsRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.List(ctx, req, c.CallOptions...)
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

func (x *TableGcpComputeFirewallsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeFirewallsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("direction").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Direction")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Allowed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("denied").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Denied")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_ranges").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SourceRanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_tags").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("TargetTags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destination_ranges").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("DestinationRanges")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_service_accounts").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SourceServiceAccounts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_tags").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SourceTags")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_service_accounts").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("TargetServiceAccounts")).Build(),
	}
}

func (x *TableGcpComputeFirewallsGenerator) GetSubTables() []*schema.Table {
	return nil
}
