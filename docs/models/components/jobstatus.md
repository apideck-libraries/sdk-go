# JobStatus

The status of the job.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.JobStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := components.JobStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `JobStatusDraft`               | draft                          |
| `JobStatusInternal`            | internal                       |
| `JobStatusPublished`           | published                      |
| `JobStatusCompleted`           | completed                      |
| `JobStatusPlaced`              | placed                         |
| `JobStatusOnHold`              | on-hold                        |
| `JobStatusPrivate`             | private                        |
| `JobStatusAcceptingCandidates` | accepting_candidates           |
| `JobStatusOpen`                | open                           |
| `JobStatusClosed`              | closed                         |
| `JobStatusArchived`            | archived                       |