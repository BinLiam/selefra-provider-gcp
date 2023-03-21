# Table: gcp_vmmigration_source_migrating_vm_cutover_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| end_time | json | X | √ |  | 
| name | string | X | √ |  | 
| state | big_int | X | √ |  | 
| state_time | json | X | √ |  | 
| error | json | X | √ |  | 
| gcp_vmmigration_source_migrating_vms_selefra_id | string | X | X | fk to gcp_vmmigration_source_migrating_vms.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_time | json | X | √ |  | 
| progress_percent | big_int | X | √ |  | 
| state_message | string | X | √ |  | 
| project_id | string | X | √ |  | 


