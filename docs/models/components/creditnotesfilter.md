# CreditNotesFilter


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           | Example                                                               |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `IDSince`                                                             | `*string`                                                             | :heavy_minus_sign:                                                    | Return records with a row ID greater than or equal to the given value | 1                                                                     |
| `UpdatedSince`                                                        | [*time.Time](https://pkg.go.dev/time#Time)                            | :heavy_minus_sign:                                                    | N/A                                                                   | 2020-09-30T07:43:32.000Z                                              |