# FileType

The type of resource. Could be file, folder or url

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.FileTypeFile

// Open enum: custom values can be created with a direct type cast
custom := components.FileType("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `FileTypeFile`   | file             |
| `FileTypeFolder` | folder           |
| `FileTypeURL`    | url              |