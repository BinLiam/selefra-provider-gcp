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

type TableGcpLivestreamInputsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpLivestreamInputsGenerator{}

func (x *TableGcpLivestreamInputsGenerator) GetTableName() string {
	return "gcp_livestream_inputs"
}

func (x *TableGcpLivestreamInputsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpLivestreamInputsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpLivestreamInputsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpLivestreamInputsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)

			gcpClient, err := livestream.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListInputs(ctx, &pb.ListInputsRequest{
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

func (x *TableGcpLivestreamInputsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpLivestreamInputsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("uri").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Uri")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("preprocessing_config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PreprocessingConfig")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("input_stream_property").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("InputStreamProperty")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("tier").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Tier")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("security_rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("SecurityRules")).Build(),
	}
}

func (x *TableGcpLivestreamInputsGenerator) GetSubTables() []*schema.Table {
	return nil
}
