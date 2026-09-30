# BatchSupportResourcesMode

`none` means this resource refuses batch writes. `native` satisfies a batch in a single downstream call against the provider's own batch endpoint, so the request counts as one request against your plan. `loop` satisfies it as a bounded sequential fan-out, one downstream call per item — so a request of N records takes roughly N times as long and counts as N requests.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BatchSupportResourcesModeNone

// Open enum: custom values can be created with a direct type cast
custom := components.BatchSupportResourcesMode("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `BatchSupportResourcesModeNone`   | none                              |
| `BatchSupportResourcesModeNative` | native                            |
| `BatchSupportResourcesModeLoop`   | loop                              |