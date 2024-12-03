# CreateCallback
(*Vault.CreateCallback*)

## Overview

### Available Operations

* [State](#state) - Create Callback State

## State

This endpoint creates a callback state that can be used to issue requests to the callback endpoint.


### Example Usage

```go
package main

import(
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
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
    res, err := s.Vault.CreateCallback.State(ctx, "pipedrive", "crm", components.CreateCallbackState{
        RedirectURI: sdkgo.String("https://example.com/callback"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCallbackStateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |                                                                                  |
| `serviceID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | Service ID of the resource to return                                             | pipedrive                                                                        |
| `unifiedAPI`                                                                     | *string*                                                                         | :heavy_check_mark:                                                               | Unified API                                                                      | crm                                                                              |
| `createCallbackState`                                                            | [components.CreateCallbackState](../../models/components/createcallbackstate.md) | :heavy_check_mark:                                                               | Callback state data                                                              |                                                                                  |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |                                                                                  |

### Response

**[*operations.VaultCreateCallbackStateResponse](../../models/operations/vaultcreatecallbackstateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |