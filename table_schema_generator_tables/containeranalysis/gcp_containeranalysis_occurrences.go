package containeranalysis

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	pb "cloud.google.com/go/containeranalysis/apiv1beta1/grafeas/grafeaspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpContaineranalysisOccurrencesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpContaineranalysisOccurrencesGenerator{}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetTableName() string {
	return "gcp_containeranalysis_occurrences"
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListOccurrencesRequest{
				Parent: "projects/" + c.ProjectId,
			}
			gcpClient, err := containeranalysis.NewGrafeasV1Beta1Client(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListOccurrences(ctx, req, c.CallOptions...)
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

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("remediation").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Remediation")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("resource").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Resource")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("note_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NoteName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
	}
}

func (x *TableGcpContaineranalysisOccurrencesGenerator) GetSubTables() []*schema.Table {
	return nil
}
