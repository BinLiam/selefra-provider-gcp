# Table: gcp_aiplatform_model_deployment_monitoring_jobs

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| enable_monitoring_pipeline_logs | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_aiplatform_job_locations_selefra_id | string | X | X | fk to gcp_aiplatform_job_locations.selefra_id | 
| display_name | string | X | √ |  | 
| model_deployment_monitoring_objective_configs | json | X | √ |  | 
| predict_instance_schema_uri | string | X | √ |  | 
| stats_anomalies_base_directory | json | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| endpoint | string | X | √ |  | 
| schedule_state | big_int | X | √ |  | 
| create_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| update_time | json | X | √ |  | 
| next_schedule_time | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| latest_monitoring_pipeline_metadata | json | X | √ |  | 
| logging_sampling_strategy | json | X | √ |  | 
| sample_predict_instance | json | X | √ |  | 
| analysis_instance_schema_uri | string | X | √ |  | 
| log_ttl | json | X | √ |  | 
| error | json | X | √ |  | 
| state | big_int | X | √ |  | 
| model_deployment_monitoring_schedule_config | json | X | √ |  | 
| model_monitoring_alert_config | json | X | √ |  | 
| bigquery_tables | json | X | √ |  | 


