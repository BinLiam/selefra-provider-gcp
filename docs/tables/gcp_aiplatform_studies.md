# Table: gcp_aiplatform_studies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| study_spec | json | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_aiplatform_vizier_locations_selefra_id | string | X | X | fk to gcp_aiplatform_vizier_locations.selefra_id | 
| state | big_int | X | √ |  | 
| create_time | json | X | √ |  | 
| inactive_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


