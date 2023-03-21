package baremetalsolution

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	baremetalsolution "cloud.google.com/go/baremetalsolution/apiv2"
	pb "cloud.google.com/go/baremetalsolution/apiv2/baremetalsolutionpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
)

type TableGcpBaremetalsolutionInstancesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpBaremetalsolutionInstancesGenerator{}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetTableName() string {
	return "gcp_baremetalsolution_instances"
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListInstancesRequest{
				Parent: "projects/" + c.ProjectId + "/locations/-",
			}
			gcpClient, err := baremetalsolution.NewClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListInstances(ctx, req, c.CallOptions...)
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

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("machine_type").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("MachineType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("luns").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Luns")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pod").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Pod")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("update_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("UpdateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("network_template").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("NetworkTemplate")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("hyperthreading_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("HyperthreadingEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("os_image").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OsImage")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("interactive_serial_console_enabled").ColumnType(schema.ColumnTypeBool).
			Extractor(column_value_extractor.StructSelector("InteractiveSerialConsoleEnabled")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("networks").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Networks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("logical_interfaces").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("LogicalInterfaces")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
	}
}

func (x *TableGcpBaremetalsolutionInstancesGenerator) GetSubTables() []*schema.Table {
	return nil
}
