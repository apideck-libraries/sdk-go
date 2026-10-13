# LineType

`account`: a ledger account. `unappropriated_result`: the result of earlier fiscal years not yet posted to an equity account, added by the connector. `unmapped`: a balance the accounting system reports without a ledger account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.LineTypeAccount

// Open enum: custom values can be created with a direct type cast
custom := components.LineType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `LineTypeAccount`              | account                        |
| `LineTypeUnappropriatedResult` | unappropriated_result          |
| `LineTypeUnmapped`             | unmapped                       |