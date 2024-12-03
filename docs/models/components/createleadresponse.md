# CreateLeadResponse

Lead created


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `StatusCode`                                                 | *int64*                                                      | :heavy_check_mark:                                           | HTTP Response Status Code                                    | 200                                                          |
| `Status`                                                     | *string*                                                     | :heavy_check_mark:                                           | HTTP Response Status                                         | OK                                                           |
| `Service`                                                    | *string*                                                     | :heavy_check_mark:                                           | Apideck ID of service provider                               | zoho-crm                                                     |
| `Resource`                                                   | *string*                                                     | :heavy_check_mark:                                           | Unified API resource name                                    | companies                                                    |
| `Operation`                                                  | *string*                                                     | :heavy_check_mark:                                           | Operation performed                                          | add                                                          |
| `Data`                                                       | [components.UnifiedID](../../models/components/unifiedid.md) | :heavy_check_mark:                                           | N/A                                                          |                                                              |