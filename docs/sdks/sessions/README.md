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

    res, err := s.Vault.Sessions.Create(ctx, &components.Session{
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
            HideResourceSettings: sdkgo.Bool(false),
            SandboxMode: sdkgo.Bool(false),
            IsolationMode: sdkgo.Bool(false),
            SessionLength: sdkgo.String("1h"),
            ShowLogs: sdkgo.Bool(true),
            ShowSuggestions: sdkgo.Bool(false),
            ShowSidebar: sdkgo.Bool(true),
            AutoRedirect: sdkgo.Bool(false),
            HideGuides: sdkgo.Bool(false),
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Session](../../models/components/session.md) | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

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