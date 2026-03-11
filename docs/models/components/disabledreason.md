# DisabledReason

Indicates why the webhook has been disabled. `retry_limit`: webhook reached its retry limit. `usage_limit`: account is over its usage limit. `delivery_url_validation_failed`: delivery URL failed validation during webhook creation or update.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.DisabledReasonNone

// Open enum: custom values can be created with a direct type cast
custom := components.DisabledReason("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `DisabledReasonNone`                        | none                                        |
| `DisabledReasonRetryLimit`                  | retry_limit                                 |
| `DisabledReasonUsageLimit`                  | usage_limit                                 |
| `DisabledReasonDeliveryURLValidationFailed` | delivery_url_validation_failed              |