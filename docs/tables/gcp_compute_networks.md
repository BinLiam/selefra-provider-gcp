# Table: gcp_compute_networks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_ula_internal_ipv6 | bool | X | √ |  | 
| id | big_int | X | √ |  | 
| peerings | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| gateway_i_pv4 | string | X | √ |  | 
| name | string | X | √ |  | 
| self_link | string | X | √ |  | 
| self_link_with_id | string | X | √ |  | 
| subnetworks | string_array | X | √ |  | 
| auto_create_subnetworks | bool | X | √ |  | 
| description | string | X | √ |  | 
| kind | string | X | √ |  | 
| network_firewall_policy_enforcement_order | string | X | √ |  | 
| routing_config | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| i_pv4_range | string | X | √ |  | 
| firewall_policy | string | X | √ |  | 
| internal_ipv6_range | string | X | √ |  | 
| mtu | big_int | X | √ |  | 
| project_id | string | X | √ |  | 


