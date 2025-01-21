# InvoiceItems
(*Accounting.InvoiceItems*)

## Overview

### Available Operations

* [List](#list) - List Invoice Items
* [Create](#create) - Create Invoice Item
* [Get](#get) - Get Invoice Item
* [Update](#update) - Update Invoice Item
* [Delete](#delete) - Delete Invoice Item

## List

List Invoice Items

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.InvoiceItems.List(ctx, operations.AccountingInvoiceItemsAllRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        Limit: sdkgo.Int64(20),
        Filter: &components.InvoiceItemsFilter{
            Name: sdkgo.String("Widgets Large"),
            Type: components.InvoiceItemTypeService.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetInvoiceItemsResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingInvoiceItemsAllRequest](../../models/operations/accountinginvoiceitemsallrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingInvoiceItemsAllResponse](../../models/operations/accountinginvoiceitemsallresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Create

Create Invoice Item

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.InvoiceItems.Create(ctx, operations.AccountingInvoiceItemsAddRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        InvoiceItem: components.InvoiceItemInput{
            Name: sdkgo.String("Model Y"),
            Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
            Code: sdkgo.String("120-C"),
            Sold: sdkgo.Bool(true),
            Purchased: sdkgo.Bool(true),
            Tracked: sdkgo.Bool(true),
            Taxable: sdkgo.Bool(true),
            InventoryDate: types.MustNewDateFromString("2020-10-30"),
            Type: components.InvoiceItemTypeTypeInventory.ToPointer(),
            SalesDetails: &components.InvoiceItemSalesDetails{
                UnitPrice: sdkgo.Float64(27500.5),
                UnitOfMeasure: sdkgo.String("pc."),
                TaxInclusive: sdkgo.Bool(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
            },
            PurchaseDetails: &components.InvoiceItemPurchaseDetails{
                UnitPrice: sdkgo.Float64(27500.5),
                UnitOfMeasure: sdkgo.String("pc."),
                TaxInclusive: sdkgo.Bool(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
            },
            Quantity: sdkgo.Float64(1),
            UnitPrice: sdkgo.Float64(27500.5),
            AssetAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            IncomeAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            ExpenseAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            TrackingCategories: []components.LinkedTrackingCategory{
                components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
                components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
            Active: sdkgo.Bool(true),
            RowVersion: sdkgo.String("1-12345"),
            PassThrough: []components.PassThroughBody{
                components.PassThroughBody{
                    ServiceID: "<id>",
                    ExtendPaths: []components.ExtendPaths{
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                    },
                },
                components.PassThroughBody{
                    ServiceID: "<id>",
                    ExtendPaths: []components.ExtendPaths{
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateInvoiceItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingInvoiceItemsAddRequest](../../models/operations/accountinginvoiceitemsaddrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingInvoiceItemsAddResponse](../../models/operations/accountinginvoiceitemsaddresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Get

Get Invoice Item

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.InvoiceItems.Get(ctx, operations.AccountingInvoiceItemsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        Fields: sdkgo.String("id,updated_at"),
        Filter: &components.InvoiceItemFilter{
            Type: components.InvoiceItemFilterInvoiceItemTypeService.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetInvoiceItemResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingInvoiceItemsOneRequest](../../models/operations/accountinginvoiceitemsonerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingInvoiceItemsOneResponse](../../models/operations/accountinginvoiceitemsoneresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Update

Update Invoice Item

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.InvoiceItems.Update(ctx, operations.AccountingInvoiceItemsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        InvoiceItem: components.InvoiceItemInput{
            Name: sdkgo.String("Model Y"),
            Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
            Code: sdkgo.String("120-C"),
            Sold: sdkgo.Bool(true),
            Purchased: sdkgo.Bool(true),
            Tracked: sdkgo.Bool(true),
            Taxable: sdkgo.Bool(true),
            InventoryDate: types.MustNewDateFromString("2020-10-30"),
            Type: components.InvoiceItemTypeTypeInventory.ToPointer(),
            SalesDetails: &components.InvoiceItemSalesDetails{
                UnitPrice: sdkgo.Float64(27500.5),
                UnitOfMeasure: sdkgo.String("pc."),
                TaxInclusive: sdkgo.Bool(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
            },
            PurchaseDetails: &components.InvoiceItemPurchaseDetails{
                UnitPrice: sdkgo.Float64(27500.5),
                UnitOfMeasure: sdkgo.String("pc."),
                TaxInclusive: sdkgo.Bool(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
            },
            Quantity: sdkgo.Float64(1),
            UnitPrice: sdkgo.Float64(27500.5),
            AssetAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            IncomeAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            ExpenseAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            TrackingCategories: []components.LinkedTrackingCategory{
                components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
                components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
                components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
            Active: sdkgo.Bool(true),
            RowVersion: sdkgo.String("1-12345"),
            PassThrough: []components.PassThroughBody{
                components.PassThroughBody{
                    ServiceID: "<id>",
                    ExtendPaths: []components.ExtendPaths{
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                    },
                },
                components.PassThroughBody{
                    ServiceID: "<id>",
                    ExtendPaths: []components.ExtendPaths{
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                        components.ExtendPaths{
                            Path: "$.nested.property",
                            Value: map[string]any{
                                "TaxClassificationRef": map[string]any{
                                    "value": "EUC-99990201-V1-00020000",
                                },
                            },
                        },
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateInvoiceItemsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingInvoiceItemsUpdateRequest](../../models/operations/accountinginvoiceitemsupdaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingInvoiceItemsUpdateResponse](../../models/operations/accountinginvoiceitemsupdateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete Invoice Item

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.InvoiceItems.Delete(ctx, operations.AccountingInvoiceItemsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteTaxRateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingInvoiceItemsDeleteRequest](../../models/operations/accountinginvoiceitemsdeleterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingInvoiceItemsDeleteResponse](../../models/operations/accountinginvoiceitemsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |