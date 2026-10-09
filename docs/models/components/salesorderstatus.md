# SalesOrderStatus

Sales order status, in order of precedence: `cancelled`; `closed` (the order is closed or completed, whether or not it was billed); `invoiced` (fully billed but not yet closed); `back_ordered`; `on_hold` (including credit hold); `draft` (including pending approval); `open` (every other active state, including partially shipped and partially invoiced); `other` for states that fit none of these.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SalesOrderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.SalesOrderStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SalesOrderStatusDraft`       | draft                         |
| `SalesOrderStatusOpen`        | open                          |
| `SalesOrderStatusOnHold`      | on_hold                       |
| `SalesOrderStatusBackOrdered` | back_ordered                  |
| `SalesOrderStatusInvoiced`    | invoiced                      |
| `SalesOrderStatusClosed`      | closed                        |
| `SalesOrderStatusCancelled`   | cancelled                     |
| `SalesOrderStatusOther`       | other                         |