# Table: gcp_compute_autoscalers

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| status_details | json | X | √ |  | 
| project_id | string | X | √ |  | 
| scaling_schedule_status | json | X | √ |  | 
| zone | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| description | string | X | √ |  | 
| id | big_int | X | √ |  | 
| region | string | X | √ |  | 
| kind | string | X | √ |  | 
| status | string | X | √ |  | 
| target | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| autoscaling_policy | json | X | √ |  | 
| recommended_size | big_int | X | √ |  | 
| self_link | string | X | √ |  | 


