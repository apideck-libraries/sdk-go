# BillPaymentCreateInputBillPaymentType

Type of payment

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentCreateInputBillPaymentTypeAccountsPayableCredit

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentCreateInputBillPaymentType("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BillPaymentCreateInputBillPaymentTypeAccountsPayableCredit`      | accounts_payable_credit                                           |
| `BillPaymentCreateInputBillPaymentTypeAccountsPayableOverpayment` | accounts_payable_overpayment                                      |
| `BillPaymentCreateInputBillPaymentTypeAccountsPayablePrepayment`  | accounts_payable_prepayment                                       |
| `BillPaymentCreateInputBillPaymentTypeAccountsPayable`            | accounts_payable                                                  |