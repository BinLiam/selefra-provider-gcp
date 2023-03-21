# Table: gcp_aiplatform_specialist_pools

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| specialist_managers_count | big_int | X | √ |  | 
| specialist_manager_emails | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_specialistpool_locations_selefra_id | string | X | X | fk to gcp_aiplatform_specialistpool_locations.selefra_id | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 
| project_id | string | X | √ |  | 
| pending_data_labeling_jobs | string_array | X | √ |  | 
| specialist_worker_emails | string_array | X | √ |  | 


