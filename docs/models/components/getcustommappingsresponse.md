# GetCustomMappingsResponse

Custom mapping


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `StatusCode`                                                            | *int64*                                                                 | :heavy_check_mark:                                                      | HTTP Response Status Code                                               | 200                                                                     |
| `Status`                                                                | *string*                                                                | :heavy_check_mark:                                                      | HTTP Response Status                                                    | OK                                                                      |
| `Data`                                                                  | [][components.CustomMapping](../../models/components/custommapping.md)  | :heavy_check_mark:                                                      | N/A                                                                     |                                                                         |
| `Raw`                                                                   | map[string]*any*                                                        | :heavy_minus_sign:                                                      | Raw response from the integration when raw=true query param is provided |                                                                         |