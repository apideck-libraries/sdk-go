# AuthType

Type of authorization used by the connector

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AuthTypeOauth2

// Open enum: custom values can be created with a direct type cast
custom := components.AuthType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `AuthTypeOauth2` | oauth2           |
| `AuthTypeAPIKey` | apiKey           |
| `AuthTypeBasic`  | basic            |
| `AuthTypeCustom` | custom           |
| `AuthTypeNone`   | none             |