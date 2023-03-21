# Table: gcp_compute_backend_services

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| circuit_breakers | json | X | √ |  | 
| consistent_hash | json | X | √ |  | 
| custom_request_headers | string_array | X | √ |  | 
| edge_security_policy | string | X | √ |  | 
| iap | json | X | √ |  | 
| connection_draining | json | X | √ |  | 
| failover_policy | json | X | √ |  | 
| port | big_int | X | √ |  | 
| session_affinity | string | X | √ |  | 
| backends | json | X | √ |  | 
| enable_c_d_n | bool | X | √ |  | 
| max_stream_duration | json | X | √ |  | 
| protocol | string | X | √ |  | 
| self_link | string | X | √ |  | 
| connection_tracking_policy | json | X | √ |  | 
| custom_response_headers | string_array | X | √ |  | 
| kind | string | X | √ |  | 
| port_name | string | X | √ |  | 
| timeout_sec | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| id | big_int | X | √ |  | 
| locality_lb_policies | json | X | √ |  | 
| locality_lb_policy | string | X | √ |  | 
| name | string | X | √ |  | 
| security_policy | string | X | √ |  | 
| security_settings | json | X | √ |  | 
| compression_mode | string | X | √ |  | 
| fingerprint | string | X | √ |  | 
| health_checks | string_array | X | √ |  | 
| log_config | json | X | √ |  | 
| outlier_detection | json | X | √ |  | 
| subsetting | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| description | string | X | √ |  | 
| service_bindings | string_array | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| affinity_cookie_ttl_sec | big_int | X | √ |  | 
| cdn_policy | json | X | √ |  | 
| load_balancing_scheme | string | X | √ |  | 
| network | string | X | √ |  | 
| region | string | X | √ |  | 


