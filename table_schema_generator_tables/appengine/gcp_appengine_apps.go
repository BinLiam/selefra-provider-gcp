package appengine

import (
	"context"
	"github.com/selefra/selefra-provider-gcp/gcp_client"

	appengine "cloud.google.com/go/appengine/apiv1"
	pb "cloud.google.com/go/appengine/apiv1/appenginepb"
	"github.com/selefra/selefra-provider-gcp/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableGcpAppengineAppsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableGcpAppengineAppsGenerator{}

func (x *TableGcpAppengineAppsGenerator) GetTableName() string {
	return "gcp_appengine_apps"
}

func (x *TableGcpAppengineAppsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableGcpAppengineAppsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableGcpAppengineAppsGenerator) GetOptions() *schema.TableOptions {
return &schema.TableOptions{}
}

func (x *TableGcpAppengineAppsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c := client.(*gcp_client.Client)
			req := &pb.GetApplicationRequest{
				Name: "apps/" + c.ProjectId,
			}
			gcpClient, err := appengine.NewApplicationsClient(ctx, c.ClientOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resp, err := gcpClient.GetApplication(ctx, req, c.CallOptions...)
			if err != nil {
				return schema.NewDiagnosticsErrorPullTable(task.Table, err)

			}
			resultChannel <- resp
			return nil
		},
	}
}

func (x *TableGcpAppengineAppsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableGcpAppengineAppsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("default_cookie_expiration").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DefaultCookieExpiration")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("feature_settings").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("FeatureSettings")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("code_bucket").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("CodeBucket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("service_account").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ServiceAccount")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("iap").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("Iap")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Name")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("Id")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("dispatch_rules").ColumnType(schema.ColumnTypeJSON).
			Extractor(column_value_extractor.StructSelector("DispatchRules")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("auth_domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AuthDomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("location_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LocationId")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("serving_status").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("ServingStatus")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("gcr_domain").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GcrDomain")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_hostname").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultHostname")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_bucket").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("DefaultBucket")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("database_type").ColumnType(schema.ColumnTypeBigInt).
			Extractor(column_value_extractor.StructSelector("DatabaseType")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("project_id").ColumnType(schema.ColumnTypeString).
			Extractor(gcp_client.ExtractorProject()).Build(),
	}
}

func (x *TableGcpAppengineAppsGenerator) GetSubTables() []*schema.Table {
	return nil
}
