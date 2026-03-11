# PaymentType

Type of payment

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaymentTypeAccountsReceivable

// Open enum: custom values can be created with a direct type cast
custom := components.PaymentType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PaymentTypeAccountsReceivable`            | accounts_receivable                        |
| `PaymentTypeAccountsPayable`               | accounts_payable                           |
| `PaymentTypeAccountsReceivableCredit`      | accounts_receivable_credit                 |
| `PaymentTypeAccountsPayableCredit`         | accounts_payable_credit                    |
| `PaymentTypeAccountsReceivableOverpayment` | accounts_receivable_overpayment            |
| `PaymentTypeAccountsPayableOverpayment`    | accounts_payable_overpayment               |
| `PaymentTypeAccountsReceivablePrepayment`  | accounts_receivable_prepayment             |
| `PaymentTypeAccountsPayablePrepayment`     | accounts_payable_prepayment                |