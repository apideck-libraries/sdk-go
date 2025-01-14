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

    res, err := s.Vault.Connections.List(ctx, sdkgo.String("crm"), sdkgo.Bool(true))
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | Service ID of the resource to return                     | pipedrive                                                |
| `unifiedAPI`                                             | *string*                                                 | :heavy_check_mark:                                       | Unified API                                              | crm                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

    res, err := s.Vault.Connections.Update(ctx, "pipedrive", "crm", components.ConnectionInput{
        Enabled: sdkgo.Bool(true),
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
                Resource: sdkgo.String("leads"),
                Defaults: []components.ConnectionDefaults{
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{
                            components.CreateFormFieldOptionFormFieldOptionGroup(
                                components.FormFieldOptionGroup{
                                    ID: sdkgo.String("1234"),
                                    Label: sdkgo.String("General Channel"),
                                    Options: []components.SimpleFormFieldOption{
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                12.5,
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueArrayOf5(
                                                []components.Five{
                                                    components.CreateFiveStr(
                                                        "team",
                                                    ),
                                                    components.CreateFiveStr(
                                                        "general",
                                                    ),
                                                },
                                            )),
                                        },
                                    },
                                },
                            ),
                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueStr(
                            "GC5000 series",
                        )),
                    },
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{
                            components.CreateFormFieldOptionSimpleFormFieldOption(
                                components.SimpleFormFieldOption{
                                    Label: sdkgo.String("General Channel"),
                                    Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                        123,
                                    )),
                                },
                            ),
                            components.CreateFormFieldOptionSimpleFormFieldOption(
                                components.SimpleFormFieldOption{
                                    Label: sdkgo.String("General Channel"),
                                    Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                        "general",
                                    )),
                                },
                            ),
                            components.CreateFormFieldOptionFormFieldOptionGroup(
                                components.FormFieldOptionGroup{
                                    ID: sdkgo.String("1234"),
                                    Label: sdkgo.String("General Channel"),
                                    Options: []components.SimpleFormFieldOption{
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                                123,
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                12.5,
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueBoolean(
                                                true,
                                            )),
                                        },
                                    },
                                },
                            ),
                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueBoolean(
                            true,
                        )),
                    },
                },
            },
            components.ConnectionConfiguration{
                Resource: sdkgo.String("leads"),
                Defaults: []components.ConnectionDefaults{
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{

                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueBoolean(
                            true,
                        )),
                    },
                },
            },
            components.ConnectionConfiguration{
                Resource: sdkgo.String("leads"),
                Defaults: []components.ConnectionDefaults{
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{
                            components.CreateFormFieldOptionFormFieldOptionGroup(
                                components.FormFieldOptionGroup{
                                    ID: sdkgo.String("1234"),
                                    Label: sdkgo.String("General Channel"),
                                    Options: []components.SimpleFormFieldOption{

                                    },
                                },
                            ),
                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueInteger(
                            10,
                        )),
                    },
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{
                            components.CreateFormFieldOptionFormFieldOptionGroup(
                                components.FormFieldOptionGroup{
                                    ID: sdkgo.String("1234"),
                                    Label: sdkgo.String("General Channel"),
                                    Options: []components.SimpleFormFieldOption{
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueArrayOf5(
                                                []components.Five{
                                                    components.CreateFiveStr(
                                                        "team",
                                                    ),
                                                    components.CreateFiveStr(
                                                        "general",
                                                    ),
                                                },
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueBoolean(
                                                true,
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                12.5,
                                            )),
                                        },
                                    },
                                },
                            ),
                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueInteger(
                            10,
                        )),
                    },
                    components.ConnectionDefaults{
                        ID: sdkgo.String("ProductInterest"),
                        Options: []components.FormFieldOption{
                            components.CreateFormFieldOptionFormFieldOptionGroup(
                                components.FormFieldOptionGroup{
                                    ID: sdkgo.String("1234"),
                                    Label: sdkgo.String("General Channel"),
                                    Options: []components.SimpleFormFieldOption{
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueArrayOf5(
                                                []components.Five{
                                                    components.CreateFiveStr(
                                                        "team",
                                                    ),
                                                    components.CreateFiveStr(
                                                        "general",
                                                    ),
                                                },
                                            )),
                                        },
                                        components.SimpleFormFieldOption{
                                            Label: sdkgo.String("General Channel"),
                                            Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                                "general",
                                            )),
                                        },
                                    },
                                },
                            ),
                            components.CreateFormFieldOptionSimpleFormFieldOption(
                                components.SimpleFormFieldOption{
                                    Label: sdkgo.String("General Channel"),
                                    Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                        123,
                                    )),
                                },
                            ),
                        },
                        Value: sdkgo.Pointer(components.CreateConnectionValueBoolean(
                            true,
                        )),
                    },
                },
            },
        },
        CustomMappings: []components.CustomMappingInput{

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

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |                                                                          |
| `serviceID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | Service ID of the resource to return                                     | pipedrive                                                                |
| `unifiedAPI`                                                             | *string*                                                                 | :heavy_check_mark:                                                       | Unified API                                                              | crm                                                                      |
| `connection`                                                             | [components.ConnectionInput](../../models/components/connectioninput.md) | :heavy_check_mark:                                                       | Fields that need to be updated on the resource                           |                                                                          |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

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

    res, err := s.Vault.Connections.Delete(ctx, "pipedrive", "crm")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | Service ID of the resource to return                     | pipedrive                                                |
| `unifiedAPI`                                             | *string*                                                 | :heavy_check_mark:                                       | Unified API                                              | crm                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

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

    res, err := s.Vault.Connections.Imports(ctx, "pipedrive", "crm", components.ConnectionImportData{
        Credentials: &components.Credentials{
            RefreshToken: sdkgo.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.cThIIoDvwdueQB468K5xDc5633seEFoqwxjF_xSJyQQ"),
            AccessToken: sdkgo.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"),
        },
        Metadata: map[string]any{
            "account": map[string]any{
                "name": "My Company",
                "id": "c01458a5-7276-41ce-bc19-639906b0450a",
            },
            "plan": "enterprise",
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        | Example                                                                            |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |                                                                                    |
| `serviceID`                                                                        | *string*                                                                           | :heavy_check_mark:                                                                 | Service ID of the resource to return                                               | pipedrive                                                                          |
| `unifiedAPI`                                                                       | *string*                                                                           | :heavy_check_mark:                                                                 | Unified API                                                                        | crm                                                                                |
| `connectionImportData`                                                             | [components.ConnectionImportData](../../models/components/connectionimportdata.md) | :heavy_check_mark:                                                                 | Fields that need to be persisted on the resource                                   |                                                                                    |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |                                                                                    |

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

    res, err := s.Vault.Connections.Token(ctx, "pipedrive", "crm", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.GetConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 | Example                                                                                                     |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |                                                                                                             |
| `serviceID`                                                                                                 | *string*                                                                                                    | :heavy_check_mark:                                                                                          | Service ID of the resource to return                                                                        | pipedrive                                                                                                   |
| `unifiedAPI`                                                                                                | *string*                                                                                                    | :heavy_check_mark:                                                                                          | Unified API                                                                                                 | crm                                                                                                         |
| `requestBody`                                                                                               | [*operations.VaultConnectionsTokenRequestBody](../../models/operations/vaultconnectionstokenrequestbody.md) | :heavy_minus_sign:                                                                                          | N/A                                                                                                         |                                                                                                             |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |                                                                                                             |

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