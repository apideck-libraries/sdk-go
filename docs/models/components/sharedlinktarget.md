# SharedLinkTarget


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ID`                                                        | *string*                                                    | :heavy_check_mark:                                          | A unique identifier for an object.                          | 12345                                                       |
| `Name`                                                      | **string*                                                   | :heavy_minus_sign:                                          | The name of the file                                        | sample.jpg                                                  |
| `Type`                                                      | [*components.FileType](../../models/components/filetype.md) | :heavy_minus_sign:                                          | The type of resource. Could be file, folder or url          | file                                                        |