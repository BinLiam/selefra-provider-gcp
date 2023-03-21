# Table: gcp_artifactregistry_tags

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| version | string | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_artifactregistry_packages_selefra_id | string | X | X | fk to gcp_artifactregistry_packages.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


