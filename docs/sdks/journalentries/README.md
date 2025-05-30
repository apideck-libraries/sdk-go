# JournalEntries
(*Accounting.JournalEntries*)

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

    res, err := s.Accounting.JournalEntries.List(ctx, operations.AccountingJournalEntriesAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.JournalEntriesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.JournalEntriesSort{
            By: components.JournalEntriesSortByUpdatedAt.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
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

    res, err := s.Accounting.JournalEntries.Create(ctx, operations.AccountingJournalEntriesAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        JournalEntry: components.JournalEntryInput{
            Title: sdkgo.String("Purchase Invoice-Inventory (USD): 2019/02/01 Batch Summary Entry"),
            CurrencyRate: sdkgo.Float64(0.69),
            Currency: components.CurrencyUsd.ToPointer(),
            CompanyID: sdkgo.String("12345"),
            LineItems: []components.JournalEntryLineItemInput{
                components.JournalEntryLineItemInput{
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    TaxAmount: sdkgo.Float64(27500),
                    SubTotal: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Type: components.JournalEntryLineItemTypeDebit,
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
                    Customer: &components.LinkedCustomerInput{
                        ID: sdkgo.String("12345"),
                        DisplayName: sdkgo.String("Windsurf Shop"),
                        Email: sdkgo.String("boring@boring.com"),
                    },
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
                    DepartmentID: sdkgo.String("12345"),
                    LocationID: sdkgo.String("12345"),
                    LineNumber: sdkgo.Int64(1),
                },
            },
            Status: components.JournalEntryStatusDraft.ToPointer(),
            Memo: sdkgo.String("Thank you for your business and have a great day!"),
            PostedAt: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            JournalSymbol: sdkgo.String("IND"),
            TaxType: sdkgo.String("sales"),
            TaxCode: sdkgo.String("1234"),
            Number: sdkgo.String("OIT00546"),
            TrackingCategories: nil,
            AccountingPeriod: sdkgo.String("01-24"),
            RowVersion: sdkgo.String("1-12345"),
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

    res, err := s.Accounting.JournalEntries.Get(ctx, operations.AccountingJournalEntriesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
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

    res, err := s.Accounting.JournalEntries.Update(ctx, operations.AccountingJournalEntriesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        JournalEntry: components.JournalEntryInput{
            Title: sdkgo.String("Purchase Invoice-Inventory (USD): 2019/02/01 Batch Summary Entry"),
            CurrencyRate: sdkgo.Float64(0.69),
            Currency: components.CurrencyUsd.ToPointer(),
            CompanyID: sdkgo.String("12345"),
            LineItems: []components.JournalEntryLineItemInput{
                components.JournalEntryLineItemInput{
                    Description: sdkgo.String("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    TaxAmount: sdkgo.Float64(27500),
                    SubTotal: sdkgo.Float64(27500),
                    TotalAmount: sdkgo.Float64(27500),
                    Type: components.JournalEntryLineItemTypeDebit,
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
                    Customer: &components.LinkedCustomerInput{
                        ID: sdkgo.String("12345"),
                        DisplayName: sdkgo.String("Windsurf Shop"),
                        Email: sdkgo.String("boring@boring.com"),
                    },
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
                    DepartmentID: sdkgo.String("12345"),
                    LocationID: sdkgo.String("12345"),
                    LineNumber: sdkgo.Int64(1),
                },
            },
            Status: components.JournalEntryStatusDraft.ToPointer(),
            Memo: sdkgo.String("Thank you for your business and have a great day!"),
            PostedAt: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            JournalSymbol: sdkgo.String("IND"),
            TaxType: sdkgo.String("sales"),
            TaxCode: sdkgo.String("1234"),
            Number: sdkgo.String("OIT00546"),
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
            AccountingPeriod: sdkgo.String("01-24"),
            RowVersion: sdkgo.String("1-12345"),
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

    res, err := s.Accounting.JournalEntries.Delete(ctx, operations.AccountingJournalEntriesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
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