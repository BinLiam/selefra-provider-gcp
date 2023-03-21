# Table: gcp_securitycenter_organization_findings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state_change | big_int | X | √ |  | 
| resource | json | X | √ |  | 
| organization_id | string | X | √ |  | 
| name | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| finding | json | X | √ |  | 


