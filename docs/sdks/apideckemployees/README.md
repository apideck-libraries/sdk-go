# Hris.Employees

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

<!-- UsageSnippet language="go" operationID="hris.employeesAll" method="get" path="/hris/employees" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Hris.Employees.List(ctx, operations.HrisEmployeesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.EmployeesFilter{
            CompanyID: sdkgo.Pointer("1234"),
            Email: sdkgo.Pointer("elon@tesla.com"),
            FirstName: sdkgo.Pointer("Elon"),
            Title: sdkgo.Pointer("Manager"),
            LastName: sdkgo.Pointer("Musk"),
            ManagerID: sdkgo.Pointer("1234"),
            EmploymentStatus: components.EmployeesFilterEmploymentStatusActive.ToPointer(),
            EmployeeNumber: sdkgo.Pointer("123456-AB"),
            DepartmentID: sdkgo.Pointer("1234"),
            City: sdkgo.Pointer("San Francisco"),
            Country: sdkgo.Pointer("US"),
        },
        Sort: &components.EmployeesSort{
            By: components.EmployeesSortByCreatedAt.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="hris.employeesAdd" method="post" path="/hris/employees" -->
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

    res, err := s.Hris.Employees.Create(ctx, operations.HrisEmployeesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Employee: components.EmployeeInput{
            ID: sdkgo.Pointer("12345"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            MiddleName: sdkgo.Pointer("D."),
            DisplayName: sdkgo.Pointer("Technoking"),
            PreferredName: sdkgo.Pointer("Elon Musk"),
            Initials: sdkgo.Pointer("EM"),
            Salutation: sdkgo.Pointer("Mr"),
            Title: sdkgo.Pointer("CEO"),
            MaritalStatus: sdkgo.Pointer("married"),
            Partner: &components.PersonInput{
                FirstName: sdkgo.Pointer("Elon"),
                LastName: sdkgo.Pointer("Musk"),
                MiddleName: sdkgo.Pointer("D."),
                Gender: components.GenderMale.ToPointer(),
                Initials: sdkgo.Pointer("EM"),
                Birthday: types.MustNewDateFromString("2000-08-12"),
                DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            },
            Division: sdkgo.Pointer("Europe"),
            DivisionID: sdkgo.Pointer("12345"),
            DepartmentID: sdkgo.Pointer("12345"),
            DepartmentName: sdkgo.Pointer("12345"),
            Team: &components.Team{
                ID: sdkgo.Pointer("1234"),
                Name: sdkgo.Pointer("Full Stack Engineers"),
            },
            CompanyID: sdkgo.Pointer("23456"),
            CompanyName: sdkgo.Pointer("SpaceX"),
            EmploymentStartDate: sdkgo.Pointer("2021-10-26"),
            EmploymentEndDate: sdkgo.Pointer("2028-10-26"),
            LeavingReason: components.LeavingReasonResigned.ToPointer(),
            EmployeeNumber: sdkgo.Pointer("123456-AB"),
            EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            Ethnicity: sdkgo.Pointer("African American"),
            Manager: &components.Manager{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Elon Musk"),
                FirstName: sdkgo.Pointer("Elon"),
                LastName: sdkgo.Pointer("Musk"),
                Email: sdkgo.Pointer("elon@musk.com"),
                EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            },
            DirectReports: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            SocialSecurityNumber: sdkgo.Pointer("123456789"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            CountryOfBirth: sdkgo.Pointer("US"),
            Description: sdkgo.Pointer("A description"),
            Gender: components.GenderMale.ToPointer(),
            Pronouns: sdkgo.Pointer("she,her"),
            PreferredLanguage: sdkgo.Pointer("EN"),
            Languages: []*string{
                sdkgo.Pointer("EN"),
            },
            Nationalities: []*string{
                sdkgo.Pointer("US"),
            },
            PhotoURL: sdkgo.Pointer("https://unavatar.io/elon-musk"),
            Timezone: sdkgo.Pointer("Europe/London"),
            Source: sdkgo.Pointer("lever"),
            SourceID: sdkgo.Pointer("12345"),
            RecordURL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
            Jobs: []components.EmployeeJobInput{
                components.EmployeeJobInput{
                    Title: sdkgo.Pointer("CEO"),
                    Role: sdkgo.Pointer("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Pointer[float64](72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Pointer(true),
                    IsManager: sdkgo.Pointer(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                },
                components.EmployeeJobInput{
                    Title: sdkgo.Pointer("CEO"),
                    Role: sdkgo.Pointer("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Pointer[float64](72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Pointer(true),
                    IsManager: sdkgo.Pointer(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                },
                components.EmployeeJobInput{
                    Title: sdkgo.Pointer("CEO"),
                    Role: sdkgo.Pointer("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Pointer[float64](72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Pointer(true),
                    IsManager: sdkgo.Pointer(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                },
            },
            Compensations: []components.EmployeeCompensationInput{
                components.EmployeeCompensationInput{
                    Rate: sdkgo.Pointer[float64](50),
                    PaymentUnit: components.PaymentUnitHour.ToPointer(),
                    FlsaStatus: components.FlsaStatusNonexempt.ToPointer(),
                    EffectiveDate: sdkgo.Pointer("2021-06-11"),
                },
            },
            WorksRemote: sdkgo.Pointer(true),
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
                components.Address{
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
            },
            PhoneNumbers: []components.PhoneNumber{
                components.PhoneNumber{
                    ID: sdkgo.Pointer("12345"),
                    CountryCode: sdkgo.Pointer("1"),
                    AreaCode: sdkgo.Pointer("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.Pointer("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
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
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
            },
            BankAccounts: []components.BankAccount2{
                components.BankAccount2{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount2AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount2{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount2AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount2{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount2AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            TaxCode: sdkgo.Pointer("1111"),
            TaxID: sdkgo.Pointer("234-32-0000"),
            DietaryPreference: sdkgo.Pointer("Veggie"),
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
            RowVersion: sdkgo.Pointer("1-12345"),
            Deleted: sdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="hris.employeesOne" method="get" path="/hris/employees/{id}" -->
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
        sdkgo.WithConsumerID("test-consumer"),
        sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
        sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
    )

    res, err := s.Hris.Employees.Get(ctx, operations.HrisEmployeesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
        Filter: &components.EmployeesOneFilter{
            CompanyID: sdkgo.Pointer("1234"),
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

<!-- UsageSnippet language="go" operationID="hris.employeesUpdate" method="patch" path="/hris/employees/{id}" -->
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

    res, err := s.Hris.Employees.Update(ctx, operations.HrisEmployeesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Employee: components.EmployeeInput{
            ID: sdkgo.Pointer("12345"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            MiddleName: sdkgo.Pointer("D."),
            DisplayName: sdkgo.Pointer("Technoking"),
            PreferredName: sdkgo.Pointer("Elon Musk"),
            Initials: sdkgo.Pointer("EM"),
            Salutation: sdkgo.Pointer("Mr"),
            Title: sdkgo.Pointer("CEO"),
            MaritalStatus: sdkgo.Pointer("married"),
            Partner: &components.PersonInput{
                FirstName: sdkgo.Pointer("Elon"),
                LastName: sdkgo.Pointer("Musk"),
                MiddleName: sdkgo.Pointer("D."),
                Gender: components.GenderMale.ToPointer(),
                Initials: sdkgo.Pointer("EM"),
                Birthday: types.MustNewDateFromString("2000-08-12"),
                DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            },
            Division: sdkgo.Pointer("Europe"),
            DivisionID: sdkgo.Pointer("12345"),
            DepartmentID: sdkgo.Pointer("12345"),
            DepartmentName: sdkgo.Pointer("12345"),
            Team: &components.Team{
                ID: sdkgo.Pointer("1234"),
                Name: sdkgo.Pointer("Full Stack Engineers"),
            },
            CompanyID: sdkgo.Pointer("23456"),
            CompanyName: sdkgo.Pointer("SpaceX"),
            EmploymentStartDate: sdkgo.Pointer("2021-10-26"),
            EmploymentEndDate: sdkgo.Pointer("2028-10-26"),
            LeavingReason: components.LeavingReasonResigned.ToPointer(),
            EmployeeNumber: sdkgo.Pointer("123456-AB"),
            EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            Ethnicity: sdkgo.Pointer("African American"),
            Manager: &components.Manager{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Elon Musk"),
                FirstName: sdkgo.Pointer("Elon"),
                LastName: sdkgo.Pointer("Musk"),
                Email: sdkgo.Pointer("elon@musk.com"),
                EmploymentStatus: components.EmploymentStatusActive.ToPointer(),
            },
            DirectReports: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            SocialSecurityNumber: sdkgo.Pointer("123456789"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            DeceasedOn: types.MustNewDateFromString("2000-08-12"),
            CountryOfBirth: sdkgo.Pointer("US"),
            Description: sdkgo.Pointer("A description"),
            Gender: components.GenderMale.ToPointer(),
            Pronouns: sdkgo.Pointer("she,her"),
            PreferredLanguage: sdkgo.Pointer("EN"),
            Languages: []*string{
                sdkgo.Pointer("EN"),
            },
            Nationalities: []*string{
                sdkgo.Pointer("US"),
            },
            PhotoURL: sdkgo.Pointer("https://unavatar.io/elon-musk"),
            Timezone: sdkgo.Pointer("Europe/London"),
            Source: sdkgo.Pointer("lever"),
            SourceID: sdkgo.Pointer("12345"),
            RecordURL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
            Jobs: []components.EmployeeJobInput{
                components.EmployeeJobInput{
                    Title: sdkgo.Pointer("CEO"),
                    Role: sdkgo.Pointer("Sales"),
                    StartDate: types.MustNewDateFromString("2020-08-12"),
                    EndDate: types.MustNewDateFromString("2020-08-12"),
                    CompensationRate: sdkgo.Pointer[float64](72000),
                    Currency: components.CurrencyUsd.ToPointer(),
                    PaymentUnit: components.PaymentUnitYear.ToPointer(),
                    HiredAt: types.MustNewDateFromString("2020-08-12"),
                    IsPrimary: sdkgo.Pointer(true),
                    IsManager: sdkgo.Pointer(true),
                    Status: components.EmployeeJobStatusActive.ToPointer(),
                    Location: &components.Address{
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
                },
            },
            Compensations: []components.EmployeeCompensationInput{
                components.EmployeeCompensationInput{
                    Rate: sdkgo.Pointer[float64](50),
                    PaymentUnit: components.PaymentUnitHour.ToPointer(),
                    FlsaStatus: components.FlsaStatusNonexempt.ToPointer(),
                    EffectiveDate: sdkgo.Pointer("2021-06-11"),
                },
            },
            WorksRemote: sdkgo.Pointer(true),
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
            PhoneNumbers: []components.PhoneNumber{
                components.PhoneNumber{
                    ID: sdkgo.Pointer("12345"),
                    CountryCode: sdkgo.Pointer("1"),
                    AreaCode: sdkgo.Pointer("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.Pointer("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
                components.PhoneNumber{
                    ID: sdkgo.Pointer("12345"),
                    CountryCode: sdkgo.Pointer("1"),
                    AreaCode: sdkgo.Pointer("323"),
                    Number: "111-111-1111",
                    Extension: sdkgo.Pointer("105"),
                    Type: components.PhoneNumberTypePrimary.ToPointer(),
                },
            },
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
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
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
            },
            BankAccounts: []components.BankAccount2{
                components.BankAccount2{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount2AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            TaxCode: sdkgo.Pointer("1111"),
            TaxID: sdkgo.Pointer("234-32-0000"),
            DietaryPreference: sdkgo.Pointer("Veggie"),
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
            RowVersion: sdkgo.Pointer("1-12345"),
            Deleted: sdkgo.Pointer(true),
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

<!-- UsageSnippet language="go" operationID="hris.employeesDelete" method="delete" path="/hris/employees/{id}" -->
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

    res, err := s.Hris.Employees.Delete(ctx, operations.HrisEmployeesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
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