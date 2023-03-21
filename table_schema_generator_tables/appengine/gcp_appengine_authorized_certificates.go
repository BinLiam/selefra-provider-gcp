package appengine

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpAppengineAuthorizedCertificatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineAuthorizedCertificatesGenerator{}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetTableName() string {
	return "gcp_appengine_authorized_certificates"
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListAuthorizedCertificatesRequest{
				Parent: "apps/" + c.ProjectId,
			}
			gcpClient, err := appengine.NewAuthorizedCertificatesClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListAuthorizedCertificates(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("visible_domain_mappings").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("VisibleDomainMappings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_mappings_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DomainMappingsCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_names").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("DomainNames")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_raw_data").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CertificateRawData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managed_certificate").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ManagedCertificate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpAppengineAuthorizedCertificatesGenerator) GetSubTables() []*schema.Table {
	return nil
}
