package cloudscheduler

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/cloudscheduler/v1"
)

type TableGcpCloudschedulerLocationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpCloudschedulerLocationsGenerator{}

func (x *TableGcpCloudschedulerLocationsGenerator) GetTableName() string {
	return "gcp_cloudscheduler_locations"
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			nextPageToken := ""
			gcpClient, err := cloudscheduler.NewService(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}

			for {
				output, err := gcpClient.Projects.Locations.List("projects/" + c.ProjectId).PageToken(nextPageToken).Context(ctx).Do()
				if err != nil {
					return schema.NewDiagnosticsErrorPullTable(task.Table, err)

				}
				resultChannel <- output.Locations
				if output.NextPageToken == "" {
					break
				}
				nextPageToken = output.NextPageToken
			}
			return nil
		},
	}
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("labels").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Labels")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("metadata").ColumnType(schema.ColumnTypeIntArray).
			Extractor(column_value_extractor.StructSelector("Metadata")).Build(),
	}
}

func (x *TableGcpCloudschedulerLocationsGenerator) GetSubTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&TableGcpCloudschedulerJobsGenerator{}),
	}
}
