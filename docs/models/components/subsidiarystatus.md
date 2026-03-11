# SubsidiaryStatus

Based on the status some functionality is enabled or disabled.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SubsidiaryStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.SubsidiaryStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SubsidiaryStatusActive`   | active                     |
| `SubsidiaryStatusInactive` | inactive                   |