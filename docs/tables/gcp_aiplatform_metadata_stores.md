# Table: gcp_aiplatform_metadata_stores

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_spec | json | X | √ |  | 
| description | string | X | √ |  | 
| state | json | X | √ |  | 
| project_id | string | X | √ |  | 
| update_time | json | X | √ |  | 
| create_time | json | X | √ |  | 
| gcp_aiplatform_metadata_locations_selefra_id | string | X | X | fk to gcp_aiplatform_metadata_locations.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 


