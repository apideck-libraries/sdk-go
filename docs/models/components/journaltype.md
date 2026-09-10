# JournalType

Normalized journal classification.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JournalTypeGeneral

// Open enum: custom values can be created with a direct type cast
custom := components.JournalType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `JournalTypeGeneral`            | general                         |
| `JournalTypeSales`              | sales                           |
| `JournalTypePurchase`           | purchase                        |
| `JournalTypeSalesCreditNote`    | sales_credit_note               |
| `JournalTypePurchaseCreditNote` | purchase_credit_note            |
| `JournalTypeCash`               | cash                            |
| `JournalTypeBank`               | bank                            |
| `JournalTypePaymentService`     | payment_service                 |
| `JournalTypeOther`              | other                           |