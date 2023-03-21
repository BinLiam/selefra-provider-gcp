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
)

type TableGcpKmsImportJobsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsImportJobsGenerator{}

func (x *TableGcpKmsImportJobsGenerator) GetTableName() string {
	return "gcp_kms_import_jobs"
}

func (x *TableGcpKmsImportJobsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsImportJobsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsImportJobsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsImportJobsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*kmspb.KeyRing)
			kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := kmsClient.ListImportJobs(ctx, &kmspb.ListImportJobsRequest{Parent: p.Name})
			for {
				key, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- key
			}
			return nil
		},
	}
}

func (x *TableGcpKmsImportJobsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpKmsImportJobsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("generate_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GenerateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attestation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attestation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_method").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ImportMethod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protection_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProtectionLevel")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_event_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExpireEventTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PublicKey")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_kms_keyrings_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_kms_keyrings.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
	}
}

func (x *TableGcpKmsImportJobsGenerator) GetSubTables() []*schema.Table {
	return nil
}
