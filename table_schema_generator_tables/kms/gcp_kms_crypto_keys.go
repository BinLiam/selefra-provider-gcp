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

type TableGcpKmsCryptoKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpKmsCryptoKeysGenerator{}

func (x *TableGcpKmsCryptoKeysGenerator) GetTableName() string {
	return "gcp_kms_crypto_keys"
}

func (x *TableGcpKmsCryptoKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpKmsCryptoKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpKmsCryptoKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpKmsCryptoKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*kmspb.KeyRing)
			kmsClient, err := kms.NewKeyManagementClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := kmsClient.ListCryptoKeys(ctx, &kmspb.ListCryptoKeysRequest{Parent: p.Name})
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

func (x *TableGcpKmsCryptoKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpKmsCryptoKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("import_only").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("ImportOnly")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destroy_scheduled_duration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DestroyScheduledDuration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation_period").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.WrapperExtractFunction(func(ctx context.Context, clientMeta *schema.ClientMeta, client any,
				task *schema.DataSourcePullTask, row *schema.Row, column *schema.Column, result any) (any, *schema.Diagnostics) {

				extractor := func() (any, error) {
					return result.(*kmspb.CryptoKey).GetRotationPeriod().AsDuration(), nil
				}
				extractResultValue, err := extractor()
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorColumnValueExtractor(task.Table, column, err)
				} else {
					return extractResultValue, nil
				}
			})).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("crypto_key_backend").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CryptoKeyBackend")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("primary").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Primary")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("purpose").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Purpose")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_template").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VersionTemplate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_kms_keyrings_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_kms_keyrings.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("next_rotation_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("NextRotationTime")).Build(),
	}
}

func (x *TableGcpKmsCryptoKeysGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpKmsCryptoKeyVersionsGenerator{}),
	}
}
