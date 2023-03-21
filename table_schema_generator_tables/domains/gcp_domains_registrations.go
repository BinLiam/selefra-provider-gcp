package domains

import (
	"context"
	"fmt"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	domains "cloud.google.com/go/domains/apiv1beta1"
	pb "cloud.google.com/go/domains/apiv1beta1/domainspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpDomainsRegistrationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpDomainsRegistrationsGenerator{}

func (x *TableGcpDomainsRegistrationsGenerator) GetTableName() string {
	return "gcp_domains_registrations"
}

func (x *TableGcpDomainsRegistrationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpDomainsRegistrationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpDomainsRegistrationsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpDomainsRegistrationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListRegistrationsRequest{
				Parent: fmt.Sprintf("projects/%s/locations/-", c.ProjectId),
			}
			gcpClient, err := domains.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListRegistrations(ctx, req, c.CallOptions...)
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

func (x *TableGcpDomainsRegistrationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpDomainsRegistrationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("issues").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("Issues")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("management_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ManagementSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dns_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DnsSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pending_contact_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("PendingContactSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expire_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExpireTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("contact_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ContactSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("domain_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DomainName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("supported_privacy").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("SupportedPrivacy")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpDomainsRegistrationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
