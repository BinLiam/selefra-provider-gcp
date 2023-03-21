# Table: gcp_batch_tasks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| status | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_batch_task_groups_selefra_id | string | X | X | fk to gcp_batch_task_groups.selefra_id | 


