# BatchItemResultStatus

The outcome for this item.

- `created` / `updated` — the record was written.
- `failed` — it was not written, and `error` says why.
- `uncertain` — the request to the connector did not complete (typically a timeout), so whether this record was written is genuinely unknown. Do not blindly retry: list the resource to find what exists first, or a retry may duplicate it.

A `created` or `updated` result can also carry an `error`: the record was written, but work that runs after the write did not complete.

Read this rather than `id` to decide what happened. `id` is present when the identifier could also be read back, which is not the same question — a written record whose id could not be parsed still reports `created`.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.BatchItemResultStatusCreated

// Open enum: custom values can be created with a direct type cast
custom := components.BatchItemResultStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `BatchItemResultStatusCreated`   | created                          |
| `BatchItemResultStatusUpdated`   | updated                          |
| `BatchItemResultStatusFailed`    | failed                           |
| `BatchItemResultStatusUncertain` | uncertain                        |