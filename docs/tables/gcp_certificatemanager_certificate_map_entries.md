# Table: gcp_certificatemanager_certificate_map_entries

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| create_time | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| state | big_int | X | √ |  | 
| project_id | string | X | √ |  | 
| gcp_certificatemanager_certificate_maps_selefra_id | string | X | X | fk to gcp_certificatemanager_certificate_maps.selefra_id | 
| name | string | X | √ |  | 
| description | string | X | √ |  | 
| update_time | json | X | √ |  | 
| labels | json | X | √ |  | 
| certificates | string_array | X | √ |  | 


