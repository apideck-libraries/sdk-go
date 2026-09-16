# GoodsReceiptStatus

Lifecycle status of the receipt. `draft` covers unposted or awaiting-validation documents, `pending_approval` covers documents submitted into an approval flow, `received` covers posted/validated/released receipts that have affected stock or the receiving ledger, `cancelled` covers voided, reversed or denied receipts.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.GoodsReceiptStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.GoodsReceiptStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `GoodsReceiptStatusDraft`           | draft                               |
| `GoodsReceiptStatusPendingApproval` | pending_approval                    |
| `GoodsReceiptStatusReceived`        | received                            |
| `GoodsReceiptStatusCancelled`       | cancelled                           |
| `GoodsReceiptStatusOther`           | other                               |