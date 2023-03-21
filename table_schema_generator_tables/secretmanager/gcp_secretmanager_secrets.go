package secretmanager

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	pb "cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpSecretmanagerSecretsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSecretmanagerSecretsGenerator{}

func (x *TableGcpSecretmanagerSecretsGenerator) GetTableName() string {
	return "gcp_secretmanager_secrets"
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListSecretsRequest{
				Parent: "projects/" + c.ProjectId,
			}
			gcpClient, err := secretmanager.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListSecrets(ctx, req, c.CallOptions...)
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

func (x *TableGcpSecretmanagerSecretsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rotation").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Rotation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("version_aliases").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("VersionAliases")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("replication").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Replication")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("topics").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Topics")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
	}
}

func (x *TableGcpSecretmanagerSecretsGenerator) GetSubTables() []*schema.Table {
	return nil
}
