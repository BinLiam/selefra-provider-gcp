package iam

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	policies "cloud.google.com/go/iam/apiv2"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	policiespb "google.golang.org/genproto/googleapis/iam/v2"
)

type TableGcpIamDenyPoliciesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpIamDenyPoliciesGenerator{}

func (x *TableGcpIamDenyPoliciesGenerator) GetTableName() string {
	return "gcp_iam_deny_policies"
}

func (x *TableGcpIamDenyPoliciesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpIamDenyPoliciesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpIamDenyPoliciesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpIamDenyPoliciesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			iamClient, err := policies.NewPoliciesClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			iamClient.CallOptions = &policies.PoliciesCallOptions{}

			parent := fmt.Sprintf("policies/cloudresourcemanager.googleapis.com%%2Fprojects%%2F%s/denypolicies", c.ProjectId)
			req := &policiespb.ListPoliciesRequest{
				Parent:   parent,
				PageSize: 1000,
			}

			it := iamClient.ListPolicies(ctx, req, c.CallOptions...)

			for {
				policy, err := it.Next()
				if err == iterator.Done {
					break
				}
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}

				resultChannel <- []*policiespb.Policy{policy}
			}

			return nil
		},
	}
}

func (x *TableGcpIamDenyPoliciesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpIamDenyPoliciesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delete_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DeleteTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Rules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("managing_authority").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ManagingAuthority")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uid").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uid")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("annotations").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Annotations")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("etag").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Etag")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpIamDenyPoliciesGenerator) GetSubTables() []*schema.Table {
	return nil
}
