# ApplicationStatus

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ApplicationStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.ApplicationStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ApplicationStatusOpen`      | open                         |
| `ApplicationStatusRejected`  | rejected                     |
| `ApplicationStatusHired`     | hired                        |
| `ApplicationStatusConverted` | converted                    |
| `ApplicationStatusOther`     | other                        |