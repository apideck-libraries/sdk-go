# Ats.Jobs

## Overview

### Available Operations

* [List](#list) - List Jobs
* [Create](#create) - Create Job
* [Get](#get) - Get Job
* [Update](#update) - Update Job
* [Delete](#delete) - Delete Job

## List

List Jobs

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.jobsAll" method="get" path="/ats/jobs" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Ats.Jobs.List(ctx, operations.AtsJobsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetJobsResponse != nil {
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.AtsJobsAllRequest](../../models/operations/atsjobsallrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AtsJobsAllResponse](../../models/operations/atsjobsallresponse.md), error**

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

Create Job

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.jobsAdd" method="post" path="/ats/jobs" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Ats.Jobs.Create(ctx, operations.AtsJobsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Job: components.JobInput{
            Slug: sdkgo.Pointer("ceo"),
            Title: sdkgo.Pointer("CEO"),
            Sequence: sdkgo.Pointer[int64](3),
            Visibility: components.VisibilityInternal.ToPointer(),
            Status: components.JobStatusCompleted.ToPointer(),
            Code: sdkgo.Pointer("123-OC"),
            Language: sdkgo.Pointer("EN"),
            EmploymentTerms: components.EmploymentTermsFullTime.ToPointer(),
            Experience: sdkgo.Pointer("Director/ Vice President"),
            Location: sdkgo.Pointer("San Francisco"),
            Remote: sdkgo.Pointer(true),
            RequisitionID: sdkgo.Pointer("abc123"),
            Department: &components.DepartmentInput{
                Name: sdkgo.Pointer("R&D"),
                Code: sdkgo.Pointer("2"),
                Description: sdkgo.Pointer("R&D"),
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
            Branch: &components.JobBranch{
                Name: sdkgo.Pointer("HQ NY"),
            },
            Recruiters: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
            },
            HiringManagers: []string{
                "123456",
            },
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Description: sdkgo.Pointer("A description"),
            Blocks: []components.Blocks{
                components.Blocks{
                    Title: sdkgo.Pointer("string"),
                    Content: sdkgo.Pointer("string"),
                },
            },
            Closing: sdkgo.Pointer("The closing section of the job description"),
            ClosingDate: types.MustNewDateFromString("2020-10-30"),
            Salary: &components.Salary{
                Min: sdkgo.Pointer[int64](8000),
                Max: sdkgo.Pointer[int64](10000),
                Currency: components.CurrencyUsd.ToPointer(),
                Interval: sdkgo.Pointer("year"),
            },
            Links: []components.JobLinks{
                components.JobLinks{
                    Type: components.JobTypeJobPortal.ToPointer(),
                    URL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
                },
            },
            Confidential: sdkgo.Pointer(false),
            AvailableToEmployees: sdkgo.Pointer(false),
            Tags: []string{
                "New",
            },
            Addresses: []components.Address{
                components.Address{
                    ID: sdkgo.Pointer("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.Pointer("HQ US"),
                    Line1: sdkgo.Pointer("Main street"),
                    Line2: sdkgo.Pointer("apt #"),
                    Line3: sdkgo.Pointer("Suite #"),
                    Line4: sdkgo.Pointer("delivery instructions"),
                    Line5: sdkgo.Pointer("Attention: Finance Dept"),
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
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                            "Uses Salesforce and Marketo",
                        )),
                    },
                ),
            },
            Deleted: sdkgo.Pointer(true),
            OwnerID: sdkgo.Pointer("54321"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.AtsJobsAddRequest](../../models/operations/atsjobsaddrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AtsJobsAddResponse](../../models/operations/atsjobsaddresponse.md), error**

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

Get Job

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.jobsOne" method="get" path="/ats/jobs/{id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Ats.Jobs.Get(ctx, operations.AtsJobsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.AtsJobsOneRequest](../../models/operations/atsjobsonerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AtsJobsOneResponse](../../models/operations/atsjobsoneresponse.md), error**

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

Update Job

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.jobsUpdate" method="patch" path="/ats/jobs/{id}" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/types"
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

    res, err := s.Ats.Jobs.Update(ctx, operations.AtsJobsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Job: components.JobInput{
            Slug: sdkgo.Pointer("ceo"),
            Title: sdkgo.Pointer("CEO"),
            Sequence: sdkgo.Pointer[int64](3),
            Visibility: components.VisibilityInternal.ToPointer(),
            Status: components.JobStatusCompleted.ToPointer(),
            Code: sdkgo.Pointer("123-OC"),
            Language: sdkgo.Pointer("EN"),
            EmploymentTerms: components.EmploymentTermsFullTime.ToPointer(),
            Experience: sdkgo.Pointer("Director/ Vice President"),
            Location: sdkgo.Pointer("San Francisco"),
            Remote: sdkgo.Pointer(true),
            RequisitionID: sdkgo.Pointer("abc123"),
            Department: &components.DepartmentInput{
                Name: sdkgo.Pointer("R&D"),
                Code: sdkgo.Pointer("2"),
                Description: sdkgo.Pointer("R&D"),
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
            Branch: &components.JobBranch{
                Name: sdkgo.Pointer("HQ NY"),
            },
            Recruiters: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
            },
            HiringManagers: []string{
                "123456",
            },
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Description: sdkgo.Pointer("A description"),
            Blocks: []components.Blocks{
                components.Blocks{
                    Title: sdkgo.Pointer("string"),
                    Content: sdkgo.Pointer("string"),
                },
            },
            Closing: sdkgo.Pointer("The closing section of the job description"),
            ClosingDate: types.MustNewDateFromString("2020-10-30"),
            Salary: &components.Salary{
                Min: sdkgo.Pointer[int64](8000),
                Max: sdkgo.Pointer[int64](10000),
                Currency: components.CurrencyUsd.ToPointer(),
                Interval: sdkgo.Pointer("year"),
            },
            Links: []components.JobLinks{
                components.JobLinks{
                    Type: components.JobTypeJobPortal.ToPointer(),
                    URL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
                },
            },
            Confidential: sdkgo.Pointer(false),
            AvailableToEmployees: sdkgo.Pointer(false),
            Tags: []string{
                "New",
            },
            Addresses: []components.Address{
                components.Address{
                    ID: sdkgo.Pointer("123"),
                    Type: components.TypePrimary.ToPointer(),
                    String: sdkgo.Pointer("25 Spring Street, Blackburn, VIC 3130"),
                    Name: sdkgo.Pointer("HQ US"),
                    Line1: sdkgo.Pointer("Main street"),
                    Line2: sdkgo.Pointer("apt #"),
                    Line3: sdkgo.Pointer("Suite #"),
                    Line4: sdkgo.Pointer("delivery instructions"),
                    Line5: sdkgo.Pointer("Attention: Finance Dept"),
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
            },
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                            "Uses Salesforce and Marketo",
                        )),
                    },
                ),
            },
            Deleted: sdkgo.Pointer(true),
            OwnerID: sdkgo.Pointer("54321"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AtsJobsUpdateRequest](../../models/operations/atsjobsupdaterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AtsJobsUpdateResponse](../../models/operations/atsjobsupdateresponse.md), error**

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

Delete Job

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.jobsDelete" method="delete" path="/ats/jobs/{id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Ats.Jobs.Delete(ctx, operations.AtsJobsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteJobResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.AtsJobsDeleteRequest](../../models/operations/atsjobsdeleterequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.AtsJobsDeleteResponse](../../models/operations/atsjobsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |