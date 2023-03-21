package clouderrorreporting

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
	pb "cloud.google.com/go/errorreporting/apiv1beta1/errorreportingpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpClouderrorreportingErrorEventsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouderrorreportingErrorEventsGenerator{}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetTableName() string {
	return "gcp_clouderrorreporting_error_events"
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListEventsRequest{
				ProjectName: "projects/" + c.ProjectId, GroupId: task.ParentRawResult.(*pb.ErrorGroupStats).Group.GroupId,
			}
			gcpClient, err := errorreporting.NewErrorStatsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListEvents(ctx, req, c.CallOptions...)
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

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("service_context").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ServiceContext")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("message").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Message")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("context").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Context")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_clouderrorreporting_error_group_stats_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_clouderrorreporting_error_group_stats.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("event_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EventTime")).Build(),
	}
}

func (x *TableGcpClouderrorreportingErrorEventsGenerator) GetSubTables() []*schema.Table {
	return nil
}
