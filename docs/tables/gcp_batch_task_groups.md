# Table: gcp_batch_task_groups

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| task_spec | json | X | √ |  | 
| task_count | big_int | X | √ |  | 
| parallelism | big_int | X | √ |  | 
| require_hosts_file | bool | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| task_environments | json | X | √ |  | 
| task_count_per_node | big_int | X | √ |  | 
| permissive_ssh | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


