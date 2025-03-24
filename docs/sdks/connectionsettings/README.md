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

    res, err := s.Vault.ConnectionSettings.List(ctx, operations.VaultConnectionSettingsAllRequest{
        UnifiedAPI: "crm",
        ServiceID: "pipedrive",
        Resource: "leads",
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.VaultConnectionSettingsAllRequest](../../models/operations/vaultconnectionsettingsallrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

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

    res, err := s.Vault.ConnectionSettings.Update(ctx, operations.VaultConnectionSettingsUpdateRequest{
        ServiceID: "pipedrive",
        UnifiedAPI: "crm",
        Resource: "leads",
        Connection: components.ConnectionInput{
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
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                                    123,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
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
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
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
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueNumber(
                                10.5,
                            )),
                        },
                        components.ConnectionDefaults{
                            ID: sdkgo.String("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
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
                                        OptionType: components.OptionTypeSimple,
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
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                            123,
                                        )),
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueBoolean(
                                true,
                            )),
                        },
                        components.ConnectionDefaults{
                            ID: sdkgo.String("ProductInterest"),
                            Options: []components.FormFieldOption{
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                            "general",
                                        )),
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                                    123,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                    12.5,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueBoolean(
                                                    true,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
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
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                                    "general",
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                    12.5,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                    12.5,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                                    "general",
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueNumber(
                                10.5,
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
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                            123,
                                        )),
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
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
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueBoolean(
                                                    true,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                                    12.5,
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueNumber(
                                            12.5,
                                        )),
                                        OptionType: components.OptionTypeSimple,
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
                                components.CreateFormFieldOptionGroup(
                                    components.FormFieldOptionGroup{
                                        ID: sdkgo.String("1234"),
                                        Label: "General Channel",
                                        Options: []components.SimpleFormFieldOption{
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
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
                                                OptionType: components.OptionTypeSimple,
                                            },
                                            components.SimpleFormFieldOption{
                                                Label: "General Channel",
                                                Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueStr(
                                                    "general",
                                                )),
                                                OptionType: components.OptionTypeSimple,
                                            },
                                        },
                                        OptionType: components.FormFieldOptionGroupOptionTypeGroup,
                                    },
                                ),
                                components.CreateFormFieldOptionSimple(
                                    components.SimpleFormFieldOption{
                                        Label: "General Channel",
                                        Value: sdkgo.Pointer(components.CreateSimpleFormFieldOptionValueInteger(
                                            123,
                                        )),
                                        OptionType: components.OptionTypeSimple,
                                    },
                                ),
                            },
                            Value: sdkgo.Pointer(components.CreateConnectionValueNumber(
                                10.5,
                            )),
                        },
                    },
                },
            },
            CustomMappings: []components.CustomMappingInput{
                components.CustomMappingInput{
                    Value: sdkgo.String("$.root.training.first_aid"),
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.VaultConnectionSettingsUpdateRequest](../../models/operations/vaultconnectionsettingsupdaterequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

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