# Table: gcp_vmmigration_source_datacenter_connectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| available_versions | json | X | √ |  | 
| create_time | json | X | √ |  | 
| registration_id | string | X | √ |  | 
| bucket | string | X | √ |  | 
| upgrade_status | json | X | √ |  | 
| project_id | string | X | √ |  | 
| update_time | json | X | √ |  | 
| version | string | X | √ |  | 
| error | json | X | √ |  | 
| service_account | string | X | √ |  | 
| state_time | json | X | √ |  | 
| appliance_software_version | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_vmmigration_sources_selefra_id | string | X | X | fk to gcp_vmmigration_sources.selefra_id | 
| name | string | X | √ |  | 
| state | big_int | X | √ |  | 
| appliance_infrastructure_version | string | X | √ |  | 


