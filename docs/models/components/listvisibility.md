# ListVisibility

The visibility of the List. Which of these values a given connector can return depends on its native sharing model — see the connector-specific gotchas below for details.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ListVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.ListVisibility("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ListVisibilityPrivate` | private                 |
| `ListVisibilityShared`  | shared                  |
| `ListVisibilityPublic`  | public                  |