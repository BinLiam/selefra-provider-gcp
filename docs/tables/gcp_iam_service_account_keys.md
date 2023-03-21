# Table: gcp_iam_service_account_keys

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| key_type | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| private_key_type | big_int | X | √ |  | 
| public_key_data | int_array | X | √ |  | 
| gcp_iam_service_accounts_selefra_id | string | X | X | fk to gcp_iam_service_accounts.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| valid_before_time | json | X | √ |  | 
| disabled | bool | X | √ |  | 
| key_algorithm | big_int | X | √ |  | 
| private_key_data | int_array | X | √ |  | 
| valid_after_time | json | X | √ |  | 
| key_origin | big_int | X | √ |  | 
| service_account_unique_id | string | X | √ |  | 


