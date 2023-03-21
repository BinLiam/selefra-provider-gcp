# Table: gcp_aiplatform_models

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| encryption_spec | json | X | √ |  | 
| version_create_time | json | X | √ |  | 
| description | string | X | √ |  | 
| artifact_uri | string | X | √ |  | 
| supported_output_storage_formats | string_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| version_id | string | X | √ |  | 
| predict_schemata | json | X | √ |  | 
| training_pipeline | string | X | √ |  | 
| supported_input_storage_formats | string_array | X | √ |  | 
| display_name | string | X | √ |  | 
| version_description | string | X | √ |  | 
| update_time | json | X | √ |  | 
| deployed_models | json | X | √ |  | 
| explanation_spec | json | X | √ |  | 
| etag | string | X | √ |  | 
| model_source_info | json | X | √ |  | 
| project_id | string | X | √ |  | 
| version_aliases | string_array | X | √ |  | 
| metadata_schema_uri | string | X | √ |  | 
| supported_export_formats | json | X | √ |  | 
| labels | json | X | √ |  | 
| version_update_time | json | X | √ |  | 
| container_spec | json | X | √ |  | 
| supported_deployment_resources_types | int_array | X | √ |  | 
| metadata | json | X | √ |  | 
| create_time | json | X | √ |  | 
| metadata_artifact | string | X | √ |  | 
| gcp_aiplatform_model_locations_selefra_id | string | X | X | fk to gcp_aiplatform_model_locations.selefra_id | 
| name | string | X | √ |  | 


