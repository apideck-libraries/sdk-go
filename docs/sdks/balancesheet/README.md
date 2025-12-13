# Accounting.BalanceSheet

## Overview

### Available Operations

* [Get](#get) - Get BalanceSheet

## Get

Get BalanceSheet

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.balanceSheetOne" method="get" path="/accounting/balance-sheet" -->
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

    res, err := s.Accounting.BalanceSheet.Get(ctx, operations.AccountingBalanceSheetOneRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Filter: &components.BalanceSheetFilter{
            StartDate: sdkgo.Pointer("2021-01-01"),
            EndDate: sdkgo.Pointer("2021-12-31"),
            PeriodCount: sdkgo.Pointer[int64](3),
            PeriodType: components.PeriodTypeMonth.ToPointer(),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBalanceSheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingBalanceSheetOneRequest](../../models/operations/accountingbalancesheetonerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingBalanceSheetOneResponse](../../models/operations/accountingbalancesheetoneresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |