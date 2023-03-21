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

type TableGcpClouddeployReleasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpClouddeployReleasesGenerator{}

func (x *TableGcpClouddeployReleasesGenerator) GetTableName() string {
	return "gcp_clouddeploy_releases"
}

func (x *TableGcpClouddeployReleasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpClouddeployReleasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpClouddeployReleasesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpClouddeployReleasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListReleasesRequest{
				Parent: task.ParentRawResult.(*pb.DeliveryPipeline).Name,
			}
			gcpClient, err := deploy.NewCloudDeployClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListReleases(ctx, req, c.CallOptions...)
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

func (x *TableGcpClouddeployReleasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpClouddeployReleasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("render_start_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RenderStartTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("render_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RenderState")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_clouddeploy_delivery_pipelines_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_clouddeploy_delivery_pipelines.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("abandoned").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Abandoned")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("skaffold_config_path").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SkaffoldConfigPath")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delivery_pipeline_snapshot").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeliveryPipelineSnapshot")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_snapshots").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetSnapshots")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("skaffold_version").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SkaffoldVersion")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_artifacts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetArtifacts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("render_end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("RenderEndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("skaffold_config_uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SkaffoldConfigUri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_artifacts").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("BuildArtifacts")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("target_renders").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("TargetRenders")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpClouddeployReleasesGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpClouddeployRolloutsGenerator{}),
	}
}
