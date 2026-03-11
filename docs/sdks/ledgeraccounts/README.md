# Accounting.LedgerAccounts

## Overview

### Available Operations

* [List](#list) - List Ledger Accounts
* [Create](#create) - Create Ledger Account
* [Get](#get) - Get Ledger Account
* [Update](#update) - Update Ledger Account
* [Delete](#delete) - Delete Ledger Account

## List

List Ledger Accounts

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.ledgerAccountsAll" method="get" path="/accounting/ledger-accounts" -->
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

    res, err := s.Accounting.LedgerAccounts.List(ctx, operations.AccountingLedgerAccountsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.LedgerAccountsFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.LedgerAccountsSort{
            By: components.LedgerAccountsSortByUpdatedAt.ToPointer(),
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
    if res.GetLedgerAccountsResponse != nil {
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
| `request`                                                                                                      | [operations.AccountingLedgerAccountsAllRequest](../../models/operations/accountingledgeraccountsallrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingLedgerAccountsAllResponse](../../models/operations/accountingledgeraccountsallresponse.md), error**

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

Create Ledger Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.ledgerAccountsAdd" method="post" path="/accounting/ledger-accounts" -->
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

    res, err := s.Accounting.LedgerAccounts.Create(ctx, operations.AccountingLedgerAccountsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        LedgerAccount: components.LedgerAccountInput{
            DisplayID: sdkgo.Pointer("1-12345"),
            Code: sdkgo.Pointer("453"),
            Classification: components.LedgerAccountClassificationAsset.ToPointer(),
            Type: components.LedgerAccountTypeBank.ToPointer(),
            SubType: sdkgo.Pointer("CHECKING_ACCOUNT"),
            Name: sdkgo.Pointer("Bank account"),
            FullyQualifiedName: sdkgo.Pointer("Asset.Bank.Checking_Account"),
            Description: sdkgo.Pointer("Main checking account"),
            OpeningBalance: sdkgo.Pointer[float64](75000),
            CurrentBalance: sdkgo.Pointer[float64](20000),
            Currency: components.CurrencyUsd.ToPointer(),
            TaxType: sdkgo.Pointer("NONE"),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Rate: sdkgo.Pointer[float64](10),
            },
            Level: sdkgo.Pointer[float64](1),
            Active: sdkgo.Pointer(true),
            Status: components.AccountStatusActive.ToPointer(),
            Header: sdkgo.Pointer(true),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.Pointer("Monzo"),
                AccountNumber: sdkgo.Pointer("123465"),
                AccountName: sdkgo.Pointer("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.Pointer("CH2989144532982975332"),
                Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.Pointer("012345678"),
                BsbNumber: sdkgo.Pointer("062-001"),
                BranchIdentifier: sdkgo.Pointer("001"),
                BankCode: sdkgo.Pointer("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            ParentAccount: &components.ParentAccount{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Bank Accounts"),
                DisplayID: sdkgo.Pointer("1-1100"),
            },
            SubAccount: sdkgo.Pointer(false),
            LastReconciliationDate: types.MustNewDateFromString("2020-09-30"),
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
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateLedgerAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingLedgerAccountsAddRequest](../../models/operations/accountingledgeraccountsaddrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingLedgerAccountsAddResponse](../../models/operations/accountingledgeraccountsaddresponse.md), error**

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

Get Ledger Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.ledgerAccountsOne" method="get" path="/accounting/ledger-accounts/{id}" -->
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

    res, err := s.Accounting.LedgerAccounts.Get(ctx, operations.AccountingLedgerAccountsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLedgerAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingLedgerAccountsOneRequest](../../models/operations/accountingledgeraccountsonerequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingLedgerAccountsOneResponse](../../models/operations/accountingledgeraccountsoneresponse.md), error**

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

Update Ledger Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.ledgerAccountsUpdate" method="patch" path="/accounting/ledger-accounts/{id}" -->
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

    res, err := s.Accounting.LedgerAccounts.Update(ctx, operations.AccountingLedgerAccountsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        LedgerAccount: components.LedgerAccountInput{
            DisplayID: sdkgo.Pointer("1-12345"),
            Code: sdkgo.Pointer("453"),
            Classification: components.LedgerAccountClassificationAsset.ToPointer(),
            Type: components.LedgerAccountTypeBank.ToPointer(),
            SubType: sdkgo.Pointer("CHECKING_ACCOUNT"),
            Name: sdkgo.Pointer("Bank account"),
            FullyQualifiedName: sdkgo.Pointer("Asset.Bank.Checking_Account"),
            Description: sdkgo.Pointer("Main checking account"),
            OpeningBalance: sdkgo.Pointer[float64](75000),
            CurrentBalance: sdkgo.Pointer[float64](20000),
            Currency: components.CurrencyUsd.ToPointer(),
            TaxType: sdkgo.Pointer("NONE"),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Rate: sdkgo.Pointer[float64](10),
            },
            Level: sdkgo.Pointer[float64](1),
            Active: sdkgo.Pointer(true),
            Status: components.AccountStatusActive.ToPointer(),
            Header: sdkgo.Pointer(true),
            BankAccount: &components.BankAccount{
                BankName: sdkgo.Pointer("Monzo"),
                AccountNumber: sdkgo.Pointer("123465"),
                AccountName: sdkgo.Pointer("SPACEX LLC"),
                AccountType: components.AccountTypeCreditCard.ToPointer(),
                Iban: sdkgo.Pointer("CH2989144532982975332"),
                Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                RoutingNumber: sdkgo.Pointer("012345678"),
                BsbNumber: sdkgo.Pointer("062-001"),
                BranchIdentifier: sdkgo.Pointer("001"),
                BankCode: sdkgo.Pointer("BNH"),
                Currency: components.CurrencyUsd.ToPointer(),
            },
            ParentAccount: &components.ParentAccount{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Bank Accounts"),
                DisplayID: sdkgo.Pointer("1-1100"),
            },
            SubAccount: sdkgo.Pointer(false),
            LastReconciliationDate: types.MustNewDateFromString("2020-09-30"),
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
    if res.UpdateLedgerAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingLedgerAccountsUpdateRequest](../../models/operations/accountingledgeraccountsupdaterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingLedgerAccountsUpdateResponse](../../models/operations/accountingledgeraccountsupdateresponse.md), error**

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

Delete Ledger Account

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.ledgerAccountsDelete" method="delete" path="/accounting/ledger-accounts/{id}" -->
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

    res, err := s.Accounting.LedgerAccounts.Delete(ctx, operations.AccountingLedgerAccountsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteLedgerAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                            | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                                | :heavy_check_mark:                                                                                                   | The context to use for the request.                                                                                  |
| `request`                                                                                                            | [operations.AccountingLedgerAccountsDeleteRequest](../../models/operations/accountingledgeraccountsdeleterequest.md) | :heavy_check_mark:                                                                                                   | The request object to use for the request.                                                                           |
| `opts`                                                                                                               | [][operations.Option](../../models/operations/option.md)                                                             | :heavy_minus_sign:                                                                                                   | The options for this request.                                                                                        |

### Response

**[*operations.AccountingLedgerAccountsDeleteResponse](../../models/operations/accountingledgeraccountsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |