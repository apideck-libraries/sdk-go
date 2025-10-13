# PurchaseOrders
(*Accounting.PurchaseOrders*)

## Overview

### Available Operations

* [List](#list) - List Purchase Orders
* [Create](#create) - Create Purchase Order
* [Get](#get) - Get Purchase Order
* [Update](#update) - Update Purchase Order
* [Delete](#delete) - Delete Purchase Order

## List

List Purchase Orders

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.purchaseOrdersAll" method="get" path="/accounting/purchase-orders" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.PurchaseOrders.List(ctx, operations.AccountingPurchaseOrdersAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Filter: &components.PurchaseOrdersFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            SupplierID: sdkgo.Pointer("1234"),
        },
        Sort: &components.PurchaseOrdersSort{
            By: components.PurchaseOrdersSortByUpdatedAt.ToPointer(),
            Direction: components.SortDirectionDesc.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPurchaseOrdersResponse != nil {
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingPurchaseOrdersAllRequest](../../models/operations/accountingpurchaseordersallrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingPurchaseOrdersAllResponse](../../models/operations/accountingpurchaseordersallresponse.md), error**

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

Create Purchase Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.purchaseOrdersAdd" method="post" path="/accounting/purchase-orders" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Accounting.PurchaseOrders.Create(ctx, operations.AccountingPurchaseOrdersAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        PurchaseOrder: components.PurchaseOrderInput{
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("123456"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.Pointer("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.Pointer("HQ US"),
                    Line1: sdkgo.Pointer("Main street"),
                    Line2: sdkgo.Pointer("apt #"),
                    Line3: sdkgo.Pointer("Suite #"),
                    Line4: sdkgo.Pointer("delivery instructions"),
                    StreetNumber: sdkgo.Pointer("25"),
                    City: sdkgo.Pointer("San Francisco"),
                    State: sdkgo.Pointer("CA"),
                    PostalCode: sdkgo.Pointer("94104"),
                    Country: sdkgo.Pointer("US"),
                    Latitude: sdkgo.Pointer("40.759211"),
                    Longitude: sdkgo.Pointer("-73.984638"),
                    County: sdkgo.Pointer("Santa Clara"),
                    ContactName: sdkgo.Pointer("Elon Musk"),
                    Salutation: sdkgo.Pointer("Mr"),
                    PhoneNumber: sdkgo.Pointer("111-111-1111"),
                    Fax: sdkgo.Pointer("122-111-1111"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Website: sdkgo.Pointer("https://elonmusk.com"),
                    Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            SubsidiaryID: sdkgo.Pointer("12345"),
            CompanyID: sdkgo.Pointer("12345"),
            Status: components.PurchaseOrderStatusOpen.ToPointer(),
            IssuedDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-09-30"),
            ExpectedArrivalDate: types.MustNewDateFromString("2020-09-30"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            SubTotal: sdkgo.Pointer[float64](27500),
            TotalTax: sdkgo.Pointer[float64](2500),
            Total: sdkgo.Pointer[float64](27500),
            TaxInclusive: sdkgo.Pointer(true),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Quantity: sdkgo.Pointer[float64](1),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Rate: sdkgo.Pointer[float64](10),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Name: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Quantity: sdkgo.Pointer[float64](1),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Rate: sdkgo.Pointer[float64](10),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Name: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Quantity: sdkgo.Pointer[float64](1),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Rate: sdkgo.Pointer[float64](10),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Name: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            ShippingAddress: &components.Address{
                ID: sdkgo.Pointer("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.Pointer("HQ US"),
                Line1: sdkgo.Pointer("Main street"),
                Line2: sdkgo.Pointer("apt #"),
                Line3: sdkgo.Pointer("Suite #"),
                Line4: sdkgo.Pointer("delivery instructions"),
                StreetNumber: sdkgo.Pointer("25"),
                City: sdkgo.Pointer("San Francisco"),
                State: sdkgo.Pointer("CA"),
                PostalCode: sdkgo.Pointer("94104"),
                Country: sdkgo.Pointer("US"),
                Latitude: sdkgo.Pointer("40.759211"),
                Longitude: sdkgo.Pointer("-73.984638"),
                County: sdkgo.Pointer("Santa Clara"),
                ContactName: sdkgo.Pointer("Elon Musk"),
                Salutation: sdkgo.Pointer("Mr"),
                PhoneNumber: sdkgo.Pointer("111-111-1111"),
                Fax: sdkgo.Pointer("122-111-1111"),
                Email: sdkgo.Pointer("elon@musk.com"),
                Website: sdkgo.Pointer("https://elonmusk.com"),
                Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                RowVersion: sdkgo.Pointer("1-12345"),
            },
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            TemplateID: sdkgo.Pointer("123456"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.Pointer("Monzo"),
                AccountNumber: sdkgo.Pointer("123465"),
                AccountName: sdkgo.Pointer("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.Pointer("CH2989144532982975332"),
                Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.Pointer("012345678"),
                BsbNumber: sdkgo.Pointer("062-001"),
                BranchIdentifier: sdkgo.Pointer("001"),
                BankCode: sdkgo.Pointer("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            AccountingByRow: sdkgo.Pointer(false),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaymentMethod: sdkgo.Pointer("cash"),
            TaxCode: sdkgo.Pointer("1234"),
            Channel: sdkgo.Pointer("email"),
            Memo: sdkgo.Pointer("Thank you for the partnership and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingPurchaseOrdersAddRequest](../../models/operations/accountingpurchaseordersaddrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingPurchaseOrdersAddResponse](../../models/operations/accountingpurchaseordersaddresponse.md), error**

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

Get Purchase Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.purchaseOrdersOne" method="get" path="/accounting/purchase-orders/{id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.PurchaseOrders.Get(ctx, operations.AccountingPurchaseOrdersOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetPurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingPurchaseOrdersOneRequest](../../models/operations/accountingpurchaseordersonerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingPurchaseOrdersOneResponse](../../models/operations/accountingpurchaseordersoneresponse.md), error**

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

Update Purchase Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.purchaseOrdersUpdate" method="patch" path="/accounting/purchase-orders/{id}" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Accounting.PurchaseOrders.Update(ctx, operations.AccountingPurchaseOrdersUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        PurchaseOrder: components.PurchaseOrderInput{
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("123456"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.Pointer("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.Pointer("HQ US"),
                    Line1: sdkgo.Pointer("Main street"),
                    Line2: sdkgo.Pointer("apt #"),
                    Line3: sdkgo.Pointer("Suite #"),
                    Line4: sdkgo.Pointer("delivery instructions"),
                    StreetNumber: sdkgo.Pointer("25"),
                    City: sdkgo.Pointer("San Francisco"),
                    State: sdkgo.Pointer("CA"),
                    PostalCode: sdkgo.Pointer("94104"),
                    Country: sdkgo.Pointer("US"),
                    Latitude: sdkgo.Pointer("40.759211"),
                    Longitude: sdkgo.Pointer("-73.984638"),
                    County: sdkgo.Pointer("Santa Clara"),
                    ContactName: sdkgo.Pointer("Elon Musk"),
                    Salutation: sdkgo.Pointer("Mr"),
                    PhoneNumber: sdkgo.Pointer("111-111-1111"),
                    Fax: sdkgo.Pointer("122-111-1111"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Website: sdkgo.Pointer("https://elonmusk.com"),
                    Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            SubsidiaryID: sdkgo.Pointer("12345"),
            CompanyID: sdkgo.Pointer("12345"),
            Status: components.PurchaseOrderStatusOpen.ToPointer(),
            IssuedDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-09-30"),
            ExpectedArrivalDate: types.MustNewDateFromString("2020-09-30"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            SubTotal: sdkgo.Pointer[float64](27500),
            TotalTax: sdkgo.Pointer[float64](2500),
            Total: sdkgo.Pointer[float64](27500),
            TaxInclusive: sdkgo.Pointer(true),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Quantity: sdkgo.Pointer[float64](1),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Rate: sdkgo.Pointer[float64](10),
                    },
                    TrackingCategories: nil,
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Quantity: sdkgo.Pointer[float64](1),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Rate: sdkgo.Pointer[float64](10),
                    },
                    TrackingCategories: nil,
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.Pointer("2389328923893298"),
                            Name: sdkgo.Pointer("employee_level"),
                            Description: sdkgo.Pointer("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            ShippingAddress: &components.Address{
                ID: sdkgo.Pointer("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.Pointer("HQ US"),
                Line1: sdkgo.Pointer("Main street"),
                Line2: sdkgo.Pointer("apt #"),
                Line3: sdkgo.Pointer("Suite #"),
                Line4: sdkgo.Pointer("delivery instructions"),
                StreetNumber: sdkgo.Pointer("25"),
                City: sdkgo.Pointer("San Francisco"),
                State: sdkgo.Pointer("CA"),
                PostalCode: sdkgo.Pointer("94104"),
                Country: sdkgo.Pointer("US"),
                Latitude: sdkgo.Pointer("40.759211"),
                Longitude: sdkgo.Pointer("-73.984638"),
                County: sdkgo.Pointer("Santa Clara"),
                ContactName: sdkgo.Pointer("Elon Musk"),
                Salutation: sdkgo.Pointer("Mr"),
                PhoneNumber: sdkgo.Pointer("111-111-1111"),
                Fax: sdkgo.Pointer("122-111-1111"),
                Email: sdkgo.Pointer("elon@musk.com"),
                Website: sdkgo.Pointer("https://elonmusk.com"),
                Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                RowVersion: sdkgo.Pointer("1-12345"),
            },
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            TemplateID: sdkgo.Pointer("123456"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.Pointer("Monzo"),
                AccountNumber: sdkgo.Pointer("123465"),
                AccountName: sdkgo.Pointer("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.Pointer("CH2989144532982975332"),
                Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.Pointer("012345678"),
                BsbNumber: sdkgo.Pointer("062-001"),
                BranchIdentifier: sdkgo.Pointer("001"),
                BankCode: sdkgo.Pointer("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            AccountingByRow: sdkgo.Pointer(false),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaymentMethod: sdkgo.Pointer("cash"),
            TaxCode: sdkgo.Pointer("1234"),
            Channel: sdkgo.Pointer("email"),
            Memo: sdkgo.Pointer("Thank you for the partnership and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
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
    if res.UpdatePurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingPurchaseOrdersUpdateRequest](../../models/operations/accountingpurchaseordersupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingPurchaseOrdersUpdateResponse](../../models/operations/accountingpurchaseordersupdateresponse.md), error**

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

Delete Purchase Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.purchaseOrdersDelete" method="delete" path="/accounting/purchase-orders/{id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Accounting.PurchaseOrders.Delete(ctx, operations.AccountingPurchaseOrdersDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeletePurchaseOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingPurchaseOrdersDeleteRequest](../../models/operations/accountingpurchaseordersdeleterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingPurchaseOrdersDeleteResponse](../../models/operations/accountingpurchaseordersdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |