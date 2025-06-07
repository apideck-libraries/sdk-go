# Companies
(*Crm.Companies*)

## Overview

### Available Operations

* [List](#list) - List companies
* [Create](#create) - Create company
* [Get](#get) - Get company
* [Update](#update) - Update company
* [Delete](#delete) - Delete company

## List

List companies

### Example Usage

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

    res, err := s.Crm.Companies.List(ctx, operations.CrmCompaniesAllRequest{
        ServiceID: sdkgo.String("salesforce"),
        Filter: &components.CompaniesFilter{
            Name: sdkgo.String("SpaceX"),
        },
        Sort: &components.CompaniesSort{
            By: components.CompaniesSortByCreatedAt.ToPointer(),
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
    if res.GetCompaniesResponse != nil {
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CrmCompaniesAllRequest](../../models/operations/crmcompaniesallrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CrmCompaniesAllResponse](../../models/operations/crmcompaniesallresponse.md), error**

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

Create company

### Example Usage

```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Crm.Companies.Create(ctx, operations.CrmCompaniesAddRequest{
        ServiceID: sdkgo.String("salesforce"),
        Company: components.CompanyInput{
            Name: sdkgo.String("SpaceX"),
            OwnerID: sdkgo.String("12345"),
            Image: sdkgo.String("https://www.spacex.com/static/images/share.jpg"),
            Description: sdkgo.String("Space Exploration Technologies Corp. is an American aerospace manufacturer, space transportation services and communications company headquartered in Hawthorne, California."),
            VatNumber: sdkgo.String("BE0689615164"),
            Currency: components.CurrencyUsd.ToPointer(),
            Status: sdkgo.String("Open"),
            Fax: sdkgo.String("+12129876543"),
            AnnualRevenue: sdkgo.String("+$35m"),
            NumberOfEmployees: sdkgo.String("500-1000"),
            Industry: sdkgo.String("Apparel"),
            Ownership: sdkgo.String("Public"),
            SalesTaxNumber: sdkgo.String("12456EN"),
            PayeeNumber: sdkgo.String("78932EN"),
            AbnOrTfn: sdkgo.String("46 115 614 695"),
            AbnBranch: sdkgo.String("123"),
            Acn: sdkgo.String("XXX XXX XXX"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
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
            Websites: []components.Website{
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
            },
            SocialLinks: []components.SocialLink{
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
            RowType: &components.CompanyRowType{
                ID: sdkgo.String("12345"),
                Name: sdkgo.String("Customer Account"),
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
            ReadOnly: sdkgo.Bool(false),
            Salutation: sdkgo.String("Mr"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
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
    if res.CreateCompanyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CrmCompaniesAddRequest](../../models/operations/crmcompaniesaddrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CrmCompaniesAddResponse](../../models/operations/crmcompaniesaddresponse.md), error**

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

Get company

### Example Usage

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

    res, err := s.Crm.Companies.Get(ctx, operations.CrmCompaniesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Fields: sdkgo.String("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCompanyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CrmCompaniesOneRequest](../../models/operations/crmcompaniesonerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.CrmCompaniesOneResponse](../../models/operations/crmcompaniesoneresponse.md), error**

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

Update company

### Example Usage

```go
package main

import(
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"os"
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

    res, err := s.Crm.Companies.Update(ctx, operations.CrmCompaniesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
        Company: components.CompanyInput{
            Name: sdkgo.String("SpaceX"),
            OwnerID: sdkgo.String("12345"),
            Image: sdkgo.String("https://www.spacex.com/static/images/share.jpg"),
            Description: sdkgo.String("Space Exploration Technologies Corp. is an American aerospace manufacturer, space transportation services and communications company headquartered in Hawthorne, California."),
            VatNumber: sdkgo.String("BE0689615164"),
            Currency: components.CurrencyUsd.ToPointer(),
            Status: sdkgo.String("Open"),
            Fax: sdkgo.String("+12129876543"),
            AnnualRevenue: sdkgo.String("+$35m"),
            NumberOfEmployees: sdkgo.String("500-1000"),
            Industry: sdkgo.String("Apparel"),
            Ownership: sdkgo.String("Public"),
            SalesTaxNumber: sdkgo.String("12456EN"),
            PayeeNumber: sdkgo.String("78932EN"),
            AbnOrTfn: sdkgo.String("46 115 614 695"),
            AbnBranch: sdkgo.String("123"),
            Acn: sdkgo.String("XXX XXX XXX"),
            FirstName: sdkgo.String("Elon"),
            LastName: sdkgo.String("Musk"),
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
            RowType: &components.CompanyRowType{
                ID: sdkgo.String("12345"),
                Name: sdkgo.String("Customer Account"),
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
            ReadOnly: sdkgo.Bool(false),
            Salutation: sdkgo.String("Mr"),
            Birthday: types.MustNewDateFromString("2000-08-12"),
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
    if res.UpdateCompanyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CrmCompaniesUpdateRequest](../../models/operations/crmcompaniesupdaterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CrmCompaniesUpdateResponse](../../models/operations/crmcompaniesupdateresponse.md), error**

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

Delete company

### Example Usage

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

    res, err := s.Crm.Companies.Delete(ctx, operations.CrmCompaniesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.String("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCompanyResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CrmCompaniesDeleteRequest](../../models/operations/crmcompaniesdeleterequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CrmCompaniesDeleteResponse](../../models/operations/crmcompaniesdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |