# Accounting.Refunds

## Overview

### Available Operations

* [List](#list) - List Refunds
* [Create](#create) - Create Refund
* [Get](#get) - Get Refund
* [Update](#update) - Update Refund
* [Delete](#delete) - Delete Refund

## List

List Refunds

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.refundsAll" method="get" path="/accounting/refunds" -->
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

    res, err := s.Accounting.Refunds.List(ctx, operations.AccountingRefundsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.RefundsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CustomerID: sdkgo.Pointer("123abc"),
        },
        Sort: &components.RefundsSort{
            By: components.RefundsSortByUpdatedAt.ToPointer(),
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
    if res.GetRefundsResponse != nil {
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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountingRefundsAllRequest](../../models/operations/accountingrefundsallrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AccountingRefundsAllResponse](../../models/operations/accountingrefundsallresponse.md), error**

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

Create Refund

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.refundsAdd" method="post" path="/accounting/refunds" -->
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

    res, err := s.Accounting.Refunds.Create(ctx, operations.AccountingRefundsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Refund: components.RefundInput{
            Number: sdkgo.Pointer("RF-00001"),
            Customer: nil,
            CompanyID: sdkgo.Pointer("12345"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](250.0),
            TotalAmount: sdkgo.Pointer[float64](49.99),
            TotalTax: sdkgo.Pointer[float64](25.0),
            RefundDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Status: components.RefundStatusPaid.ToPointer(),
            Type: components.RefundTypeRefundReceipt.ToPointer(),
            PaymentMethod: sdkgo.Pointer("cash"),
            PaymentMethodReference: sdkgo.Pointer("123456"),
            PaymentMethodID: sdkgo.Pointer("12345"),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Bank account"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500.0),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    CategoryID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    SubsidiaryID: sdkgo.Pointer("12345"),
                    ShippingID: sdkgo.Pointer("12345"),
                    Memo: sdkgo.Pointer("Some memo"),
                    Prepaid: sdkgo.Pointer(true),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxApplicableOn: sdkgo.Pointer("Domestic_Purchase_of_Goods_and_Services"),
                    TaxRecoverability: sdkgo.Pointer("Fully_Recoverable"),
                    TaxMethod: sdkgo.Pointer("Due_to_Supplier"),
                    Worktags: []*components.LinkedWorktag{
                        &components.LinkedWorktag{
                            ID: sdkgo.Pointer("123456"),
                            Value: sdkgo.Pointer("New York"),
                        },
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
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
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
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
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Note: sdkgo.Pointer("Refund for returned items"),
            CustomerMemo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            Reference: sdkgo.Pointer("REF-123456"),
            BillingAddress: &components.Address{
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
            ShippingAddress: &components.Address{
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
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
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
    if res.CreateRefundResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountingRefundsAddRequest](../../models/operations/accountingrefundsaddrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AccountingRefundsAddResponse](../../models/operations/accountingrefundsaddresponse.md), error**

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

Get Refund

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.refundsOne" method="get" path="/accounting/refunds/{id}" -->
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

    res, err := s.Accounting.Refunds.Get(ctx, operations.AccountingRefundsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetRefundResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.AccountingRefundsOneRequest](../../models/operations/accountingrefundsonerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.AccountingRefundsOneResponse](../../models/operations/accountingrefundsoneresponse.md), error**

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

Update Refund

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.refundsUpdate" method="patch" path="/accounting/refunds/{id}" -->
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

    res, err := s.Accounting.Refunds.Update(ctx, operations.AccountingRefundsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Refund: components.RefundInput{
            Number: sdkgo.Pointer("RF-00001"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            CompanyID: sdkgo.Pointer("12345"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](250.0),
            TotalAmount: sdkgo.Pointer[float64](49.99),
            TotalTax: sdkgo.Pointer[float64](25.0),
            RefundDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            Status: components.RefundStatusPaid.ToPointer(),
            Type: components.RefundTypeRefundReceipt.ToPointer(),
            PaymentMethod: sdkgo.Pointer("cash"),
            PaymentMethodReference: sdkgo.Pointer("123456"),
            PaymentMethodID: sdkgo.Pointer("12345"),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Bank account"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500.0),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    CategoryID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    SubsidiaryID: sdkgo.Pointer("12345"),
                    ShippingID: sdkgo.Pointer("12345"),
                    Memo: sdkgo.Pointer("Some memo"),
                    Prepaid: sdkgo.Pointer(true),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxApplicableOn: sdkgo.Pointer("Domestic_Purchase_of_Goods_and_Services"),
                    TaxRecoverability: sdkgo.Pointer("Fully_Recoverable"),
                    TaxMethod: sdkgo.Pointer("Due_to_Supplier"),
                    Worktags: []*components.LinkedWorktag{
                        &components.LinkedWorktag{
                            ID: sdkgo.Pointer("123456"),
                            Value: sdkgo.Pointer("New York"),
                        },
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
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
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
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
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Note: sdkgo.Pointer("Refund for returned items"),
            CustomerMemo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            Reference: sdkgo.Pointer("REF-123456"),
            BillingAddress: &components.Address{
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
            ShippingAddress: &components.Address{
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
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
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
    if res.UpdateRefundResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AccountingRefundsUpdateRequest](../../models/operations/accountingrefundsupdaterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.AccountingRefundsUpdateResponse](../../models/operations/accountingrefundsupdateresponse.md), error**

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

Delete Refund

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.refundsDelete" method="delete" path="/accounting/refunds/{id}" -->
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

    res, err := s.Accounting.Refunds.Delete(ctx, operations.AccountingRefundsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteRefundResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.AccountingRefundsDeleteRequest](../../models/operations/accountingrefundsdeleterequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.AccountingRefundsDeleteResponse](../../models/operations/accountingrefundsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |