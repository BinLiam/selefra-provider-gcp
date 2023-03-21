# Table: gcp_run_services

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| api_version | string | X | √ |  | 
| kind | string | X | √ |  | 
| metadata | json | X | √ |  | 
| spec | json | X | √ |  | 
| status | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_run_locations_selefra_id | string | X | X | fk to gcp_run_locations.selefra_id | 


