# Table: gcp_baremetalsolution_volume_luns

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| state | big_int | X | √ |  | 
| shareable | bool | X | √ |  | 
| boot_lun | bool | X | √ |  | 
| storage_type | big_int | X | √ |  | 
| wwid | string | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| id | string | X | √ |  | 
| size_gb | big_int | X | √ |  | 
| multiprotocol_type | big_int | X | √ |  | 
| storage_volume | string | X | √ |  | 
| gcp_baremetalsolution_volumes_selefra_id | string | X | X | fk to gcp_baremetalsolution_volumes.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 


