package vpcaccess

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	vpcaccess "cloud.google.com/go/vpcaccess/apiv1"
	pb "cloud.google.com/go/vpcaccess/apiv1/vpcaccesspb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

type TableGcpVpcaccessConnectorsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpVpcaccessConnectorsGenerator{}

func (x *TableGcpVpcaccessConnectorsGenerator) GetTableName() string {
	return "gcp_vpcaccess_connectors"
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			parentItem := task.ParentRawResult.(*locationpb.Location)

			gcpClient, err := vpcaccess.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			it := gcpClient.ListConnectors(ctx, &pb.ListConnectorsRequest{
				Parent: parentItem.Name,
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

func (x *TableGcpVpcaccessConnectorsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_vpcaccess_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_vpcaccess_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_throughput").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxThroughput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("subnet").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Subnet")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("machine_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MachineType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("max_instances").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MaxInstances")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ip_cidr_range").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("IpCidrRange")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_throughput").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinThroughput")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Network")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("connected_projects").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("ConnectedProjects")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("min_instances").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("MinInstances")).Build(),
	}
}

func (x *TableGcpVpcaccessConnectorsGenerator) GetSubTables() []*schema.Table {
	return nil
}
