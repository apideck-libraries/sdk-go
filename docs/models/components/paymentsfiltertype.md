# PaymentsFilterType

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentsFilterTypeAccountsReceivable

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentsFilterType("custom_value")
```


## Values

| Name                                              | Value                                             |
| ------------------------------------------------- | ------------------------------------------------- |
| `PaymentsFilterTypeAccountsReceivable`            | accounts_receivable                               |
| `PaymentsFilterTypeAccountsPayable`               | accounts_payable                                  |
| `PaymentsFilterTypeAccountsReceivableCredit`      | accounts_receivable_credit                        |
| `PaymentsFilterTypeAccountsPayableCredit`         | accounts_payable_credit                           |
| `PaymentsFilterTypeAccountsReceivableOverpayment` | accounts_receivable_overpayment                   |
| `PaymentsFilterTypeAccountsPayableOverpayment`    | accounts_payable_overpayment                      |
| `PaymentsFilterTypeAccountsReceivablePrepayment`  | accounts_receivable_prepayment                    |
| `PaymentsFilterTypeAccountsPayablePrepayment`     | accounts_payable_prepayment                       |