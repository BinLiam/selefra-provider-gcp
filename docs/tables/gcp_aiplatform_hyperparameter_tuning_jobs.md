# Table: gcp_aiplatform_hyperparameter_tuning_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| trials | json | X | √ |  | 
| error | json | X | √ |  | 
| end_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| display_name | string | X | √ |  | 
| study_spec | json | X | √ |  | 
| max_trial_count | big_int | X | √ |  | 
| trial_job_spec | json | X | √ |  | 
| state | big_int | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| parallel_trial_count | big_int | X | √ |  | 
| start_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_job_locations_selefra_id | string | X | X | fk to gcp_aiplatform_job_locations.selefra_id | 
| max_failed_trial_count | big_int | X | √ |  | 
| create_time | json | X | √ |  | 


