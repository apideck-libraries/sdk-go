# APIType

Indicates whether the API is a Unified API. If unified_api is false, the API is a Platform API.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.APITypePlatform

// Open enum: custom values can be created with a direct type cast
custom := components.APIType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `APITypePlatform` | platform          |
| `APITypeUnified`  | unified           |