# FlsaStatus

The FLSA status for this compensation.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FlsaStatusExempt

// Open enum: custom values can be created with a direct type cast
custom := components.FlsaStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `FlsaStatusExempt`            | exempt                        |
| `FlsaStatusSalariedNonexempt` | salaried-nonexempt            |
| `FlsaStatusNonexempt`         | nonexempt                     |
| `FlsaStatusOwner`             | owner                         |
| `FlsaStatusOther`             | other                         |