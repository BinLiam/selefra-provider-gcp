package billing

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	pb "cloud.google.com/go/billing/apiv1/billingpb"
	budgets "cloud.google.com/go/billing/budgets/apiv1"
	"cloud.google.com/go/billing/budgets/apiv1/budgetspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBillingBudgetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBillingBudgetsGenerator{}

func (x *TableGcpBillingBudgetsGenerator) GetTableName() string {
	return "gcp_billing_budgets"
}

func (x *TableGcpBillingBudgetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBillingBudgetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBillingBudgetsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBillingBudgetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &budgetspb.ListBudgetsRequest{
				Parent: task.ParentRawResult.(*pb.BillingAccount).Name,
			}
			gcpClient, err := budgets.NewBudgetClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListBudgets(ctx, req, c.CallOptions...)
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

func (x *TableGcpBillingBudgetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpBillingBudgetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("notifications_rule").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NotificationsRule")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_billing_billing_accounts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_billing_billing_accounts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("budget_filter").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BudgetFilter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("threshold_rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ThresholdRules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("amount").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Amount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
	}
}

func (x *TableGcpBillingBudgetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
