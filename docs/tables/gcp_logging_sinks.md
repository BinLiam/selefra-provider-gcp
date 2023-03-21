# Table: gcp_logging_sinks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| output_version_format | big_int | X | √ |  | 
| writer_identity | string | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| name | string | X | √ |  | 
| filter | string | X | √ |  | 
| description | string | X | √ |  | 
| include_children | bool | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| destination | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| exclusions | json | X | √ |  | 


