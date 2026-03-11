# RequestType

The type of request

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.RequestTypeVacation

// Open enum: custom values can be created with a direct type cast
custom := components.RequestType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `RequestTypeVacation`    | vacation                 |
| `RequestTypeSick`        | sick                     |
| `RequestTypePersonal`    | personal                 |
| `RequestTypeJuryDuty`    | jury_duty                |
| `RequestTypeVolunteer`   | volunteer                |
| `RequestTypeBereavement` | bereavement              |
| `RequestTypeOther`       | other                    |