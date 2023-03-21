# Table: gcp_clouddeploy_rollouts

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| etag | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| description | string | X | √ |  | 
| create_time | json | X | √ |  | 
| deploying_build | string | X | √ |  | 
| deploy_failure_cause | big_int | X | √ |  | 
| phases | json | X | √ |  | 
| metadata | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_clouddeploy_releases_selefra_id | string | X | X | fk to gcp_clouddeploy_releases.selefra_id | 
| approve_time | json | X | √ |  | 
| approval_state | big_int | X | √ |  | 
| state | big_int | X | √ |  | 
| target_id | string | X | √ |  | 
| name | string | X | √ |  | 
| uid | string | X | √ |  | 
| annotations | json | X | √ |  | 
| deploy_end_time | json | X | √ |  | 
| failure_reason | string | X | √ |  | 
| labels | json | X | √ |  | 
| enqueue_time | json | X | √ |  | 
| deploy_start_time | json | X | √ |  | 


