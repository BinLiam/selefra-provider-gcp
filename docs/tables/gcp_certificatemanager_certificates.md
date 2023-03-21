# Table: gcp_certificatemanager_certificates

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| name | string | X | √ |  | 
| create_time | json | X | √ |  | 
| update_time | json | X | √ |  | 
| san_dnsnames | string_array | X | √ |  | 
| expire_time | json | X | √ |  | 
| description | string | X | √ |  | 
| labels | json | X | √ |  | 
| pem_certificate | string | X | √ |  | 
| scope | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


