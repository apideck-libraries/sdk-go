# SupplierCreateInputStatus

Supplier status

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.SupplierCreateInputStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.SupplierCreateInputStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `SupplierCreateInputStatusActive`             | active                                        |
| `SupplierCreateInputStatusInactive`           | inactive                                      |
| `SupplierCreateInputStatusArchived`           | archived                                      |
| `SupplierCreateInputStatusGdprErasureRequest` | gdpr-erasure-request                          |
| `SupplierCreateInputStatusUnknown`            | unknown                                       |