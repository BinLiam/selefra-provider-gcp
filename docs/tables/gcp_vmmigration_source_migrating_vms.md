# Table: gcp_vmmigration_source_migrating_vms

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| last_sync | json | X | √ |  | 
| recent_clone_jobs | json | X | √ |  | 
| project_id | string | X | √ |  | 
| update_time | json | X | √ |  | 
| state_time | json | X | √ |  | 
| source_vm_id | string | X | √ |  | 
| state | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| error | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_vmmigration_sources_selefra_id | string | X | X | fk to gcp_vmmigration_sources.selefra_id | 
| display_name | string | X | √ |  | 
| policy | json | X | √ |  | 
| create_time | json | X | √ |  | 
| current_sync_info | json | X | √ |  | 
| group | string | X | √ |  | 
| recent_cutover_jobs | json | X | √ |  | 


