# Table: gcp_compute_firewalls

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| direction | string | X | √ |  | 
| project_id | string | X | √ |  | 
| description | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| id | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| log_config | json | X | √ |  | 
| self_link | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| allowed | json | X | √ |  | 
| denied | json | X | √ |  | 
| name | string | X | √ |  | 
| network | string | X | √ |  | 
| source_ranges | string_array | X | √ |  | 
| target_tags | string_array | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| destination_ranges | string_array | X | √ |  | 
| priority | big_int | X | √ |  | 
| source_service_accounts | string_array | X | √ |  | 
| source_tags | string_array | X | √ |  | 
| target_service_accounts | string_array | X | √ |  | 


