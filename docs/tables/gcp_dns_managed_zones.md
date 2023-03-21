# Table: gcp_dns_managed_zones

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| name_servers | string_array | X | √ |  | 
| peering_config | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| cloud_logging_config | json | X | √ |  | 
| private_visibility_config | json | X | √ |  | 
| creation_time | string | X | √ |  | 
| service_directory_config | json | X | √ |  | 
| dnssec_config | json | X | √ |  | 
| dns_name | string | X | √ |  | 
| forwarding_config | json | X | √ |  | 
| id | big_int | X | √ |  | 
| kind | string | X | √ |  | 
| labels | json | X | √ |  | 
| name_server_set | string | X | √ |  | 
| reverse_lookup_config | json | X | √ |  | 
| description | string | X | √ |  | 
| visibility | string | X | √ |  | 


