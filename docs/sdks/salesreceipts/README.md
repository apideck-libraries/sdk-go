# Accounting.SalesReceipts

## Overview

### Available Operations

* [List](#list) - List Sales Receipts
* [Create](#create) - Create Sales Receipt
* [Get](#get) - Get Sales Receipt
* [Update](#update) - Update Sales Receipt
* [Delete](#delete) - Delete Sales Receipt

## List

List Sales Receipts

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesReceiptsAll" method="get" path="/accounting/sales-receipts" -->
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

    res, err := s.Accounting.SalesReceipts.List(ctx, operations.AccountingSalesReceiptsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.SalesReceiptsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CustomerID: sdkgo.Pointer("123abc"),
        },
        Sort: &components.SalesReceiptsSort{
            By: components.SalesReceiptsSortByUpdatedAt.ToPointer(),
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
    if res.GetSalesReceiptsResponse != nil {
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AccountingSalesReceiptsAllRequest](../../models/operations/accountingsalesreceiptsallrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AccountingSalesReceiptsAllResponse](../../models/operations/accountingsalesreceiptsallresponse.md), error**

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

Create Sales Receipt

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesReceiptsAdd" method="post" path="/accounting/sales-receipts" -->
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

    res, err := s.Accounting.SalesReceipts.Create(ctx, operations.AccountingSalesReceiptsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        SalesReceipt: components.SalesReceiptInput{
            Number: sdkgo.Pointer("SR-00001"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](250.0),
            TotalAmount: sdkgo.Pointer[float64](49.99),
            TotalTax: sdkgo.Pointer[float64](25.0),
            TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
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
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Note: sdkgo.Pointer("Thank you for your purchase"),
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
    if res.CreateSalesReceiptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AccountingSalesReceiptsAddRequest](../../models/operations/accountingsalesreceiptsaddrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AccountingSalesReceiptsAddResponse](../../models/operations/accountingsalesreceiptsaddresponse.md), error**

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

Get Sales Receipt

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesReceiptsOne" method="get" path="/accounting/sales-receipts/{id}" -->
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

    res, err := s.Accounting.SalesReceipts.Get(ctx, operations.AccountingSalesReceiptsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSalesReceiptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.AccountingSalesReceiptsOneRequest](../../models/operations/accountingsalesreceiptsonerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.AccountingSalesReceiptsOneResponse](../../models/operations/accountingsalesreceiptsoneresponse.md), error**

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

Update Sales Receipt

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesReceiptsUpdate" method="patch" path="/accounting/sales-receipts/{id}" -->
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

    res, err := s.Accounting.SalesReceipts.Update(ctx, operations.AccountingSalesReceiptsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        SalesReceipt: components.SalesReceiptInput{
            Number: sdkgo.Pointer("SR-00001"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](250.0),
            TotalAmount: sdkgo.Pointer[float64](49.99),
            TotalTax: sdkgo.Pointer[float64](25.0),
            TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
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
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Note: sdkgo.Pointer("Thank you for your purchase"),
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
    if res.UpdateSalesReceiptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.AccountingSalesReceiptsUpdateRequest](../../models/operations/accountingsalesreceiptsupdaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.AccountingSalesReceiptsUpdateResponse](../../models/operations/accountingsalesreceiptsupdateresponse.md), error**

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

Delete Sales Receipt

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesReceiptsDelete" method="delete" path="/accounting/sales-receipts/{id}" -->
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

    res, err := s.Accounting.SalesReceipts.Delete(ctx, operations.AccountingSalesReceiptsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteSalesReceiptResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.AccountingSalesReceiptsDeleteRequest](../../models/operations/accountingsalesreceiptsdeleterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.AccountingSalesReceiptsDeleteResponse](../../models/operations/accountingsalesreceiptsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |