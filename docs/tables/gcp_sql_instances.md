# Table: gcp_sql_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| database_installed_version | string | X | √ |  | 
| gce_zone | string | X | √ |  | 
| name | string | X | √ |  | 
| root_password | string | X | √ |  | 
| backend_type | string | X | √ |  | 
| ip_addresses | json | X | √ |  | 
| ipv6_address | string | X | √ |  | 
| kind | string | X | √ |  | 
| maintenance_version | string | X | √ |  | 
| replica_configuration | json | X | √ |  | 
| replica_names | string_array | X | √ |  | 
| self_link | string | X | √ |  | 
| service_account_email_address | string | X | √ |  | 
| create_time | string | X | √ |  | 
| disk_encryption_configuration | json | X | √ |  | 
| disk_encryption_status | json | X | √ |  | 
| server_ca_cert | json | X | √ |  | 
| project_id | string | X | √ |  | 
| current_disk_size | big_int | X | √ |  | 
| database_version | string | X | √ |  | 
| settings | json | X | √ |  | 
| etag | string | X | √ |  | 
| failover_replica | json | X | √ |  | 
| master_instance_name | string | X | √ |  | 
| on_premises_configuration | json | X | √ |  | 
| region | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| instance_type | string | X | √ |  | 
| max_disk_size | big_int | X | √ |  | 
| out_of_disk_report | json | X | √ |  | 
| scheduled_maintenance | json | X | √ |  | 
| available_maintenance_versions | string_array | X | √ |  | 
| connection_name | string | X | √ |  | 
| project | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| secondary_gce_zone | string | X | √ |  | 
| state | string | X | √ |  | 
| suspension_reason | string_array | X | √ |  | 


