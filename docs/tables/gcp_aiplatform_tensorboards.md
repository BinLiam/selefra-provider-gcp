# Table: gcp_aiplatform_tensorboards

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| labels | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_tensorboard_locations_selefra_id | string | X | X | fk to gcp_aiplatform_tensorboard_locations.selefra_id | 
| encryption_spec | json | X | √ |  | 
| blob_storage_path_prefix | string | X | √ |  | 
| run_count | big_int | X | √ |  | 
| update_time | json | X | √ |  | 
| etag | string | X | √ |  | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| description | string | X | √ |  | 
| create_time | json | X | √ |  | 


