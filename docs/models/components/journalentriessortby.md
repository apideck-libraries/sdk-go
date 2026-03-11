# JournalEntriesSortBy

The field on which to sort the Journal Entries.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalEntriesSortByCreatedAt

// Open enum: custom values can be created with a direct type cast
custom := components.JournalEntriesSortBy("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `JournalEntriesSortByCreatedAt` | created_at                      |
| `JournalEntriesSortByUpdatedAt` | updated_at                      |