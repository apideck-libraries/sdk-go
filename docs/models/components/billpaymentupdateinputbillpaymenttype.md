# BillPaymentUpdateInputBillPaymentType

Type of payment

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BillPaymentUpdateInputBillPaymentTypeAccountsPayableCredit

// Open enum: custom values can be created with a direct type cast
custom := components.BillPaymentUpdateInputBillPaymentType("custom_value")
```


## Values

| Name                                                              | Value                                                             |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `BillPaymentUpdateInputBillPaymentTypeAccountsPayableCredit`      | accounts_payable_credit                                           |
| `BillPaymentUpdateInputBillPaymentTypeAccountsPayableOverpayment` | accounts_payable_overpayment                                      |
| `BillPaymentUpdateInputBillPaymentTypeAccountsPayablePrepayment`  | accounts_payable_prepayment                                       |
| `BillPaymentUpdateInputBillPaymentTypeAccountsPayable`            | accounts_payable                                                  |