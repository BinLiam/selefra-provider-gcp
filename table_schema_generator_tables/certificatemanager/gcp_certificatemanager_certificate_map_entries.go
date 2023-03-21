package certificatemanager

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
	pb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpCertificatemanagerCertificateMapEntriesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCertificatemanagerCertificateMapEntriesGenerator{}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetTableName() string {
	return "gcp_certificatemanager_certificate_map_entries"
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListCertificateMapEntriesRequest{
				Parent: task.ParentRawResult.(*pb.CertificateMap).Name,
			}
			gcpClient, err := certificatemanager.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListCertificateMapEntries(ctx, req, c.CallOptions...)
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

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_certificatemanager_certificate_maps_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_certificatemanager_certificate_maps.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificates").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("Certificates")).Build(),
	}
}

func (x *TableGcpCertificatemanagerCertificateMapEntriesGenerator) GetSubTables() []*schema.Table {
	return nil
}
