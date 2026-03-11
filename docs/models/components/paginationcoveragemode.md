# PaginationCoverageMode

How pagination is implemented on this connector. Native mode means Apideck is using the pagination parameters of the connector. With virtual pagination, the connector does not support pagination, but Apideck emulates it.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.PaginationCoverageModeNative

// Open enum: custom values can be created with a direct type cast
custom := components.PaginationCoverageMode("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `PaginationCoverageModeNative`  | native                          |
| `PaginationCoverageModeVirtual` | virtual                         |