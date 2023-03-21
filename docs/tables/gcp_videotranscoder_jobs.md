# Table: gcp_videotranscoder_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| output_uri | string | X | √ |  | 
| end_time | json | X | √ |  | 
| ttl_after_completion_days | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| error | json | X | √ |  | 
| name | string | X | √ |  | 
| state | big_int | X | √ |  | 
| create_time | json | X | √ |  | 
| start_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| input_uri | string | X | √ |  | 


