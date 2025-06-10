# Invoices
(*Accounting.Invoices*)

## Overview

### Available Operations

* [List](#list) - List Invoices
* [Create](#create) - Create Invoice
* [Get](#get) - Get Invoice
* [Update](#update) - Update Invoice
* [Delete](#delete) - Delete Invoice

## List

List Invoices

### Example Usage

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

    res, err := s.Accounting.Invoices.List(ctx, operations.AccountingInvoicesAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.InvoicesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CreatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            Number: sdkgo.String("OIT00546"),
        },
        Sort: &components.InvoicesSort{
            By: components.InvoicesSortByUpdatedAt.ToPointer(),
            Direction: components.SortDirectionDesc.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetInvoicesResponse != nil {
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingInvoicesAllRequest](../../models/operations/accountinginvoicesallrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingInvoicesAllResponse](../../models/operations/accountinginvoicesallresponse.md), error**

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

Create Invoice

### Example Usage

```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Accounting.Invoices.Create(ctx, operations.AccountingInvoicesAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        Invoice: components.InvoiceInput{
            Type: components.InvoiceTypeService.ToPointer(),
            Number: sdkgo.String("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            CompanyID: sdkgo.String("12345"),
            InvoiceDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-09-30"),
            Terms: sdkgo.String("Net 30 days"),
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            Status: components.InvoiceStatusDraft.ToPointer(),
            InvoiceSent: sdkgo.Bool(true),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            DiscountPercentage: sdkgo.Float64(5.5),
            DiscountAmount: sdkgo.Float64(25),
            Total: sdkgo.Float64(27500),
            Balance: sdkgo.Float64(27500),
            Deposit: sdkgo.Float64(0),
            CustomerMemo: sdkgo.String("Thank you for your business and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
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
                    LocationID: sdkgo.String("12345"),
                    DepartmentID: sdkgo.String("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.String("123456"),
                            Name: sdkgo.String("New York"),
                        },
                    },
                    LedgerAccount: nil,
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
            },
            BillingAddress: &components.Address{
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
            TemplateID: sdkgo.String("123456"),
            SourceDocumentURL: sdkgo.String("https://www.invoicesolution.com/invoice/123456"),
            PaymentAllocations: []components.PaymentAllocations{
                components.PaymentAllocations{
                    ID: sdkgo.String("123456"),
                    AllocatedAmount: sdkgo.Float64(1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
            },
            PaymentMethod: sdkgo.String("cash"),
            Channel: sdkgo.String("email"),
            Language: sdkgo.String("EN"),
            AccountingByRow: sdkgo.Bool(false),
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
    if res.CreateInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingInvoicesAddRequest](../../models/operations/accountinginvoicesaddrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingInvoicesAddResponse](../../models/operations/accountinginvoicesaddresponse.md), error**

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

Get Invoice

### Example Usage

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

    res, err := s.Accounting.Invoices.Get(ctx, operations.AccountingInvoicesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingInvoicesOneRequest](../../models/operations/accountinginvoicesonerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingInvoicesOneResponse](../../models/operations/accountinginvoicesoneresponse.md), error**

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

Update Invoice

### Example Usage

```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Accounting.Invoices.Update(ctx, operations.AccountingInvoicesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Invoice: components.InvoiceInput{
            Type: components.InvoiceTypeService.ToPointer(),
            Number: sdkgo.String("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            CompanyID: sdkgo.String("12345"),
            InvoiceDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-09-30"),
            Terms: sdkgo.String("Net 30 days"),
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            Status: components.InvoiceStatusDraft.ToPointer(),
            InvoiceSent: sdkgo.Bool(true),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            DiscountPercentage: sdkgo.Float64(5.5),
            DiscountAmount: sdkgo.Float64(25),
            Total: sdkgo.Float64(27500),
            Balance: sdkgo.Float64(27500),
            Deposit: sdkgo.Float64(0),
            CustomerMemo: sdkgo.String("Thank you for your business and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
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
                    LocationID: sdkgo.String("12345"),
                    DepartmentID: sdkgo.String("12345"),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.String("12344"),
                        Code: sdkgo.String("120-C"),
                        Name: sdkgo.String("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.String("123456"),
                        Rate: sdkgo.Float64(10),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.String("123456"),
                            Name: sdkgo.String("New York"),
                        },
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.String("123456"),
                            Name: sdkgo.String("New York"),
                        },
                        &components.LinkedTrackingCategory{
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
                            Value: nil,
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: nil,
                        },
                        components.CustomField{
                            ID: sdkgo.String("2389328923893298"),
                            Name: sdkgo.String("employee_level"),
                            Description: sdkgo.String("Employee Level"),
                            Value: nil,
                        },
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            BillingAddress: &components.Address{
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
            TemplateID: sdkgo.String("123456"),
            SourceDocumentURL: sdkgo.String("https://www.invoicesolution.com/invoice/123456"),
            PaymentAllocations: []components.PaymentAllocations{
                components.PaymentAllocations{
                    ID: sdkgo.String("123456"),
                    AllocatedAmount: sdkgo.Float64(1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
                components.PaymentAllocations{
                    ID: sdkgo.String("123456"),
                    AllocatedAmount: sdkgo.Float64(1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
            },
            PaymentMethod: sdkgo.String("cash"),
            Channel: sdkgo.String("email"),
            Language: sdkgo.String("EN"),
            AccountingByRow: sdkgo.Bool(false),
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
                    Value: nil,
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
    if res.UpdateInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingInvoicesUpdateRequest](../../models/operations/accountinginvoicesupdaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingInvoicesUpdateResponse](../../models/operations/accountinginvoicesupdateresponse.md), error**

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

Delete Invoice

### Example Usage

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

    res, err := s.Accounting.Invoices.Delete(ctx, operations.AccountingInvoicesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteInvoiceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingInvoicesDeleteRequest](../../models/operations/accountinginvoicesdeleterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingInvoicesDeleteResponse](../../models/operations/accountinginvoicesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |