package provider

import (
	"context"

	"github.com/selefra/selefra-provider-gcp/gcp_client"

	"github.com/selefra/selefra-provider-sdk/provider"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/spf13/viper"
)

const Version = "v0.0.7"

func GetProvider() *provider.Provider {
	return &provider.Provider{
		Name:      "gcp",
		Version:   Version,
		TableList: GenTables(),
		ClientMeta: schema.ClientMeta{
			InitClient: func(ctx context.Context, clientMeta *schema.ClientMeta, config *viper.Viper) ([]any, *schema.Diagnostics) {
				var gcpConfig gcp_client.Config
				err := config.Unmarshal(&gcpConfig)
				if err != nil {
					return nil, schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}

				clients, err := gcp_client.NewClients(gcpConfig)

				if err != nil {
					clientMeta.ErrorF("new clients err: %s", err.Error())
					return nil, schema.NewDiagnostics().AddError(err)
				}

				if len(clients) == 0 {
					return nil, schema.NewDiagnostics().AddErrorMsg("account information not found")
				}

				hash := make(map[string]bool)
				res := make([]interface{}, 0, len(clients))
				for i := range clients {
					if hash[clients[i].ProjectId] {
						continue
					}
					res = append(res, clients[i])
					hash[clients[i].ProjectId] = true
				}
				return res, nil
			},
		},
		ConfigMeta: provider.ConfigMeta{
			GetDefaultConfigTemplate: func(ctx context.Context) string {
				return `# service_account_env_key: <YOUR_GCP_JSON>`
			},
			Validation: func(ctx context.Context, config *viper.Viper) *schema.Diagnostics {
				var gcpConfig gcp_client.Config
				err := config.Unmarshal(&gcpConfig)
				if err != nil {
					return schema.NewDiagnostics().AddErrorMsg("analysis config err: %s", err.Error())
				}
				return nil
			},
		},
		TransformerMeta: schema.TransformerMeta{
			DefaultColumnValueConvertorBlackList: []string{
				"",
				"N/A",
				"not_supported",
			},
			DataSourcePullResultAutoExpand: true,
		},
		ErrorsHandlerMeta: schema.ErrorsHandlerMeta{

			IgnoredErrors: []schema.IgnoredError{schema.IgnoredErrorOnSaveResult},
		},
	}
}
