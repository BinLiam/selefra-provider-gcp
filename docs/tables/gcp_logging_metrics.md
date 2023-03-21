# Table: gcp_logging_metrics

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| update_time | json | X | √ |  | 
| version | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| filter | string | X | √ |  | 
| disabled | bool | X | √ |  | 
| label_extractors | json | X | √ |  | 
| bucket_options | json | X | √ |  | 
| description | string | X | √ |  | 
| metric_descriptor | json | X | √ |  | 
| value_extractor | string | X | √ |  | 
| create_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


