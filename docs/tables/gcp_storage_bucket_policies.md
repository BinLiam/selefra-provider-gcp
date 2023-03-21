# Table: gcp_storage_bucket_policies

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| bucket_name | string | X | √ |  | 
| gcp_storage_buckets_selefra_id | string | X | X | fk to gcp_storage_buckets.selefra_id | 
| selefra_id | string | √ | √ | random id | 
| bindings | json | X | √ |  | 
| project_id | string | X | √ |  | 


