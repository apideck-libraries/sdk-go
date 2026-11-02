# DepreciationMethod

The depreciation method of the assets, normalised across accounting systems: declining-balance variants (for example 150% or 200%) are all `declining_balance`, and the accounting system's own value is in `downstream_depreciation_method`. Set only when every asset recorded in the account uses the same method.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.DepreciationMethodOther

// Open enum: custom values can be created with a direct type cast
custom := components.DepreciationMethod("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `DepreciationMethodOther`            | other                                |
| `DepreciationMethodStraightLine`     | straight_line                        |
| `DepreciationMethodDecliningBalance` | declining_balance                    |
| `DepreciationMethodNone`             | none                                 |