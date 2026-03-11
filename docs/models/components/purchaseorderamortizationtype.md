# PurchaseOrderAmortizationType

Type of amortization

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PurchaseOrderAmortizationTypeManual

// Open enum: custom values can be created with a direct type cast
custom := components.PurchaseOrderAmortizationType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PurchaseOrderAmortizationTypeManual`   | manual                                  |
| `PurchaseOrderAmortizationTypeReceipt`  | receipt                                 |
| `PurchaseOrderAmortizationTypeSchedule` | schedule                                |
| `PurchaseOrderAmortizationTypeOther`    | other                                   |