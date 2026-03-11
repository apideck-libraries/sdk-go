# EcommerceOrderPaymentStatus

Current payment status of the order.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.EcommerceOrderPaymentStatusPending

// Open enum: custom values can be created with a direct type cast
custom := components.EcommerceOrderPaymentStatus("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `EcommerceOrderPaymentStatusPending`           | pending                                        |
| `EcommerceOrderPaymentStatusAuthorized`        | authorized                                     |
| `EcommerceOrderPaymentStatusPaid`              | paid                                           |
| `EcommerceOrderPaymentStatusPartial`           | partial                                        |
| `EcommerceOrderPaymentStatusRefunded`          | refunded                                       |
| `EcommerceOrderPaymentStatusVoided`            | voided                                         |
| `EcommerceOrderPaymentStatusUnknown`           | unknown                                        |
| `EcommerceOrderPaymentStatusPartiallyRefunded` | partially_refunded                             |