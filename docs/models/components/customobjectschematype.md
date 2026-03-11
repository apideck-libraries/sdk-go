# CustomObjectSchemaType

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CustomObjectSchemaTypeString

// Open enum: custom values can be created with a direct type cast
custom := components.CustomObjectSchemaType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `CustomObjectSchemaTypeString`      | string                              |
| `CustomObjectSchemaTypeNumber`      | number                              |
| `CustomObjectSchemaTypeInteger`     | integer                             |
| `CustomObjectSchemaTypeBoolean`     | boolean                             |
| `CustomObjectSchemaTypeDate`        | date                                |
| `CustomObjectSchemaTypeDatetime`    | datetime                            |
| `CustomObjectSchemaTypeCurrency`    | currency                            |
| `CustomObjectSchemaTypeEmail`       | email                               |
| `CustomObjectSchemaTypePhone`       | phone                               |
| `CustomObjectSchemaTypeReference`   | reference                           |
| `CustomObjectSchemaTypeSelect`      | select                              |
| `CustomObjectSchemaTypeMultiselect` | multiselect                         |