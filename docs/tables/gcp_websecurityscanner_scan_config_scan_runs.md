# Table: gcp_websecurityscanner_scan_config_scan_runs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| urls_tested_count | big_int | X | √ |  | 
| progress_percent | big_int | X | √ |  | 
| warning_traces | json | X | √ |  | 
| gcp_websecurityscanner_scan_configs_selefra_id | string | X | X | fk to gcp_websecurityscanner_scan_configs.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| name | string | X | √ |  | 
| has_vulnerabilities | bool | X | √ |  | 
| error_trace | json | X | √ |  | 
| project_id | string | X | √ |  | 
| execution_state | big_int | X | √ |  | 
| result_state | big_int | X | √ |  | 
| start_time | json | X | √ |  | 
| end_time | json | X | √ |  | 
| urls_crawled_count | big_int | X | √ |  | 


