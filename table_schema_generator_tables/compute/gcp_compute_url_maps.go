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

type TableGcpComputeUrlMapsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeUrlMapsGenerator{}

func (x *TableGcpComputeUrlMapsGenerator) GetTableName() string {
	return "gcp_compute_url_maps"
}

func (x *TableGcpComputeUrlMapsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeUrlMapsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeUrlMapsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeUrlMapsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListUrlMapsRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewUrlMapsRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.AggregatedList(ctx, req, c.CallOptions...)
			for {
				resp, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- resp.Value.UrlMaps
			}
			return nil
		},
	}
}

func (x *TableGcpComputeUrlMapsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeUrlMapsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_url_redirect").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultUrlRedirect")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("header_action").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HeaderAction")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultService")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Fingerprint")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tests").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Tests")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_route_action").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultRouteAction")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host_rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("HostRules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("path_matchers").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PathMatchers")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
	}
}

func (x *TableGcpComputeUrlMapsGenerator) GetSubTables() []*schema.Table {
	return nil
}
