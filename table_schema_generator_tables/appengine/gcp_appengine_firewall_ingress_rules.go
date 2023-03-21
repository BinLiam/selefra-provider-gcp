package appengine

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpAppengineFirewallIngressRulesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineFirewallIngressRulesGenerator{}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetTableName() string {
	return "gcp_appengine_firewall_ingress_rules"
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListIngressRulesRequest{
				Parent: "apps/" + c.ProjectId,
			}
			gcpClient, err := appengine.NewFirewallClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListIngressRules(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("priority").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Priority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("action").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Action")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("source_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SourceRange")).Build(),
	}
}

func (x *TableGcpAppengineFirewallIngressRulesGenerator) GetSubTables() []*schema.Table {
	return nil
}
