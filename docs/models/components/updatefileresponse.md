# UpdateFileResponse

Files


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `StatusCode`                                                 | *int64*                                                      | :heavy_check_mark:                                           | HTTP Response Status Code                                    | 200                                                          |
| `Status`                                                     | *string*                                                     | :heavy_check_mark:                                           | HTTP Response Status                                         | OK                                                           |
| `Service`                                                    | *string*                                                     | :heavy_check_mark:                                           | Apideck ID of service provider                               | google-drive                                                 |
| `Resource`                                                   | *string*                                                     | :heavy_check_mark:                                           | Unified API resource name                                    | files                                                        |
| `Operation`                                                  | *string*                                                     | :heavy_check_mark:                                           | Operation performed                                          | update                                                       |
| `Data`                                                       | [components.UnifiedID](../../models/components/unifiedid.md) | :heavy_check_mark:                                           | N/A                                                          |                                                              |