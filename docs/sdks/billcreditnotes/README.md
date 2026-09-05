# Accounting.BillCreditNotes

## Overview

### Available Operations

* [List](#list) - List Bill Credit Notes
* [Create](#create) - Create Bill Credit Note
* [Get](#get) - Get Bill Credit Note
* [Update](#update) - Update Bill Credit Note
* [Delete](#delete) - Delete Bill Credit Note

## List

List Bill Credit Notes

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.billCreditNotesAll" method="get" path="/accounting/bill-credit-notes" -->
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

    res, err := s.Accounting.BillCreditNotes.List(ctx, operations.AccountingBillCreditNotesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.BillCreditNotesFilter{
            Ids: sdkgo.Pointer("12345,67890"),
            IDSince: sdkgo.Pointer("1"),
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CreatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            Number: sdkgo.Pointer("OIT00546"),
            SupplierID: sdkgo.Pointer("123abc"),
        },
        Sort: &components.BillCreditNotesSort{
            By: components.BillCreditNotesSortByUpdatedAt.ToPointer(),
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
    if res.GetBillCreditNotesResponse != nil {
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingBillCreditNotesAllRequest](../../models/operations/accountingbillcreditnotesallrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingBillCreditNotesAllResponse](../../models/operations/accountingbillcreditnotesallresponse.md), error**

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

Create Bill Credit Note

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.billCreditNotesAdd" method="post" path="/accounting/bill-credit-notes" -->
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

    res, err := s.Accounting.BillCreditNotes.Create(ctx, operations.AccountingBillCreditNotesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        BillCreditNote: components.BillCreditNoteInput{
            Number: sdkgo.Pointer("OIT00546"),
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
                    Line5: sdkgo.Pointer("Attention: Finance Dept"),
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
            Subsidiary: &components.LinkedSubsidiaryInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
            },
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500.0),
            TotalAmount: 49.99,
            TotalTax: sdkgo.Pointer[float64](2500.0),
            TaxCode: sdkgo.Pointer("1234"),
            Balance: sdkgo.Pointer[float64](27500.0),
            RemainingCredit: sdkgo.Pointer[float64](27500.0),
            Status: components.BillCreditNoteStatusAuthorised.ToPointer(),
            Reference: sdkgo.Pointer("123456"),
            DateIssued: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            DatePaid: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Type: components.BillCreditNoteTypeAccountsPayableCredit.ToPointer(),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Bank account"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            LineItems: []components.BillCreditNoteLineItemInput{
                components.BillCreditNoteLineItemInput{
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Returned goods credit"),
                    Type: components.LineItemTypeExpenseAccount.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27.5),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    Location: &components.LinkedLocationInput{
                        ID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("New York Office"),
                    },
                    Department: nil,
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
                    },
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Code: sdkgo.Pointer("100"),
                            Name: sdkgo.Pointer("New York"),
                            ParentID: sdkgo.Pointer("123456"),
                            ParentName: sdkgo.Pointer("New York"),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            Allocations: []components.AllocationInput{
                components.AllocationInput{
                    ID: sdkgo.Pointer("123456"),
                    Amount: sdkgo.Pointer[float64](49.99),
                    AllocationID: sdkgo.Pointer("123456"),
                },
            },
            Note: sdkgo.Pointer("Some notes about this bill credit note"),
            Terms: sdkgo.Pointer("Some terms about this bill credit note"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Code: sdkgo.Pointer("100"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        RefName: sdkgo.Pointer("Marketing"),
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingBillCreditNotesAddRequest](../../models/operations/accountingbillcreditnotesaddrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingBillCreditNotesAddResponse](../../models/operations/accountingbillcreditnotesaddresponse.md), error**

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

Get Bill Credit Note

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.billCreditNotesOne" method="get" path="/accounting/bill-credit-notes/{id}" -->
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

    res, err := s.Accounting.BillCreditNotes.Get(ctx, operations.AccountingBillCreditNotesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingBillCreditNotesOneRequest](../../models/operations/accountingbillcreditnotesonerequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingBillCreditNotesOneResponse](../../models/operations/accountingbillcreditnotesoneresponse.md), error**

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

Update Bill Credit Note

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.billCreditNotesUpdate" method="patch" path="/accounting/bill-credit-notes/{id}" -->
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

    res, err := s.Accounting.BillCreditNotes.Update(ctx, operations.AccountingBillCreditNotesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        BillCreditNote: components.BillCreditNoteInput{
            Number: sdkgo.Pointer("OIT00546"),
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
                    Line5: sdkgo.Pointer("Attention: Finance Dept"),
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
            Subsidiary: &components.LinkedSubsidiaryInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
            },
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500.0),
            TotalAmount: 49.99,
            TotalTax: sdkgo.Pointer[float64](2500.0),
            TaxCode: sdkgo.Pointer("1234"),
            Balance: sdkgo.Pointer[float64](27500.0),
            RemainingCredit: sdkgo.Pointer[float64](27500.0),
            Status: components.BillCreditNoteStatusAuthorised.ToPointer(),
            Reference: sdkgo.Pointer("123456"),
            DateIssued: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            DatePaid: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Type: components.BillCreditNoteTypeAccountsPayableCredit.ToPointer(),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Bank account"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            LineItems: []components.BillCreditNoteLineItemInput{
                components.BillCreditNoteLineItemInput{
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Returned goods credit"),
                    Type: components.LineItemTypeExpenseAccount.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27.5),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    Location: &components.LinkedLocationInput{
                        ID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("New York Office"),
                    },
                    Department: &components.LinkedDepartmentInput{
                        DisplayID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Acme Inc."),
                    },
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
                    },
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Code: sdkgo.Pointer("100"),
                            Name: sdkgo.Pointer("New York"),
                            ParentID: sdkgo.Pointer("123456"),
                            ParentName: sdkgo.Pointer("New York"),
                        },
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            Allocations: []components.AllocationInput{
                components.AllocationInput{
                    ID: sdkgo.Pointer("123456"),
                    Amount: sdkgo.Pointer[float64](49.99),
                    AllocationID: sdkgo.Pointer("123456"),
                },
            },
            Note: sdkgo.Pointer("Some notes about this bill credit note"),
            Terms: sdkgo.Pointer("Some terms about this bill credit note"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Code: sdkgo.Pointer("100"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        RefName: sdkgo.Pointer("Marketing"),
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AccountingBillCreditNotesUpdateRequest](../../models/operations/accountingbillcreditnotesupdaterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AccountingBillCreditNotesUpdateResponse](../../models/operations/accountingbillcreditnotesupdateresponse.md), error**

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

Delete Bill Credit Note

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.billCreditNotesDelete" method="delete" path="/accounting/bill-credit-notes/{id}" -->
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

    res, err := s.Accounting.BillCreditNotes.Delete(ctx, operations.AccountingBillCreditNotesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBillCreditNoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AccountingBillCreditNotesDeleteRequest](../../models/operations/accountingbillcreditnotesdeleterequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AccountingBillCreditNotesDeleteResponse](../../models/operations/accountingbillcreditnotesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |