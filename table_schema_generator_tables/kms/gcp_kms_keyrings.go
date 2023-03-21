package kms

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	kms "cloud.google.com/go/kms/apiv1"
	"cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

type TableGcpKmsKeyringsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsKeyringsGenerator{}

func (x *TableGcpKmsKeyringsGenerator) GetTableName() string {
	return "gcp_kms_keyrings"
}

func (x *TableGcpKmsKeyringsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsKeyringsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsKeyringsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsKeyringsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*locationpb.Location)
			kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := kmsClient.ListKeyRings(ctx, &kmspb.ListKeyRingsRequest{Parent: p.Name})
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

func (x *TableGcpKmsKeyringsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpKmsKeyringsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_kms_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_kms_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpKmsKeyringsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpKmsCryptoKeysGenerator{}),
		table_schema_generator.GenTableSchema(&TableGcpKmsImportJobsGenerator{}),
	}
}
