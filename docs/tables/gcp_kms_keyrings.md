# Table: gcp_kms_keyrings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_kms_locations_selefra_id | string | X | X | fk to gcp_kms_locations.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


