# Table: gcp_cloudiot_device_registries

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| mqtt_config | json | X | √ |  | 
| log_level | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| credentials | json | X | √ |  | 
| project_id | string | X | √ |  | 
| id | string | X | √ |  | 
| name | string | X | √ |  | 
| event_notification_configs | json | X | √ |  | 
| state_notification_config | json | X | √ |  | 
| http_config | json | X | √ |  | 


