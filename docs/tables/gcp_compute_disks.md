# Table: gcp_compute_disks

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| provisioned_iops | big_int | X | √ |  | 
| architecture | string | X | √ |  | 
| last_attach_timestamp | string | X | √ |  | 
| location_hint | string | X | √ |  | 
| replica_zones | string_array | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| source_disk_id | string | X | √ |  | 
| id | big_int | X | √ |  | 
| size_gb | big_int | X | √ |  | 
| source_disk | string | X | √ |  | 
| source_image | string | X | √ |  | 
| source_snapshot_id | string | X | √ |  | 
| status | string | X | √ |  | 
| users | string_array | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| region | string | X | √ |  | 
| source_image_id | string | X | √ |  | 
| source_snapshot | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| description | string | X | √ |  | 
| physical_block_size_bytes | big_int | X | √ |  | 
| self_link | string | X | √ |  | 
| source_image_encryption_key | json | X | √ |  | 
| type | string | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| licenses | string_array | X | √ |  | 
| resource_policies | string_array | X | √ |  | 
| labels | json | X | √ |  | 
| source_snapshot_encryption_key | json | X | √ |  | 
| source_storage_object | string | X | √ |  | 
| zone | string | X | √ |  | 
| guest_os_features | json | X | √ |  | 
| kind | string | X | √ |  | 
| last_detach_timestamp | string | X | √ |  | 
| license_codes | int_array | X | √ |  | 
| name | string | X | √ |  | 
| options | string | X | √ |  | 
| params | json | X | √ |  | 
| project_id | string | X | √ |  | 
| disk_encryption_key | json | X | √ |  | 


