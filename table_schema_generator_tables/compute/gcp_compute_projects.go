package compute

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpComputeProjectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeProjectsGenerator{}

func (x *TableGcpComputeProjectsGenerator) GetTableName() string {
	return "gcp_compute_projects"
}

func (x *TableGcpComputeProjectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeProjectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeProjectsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeProjectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.GetProjectRequest{
				Project: c.ProjectId,
			}
			computeProjectsClient, err := compute.NewProjectsRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resp, err := computeProjectsClient.Get(ctx, req)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp
			return nil
		},
	}
}

func (x *TableGcpComputeProjectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeProjectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("quotas").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Quotas")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("usage_export_location").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UsageExportLocation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("common_instance_metadata").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CommonInstanceMetadata")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_network_tier").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultNetworkTier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enabled_features").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("EnabledFeatures")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("xpn_project_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("XpnProjectStatus")).Build(),
	}
}

func (x *TableGcpComputeProjectsGenerator) GetSubTables() []*schema.Table {
	return nil
}
