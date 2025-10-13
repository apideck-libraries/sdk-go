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

<!-- UsageSnippet language="go" operationID="crm.companiesAll" method="get" path="/crm/companies" -->
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

    res, err := s.Crm.Companies.List(ctx, operations.CrmCompaniesAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Filter: &components.CompaniesFilter{
            Name: sdkgo.Pointer("SpaceX"),
        },
        Sort: &components.CompaniesSort{
            By: components.CompaniesSortByCreatedAt.ToPointer(),
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

<!-- UsageSnippet language="go" operationID="crm.companiesAdd" method="post" path="/crm/companies" -->
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

    res, err := s.Crm.Companies.Create(ctx, operations.CrmCompaniesAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        Company1: components.Company1Input{
            Name: sdkgo.Pointer("SpaceX"),
            OwnerID: sdkgo.Pointer("12345"),
            Image: sdkgo.Pointer("https://www.spacex.com/static/images/share.jpg"),
            Description: sdkgo.Pointer("Space Exploration Technologies Corp. is an American aerospace manufacturer, space transportation services and communications company headquartered in Hawthorne, California."),
            VatNumber: sdkgo.Pointer("BE0689615164"),
            Currency: components.CurrencyUsd.ToPointer(),
            Status: sdkgo.Pointer("Open"),
            Fax: sdkgo.Pointer("+12129876543"),
            AnnualRevenue: sdkgo.Pointer("+$35m"),
            NumberOfEmployees: sdkgo.Pointer("500-1000"),
            Industry: sdkgo.Pointer("Apparel"),
            Ownership: sdkgo.Pointer("Public"),
            SalesTaxNumber: sdkgo.Pointer("12456EN"),
            PayeeNumber: sdkgo.Pointer("78932EN"),
            AbnOrTfn: sdkgo.Pointer("46 115 614 695"),
            AbnBranch: sdkgo.Pointer("123"),
            Acn: sdkgo.Pointer("XXX XXX XXX"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            BankAccounts: []components.BankAccount1{
                components.BankAccount1{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount1AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount1{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount1AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount1{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount1AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            Websites: []components.Website{
                components.Website{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
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
            },
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
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
            RowType: &components.CompanyRowType{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Customer Account"),
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
            Tags: []string{
                "New",
            },
            ReadOnly: sdkgo.Pointer(false),
            Salutation: sdkgo.Pointer("Mr"),
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

<!-- UsageSnippet language="go" operationID="crm.companiesOne" method="get" path="/crm/companies/{id}" -->
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

    res, err := s.Crm.Companies.Get(ctx, operations.CrmCompaniesOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Fields: sdkgo.Pointer("id,updated_at"),
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

<!-- UsageSnippet language="go" operationID="crm.companiesUpdate" method="patch" path="/crm/companies/{id}" -->
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

    res, err := s.Crm.Companies.Update(ctx, operations.CrmCompaniesUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Company1: components.Company1Input{
            Name: sdkgo.Pointer("SpaceX"),
            OwnerID: sdkgo.Pointer("12345"),
            Image: sdkgo.Pointer("https://www.spacex.com/static/images/share.jpg"),
            Description: sdkgo.Pointer("Space Exploration Technologies Corp. is an American aerospace manufacturer, space transportation services and communications company headquartered in Hawthorne, California."),
            VatNumber: sdkgo.Pointer("BE0689615164"),
            Currency: components.CurrencyUsd.ToPointer(),
            Status: sdkgo.Pointer("Open"),
            Fax: sdkgo.Pointer("+12129876543"),
            AnnualRevenue: sdkgo.Pointer("+$35m"),
            NumberOfEmployees: sdkgo.Pointer("500-1000"),
            Industry: sdkgo.Pointer("Apparel"),
            Ownership: sdkgo.Pointer("Public"),
            SalesTaxNumber: sdkgo.Pointer("12456EN"),
            PayeeNumber: sdkgo.Pointer("78932EN"),
            AbnOrTfn: sdkgo.Pointer("46 115 614 695"),
            AbnBranch: sdkgo.Pointer("123"),
            Acn: sdkgo.Pointer("XXX XXX XXX"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            BankAccounts: []components.BankAccount1{
                components.BankAccount1{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.BankAccount1AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            Websites: []components.Website{
                components.Website{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
                components.Website{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
                },
                components.Website{
                    ID: sdkgo.Pointer("12345"),
                    URL: "http://example.com",
                    Type: components.WebsiteTypePrimary.ToPointer(),
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
            SocialLinks: []components.SocialLink{
                components.SocialLink{
                    ID: sdkgo.Pointer("12345"),
                    URL: "https://www.twitter.com/apideck",
                    Type: sdkgo.Pointer("twitter"),
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
            RowType: &components.CompanyRowType{
                ID: sdkgo.Pointer("12345"),
                Name: sdkgo.Pointer("Customer Account"),
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
            Tags: []string{
                "New",
            },
            ReadOnly: sdkgo.Pointer(false),
            Salutation: sdkgo.Pointer("Mr"),
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

<!-- UsageSnippet language="go" operationID="crm.companiesDelete" method="delete" path="/crm/companies/{id}" -->
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

    res, err := s.Crm.Companies.Delete(ctx, operations.CrmCompaniesDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
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