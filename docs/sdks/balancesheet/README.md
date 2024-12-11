# BalanceSheet
(*Accounting.BalanceSheet*)

## Overview

### Available Operations

* [Get](#get) - Get BalanceSheet

## Get

Get BalanceSheet

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Accounting.BalanceSheet.Get(ctx, sdkgo.String("salesforce"), map[string]any{
        "search": "San Francisco",
    }, &components.BalanceSheetFilter{
        StartDate: sdkgo.String("2021-01-01"),
        EndDate: sdkgo.String("2021-12-31"),
        PeriodCount: sdkgo.Int64(3),
        PeriodType: components.PeriodTypeMonth.ToPointer(),
    }, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetBalanceSheetResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                         | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       | Example                                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                             | :heavy_check_mark:                                                                                                                                | The context to use for the request.                                                                                                               |                                                                                                                                                   |
| `serviceID`                                                                                                                                       | **string*                                                                                                                                         | :heavy_minus_sign:                                                                                                                                | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.     | salesforce                                                                                                                                        |
| `passThrough`                                                                                                                                     | map[string]*any*                                                                                                                                  | :heavy_minus_sign:                                                                                                                                | Optional unmapped key/values that will be passed through to downstream as query parameters. Ie: ?pass_through[search]=leads becomes ?search=leads | {<br/>"search": "San Francisco"<br/>}                                                                                                             |
| `filter`                                                                                                                                          | [*components.BalanceSheetFilter](../../models/components/balancesheetfilter.md)                                                                   | :heavy_minus_sign:                                                                                                                                | Apply filters                                                                                                                                     | {<br/>"start_date": "2021-01-01",<br/>"end_date": "2021-12-31",<br/>"period_count": 3,<br/>"period_type": "month"<br/>}                           |
| `raw`                                                                                                                                             | **bool*                                                                                                                                           | :heavy_minus_sign:                                                                                                                                | Include raw response. Mostly used for debugging purposes                                                                                          |                                                                                                                                                   |
| `opts`                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                          | :heavy_minus_sign:                                                                                                                                | The options for this request.                                                                                                                     |                                                                                                                                                   |

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