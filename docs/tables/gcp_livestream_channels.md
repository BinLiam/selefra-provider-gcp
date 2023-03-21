# Table: gcp_livestream_channels

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| manifests | json | X | √ |  | 
| project_id | string | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| active_input | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| elementary_streams | json | X | √ |  | 
| mux_streams | json | X | √ |  | 
| streaming_error | json | X | √ |  | 
| name | string | X | √ |  | 
| input_attachments | json | X | √ |  | 
| sprite_sheets | json | X | √ |  | 
| log_config | json | X | √ |  | 
| create_time | json | X | √ |  | 
| output | json | X | √ |  | 
| streaming_state | big_int | X | √ |  | 


