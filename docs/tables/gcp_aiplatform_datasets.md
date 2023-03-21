# Table: gcp_aiplatform_datasets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| etag | string | X | √ |  | 
| metadata_artifact | string | X | √ |  | 
| gcp_aiplatform_dataset_locations_selefra_id | string | X | X | fk to gcp_aiplatform_dataset_locations.selefra_id | 
| display_name | string | X | √ |  | 
| metadata_schema_uri | string | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| description | string | X | √ |  | 
| create_time | json | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| project_id | string | X | √ |  | 
| metadata | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 


