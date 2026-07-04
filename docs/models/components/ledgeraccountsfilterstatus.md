# LedgerAccountsFilterStatus

Filter by account status. Supported only on a subset of connectors (e.g. NetSuite); connectors that do not support it reject `filter[status]` with a `400 UnsupportedFiltersError` — read the account's `status` field in the response and filter client-side instead. See the error's `supported_filters` or the connector's supported filters.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LedgerAccountsFilterStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.LedgerAccountsFilterStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `LedgerAccountsFilterStatusActive`   | active                               |
| `LedgerAccountsFilterStatusInactive` | inactive                             |