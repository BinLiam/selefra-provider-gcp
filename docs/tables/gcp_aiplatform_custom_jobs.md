# Table: gcp_aiplatform_custom_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| update_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| end_time | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| web_access_uris | json | X | √ |  | 
| gcp_aiplatform_job_locations_selefra_id | string | X | X | fk to gcp_aiplatform_job_locations.selefra_id | 
| display_name | string | X | √ |  | 
| job_spec | json | X | √ |  | 
| state | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_time | json | X | √ |  | 
| error | json | X | √ |  | 
| labels | json | X | √ |  | 
| start_time | json | X | √ |  | 


