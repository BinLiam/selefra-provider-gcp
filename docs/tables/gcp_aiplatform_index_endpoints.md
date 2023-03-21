# Table: gcp_aiplatform_index_endpoints

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_indexendpoint_locations_selefra_id | string | X | X | fk to gcp_aiplatform_indexendpoint_locations.selefra_id | 
| description | string | X | √ |  | 
| update_time | json | X | √ |  | 
| network | string | X | √ |  | 
| enable_private_service_connect | bool | X | √ |  | 
| labels | json | X | √ |  | 
| create_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| deployed_indexes | json | X | √ |  | 
| etag | string | X | √ |  | 


