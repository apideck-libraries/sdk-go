# EcommerceOrderStatus

Current status of the order.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EcommerceOrderStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.EcommerceOrderStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `EcommerceOrderStatusActive`    | active                          |
| `EcommerceOrderStatusCompleted` | completed                       |
| `EcommerceOrderStatusCancelled` | cancelled                       |
| `EcommerceOrderStatusArchived`  | archived                        |
| `EcommerceOrderStatusUnknown`   | unknown                         |
| `EcommerceOrderStatusOther`     | other                           |