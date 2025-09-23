# Applicants
(*Ats.Applicants*)

## Overview

### Available Operations

* [List](#list) - List Applicants
* [Create](#create) - Create Applicant
* [Get](#get) - Get Applicant
* [Update](#update) - Update Applicant
* [Delete](#delete) - Delete Applicant

## List

List Applicants

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.applicantsAll" method="get" path="/ats/applicants" -->
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

    res, err := s.Ats.Applicants.List(ctx, operations.AtsApplicantsAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.ApplicantsFilter{
            JobID: sdkgo.Pointer("1234"),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicantsResponse != nil {
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
| `request`                                                                                | [operations.AtsApplicantsAllRequest](../../models/operations/atsapplicantsallrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.AtsApplicantsAllResponse](../../models/operations/atsapplicantsallresponse.md), error**

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

Create Applicant

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.applicantsAdd" method="post" path="/ats/applicants" -->
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

    res, err := s.Ats.Applicants.Create(ctx, operations.AtsApplicantsAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Applicant: components.ApplicantInput{
            Name: sdkgo.Pointer("Elon Musk"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            MiddleName: sdkgo.Pointer("D."),
            Initials: sdkgo.Pointer("EM"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            CoverLetter: sdkgo.Pointer("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            PhotoURL: sdkgo.Pointer("https://unavatar.io/elon-musk"),
            Headline: sdkgo.Pointer("PepsiCo, Inc, Central Perk"),
            Title: sdkgo.Pointer("CEO"),
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
            },
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
            Websites: []components.Websites{
                components.Websites{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.ApplicantTypePrimary.ToPointer(),
                },
                components.Websites{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.ApplicantTypePrimary.ToPointer(),
                },
            },
            SocialLinks: []components.SocialLinks{
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
            },
            StageID: sdkgo.Pointer("12345"),
            RecruiterID: sdkgo.Pointer("12345"),
            CoordinatorID: sdkgo.Pointer("12345"),
            ApplicationIds: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Applications: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Sources: []string{
                "Job site",
            },
            Confidential: sdkgo.Pointer(false),
            Anonymized: sdkgo.Pointer(true),
            Tags: []string{
                "New",
            },
            Archived: sdkgo.Pointer(false),
            OwnerID: sdkgo.Pointer("54321"),
            RecordURL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
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
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AtsApplicantsAddRequest](../../models/operations/atsapplicantsaddrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.AtsApplicantsAddResponse](../../models/operations/atsapplicantsaddresponse.md), error**

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

Get Applicant

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.applicantsOne" method="get" path="/ats/applicants/{id}" -->
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

    res, err := s.Ats.Applicants.Get(ctx, operations.AtsApplicantsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.AtsApplicantsOneRequest](../../models/operations/atsapplicantsonerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.AtsApplicantsOneResponse](../../models/operations/atsapplicantsoneresponse.md), error**

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

Update Applicant

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.applicantsUpdate" method="patch" path="/ats/applicants/{id}" -->
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

    res, err := s.Ats.Applicants.Update(ctx, operations.AtsApplicantsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Applicant: components.ApplicantInput{
            Name: sdkgo.Pointer("Elon Musk"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            MiddleName: sdkgo.Pointer("D."),
            Initials: sdkgo.Pointer("EM"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
            CoverLetter: sdkgo.Pointer("I submit this application to express my sincere interest in the API developer position. In the previous role, I was responsible for leadership and ..."),
            PhotoURL: sdkgo.Pointer("https://unavatar.io/elon-musk"),
            Headline: sdkgo.Pointer("PepsiCo, Inc, Central Perk"),
            Title: sdkgo.Pointer("CEO"),
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
                },
            },
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.Pointer("2389328923893298"),
                    Name: sdkgo.Pointer("employee_level"),
                    Description: sdkgo.Pointer("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
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
            Websites: []components.Websites{
                components.Websites{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.ApplicantTypePrimary.ToPointer(),
                },
                components.Websites{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.ApplicantTypePrimary.ToPointer(),
                },
            },
            SocialLinks: []components.SocialLinks{
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
                components.SocialLinks{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
                },
            },
            StageID: sdkgo.Pointer("12345"),
            RecruiterID: sdkgo.Pointer("12345"),
            CoordinatorID: sdkgo.Pointer("12345"),
            ApplicationIds: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Applications: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Followers: []string{
                "a0d636c6-43b3-4bde-8c70-85b707d992f4",
                "a98lfd96-43b3-4bde-8c70-85b707d992e6",
            },
            Sources: []string{
                "Job site",
            },
            Confidential: sdkgo.Pointer(false),
            Anonymized: sdkgo.Pointer(true),
            Tags: []string{
                "New",
            },
            Archived: sdkgo.Pointer(false),
            OwnerID: sdkgo.Pointer("54321"),
            RecordURL: sdkgo.Pointer("https://app.intercom.io/contacts/12345"),
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
    if res.UpdateApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AtsApplicantsUpdateRequest](../../models/operations/atsapplicantsupdaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AtsApplicantsUpdateResponse](../../models/operations/atsapplicantsupdateresponse.md), error**

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

Delete Applicant

### Example Usage

<!-- UsageSnippet language="go" operationID="ats.applicantsDelete" method="delete" path="/ats/applicants/{id}" -->
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

    res, err := s.Ats.Applicants.Delete(ctx, operations.AtsApplicantsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteApplicantResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.AtsApplicantsDeleteRequest](../../models/operations/atsapplicantsdeleterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.AtsApplicantsDeleteResponse](../../models/operations/atsapplicantsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |