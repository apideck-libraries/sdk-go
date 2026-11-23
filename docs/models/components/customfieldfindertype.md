# CustomFieldFinderType

The declared data type of the custom field. Only returned by connectors that can read the field's definition; omitted otherwise.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.CustomFieldFinderTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.CustomFieldFinderType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `CustomFieldFinderTypeText`           | text                                  |
| `CustomFieldFinderTypeCheckbox`       | checkbox                              |
| `CustomFieldFinderTypeTel`            | tel                                   |
| `CustomFieldFinderTypeEmail`          | email                                 |
| `CustomFieldFinderTypeURL`            | url                                   |
| `CustomFieldFinderTypeTextarea`       | textarea                              |
| `CustomFieldFinderTypeSelect`         | select                                |
| `CustomFieldFinderTypeFilteredSelect` | filtered-select                       |
| `CustomFieldFinderTypeMultiSelect`    | multi-select                          |
| `CustomFieldFinderTypeDatetime`       | datetime                              |
| `CustomFieldFinderTypeDate`           | date                                  |
| `CustomFieldFinderTypeTime`           | time                                  |
| `CustomFieldFinderTypeNumber`         | number                                |
| `CustomFieldFinderTypePassword`       | password                              |