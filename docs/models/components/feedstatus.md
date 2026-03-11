# FeedStatus

Current status of the bank feed.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FeedStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.FeedStatus("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `FeedStatusPending`  | pending              |
| `FeedStatusRejected` | rejected             |