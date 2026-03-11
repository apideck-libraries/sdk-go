# TicketPriority

Priority of the ticket

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TicketPriorityLow

// Open enum: custom values can be created with a direct type cast
custom := components.TicketPriority("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TicketPriorityLow`    | low                    |
| `TicketPriorityNormal` | normal                 |
| `TicketPriorityHigh`   | high                   |
| `TicketPriorityUrgent` | urgent                 |