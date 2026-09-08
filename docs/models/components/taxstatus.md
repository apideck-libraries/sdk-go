# TaxStatus

The tax applicability of the product: `taxable` (the product is taxed), `shipping` (only the shipping is taxed, the product itself is exempt) or `none` (neither is taxed).

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TaxStatusTaxable

// Open enum: custom values can be created with a direct type cast
custom := components.TaxStatus("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `TaxStatusTaxable`  | taxable             |
| `TaxStatusShipping` | shipping            |
| `TaxStatusNone`     | none                |