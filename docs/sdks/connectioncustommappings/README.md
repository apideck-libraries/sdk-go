# ConnectionCustomMappings
(*Vault.ConnectionCustomMappings*)

## Overview

### Available Operations

* [List](#list) - List connection custom mappings

## List

This endpoint returns a list of custom mappings for a connection.

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Vault.ConnectionCustomMappings.List(ctx, "crm", "pipedrive", "leads", sdkgo.String("1234"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCustomMappingsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                          | Type                                                                                                                                                                               | Required                                                                                                                                                                           | Description                                                                                                                                                                        | Example                                                                                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                              | :heavy_check_mark:                                                                                                                                                                 | The context to use for the request.                                                                                                                                                |                                                                                                                                                                                    |
| `unifiedAPI`                                                                                                                                                                       | *string*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                 | Unified API                                                                                                                                                                        | crm                                                                                                                                                                                |
| `serviceID`                                                                                                                                                                        | *string*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                 | Service ID of the resource to return                                                                                                                                               | pipedrive                                                                                                                                                                          |
| `resource`                                                                                                                                                                         | *string*                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                 | Name of the resource (plural)                                                                                                                                                      | leads                                                                                                                                                                              |
| `resourceID`                                                                                                                                                                       | **string*                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                 | This is the id of the resource you want to fetch when listing custom fields. For example, if you want to fetch custom fields for a specific contact, you would use the contact id. | 1234                                                                                                                                                                               |
| `opts`                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                 | The options for this request.                                                                                                                                                      |                                                                                                                                                                                    |

### Response

**[*operations.VaultConnectionCustomMappingsAllResponse](../../models/operations/vaultconnectioncustommappingsallresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |