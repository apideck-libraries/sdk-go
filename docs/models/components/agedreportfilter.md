# AgedReportFilter


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CustomerID`                                              | **string*                                                 | :heavy_minus_sign:                                        | Filter by customer id                                     | 123abc                                                    |
| `SupplierID`                                              | **string*                                                 | :heavy_minus_sign:                                        | Filter by supplier id                                     | 123abc                                                    |
| `ReportAsOfDate`                                          | **string*                                                 | :heavy_minus_sign:                                        | The cutoff date for considering transactions              | 2024-01-01                                                |
| `PeriodCount`                                             | **int64*                                                  | :heavy_minus_sign:                                        | Number of periods to split the aged creditors report into | 3                                                         |
| `PeriodLength`                                            | **int64*                                                  | :heavy_minus_sign:                                        | Length of each period in days                             | 30                                                        |