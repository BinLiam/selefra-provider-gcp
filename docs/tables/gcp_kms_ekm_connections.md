# Table: gcp_kms_ekm_connections

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| gcp_kms_locations_selefra_id | string | X | X | fk to gcp_kms_locations.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| service_resolvers | json | X | √ |  | 
| etag | string | X | √ |  | 


