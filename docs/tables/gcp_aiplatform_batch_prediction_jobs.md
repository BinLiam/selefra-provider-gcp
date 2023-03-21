# Table: gcp_aiplatform_batch_prediction_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| resources_consumed | json | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| project_id | string | X | √ |  | 
| model | string | X | √ |  | 
| service_account | string | X | √ |  | 
| partial_failures | json | X | √ |  | 
| error | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| model_version_id | string | X | √ |  | 
| model_parameters | json | X | √ |  | 
| explanation_spec | json | X | √ |  | 
| end_time | json | X | √ |  | 
| manual_batch_tuning_parameters | json | X | √ |  | 
| generate_explanation | bool | X | √ |  | 
| state | big_int | X | √ |  | 
| output_config | json | X | √ |  | 
| dedicated_resources | json | X | √ |  | 
| completion_stats | json | X | √ |  | 
| input_config | json | X | √ |  | 
| create_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| unmanaged_container_model | json | X | √ |  | 
| output_info | json | X | √ |  | 
| start_time | json | X | √ |  | 
| gcp_aiplatform_job_locations_selefra_id | string | X | X | fk to gcp_aiplatform_job_locations.selefra_id | 


