# InvoicesFilter


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `UpdatedSince`                             | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        | 2020-09-30T07:43:32.000Z                   |
| `CreatedSince`                             | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        | 2020-09-30T07:43:32.000Z                   |
| `Number`                                   | **string*                                  | :heavy_minus_sign:                         | Invoice number to search for               | OIT00546                                   |
| `SupplierID`                               | **string*                                  | :heavy_minus_sign:                         | Supplier ID to filter invoices by          | 123                                        |