# Leads
(*Crm.Leads*)

## Overview

### Available Operations

* [List](#list) - List leads
* [Create](#create) - Create lead
* [Get](#get) - Get lead
* [Update](#update) - Update lead
* [Delete](#delete) - Delete lead

## List

List leads

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

    res, err := s.Crm.Leads.List(ctx, operations.CrmLeadsAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.LeadsFilter{
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            Email: sdkgo.String("elon@tesla.com"),
            PhoneNumber: sdkgo.String("1234567890"),
        },
        Sort: &components.LeadsSort{
            By: components.LeadsSortByCreatedAt.ToPointer(),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLeadsResponse != nil {
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
| `request`                                                                      | [operations.CrmLeadsAllRequest](../../models/operations/crmleadsallrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmLeadsAllResponse](../../models/operations/crmleadsallresponse.md), error**

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

Create lead

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

    res, err := s.Crm.Leads.Create(ctx, operations.CrmLeadsAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        Lead: components.LeadInput{
            Name: "Elon Musk",
            CompanyName: sdkgo.String("Spacex"),
            OwnerID: sdkgo.String("54321"),
            OwnerName: sdkgo.String("John Doe"),
            CompanyID: sdkgo.String("2"),
            LeadID: sdkgo.String("2"),
            LeadSource: sdkgo.String("Cold Call"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            Description: sdkgo.String("A thinker"),
            Prefix: sdkgo.String("Sir"),
            Title: sdkgo.String("CEO"),
            Language: sdkgo.String("EN"),
            Status: sdkgo.String("New"),
            MonetaryAmount: sdkgo.Float64(75000),
            Currency: components.CurrencyUsd.ToPointer(),
            Fax: sdkgo.String("+12129876543"),
            Websites: []components.Website{
                components.Website{
                    ID: sdkgo.String("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
                components.Website{
                    ID: sdkgo.String("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
            },
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
            PhoneNumbers: []components.PhoneNumber{
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
            CustomFields: []components.CustomField{
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
            Tags: []string{
                "New",
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
    if res.CreateLeadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CrmLeadsAddRequest](../../models/operations/crmleadsaddrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmLeadsAddResponse](../../models/operations/crmleadsaddresponse.md), error**

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

Get lead

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

    res, err := s.Crm.Leads.Get(ctx, operations.CrmLeadsOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetLeadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CrmLeadsOneRequest](../../models/operations/crmleadsonerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CrmLeadsOneResponse](../../models/operations/crmleadsoneresponse.md), error**

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

Update lead

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

    res, err := s.Crm.Leads.Update(ctx, operations.CrmLeadsUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Lead: components.LeadInput{
            Name: "Elon Musk",
            CompanyName: sdkgo.String("Spacex"),
            OwnerID: sdkgo.String("54321"),
            OwnerName: sdkgo.String("John Doe"),
            CompanyID: sdkgo.String("2"),
            LeadID: sdkgo.String("2"),
            LeadSource: sdkgo.String("Cold Call"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
            Description: sdkgo.String("A thinker"),
            Prefix: sdkgo.String("Sir"),
            Title: sdkgo.String("CEO"),
            Language: sdkgo.String("EN"),
            Status: sdkgo.String("New"),
            MonetaryAmount: sdkgo.Float64(75000),
            Currency: components.CurrencyUsd.ToPointer(),
            Fax: sdkgo.String("+12129876543"),
            Websites: []components.Website{
                components.Website{
                    ID: sdkgo.String("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
                components.Website{
                    ID: sdkgo.String("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
                components.Website{
                    ID: sdkgo.String("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
            },
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
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
                components.CustomField{
                    ID: sdkgo.String("2389328923893298"),
                    Name: sdkgo.String("employee_level"),
                    Description: sdkgo.String("Employee Level"),
                    Value: sdkgo.Pointer(components.CreateValueStr(
                        "Uses Salesforce and Marketo",
                    )),
                },
            },
            Tags: []string{
                "New",
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
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateLeadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CrmLeadsUpdateRequest](../../models/operations/crmleadsupdaterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CrmLeadsUpdateResponse](../../models/operations/crmleadsupdateresponse.md), error**

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

Delete lead

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

    res, err := s.Crm.Leads.Delete(ctx, operations.CrmLeadsDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteLeadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CrmLeadsDeleteRequest](../../models/operations/crmleadsdeleterequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.CrmLeadsDeleteResponse](../../models/operations/crmleadsdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |