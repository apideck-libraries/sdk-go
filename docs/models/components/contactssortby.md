# ContactsSortBy

The field on which to sort the Contacts

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ContactsSortByCreatedAt

// Open enum: custom values can be created with a direct type cast
custom := components.ContactsSortBy("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ContactsSortByCreatedAt` | created_at                |
| `ContactsSortByUpdatedAt` | updated_at                |
| `ContactsSortByName`      | name                      |
| `ContactsSortByFirstName` | first_name                |
| `ContactsSortByLastName`  | last_name                 |
| `ContactsSortByEmail`     | email                     |