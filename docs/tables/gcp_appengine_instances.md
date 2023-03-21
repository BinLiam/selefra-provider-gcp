# Table: gcp_appengine_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| vm_id | string | X | √ |  | 
| start_time | json | X | √ |  | 
| memory_usage | big_int | X | √ |  | 
| gcp_appengine_versions_selefra_id | string | X | X | fk to gcp_appengine_versions.selefra_id | 
| app_engine_release | string | X | √ |  | 
| vm_name | string | X | √ |  | 
| average_latency | big_int | X | √ |  | 
| vm_ip | string | X | √ |  | 
| vm_liveness | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| id | string | X | √ |  | 
| errors | big_int | X | √ |  | 
| qps | float | X | √ |  | 
| vm_status | string | X | √ |  | 
| vm_debug_enabled | bool | X | √ |  | 
| name | string | X | √ |  | 
| availability | big_int | X | √ |  | 
| vm_zone_name | string | X | √ |  | 
| requests | big_int | X | √ |  | 


