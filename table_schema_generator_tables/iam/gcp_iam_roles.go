package iam

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	iamadmin "cloud.google.com/go/iam/admin/apiv1"
	iampb "cloud.google.com/go/iam/admin/apiv1/adminpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpIamRolesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpIamRolesGenerator{}

func (x *TableGcpIamRolesGenerator) GetTableName() string {
	return "gcp_iam_roles"
}

func (x *TableGcpIamRolesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpIamRolesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpIamRolesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpIamRolesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""

			iamClient, err := iamadmin.NewIamClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			iamClient.CallOptions = &iamadmin.IamCallOptions{}

			for {
				req := &iampb.ListRolesRequest{
					PageSize:  1000,
					PageToken: nextPageToken,
					Parent:    fmt.Sprintf("projects/%s", c.ProjectId),
				}
				resp, err := iamClient.ListRoles(ctx, req, c.CallOptions...)
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- resp.Roles

				if resp.NextPageToken == "" {
					break
				}
				nextPageToken = resp.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpIamRolesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpIamRolesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("included_permissions").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("IncludedPermissions")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("deleted").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Deleted")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("title").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Title")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stage").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Stage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
	}
}

func (x *TableGcpIamRolesGenerator) GetSubTables() []*schema.Table {
	return nil
}
