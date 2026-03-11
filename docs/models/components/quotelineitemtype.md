# QuoteLineItemType

Item type

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.QuoteLineItemTypeSalesItem

// Open enum: custom values can be created with a direct type cast
custom := components.QuoteLineItemType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `QuoteLineItemTypeSalesItem` | sales_item                   |
| `QuoteLineItemTypeDiscount`  | discount                     |
| `QuoteLineItemTypeInfo`      | info                         |
| `QuoteLineItemTypeSubTotal`  | sub_total                    |
| `QuoteLineItemTypeService`   | service                      |
| `QuoteLineItemTypeOther`     | other                        |