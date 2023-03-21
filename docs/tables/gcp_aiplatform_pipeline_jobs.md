# Table: gcp_aiplatform_pipeline_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| update_time | json | X | √ |  | 
| state | big_int | X | √ |  | 
| job_detail | json | X | √ |  | 
| error | json | X | √ |  | 
| project_id | string | X | √ |  | 
| labels | json | X | √ |  | 
| template_metadata | json | X | √ |  | 
| create_time | json | X | √ |  | 
| start_time | json | X | √ |  | 
| pipeline_spec | json | X | √ |  | 
| runtime_config | json | X | √ |  | 
| network | string | X | √ |  | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| end_time | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| service_account | string | X | √ |  | 
| template_uri | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_pipeline_locations_selefra_id | string | X | X | fk to gcp_aiplatform_pipeline_locations.selefra_id | 


