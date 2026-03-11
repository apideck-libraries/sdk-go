# ConnectorAuthType

Type of authorization used by the connector

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConnectorAuthTypeOauth2

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorAuthType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ConnectorAuthTypeOauth2` | oauth2                    |
| `ConnectorAuthTypeAPIKey` | apiKey                    |
| `ConnectorAuthTypeBasic`  | basic                     |
| `ConnectorAuthTypeCustom` | custom                    |
| `ConnectorAuthTypeNone`   | none                      |