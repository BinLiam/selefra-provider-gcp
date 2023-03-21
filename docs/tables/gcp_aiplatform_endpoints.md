# Table: gcp_aiplatform_endpoints

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| labels | json | X | √ |  | 
| gcp_aiplatform_endpoint_locations_selefra_id | string | X | X | fk to gcp_aiplatform_endpoint_locations.selefra_id | 
| deployed_models | json | X | √ |  | 
| create_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| traffic_split | json | X | √ |  | 
| etag | string | X | √ |  | 
| enable_private_service_connect | bool | X | √ |  | 
| model_deployment_monitoring_job | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| display_name | string | X | √ |  | 
| update_time | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| network | string | X | √ |  | 
| predict_request_response_logging_config | json | X | √ |  | 


