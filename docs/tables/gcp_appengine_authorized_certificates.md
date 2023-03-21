# Table: gcp_appengine_authorized_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| id | string | X | √ |  | 
| display_name | string | X | √ |  | 
| visible_domain_mappings | string_array | X | √ |  | 
| domain_mappings_count | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| domain_names | string_array | X | √ |  | 
| expire_time | json | X | √ |  | 
| certificate_raw_data | json | X | √ |  | 
| managed_certificate | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


