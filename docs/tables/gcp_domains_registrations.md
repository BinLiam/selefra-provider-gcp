# Table: gcp_domains_registrations

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| issues | int_array | X | √ |  | 
| labels | json | X | √ |  | 
| management_settings | json | X | √ |  | 
| dns_settings | json | X | √ |  | 
| pending_contact_settings | json | X | √ |  | 
| create_time | json | X | √ |  | 
| expire_time | json | X | √ |  | 
| state | big_int | X | √ |  | 
| contact_settings | json | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| domain_name | string | X | √ |  | 
| supported_privacy | int_array | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 


