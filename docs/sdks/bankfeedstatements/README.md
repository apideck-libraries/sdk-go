# BankFeedStatements
(*Accounting.BankFeedStatements*)

## Overview

### Available Operations

* [List](#list) - List Bank Feed Statements
* [Create](#create) - Create Bank Feed Statement
* [Get](#get) - Get Bank Feed Statement
* [Update](#update) - Update Bank Feed Statement
* [Delete](#delete) - Delete Bank Feed Statement

## List

List Bank Feed Statements

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankFeedStatementsAll" method="get" path="/accounting/bank-feed-statements" -->
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

    res, err := s.Accounting.BankFeedStatements.List(ctx, operations.AccountingBankFeedStatementsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBankFeedStatementsResponse != nil {
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

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AccountingBankFeedStatementsAllRequest](../../models/operations/accountingbankfeedstatementsallrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AccountingBankFeedStatementsAllResponse](../../models/operations/accountingbankfeedstatementsallresponse.md), error**

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

Create Bank Feed Statement

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankFeedStatementsAdd" method="post" path="/accounting/bank-feed-statements" -->
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

    res, err := s.Accounting.BankFeedStatements.Create(ctx, operations.AccountingBankFeedStatementsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        BankFeedStatement: components.BankFeedStatementInput{
            BankFeedAccountID: sdkgo.Pointer("acc_456"),
            Status: components.StatementStatusPending.ToPointer(),
            StartDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            EndDate: types.MustNewTimeFromString("2025-01-31T12:00:00.000Z"),
            StartBalance: sdkgo.Pointer[float64](10500.25),
            StartBalanceCreditOrDebit: components.CreditOrDebitDebit.ToPointer(),
            EndBalance: sdkgo.Pointer[float64](9800.5),
            EndBalanceCreditOrDebit: components.CreditOrDebitDebit.ToPointer(),
            Transactions: []components.Transactions{
                components.Transactions{
                    PostedDate: types.MustTimeFromString("2025-01-15T12:00:00.000Z"),
                    Description: sdkgo.Pointer("Payment received from ACME Corp"),
                    Amount: 250,
                    CreditOrDebit: components.CreditOrDebitDebit,
                    SourceTransactionID: "txn_987",
                    Counterparty: sdkgo.Pointer("ACME Corp"),
                    Reference: sdkgo.Pointer("INV-2025-01"),
                    TransactionType: components.BankFeedStatementTransactionTypePayment.ToPointer(),
                },
                components.Transactions{
                    PostedDate: types.MustTimeFromString("2025-01-15T12:00:00.000Z"),
                    Description: sdkgo.Pointer("Payment received from ACME Corp"),
                    Amount: 250,
                    CreditOrDebit: components.CreditOrDebitDebit,
                    SourceTransactionID: "txn_987",
                    Counterparty: sdkgo.Pointer("ACME Corp"),
                    Reference: sdkgo.Pointer("INV-2025-01"),
                    TransactionType: components.BankFeedStatementTransactionTypePayment.ToPointer(),
                },
                components.Transactions{
                    PostedDate: types.MustTimeFromString("2025-01-15T12:00:00.000Z"),
                    Description: sdkgo.Pointer("Payment received from ACME Corp"),
                    Amount: 250,
                    CreditOrDebit: components.CreditOrDebitDebit,
                    SourceTransactionID: "txn_987",
                    Counterparty: sdkgo.Pointer("ACME Corp"),
                    Reference: sdkgo.Pointer("INV-2025-01"),
                    TransactionType: components.BankFeedStatementTransactionTypePayment.ToPointer(),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBankFeedStatementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AccountingBankFeedStatementsAddRequest](../../models/operations/accountingbankfeedstatementsaddrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AccountingBankFeedStatementsAddResponse](../../models/operations/accountingbankfeedstatementsaddresponse.md), error**

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

Get Bank Feed Statement

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankFeedStatementsOne" method="get" path="/accounting/bank-feed-statements/{id}" -->
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

    res, err := s.Accounting.BankFeedStatements.Get(ctx, operations.AccountingBankFeedStatementsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBankFeedStatementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                  | :heavy_check_mark:                                                                                                     | The context to use for the request.                                                                                    |
| `request`                                                                                                              | [operations.AccountingBankFeedStatementsOneRequest](../../models/operations/accountingbankfeedstatementsonerequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |
| `opts`                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                               | :heavy_minus_sign:                                                                                                     | The options for this request.                                                                                          |

### Response

**[*operations.AccountingBankFeedStatementsOneResponse](../../models/operations/accountingbankfeedstatementsoneresponse.md), error**

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

Update Bank Feed Statement

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankFeedStatementsUpdate" method="patch" path="/accounting/bank-feed-statements/{id}" -->
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

    res, err := s.Accounting.BankFeedStatements.Update(ctx, operations.AccountingBankFeedStatementsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        BankFeedStatement: components.BankFeedStatementInput{
            BankFeedAccountID: sdkgo.Pointer("acc_456"),
            Status: components.StatementStatusPending.ToPointer(),
            StartDate: types.MustNewTimeFromString("2021-05-01T12:00:00.000Z"),
            EndDate: types.MustNewTimeFromString("2025-01-31T12:00:00.000Z"),
            StartBalance: sdkgo.Pointer[float64](10500.25),
            StartBalanceCreditOrDebit: components.CreditOrDebitDebit.ToPointer(),
            EndBalance: sdkgo.Pointer[float64](9800.5),
            EndBalanceCreditOrDebit: components.CreditOrDebitDebit.ToPointer(),
            Transactions: []components.Transactions{
                components.Transactions{
                    PostedDate: types.MustTimeFromString("2025-01-15T12:00:00.000Z"),
                    Description: sdkgo.Pointer("Payment received from ACME Corp"),
                    Amount: 250,
                    CreditOrDebit: components.CreditOrDebitDebit,
                    SourceTransactionID: "txn_987",
                    Counterparty: sdkgo.Pointer("ACME Corp"),
                    Reference: sdkgo.Pointer("INV-2025-01"),
                    TransactionType: components.BankFeedStatementTransactionTypePayment.ToPointer(),
                },
                components.Transactions{
                    PostedDate: types.MustTimeFromString("2025-01-15T12:00:00.000Z"),
                    Description: sdkgo.Pointer("Payment received from ACME Corp"),
                    Amount: 250,
                    CreditOrDebit: components.CreditOrDebitDebit,
                    SourceTransactionID: "txn_987",
                    Counterparty: sdkgo.Pointer("ACME Corp"),
                    Reference: sdkgo.Pointer("INV-2025-01"),
                    TransactionType: components.BankFeedStatementTransactionTypePayment.ToPointer(),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateBankFeedStatementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.AccountingBankFeedStatementsUpdateRequest](../../models/operations/accountingbankfeedstatementsupdaterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.AccountingBankFeedStatementsUpdateResponse](../../models/operations/accountingbankfeedstatementsupdateresponse.md), error**

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

Delete Bank Feed Statement

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankFeedStatementsDelete" method="delete" path="/accounting/bank-feed-statements/{id}" -->
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

    res, err := s.Accounting.BankFeedStatements.Delete(ctx, operations.AccountingBankFeedStatementsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBankFeedStatementResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                    | Type                                                                                                                         | Required                                                                                                                     | Description                                                                                                                  |
| ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                        | :heavy_check_mark:                                                                                                           | The context to use for the request.                                                                                          |
| `request`                                                                                                                    | [operations.AccountingBankFeedStatementsDeleteRequest](../../models/operations/accountingbankfeedstatementsdeleterequest.md) | :heavy_check_mark:                                                                                                           | The request object to use for the request.                                                                                   |
| `opts`                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                     | :heavy_minus_sign:                                                                                                           | The options for this request.                                                                                                |

### Response

**[*operations.AccountingBankFeedStatementsDeleteResponse](../../models/operations/accountingbankfeedstatementsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |