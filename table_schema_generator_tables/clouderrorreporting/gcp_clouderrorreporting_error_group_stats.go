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

type TableGcpClouderrorreportingErrorGroupStatsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouderrorreportingErrorGroupStatsGenerator{}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetTableName() string {
	return "gcp_clouderrorreporting_error_group_stats"
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListGroupStatsRequest{
				ProjectName: "projects/" + c.ProjectId,
			}
			gcpClient, err := errorreporting.NewErrorStatsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListGroupStats(ctx, req, c.CallOptions...)
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

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("affected_users_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("AffectedUsersCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("timed_counts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TimedCounts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_seen_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LastSeenTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("affected_services").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AffectedServices")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("representative").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Representative")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("group").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Group")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Count")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("first_seen_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FirstSeenTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("num_affected_services").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("NumAffectedServices")).Build(),
	}
}

func (x *TableGcpClouderrorreportingErrorGroupStatsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpClouderrorreportingErrorEventsGenerator{}),
	}
}
