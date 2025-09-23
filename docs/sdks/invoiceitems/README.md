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

<!-- UsageSnippet language="go" operationID="accounting.invoiceItemsAll" method="get" path="/accounting/invoice-items" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.InvoiceItems.List(ctx, operations.AccountingInvoiceItemsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.InvoiceItemsFilter{
            Name: sdkgo.Pointer("Widgets Large"),
            Type: components.InvoiceItemTypeService.ToPointer(),
        },
        Sort: &components.InvoiceItemsSort{
            By: components.InvoiceItemsSortByUpdatedAt.ToPointer(),
            Direction: components.SortDirectionDesc.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="accounting.invoiceItemsAdd" method="post" path="/accounting/invoice-items" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.InvoiceItems.Create(ctx, operations.AccountingInvoiceItemsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        InvoiceItem: components.InvoiceItemInput{
            Name: sdkgo.Pointer("Model Y"),
            Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
            Code: sdkgo.Pointer("120-C"),
            Sold: sdkgo.Pointer(true),
            Purchased: sdkgo.Pointer(true),
            Tracked: sdkgo.Pointer(true),
            Taxable: sdkgo.Pointer(true),
            InventoryDate: types.MustNewDateFromString("2020-10-30"),
            Type: components.InvoiceItemTypeTypeInventory.ToPointer(),
            SalesDetails: &components.InvoiceItemSalesDetails{
                UnitPrice: sdkgo.Pointer[float64](27500.5),
                UnitOfMeasure: sdkgo.Pointer("pc."),
                TaxInclusive: sdkgo.Pointer(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.Pointer("123456"),
                    Rate: sdkgo.Pointer[float64](10),
                },
            },
            PurchaseDetails: &components.InvoiceItemPurchaseDetails{
                UnitPrice: sdkgo.Pointer[float64](27500.5),
                UnitOfMeasure: sdkgo.Pointer("pc."),
                TaxInclusive: sdkgo.Pointer(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.Pointer("123456"),
                    Rate: sdkgo.Pointer[float64](10),
                },
            },
            Quantity: sdkgo.Pointer[float64](1),
            UnitPrice: sdkgo.Pointer[float64](27500.5),
            AssetAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            IncomeAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            ExpenseAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            TrackingCategories: nil,
            Active: sdkgo.Pointer(true),
            DepartmentID: sdkgo.Pointer("12345"),
            LocationID: sdkgo.Pointer("12345"),
            SubsidiaryID: sdkgo.Pointer("12345"),
            TaxScheduleID: sdkgo.Pointer("123456"),
            RowVersion: sdkgo.Pointer("1-12345"),
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

<!-- UsageSnippet language="go" operationID="accounting.invoiceItemsOne" method="get" path="/accounting/invoice-items/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.InvoiceItems.Get(ctx, operations.AccountingInvoiceItemsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="accounting.invoiceItemsUpdate" method="patch" path="/accounting/invoice-items/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.InvoiceItems.Update(ctx, operations.AccountingInvoiceItemsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        InvoiceItem: components.InvoiceItemInput{
            Name: sdkgo.Pointer("Model Y"),
            Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
            Code: sdkgo.Pointer("120-C"),
            Sold: sdkgo.Pointer(true),
            Purchased: sdkgo.Pointer(true),
            Tracked: sdkgo.Pointer(true),
            Taxable: sdkgo.Pointer(true),
            InventoryDate: types.MustNewDateFromString("2020-10-30"),
            Type: components.InvoiceItemTypeTypeInventory.ToPointer(),
            SalesDetails: &components.InvoiceItemSalesDetails{
                UnitPrice: sdkgo.Pointer[float64](27500.5),
                UnitOfMeasure: sdkgo.Pointer("pc."),
                TaxInclusive: sdkgo.Pointer(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.Pointer("123456"),
                    Rate: sdkgo.Pointer[float64](10),
                },
            },
            PurchaseDetails: &components.InvoiceItemPurchaseDetails{
                UnitPrice: sdkgo.Pointer[float64](27500.5),
                UnitOfMeasure: sdkgo.Pointer("pc."),
                TaxInclusive: sdkgo.Pointer(true),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.Pointer("123456"),
                    Rate: sdkgo.Pointer[float64](10),
                },
            },
            Quantity: sdkgo.Pointer[float64](1),
            UnitPrice: sdkgo.Pointer[float64](27500.5),
            AssetAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            IncomeAccount: nil,
            ExpenseAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
            Active: sdkgo.Pointer(true),
            DepartmentID: sdkgo.Pointer("12345"),
            LocationID: sdkgo.Pointer("12345"),
            SubsidiaryID: sdkgo.Pointer("12345"),
            TaxScheduleID: sdkgo.Pointer("123456"),
            RowVersion: sdkgo.Pointer("1-12345"),
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

<!-- UsageSnippet language="go" operationID="accounting.invoiceItemsDelete" method="delete" path="/accounting/invoice-items/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.InvoiceItems.Delete(ctx, operations.AccountingInvoiceItemsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
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