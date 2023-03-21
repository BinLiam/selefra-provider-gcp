package videotranscoder

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	transcoder "cloud.google.com/go/video/transcoder/apiv1"
	pb "cloud.google.com/go/video/transcoder/apiv1/transcoderpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpVideotranscoderJobTemplatesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVideotranscoderJobTemplatesGenerator{}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetTableName() string {
	return "gcp_videotranscoder_job_templates"
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := transcoder.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListJobTemplates(ctx, &pb.ListJobTemplatesRequest{
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

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Config")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpVideotranscoderJobTemplatesGenerator) GetSubTables() []*schema.Table {
	return nil
}