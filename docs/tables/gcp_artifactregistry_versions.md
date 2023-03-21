# Table: gcp_artifactregistry_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| update_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_artifactregistry_packages_selefra_id | string | X | X | fk to gcp_artifactregistry_packages.selefra_id | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| metadata | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_time | json | X | √ |  | 
| related_tags | json | X | √ |  | 


