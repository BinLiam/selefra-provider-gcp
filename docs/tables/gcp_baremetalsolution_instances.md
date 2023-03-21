# Table: gcp_baremetalsolution_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| machine_type | string | X | √ |  | 
| luns | json | X | √ |  | 
| pod | string | X | √ |  | 
| name | string | X | √ |  | 
| update_time | json | X | √ |  | 
| network_template | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| hyperthreading_enabled | bool | X | √ |  | 
| os_image | string | X | √ |  | 
| state | big_int | X | √ |  | 
| interactive_serial_console_enabled | bool | X | √ |  | 
| labels | json | X | √ |  | 
| networks | json | X | √ |  | 
| logical_interfaces | json | X | √ |  | 
| project_id | string | X | √ |  | 
| id | string | X | √ |  | 
| create_time | json | X | √ |  | 


