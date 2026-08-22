# AccountingPaymentMethodType

The type of payment method.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.AccountingPaymentMethodTypeCash

// Open enum: custom values can be created with a direct type cast
custom := components.AccountingPaymentMethodType("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AccountingPaymentMethodTypeCash`         | cash                                      |
| `AccountingPaymentMethodTypeCheck`        | check                                     |
| `AccountingPaymentMethodTypeCreditCard`   | credit_card                               |
| `AccountingPaymentMethodTypeDebitCard`    | debit_card                                |
| `AccountingPaymentMethodTypeBankTransfer` | bank_transfer                             |
| `AccountingPaymentMethodTypeElectronic`   | electronic                                |
| `AccountingPaymentMethodTypeOther`        | other                                     |
| `AccountingPaymentMethodTypeUnknown`      | unknown                                   |