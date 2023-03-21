package compute

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	compute "cloud.google.com/go/compute/apiv1"
	pb "cloud.google.com/go/compute/apiv1/computepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpComputeInterconnectsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpComputeInterconnectsGenerator{}

func (x *TableGcpComputeInterconnectsGenerator) GetTableName() string {
	return "gcp_compute_interconnects"
}

func (x *TableGcpComputeInterconnectsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpComputeInterconnectsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpComputeInterconnectsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableGcpComputeInterconnectsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListInterconnectsRequest{
				Project: c.ProjectId,
			}
			gcpClient, err := compute.NewInterconnectsRESTClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.List(ctx, req, c.CallOptions...)
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

func (x *TableGcpComputeInterconnectsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpComputeInterconnectsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("expected_outages").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("ExpectedOutages")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("kind").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Kind")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("noc_contact_email").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NocContactEmail")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("admin_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("AdminEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("circuit_infos").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CircuitInfos")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("creation_timestamp").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CreationTimestamp")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("customer_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CustomerName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Description")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioned_link_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ProvisionedLinkCount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("operational_status").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OperationalStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("satisfies_pzs").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("SatisfiesPzs")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("self_link").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("SelfLink")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("google_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GoogleIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interconnect_attachments").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("InterconnectAttachments")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interconnect_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InterconnectType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("peer_ip_address").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PeerIpAddress")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("google_reference_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GoogleReferenceId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("link_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LinkType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Location")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("requested_link_count").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("RequestedLinkCount")).Build(),
	}
}

func (x *TableGcpComputeInterconnectsGenerator) GetSubTables() []*schema.Table {
	return nil
}
