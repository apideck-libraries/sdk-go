# TrackingCategoriesMode

The mode of tracking categories for the company on transactions

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TrackingCategoriesModeTransaction

// Open enum: custom values can be created with a direct type cast
custom := components.TrackingCategoriesMode("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `TrackingCategoriesModeTransaction` | transaction                         |
| `TrackingCategoriesModeLine`        | line                                |
| `TrackingCategoriesModeBoth`        | both                                |
| `TrackingCategoriesModeDisabled`    | disabled                            |