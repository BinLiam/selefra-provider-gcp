# Table: gcp_kms_crypto_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| import_only | bool | X | √ |  | 
| destroy_scheduled_duration | json | X | √ |  | 
| project_id | string | X | √ |  | 
| rotation_period | big_int | X | √ |  | 
| create_time | json | X | √ |  | 
| crypto_key_backend | string | X | √ |  | 
| primary | json | X | √ |  | 
| purpose | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| version_template | json | X | √ |  | 
| labels | json | X | √ |  | 
| gcp_kms_keyrings_selefra_id | string | X | X | fk to gcp_kms_keyrings.selefra_id | 
| next_rotation_time | json | X | √ |  | 


