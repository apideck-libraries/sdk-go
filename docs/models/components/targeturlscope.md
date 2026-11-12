# TargetURLScope

Where the vendor stores the webhook target URL. `app` means one URL per vendor developer app, which customers on Apideck's own credentials cannot configure. `tenant` means each customer's vendor account holds its own URL.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TargetURLScopeApp

// Open enum: custom values can be created with a direct type cast
custom := components.TargetURLScope("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `TargetURLScopeApp`    | app                    |
| `TargetURLScopeTenant` | tenant                 |