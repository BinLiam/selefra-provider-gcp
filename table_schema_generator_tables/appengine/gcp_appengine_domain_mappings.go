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

type TableGcpAppengineDomainMappingsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineDomainMappingsGenerator{}

func (x *TableGcpAppengineDomainMappingsGenerator) GetTableName() string {
	return "gcp_appengine_domain_mappings"
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListDomainMappingsRequest{
				Parent: "apps/" + c.ProjectId,
			}
			gcpClient, err := appengine.NewDomainMappingsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListDomainMappings(ctx, req, c.CallOptions...)
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

func (x *TableGcpAppengineDomainMappingsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("resource_records").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ResourceRecords")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ssl_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SslSettings")).Build(),
	}
}

func (x *TableGcpAppengineDomainMappingsGenerator) GetSubTables() []*schema.Table {
	return nil
}
