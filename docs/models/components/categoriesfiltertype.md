# CategoriesFilterType

The type of the category.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CategoriesFilterTypeSupplier

// Open enum: custom values can be created with a direct type cast
custom := components.CategoriesFilterType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CategoriesFilterTypeSupplier` | supplier                       |
| `CategoriesFilterTypeExpense`  | expense                        |
| `CategoriesFilterTypeRevenue`  | revenue                        |
| `CategoriesFilterTypeCustomer` | customer                       |