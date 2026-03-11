# Units

The unit of time off requested. Possible values include: `hours`, `days`, or `other`.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.UnitsDays

// Open enum: custom values can be created with a direct type cast
custom := components.Units("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `UnitsDays`  | days         |
| `UnitsHours` | hours        |
| `UnitsOther` | other        |