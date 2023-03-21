# Table: gcp_appengine_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| instance_class | string | X | √ |  | 
| readiness_check | json | X | √ |  | 
| vpc_access_connector | json | X | √ |  | 
| project_id | string | X | √ |  | 
| resources | json | X | √ |  | 
| runtime | string | X | √ |  | 
| vm | bool | X | √ |  | 
| runtime_api_version | string | X | √ |  | 
| handlers | json | X | √ |  | 
| error_handlers | json | X | √ |  | 
| env_variables | json | X | √ |  | 
| runtime_channel | string | X | √ |  | 
| beta_settings | json | X | √ |  | 
| disk_usage_bytes | big_int | X | √ |  | 
| libraries | json | X | √ |  | 
| default_expiration | json | X | √ |  | 
| health_check | json | X | √ |  | 
| entrypoint | json | X | √ |  | 
| network | json | X | √ |  | 
| create_time | json | X | √ |  | 
| runtime_main_executable_path | string | X | √ |  | 
| deployment | json | X | √ |  | 
| version_url | string | X | √ |  | 
| inbound_services | int_array | X | √ |  | 
| app_engine_apis | bool | X | √ |  | 
| env | string | X | √ |  | 
| created_by | string | X | √ |  | 
| api_config | json | X | √ |  | 
| build_env_variables | json | X | √ |  | 
| gcp_appengine_services_selefra_id | string | X | X | fk to gcp_appengine_services.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| threadsafe | bool | X | √ |  | 
| serving_status | big_int | X | √ |  | 
| service_account | string | X | √ |  | 
| liveness_check | json | X | √ |  | 
| endpoints_api_service | json | X | √ |  | 
| id | string | X | √ |  | 
| zones | string_array | X | √ |  | 
| nobuild_files_regex | string | X | √ |  | 


