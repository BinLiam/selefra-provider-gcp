package binaryauthorization

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	binaryauthorization "cloud.google.com/go/binaryauthorization/apiv1"
	pb "cloud.google.com/go/binaryauthorization/apiv1/binaryauthorizationpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBinaryauthorizationAssertorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBinaryauthorizationAssertorsGenerator{}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetTableName() string {
	return "gcp_binaryauthorization_assertors"
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListAttestorsRequest{
				Parent: "projects/" + c.ProjectId,
			}
			gcpClient, err := binaryauthorization.NewBinauthzManagementClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListAttestors(ctx, req, c.CallOptions...)
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

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
	}
}

func (x *TableGcpBinaryauthorizationAssertorsGenerator) GetSubTables() []*schema.Table {
	return nil
}
