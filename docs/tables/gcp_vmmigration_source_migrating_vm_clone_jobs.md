# Table: gcp_vmmigration_source_migrating_vm_clone_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_time | json | X | √ |  | 
| end_time | json | X | √ |  | 
| state | big_int | X | √ |  | 
| state_time | json | X | √ |  | 
| error | json | X | √ |  | 
| name | string | X | √ |  | 
| gcp_vmmigration_source_migrating_vms_selefra_id | string | X | X | fk to gcp_vmmigration_source_migrating_vms.selefra_id | 


