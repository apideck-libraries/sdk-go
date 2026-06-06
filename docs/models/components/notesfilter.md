# NotesFilter


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Title`                                    | `*string`                                  | :heavy_minus_sign:                         | Title of the note to filter on             | Follow up call                             |
| `OwnerID`                                  | `*string`                                  | :heavy_minus_sign:                         | Owner ID to filter on                      | 1234                                       |
| `UpdatedSince`                             | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        | 2020-09-30T07:43:32.000Z                   |
| `CreatedSince`                             | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        | 2020-09-30T07:43:32.000Z                   |