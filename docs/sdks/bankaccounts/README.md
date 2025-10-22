# BankAccounts
(*Accounting.BankAccounts*)

## Overview

### Available Operations

* [List](#list) - List Bank Accounts
* [Create](#create) - Create Bank Account
* [Get](#get) - Get Bank Account
* [Update](#update) - Update Bank Account
* [Delete](#delete) - Delete Bank Account

## List

List Bank Accounts

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankAccountsAll" method="get" path="/accounting/bank-accounts" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Accounting.BankAccounts.List(ctx, operations.AccountingBankAccountsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.BankAccountsFilter{
            Name: sdkgo.Pointer("Main Operating"),
            Status: components.BankAccountsFilterStatusActive.ToPointer(),
        },
        Sort: &components.BankAccountsSort{},
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBankAccountsResponse != nil {
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingBankAccountsAllRequest](../../models/operations/accountingbankaccountsallrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingBankAccountsAllResponse](../../models/operations/accountingbankaccountsallresponse.md), error**

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

Create Bank Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankAccountsAdd" method="post" path="/accounting/bank-accounts" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Accounting.BankAccounts.Create(ctx, operations.AccountingBankAccountsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        AccountingBankAccount: components.AccountingBankAccountInput{
            DisplayID: sdkgo.Pointer("BA-001"),
            Name: sdkgo.Pointer("Main Operating Account"),
            AccountNumber: sdkgo.Pointer("123465"),
            AccountType: components.AccountingBankAccountAccountTypeChecking.ToPointer(),
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            BankName: sdkgo.Pointer("Chase Bank"),
            Currency: components.CurrencyUsd.ToPointer(),
            Balance: sdkgo.Pointer[float64](25000),
            AvailableBalance: sdkgo.Pointer[float64](24500),
            OverdraftLimit: sdkgo.Pointer[float64](5000),
            RoutingNumber: sdkgo.Pointer("021000021"),
            Iban: sdkgo.Pointer("GB33BUKB20201555555555"),
            Bic: sdkgo.Pointer("CHASUS33"),
            BsbNumber: sdkgo.Pointer("062-001"),
            BranchIdentifier: sdkgo.Pointer("001"),
            BankCode: sdkgo.Pointer("BNH"),
            Country: sdkgo.Pointer("US"),
            Status: components.AccountingBankAccountStatusActive.ToPointer(),
            Description: sdkgo.Pointer("Primary operating account for daily transactions"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingBankAccountsAddRequest](../../models/operations/accountingbankaccountsaddrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingBankAccountsAddResponse](../../models/operations/accountingbankaccountsaddresponse.md), error**

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

Get Bank Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankAccountsOne" method="get" path="/accounting/bank-accounts/{id}" -->
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

    res, err := s.Accounting.BankAccounts.Get(ctx, operations.AccountingBankAccountsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingBankAccountsOneRequest](../../models/operations/accountingbankaccountsonerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingBankAccountsOneResponse](../../models/operations/accountingbankaccountsoneresponse.md), error**

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

Update Bank Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankAccountsUpdate" method="patch" path="/accounting/bank-accounts/{id}" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Accounting.BankAccounts.Update(ctx, operations.AccountingBankAccountsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        AccountingBankAccount: components.AccountingBankAccountInput{
            DisplayID: sdkgo.Pointer("BA-001"),
            Name: sdkgo.Pointer("Main Operating Account"),
            AccountNumber: sdkgo.Pointer("123465"),
            AccountType: components.AccountingBankAccountAccountTypeChecking.ToPointer(),
            LedgerAccount: &components.LinkedLedgerAccountInput{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
                ParentID: sdkgo.Pointer("123456"),
                DisplayID: sdkgo.Pointer("123456"),
            },
            BankName: sdkgo.Pointer("Chase Bank"),
            Currency: components.CurrencyUsd.ToPointer(),
            Balance: sdkgo.Pointer[float64](25000),
            AvailableBalance: sdkgo.Pointer[float64](24500),
            OverdraftLimit: sdkgo.Pointer[float64](5000),
            RoutingNumber: sdkgo.Pointer("021000021"),
            Iban: sdkgo.Pointer("GB33BUKB20201555555555"),
            Bic: sdkgo.Pointer("CHASUS33"),
            BsbNumber: sdkgo.Pointer("062-001"),
            BranchIdentifier: sdkgo.Pointer("001"),
            BankCode: sdkgo.Pointer("BNH"),
            Country: sdkgo.Pointer("US"),
            Status: components.AccountingBankAccountStatusActive.ToPointer(),
            Description: sdkgo.Pointer("Primary operating account for daily transactions"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingBankAccountsUpdateRequest](../../models/operations/accountingbankaccountsupdaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingBankAccountsUpdateResponse](../../models/operations/accountingbankaccountsupdateresponse.md), error**

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

Delete Bank Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.bankAccountsDelete" method="delete" path="/accounting/bank-accounts/{id}" -->
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

    res, err := s.Accounting.BankAccounts.Delete(ctx, operations.AccountingBankAccountsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteBankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.AccountingBankAccountsDeleteRequest](../../models/operations/accountingbankaccountsdeleterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.AccountingBankAccountsDeleteResponse](../../models/operations/accountingbankaccountsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |