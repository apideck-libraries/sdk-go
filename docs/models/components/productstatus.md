# ProductStatus

The current status of the product (active or archived).

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ProductStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.ProductStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ProductStatusActive`   | active                  |
| `ProductStatusArchived` | archived                |