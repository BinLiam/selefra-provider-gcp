package billing

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	billing "cloud.google.com/go/billing/apiv1"
	pb "cloud.google.com/go/billing/apiv1/billingpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBillingServicesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBillingServicesGenerator{}

func (x *TableGcpBillingServicesGenerator) GetTableName() string {
	return "gcp_billing_services"
}

func (x *TableGcpBillingServicesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBillingServicesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBillingServicesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpBillingServicesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListServicesRequest{}
			gcpClient, err := billing.NewCloudCatalogClient(ctx, c.ClientOptions...)
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

func (x *TableGcpBillingServicesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return gcp_client.ExpandByProjects()
}

func (x *TableGcpBillingServicesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("business_entity_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("BusinessEntityName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpBillingServicesGenerator) GetSubTables() []*schema.Table {
	return nil
}
