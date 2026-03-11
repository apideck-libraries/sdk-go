# Priority

Priority level of the project

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PriorityLow

// Open enum: custom values can be created with a direct type cast
custom := components.Priority("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `PriorityLow`      | low                |
| `PriorityMedium`   | medium             |
| `PriorityHigh`     | high               |
| `PriorityCritical` | critical           |