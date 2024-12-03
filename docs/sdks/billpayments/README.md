# BillPayments
(*Accounting.BillPayments*)

## Overview

### Available Operations

* [List](#list) - List Bill Payments
* [Create](#create) - Create Bill Payment
* [Get](#get) - Get Bill Payment
* [Update](#update) - Update Bill Payment
* [Delete](#delete) - Delete Bill Payment

## List

List Bill Payments

### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"context"
	"log"
)

func main() {
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    ctx := context.Background()
    res, err := s.Accounting.BillPayments.List(ctx, operations.AccountingBillPaymentsAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.PaymentsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.PaymentsSort{
            By: components.PaymentsSortByUpdatedAt.ToPointer(),
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
    if res.GetBillPaymentsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingBillPaymentsAllRequest](../../models/operations/accountingbillpaymentsallrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingBillPaymentsAllResponse](../../models/operations/accountingbillpaymentsallresponse.md), error**

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

Create Bill Payment

### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
	"context"
	"log"
)

func main() {
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    ctx := context.Background()
    res, err := s.Accounting.BillPayments.Create(ctx, components.BillPaymentInput{
        Currency: components.CurrencyUsd.ToPointer(),
        CurrencyRate: sdkgo.Float64(0.69),
        TotalAmount: sdkgo.Float64(49.99),
        Reference: sdkgo.String("123456"),
        PaymentMethod: sdkgo.String("cash"),
        PaymentMethodReference: sdkgo.String("123456"),
        PaymentMethodID: sdkgo.String("12345"),
        Account: &components.LinkedLedgerAccountInput{
            ID: sdkgo.String("123456"),
            NominalCode: sdkgo.String("N091"),
            Code: sdkgo.String("453"),
        },
        TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
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
        CompanyID: sdkgo.String("12345"),
        Reconciled: sdkgo.Bool(true),
        Status: components.PaymentStatusAuthorised.ToPointer(),
        Type: components.BillPaymentTypeAccountsPayable.ToPointer(),
        Allocations: []components.BillPaymentAllocations{
            components.BillPaymentAllocations{
                ID: sdkgo.String("12345"),
                Type: components.BillPaymentAllocationTypeBill.ToPointer(),
                Amount: sdkgo.Float64(49.99),
                AllocationID: sdkgo.String("123456"),
            },
        },
        Note: sdkgo.String("Some notes about this transaction"),
        Number: sdkgo.String("123456"),
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
        DisplayID: sdkgo.String("123456"),
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
    }, nil, sdkgo.String("salesforce"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBillPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `billPayment`                                                                                                                                 | [components.BillPaymentInput](../../models/components/billpaymentinput.md)                                                                    | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingBillPaymentsAddResponse](../../models/operations/accountingbillpaymentsaddresponse.md), error**

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

Get Bill Payment

### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"context"
	"log"
)

func main() {
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    ctx := context.Background()
    res, err := s.Accounting.BillPayments.Get(ctx, "<id>", sdkgo.String("salesforce"), nil, sdkgo.String("id,updated_at"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBillPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `id`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | ID of the record you are acting upon.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `serviceID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | salesforce                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `raw`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | **bool*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Include raw response. Mostly used for debugging purposes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `fields`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded. | id,updated_at                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

### Response

**[*operations.AccountingBillPaymentsOneResponse](../../models/operations/accountingbillpaymentsoneresponse.md), error**

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

Update Bill Payment

### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
	"context"
	"log"
)

func main() {
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    ctx := context.Background()
    res, err := s.Accounting.BillPayments.Update(ctx, "<id>", components.BillPaymentInput{
        Currency: components.CurrencyUsd.ToPointer(),
        CurrencyRate: sdkgo.Float64(0.69),
        TotalAmount: sdkgo.Float64(49.99),
        Reference: sdkgo.String("123456"),
        PaymentMethod: sdkgo.String("cash"),
        PaymentMethodReference: sdkgo.String("123456"),
        PaymentMethodID: sdkgo.String("12345"),
        Account: &components.LinkedLedgerAccountInput{
            ID: sdkgo.String("123456"),
            NominalCode: sdkgo.String("N091"),
            Code: sdkgo.String("453"),
        },
        TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
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
        CompanyID: sdkgo.String("12345"),
        Reconciled: sdkgo.Bool(true),
        Status: components.PaymentStatusAuthorised.ToPointer(),
        Type: components.BillPaymentTypeAccountsPayable.ToPointer(),
        Allocations: []components.BillPaymentAllocations{
            components.BillPaymentAllocations{
                ID: sdkgo.String("12345"),
                Type: components.BillPaymentAllocationTypeBill.ToPointer(),
                Amount: sdkgo.Float64(49.99),
                AllocationID: sdkgo.String("123456"),
            },
            components.BillPaymentAllocations{
                ID: sdkgo.String("12345"),
                Type: components.BillPaymentAllocationTypeBill.ToPointer(),
                Amount: sdkgo.Float64(49.99),
                AllocationID: sdkgo.String("123456"),
            },
            components.BillPaymentAllocations{
                ID: sdkgo.String("12345"),
                Type: components.BillPaymentAllocationTypeBill.ToPointer(),
                Amount: sdkgo.Float64(49.99),
                AllocationID: sdkgo.String("123456"),
            },
        },
        Note: sdkgo.String("Some notes about this transaction"),
        Number: sdkgo.String("123456"),
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
        CustomFields: []components.CustomField{
            components.CustomField{
                ID: sdkgo.String("2389328923893298"),
                Name: sdkgo.String("employee_level"),
                Description: sdkgo.String("Employee Level"),
            },
        },
        RowVersion: sdkgo.String("1-12345"),
        DisplayID: sdkgo.String("123456"),
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
        },
    }, sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateBillPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `id`                                                                                                                                          | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | ID of the record you are acting upon.                                                                                                         |                                                                                                                                               |
| `billPayment`                                                                                                                                 | [components.BillPaymentInput](../../models/components/billpaymentinput.md)                                                                    | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingBillPaymentsUpdateResponse](../../models/operations/accountingbillpaymentsupdateresponse.md), error**

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

Delete Bill Payment

### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"context"
	"log"
)

func main() {
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    ctx := context.Background()
    res, err := s.Accounting.BillPayments.Delete(ctx, "<id>", sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBillPaymentResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `id`                                                                                                                                          | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | ID of the record you are acting upon.                                                                                                         |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingBillPaymentsDeleteResponse](../../models/operations/accountingbillpaymentsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |