# ConnectorDocs
(*Connector.ConnectorDocs*)

## Overview

### Available Operations

* [Get](#get) - Get Connector Doc content

## Get

Get Connector Doc content

### Example Usage

<!-- UsageSnippet language="go" operationID="connector.connectorDocsOne" method="get" path="/connector/connectors/{id}/docs/{doc_id}" -->
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
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Connector.ConnectorDocs.Get(ctx, "<id>", "application_owner+oauth_credentials")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectorDocResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of the record you are acting upon.                    |                                                          |
| `docID`                                                  | *string*                                                 | :heavy_check_mark:                                       | ID of the Doc                                            | application_owner+oauth_credentials                      |
| `appID`                                                  | **string*                                                | :heavy_minus_sign:                                       | The ID of your Unify application                         | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.ConnectorConnectorDocsOneResponse](../../models/operations/connectorconnectordocsoneresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |