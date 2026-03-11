# ConsentState

The current consent state of the connection

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.ConsentStateImplicit

// Open enum: custom values can be created with a direct type cast
custom := components.ConsentState("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ConsentStateImplicit`          | implicit                        |
| `ConsentStatePending`           | pending                         |
| `ConsentStateGranted`           | granted                         |
| `ConsentStateDenied`            | denied                          |
| `ConsentStateRevoked`           | revoked                         |
| `ConsentStateRequiresReconsent` | requires_reconsent              |