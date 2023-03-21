# Table: gcp_websecurityscanner_scan_config_scan_run_crawled_urls

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| scan_run_name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_websecurityscanner_scan_config_scan_runs_selefra_id | string | X | X | fk to gcp_websecurityscanner_scan_config_scan_runs.selefra_id | 
| http_method | string | X | √ |  | 
| url | string | X | √ |  | 
| body | string | X | √ |  | 
| project_id | string | X | √ |  | 


