# FulfillmentStatus

Current fulfillment status of the order.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FulfillmentStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.FulfillmentStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `FulfillmentStatusPending`   | pending                      |
| `FulfillmentStatusShipped`   | shipped                      |
| `FulfillmentStatusPartial`   | partial                      |
| `FulfillmentStatusDelivered` | delivered                    |
| `FulfillmentStatusCancelled` | cancelled                    |
| `FulfillmentStatusReturned`  | returned                     |
| `FulfillmentStatusUnknown`   | unknown                      |