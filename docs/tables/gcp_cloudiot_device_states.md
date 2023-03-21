# Table: gcp_cloudiot_device_states

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| update_time | json | X | √ |  | 
| binary_data | int_array | X | √ |  | 
| project_id | string | X | √ |  | 
| device_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_cloudiot_devices_selefra_id | string | X | X | fk to gcp_cloudiot_devices.selefra_id | 


