# Table: gcp_container_clusters

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| monitoring_service | string | X | √ |  | 
| confidential_nodes | json | X | √ |  | 
| zone | string | X | √ |  | 
| location | string | X | √ |  | 
| id | string | X | √ |  | 
| enable_tpu | bool | X | √ |  | 
| resource_labels | json | X | √ |  | 
| label_fingerprint | string | X | √ |  | 
| mesh_certificates | json | X | √ |  | 
| status | big_int | X | √ |  | 
| status_message | string | X | √ |  | 
| workload_identity_config | json | X | √ |  | 
| initial_node_count | big_int | X | √ |  | 
| subnetwork | string | X | √ |  | 
| network_policy | json | X | √ |  | 
| expire_time | string | X | √ |  | 
| authenticator_groups_config | json | X | √ |  | 
| tpu_ipv4_cidr_block | string | X | √ |  | 
| selefra_id | string | √ | √ | random id | 
| network | string | X | √ |  | 
| ip_allocation_policy | json | X | √ |  | 
| cost_management_config | json | X | √ |  | 
| autopilot | json | X | √ |  | 
| maintenance_policy | json | X | √ |  | 
| autoscaling | json | X | √ |  | 
| node_ipv4_cidr_size | big_int | X | √ |  | 
| logging_config | json | X | √ |  | 
| current_node_version | string | X | √ |  | 
| node_pool_defaults | json | X | √ |  | 
| node_pool_auto_config | json | X | √ |  | 
| enable_kubernetes_alpha | bool | X | √ |  | 
| network_config | json | X | √ |  | 
| resource_usage_export_config | json | X | √ |  | 
| endpoint | string | X | √ |  | 
| initial_cluster_version | string | X | √ |  | 
| logging_service | string | X | √ |  | 
| shielded_nodes | json | X | √ |  | 
| notification_config | json | X | √ |  | 
| monitoring_config | json | X | √ |  | 
| legacy_abac | json | X | √ |  | 
| vertical_pod_autoscaling | json | X | √ |  | 
| node_config | json | X | √ |  | 
| identity_service_config | json | X | √ |  | 
| create_time | string | X | √ |  | 
| project_id | string | X | √ |  | 
| self_link | string | X | √ |  | 
| current_master_version | string | X | √ |  | 
| services_ipv4_cidr | string | X | √ |  | 
| description | string | X | √ |  | 
| cluster_ipv4_cidr | string | X | √ |  | 
| node_pools | json | X | √ |  | 
| private_cluster_config | json | X | √ |  | 
| database_encryption | json | X | √ |  | 
| instance_group_urls | string_array | X | √ |  | 
| master_auth | json | X | √ |  | 
| locations | string_array | X | √ |  | 
| binary_authorization | json | X | √ |  | 
| name | string | X | √ |  | 
| addons_config | json | X | √ |  | 
| master_authorized_networks_config | json | X | √ |  | 
| default_max_pods_constraint | json | X | √ |  | 
| conditions | json | X | √ |  | 
| release_channel | json | X | √ |  | 
| current_node_count | big_int | X | √ |  | 


