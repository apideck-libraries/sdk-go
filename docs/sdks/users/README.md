# Users
(*Crm.Users*)

## Overview

### Available Operations

* [List](#list) - List users
* [Create](#create) - Create user
* [Get](#get) - Get user
* [Update](#update) - Update user
* [Delete](#delete) - Delete user

## List

List users

### Example Usage

<!-- UsageSnippet language="go" operationID="crm.usersAll" method="get" path="/crm/users" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Crm.Users.List(ctx, operations.CrmUsersAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetUsersResponse != nil {
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CrmUsersAllRequest](../../models/operations/crmusersallrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmUsersAllResponse](../../models/operations/crmusersallresponse.md), error**

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

Create user

### Example Usage

<!-- UsageSnippet language="go" operationID="crm.usersAdd" method="post" path="/crm/users" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Crm.Users.Create(ctx, operations.CrmUsersAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        User: components.UserInput{
            ParentID: sdkgo.String("54321"),
            Username: sdkgo.String("masterofcoin"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            Title: sdkgo.String("CEO"),
            Division: sdkgo.String("Europe"),
            CompanyName: sdkgo.String("SpaceX"),
            EmployeeNumber: sdkgo.String("123456-AB"),
            Description: sdkgo.String("A description"),
            Image: sdkgo.String("https://logo.clearbit.com/spacex.com?s=128"),
            Language: sdkgo.String("EN"),
            Status: sdkgo.String("active"),
            Password: sdkgo.String("supersecretpassword"),
            Addresses: []components.Address{
                components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            PhoneNumbers: []components.PhoneNumber{
                components.PhoneNumber{
                    ID: sdkgo.String("12345"),
                    CountryCode: sdkgo.String("1"),
                    AreaCode: sdkgo.String("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.String("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
                components.PhoneNumber{
                    ID: sdkgo.String("12345"),
                    CountryCode: sdkgo.String("1"),
                    AreaCode: sdkgo.String("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.String("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.String("123"),
                    Email: sdkgo.String("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
                components.Email{
                    ID: sdkgo.String("123"),
                    Email: sdkgo.String("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
            },
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CrmUsersAddRequest](../../models/operations/crmusersaddrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmUsersAddResponse](../../models/operations/crmusersaddresponse.md), error**

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

Get user

### Example Usage

<!-- UsageSnippet language="go" operationID="crm.usersOne" method="get" path="/crm/users/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Crm.Users.Get(ctx, operations.CrmUsersOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CrmUsersOneRequest](../../models/operations/crmusersonerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmUsersOneResponse](../../models/operations/crmusersoneresponse.md), error**

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

Update user

### Example Usage

<!-- UsageSnippet language="go" operationID="crm.usersUpdate" method="patch" path="/crm/users/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Crm.Users.Update(ctx, operations.CrmUsersUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        User: components.UserInput{
            ParentID: sdkgo.String("54321"),
            Username: sdkgo.String("masterofcoin"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            Title: sdkgo.String("CEO"),
            Division: sdkgo.String("Europe"),
            CompanyName: sdkgo.String("SpaceX"),
            EmployeeNumber: sdkgo.String("123456-AB"),
            Description: sdkgo.String("A description"),
            Image: sdkgo.String("https://logo.clearbit.com/spacex.com?s=128"),
            Language: sdkgo.String("EN"),
            Status: sdkgo.String("active"),
            Password: sdkgo.String("supersecretpassword"),
            Addresses: []components.Address{
                components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
                components.Address{
                    ID: sdkgo.String("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.String("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.String("HQ US"),
                    Line1: sdkgo.String("Main street"),
                    Line2: sdkgo.String("apt #"),
                    Line3: sdkgo.String("Suite #"),
                    Line4: sdkgo.String("delivery instructions"),
                    StreetNumber: sdkgo.String("25"),
                    City: sdkgo.String("San Francisco"),
                    State: sdkgo.String("CA"),
                    PostalCode: sdkgo.String("94104"),
                    Country: sdkgo.String("US"),
                    Latitude: sdkgo.String("40.759211"),
                    Longitude: sdkgo.String("-73.984638"),
                    County: sdkgo.String("Santa Clara"),
                    ContactName: sdkgo.String("Elon Musk"),
                    Salutation: sdkgo.String("Mr"),
                    PhoneNumber: sdkgo.String("111-111-1111"),
                    Fax: sdkgo.String("122-111-1111"),
                    Email: sdkgo.String("elon@musk.com"),
                    Website: sdkgo.String("https://elonmusk.com"),
                    Notes: sdkgo.String("Address notes or delivery instructions."),
                    RowVersion: sdkgo.String("1-12345"),
                },
            },
            PhoneNumbers: []components.PhoneNumber{
                components.PhoneNumber{
                    ID: sdkgo.String("12345"),
                    CountryCode: sdkgo.String("1"),
                    AreaCode: sdkgo.String("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.String("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
                components.PhoneNumber{
                    ID: sdkgo.String("12345"),
                    CountryCode: sdkgo.String("1"),
                    AreaCode: sdkgo.String("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.String("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.String("123"),
                    Email: sdkgo.String("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
                components.Email{
                    ID: sdkgo.String("123"),
                    Email: sdkgo.String("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
            },
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
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CrmUsersUpdateRequest](../../models/operations/crmusersupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CrmUsersUpdateResponse](../../models/operations/crmusersupdateresponse.md), error**

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

Delete user

### Example Usage

<!-- UsageSnippet language="go" operationID="crm.usersDelete" method="delete" path="/crm/users/{id}" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkgo.New(
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Crm.Users.Delete(ctx, operations.CrmUsersDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteUserResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CrmUsersDeleteRequest](../../models/operations/crmusersdeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CrmUsersDeleteResponse](../../models/operations/crmusersdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |