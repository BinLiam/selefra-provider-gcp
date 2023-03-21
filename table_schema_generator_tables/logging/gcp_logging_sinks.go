package logging

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	logging "cloud.google.com/go/logging/apiv2"
	pb "cloud.google.com/go/logging/apiv2/loggingpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpLoggingSinksGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpLoggingSinksGenerator{}

func (x *TableGcpLoggingSinksGenerator) GetTableName() string {
	return "gcp_logging_sinks"
}

func (x *TableGcpLoggingSinksGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpLoggingSinksGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpLoggingSinksGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpLoggingSinksGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListSinksRequest{
				Parent: "projects/" + c.ProjectId,
			}
			gcpClient, err := logging.NewConfigClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListSinks(ctx, req, c.CallOptions...)
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

func (x *TableGcpLoggingSinksGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpLoggingSinksGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("output_version_format").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("OutputVersionFormat")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("writer_identity").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WriterIdentity")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("filter").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Filter")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("include_children").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("IncludeChildren")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("destination").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Destination")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("disabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("Disabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("exclusions").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Exclusions")).Build(),
	}
}

func (x *TableGcpLoggingSinksGenerator) GetSubTables() []*schema.Table {
	return nil
}
