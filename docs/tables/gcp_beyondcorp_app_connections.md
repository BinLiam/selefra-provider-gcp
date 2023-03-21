# Table: gcp_beyondcorp_app_connections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| type | big_int | X | √ |  | 
| application_endpoint | json | X | √ |  | 
| state | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gateway | json | X | √ |  | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| display_name | string | X | √ |  | 
| uid | string | X | √ |  | 
| connectors | string_array | X | √ |  | 


