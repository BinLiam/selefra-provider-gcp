package resourcemanager

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	resourcemanager "cloud.google.com/go/resourcemanager/apiv3"
	pb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpResourcemanagerFoldersGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpResourcemanagerFoldersGenerator{}

func (x *TableGcpResourcemanagerFoldersGenerator) GetTableName() string {
	return "gcp_resourcemanager_folders"
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			fClient, err := resourcemanager.NewFoldersClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			req := &pb.ListFoldersRequest{
				Parent: "organizations/" + c.OrgId,
			}
			it := fClient.ListFolders(ctx, req)
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

func (x *TableGcpResourcemanagerFoldersGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandOrgMultiplex()
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DeleteTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Parent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorOrganization()).Build(),
	}
}

func (x *TableGcpResourcemanagerFoldersGenerator) GetSubTables() []*schema.Table {
	return nil
}
