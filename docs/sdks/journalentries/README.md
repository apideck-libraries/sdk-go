# Accounting.JournalEntries

## Overview

### Available Operations

* [List](#list) - List Journal Entries
* [Create](#create) - Create Journal Entry
* [Get](#get) - Get Journal Entry
* [Update](#update) - Update Journal Entry
* [Delete](#delete) - Delete Journal Entry

## List

List Journal Entries

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.journalEntriesAll" method="get" path="/accounting/journal-entries" -->
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

    res, err := s.Accounting.JournalEntries.List(ctx, operations.AccountingJournalEntriesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.JournalEntriesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.JournalEntriesSort{
            By: components.JournalEntriesSortByUpdatedAt.ToPointer(),
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
    if res.GetJournalEntriesResponse != nil {
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
| `request`                                                                                                      | [operations.AccountingJournalEntriesAllRequest](../../models/operations/accountingjournalentriesallrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingJournalEntriesAllResponse](../../models/operations/accountingjournalentriesallresponse.md), error**

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

Create Journal Entry

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.journalEntriesAdd" method="post" path="/accounting/journal-entries" -->
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

    res, err := s.Accounting.JournalEntries.Create(ctx, operations.AccountingJournalEntriesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        JournalEntry: components.JournalEntryInput{
            Title: sdkgo.Pointer("Purchase Invoice-Inventory (USD): 2019/02/01 Batch Summary Entry"),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            Currency: components.CurrencyUsd.ToPointer(),
            CompanyID: sdkgo.Pointer("12345"),
            LineItems: []components.JournalEntryLineItemInput{
                components.JournalEntryLineItemInput{
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    SubTotal: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Type: components.JournalEntryLineItemTypeDebit,
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
                    Customer: &components.LinkedCustomerInput{
                        ID: sdkgo.Pointer("12345"),
                        DisplayName: sdkgo.Pointer("Windsurf Shop"),
                        Email: sdkgo.Pointer("boring@boring.com"),
                    },
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
                    DepartmentID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    LineNumber: sdkgo.Pointer[int64](1),
                },
            },
            Status: components.JournalEntryStatusDraft.ToPointer(),
            Memo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            PostedAt: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            JournalSymbol: sdkgo.Pointer("IND"),
            TaxType: sdkgo.Pointer("sales"),
            TaxCode: sdkgo.Pointer("1234"),
            Number: sdkgo.Pointer("OIT00546"),
            TrackingCategories: nil,
            AccountingPeriod: sdkgo.Pointer("01-24"),
            RowVersion: sdkgo.Pointer("1-12345"),
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
    if res.CreateJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingJournalEntriesAddRequest](../../models/operations/accountingjournalentriesaddrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingJournalEntriesAddResponse](../../models/operations/accountingjournalentriesaddresponse.md), error**

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

Get Journal Entry

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.journalEntriesOne" method="get" path="/accounting/journal-entries/{id}" -->
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

    res, err := s.Accounting.JournalEntries.Get(ctx, operations.AccountingJournalEntriesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingJournalEntriesOneRequest](../../models/operations/accountingjournalentriesonerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingJournalEntriesOneResponse](../../models/operations/accountingjournalentriesoneresponse.md), error**

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

Update Journal Entry

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.journalEntriesUpdate" method="patch" path="/accounting/journal-entries/{id}" -->
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

    res, err := s.Accounting.JournalEntries.Update(ctx, operations.AccountingJournalEntriesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        JournalEntry: components.JournalEntryInput{
            Title: sdkgo.Pointer("Purchase Invoice-Inventory (USD): 2019/02/01 Batch Summary Entry"),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            Currency: components.CurrencyUsd.ToPointer(),
            CompanyID: sdkgo.Pointer("12345"),
            LineItems: []components.JournalEntryLineItemInput{
                components.JournalEntryLineItemInput{
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    TaxAmount: sdkgo.Pointer[float64](27500),
                    SubTotal: sdkgo.Pointer[float64](27500),
                    TotalAmount: sdkgo.Pointer[float64](27500),
                    Type: components.JournalEntryLineItemTypeDebit,
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
                    Customer: &components.LinkedCustomerInput{
                        ID: sdkgo.Pointer("12345"),
                        DisplayName: sdkgo.Pointer("Windsurf Shop"),
                        Email: sdkgo.Pointer("boring@boring.com"),
                    },
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
                    DepartmentID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    LineNumber: sdkgo.Pointer[int64](1),
                },
            },
            Status: components.JournalEntryStatusDraft.ToPointer(),
            Memo: sdkgo.Pointer("Thank you for your business and have a great day!"),
            PostedAt: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            JournalSymbol: sdkgo.Pointer("IND"),
            TaxType: sdkgo.Pointer("sales"),
            TaxCode: sdkgo.Pointer("1234"),
            Number: sdkgo.Pointer("OIT00546"),
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                },
            },
            AccountingPeriod: sdkgo.Pointer("01-24"),
            RowVersion: sdkgo.Pointer("1-12345"),
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
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingJournalEntriesUpdateRequest](../../models/operations/accountingjournalentriesupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingJournalEntriesUpdateResponse](../../models/operations/accountingjournalentriesupdateresponse.md), error**

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

Delete Journal Entry

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.journalEntriesDelete" method="delete" path="/accounting/journal-entries/{id}" -->
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

    res, err := s.Accounting.JournalEntries.Delete(ctx, operations.AccountingJournalEntriesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteJournalEntryResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingJournalEntriesDeleteRequest](../../models/operations/accountingjournalentriesdeleterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingJournalEntriesDeleteResponse](../../models/operations/accountingjournalentriesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |