# Type

The type of address.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.TypePrimary

// Open enum: custom values can be created with a direct type cast
custom := components.Type("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `TypePrimary`   | primary         |
| `TypeSecondary` | secondary       |
| `TypeHome`      | home            |
| `TypeOffice`    | office          |
| `TypeShipping`  | shipping        |
| `TypeBilling`   | billing         |
| `TypeWork`      | work            |
| `TypeOther`     | other           |