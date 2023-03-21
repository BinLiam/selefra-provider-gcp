# Table: gcp_clouddeploy_releases

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| description | string | X | √ |  | 
| annotations | json | X | √ |  | 
| labels | json | X | √ |  | 
| create_time | json | X | √ |  | 
| render_start_time | json | X | √ |  | 
| render_state | big_int | X | √ |  | 
| gcp_clouddeploy_delivery_pipelines_selefra_id | string | X | X | fk to gcp_clouddeploy_delivery_pipelines.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| abandoned | bool | X | √ |  | 
| name | string | X | √ |  | 
| uid | string | X | √ |  | 
| skaffold_config_path | string | X | √ |  | 
| delivery_pipeline_snapshot | json | X | √ |  | 
| target_snapshots | json | X | √ |  | 
| etag | string | X | √ |  | 
| skaffold_version | string | X | √ |  | 
| target_artifacts | json | X | √ |  | 
| render_end_time | json | X | √ |  | 
| skaffold_config_uri | string | X | √ |  | 
| build_artifacts | json | X | √ |  | 
| target_renders | json | X | √ |  | 
| project_id | string | X | √ |  | 


