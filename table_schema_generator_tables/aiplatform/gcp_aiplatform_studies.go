package aiplatform

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	pb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/cloud/location"
)

type TableGcpAiplatformStudiesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAiplatformStudiesGenerator{}

func (x *TableGcpAiplatformStudiesGenerator) GetTableName() string {
	return "gcp_aiplatform_studies"
}

func (x *TableGcpAiplatformStudiesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAiplatformStudiesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAiplatformStudiesGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAiplatformStudiesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.ListStudiesRequest{
				Parent: task.ParentRawResult.(*location.Location).Name,
			}
			if filterStudiesLocation(task) {
				return nil
			}

			clientOptions := c.ClientOptions
			clientOptions = append([]option.ClientOption{option.WithEndpoint(task.ParentRawResult.(*location.Location).LocationId + "-aiplatform.googleapis.com:443")}, clientOptions...)
			gcpClient, err := aiplatform.NewVizierClient(ctx, clientOptions...)

			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			it := gcpClient.ListStudies(ctx, req, c.CallOptions...)
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

func filterStudiesLocation(task *schema.DataSourcePullTask) bool {
	if filterLocation(task.ParentRawResult.(*location.Location).LocationId) {
		return true
	}

	locationId := task.ParentRawResult.(*location.Location).LocationId

	toFilter := []string{"asia-southeast2", "europe-central2", "me-west1", "us-south1", "us-west3"}
	for _, f := range toFilter {
		if locationId == f {
			return true
		}
	}
	return false
}

func (x *TableGcpAiplatformStudiesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAiplatformStudiesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("display_name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DisplayName")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("study_spec").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("StudySpec")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcp_aiplatform_vizier_locations_selefra_id").ColumnType(schema.ColumnTypeString).SetNotNull().Description("fk to gcp_aiplatform_vizier_locations.selefra_id").
			Extractor(column_value_extractor.ParentColumnValue("selefra_id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("State")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("create_time").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("CreateTime")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("inactive_reason").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("InactiveReason")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableGcpAiplatformStudiesGenerator) GetSubTables() []*schema.Table {
	return nil
}
