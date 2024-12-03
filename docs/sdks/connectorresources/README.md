# ConnectorResources
(*Connector.ConnectorResources*)

## Overview

### Available Operations

* [Get](#get) - Get Connector Resource

## Get

Get Connector Resource

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
    res, err := s.Connector.ConnectorResources.Get(ctx, "<id>", "<id>", components.UnifiedAPIIDCrm.ToPointer())
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectorResourceResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |                                                                                                         |
| `id`                                                                                                    | *string*                                                                                                | :heavy_check_mark:                                                                                      | ID of the record you are acting upon.                                                                   |                                                                                                         |
| `resourceID`                                                                                            | *string*                                                                                                | :heavy_check_mark:                                                                                      | ID of the resource you are acting upon.                                                                 |                                                                                                         |
| `unifiedAPI`                                                                                            | [*components.UnifiedAPIID](../../models/components/unifiedapiid.md)                                     | :heavy_minus_sign:                                                                                      | Specify unified API for the connector resource. This is useful when a resource appears in multiple APIs | crm                                                                                                     |
| `opts`                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                | :heavy_minus_sign:                                                                                      | The options for this request.                                                                           |                                                                                                         |

### Response

**[*operations.ConnectorConnectorResourcesOneResponse](../../models/operations/connectorconnectorresourcesoneresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |