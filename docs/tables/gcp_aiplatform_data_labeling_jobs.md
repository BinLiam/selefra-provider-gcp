# Table: gcp_aiplatform_data_labeling_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| project_id | string | X | √ |  | 
| labeler_count | big_int | X | √ |  | 
| inputs_schema_uri | string | X | √ |  | 
| update_time | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| instruction_uri | string | X | √ |  | 
| current_spend | json | X | √ |  | 
| error | json | X | √ |  | 
| specialist_pools | string_array | X | √ |  | 
| active_learning_config | json | X | √ |  | 
| datasets | string_array | X | √ |  | 
| inputs | json | X | √ |  | 
| annotation_labels | json | X | √ |  | 
| state | big_int | X | √ |  | 
| labeling_progress | big_int | X | √ |  | 
| create_time | json | X | √ |  | 
| gcp_aiplatform_job_locations_selefra_id | string | X | X | fk to gcp_aiplatform_job_locations.selefra_id | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 


