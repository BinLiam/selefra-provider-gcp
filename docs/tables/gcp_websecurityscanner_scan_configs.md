# Table: gcp_websecurityscanner_scan_configs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| authentication | json | X | √ |  | 
| user_agent | big_int | X | √ |  | 
| export_to_security_command_center | big_int | X | √ |  | 
| name | string | X | √ |  | 
| starting_urls | string_array | X | √ |  | 
| risk_level | big_int | X | √ |  | 
| static_ip_scan | bool | X | √ |  | 
| project_id | string | X | √ |  | 
| display_name | string | X | √ |  | 
| max_qps | big_int | X | √ |  | 
| schedule | json | X | √ |  | 
| managed_scan | bool | X | √ |  | 
| ignore_http_status_errors | bool | X | √ |  | 
| blacklist_patterns | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


