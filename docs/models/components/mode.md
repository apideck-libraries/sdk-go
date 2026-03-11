# Mode

Mode of the webhook support.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ModeNative

// Open enum: custom values can be created with a direct type cast
custom := components.Mode("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `ModeNative`  | native        |
| `ModeVirtual` | virtual       |
| `ModeNone`    | none          |