package translate

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	translate "cloud.google.com/go/translate/apiv3"
	pb "cloud.google.com/go/translate/apiv3/translatepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpTranslateGlossariesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpTranslateGlossariesGenerator{}

func (x *TableGcpTranslateGlossariesGenerator) GetTableName() string {
	return "gcp_translate_glossaries"
}

func (x *TableGcpTranslateGlossariesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpTranslateGlossariesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpTranslateGlossariesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpTranslateGlossariesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := translate.NewTranslationClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListGlossaries(ctx, &pb.ListGlossariesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/us-central1",
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

func (x *TableGcpTranslateGlossariesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpTranslateGlossariesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InputConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("entry_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("EntryCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("submit_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SubmitTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("end_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("EndTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpTranslateGlossariesGenerator) GetSubTables() []*schema.Table {
	return nil
}
