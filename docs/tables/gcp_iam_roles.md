# Table: gcp_iam_roles

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| included_permissions | string_array | X | √ |  | 
| deleted | bool | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| title | string | X | √ |  | 
| stage | big_int | X | √ |  | 
| etag | int_array | X | √ |  | 


