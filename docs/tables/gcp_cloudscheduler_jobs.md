# Table: gcp_cloudscheduler_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| pubsub_target | json | X | √ |  | 
| retry_config | json | X | √ |  | 
| status | json | X | √ |  | 
| attempt_deadline | string | X | √ |  | 
| description | string | X | √ |  | 
| http_target | json | X | √ |  | 
| last_attempt_time | string | X | √ |  | 
| gcp_cloudscheduler_locations_selefra_id | string | X | X | fk to gcp_cloudscheduler_locations.selefra_id | 
| name | string | X | √ |  | 
| schedule | string | X | √ |  | 
| schedule_time | string | X | √ |  | 
| time_zone | string | X | √ |  | 
| state | string | X | √ |  | 
| user_update_time | string | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| app_engine_http_target | json | X | √ |  | 


