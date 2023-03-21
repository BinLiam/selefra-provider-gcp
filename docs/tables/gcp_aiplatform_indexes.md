# Table: gcp_aiplatform_indexes

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| etag | string | X | √ |  | 
| labels | json | X | √ |  | 
| update_time | json | X | √ |  | 
| index_update_method | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_aiplatform_index_locations_selefra_id | string | X | X | fk to gcp_aiplatform_index_locations.selefra_id | 
| description | string | X | √ |  | 
| metadata_schema_uri | string | X | √ |  | 
| metadata | json | X | √ |  | 
| index_stats | json | X | √ |  | 
| display_name | string | X | √ |  | 
| deployed_indexes | json | X | √ |  | 
| create_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


