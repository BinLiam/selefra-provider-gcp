# Table: gcp_artifactregistry_packages

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| gcp_artifactregistry_repositories_selefra_id | string | X | X | fk to gcp_artifactregistry_repositories.selefra_id | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


