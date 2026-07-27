# ListType

Whether the List is static (a fixed set of records) or dynamic (a saved segment that is automatically kept up to date based on filter criteria).

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ListTypeStatic

// Open enum: custom values can be created with a direct type cast
custom := components.ListType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `ListTypeStatic`  | static            |
| `ListTypeDynamic` | dynamic           |