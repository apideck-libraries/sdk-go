# Accounting.Customers

## Overview

### Available Operations

* [List](#list) - List Customers
* [Create](#create) - Create Customer
* [Get](#get) - Get Customer
* [Update](#update) - Update Customer
* [Delete](#delete) - Delete Customer

## List

List Customers

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.customersAll" method="get" path="/accounting/customers" -->
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

    res, err := s.Accounting.Customers.List(ctx, operations.AccountingCustomersAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.CustomersFilter{
            CompanyName: sdkgo.Pointer("SpaceX"),
            DisplayName: sdkgo.Pointer("Elon Musk"),
            FirstName: sdkgo.Pointer("Elon"),
            LastName: sdkgo.Pointer("Musk"),
            Email: sdkgo.Pointer("elon@musk.com"),
            Status: components.CustomersFilterStatusActive.ToPointer(),
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            SupplierID: sdkgo.Pointer("123"),
        },
        Sort: &components.CustomersSort{
            By: components.CustomersSortByUpdatedAt.ToPointer(),
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
    if res.GetCustomersResponse != nil {
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountingCustomersAllRequest](../../models/operations/accountingcustomersallrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AccountingCustomersAllResponse](../../models/operations/accountingcustomersallresponse.md), error**

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

Create Customer

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.customersAdd" method="post" path="/accounting/customers" -->
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

    res, err := s.Accounting.Customers.Create(ctx, operations.AccountingCustomersAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Customer: components.CustomerInput{
            DisplayID: sdkgo.Pointer("EMP00101"),
            DisplayName: sdkgo.Pointer("Windsurf Shop"),
            CompanyName: sdkgo.Pointer("SpaceX"),
            CompanyID: sdkgo.Pointer("12345"),
            Title: sdkgo.Pointer("CEO"),
            FirstName: sdkgo.Pointer("Elon"),
            MiddleName: sdkgo.Pointer("D."),
            LastName: sdkgo.Pointer("Musk"),
            Suffix: sdkgo.Pointer("Jr."),
            Individual: sdkgo.Pointer(true),
            Project: sdkgo.Pointer(false),
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
            },
            Emails: []components.Email{
                components.Email{
                    ID: sdkgo.Pointer("123"),
                    Email: sdkgo.Pointer("elon@musk.com"),
                    Type: components.EmailTypePrimary.ToPointer(),
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
            BankAccounts: []components.BankAccount{
                components.BankAccount{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            Notes: sdkgo.Pointer("Some notes about this customer"),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Rate: sdkgo.Pointer[float64](10),
            },
            TaxNumber: sdkgo.Pointer("US123945459"),
            Currency: components.CurrencyUsd.ToPointer(),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            Parent: &components.LinkedParentCustomer{
                ID: "12345",
                Name: sdkgo.Pointer("Windsurf Shop"),
            },
            Status: components.CustomerStatusStatusActive.ToPointer(),
            PaymentMethod: sdkgo.Pointer("cash"),
            Channel: sdkgo.Pointer("email"),
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: nil,
                    },
                ),
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: nil,
                    },
                ),
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        Description: sdkgo.Pointer("Employee Level"),
                        Value: nil,
                    },
                ),
            },
            RowVersion: sdkgo.Pointer("1-12345"),
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
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountingCustomersAddRequest](../../models/operations/accountingcustomersaddrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AccountingCustomersAddResponse](../../models/operations/accountingcustomersaddresponse.md), error**

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

Get Customer

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.customersOne" method="get" path="/accounting/customers/{id}" -->
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

    res, err := s.Accounting.Customers.Get(ctx, operations.AccountingCustomersOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Fields: sdkgo.Pointer("id,updated_at"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.AccountingCustomersOneRequest](../../models/operations/accountingcustomersonerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.AccountingCustomersOneResponse](../../models/operations/accountingcustomersoneresponse.md), error**

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

Update Customer

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.customersUpdate" method="patch" path="/accounting/customers/{id}" -->
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

    res, err := s.Accounting.Customers.Update(ctx, operations.AccountingCustomersUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        Customer: components.CustomerInput{
            DisplayID: sdkgo.Pointer("EMP00101"),
            DisplayName: sdkgo.Pointer("Windsurf Shop"),
            CompanyName: sdkgo.Pointer("SpaceX"),
            CompanyID: sdkgo.Pointer("12345"),
            Title: sdkgo.Pointer("CEO"),
            FirstName: sdkgo.Pointer("Elon"),
            MiddleName: sdkgo.Pointer("D."),
            LastName: sdkgo.Pointer("Musk"),
            Suffix: sdkgo.Pointer("Jr."),
            Individual: sdkgo.Pointer(true),
            Project: sdkgo.Pointer(false),
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
            BankAccounts: []components.BankAccount{
                components.BankAccount{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
                components.BankAccount{
                    BankName: sdkgo.Pointer("Monzo"),
                    AccountNumber: sdkgo.Pointer("123465"),
                    AccountName: sdkgo.Pointer("SPACEX LLC"),
                    AccountType: components.AccountTypeCreditCard.ToPointer(),
                    Iban: sdkgo.Pointer("CH2989144532982975332"),
                    Bic: sdkgo.Pointer("AUDSCHGGXXX"),
                    RoutingNumber: sdkgo.Pointer("012345678"),
                    BsbNumber: sdkgo.Pointer("062-001"),
                    BranchIdentifier: sdkgo.Pointer("001"),
                    BankCode: sdkgo.Pointer("BNH"),
                    Currency: components.CurrencyUsd.ToPointer(),
                },
            },
            Notes: sdkgo.Pointer("Some notes about this customer"),
            TaxRate: &components.LinkedTaxRateInput{
                ID: sdkgo.Pointer("123456"),
                Rate: sdkgo.Pointer[float64](10),
            },
            TaxNumber: sdkgo.Pointer("US123945459"),
            Currency: components.CurrencyUsd.ToPointer(),
            Account: &components.LinkedLedgerAccount{
                ID: sdkgo.Pointer("123456"),
                NominalCode: sdkgo.Pointer("N091"),
                Code: sdkgo.Pointer("453"),
            },
            Parent: &components.LinkedParentCustomer{
                ID: "12345",
                Name: sdkgo.Pointer("Windsurf Shop"),
            },
            Status: components.CustomerStatusStatusActive.ToPointer(),
            PaymentMethod: sdkgo.Pointer("cash"),
            Channel: sdkgo.Pointer("email"),
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
    if res.UpdateCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingCustomersUpdateRequest](../../models/operations/accountingcustomersupdaterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingCustomersUpdateResponse](../../models/operations/accountingcustomersupdateresponse.md), error**

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

Delete Customer

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.customersDelete" method="delete" path="/accounting/customers/{id}" -->
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

    res, err := s.Accounting.Customers.Delete(ctx, operations.AccountingCustomersDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteCustomerResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.AccountingCustomersDeleteRequest](../../models/operations/accountingcustomersdeleterequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |
| `opts`                                                                                                     | [][operations.Option](../../models/operations/option.md)                                                   | :heavy_minus_sign:                                                                                         | The options for this request.                                                                              |

### Response

**[*operations.AccountingCustomersDeleteResponse](../../models/operations/accountingcustomersdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |