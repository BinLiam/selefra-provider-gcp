# Table: gcp_aiplatform_featurestores

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_spec | json | X | √ |  | 
| project_id | string | X | √ |  | 
| update_time | json | X | √ |  | 
| etag | string | X | √ |  | 
| labels | json | X | √ |  | 
| online_serving_config | json | X | √ |  | 
| state | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| gcp_aiplatform_featurestore_locations_selefra_id | string | X | X | fk to gcp_aiplatform_featurestore_locations.selefra_id | 


