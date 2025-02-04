# Sessions
(*Vault.Sessions*)

## Overview

### Available Operations

* [Create](#create) - Create Session

## Create

Making a POST request to this endpoint will initiate a Hosted Vault session. Redirect the consumer to the returned
URL to allow temporary access to manage their integrations and settings.

Note: This is a short lived token that will expire after 1 hour (TTL: 3600).


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

    res, err := s.Vault.Sessions.Create(ctx, sdkgo.String("test-consumer"), sdkgo.String("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"), &components.Session{
        ConsumerMetadata: &components.ConsumerMetadata{
            AccountName: sdkgo.String("SpaceX"),
            UserName: sdkgo.String("Elon Musk"),
            Email: sdkgo.String("elon@musk.com"),
            Image: sdkgo.String("https://www.spacex.com/static/images/share.jpg"),
        },
        RedirectURI: sdkgo.String("https://mysaas.com/dashboard"),
        Settings: &components.SessionSettings{
            UnifiedApis: []components.UnifiedAPIID{
                components.UnifiedAPIIDCrm,
            },
        },
        Theme: &components.Theme{
            Favicon: sdkgo.String("https://res.cloudinary.com/apideck/icons/intercom"),
            Logo: sdkgo.String("https://res.cloudinary.com/apideck/icons/intercom"),
            PrimaryColor: sdkgo.String("#286efa"),
            SidepanelBackgroundColor: sdkgo.String("#286efa"),
            SidepanelTextColor: sdkgo.String("#FFFFFF"),
            VaultName: sdkgo.String("Intercom"),
            PrivacyURL: sdkgo.String("https://compliance.apideck.com/privacy-policy"),
            TermsURL: sdkgo.String("https://www.termsfeed.com/terms-conditions/957c85c1b089ae9e3219c83eff65377e"),
        },
        CustomConsumerSettings: map[string]any{
            "feature_flag_1": true,
            "tax_rates": []any{
                map[string]any{
                    "id": "6",
                    "label": "6%",
                },
                map[string]any{
                    "id": "21",
                    "label": "21%",
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSessionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                | Example                                                    |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |                                                            |
| `consumerID`                                               | **string*                                                  | :heavy_minus_sign:                                         | ID of the consumer which you want to get or push data from | test-consumer                                              |
| `appID`                                                    | **string*                                                  | :heavy_minus_sign:                                         | The ID of your Unify application                           | dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX                    |
| `session`                                                  | [*components.Session](../../models/components/session.md)  | :heavy_minus_sign:                                         | Additional redirect uri and/or consumer metadata           |                                                            |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |                                                            |

### Response

**[*operations.VaultSessionsCreateResponse](../../models/operations/vaultsessionscreateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |