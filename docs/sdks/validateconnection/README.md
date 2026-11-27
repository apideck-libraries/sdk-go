# Vault.ValidateConnection

## Overview

### Available Operations

* [State](#state) - Validate Connection State

## State

This endpoint checks a connection again and returns its updated state.
For OAuth2 connectors and custom connectors with an OAuth grant (except connectors used only for authentication), this operation first forces the refresh flow for an access token regardless of its expiry.
Then, for connectors that support connection validation (`validation_support: true`), it sends the connector's validation request with the stored credentials.
The validation request is not sent when the connection is disabled, does not have the settings that the request needs, or still needs OAuth authorization.
When the connector has no OAuth refresh and no connection validation, nothing is sent to the integration, and the state is computed from the stored connection.
When the integration rejects the forced token refresh with a non-recoverable error, and no other process saved a newer token in the meantime, the connection is marked as not valid, `vault.connection.token_refresh.failed` is triggered, and the operation returns an error. Its `state` becomes `invalid` only when the connector supports connection validation and no refresh-token retention is active (`credentials_expire_at` is not a future time).

The result is saved on the connection. When the connection `state` changes, `vault.connection.updated` is triggered, and `vault.connection.callable` too when the connection becomes callable. Use this operation to check an `invalid` connection again: reading the connection does not check it.

Note:
  - Do not include any credentials in the request body. The operation uses the credentials stored on the connection.
  - If a refresh token flow was performed and successful, the new access token will then be used for subsequent API requests.
  - A connection that is waiting for confirmation is returned with its stored state, without a refresh or a validation request.


### Example Usage

<!-- UsageSnippet language="go" operationID="vault.validateConnectionState" method="post" path="/vault/connections/{unified_api}/{service_id}/validate" -->
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