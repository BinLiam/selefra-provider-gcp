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

type TableGcpComputeTargetSslProxiesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeTargetSslProxiesGenerator{}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetTableName() string {
	return "gcp_compute_target_ssl_proxies"
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListTargetSslProxiesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewTargetSslProxiesRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.List(ctx, req, c.CallOptions...)
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

func (x *TableGcpComputeTargetSslProxiesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_map").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CertificateMap")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_policy").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SslPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("proxy_header").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ProxyHeader")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Service")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_certificates").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SslCertificates")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpComputeTargetSslProxiesGenerator) GetSubTables() []*schema.Table {
	return nil
}
