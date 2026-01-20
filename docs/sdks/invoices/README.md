# Accounting.Invoices

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

<!-- UsageSnippet language="go" operationID="accounting.invoicesAll" method="get" path="/accounting/invoices" -->
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

    res, err := s.Accounting.Invoices.List(ctx, operations.AccountingInvoicesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.InvoicesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CreatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            Number: sdkgo.Pointer("OIT00546"),
        },
        Sort: &components.InvoicesSort{
            By: components.InvoicesSortByUpdatedAt.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="accounting.invoicesAdd" method="post" path="/accounting/invoices" -->
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

    res, err := s.Accounting.Invoices.Create(ctx, operations.AccountingInvoicesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Invoice: components.InvoiceInput{
            Type: components.InvoiceTypeService.ToPointer(),
            Number: sdkgo.Pointer("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            CompanyID: sdkgo.Pointer("12345"),
            InvoiceDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-09-30"),
            Terms: sdkgo.Pointer("Net 30 days"),
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("123456"),
            Status: components.InvoiceStatusDraft.ToPointer(),
            InvoiceSent: sdkgo.Pointer(true),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500),
            TotalTax: sdkgo.Pointer[float64](2500),
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25),
            Total: sdkgo.Pointer[float64](27500),
            Balance: sdkgo.Pointer[float64](27500),
            Deposit: sdkgo.Pointer[float64](0),
            CustomerMemo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
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
                    LedgerAccount: nil,
                    CustomFields: []components.CustomField{
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                                    "Uses Salesforce and Marketo",
                                )),
                            },
                        ),
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                                    "Uses Salesforce and Marketo",
                                )),
                            },
                        ),
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            BillingAddress: &components.Address{
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
            TemplateID: sdkgo.Pointer("123456"),
            SourceDocumentURL: sdkgo.Pointer("https://www.invoicesolution.com/invoice/123456"),
            PaymentAllocations: []components.PaymentAllocations{
                components.PaymentAllocations{
                    ID: sdkgo.Pointer("123456"),
                    AllocatedAmount: sdkgo.Pointer[float64](1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
            },
            PaymentMethod: sdkgo.Pointer("cash"),
            Channel: sdkgo.Pointer("email"),
            Language: sdkgo.Pointer("EN"),
            AccountingByRow: sdkgo.Pointer(false),
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
            LedgerAccount: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                            "Uses Salesforce and Marketo",
                        )),
                    },
                ),
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                            "Uses Salesforce and Marketo",
                        )),
                    },
                ),
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                            "Uses Salesforce and Marketo",
                        )),
                    },
                ),
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

<!-- UsageSnippet language="go" operationID="accounting.invoicesOne" method="get" path="/accounting/invoices/{id}" -->
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

    res, err := s.Accounting.Invoices.Get(ctx, operations.AccountingInvoicesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="accounting.invoicesUpdate" method="patch" path="/accounting/invoices/{id}" -->
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

    res, err := s.Accounting.Invoices.Update(ctx, operations.AccountingInvoicesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Invoice: components.InvoiceInput{
            Type: components.InvoiceTypeService.ToPointer(),
            Number: sdkgo.Pointer("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            CompanyID: sdkgo.Pointer("12345"),
            InvoiceDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-09-30"),
            Terms: sdkgo.Pointer("Net 30 days"),
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("123456"),
            Status: components.InvoiceStatusDraft.ToPointer(),
            InvoiceSent: sdkgo.Pointer(true),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500),
            TotalTax: sdkgo.Pointer[float64](2500),
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25),
            Total: sdkgo.Pointer[float64](27500),
            Balance: sdkgo.Pointer[float64](27500),
            Deposit: sdkgo.Pointer[float64](0),
            CustomerMemo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
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
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Name: sdkgo.Pointer("New York"),
                        },
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Name: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                    },
                    CustomFields: []components.CustomField{
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: nil,
                            },
                        ),
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: nil,
                            },
                        ),
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: nil,
                            },
                        ),
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            BillingAddress: &components.Address{
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
            TemplateID: sdkgo.Pointer("123456"),
            SourceDocumentURL: sdkgo.Pointer("https://www.invoicesolution.com/invoice/123456"),
            PaymentAllocations: []components.PaymentAllocations{
                components.PaymentAllocations{
                    ID: sdkgo.Pointer("123456"),
                    AllocatedAmount: sdkgo.Pointer[float64](1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
                components.PaymentAllocations{
                    ID: sdkgo.Pointer("123456"),
                    AllocatedAmount: sdkgo.Pointer[float64](1000),
                    Date: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
                },
            },
            PaymentMethod: sdkgo.Pointer("cash"),
            Channel: sdkgo.Pointer("email"),
            Language: sdkgo.Pointer("EN"),
            AccountingByRow: sdkgo.Pointer(false),
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
            LedgerAccount: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: nil,
                    },
                ),
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

<!-- UsageSnippet language="go" operationID="accounting.invoicesDelete" method="delete" path="/accounting/invoices/{id}" -->
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

    res, err := s.Accounting.Invoices.Delete(ctx, operations.AccountingInvoicesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
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