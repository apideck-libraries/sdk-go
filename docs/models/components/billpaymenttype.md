# BillPaymentType

Type of payment

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentTypeAccountsPayableCredit

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `BillPaymentTypeAccountsPayableCredit`      | accounts_payable_credit                     |
| `BillPaymentTypeAccountsPayableOverpayment` | accounts_payable_overpayment                |
| `BillPaymentTypeAccountsPayablePrepayment`  | accounts_payable_prepayment                 |
| `BillPaymentTypeAccountsPayable`            | accounts_payable                            |