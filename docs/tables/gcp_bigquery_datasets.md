# Table: gcp_bigquery_datasets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| access | json | X | √ |  | 
| dataset_reference | json | X | √ |  | 
| etag | string | X | √ |  | 
| storage_billing_model | string | X | √ |  | 
| creation_time | big_int | X | √ |  | 
| default_collation | string | X | √ |  | 
| default_table_expiration_ms | big_int | X | √ |  | 
| is_case_insensitive | bool | X | √ |  | 
| kind | string | X | √ |  | 
| max_time_travel_hours | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| default_encryption_configuration | json | X | √ |  | 
| default_partition_expiration_ms | big_int | X | √ |  | 
| description | string | X | √ |  | 
| friendly_name | string | X | √ |  | 
| labels | json | X | √ |  | 
| last_modified_time | big_int | X | √ |  | 
| location | string | X | √ |  | 
| self_link | string | X | √ |  | 
| tags | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


