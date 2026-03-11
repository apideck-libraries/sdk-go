# RefundType

Type of refund. `refund_receipt` for itemized refunds with product/service lines and payment (QBO RefundReceipt, NetSuite CashRefund). `cash_refund` for cash-out refunds with GL distribution or allocations (Sage Intacct). `credit_note_refund` for refunds applied against a credit note (Zoho Books).

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.RefundTypeRefundReceipt

// Open enum: custom values can be created with a direct type cast
custom := components.RefundType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `RefundTypeRefundReceipt`    | refund_receipt               |
| `RefundTypeCashRefund`       | cash_refund                  |
| `RefundTypeCreditNoteRefund` | credit_note_refund           |