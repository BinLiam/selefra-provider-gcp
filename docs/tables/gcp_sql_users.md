# Table: gcp_sql_users

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| type | string | X | √ |  | 
| project_id | string | X | √ |  | 
| etag | string | X | √ |  | 
| host | string | X | √ |  | 
| password_policy | json | X | √ |  | 
| sqlserver_user_details | json | X | √ |  | 
| dual_password_type | string | X | √ |  | 
| password | string | X | √ |  | 
| instance | string | X | √ |  | 
| project | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| kind | string | X | √ |  | 
| gcp_sql_instances_selefra_id | string | X | X | fk to gcp_sql_instances.selefra_id | 


