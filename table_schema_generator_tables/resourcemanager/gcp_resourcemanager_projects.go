package resourcemanager

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	"cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpResourcemanagerProjectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpResourcemanagerProjectsGenerator{}

func (x *TableGcpResourcemanagerProjectsGenerator) GetTableName() string {
	return "gcp_resourcemanager_projects"
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &resourcemanagerpb.GetProjectRequest{
				Name: "projects/" + c.ProjectId,
			}
			projectsClient, err := resourcemanager.NewProjectsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			output, err := projectsClient.GetProject(ctx, req)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- output
			return nil
		},
	}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeleteTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Parent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
	}
}

func (x *TableGcpResourcemanagerProjectsGenerator) GetSubTables() []*schema.Table {
	return nil
}
