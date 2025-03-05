# Bills
(*Accounting.Bills*)

## Overview

### Available Operations

* [List](#list) - List Bills
* [Create](#create) - Create Bill
* [Get](#get) - Get Bill
* [Update](#update) - Update Bill
* [Delete](#delete) - Delete Bill

## List

List Bills

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

    res, err := s.Accounting.Bills.List(ctx, operations.AccountingBillsAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.BillsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.BillsSort{
            By: components.ByUpdatedAt.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBillsResponse != nil {
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AccountingBillsAllRequest](../../models/operations/accountingbillsallrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AccountingBillsAllResponse](../../models/operations/accountingbillsallresponse.md), error**

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

Create Bill

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

    res, err := s.Accounting.Bills.Create(ctx, operations.AccountingBillsAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        Bill: components.BillInput{
            BillNumber: sdkgo.String("10001"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.AddressTypePrimary.ToPointer(),
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
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            BillDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaidDate: types.MustNewDateFromString("2020-10-30"),
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            LineItems: []components.BillLineItemInput{
                components.BillLineItemInput{
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.BillLineItemTypeExpenseAccount.ToPointer(),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
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
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.BillLineItemInput{
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.BillLineItemTypeExpenseAccount.ToPointer(),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
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
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            Terms: sdkgo.String("Net 30 days"),
            Balance: sdkgo.Float64(27500),
            Deposit: sdkgo.Float64(0),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            Total: sdkgo.Float64(27500),
            TaxCode: sdkgo.String("1234"),
            Notes: sdkgo.String("Some notes about this bill."),
            Status: components.BillStatusDraft.ToPointer(),
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
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
            DiscountPercentage: sdkgo.Float64(5.5),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
            RowVersion: sdkgo.String("1-12345"),
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
            },
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
            AccountingPeriod: sdkgo.String("01-24"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AccountingBillsAddRequest](../../models/operations/accountingbillsaddrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AccountingBillsAddResponse](../../models/operations/accountingbillsaddresponse.md), error**

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

Get Bill

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

    res, err := s.Accounting.Bills.Get(ctx, operations.AccountingBillsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.AccountingBillsOneRequest](../../models/operations/accountingbillsonerequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AccountingBillsOneResponse](../../models/operations/accountingbillsoneresponse.md), error**

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

Update Bill

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

    res, err := s.Accounting.Bills.Update(ctx, operations.AccountingBillsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Bill: components.BillInput{
            BillNumber: sdkgo.String("10001"),
            Supplier: &components.LinkedSupplierInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Address: &components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.AddressTypePrimary.ToPointer(),
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
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            BillDate: types.MustNewDateFromString("2020-09-30"),
            DueDate: types.MustNewDateFromString("2020-10-30"),
            PaidDate: types.MustNewDateFromString("2020-10-30"),
            PoNumber: sdkgo.String("90000117"),
            Reference: sdkgo.String("123456"),
            LineItems: []components.BillLineItemInput{
                components.BillLineItemInput{
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.BillLineItemTypeExpenseAccount.ToPointer(),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
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
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.BillLineItemInput{
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.BillLineItemTypeExpenseAccount.ToPointer(),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
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
                    },
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.BillLineItemInput{
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.BillLineItemTypeExpenseAccount.ToPointer(),
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
                    LedgerAccount: &components.LinkedLedgerAccountInput{
                        ID: sdkgo.String("123456"),
                        NominalCode: sdkgo.String("N091"),
                        Code: sdkgo.String("453"),
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
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            Terms: sdkgo.String("Net 30 days"),
            Balance: sdkgo.Float64(27500),
            Deposit: sdkgo.Float64(0),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            Total: sdkgo.Float64(27500),
            TaxCode: sdkgo.String("1234"),
            Notes: sdkgo.String("Some notes about this bill."),
            Status: components.BillStatusDraft.ToPointer(),
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.String("123456"),
                NominalCode: sdkgo.String("N091"),
                Code: sdkgo.String("453"),
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
            DiscountPercentage: sdkgo.Float64(5.5),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                },
            },
            RowVersion: sdkgo.String("1-12345"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueFour(
                        components.Four{},
                    )),
                },
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
            AccountingPeriod: sdkgo.String("01-24"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingBillsUpdateRequest](../../models/operations/accountingbillsupdaterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingBillsUpdateResponse](../../models/operations/accountingbillsupdateresponse.md), error**

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

Delete Bill

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

    res, err := s.Accounting.Bills.Delete(ctx, operations.AccountingBillsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBillResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingBillsDeleteRequest](../../models/operations/accountingbillsdeleterequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingBillsDeleteResponse](../../models/operations/accountingbillsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |