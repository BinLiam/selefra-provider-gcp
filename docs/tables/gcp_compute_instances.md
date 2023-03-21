# Table: gcp_compute_instances

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| start_restricted | bool | X | √ |  | 
| advanced_machine_features | json | X | √ |  | 
| creation_timestamp | string | X | √ |  | 
| description | string | X | √ |  | 
| last_suspended_timestamp | string | X | √ |  | 
| machine_type | string | X | √ |  | 
| source_machine_image | string | X | √ |  | 
| source_machine_image_encryption_key | json | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| disks | json | X | √ |  | 
| last_stop_timestamp | string | X | √ |  | 
| network_performance_config | json | X | √ |  | 
| params | json | X | √ |  | 
| scheduling | json | X | √ |  | 
| shielded_instance_config | json | X | √ |  | 
| deletion_protection | bool | X | √ |  | 
| id | big_int | X | √ |  | 
| reservation_affinity | json | X | √ |  | 
| satisfies_pzs | bool | X | √ |  | 
| self_link | string | X | √ |  | 
| display_device | json | X | √ |  | 
| last_start_timestamp | string | X | √ |  | 
| min_cpu_platform | string | X | √ |  | 
| name | string | X | √ |  | 
| network_interfaces | json | X | √ |  | 
| confidential_instance_config | json | X | √ |  | 
| fingerprint | string | X | √ |  | 
| resource_status | json | X | √ |  | 
| status | string | X | √ |  | 
| zone | string | X | √ |  | 
| project_id | string | X | √ |  | 
| cpu_platform | string | X | √ |  | 
| key_revocation_action_type | string | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| labels | json | X | √ |  | 
| metadata | json | X | √ |  | 
| tags | json | X | √ |  | 
| guest_accelerators | json | X | √ |  | 
| resource_policies | string_array | X | √ |  | 
| shielded_instance_integrity_policy | json | X | √ |  | 
| status_message | string | X | √ |  | 
| can_ip_forward | bool | X | √ |  | 
| hostname | string | X | √ |  | 
| kind | string | X | √ |  | 
| private_ipv6_google_access | string | X | √ |  | 
| service_accounts | json | X | √ |  | 


