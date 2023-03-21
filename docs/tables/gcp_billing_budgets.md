# Table: gcp_billing_budgets

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| notifications_rule | json | X | √ |  | 
| gcp_billing_billing_accounts_selefra_id | string | X | X | fk to gcp_billing_billing_accounts.selefra_id | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| budget_filter | json | X | √ |  | 
| threshold_rules | json | X | √ |  | 
| amount | json | X | √ |  | 
| etag | string | X | √ |  | 
| project_id | string | X | √ |  | 
| name | string | X | √ |  | 
| display_name | string | X | √ |  | 


