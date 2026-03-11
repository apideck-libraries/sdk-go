# AccountType

The type of bank account.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AccountTypeBankAccount

// Open enum: custom values can be created with a direct type cast
custom := components.AccountType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AccountTypeBankAccount` | bank_account             |
| `AccountTypeCreditCard`  | credit_card              |
| `AccountTypeOther`       | other                    |