# Table: gcp_kms_import_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| generate_time | json | X | √ |  | 
| state | big_int | X | √ |  | 
| attestation | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| create_time | json | X | √ |  | 
| import_method | big_int | X | √ |  | 
| protection_level | big_int | X | √ |  | 
| expire_time | json | X | √ |  | 
| expire_event_time | json | X | √ |  | 
| public_key | json | X | √ |  | 
| gcp_kms_keyrings_selefra_id | string | X | X | fk to gcp_kms_keyrings.selefra_id | 
| name | string | X | √ |  | 


