# TaxType

The tax applicability of this line item. Overrides the root-level tax_type for this line.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TaxTypeSales

// Open enum: custom values can be created with a direct type cast
custom := components.TaxType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `TaxTypeSales`    | sales             |
| `TaxTypePurchase` | purchase          |