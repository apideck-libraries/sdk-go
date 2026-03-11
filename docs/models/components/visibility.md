# Visibility

The visibility of the job

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.VisibilityDraft

// Open enum: custom values can be created with a direct type cast
custom := components.Visibility("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `VisibilityDraft`    | draft                |
| `VisibilityPublic`   | public               |
| `VisibilityInternal` | internal             |