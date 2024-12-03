# GetLogsResponse

Logs


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `StatusCode`                                                | *int64*                                                     | :heavy_check_mark:                                          | HTTP Response Status Code                                   | 200                                                         |
| `Status`                                                    | *string*                                                    | :heavy_check_mark:                                          | HTTP Response Status                                        | OK                                                          |
| `Data`                                                      | [][components.Log](../../models/components/log.md)          | :heavy_check_mark:                                          | N/A                                                         |                                                             |
| `Meta`                                                      | [*components.Meta](../../models/components/meta.md)         | :heavy_minus_sign:                                          | Response metadata                                           |                                                             |
| `Links`                                                     | [*components.Links](../../models/components/links.md)       | :heavy_minus_sign:                                          | Links to navigate to previous or next pages through the API |                                                             |