# Connections
(*Vault.Connections*)

## Overview

### Available Operations

* [List](#list) - Get all connections
* [Get](#get) - Get connection
* [Update](#update) - Update connection
* [Delete](#delete) - Deletes a connection
* [Imports](#imports) - Import connection
* [Token](#token) - Authorize Access Token

## List

This endpoint includes all the configured integrations and contains the required assets
to build an integrations page where your users can install integrations.
OAuth2 supported integrations will contain authorize and revoke links to handle the authentication flows.


### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsAll" method="get" path="/vault/connections" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Vault.Connections.List(ctx, sdkgo.Pointer("crm"), sdkgo.Pointer(true))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectionsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |                                                                |
| `consumerID`                                                   | **string*                                                      | :heavy_minus_sign:                                             | ID of the consumer which you want to get or push data from     | test-consumer                                                  |
| `appID`                                                        | **string*                                                      | :heavy_minus_sign:                                             | The ID of your Unify application                               | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                        |
| `api`                                                          | **string*                                                      | :heavy_minus_sign:                                             | Scope results to Unified API                                   | crm                                                            |
| `configured`                                                   | **bool*                                                        | :heavy_minus_sign:                                             | Scopes results to connections that have been configured or not | true                                                           |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |                                                                |

### Response

**[*operations.VaultConnectionsAllResponse](../../models/operations/vaultconnectionsallresponse.md), error**

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

Get a connection

### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsOne" method="get" path="/vault/connections/{unified_api}/{service_id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Vault.Connections.Get(ctx, "pipedrive", "crm")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `serviceID`                                                | *string*                                                   | :heavy_check_mark:                                         | Service ID of the resource to return                       | pipedrive                                                  |
| `unifiedAPI`                                               | *string*                                                   | :heavy_check_mark:                                         | Unified API                                                | crm                                                        |
| `consumerID`                                               | **string*                                                  | :heavy_minus_sign:                                         | ID of the consumer which you want to get or push data from | test-consumer                                              |
| `appID`                                                    | **string*                                                  | :heavy_minus_sign:                                         | The ID of your Unify application                           | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                    |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.VaultConnectionsOneResponse](../../models/operations/vaultconnectionsoneresponse.md), error**

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

Update a connection

### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsUpdate" method="patch" path="/vault/connections/{unified_api}/{service_id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Vault.Connections.Update(ctx, operations.VaultConnectionsUpdateRequest{
        ServiceID: "pipedrive",
        UnifiedAPI: "crm",
        Connection: components.ConnectionInput{
            Enabled: sdkgo.Pointer(true),
            Settings: map[string]any{
                "instance_url": "https://eu28.salesforce.com",
                "api_key": "12345xxxxxx",
            },
            Metadata: map[string]any{
                "account": map[string]any{
                    "name": "My Company",
                    "id": "c01458a5-7276-41ce-bc19-639906b0450a",
                },
                "plan": "enterprise",
            },
            Configuration: []components.ConnectionConfiguration{
                components.ConnectionConfiguration{
                    Resource: sdkgo.Pointer("leads"),
                    Defaults: []components.ConnectionDefaults{
                        components.ConnectionDefaults{
                            ID: sdkgo.Pointer("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{},
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueStr(
                                "GC5000 series",
                            )),
                        },
                        components.ConnectionDefaults{
                            ID: sdkgo.Pointer("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{},
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueStr(
                                "GC5000 series",
                            )),
                        },
                    },
                },
                components.ConnectionConfiguration{
                    Resource: sdkgo.Pointer("leads"),
                    Defaults: []components.ConnectionDefaults{
                        components.ConnectionDefaults{
                            ID: sdkgo.Pointer("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{},
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueStr(
                                "GC5000 series",
                            )),
                        },
                        components.ConnectionDefaults{
                            ID: sdkgo.Pointer("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{},
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueStr(
                                "GC5000 series",
                            )),
                        },
                    },
                },
            },
            CustomMappings: []components.CustomMappingInput{
                components.CustomMappingInput{
                    Value: sdkgo.Pointer("$.root.training.first_aid"),
                },
                components.CustomMappingInput{
                    Value: sdkgo.Pointer("$.root.training.first_aid"),
                },
                components.CustomMappingInput{
                    Value: sdkgo.Pointer("$.root.training.first_aid"),
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.VaultConnectionsUpdateRequest](../../models/operations/vaultconnectionsupdaterequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.VaultConnectionsUpdateResponse](../../models/operations/vaultconnectionsupdateresponse.md), error**

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

Deletes a connection

### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsDelete" method="delete" path="/vault/connections/{unified_api}/{service_id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Vault.Connections.Delete(ctx, "pipedrive", "crm")
    if err != nil {
        log.Fatal(err)
    }
    if res.UnexpectedErrorResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `serviceID`                                                | *string*                                                   | :heavy_check_mark:                                         | Service ID of the resource to return                       | pipedrive                                                  |
| `unifiedAPI`                                               | *string*                                                   | :heavy_check_mark:                                         | Unified API                                                | crm                                                        |
| `consumerID`                                               | **string*                                                  | :heavy_minus_sign:                                         | ID of the consumer which you want to get or push data from | test-consumer                                              |
| `appID`                                                    | **string*                                                  | :heavy_minus_sign:                                         | The ID of your Unify application                           | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                    |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.VaultConnectionsDeleteResponse](../../models/operations/vaultconnectionsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Imports

Import an authorized connection.


### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsImport" method="post" path="/vault/connections/{unified_api}/{service_id}/import" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Vault.Connections.Imports(ctx, operations.VaultConnectionsImportRequest{
        ServiceID: "pipedrive",
        UnifiedAPI: "crm",
        ConnectionImportData: components.ConnectionImportData{
            Credentials: &components.Credentials{
                RefreshToken: sdkgo.Pointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cThIIoDvwdueQB468K5xDc5633seEFoqwxjF_xSJyQQ"),
                AccessToken: sdkgo.Pointer("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"),
            },
            Settings: map[string]any{
                "instance_url": "https://eu28.salesforce.com",
            },
            Metadata: map[string]any{
                "account": map[string]any{
                    "name": "My Company",
                    "id": "c01458a5-7276-41ce-bc19-639906b0450a",
                },
                "plan": "enterprise",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.VaultConnectionsImportRequest](../../models/operations/vaultconnectionsimportrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.VaultConnectionsImportResponse](../../models/operations/vaultconnectionsimportresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Token

Triggers exchanging persisted connection credentials for an access token and store it in Vault. Currently supported for connections with the `client_credentials` or `password` OAuth grant type.

Note:
  - Do not include any credentials in the request body. This operation does not persist changes, but only triggers the exchange of persisted connection credentials for an access token.
  - The access token will not be returned in the response. A 200 response code indicates the authorization was successful and that a valid access token was stored on the connection.
  - The access token will be used for subsequent API requests.


### Example Usage

<!-- UsageSnippet language="go" operationID="vault.connectionsToken" method="post" path="/vault/connections/{unified_api}/{service_id}/token" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Vault.Connections.Token(ctx, operations.VaultConnectionsTokenRequest{
        ServiceID: "pipedrive",
        UnifiedAPI: "crm",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.VaultConnectionsTokenRequest](../../models/operations/vaultconnectionstokenrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.VaultConnectionsTokenResponse](../../models/operations/vaultconnectionstokenresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |