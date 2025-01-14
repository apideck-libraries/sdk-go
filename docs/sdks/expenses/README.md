# Expenses
(*Accounting.Expenses*)

## Overview

### Available Operations

* [List](#list) - List Expenses
* [Create](#create) - Create Expense
* [Get](#get) - Get Expense
* [Update](#update) - Update Expense
* [Delete](#delete) - Delete Expense

## List

List Expenses

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.Expenses.List(ctx, nil, sdkgo.String("salesforce"), nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetExpensesResponse != nil {
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

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `cursor`                                                                                                                                      | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Cursor to start from. You can find cursors for next/previous pages in the meta.cursors property of the response.                              |                                                                                                                                               |
| `limit`                                                                                                                                       | **int64*                                                                                                                                      | :heavy_minus_sign:                                                                                                                            | Number of results to return. Minimum 1, Maximum 200, Default 20                                                                               |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingExpensesAllResponse](../../models/operations/accountingexpensesallresponse.md), error**

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

Create Expense

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.Expenses.Create(ctx, components.ExpenseInput{
        Number: sdkgo.String("OIT00546"),
        TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
        AccountID: "123456",
        CustomerID: sdkgo.String("12345"),
        SupplierID: sdkgo.String("12345"),
        CompanyID: sdkgo.String("12345"),
        DepartmentID: sdkgo.String("12345"),
        Currency: components.CurrencyUsd.ToPointer(),
        CurrencyRate: sdkgo.Float64(0.69),
        Type: components.ExpenseTypeExpense.ToPointer(),
        Memo: sdkgo.String("For travel expenses incurred on 2024-05-15"),
        TaxRate: &components.LinkedTaxRateInput{
            ID: sdkgo.String("123456"),
            Rate: sdkgo.Float64(10),
        },
        TotalAmount: sdkgo.Float64(275),
        LineItems: []components.ExpenseLineItemInput{
            components.ExpenseLineItemInput{
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
                AccountID: sdkgo.String("123456"),
                CustomerID: sdkgo.String("12345"),
                DepartmentID: sdkgo.String("12345"),
                LocationID: sdkgo.String("12345"),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
                Description: sdkgo.String("Travel US."),
                TotalAmount: sdkgo.Float64(275),
                Billable: sdkgo.Bool(true),
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
        PassThrough: []components.PassThroughBody{

        },
    }, nil, sdkgo.String("salesforce"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `expense`                                                                                                                                     | [components.ExpenseInput](../../models/components/expenseinput.md)                                                                            | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingExpensesAddResponse](../../models/operations/accountingexpensesaddresponse.md), error**

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

Get Expense

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.Expenses.Get(ctx, "<id>", sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetExpenseResponse != nil {
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

**[*operations.AccountingExpensesOneResponse](../../models/operations/accountingexpensesoneresponse.md), error**

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

Update Expense

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
	"github.com/apideck-libraries/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.Expenses.Update(ctx, "<id>", components.ExpenseInput{
        Number: sdkgo.String("OIT00546"),
        TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
        AccountID: "123456",
        CustomerID: sdkgo.String("12345"),
        SupplierID: sdkgo.String("12345"),
        CompanyID: sdkgo.String("12345"),
        DepartmentID: sdkgo.String("12345"),
        Currency: components.CurrencyUsd.ToPointer(),
        CurrencyRate: sdkgo.Float64(0.69),
        Type: components.ExpenseTypeExpense.ToPointer(),
        Memo: sdkgo.String("For travel expenses incurred on 2024-05-15"),
        TaxRate: &components.LinkedTaxRateInput{
            ID: sdkgo.String("123456"),
            Rate: sdkgo.Float64(10),
        },
        TotalAmount: sdkgo.Float64(275),
        LineItems: []components.ExpenseLineItemInput{
            components.ExpenseLineItemInput{
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
                AccountID: sdkgo.String("123456"),
                CustomerID: sdkgo.String("12345"),
                DepartmentID: sdkgo.String("12345"),
                LocationID: sdkgo.String("12345"),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
                Description: sdkgo.String("Travel US."),
                TotalAmount: sdkgo.Float64(275),
                Billable: sdkgo.Bool(true),
            },
            components.ExpenseLineItemInput{
                TrackingCategories: []components.LinkedTrackingCategory{
                    components.LinkedTrackingCategory{
                        ID: sdkgo.String("123456"),
                        Name: sdkgo.String("New York"),
                    },
                },
                AccountID: sdkgo.String("123456"),
                CustomerID: sdkgo.String("12345"),
                DepartmentID: sdkgo.String("12345"),
                LocationID: sdkgo.String("12345"),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
                Description: sdkgo.String("Travel US."),
                TotalAmount: sdkgo.Float64(275),
                Billable: sdkgo.Bool(true),
            },
            components.ExpenseLineItemInput{
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
                AccountID: sdkgo.String("123456"),
                CustomerID: sdkgo.String("12345"),
                DepartmentID: sdkgo.String("12345"),
                LocationID: sdkgo.String("12345"),
                TaxRate: &components.LinkedTaxRateInput{
                    ID: sdkgo.String("123456"),
                    Rate: sdkgo.Float64(10),
                },
                Description: sdkgo.String("Travel US."),
                TotalAmount: sdkgo.Float64(275),
                Billable: sdkgo.Bool(true),
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
            components.CustomField{
                ID: sdkgo.String("2389328923893298"),
                Name: sdkgo.String("employee_level"),
                Description: sdkgo.String("Employee Level"),
                Value: sdkgo.Pointer(components.CreateValueFour(
                    components.Four{},
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
            components.PassThroughBody{
                ServiceID: "<id>",
                ExtendPaths: []components.ExtendPaths{

                },
            },
        },
    }, sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateExpenseResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `id`                                                                                                                                          | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | ID of the record you are acting upon.                                                                                                         |                                                                                                                                               |
| `expense`                                                                                                                                     | [components.ExpenseInput](../../models/components/expenseinput.md)                                                                            | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.AccountingExpensesUpdateResponse](../../models/operations/accountingexpensesupdateresponse.md), error**

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

Delete Expense

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Accounting.Expenses.Delete(ctx, "<id>", sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteExpenseResponse != nil {
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

**[*operations.AccountingExpensesDeleteResponse](../../models/operations/accountingexpensesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |