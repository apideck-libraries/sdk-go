# PathMatchMode

How the path filter is matched. CONTAINS matches the path anywhere; STARTS_WITH / ENDS_WITH anchor the match; EXACT requires the whole path to match. Only applied when path is set.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PathMatchModeContains

// Open enum: custom values can be created with a direct type cast
custom := components.PathMatchMode("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `PathMatchModeContains`   | CONTAINS                  |
| `PathMatchModeStartsWith` | STARTS_WITH               |
| `PathMatchModeEndsWith`   | ENDS_WITH                 |
| `PathMatchModeExact`      | EXACT                     |