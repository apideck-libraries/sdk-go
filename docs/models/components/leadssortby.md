# LeadsSortBy

The field on which to sort the Leads

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LeadsSortByCreatedAt

// Open enum: custom values can be created with a direct type cast
custom := components.LeadsSortBy("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `LeadsSortByCreatedAt` | created_at             |
| `LeadsSortByUpdatedAt` | updated_at             |
| `LeadsSortByName`      | name                   |
| `LeadsSortByFirstName` | first_name             |
| `LeadsSortByLastName`  | last_name              |
| `LeadsSortByEmail`     | email                  |