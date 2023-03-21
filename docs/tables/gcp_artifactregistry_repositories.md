# Table: gcp_artifactregistry_repositories

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| format | big_int | X | √ |  | 
| description | string | X | √ |  | 
| labels | json | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| kms_key_name | string | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_artifactregistry_locations_selefra_id | string | X | X | fk to gcp_artifactregistry_locations.selefra_id | 


