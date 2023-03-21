package iam

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	iamadmin "cloud.google.com/go/iam/admin/apiv1"
	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpIamServiceAccountKeysGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpIamServiceAccountKeysGenerator{}

func (x *TableGcpIamServiceAccountKeysGenerator) GetTableName() string {
	return "gcp_iam_service_account_keys"
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := task.ParentRawResult.(*iampb.ServiceAccount)
			iamClient, err := iamadmin.NewIamClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			iamClient.CallOptions = &iamadmin.IamCallOptions{}

			req := &iampb.ListServiceAccountKeysRequest{
				Name: p.Name,
			}

			output, err := iamClient.ListServiceAccountKeys(ctx, req, c.CallOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			resultChannel <- output.Keys
			return nil
		},
	}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("key_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("KeyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_key_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("PrivateKeyType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key_data").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("PublicKeyData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_iam_service_accounts_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_iam_service_accounts.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_before_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidBeforeTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_algorithm").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("KeyAlgorithm")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_key_data").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("PrivateKeyData")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("valid_after_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ValidAfterTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("key_origin").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("KeyOrigin")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account_unique_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.ParentColumnValue("unique_id")).Build(),
	}
}

func (x *TableGcpIamServiceAccountKeysGenerator) GetSubTables() []*schema.Table {
	return nil
}
