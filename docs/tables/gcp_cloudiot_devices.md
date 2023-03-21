# Table: gcp_cloudiot_devices

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| blocked | bool | X | √ |  | 
| last_error_time | json | X | √ |  | 
| state | json | X | √ |  | 
| log_level | big_int | X | √ |  | 
| id | string | X | √ |  | 
| num_id | big_int | X | √ |  | 
| last_state_time | json | X | √ |  | 
| last_config_send_time | json | X | √ |  | 
| config | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_cloudiot_device_registries_selefra_id | string | X | X | fk to gcp_cloudiot_device_registries.selefra_id | 
| credentials | json | X | √ |  | 
| last_event_time | json | X | √ |  | 
| last_config_ack_time | json | X | √ |  | 
| last_error_status | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| last_heartbeat_time | json | X | √ |  | 
| metadata | json | X | √ |  | 
| gateway_config | json | X | √ |  | 


