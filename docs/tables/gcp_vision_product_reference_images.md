# Table: gcp_vision_product_reference_images

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| bounding_polys | json | X | √ |  | 
| project_id | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| gcp_vision_products_selefra_id | string | X | X | fk to gcp_vision_products.selefra_id | 
| name | string | X | √ |  | 
| uri | string | X | √ |  | 


