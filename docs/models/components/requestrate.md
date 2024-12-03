# RequestRate

The rate at which requests for resources will be made to downstream.


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Rate`                                             | *int64*                                            | :heavy_check_mark:                                 | The number of requests per window unit.            |
| `Size`                                             | *int64*                                            | :heavy_check_mark:                                 | Size of request window.                            |
| `Unit`                                             | [components.Unit](../../models/components/unit.md) | :heavy_check_mark:                                 | The window unit for the rate.                      |