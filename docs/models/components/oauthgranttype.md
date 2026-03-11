# OAuthGrantType

OAuth grant type used by the connector. More info: https://oauth.net/2/grant-types

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.OAuthGrantTypeAuthorizationCode

// Open enum: custom values can be created with a direct type cast
custom := components.OAuthGrantType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `OAuthGrantTypeAuthorizationCode` | authorization_code                |
| `OAuthGrantTypeClientCredentials` | client_credentials                |
| `OAuthGrantTypePassword`          | password                          |