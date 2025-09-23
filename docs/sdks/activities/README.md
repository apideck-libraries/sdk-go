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

<!-- UsageSnippet language="go" operationID="crm.activitiesAll" method="get" path="/crm/activities" -->
```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Crm.Activities.List(ctx, operations.CrmActivitiesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
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
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="crm.activitiesAdd" method="post" path="/crm/activities" -->
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

    res, err := s.Crm.Activities.Create(ctx, operations.CrmActivitiesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Activity: components.ActivityInput{
            ActivityDatetime: sdkgo.Pointer("2021-05-01T12:00:00.000Z"),
            DurationSeconds: sdkgo.Pointer[int64](1800),
            UserID: sdkgo.Pointer("12345"),
            AccountID: sdkgo.Pointer("12345"),
            ContactID: sdkgo.Pointer("12345"),
            CompanyID: sdkgo.Pointer("12345"),
            OpportunityID: sdkgo.Pointer("12345"),
            LeadID: sdkgo.Pointer("12345"),
            OwnerID: sdkgo.Pointer("12345"),
            CampaignID: sdkgo.Pointer("12345"),
            CaseID: sdkgo.Pointer("12345"),
            AssetID: sdkgo.Pointer("12345"),
            ContractID: sdkgo.Pointer("12345"),
            ProductID: sdkgo.Pointer("12345"),
            SolutionID: sdkgo.Pointer("12345"),
            CustomObjectID: sdkgo.Pointer("12345"),
            Type: components.ActivityTypeMeeting.ToPointer(),
            Title: sdkgo.Pointer("Meeting"),
            Description: sdkgo.Pointer("More info about the meeting"),
            Note: sdkgo.Pointer("An internal note about the meeting"),
            Location: sdkgo.Pointer("Space"),
            LocationAddress: &components.Address{
                ID: sdkgo.Pointer("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.Pointer("HQ US"),
                Line1: sdkgo.Pointer("Main street"),
                Line2: sdkgo.Pointer("apt #"),
                Line3: sdkgo.Pointer("Suite #"),
                Line4: sdkgo.Pointer("delivery instructions"),
                StreetNumber: sdkgo.Pointer("25"),
                City: sdkgo.Pointer("San Francisco"),
                State: sdkgo.Pointer("CA"),
                PostalCode: sdkgo.Pointer("94104"),
                Country: sdkgo.Pointer("US"),
                Latitude: sdkgo.Pointer("40.759211"),
                Longitude: sdkgo.Pointer("-73.984638"),
                County: sdkgo.Pointer("Santa Clara"),
                ContactName: sdkgo.Pointer("Elon Musk"),
                Salutation: sdkgo.Pointer("Mr"),
                PhoneNumber: sdkgo.Pointer("111-111-1111"),
                Fax: sdkgo.Pointer("122-111-1111"),
                Email: sdkgo.Pointer("elon@musk.com"),
                Website: sdkgo.Pointer("https://elonmusk.com"),
                Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                RowVersion: sdkgo.Pointer("1-12345"),
            },
            AllDayEvent: sdkgo.Pointer(false),
            Private: sdkgo.Pointer(true),
            GroupEvent: sdkgo.Pointer(true),
            EventSubType: sdkgo.Pointer("debrief"),
            GroupEventType: sdkgo.Pointer("Proposed"),
            Child: sdkgo.Pointer(false),
            Archived: sdkgo.Pointer(false),
            Deleted: sdkgo.Pointer(false),
            ShowAs: components.ShowAsBusy.ToPointer(),
            Done: sdkgo.Pointer(false),
            StartDatetime: sdkgo.Pointer("2021-05-01T12:00:00.000Z"),
            EndDatetime: sdkgo.Pointer("2021-05-01T12:30:00.000Z"),
            ActivityDate: sdkgo.Pointer("2021-05-01"),
            EndDate: sdkgo.Pointer("2021-05-01"),
            Recurrent: sdkgo.Pointer(false),
            ReminderDatetime: sdkgo.Pointer("2021-05-01T17:00:00.000Z"),
            ReminderSet: sdkgo.Pointer(false),
            VideoConferenceURL: sdkgo.Pointer("https://us02web.zoom.us/j/88120759396"),
            VideoConferenceID: sdkgo.Pointer("zoom:88120759396"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: nil,
                },
            },
            Attendees: []components.ActivityAttendeeInput{
                components.ActivityAttendeeInput{
                    Name: sdkgo.Pointer("Elon Musk"),
                    FirstName: sdkgo.Pointer("Elon"),
                    MiddleName: sdkgo.Pointer("D."),
                    LastName: sdkgo.Pointer("Musk"),
                    Prefix: sdkgo.Pointer("Mr."),
                    Suffix: sdkgo.Pointer("PhD"),
                    EmailAddress: sdkgo.Pointer("elon@musk.com"),
                    IsOrganizer: sdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="crm.activitiesOne" method="get" path="/crm/activities/{id}" -->
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

    res, err := s.Crm.Activities.Get(ctx, operations.CrmActivitiesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="crm.activitiesUpdate" method="patch" path="/crm/activities/{id}" -->
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

    res, err := s.Crm.Activities.Update(ctx, operations.CrmActivitiesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Activity: components.ActivityInput{
            ActivityDatetime: sdkgo.Pointer("2021-05-01T12:00:00.000Z"),
            DurationSeconds: sdkgo.Pointer[int64](1800),
            UserID: sdkgo.Pointer("12345"),
            AccountID: sdkgo.Pointer("12345"),
            ContactID: sdkgo.Pointer("12345"),
            CompanyID: sdkgo.Pointer("12345"),
            OpportunityID: sdkgo.Pointer("12345"),
            LeadID: sdkgo.Pointer("12345"),
            OwnerID: sdkgo.Pointer("12345"),
            CampaignID: sdkgo.Pointer("12345"),
            CaseID: sdkgo.Pointer("12345"),
            AssetID: sdkgo.Pointer("12345"),
            ContractID: sdkgo.Pointer("12345"),
            ProductID: sdkgo.Pointer("12345"),
            SolutionID: sdkgo.Pointer("12345"),
            CustomObjectID: sdkgo.Pointer("12345"),
            Type: components.ActivityTypeMeeting.ToPointer(),
            Title: sdkgo.Pointer("Meeting"),
            Description: sdkgo.Pointer("More info about the meeting"),
            Note: sdkgo.Pointer("An internal note about the meeting"),
            Location: sdkgo.Pointer("Space"),
            LocationAddress: &components.Address{
                ID: sdkgo.Pointer("123"),
                Type: components.TypePrimary.ToPointer(),
                String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                Name: sdkgo.Pointer("HQ US"),
                Line1: sdkgo.Pointer("Main street"),
                Line2: sdkgo.Pointer("apt #"),
                Line3: sdkgo.Pointer("Suite #"),
                Line4: sdkgo.Pointer("delivery instructions"),
                StreetNumber: sdkgo.Pointer("25"),
                City: sdkgo.Pointer("San Francisco"),
                State: sdkgo.Pointer("CA"),
                PostalCode: sdkgo.Pointer("94104"),
                Country: sdkgo.Pointer("US"),
                Latitude: sdkgo.Pointer("40.759211"),
                Longitude: sdkgo.Pointer("-73.984638"),
                County: sdkgo.Pointer("Santa Clara"),
                ContactName: sdkgo.Pointer("Elon Musk"),
                Salutation: sdkgo.Pointer("Mr"),
                PhoneNumber: sdkgo.Pointer("111-111-1111"),
                Fax: sdkgo.Pointer("122-111-1111"),
                Email: sdkgo.Pointer("elon@musk.com"),
                Website: sdkgo.Pointer("https://elonmusk.com"),
                Notes: sdkgo.Pointer("Address notes or delivery instructions."),
                RowVersion: sdkgo.Pointer("1-12345"),
            },
            AllDayEvent: sdkgo.Pointer(false),
            Private: sdkgo.Pointer(true),
            GroupEvent: sdkgo.Pointer(true),
            EventSubType: sdkgo.Pointer("debrief"),
            GroupEventType: sdkgo.Pointer("Proposed"),
            Child: sdkgo.Pointer(false),
            Archived: sdkgo.Pointer(false),
            Deleted: sdkgo.Pointer(false),
            ShowAs: components.ShowAsBusy.ToPointer(),
            Done: sdkgo.Pointer(false),
            StartDatetime: sdkgo.Pointer("2021-05-01T12:00:00.000Z"),
            EndDatetime: sdkgo.Pointer("2021-05-01T12:30:00.000Z"),
            ActivityDate: sdkgo.Pointer("2021-05-01"),
            EndDate: sdkgo.Pointer("2021-05-01"),
            Recurrent: sdkgo.Pointer(false),
            ReminderDatetime: sdkgo.Pointer("2021-05-01T17:00:00.000Z"),
            ReminderSet: sdkgo.Pointer(false),
            VideoConferenceURL: sdkgo.Pointer("https://us02web.zoom.us/j/88120759396"),
            VideoConferenceID: sdkgo.Pointer("zoom:88120759396"),
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
            Attendees: []components.ActivityAttendeeInput{
                components.ActivityAttendeeInput{
                    Name: sdkgo.Pointer("Elon Musk"),
                    FirstName: sdkgo.Pointer("Elon"),
                    MiddleName: sdkgo.Pointer("D."),
                    LastName: sdkgo.Pointer("Musk"),
                    Prefix: sdkgo.Pointer("Mr."),
                    Suffix: sdkgo.Pointer("PhD"),
                    EmailAddress: sdkgo.Pointer("elon@musk.com"),
                    IsOrganizer: sdkgo.Pointer(true),
                    Status: components.ActivityAttendeeStatusAccepted.ToPointer(),
                },
                components.ActivityAttendeeInput{
                    Name: sdkgo.Pointer("Elon Musk"),
                    FirstName: sdkgo.Pointer("Elon"),
                    MiddleName: sdkgo.Pointer("D."),
                    LastName: sdkgo.Pointer("Musk"),
                    Prefix: sdkgo.Pointer("Mr."),
                    Suffix: sdkgo.Pointer("PhD"),
                    EmailAddress: sdkgo.Pointer("elon@musk.com"),
                    IsOrganizer: sdkgo.Pointer(true),
                    Status: components.ActivityAttendeeStatusAccepted.ToPointer(),
                },
                components.ActivityAttendeeInput{
                    Name: sdkgo.Pointer("Elon Musk"),
                    FirstName: sdkgo.Pointer("Elon"),
                    MiddleName: sdkgo.Pointer("D."),
                    LastName: sdkgo.Pointer("Musk"),
                    Prefix: sdkgo.Pointer("Mr."),
                    Suffix: sdkgo.Pointer("PhD"),
                    EmailAddress: sdkgo.Pointer("elon@musk.com"),
                    IsOrganizer: sdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="crm.activitiesDelete" method="delete" path="/crm/activities/{id}" -->
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

    res, err := s.Crm.Activities.Delete(ctx, operations.CrmActivitiesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
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