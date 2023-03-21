package monitoring

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	pb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpMonitoringAlertPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpMonitoringAlertPoliciesGenerator{}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetTableName() string {
	return "gcp_monitoring_alert_policies"
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListAlertPoliciesRequest{
				Name: "projects/" + c.ProjectId,
			}
			gcpClient, err := monitoring.NewAlertPolicyClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListAlertPolicies(ctx, req, c.CallOptions...)
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

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("combiner").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Combiner")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("validity").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Validity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Enabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_record").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreationRecord")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("documentation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Documentation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UserLabels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("conditions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Conditions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("notification_channels").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("NotificationChannels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mutation_record").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MutationRecord")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("alert_strategy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AlertStrategy")).Build(),
	}
}

func (x *TableGcpMonitoringAlertPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
