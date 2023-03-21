# Table: gcp_kms_crypto_key_versions

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| destroy_time | json | X | √ |  | 
| reimport_eligible | bool | X | √ |  | 
| name | string | X | √ |  | 
| destroy_event_time | json | X | √ |  | 
| import_job | string | X | √ |  | 
| external_protection_level_options | json | X | √ |  | 
| project_id | string | X | √ |  | 
| generate_time | json | X | √ |  | 
| import_time | json | X | √ |  | 
| gcp_kms_crypto_keys_selefra_id | string | X | X | fk to gcp_kms_crypto_keys.selefra_id | 
| state | big_int | X | √ |  | 
| algorithm | big_int | X | √ |  | 
| attestation | json | X | √ |  | 
| create_time | json | X | √ |  | 
| import_failure_reason | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| protection_level | big_int | X | √ |  | 


