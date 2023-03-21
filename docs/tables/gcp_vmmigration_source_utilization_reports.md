# Table: gcp_vmmigration_source_utilization_reports

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| error | json | X | √ |  | 
| create_time | json | X | √ |  | 
| frame_end_time | json | X | √ |  | 
| vms | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| state_time | json | X | √ |  | 
| display_name | string | X | √ |  | 
| state | big_int | X | √ |  | 
| time_frame | big_int | X | √ |  | 
| vm_count | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_vmmigration_sources_selefra_id | string | X | X | fk to gcp_vmmigration_sources.selefra_id | 
| name | string | X | √ |  | 


