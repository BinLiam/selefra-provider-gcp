package sql

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	sql "google.golang.org/api/sqladmin/v1beta4"
)

type TableGcpSqlUsersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSqlUsersGenerator{}

func (x *TableGcpSqlUsersGenerator) GetTableName() string {
	return "gcp_sql_users"
}

func (x *TableGcpSqlUsersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSqlUsersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSqlUsersGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpSqlUsersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			sqlClient, err := sql.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			instance := task.ParentRawResult.(*sql.DatabaseInstance)
			output, err := sqlClient.Users.List(c.ProjectId, instance.Name).Do()
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output.Items
			return nil
		},
	}
}

func (x *TableGcpSqlUsersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpSqlUsersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("host").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Host")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password_policy").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PasswordPolicy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sqlserver_user_details").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SqlserverUserDetails")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dual_password_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DualPasswordType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("password").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Password")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("instance").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Instance")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Project")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_sql_instances_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_sql_instances.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
	}
}

func (x *TableGcpSqlUsersGenerator) GetSubTables() []*schema.Table {
	return nil
}
