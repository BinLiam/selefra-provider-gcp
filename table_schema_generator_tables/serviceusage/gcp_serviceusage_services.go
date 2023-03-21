package serviceusage

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	serviceusage "cloud.google.com/go/serviceusage/apiv1"
	pb "cloud.google.com/go/serviceusage/apiv1/serviceusagepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpServiceusageServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpServiceusageServicesGenerator{}

func (x *TableGcpServiceusageServicesGenerator) GetTableName() string {
	return "gcp_serviceusage_services"
}

func (x *TableGcpServiceusageServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpServiceusageServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpServiceusageServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpServiceusageServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			if len(c.EnabledServices) > 0 {
				for _, svc := range c.EnabledServices[c.ProjectId] {
					resultChannel <- svc.(*pb.Service)
				}
				return nil
			}
			req := &pb.ListServicesRequest{
				Parent:   "projects/" + c.ProjectId,
				PageSize: 200,
				Filter:   "state:ENABLED",
			}
			gcpClient, err := serviceusage.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListServices(ctx, req, c.CallOptions...)
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

func (x *TableGcpServiceusageServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpServiceusageServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("parent").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Parent")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Config")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpServiceusageServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
