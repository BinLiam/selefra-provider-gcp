# Table: gcp_websecurityscanner_scan_config_scan_run_findings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| body | string | X | √ |  | 
| frame_url | string | X | √ |  | 
| violating_resource | json | X | √ |  | 
| vulnerable_headers | json | X | √ |  | 
| xss | json | X | √ |  | 
| severity | big_int | X | √ |  | 
| reproduction_url | string | X | √ |  | 
| final_url | string | X | √ |  | 
| form | json | X | √ |  | 
| outdated_library | json | X | √ |  | 
| name | string | X | √ |  | 
| tracking_id | string | X | √ |  | 
| xxe | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_websecurityscanner_scan_config_scan_runs_selefra_id | string | X | X | fk to gcp_websecurityscanner_scan_config_scan_runs.selefra_id | 
| finding_type | string | X | √ |  | 
| http_method | string | X | √ |  | 
| fuzzed_url | string | X | √ |  | 
| description | string | X | √ |  | 
| vulnerable_parameters | json | X | √ |  | 


