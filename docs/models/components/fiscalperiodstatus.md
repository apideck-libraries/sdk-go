# FiscalPeriodStatus

Whether transactions can be posted into the period. `open`: postings are accepted on at least the last day of the period; a lock date inside the period still rejects earlier dates (see `locked_through`). `locked`: posting is blocked for some users or transaction types only, for example everyone except advisers or administrators. `closed`: posting is blocked for all users until the period or year is reopened. `future`: the accounting system reports the period as not yet opened. Null or absent means the connector does not expose the period's status.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FiscalPeriodStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.FiscalPeriodStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `FiscalPeriodStatusOpen`   | open                       |
| `FiscalPeriodStatusLocked` | locked                     |
| `FiscalPeriodStatusClosed` | closed                     |
| `FiscalPeriodStatusFuture` | future                     |