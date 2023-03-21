# Table: gcp_artifactregistry_files

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| size_bytes | big_int | X | √ |  | 
| update_time | json | X | √ |  | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| owner | string | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_artifactregistry_repositories_selefra_id | string | X | X | fk to gcp_artifactregistry_repositories.selefra_id | 
| hashes | json | X | √ |  | 


