# Table: gcp_bigquery_tables

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| clone_definition | json | X | √ |  | 
| last_modified_time | big_int | X | √ |  | 
| table_reference | json | X | √ |  | 
| time_partitioning | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| streaming_buffer | json | X | √ |  | 
| description | string | X | √ |  | 
| external_data_configuration | json | X | √ |  | 
| max_staleness | string | X | √ |  | 
| num_long_term_bytes | big_int | X | √ |  | 
| num_active_physical_bytes | big_int | X | √ |  | 
| num_long_term_logical_bytes | big_int | X | √ |  | 
| require_partition_filter | bool | X | √ |  | 
| view | json | X | √ |  | 
| project_id | string | X | √ |  | 
| clustering | json | X | √ |  | 
| kind | string | X | √ |  | 
| model | json | X | √ |  | 
| num_rows | big_int | X | √ |  | 
| range_partitioning | json | X | √ |  | 
| snapshot_definition | json | X | √ |  | 
| id | string | X | √ |  | 
| etag | string | X | √ |  | 
| friendly_name | string | X | √ |  | 
| self_link | string | X | √ |  | 
| gcp_bigquery_datasets_selefra_id | string | X | X | fk to gcp_bigquery_datasets.selefra_id | 
| num_long_term_physical_bytes | big_int | X | √ |  | 
| num_partitions | big_int | X | √ |  | 
| creation_time | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| location | string | X | √ |  | 
| num_bytes | big_int | X | √ |  | 
| num_physical_bytes | big_int | X | √ |  | 
| num_active_logical_bytes | big_int | X | √ |  | 
| num_time_travel_physical_bytes | big_int | X | √ |  | 
| type | string | X | √ |  | 
| schema | json | X | √ |  | 
| default_collation | string | X | √ |  | 
| encryption_configuration | json | X | √ |  | 
| expiration_time | big_int | X | √ |  | 
| materialized_view | json | X | √ |  | 
| num_total_logical_bytes | big_int | X | √ |  | 
| num_total_physical_bytes | big_int | X | √ |  | 


