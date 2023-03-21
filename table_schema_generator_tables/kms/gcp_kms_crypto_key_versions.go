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

type TableGcpKmsCryptoKeyVersionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsCryptoKeyVersionsGenerator{}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetTableName() string {
	return "gcp_kms_crypto_key_versions"
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*kmspb.CryptoKey)
			kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := kmsClient.ListCryptoKeyVersions(ctx, &kmspb.ListCryptoKeyVersionsRequest{Parent: p.Name})
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

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("destroy_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DestroyTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("reimport_eligible").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ReimportEligible")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destroy_event_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DestroyEventTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_job").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImportJob")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("external_protection_level_options").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExternalProtectionLevelOptions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("generate_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("GenerateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ImportTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_kms_crypto_keys_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_kms_crypto_keys.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("algorithm").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Algorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("attestation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Attestation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("import_failure_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ImportFailureReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("protection_level").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProtectionLevel")).Build(),
	}
}

func (x *TableGcpKmsCryptoKeyVersionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
