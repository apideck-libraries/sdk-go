# CreditNotes
(*Accounting.CreditNotes*)

## Overview

### Available Operations

* [List](#list) - List Credit Notes
* [Create](#create) - Create Credit Note
* [Get](#get) - Get Credit Note
* [Update](#update) - Update Credit Note
* [Delete](#delete) - Delete Credit Note

## List

List Credit Notes

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

    res, err := s.Accounting.CreditNotes.List(ctx, operations.AccountingCreditNotesAllRequest{
        ConsumerID: sdkgo.String("test-consumer"),
        AppID: sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.CreditNotesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.CreditNotesSort{
            By: components.CreditNotesSortByUpdatedAt.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCreditNotesResponse != nil {
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingCreditNotesAllRequest](../../models/operations/accountingcreditnotesallrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingCreditNotesAllResponse](../../models/operations/accountingcreditnotesallresponse.md), error**

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

Create Credit Note

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

    res, err := s.Accounting.CreditNotes.Create(ctx, operations.AccountingCreditNotesAddRequest{
        ConsumerID: sdkgo.String("test-consumer"),
        AppID: sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        ServiceID: sdkgo.String("salesforce"),
        CreditNote: components.CreditNoteInput{
            Number: sdkgo.String("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            CompanyID: sdkgo.String("12345"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalAmount: 49.99,
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            Balance: sdkgo.Float64(27500),
            RemainingCredit: sdkgo.Float64(27500),
            Status: components.CreditNoteStatusAuthorised.ToPointer(),
            Reference: sdkgo.String("123456"),
            DateIssued: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            DatePaid: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Type: components.CreditNoteTypeAccountsReceivableCredit.ToPointer(),
            Account: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
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
            Allocations: []components.AllocationInput{
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
            },
            Note: sdkgo.String("Some notes about this credit note"),
            Terms: sdkgo.String("Some terms about this credit note"),
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
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueArrayOfStr(
                        []string{
                            "<value>",
                            "<value>",
                            "<value>",
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
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingCreditNotesAddRequest](../../models/operations/accountingcreditnotesaddrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingCreditNotesAddResponse](../../models/operations/accountingcreditnotesaddresponse.md), error**

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

Get Credit Note

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

    res, err := s.Accounting.CreditNotes.Get(ctx, operations.AccountingCreditNotesOneRequest{
        ID: "<id>",
        ConsumerID: sdkgo.String("test-consumer"),
        AppID: sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingCreditNotesOneRequest](../../models/operations/accountingcreditnotesonerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingCreditNotesOneResponse](../../models/operations/accountingcreditnotesoneresponse.md), error**

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

Update Credit Note

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

    res, err := s.Accounting.CreditNotes.Update(ctx, operations.AccountingCreditNotesUpdateRequest{
        ID: "<id>",
        ConsumerID: sdkgo.String("test-consumer"),
        AppID: sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        ServiceID: sdkgo.String("salesforce"),
        CreditNote: components.CreditNoteInput{
            Number: sdkgo.String("OIT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            CompanyID: sdkgo.String("12345"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalAmount: 49.99,
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            Balance: sdkgo.Float64(27500),
            RemainingCredit: sdkgo.Float64(27500),
            Status: components.CreditNoteStatusAuthorised.ToPointer(),
            Reference: sdkgo.String("123456"),
            DateIssued: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            DatePaid: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Type: components.CreditNoteTypeAccountsReceivableCredit.ToPointer(),
            Account: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
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
            Allocations: []components.AllocationInput{
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
                components.AllocationInput{
                    ID: sdkgo.String("123456"),
                    Amount: sdkgo.Float64(49.99),
                    AllocationID: sdkgo.String("123456"),
                },
            },
            Note: sdkgo.String("Some notes about this credit note"),
            Terms: sdkgo.String("Some terms about this credit note"),
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
            CustomFields: []components.CustomField{
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
    if res.UpdateCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingCreditNotesUpdateRequest](../../models/operations/accountingcreditnotesupdaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingCreditNotesUpdateResponse](../../models/operations/accountingcreditnotesupdateresponse.md), error**

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

Delete Credit Note

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

    res, err := s.Accounting.CreditNotes.Delete(ctx, operations.AccountingCreditNotesDeleteRequest{
        ID: "<id>",
        ConsumerID: sdkgo.String("test-consumer"),
        AppID: sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingCreditNotesDeleteRequest](../../models/operations/accountingcreditnotesdeleterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingCreditNotesDeleteResponse](../../models/operations/accountingcreditnotesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |