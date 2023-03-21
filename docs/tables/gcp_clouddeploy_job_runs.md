# Table: gcp_clouddeploy_job_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uid | string | X | √ |  | 
| phase_id | string | X | √ |  | 
| start_time | json | X | √ |  | 
| end_time | json | X | √ |  | 
| etag | string | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| state | big_int | X | √ |  | 
| gcp_clouddeploy_rollouts_selefra_id | string | X | X | fk to gcp_clouddeploy_rollouts.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| job_id | string | X | √ |  | 


