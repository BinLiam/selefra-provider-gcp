# Table: gcp_artifactregistry_docker_images

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| uri | string | X | √ |  | 
| tags | string_array | X | √ |  | 
| image_size_bytes | big_int | X | √ |  | 
| build_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| upload_time | json | X | √ |  | 
| media_type | string | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_artifactregistry_repositories_selefra_id | string | X | X | fk to gcp_artifactregistry_repositories.selefra_id | 


