package livestream

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	livestream "cloud.google.com/go/video/livestream/apiv1"
	pb "cloud.google.com/go/video/livestream/apiv1/livestreampb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpLivestreamChannelsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpLivestreamChannelsGenerator{}

func (x *TableGcpLivestreamChannelsGenerator) GetTableName() string {
	return "gcp_livestream_channels"
}

func (x *TableGcpLivestreamChannelsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpLivestreamChannelsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpLivestreamChannelsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpLivestreamChannelsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := livestream.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListChannels(ctx, &pb.ListChannelsRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}, c.CallOptions...)
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

func (x *TableGcpLivestreamChannelsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpLivestreamChannelsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("manifests").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Manifests")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("active_input").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ActiveInput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("elementary_streams").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ElementaryStreams")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("mux_streams").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("MuxStreams")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("streaming_error").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StreamingError")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_attachments").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InputAttachments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sprite_sheets").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SpriteSheets")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("log_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Output")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("streaming_state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("StreamingState")).Build(),
	}
}

func (x *TableGcpLivestreamChannelsGenerator) GetSubTables() []*schema.Table {
	return nil
}
