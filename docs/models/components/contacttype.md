# ContactType

The type of the contact.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ContactTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := components.ContactType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ContactTypeCustomer` | customer              |
| `ContactTypeSupplier` | supplier              |
| `ContactTypeEmployee` | employee              |
| `ContactTypePersonal` | personal              |