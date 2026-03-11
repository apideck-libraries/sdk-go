# SupplierStatus

Supplier status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SupplierStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.SupplierStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `SupplierStatusActive`             | active                             |
| `SupplierStatusInactive`           | inactive                           |
| `SupplierStatusArchived`           | archived                           |
| `SupplierStatusGdprErasureRequest` | gdpr-erasure-request               |
| `SupplierStatusUnknown`            | unknown                            |