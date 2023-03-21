# Table: gcp_vpcaccess_connectors

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| gcp_vpcaccess_locations_selefra_id | string | X | X | fk to gcp_vpcaccess_locations.selefra_id | 
| state | big_int | X | √ |  | 
| max_throughput | big_int | X | √ |  | 
| subnet | json | X | √ |  | 
| machine_type | string | X | √ |  | 
| max_instances | big_int | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| ip_cidr_range | string | X | √ |  | 
| min_throughput | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| network | string | X | √ |  | 
| connected_projects | string_array | X | √ |  | 
| min_instances | big_int | X | √ |  | 


