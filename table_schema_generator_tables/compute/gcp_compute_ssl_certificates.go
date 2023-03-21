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

type TableGcpComputeSslCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeSslCertificatesGenerator{}

func (x *TableGcpComputeSslCertificatesGenerator) GetTableName() string {
	return "gcp_compute_ssl_certificates"
}

func (x *TableGcpComputeSslCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeSslCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeSslCertificatesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeSslCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.AggregatedListSslCertificatesRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewSslCertificatesRESTClient(ctx, c.ClientOptions...)
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

				resultChannel <- resp.Value.SslCertificates
			}
			return nil
		},
	}
}

func (x *TableGcpComputeSslCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeSslCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("certificate").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Certificate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Region")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subject_alternative_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("SubjectAlternativeNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Managed")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_key").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PrivateKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_managed").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SelfManaged")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpComputeSslCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
