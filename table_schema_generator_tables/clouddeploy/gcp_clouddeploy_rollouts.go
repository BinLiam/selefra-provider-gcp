package clouddeploy

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	deploy "cloud.google.com/go/deploy/apiv1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpClouddeployRolloutsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouddeployRolloutsGenerator{}

func (x *TableGcpClouddeployRolloutsGenerator) GetTableName() string {
	return "gcp_clouddeploy_rollouts"
}

func (x *TableGcpClouddeployRolloutsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouddeployRolloutsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouddeployRolloutsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpClouddeployRolloutsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListRolloutsRequest{
				Parent: task.ParentRawResult.(*pb.Release).Name,
			}
			gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListRollouts(ctx, req, c.CallOptions...)
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

func (x *TableGcpClouddeployRolloutsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouddeployRolloutsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deploying_build").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeployingBuild")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deploy_failure_cause").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DeployFailureCause")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("phases").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Phases")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_clouddeploy_releases_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_clouddeploy_releases.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approve_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ApproveTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("approval_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ApprovalState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("TargetId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deploy_end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeployEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("FailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enqueue_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EnqueueTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deploy_start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeployStartTime")).Build(),
	}
}

func (x *TableGcpClouddeployRolloutsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpClouddeployJobRunsGenerator{}),
	}
}
