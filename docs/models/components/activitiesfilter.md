# ActivitiesFilter


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `CompanyID`                                | **string*                                  | :heavy_minus_sign:                         | Company ID to filter on                    | 1234                                       |
| `OwnerID`                                  | **string*                                  | :heavy_minus_sign:                         | Owner ID to filter on                      | 1234                                       |
| `ContactID`                                | **string*                                  | :heavy_minus_sign:                         | Primary contact ID to filter on            | 1234                                       |
| `UpdatedSince`                             | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        | 2020-09-30T07:43:32.000Z                   |
| `Type`                                     | **string*                                  | :heavy_minus_sign:                         | Type to filter on                          | Task                                       |