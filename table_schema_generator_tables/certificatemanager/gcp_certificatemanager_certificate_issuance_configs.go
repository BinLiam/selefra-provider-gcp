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

type TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator{}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetTableName() string {
	return "gcp_certificatemanager_certificate_issuance_configs"
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListCertificateIssuanceConfigsRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := certificatemanager.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListCertificateIssuanceConfigs(ctx, req, c.CallOptions...)
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

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_window_percentage").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RotationWindowPercentage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("KeyAlgorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("lifetime").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Lifetime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("certificate_authority_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CertificateAuthorityConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpCertificatemanagerCertificateIssuanceConfigsGenerator) GetSubTables() []*schema.Table {
	return nil
}
