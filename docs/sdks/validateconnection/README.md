# ValidateConnection
(*Vault.ValidateConnection*)

## Overview

### Available Operations

* [State](#state) - Validate Connection State

## State

This endpoint validates the current state of a given connection. This will perform different checks based on the connection auth type. For basic and apiKey auth types, the presence of required fields is checked.
For connectors that implement OAuth2, this operation forces the refresh flow for an access token regardless of its expiry.

Note:
  - Do not include any credentials in the request body. This operation does not persist changes, but only triggers the validation of connection state.
  - If a refresh token flow was performed and successful, the new access token will then be used for subsequent API requests.


### Example Usage

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
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Vault.ValidateConnection.State(ctx, operations.VaultValidateConnectionStateRequest{
        ServiceID: "pipedrive",
        UnifiedAPI: "crm",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ValidateConnectionStateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.VaultValidateConnectionStateRequest](../../models/operations/vaultvalidateconnectionstaterequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.VaultValidateConnectionStateResponse](../../models/operations/vaultvalidateconnectionstateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |