# CustomerCreateInputStatus

Customer status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CustomerCreateInputStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.CustomerCreateInputStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `CustomerCreateInputStatusActive`             | active                                        |
| `CustomerCreateInputStatusInactive`           | inactive                                      |
| `CustomerCreateInputStatusArchived`           | archived                                      |
| `CustomerCreateInputStatusGdprErasureRequest` | gdpr-erasure-request                          |
| `CustomerCreateInputStatusUnknown`            | unknown                                       |