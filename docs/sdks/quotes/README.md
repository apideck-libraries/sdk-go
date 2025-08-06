# Quotes
(*Accounting.Quotes*)

## Overview

### Available Operations

* [List](#list) - List Quotes
* [Create](#create) - Create Quote
* [Get](#get) - Get Quote
* [Update](#update) - Update Quote
* [Delete](#delete) - Delete Quote

## List

List Quotes

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.quotesAll" method="get" path="/accounting/quotes" -->
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

    res, err := s.Accounting.Quotes.List(ctx, operations.AccountingQuotesAllRequest{
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetQuotesResponse != nil {
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountingQuotesAllRequest](../../models/operations/accountingquotesallrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AccountingQuotesAllResponse](../../models/operations/accountingquotesallresponse.md), error**

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

Create Quote

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.quotesAdd" method="post" path="/accounting/quotes" -->
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

    res, err := s.Accounting.Quotes.Create(ctx, operations.AccountingQuotesAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        Quote: components.QuoteInput{
            Number: sdkgo.String("QT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            SalesOrderID: sdkgo.String("123456"),
            CompanyID: sdkgo.String("12345"),
            ProjectID: sdkgo.String("12345"),
            QuoteDate: types.MustNewDateFromString("2020-09-30"),
            ExpiryDate: types.MustNewDateFromString("2020-10-30"),
            Terms: sdkgo.String("Valid for 30 days"),
            Reference: sdkgo.String("INV-2024-001"),
            Status: components.QuoteStatusDraft.ToPointer(),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            DiscountPercentage: sdkgo.Float64(5.5),
            DiscountAmount: sdkgo.Float64(25),
            Total: sdkgo.Float64(27500),
            CustomerMemo: sdkgo.String("Thank you for considering our services!"),
            LineItems: []components.QuoteLineItemInput{
                components.QuoteLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.QuoteLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    CategoryID: sdkgo.String("12345"),
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
                            ParentID: sdkgo.String("123456"),
                            ParentName: sdkgo.String("New York"),
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
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                    ParentID: sdkgo.String("123456"),
                    ParentName: sdkgo.String("New York"),
                },
            },
            TemplateID: sdkgo.String("123456"),
            SourceDocumentURL: sdkgo.String("https://www.quotesolution.com/quote/123456"),
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
    if res.CreateQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountingQuotesAddRequest](../../models/operations/accountingquotesaddrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AccountingQuotesAddResponse](../../models/operations/accountingquotesaddresponse.md), error**

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

Get Quote

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.quotesOne" method="get" path="/accounting/quotes/{id}" -->
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

    res, err := s.Accounting.Quotes.Get(ctx, operations.AccountingQuotesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AccountingQuotesOneRequest](../../models/operations/accountingquotesonerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AccountingQuotesOneResponse](../../models/operations/accountingquotesoneresponse.md), error**

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

Update Quote

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.quotesUpdate" method="patch" path="/accounting/quotes/{id}" -->
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

    res, err := s.Accounting.Quotes.Update(ctx, operations.AccountingQuotesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Quote: components.QuoteInput{
            Number: sdkgo.String("QT00546"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.String("12345"),
                DisplayName: sdkgo.String("Windsurf Shop"),
                Email: sdkgo.String("boring@boring.com"),
            },
            SalesOrderID: sdkgo.String("123456"),
            CompanyID: sdkgo.String("12345"),
            ProjectID: sdkgo.String("12345"),
            QuoteDate: types.MustNewDateFromString("2020-09-30"),
            ExpiryDate: types.MustNewDateFromString("2020-10-30"),
            Terms: sdkgo.String("Valid for 30 days"),
            Reference: sdkgo.String("INV-2024-001"),
            Status: components.QuoteStatusDraft.ToPointer(),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Float64(0.69),
            TaxInclusive: sdkgo.Bool(true),
            SubTotal: sdkgo.Float64(27500),
            TotalTax: sdkgo.Float64(2500),
            TaxCode: sdkgo.String("1234"),
            DiscountPercentage: sdkgo.Float64(5.5),
            DiscountAmount: sdkgo.Float64(25),
            Total: sdkgo.Float64(27500),
            CustomerMemo: sdkgo.String("Thank you for considering our services!"),
            LineItems: []components.QuoteLineItemInput{
                components.QuoteLineItemInput{
                    ID: sdkgo.String("12345"),
                    RowID: sdkgo.String("12345"),
                    Code: sdkgo.String("120-C"),
                    LineNumber: sdkgo.Int64(1),
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.QuoteLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Quantity: sdkgo.Float64(1),
                    UnitPrice: sdkgo.Float64(27500.5),
                    UnitOfMeasure: sdkgo.String("pc."),
                    DiscountPercentage: sdkgo.Float64(0.01),
                    DiscountAmount: sdkgo.Float64(19.99),
                    CategoryID: sdkgo.String("12345"),
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
                    TrackingCategories: nil,
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
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.String("123456"),
                    Name: sdkgo.String("New York"),
                    ParentID: sdkgo.String("123456"),
                    ParentName: sdkgo.String("New York"),
                },
            },
            TemplateID: sdkgo.String("123456"),
            SourceDocumentURL: sdkgo.String("https://www.quotesolution.com/quote/123456"),
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
    if res.UpdateQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountingQuotesUpdateRequest](../../models/operations/accountingquotesupdaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AccountingQuotesUpdateResponse](../../models/operations/accountingquotesupdateresponse.md), error**

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

Delete Quote

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.quotesDelete" method="delete" path="/accounting/quotes/{id}" -->
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

    res, err := s.Accounting.Quotes.Delete(ctx, operations.AccountingQuotesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteQuoteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountingQuotesDeleteRequest](../../models/operations/accountingquotesdeleterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AccountingQuotesDeleteResponse](../../models/operations/accountingquotesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |