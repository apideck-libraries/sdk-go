# GetSalesOrderResponse

Sales Orders


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `StatusCode`                                                   | `int64`                                                        | :heavy_check_mark:                                             | HTTP Response Status Code                                      | 200                                                            |
| `Status`                                                       | `string`                                                       | :heavy_check_mark:                                             | HTTP Response Status                                           | OK                                                             |
| `Service`                                                      | `string`                                                       | :heavy_check_mark:                                             | Apideck ID of service provider                                 | acumatica                                                      |
| `Resource`                                                     | `string`                                                       | :heavy_check_mark:                                             | Unified API resource name                                      | SalesOrders                                                    |
| `Operation`                                                    | `string`                                                       | :heavy_check_mark:                                             | Operation performed                                            | one                                                            |
| `Data`                                                         | [components.SalesOrder](../../models/components/salesorder.md) | :heavy_check_mark:                                             | N/A                                                            |                                                                |
| `Meta`                                                         | [*components.Meta](../../models/components/meta.md)            | :heavy_minus_sign:                                             | Response metadata                                              |                                                                |