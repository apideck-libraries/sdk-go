# Gender

The gender represents the gender identity of a person.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.GenderMale

// Open enum: custom values can be created with a direct type cast
custom := components.Gender("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `GenderMale`         | male                 |
| `GenderFemale`       | female               |
| `GenderUnisex`       | unisex               |
| `GenderOther`        | other                |
| `GenderNotSpecified` | not_specified        |