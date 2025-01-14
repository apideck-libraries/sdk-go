# ConnectionSettings
(*Vault.ConnectionSettings*)

## Overview

### Available Operations

* [List](#list) - Get resource settings
* [Update](#update) - Update settings

## List

This endpoint returns custom settings and their defaults required by connection for a given resource.


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

    res, err := s.Vault.ConnectionSettings.List(ctx, "crm", "pipedrive", "leads")
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
| `unifiedAPI`                                             | *string*                                                 | :heavy_check_mark:                                       | Unified API                                              | crm                                                      |
| `serviceID`                                              | *string*                                                 | :heavy_check_mark:                                       | Service ID of the resource to return                     | pipedrive                                                |
| `resource`                                               | *string*                                                 | :heavy_check_mark:                                       | Name of the resource (plural)                            | leads                                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.VaultConnectionSettingsAllResponse](../../models/operations/vaultconnectionsettingsallresponse.md), error**

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

Update default values for a connection's resource settings

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

    res, err := s.Vault.ConnectionSettings.Update(ctx, "pipedrive", "crm", "leads", components.ConnectionInput{
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
| `resource`                                                               | *string*                                                                 | :heavy_check_mark:                                                       | Name of the resource (plural)                                            | leads                                                                    |
| `connection`                                                             | [components.ConnectionInput](../../models/components/connectioninput.md) | :heavy_check_mark:                                                       | Fields that need to be updated on the resource                           |                                                                          |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |                                                                          |

### Response

**[*operations.VaultConnectionSettingsUpdateResponse](../../models/operations/vaultconnectionsettingsupdateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |