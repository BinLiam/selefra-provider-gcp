package cloudscheduler

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	scheduler "cloud.google.com/go/scheduler/apiv1"
	"cloud.google.com/go/scheduler/apiv1/schedulerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	locationspb "google.golang.org/api/cloudscheduler/v1"
	"google.golang.org/api/iterator"
)

type TableGcpCloudschedulerJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudschedulerJobsGenerator{}

func (x *TableGcpCloudschedulerJobsGenerator) GetTableName() string {
	return "gcp_cloudscheduler_jobs"
}

func (x *TableGcpCloudschedulerJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudschedulerJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudschedulerJobsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCloudschedulerJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*locationspb.Location)
			gcpClient, err := scheduler.NewCloudSchedulerClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListJobs(ctx, &schedulerpb.ListJobsRequest{Parent: p.Name})
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

func (x *TableGcpCloudschedulerJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudschedulerJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("pubsub_target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PubsubTarget")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("retry_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RetryConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Status")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attempt_deadline").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AttemptDeadline")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("http_target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HttpTarget")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("last_attempt_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LastAttemptTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_cloudscheduler_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_cloudscheduler_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Schedule")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("schedule_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ScheduleTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("time_zone").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TimeZone")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user_update_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UserUpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app_engine_http_target").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AppEngineHttpTarget")).Build(),
	}
}

func (x *TableGcpCloudschedulerJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
