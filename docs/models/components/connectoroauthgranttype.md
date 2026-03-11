# ConnectorOauthGrantType

OAuth grant type used by the connector. More info: https://oauth.net/2/grant-types

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectorOauthGrantTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorOauthGrantType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ConnectorOauthGrantTypeAuthorizationCode` | authorization_code                         |
| `ConnectorOauthGrantTypeClientCredentials` | client_credentials                         |
| `ConnectorOauthGrantTypePassword`          | password                                   |