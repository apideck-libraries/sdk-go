# Drives
(*FileStorage.Drives*)

## Overview

### Available Operations

* [List](#list) - List Drives
* [Create](#create) - Create Drive
* [Get](#get) - Get Drive
* [Update](#update) - Update Drive
* [Delete](#delete) - Delete Drive

## List

List Drives

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.FileStorage.Drives.List(ctx, operations.FileStorageDrivesAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.DrivesFilter{
            GroupID: sdkgo.String("1234"),
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDrivesResponse != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.FileStorageDrivesAllRequest](../../models/operations/filestoragedrivesallrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |
| `opts`                                                                                           | [][operations.Option](../../models/operations/option.md)                                         | :heavy_minus_sign:                                                                               | The options for this request.                                                                    |

### Response

**[*operations.FileStorageDrivesAllResponse](../../models/operations/filestoragedrivesallresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Create

Create Drive

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.FileStorage.Drives.Create(ctx, components.DriveInput{
        Name: "Project Resources",
        Description: sdkgo.String("A description"),
        PassThrough: []components.PassThroughBody{
            components.PassThroughBody{
                ServiceID: "<id>",
                ExtendPaths: []components.ExtendPaths{
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                },
            },
        },
    }, nil, sdkgo.String("salesforce"))
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateDriveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `drive`                                                                                                                                       | [components.DriveInput](../../models/components/driveinput.md)                                                                                | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.FileStorageDrivesAddResponse](../../models/operations/filestoragedrivesaddresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Get

Get Drive

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.FileStorage.Drives.Get(ctx, "<id>", sdkgo.String("salesforce"), nil, sdkgo.String("id,updated_at"))
    if err != nil {
        log.Fatal(err)
    }
    if res.GetDriveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `id`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | ID of the record you are acting upon.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `serviceID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | salesforce                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `raw`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | **bool*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Include raw response. Mostly used for debugging purposes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `fields`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The 'fields' parameter allows API users to specify the fields they want to include in the API response. If this parameter is not present, the API will return all available fields. If this parameter is present, only the fields specified in the comma-separated string will be included in the response. Nested properties can also be requested by using a dot notation. <br /><br />Example: `fields=name,email,addresses.city`<br /><br />In the example above, the response will only include the fields "name", "email" and "addresses.city". If any other fields are available, they will be excluded. | id,updated_at                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

### Response

**[*operations.FileStorageDrivesOneResponse](../../models/operations/filestoragedrivesoneresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Update

Update Drive

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.FileStorage.Drives.Update(ctx, "<id>", components.DriveInput{
        Name: "Project Resources",
        Description: sdkgo.String("A description"),
        PassThrough: []components.PassThroughBody{
            components.PassThroughBody{
                ServiceID: "<id>",
                ExtendPaths: []components.ExtendPaths{
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                },
            },
            components.PassThroughBody{
                ServiceID: "<id>",
                ExtendPaths: []components.ExtendPaths{
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                },
            },
            components.PassThroughBody{
                ServiceID: "<id>",
                ExtendPaths: []components.ExtendPaths{
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                    components.ExtendPaths{
                        Path: "$.nested.property",
                        Value: map[string]any{
                            "TaxClassificationRef": map[string]any{
                                "value": "EUC-99990201-V1-00020000",
                            },
                        },
                    },
                },
            },
        },
    }, sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateDriveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `id`                                                                                                                                          | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | ID of the record you are acting upon.                                                                                                         |                                                                                                                                               |
| `drive`                                                                                                                                       | [components.DriveInput](../../models/components/driveinput.md)                                                                                | :heavy_check_mark:                                                                                                                            | N/A                                                                                                                                           |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.FileStorageDrivesUpdateResponse](../../models/operations/filestoragedrivesupdateresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |

## Delete

Delete Drive

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkgo.New(
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.FileStorage.Drives.Delete(ctx, "<id>", sdkgo.String("salesforce"), nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteDriveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                     | Type                                                                                                                                          | Required                                                                                                                                      | Description                                                                                                                                   | Example                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                         | :heavy_check_mark:                                                                                                                            | The context to use for the request.                                                                                                           |                                                                                                                                               |
| `id`                                                                                                                                          | *string*                                                                                                                                      | :heavy_check_mark:                                                                                                                            | ID of the record you are acting upon.                                                                                                         |                                                                                                                                               |
| `serviceID`                                                                                                                                   | **string*                                                                                                                                     | :heavy_minus_sign:                                                                                                                            | Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API. | salesforce                                                                                                                                    |
| `raw`                                                                                                                                         | **bool*                                                                                                                                       | :heavy_minus_sign:                                                                                                                            | Include raw response. Mostly used for debugging purposes                                                                                      |                                                                                                                                               |
| `opts`                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                      | :heavy_minus_sign:                                                                                                                            | The options for this request.                                                                                                                 |                                                                                                                                               |

### Response

**[*operations.FileStorageDrivesDeleteResponse](../../models/operations/filestoragedrivesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |