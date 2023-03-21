# Table: gcp_cloudiot_device_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| gcp_cloudiot_devices_selefra_id | string | X | X | fk to gcp_cloudiot_devices.selefra_id | 
| version | big_int | X | √ |  | 
| cloud_update_time | json | X | √ |  | 
| device_ack_time | json | X | √ |  | 
| binary_data | int_array | X | √ |  | 
| project_id | string | X | √ |  | 
| device_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


