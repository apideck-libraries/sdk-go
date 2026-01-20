# Accounting.Projects

## Overview

### Available Operations

* [List](#list) - List projects
* [Create](#create) - Create project
* [Get](#get) - Get project
* [Update](#update) - Update project
* [Delete](#delete) - Delete project

## List

List projects

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.projectsAll" method="get" path="/accounting/projects" -->
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

    res, err := s.Accounting.Projects.List(ctx, operations.AccountingProjectsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.ProjectsFilter{
            Name: sdkgo.Pointer("Website Redesign"),
            Status: components.ProjectStatusActive.ToPointer(),
            CustomerID: sdkgo.Pointer("12345"),
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
        },
        Sort: &components.ProjectsSort{
            By: components.ProjectsSortSortByName.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetProjectsResponse != nil {
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingProjectsAllRequest](../../models/operations/accountingprojectsallrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingProjectsAllResponse](../../models/operations/accountingprojectsallresponse.md), error**

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

Create project

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.projectsAdd" method="post" path="/accounting/projects" -->
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

    res, err := s.Accounting.Projects.Create(ctx, operations.AccountingProjectsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Project: components.ProjectInput{
            Name: "Website Redesign Project",
            DisplayID: sdkgo.Pointer("PROJ-001"),
            ReferenceID: sdkgo.Pointer("WD-REF-2024-001"),
            Description: sdkgo.Pointer("Complete redesign of the company website including new branding and improved user experience"),
            Status: components.ProjectProjectStatusActive.ToPointer(),
            Active: sdkgo.Pointer(true),
            ProjectType: components.ProjectTypeClientProject.ToPointer(),
            Priority: components.PriorityHigh.ToPointer(),
            CompletionPercentage: sdkgo.Pointer[float64](75.5),
            StartDate: types.MustNewDateFromString("2024-01-15"),
            EndDate: types.MustNewDateFromString("2024-06-30"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            Department: &components.ProjectDepartment{
                ID: sdkgo.Pointer("DEPT-001"),
                Name: sdkgo.Pointer("Marketing"),
            },
            CompanyID: sdkgo.Pointer("12345"),
            OwnerID: sdkgo.Pointer("12345"),
            ParentProject: &components.ParentProject{
                ID: sdkgo.Pointer("PROJ-PARENT-001"),
                Name: sdkgo.Pointer("Company Website Overhaul"),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            BudgetAmount: sdkgo.Pointer[float64](50000),
            ApprovedAmount: sdkgo.Pointer[float64](48000),
            BudgetHours: sdkgo.Pointer[float64](400),
            HourlyRate: sdkgo.Pointer[float64](125),
            BillingMethod: components.BillingMethodTimeAndMaterials.ToPointer(),
            Phase: components.ProjectPhaseExecution.ToPointer(),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Code: sdkgo.Pointer("N-T"),
                Rate: sdkgo.Pointer[float64](10),
            },
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            Tags: []string{
                "website",
                "redesign",
                "urgent",
            },
            Notes: sdkgo.Pointer("Client has requested modern design with mobile-first approach"),
            ContractNumber: sdkgo.Pointer("CNT-2024-001"),
            ProfitMargin: sdkgo.Pointer[float64](15.5),
            ScheduleStatus: components.ScheduleStatusOnSchedule.ToPointer(),
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
            TeamSize: sdkgo.Pointer[int64](8),
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
            RowVersion: sdkgo.Pointer("1-12345"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingProjectsAddRequest](../../models/operations/accountingprojectsaddrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingProjectsAddResponse](../../models/operations/accountingprojectsaddresponse.md), error**

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

Get project

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.projectsOne" method="get" path="/accounting/projects/{id}" -->
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

    res, err := s.Accounting.Projects.Get(ctx, operations.AccountingProjectsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.AccountingProjectsOneRequest](../../models/operations/accountingprojectsonerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.AccountingProjectsOneResponse](../../models/operations/accountingprojectsoneresponse.md), error**

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

Update project

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.projectsUpdate" method="patch" path="/accounting/projects/{id}" -->
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

    res, err := s.Accounting.Projects.Update(ctx, operations.AccountingProjectsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Project: components.ProjectInput{
            Name: "Website Redesign Project",
            DisplayID: sdkgo.Pointer("PROJ-001"),
            ReferenceID: sdkgo.Pointer("WD-REF-2024-001"),
            Description: sdkgo.Pointer("Complete redesign of the company website including new branding and improved user experience"),
            Status: components.ProjectProjectStatusActive.ToPointer(),
            Active: sdkgo.Pointer(true),
            ProjectType: components.ProjectTypeClientProject.ToPointer(),
            Priority: components.PriorityHigh.ToPointer(),
            CompletionPercentage: sdkgo.Pointer[float64](75.5),
            StartDate: types.MustNewDateFromString("2024-01-15"),
            EndDate: types.MustNewDateFromString("2024-06-30"),
            Customer: nil,
            Department: &components.ProjectDepartment{
                ID: sdkgo.Pointer("DEPT-001"),
                Name: sdkgo.Pointer("Marketing"),
            },
            CompanyID: sdkgo.Pointer("12345"),
            OwnerID: sdkgo.Pointer("12345"),
            ParentProject: &components.ParentProject{
                ID: sdkgo.Pointer("PROJ-PARENT-001"),
                Name: sdkgo.Pointer("Company Website Overhaul"),
            },
            Currency: components.CurrencyUsd.ToPointer(),
            BudgetAmount: sdkgo.Pointer[float64](50000),
            ApprovedAmount: sdkgo.Pointer[float64](48000),
            BudgetHours: sdkgo.Pointer[float64](400),
            HourlyRate: sdkgo.Pointer[float64](125),
            BillingMethod: components.BillingMethodTimeAndMaterials.ToPointer(),
            Phase: components.ProjectPhaseExecution.ToPointer(),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Code: sdkgo.Pointer("N-T"),
                Rate: sdkgo.Pointer[float64](10),
            },
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            Tags: []string{
                "website",
                "redesign",
                "urgent",
            },
            Notes: sdkgo.Pointer("Client has requested modern design with mobile-first approach"),
            ContractNumber: sdkgo.Pointer("CNT-2024-001"),
            ProfitMargin: sdkgo.Pointer[float64](15.5),
            ScheduleStatus: components.ScheduleStatusOnSchedule.ToPointer(),
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
            TeamSize: sdkgo.Pointer[int64](8),
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
            RowVersion: sdkgo.Pointer("1-12345"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingProjectsUpdateRequest](../../models/operations/accountingprojectsupdaterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingProjectsUpdateResponse](../../models/operations/accountingprojectsupdateresponse.md), error**

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

Delete project

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.projectsDelete" method="delete" path="/accounting/projects/{id}" -->
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

    res, err := s.Accounting.Projects.Delete(ctx, operations.AccountingProjectsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteProjectResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingProjectsDeleteRequest](../../models/operations/accountingprojectsdeleterequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingProjectsDeleteResponse](../../models/operations/accountingprojectsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |