# Table: gcp_batch_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notifications | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| priority | big_int | X | √ |  | 
| allocation_policy | json | X | √ |  | 
| create_time | json | X | √ |  | 
| logs_policy | json | X | √ |  | 
| uid | string | X | √ |  | 
| task_groups | json | X | √ |  | 
| labels | json | X | √ |  | 
| status | json | X | √ |  | 
| update_time | json | X | √ |  | 


