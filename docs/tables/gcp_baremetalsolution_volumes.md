# Table: gcp_baremetalsolution_volumes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| emergency_size_gib | big_int | X | √ |  | 
| auto_grown_size_gib | big_int | X | √ |  | 
| snapshot_reservation_detail | json | X | √ |  | 
| snapshot_enabled | bool | X | √ |  | 
| id | string | X | √ |  | 
| state | big_int | X | √ |  | 
| current_size_gib | big_int | X | √ |  | 
| requested_size_gib | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| storage_type | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| pod | string | X | √ |  | 
| remaining_space_gib | big_int | X | √ |  | 
| snapshot_auto_delete_behavior | big_int | X | √ |  | 
| labels | json | X | √ |  | 


