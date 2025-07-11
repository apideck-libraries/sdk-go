# ConsumerRequestCounts
(*Vault.ConsumerRequestCounts*)

## Overview

### Available Operations

* [List](#list) - Consumer request counts

## List

Get consumer request counts within a given datetime range.


### Example Usage

```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Vault.ConsumerRequestCounts.List(ctx, "test_user_id", "2021-05-01T12:00:00.000Z", "2021-05-30T12:00:00.000Z")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConsumerRequestCountsInDateRangeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `consumerID`                                             | *string*                                                 | :heavy_check_mark:                                       | ID of the consumer to return                             | test_user_id                                             |
| `startDatetime`                                          | *string*                                                 | :heavy_check_mark:                                       | Scopes results to requests that happened after datetime  | 2021-05-01T12:00:00.000Z                                 |
| `endDatetime`                                            | *string*                                                 | :heavy_check_mark:                                       | Scopes results to requests that happened before datetime | 2021-05-30T12:00:00.000Z                                 |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | The ID of your Unify application                         | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.VaultConsumerRequestCountsAllResponse](../../models/operations/vaultconsumerrequestcountsallresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |