package appengine

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpAppengineAuthorizedDomainsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineAuthorizedDomainsGenerator{}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetTableName() string {
	return "gcp_appengine_authorized_domains"
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListAuthorizedDomainsRequest{
				Parent: "apps/" + c.ProjectId,
			}
			gcpClient, err := appengine.NewAuthorizedDomainsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListAuthorizedDomains(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpAppengineAuthorizedDomainsGenerator) GetSubTables() []*schema.Table {
	return nil
}
