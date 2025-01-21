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

    res, err := s.Accounting.PurchaseOrders.List(ctx, operations.AccountingPurchaseOrdersAllRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Limit: sdkgo.Int64(20),
        Filter: &components.PurchaseOrdersFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            SupplierID: sdkgo.String("1234"),
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
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.PurchaseOrders.Create(ctx, operations.AccountingPurchaseOrdersAddRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        PurchaseOrder: components.PurchaseOrderInput{
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            CompanyID: sdkgo.String("12345"),
            Status: components.PurchaseOrderStatusOpen.ToPointer(),
            IssuedDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-09-30"),
            ExpectedArrivalDate: types.MustNewDateFromString("2020-09-30"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            Total: sdkgo.Float64(27500),
            TaxInclusive: sdkgo.Bool(true),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    LocationID: sdkgo.String("1234"),
                    DepartmentID: sdkgo.String("1234"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    LocationID: sdkgo.String("1234"),
                    DepartmentID: sdkgo.String("1234"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueNumber(
                                10,
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            ShippingAddress: &components.Address{
                ID: sdkgo.String("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.String("HQ US"),
                Line1: sdkgo.String("Main street"),
                Line2: sdkgo.String("apt #"),
                Line3: sdkgo.String("Suite #"),
                Line4: sdkgo.String("delivery instructions"),
                StreetNumber: sdkgo.String("25"),
                City: sdkgo.String("San Francisco"),
                State: sdkgo.String("CA"),
                PostalCode: sdkgo.String("94104"),
                Country: sdkgo.String("US"),
                Latitude: sdkgo.String("40.759211"),
                Longitude: sdkgo.String("-73.984638"),
                County: sdkgo.String("Santa Clara"),
                ContactName: sdkgo.String("Elon Musk"),
                Salutation: sdkgo.String("Mr"),
                PhoneNumber: sdkgo.String("111-111-1111"),
                Fax: sdkgo.String("122-111-1111"),
                Email: sdkgo.String("elon@musk.com"),
                Website: sdkgo.String("https://elonmusk.com"),
                Notes: sdkgo.String("Address notes or delivery instructions."),
                RowVersion: sdkgo.String("1-12345"),
            },
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            TemplateID: sdkgo.String("123456"),
            DiscountPercentage: sdkgo.Float64(5.5),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.String("Monzo"),
                AccountNumber: sdkgo.String("123465"),
                AccountName: sdkgo.String("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.String("CH2989144532982975332"),
                Bic: sdkgo.String("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.String("012345678"),
                BsbNumber: sdkgo.String("062-001"),
                BranchIdentifier: sdkgo.String("001"),
                BankCode: sdkgo.String("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            AccountingByRow: sdkgo.Bool(false),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaymentMethod: sdkgo.String("cash"),
            TaxCode: sdkgo.String("1234"),
            Channel: sdkgo.String("email"),
            Memo: sdkgo.String("Thank you for the partnership and have a great day!"),
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

    res, err := s.Accounting.PurchaseOrders.Get(ctx, operations.AccountingPurchaseOrdersOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
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
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.PurchaseOrders.Update(ctx, operations.AccountingPurchaseOrdersUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        PurchaseOrder: components.PurchaseOrderInput{
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            CompanyID: sdkgo.String("12345"),
            Status: components.PurchaseOrderStatusOpen.ToPointer(),
            IssuedDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-09-30"),
            ExpectedArrivalDate: types.MustNewDateFromString("2020-09-30"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            Total: sdkgo.Float64(27500),
            TaxInclusive: sdkgo.Bool(true),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    LocationID: sdkgo.String("1234"),
                    DepartmentID: sdkgo.String("1234"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueArrayOf6(
                                []components.Six{
                                    components.Six{},
                                    components.Six{},
                                },
                            )),
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueBoolean(
                                true,
                            )),
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    LocationID: sdkgo.String("1234"),
                    DepartmentID: sdkgo.String("1234"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueArrayOf6(
                                []components.Six{
                                    components.Six{},
                                },
                            )),
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.InvoiceLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    LocationID: sdkgo.String("1234"),
                    DepartmentID: sdkgo.String("1234"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
                    },
                    TrackingCategories: []components.LinkedTrackingCategory{
                        components.LinkedTrackingCategory{
                            ID: sdkgo.String("123456"),
                            Name: sdkgo.String("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: sdkgo.Pointer(components.CreateValueStr(
                                "Uses Salesforce and Marketo",
                            )),
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            ShippingAddress: &components.Address{
                ID: sdkgo.String("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.String("HQ US"),
                Line1: sdkgo.String("Main street"),
                Line2: sdkgo.String("apt #"),
                Line3: sdkgo.String("Suite #"),
                Line4: sdkgo.String("delivery instructions"),
                StreetNumber: sdkgo.String("25"),
                City: sdkgo.String("San Francisco"),
                State: sdkgo.String("CA"),
                PostalCode: sdkgo.String("94104"),
                Country: sdkgo.String("US"),
                Latitude: sdkgo.String("40.759211"),
                Longitude: sdkgo.String("-73.984638"),
                County: sdkgo.String("Santa Clara"),
                ContactName: sdkgo.String("Elon Musk"),
                Salutation: sdkgo.String("Mr"),
                PhoneNumber: sdkgo.String("111-111-1111"),
                Fax: sdkgo.String("122-111-1111"),
                Email: sdkgo.String("elon@musk.com"),
                Website: sdkgo.String("https://elonmusk.com"),
                Notes: sdkgo.String("Address notes or delivery instructions."),
                RowVersion: sdkgo.String("1-12345"),
            },
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
            },
            TemplateID: sdkgo.String("123456"),
            DiscountPercentage: sdkgo.Float64(5.5),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.String("Monzo"),
                AccountNumber: sdkgo.String("123465"),
                AccountName: sdkgo.String("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.String("CH2989144532982975332"),
                Bic: sdkgo.String("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.String("012345678"),
                BsbNumber: sdkgo.String("062-001"),
                BranchIdentifier: sdkgo.String("001"),
                BankCode: sdkgo.String("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            AccountingByRow: sdkgo.Bool(false),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaymentMethod: sdkgo.String("cash"),
            TaxCode: sdkgo.String("1234"),
            Channel: sdkgo.String("email"),
            Memo: sdkgo.String("Thank you for the partnership and have a great day!"),
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

    res, err := s.Accounting.PurchaseOrders.Delete(ctx, operations.AccountingPurchaseOrdersDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
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