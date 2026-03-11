# PurchaseOrderStatus

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PurchaseOrderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.PurchaseOrderStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PurchaseOrderStatusDraft`   | draft                        |
| `PurchaseOrderStatusOpen`    | open                         |
| `PurchaseOrderStatusClosed`  | closed                       |
| `PurchaseOrderStatusDeleted` | deleted                      |
| `PurchaseOrderStatusBilled`  | billed                       |
| `PurchaseOrderStatusOther`   | other                        |