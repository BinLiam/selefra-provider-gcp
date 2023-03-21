# Table: gcp_compute_images

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| source_disk_encryption_key | json | X | √ |  | 
| source_snapshot_encryption_key | json | X | √ |  | 
| deprecated | json | X | √ |  | 
| raw_disk | json | X | √ |  | 
| shielded_instance_initial_state | json | X | √ |  | 
| source_disk | string | X | √ |  | 
| kind | string | X | √ |  | 
| status | string | X | √ |  | 
| source_image_encryption_key | json | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| license_codes | int_array | X | √ |  | 
| source_disk_id | string | X | √ |  | 
| source_image | string | X | √ |  | 
| family | string | X | √ |  | 
| self_link | string | X | √ |  | 
| storage_locations | string_array | X | √ |  | 
| id | big_int | X | √ |  | 
| source_snapshot | string | X | √ |  | 
| project_id | string | X | √ |  | 
| architecture | string | X | √ |  | 
| archive_size_bytes | big_int | X | √ |  | 
| description | string | X | √ |  | 
| guest_os_features | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| disk_size_gb | big_int | X | √ |  | 
| labels | json | X | √ |  | 
| source_image_id | string | X | √ |  | 
| source_snapshot_id | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| source_type | string | X | √ |  | 
| image_encryption_key | json | X | √ |  | 
| licenses | string_array | X | √ |  | 
| name | string | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 


