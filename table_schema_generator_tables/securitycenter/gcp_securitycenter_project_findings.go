package securitycenter

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpSecuritycenterProjectFindingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpSecuritycenterProjectFindingsGenerator{}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetTableName() string {
	return "gcp_securitycenter_project_findings"
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			p := "projects/" + c.ProjectId + "/sources/-"
			maybeError := fetchFindings(p)(ctx, client, resultChannel)
			return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
		},
	}
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("finding").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Finding")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state_change").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StateChange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Finding.Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpSecuritycenterProjectFindingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
