# Activities
(*Crm.Activities*)

## Overview

### Available Operations

* [List](#list) - List activities
* [Create](#create) - Create activity
* [Get](#get) - Get activity
* [Update](#update) - Update activity
* [Delete](#delete) - Delete activity

## List

List activities

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Crm.Activities.List(ctx, operations.CrmActivitiesAllRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        Limit: sdkgo.Int64(20),
        Filter: &components.ActivitiesFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.ActivitiesSort{
            By: components.ActivitiesSortByCreatedAt.ToPointer(),
            Direction: components.SortDirectionDesc.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetActivitiesResponse != nil {
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CrmActivitiesAllRequest](../../models/operations/crmactivitiesallrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CrmActivitiesAllResponse](../../models/operations/crmactivitiesallresponse.md), error**

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

Create activity

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

    res, err := s.Crm.Activities.Create(ctx, operations.CrmActivitiesAddRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        Activity: components.ActivityInput{
            ActivityDatetime: sdkgo.String("2021-05-01T12:00:00.000Z"),
            DurationSeconds: sdkgo.Int64(1800),
            UserID: sdkgo.String("12345"),
            AccountID: sdkgo.String("12345"),
            ContactID: sdkgo.String("12345"),
            CompanyID: sdkgo.String("12345"),
            OpportunityID: sdkgo.String("12345"),
            LeadID: sdkgo.String("12345"),
            OwnerID: sdkgo.String("12345"),
            CampaignID: sdkgo.String("12345"),
            CaseID: sdkgo.String("12345"),
            AssetID: sdkgo.String("12345"),
            ContractID: sdkgo.String("12345"),
            ProductID: sdkgo.String("12345"),
            SolutionID: sdkgo.String("12345"),
            CustomObjectID: sdkgo.String("12345"),
            Type: components.ActivityTypeMeeting.ToPointer(),
            Title: sdkgo.String("Meeting"),
            Description: sdkgo.String("More info about the meeting"),
            Note: sdkgo.String("An internal note about the meeting"),
            Location: sdkgo.String("Space"),
            LocationAddress: &components.Address{
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
            AllDayEvent: sdkgo.Bool(false),
            Private: sdkgo.Bool(true),
            GroupEvent: sdkgo.Bool(true),
            EventSubType: sdkgo.String("debrief"),
            GroupEventType: sdkgo.String("Proposed"),
            Child: sdkgo.Bool(false),
            Archived: sdkgo.Bool(false),
            Deleted: sdkgo.Bool(false),
            ShowAs: components.ShowAsBusy.ToPointer(),
            Done: sdkgo.Bool(false),
            StartDatetime: sdkgo.String("2021-05-01T12:00:00.000Z"),
            EndDatetime: sdkgo.String("2021-05-01T12:30:00.000Z"),
            ActivityDate: sdkgo.String("2021-05-01"),
            EndDate: sdkgo.String("2021-05-01"),
            Recurrent: sdkgo.Bool(false),
            ReminderDatetime: sdkgo.String("2021-05-01T17:00:00.000Z"),
            ReminderSet: sdkgo.Bool(false),
            VideoConferenceURL: sdkgo.String("https://us02web.zoom.us/j/88120759396"),
            VideoConferenceID: sdkgo.String("zoom:88120759396"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueFour(
                        components.Four{},
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueBoolean(
                        true,
                    )),
                },
            },
            Attendees: []components.ActivityAttendeeInput{
                components.ActivityAttendeeInput{
                    Name: sdkgo.String("Elon Musk"),
                    FirstName: sdkgo.String("Elon"),
                    MiddleName: sdkgo.String("D."),
                    LastName: sdkgo.String("Musk"),
                    Prefix: sdkgo.String("Mr."),
                    Suffix: sdkgo.String("PhD"),
                    EmailAddress: sdkgo.String("elon@musk.com"),
                    IsOrganizer: sdkgo.Bool(true),
                    Status: components.ActivityAttendeeStatusAccepted.ToPointer(),
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
    if res.CreateActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CrmActivitiesAddRequest](../../models/operations/crmactivitiesaddrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CrmActivitiesAddResponse](../../models/operations/crmactivitiesaddresponse.md), error**

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

Get activity

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Crm.Activities.Get(ctx, operations.CrmActivitiesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CrmActivitiesOneRequest](../../models/operations/crmactivitiesonerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.CrmActivitiesOneResponse](../../models/operations/crmactivitiesoneresponse.md), error**

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

Update activity

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

    res, err := s.Crm.Activities.Update(ctx, operations.CrmActivitiesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        Activity: components.ActivityInput{
            ActivityDatetime: sdkgo.String("2021-05-01T12:00:00.000Z"),
            DurationSeconds: sdkgo.Int64(1800),
            UserID: sdkgo.String("12345"),
            AccountID: sdkgo.String("12345"),
            ContactID: sdkgo.String("12345"),
            CompanyID: sdkgo.String("12345"),
            OpportunityID: sdkgo.String("12345"),
            LeadID: sdkgo.String("12345"),
            OwnerID: sdkgo.String("12345"),
            CampaignID: sdkgo.String("12345"),
            CaseID: sdkgo.String("12345"),
            AssetID: sdkgo.String("12345"),
            ContractID: sdkgo.String("12345"),
            ProductID: sdkgo.String("12345"),
            SolutionID: sdkgo.String("12345"),
            CustomObjectID: sdkgo.String("12345"),
            Type: components.ActivityTypeMeeting.ToPointer(),
            Title: sdkgo.String("Meeting"),
            Description: sdkgo.String("More info about the meeting"),
            Note: sdkgo.String("An internal note about the meeting"),
            Location: sdkgo.String("Space"),
            LocationAddress: &components.Address{
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
            AllDayEvent: sdkgo.Bool(false),
            Private: sdkgo.Bool(true),
            GroupEvent: sdkgo.Bool(true),
            EventSubType: sdkgo.String("debrief"),
            GroupEventType: sdkgo.String("Proposed"),
            Child: sdkgo.Bool(false),
            Archived: sdkgo.Bool(false),
            Deleted: sdkgo.Bool(false),
            ShowAs: components.ShowAsBusy.ToPointer(),
            Done: sdkgo.Bool(false),
            StartDatetime: sdkgo.String("2021-05-01T12:00:00.000Z"),
            EndDatetime: sdkgo.String("2021-05-01T12:30:00.000Z"),
            ActivityDate: sdkgo.String("2021-05-01"),
            EndDate: sdkgo.String("2021-05-01"),
            Recurrent: sdkgo.Bool(false),
            ReminderDatetime: sdkgo.String("2021-05-01T17:00:00.000Z"),
            ReminderSet: sdkgo.Bool(false),
            VideoConferenceURL: sdkgo.String("https://us02web.zoom.us/j/88120759396"),
            VideoConferenceID: sdkgo.String("zoom:88120759396"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueFour(
                        components.Four{},
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueBoolean(
                        true,
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueArrayOf6(
                        []components.Six{
                            components.Six{},
                            components.Six{},
                        },
                    )),
                },
            },
            Attendees: []components.ActivityAttendeeInput{
                components.ActivityAttendeeInput{
                    Name: sdkgo.String("Elon Musk"),
                    FirstName: sdkgo.String("Elon"),
                    MiddleName: sdkgo.String("D."),
                    LastName: sdkgo.String("Musk"),
                    Prefix: sdkgo.String("Mr."),
                    Suffix: sdkgo.String("PhD"),
                    EmailAddress: sdkgo.String("elon@musk.com"),
                    IsOrganizer: sdkgo.Bool(true),
                    Status: components.ActivityAttendeeStatusAccepted.ToPointer(),
                },
                components.ActivityAttendeeInput{
                    Name: sdkgo.String("Elon Musk"),
                    FirstName: sdkgo.String("Elon"),
                    MiddleName: sdkgo.String("D."),
                    LastName: sdkgo.String("Musk"),
                    Prefix: sdkgo.String("Mr."),
                    Suffix: sdkgo.String("PhD"),
                    EmailAddress: sdkgo.String("elon@musk.com"),
                    IsOrganizer: sdkgo.Bool(true),
                    Status: components.ActivityAttendeeStatusAccepted.ToPointer(),
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
    if res.UpdateActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CrmActivitiesUpdateRequest](../../models/operations/crmactivitiesupdaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CrmActivitiesUpdateResponse](../../models/operations/crmactivitiesupdateresponse.md), error**

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

Delete activity

### Example Usage

```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Crm.Activities.Delete(ctx, operations.CrmActivitiesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteActivityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CrmActivitiesDeleteRequest](../../models/operations/crmactivitiesdeleterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CrmActivitiesDeleteResponse](../../models/operations/crmactivitiesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |