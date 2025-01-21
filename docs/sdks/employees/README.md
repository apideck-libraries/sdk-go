# Employees
(*Hris.Employees*)

## Overview

### Available Operations

* [List](#list) - List Employees
* [Create](#create) - Create Employee
* [Get](#get) - Get Employee
* [Update](#update) - Update Employee
* [Delete](#delete) - Delete Employee

## List

Apideck operates as a stateless Unified API, which means that the list endpoint only provides a portion of the employee model. This is due to the fact that most HRIS systems do not readily provide all data in every call. However, you can access the complete employee model through an employee detail call.

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

    res, err := s.Hris.Employees.List(ctx, operations.HrisEmployeesAllRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        Limit: sdkgo.Int64(20),
        Filter: &components.EmployeesFilter{
            CompanyID: sdkgo.String("1234"),
            Email: sdkgo.String("elon@tesla.com"),
            FirstName: sdkgo.String("Elon"),
            Title: sdkgo.String("Manager"),
            LastName: sdkgo.String("Musk"),
            ManagerID: sdkgo.String("1234"),
            EmploymentStatus: components.EmployeesFilterEmploymentStatusActive.ToPointer(),
            EmployeeNumber: sdkgo.String("123456-AB"),
            DepartmentID: sdkgo.String("1234"),
        },
        Sort: &components.EmployeesSort{
            By: components.EmployeesSortByCreatedAt.ToPointer(),
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
    if res.GetEmployeesResponse != nil {
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
| `request`                                                                                | [operations.HrisEmployeesAllRequest](../../models/operations/hrisemployeesallrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.HrisEmployeesAllResponse](../../models/operations/hrisemployeesallresponse.md), error**

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

Create Employee

### Example Usage

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
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Hris.Employees.Create(ctx, operations.HrisEmployeesAddRequest{
        Raw: sdkgo.Bool(false),
        ServiceID: sdkgo.String("salesforce"),
        Employee: components.EmployeeInput{
            ID: sdkgo.String("12345"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            MiddleName: sdkgo.String("D."),
            DisplayName: sdkgo.String("Technoking"),
            PreferredName: sdkgo.String("Elon Musk"),
            Initials: sdkgo.String("EM"),
            Salutation: sdkgo.String("Mr"),
            Title: sdkgo.String("CEO"),
            MaritalStatus: sdkgo.String("married"),
            Partner: &components.PersonInput{
                FirstName: sdkgo.String("Elon"),
                LastName: sdkgo.String("Musk"),
                MiddleName: sdkgo.String("D."),
                Gender: components.GenderMale.ToPointer(),
                Initials: sdkgo.String("EM"),
                Birthday: types.MustNewDateFromString("2000-08-12"),
                DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            },
            Division: sdkgo.String("Europe"),
            DivisionID: sdkgo.String("12345"),
            DepartmentID: sdkgo.String("12345"),
            DepartmentName: sdkgo.String("12345"),
            Team: &components.Team{
                ID: sdkgo.String("1234"),
                Name: sdkgo.String("Full Stack Engineers"),
            },
            CompanyID: sdkgo.String("23456"),
            CompanyName: sdkgo.String("SpaceX"),
            EmploymentStartDate: sdkgo.String("2021-10-26"),
            EmploymentEndDate: sdkgo.String("2028-10-26"),
            LeavingReason: components.LeavingReasonResigned.ToPointer(),
            EmployeeNumber: sdkgo.String("123456-AB"),
            EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            Ethnicity: sdkgo.String("African American"),
            Manager: &components.Manager{
                ID: sdkgo.String("12345"),
                Name: sdkgo.String("Elon Musk"),
                FirstName: sdkgo.String("Elon"),
                LastName: sdkgo.String("Musk"),
                Email: sdkgo.String("elon@musk.com"),
                EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            },
            DirectReports: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            SocialSecurityNumber: sdkgo.String("123456789"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            CountryOfBirth: sdkgo.String("US"),
            Description: sdkgo.String("A description"),
            Gender: components.GenderMale.ToPointer(),
            Pronouns: sdkgo.String("she,her"),
            PreferredLanguage: sdkgo.String("EN"),
            Languages: []string{
                "EN",
            },
            Nationalities: []string{
                "US",
            },
            PhotoURL: sdkgo.String("https://unavatar.io/elon-musk"),
            Timezone: sdkgo.String("Europe/London"),
            Source: sdkgo.String("lever"),
            SourceID: sdkgo.String("12345"),
            RecordURL: sdkgo.String("https://app.intercom.io/contacts/12345"),
            Jobs: []components.EmployeeJobInput{
                components.EmployeeJobInput{
                    Title: sdkgo.String("CEO"),
                    Role: sdkgo.String("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Float64(72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Bool(true),
                    IsManager: sdkgo.Bool(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                components.EmployeeJobInput{
                    Title: sdkgo.String("CEO"),
                    Role: sdkgo.String("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Float64(72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Bool(true),
                    IsManager: sdkgo.Bool(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
            },
            Compensations: []components.EmployeeCompensationInput{
                components.EmployeeCompensationInput{
                    Rate: sdkgo.Float64(50),
                    PaymentUnit: components.PaymentUnitHour.ToPointer(),
                    FlsaStatus: components.FlsaStatusNonexempt.ToPointer(),
                    EffectiveDate: sdkgo.String("2021-06-11"),
                },
            },
            WorksRemote: sdkgo.Bool(true),
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
            },
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueArrayOfStr(
                        []string{
                            "<value>",
                            "<value>",
                            "<value>",
                        },
                    )),
                },
            },
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.String("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.String("twitter"),
                },
            },
            BankAccounts: []components.BankAccount{
                components.BankAccount{
                    BankName: sdkgo.String("Monzo"),
                    AccountNumber: sdkgo.String("123465"),
                    AccountName: sdkgo.String("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.String("CH2989144532982975332"),
                    Bic: sdkgo.String("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.String("012345678"),
                    BsbNumber: sdkgo.String("062-001"),
                    BranchIdentifier: sdkgo.String("001"),
                    BankCode: sdkgo.String("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            TaxCode: sdkgo.String("1111"),
            TaxID: sdkgo.String("234-32-0000"),
            DietaryPreference: sdkgo.String("Veggie"),
            FoodAllergies: []string{
                "No allergies",
            },
            ProbationPeriod: &components.ProbationPeriod{
                StartDate: types.MustNewDateFromString("2021-10-01"),
                EndDate: types.MustNewDateFromString("2021-11-28"),
            },
            Tags: []string{
                "New",
            },
            RowVersion: sdkgo.String("1-12345"),
            Deleted: sdkgo.Bool(true),
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
    if res.CreateEmployeeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.HrisEmployeesAddRequest](../../models/operations/hrisemployeesaddrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.HrisEmployeesAddResponse](../../models/operations/hrisemployeesaddresponse.md), error**

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

Get Employee

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

    res, err := s.Hris.Employees.Get(ctx, operations.HrisEmployeesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        Fields: sdkgo.String("id,updated_at"),
        Filter: &components.EmployeesOneFilter{
            CompanyID: sdkgo.String("1234"),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetEmployeeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.HrisEmployeesOneRequest](../../models/operations/hrisemployeesonerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.HrisEmployeesOneResponse](../../models/operations/hrisemployeesoneresponse.md), error**

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

Update Employee

### Example Usage

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
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
    )

    res, err := s.Hris.Employees.Update(ctx, operations.HrisEmployeesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
        Employee: components.EmployeeInput{
            ID: sdkgo.String("12345"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            MiddleName: sdkgo.String("D."),
            DisplayName: sdkgo.String("Technoking"),
            PreferredName: sdkgo.String("Elon Musk"),
            Initials: sdkgo.String("EM"),
            Salutation: sdkgo.String("Mr"),
            Title: sdkgo.String("CEO"),
            MaritalStatus: sdkgo.String("married"),
            Partner: &components.PersonInput{
                FirstName: sdkgo.String("Elon"),
                LastName: sdkgo.String("Musk"),
                MiddleName: sdkgo.String("D."),
                Gender: components.GenderMale.ToPointer(),
                Initials: sdkgo.String("EM"),
                Birthday: types.MustNewDateFromString("2000-08-12"),
                DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            },
            Division: sdkgo.String("Europe"),
            DivisionID: sdkgo.String("12345"),
            DepartmentID: sdkgo.String("12345"),
            DepartmentName: sdkgo.String("12345"),
            Team: &components.Team{
                ID: sdkgo.String("1234"),
                Name: sdkgo.String("Full Stack Engineers"),
            },
            CompanyID: sdkgo.String("23456"),
            CompanyName: sdkgo.String("SpaceX"),
            EmploymentStartDate: sdkgo.String("2021-10-26"),
            EmploymentEndDate: sdkgo.String("2028-10-26"),
            LeavingReason: components.LeavingReasonResigned.ToPointer(),
            EmployeeNumber: sdkgo.String("123456-AB"),
            EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            Ethnicity: sdkgo.String("African American"),
            Manager: &components.Manager{
                ID: sdkgo.String("12345"),
                Name: sdkgo.String("Elon Musk"),
                FirstName: sdkgo.String("Elon"),
                LastName: sdkgo.String("Musk"),
                Email: sdkgo.String("elon@musk.com"),
                EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            },
            DirectReports: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            SocialSecurityNumber: sdkgo.String("123456789"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            CountryOfBirth: sdkgo.String("US"),
            Description: sdkgo.String("A description"),
            Gender: components.GenderMale.ToPointer(),
            Pronouns: sdkgo.String("she,her"),
            PreferredLanguage: sdkgo.String("EN"),
            Languages: []string{
                "EN",
            },
            Nationalities: []string{
                "US",
            },
            PhotoURL: sdkgo.String("https://unavatar.io/elon-musk"),
            Timezone: sdkgo.String("Europe/London"),
            Source: sdkgo.String("lever"),
            SourceID: sdkgo.String("12345"),
            RecordURL: sdkgo.String("https://app.intercom.io/contacts/12345"),
            Jobs: []components.EmployeeJobInput{
                components.EmployeeJobInput{
                    Title: sdkgo.String("CEO"),
                    Role: sdkgo.String("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Float64(72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Bool(true),
                    IsManager: sdkgo.Bool(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                components.EmployeeJobInput{
                    Title: sdkgo.String("CEO"),
                    Role: sdkgo.String("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Float64(72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Bool(true),
                    IsManager: sdkgo.Bool(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                components.EmployeeJobInput{
                    Title: sdkgo.String("CEO"),
                    Role: sdkgo.String("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Float64(72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Bool(true),
                    IsManager: sdkgo.Bool(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
            },
            Compensations: []components.EmployeeCompensationInput{
                components.EmployeeCompensationInput{
                    Rate: sdkgo.Float64(50),
                    PaymentUnit: components.PaymentUnitHour.ToPointer(),
                    FlsaStatus: components.FlsaStatusNonexempt.ToPointer(),
                    EffectiveDate: sdkgo.String("2021-06-11"),
                },
            },
            WorksRemote: sdkgo.Bool(true),
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
                components.Email{
                    ID: sdkgo.String("123"),
                    Email: sdkgo.String("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
            },
            CustomFields: []components.CustomField{
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
                    Value: sdkgo.Pointer(components.CreateValueFour(
                        components.Four{},
                    )),
                },
            },
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.String("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.String("twitter"),
                },
                components.SocialLink{
                    ID: sdkgo.String("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.String("twitter"),
                },
                components.SocialLink{
                    ID: sdkgo.String("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.String("twitter"),
                },
            },
            BankAccounts: []components.BankAccount{
                components.BankAccount{
                    BankName: sdkgo.String("Monzo"),
                    AccountNumber: sdkgo.String("123465"),
                    AccountName: sdkgo.String("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.String("CH2989144532982975332"),
                    Bic: sdkgo.String("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.String("012345678"),
                    BsbNumber: sdkgo.String("062-001"),
                    BranchIdentifier: sdkgo.String("001"),
                    BankCode: sdkgo.String("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount{
                    BankName: sdkgo.String("Monzo"),
                    AccountNumber: sdkgo.String("123465"),
                    AccountName: sdkgo.String("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.String("CH2989144532982975332"),
                    Bic: sdkgo.String("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.String("012345678"),
                    BsbNumber: sdkgo.String("062-001"),
                    BranchIdentifier: sdkgo.String("001"),
                    BankCode: sdkgo.String("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            TaxCode: sdkgo.String("1111"),
            TaxID: sdkgo.String("234-32-0000"),
            DietaryPreference: sdkgo.String("Veggie"),
            FoodAllergies: []string{
                "No allergies",
            },
            ProbationPeriod: &components.ProbationPeriod{
                StartDate: types.MustNewDateFromString("2021-10-01"),
                EndDate: types.MustNewDateFromString("2021-11-28"),
            },
            Tags: []string{
                "New",
            },
            RowVersion: sdkgo.String("1-12345"),
            Deleted: sdkgo.Bool(true),
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
    if res.UpdateEmployeeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.HrisEmployeesUpdateRequest](../../models/operations/hrisemployeesupdaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.HrisEmployeesUpdateResponse](../../models/operations/hrisemployeesupdateresponse.md), error**

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

Delete Employee

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

    res, err := s.Hris.Employees.Delete(ctx, operations.HrisEmployeesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Raw: sdkgo.Bool(false),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteEmployeeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.HrisEmployeesDeleteRequest](../../models/operations/hrisemployeesdeleterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.HrisEmployeesDeleteResponse](../../models/operations/hrisemployeesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |