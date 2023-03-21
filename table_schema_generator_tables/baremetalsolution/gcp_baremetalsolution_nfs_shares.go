package baremetalsolution

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBaremetalsolutionNfsSharesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBaremetalsolutionNfsSharesGenerator{}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetTableName() string {
	return "gcp_baremetalsolution_nfs_shares"
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListNfsSharesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListNfsShares(ctx, req, c.CallOptions...)
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

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("volume").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Volume")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allowed_clients").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("AllowedClients")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("nfs_share_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NfsShareId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
	}
}

func (x *TableGcpBaremetalsolutionNfsSharesGenerator) GetSubTables() []*schema.Table {
	return nil
}