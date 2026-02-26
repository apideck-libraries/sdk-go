# Accounting.ExpenseReports

## Overview

### Available Operations

* [List](#list) - List Expense Reports
* [Create](#create) - Create Expense Report
* [Get](#get) - Get Expense Report
* [Update](#update) - Update Expense Report
* [Delete](#delete) - Delete Expense Report

## List

List Expense Reports

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.expenseReportsAll" method="get" path="/accounting/expense-reports" -->
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

    res, err := s.Accounting.ExpenseReports.List(ctx, operations.AccountingExpenseReportsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
        Filter: &components.ExpenseReportsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            Status: components.ExpenseReportsFilterStatusSubmitted.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetExpenseReportsResponse != nil {
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
| `request`                                                                                                      | [operations.AccountingExpenseReportsAllRequest](../../models/operations/accountingexpensereportsallrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingExpenseReportsAllResponse](../../models/operations/accountingexpensereportsallresponse.md), error**

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

Create Expense Report

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.expenseReportsAdd" method="post" path="/accounting/expense-reports" -->
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

    res, err := s.Accounting.ExpenseReports.Create(ctx, operations.AccountingExpenseReportsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        ExpenseReport: components.ExpenseReportInput{
            DisplayID: sdkgo.Pointer("123456"),
            Number: sdkgo.Pointer("ER-001"),
            Title: sdkgo.Pointer("Q1 Business Travel"),
            Employee: components.ExpenseReportEmployee{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("John Doe"),
            },
            Status: components.ExpenseReportStatusSubmitted.ToPointer(),
            TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            PostingDate: types.MustNewDateFromString("2024-06-01"),
            DueDate: types.MustNewDateFromString("2024-06-15"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            SubTotal: sdkgo.Pointer[float64](250),
            TotalTax: sdkgo.Pointer[float64](25),
            TotalAmount: sdkgo.Pointer[float64](1250.75),
            ReimbursableAmount: sdkgo.Pointer[float64](1100),
            Memo: sdkgo.Pointer("Business travel expenses for Q1 client meetings"),
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
            },
            Account: nil,
            AccountingPeriod: &components.AccountingPeriod{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Q1 2024"),
            },
            LineItems: []components.ExpenseReportLineItemInput{},
            Subsidiary: &components.LinkedSubsidiaryInput{
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
            TaxInclusive: sdkgo.Pointer(true),
            ApprovedBy: &components.ApprovedBy{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Jane Smith"),
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
    if res.CreateExpenseReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingExpenseReportsAddRequest](../../models/operations/accountingexpensereportsaddrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingExpenseReportsAddResponse](../../models/operations/accountingexpensereportsaddresponse.md), error**

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

Get Expense Report

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.expenseReportsOne" method="get" path="/accounting/expense-reports/{id}" -->
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

    res, err := s.Accounting.ExpenseReports.Get(ctx, operations.AccountingExpenseReportsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetExpenseReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingExpenseReportsOneRequest](../../models/operations/accountingexpensereportsonerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingExpenseReportsOneResponse](../../models/operations/accountingexpensereportsoneresponse.md), error**

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

Update Expense Report

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.expenseReportsUpdate" method="patch" path="/accounting/expense-reports/{id}" -->
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

    res, err := s.Accounting.ExpenseReports.Update(ctx, operations.AccountingExpenseReportsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        ExpenseReport: components.ExpenseReportInput{
            DisplayID: sdkgo.Pointer("123456"),
            Number: sdkgo.Pointer("ER-001"),
            Title: sdkgo.Pointer("Q1 Business Travel"),
            Employee: components.ExpenseReportEmployee{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("John Doe"),
            },
            Status: components.ExpenseReportStatusSubmitted.ToPointer(),
            TransactionDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            PostingDate: types.MustNewDateFromString("2024-06-01"),
            DueDate: types.MustNewDateFromString("2024-06-15"),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            SubTotal: sdkgo.Pointer[float64](250),
            TotalTax: sdkgo.Pointer[float64](25),
            TotalAmount: sdkgo.Pointer[float64](1250.75),
            ReimbursableAmount: sdkgo.Pointer[float64](1100),
            Memo: sdkgo.Pointer("Business travel expenses for Q1 client meetings"),
            Department: &components.LinkedDepartmentInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            Location: &components.LinkedLocationInput{
                ID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("New York Office"),
            },
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Bank account"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            AccountingPeriod: &components.AccountingPeriod{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Q1 2024"),
            },
            LineItems: []components.ExpenseReportLineItemInput{},
            Subsidiary: &components.LinkedSubsidiaryInput{
                DisplayID: sdkgo.Pointer("123456"),
                Name: sdkgo.Pointer("Acme Inc."),
            },
            TrackingCategories: nil,
            TaxInclusive: sdkgo.Pointer(true),
            ApprovedBy: &components.ApprovedBy{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Jane Smith"),
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
    if res.UpdateExpenseReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingExpenseReportsUpdateRequest](../../models/operations/accountingexpensereportsupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingExpenseReportsUpdateResponse](../../models/operations/accountingexpensereportsupdateresponse.md), error**

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

Delete Expense Report

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.expenseReportsDelete" method="delete" path="/accounting/expense-reports/{id}" -->
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

    res, err := s.Accounting.ExpenseReports.Delete(ctx, operations.AccountingExpenseReportsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteExpenseReportResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingExpenseReportsDeleteRequest](../../models/operations/accountingexpensereportsdeleterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingExpenseReportsDeleteResponse](../../models/operations/accountingexpensereportsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |