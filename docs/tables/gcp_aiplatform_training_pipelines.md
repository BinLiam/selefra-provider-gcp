# Table: gcp_aiplatform_training_pipelines

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| training_task_inputs | json | X | √ |  | 
| training_task_metadata | json | X | √ |  | 
| parent_model | string | X | √ |  | 
| error | json | X | √ |  | 
| encryption_spec | json | X | √ |  | 
| name | string | X | √ |  | 
| model_id | string | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| input_data_config | json | X | √ |  | 
| training_task_definition | string | X | √ |  | 
| start_time | json | X | √ |  | 
| gcp_aiplatform_pipeline_locations_selefra_id | string | X | X | fk to gcp_aiplatform_pipeline_locations.selefra_id | 
| labels | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| display_name | string | X | √ |  | 
| model_to_upload | json | X | √ |  | 
| state | big_int | X | √ |  | 
| end_time | json | X | √ |  | 


