# Table: gcp_clouderrorreporting_error_events

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| service_context | json | X | √ |  | 
| message | string | X | √ |  | 
| context | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| gcp_clouderrorreporting_error_group_stats_selefra_id | string | X | X | fk to gcp_clouderrorreporting_error_group_stats.selefra_id | 
| event_time | json | X | √ |  | 


