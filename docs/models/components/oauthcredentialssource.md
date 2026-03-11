# OauthCredentialsSource

Location of the OAuth client credentials. For most connectors the OAuth client credentials are stored on integration and managed by the application owner. For others they are stored on connection and managed by the consumer in Vault.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.OauthCredentialsSourceIntegration

// Open enum: custom values can be created with a direct type cast
custom := components.OauthCredentialsSource("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `OauthCredentialsSourceIntegration` | integration                         |
| `OauthCredentialsSourceConnection`  | connection                          |