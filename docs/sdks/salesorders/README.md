# Accounting.SalesOrders

## Overview

### Available Operations

* [List](#list) - List Sales Orders
* [Create](#create) - Create Sales Order
* [Get](#get) - Get Sales Order
* [Update](#update) - Update Sales Order
* [Delete](#delete) - Delete Sales Order

## List

List Sales Orders

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesOrdersAll" method="get" path="/accounting/sales-orders" -->
```go
package main

import(
	"context"
	"os"
	sdkgo "github.com/apideck-libraries/sdk-go"
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

    res, err := s.Accounting.SalesOrders.List(ctx, operations.AccountingSalesOrdersAllRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        Filter: &components.SalesOrdersFilter{
            UpdatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            CreatedSince: types.MustNewTimeFromString("2020-09-30T07:43:32.000Z"),
            Number: sdkgo.Pointer("SO000123"),
            CustomerID: sdkgo.Pointer("123abc"),
        },
        PassThrough: map[string]any{
            "search": "San Francisco",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSalesOrdersResponse != nil {
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

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingSalesOrdersAllRequest](../../models/operations/accountingsalesordersallrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingSalesOrdersAllResponse](../../models/operations/accountingsalesordersallresponse.md), error**

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

Create Sales Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesOrdersAdd" method="post" path="/accounting/sales-orders" -->
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

    res, err := s.Accounting.SalesOrders.Create(ctx, operations.AccountingSalesOrdersAddRequest{
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        SalesOrder: components.SalesOrderInput{
            Number: sdkgo.Pointer("SO000123"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            QuoteID: sdkgo.Pointer("123456"),
            CompanyID: sdkgo.Pointer("12345"),
            DepartmentID: sdkgo.Pointer("12345"),
            LocationID: sdkgo.Pointer("12345"),
            ProjectID: sdkgo.Pointer("12345"),
            SubsidiaryID: sdkgo.Pointer("12345"),
            OrderDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-10-15"),
            OrderType: sdkgo.Pointer("SO"),
            Terms: sdkgo.Pointer("Net 30"),
            TermsID: sdkgo.Pointer("12345"),
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("INV-2024-001"),
            Status: components.SalesOrderStatusOpen.ToPointer(),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500.0),
            TotalTax: sdkgo.Pointer[float64](2500.0),
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Total: sdkgo.Pointer[float64](30000.0),
            ShippingMethod: sdkgo.Pointer("FEDEX"),
            PaymentMethod: sdkgo.Pointer("cash"),
            CustomerMemo: sdkgo.Pointer("Thank you for your order!"),
            Notes: sdkgo.Pointer("Ship with next consignment"),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500.0),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    CategoryID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    SubsidiaryID: sdkgo.Pointer("12345"),
                    ShippingID: sdkgo.Pointer("12345"),
                    Memo: sdkgo.Pointer("Some memo"),
                    Prepaid: sdkgo.Pointer(true),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxApplicableOn: sdkgo.Pointer("Domestic_Purchase_of_Goods_and_Services"),
                    TaxRecoverability: sdkgo.Pointer("Fully_Recoverable"),
                    TaxMethod: sdkgo.Pointer("Due_to_Supplier"),
                    Worktags: []*components.LinkedWorktag{
                        &components.LinkedWorktag{
                            ID: sdkgo.Pointer("123456"),
                            Value: sdkgo.Pointer("New York"),
                        },
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Code: sdkgo.Pointer("100"),
                            Name: sdkgo.Pointer("New York"),
                            ParentID: sdkgo.Pointer("123456"),
                            ParentName: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                    },
                    CustomFields: []components.CustomField{
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                RefName: sdkgo.Pointer("Marketing"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                                    "Uses Salesforce and Marketo",
                                )),
                            },
                        ),
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            BillingAddress: &components.Address{
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
            ShippingAddress: &components.Address{
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
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Code: sdkgo.Pointer("100"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            TemplateID: sdkgo.Pointer("123456"),
            SourceDocumentURL: sdkgo.Pointer("https://www.ordersolution.com/order/123456"),
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        RefName: sdkgo.Pointer("Marketing"),
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateSalesOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingSalesOrdersAddRequest](../../models/operations/accountingsalesordersaddrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingSalesOrdersAddResponse](../../models/operations/accountingsalesordersaddresponse.md), error**

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

Get Sales Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesOrdersOne" method="get" path="/accounting/sales-orders/{id}" -->
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

    res, err := s.Accounting.SalesOrders.Get(ctx, operations.AccountingSalesOrdersOneRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.GetSalesOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.AccountingSalesOrdersOneRequest](../../models/operations/accountingsalesordersonerequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.AccountingSalesOrdersOneResponse](../../models/operations/accountingsalesordersoneresponse.md), error**

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

Update Sales Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesOrdersUpdate" method="patch" path="/accounting/sales-orders/{id}" -->
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

    res, err := s.Accounting.SalesOrders.Update(ctx, operations.AccountingSalesOrdersUpdateRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
        SalesOrder: components.SalesOrderInput{
            Number: sdkgo.Pointer("SO000123"),
            Customer: &components.LinkedCustomerInput{
                ID: sdkgo.Pointer("12345"),
                DisplayName: sdkgo.Pointer("Windsurf Shop"),
                Email: sdkgo.Pointer("boring@boring.com"),
            },
            QuoteID: sdkgo.Pointer("123456"),
            CompanyID: sdkgo.Pointer("12345"),
            DepartmentID: sdkgo.Pointer("12345"),
            LocationID: sdkgo.Pointer("12345"),
            ProjectID: sdkgo.Pointer("12345"),
            SubsidiaryID: sdkgo.Pointer("12345"),
            OrderDate: types.MustNewDateFromString("2020-09-30"),
            DeliveryDate: types.MustNewDateFromString("2020-10-15"),
            OrderType: sdkgo.Pointer("SO"),
            Terms: sdkgo.Pointer("Net 30"),
            TermsID: sdkgo.Pointer("12345"),
            PoNumber: sdkgo.Pointer("90000117"),
            Reference: sdkgo.Pointer("INV-2024-001"),
            Status: components.SalesOrderStatusOpen.ToPointer(),
            Currency: components.CurrencyUsd.ToPointer(),
            CurrencyRate: sdkgo.Pointer[float64](0.69),
            TaxInclusive: sdkgo.Pointer(true),
            SubTotal: sdkgo.Pointer[float64](27500.0),
            TotalTax: sdkgo.Pointer[float64](2500.0),
            TaxCode: sdkgo.Pointer("1234"),
            DiscountPercentage: sdkgo.Pointer[float64](5.5),
            DiscountAmount: sdkgo.Pointer[float64](25.0),
            Total: sdkgo.Pointer[float64](30000.0),
            ShippingMethod: sdkgo.Pointer("FEDEX"),
            PaymentMethod: sdkgo.Pointer("cash"),
            CustomerMemo: sdkgo.Pointer("Thank you for your order!"),
            Notes: sdkgo.Pointer("Ship with next consignment"),
            LineItems: []components.InvoiceLineItemInput{
                components.InvoiceLineItemInput{
                    ID: sdkgo.Pointer("12345"),
                    RowID: sdkgo.Pointer("12345"),
                    Code: sdkgo.Pointer("120-C"),
                    LineNumber: sdkgo.Pointer[int64](1),
                    Description: sdkgo.Pointer("Model Y is a fully electric, mid-size SUV, with seating for up to seven, dual motor AWD and unparalleled protection."),
                    Type: components.InvoiceLineItemTypeSalesItem.ToPointer(),
                    TaxAmount: sdkgo.Pointer[float64](27500.0),
                    TotalAmount: sdkgo.Pointer[float64](27500.0),
                    Quantity: sdkgo.Pointer[float64](1.0),
                    UnitPrice: sdkgo.Pointer[float64](27500.5),
                    UnitOfMeasure: sdkgo.Pointer("pc."),
                    DiscountPercentage: sdkgo.Pointer[float64](0.01),
                    DiscountAmount: sdkgo.Pointer[float64](19.99),
                    ServiceDate: types.MustNewDateFromString("2024-01-15"),
                    CategoryID: sdkgo.Pointer("12345"),
                    LocationID: sdkgo.Pointer("12345"),
                    DepartmentID: sdkgo.Pointer("12345"),
                    SubsidiaryID: sdkgo.Pointer("12345"),
                    ShippingID: sdkgo.Pointer("12345"),
                    Memo: sdkgo.Pointer("Some memo"),
                    Prepaid: sdkgo.Pointer(true),
                    Item: &components.LinkedInvoiceItem{
                        ID: sdkgo.Pointer("12344"),
                        Code: sdkgo.Pointer("120-C"),
                        Name: sdkgo.Pointer("Model Y"),
                    },
                    TaxApplicableOn: sdkgo.Pointer("Domestic_Purchase_of_Goods_and_Services"),
                    TaxRecoverability: sdkgo.Pointer("Fully_Recoverable"),
                    TaxMethod: sdkgo.Pointer("Due_to_Supplier"),
                    Worktags: []*components.LinkedWorktag{
                        &components.LinkedWorktag{
                            ID: sdkgo.Pointer("123456"),
                            Value: sdkgo.Pointer("New York"),
                        },
                    },
                    TaxRate: &components.LinkedTaxRateInput{
                        ID: sdkgo.Pointer("123456"),
                        Code: sdkgo.Pointer("N-T"),
                        Rate: sdkgo.Pointer[float64](10.0),
                    },
                    TrackingCategories: []*components.LinkedTrackingCategory{
                        &components.LinkedTrackingCategory{
                            ID: sdkgo.Pointer("123456"),
                            Code: sdkgo.Pointer("100"),
                            Name: sdkgo.Pointer("New York"),
                            ParentID: sdkgo.Pointer("123456"),
                            ParentName: sdkgo.Pointer("New York"),
                        },
                    },
                    LedgerAccount: &components.LinkedLedgerAccount{
                        ID: sdkgo.Pointer("123456"),
                        Name: sdkgo.Pointer("Bank account"),
                        NominalCode: sdkgo.Pointer("N091"),
                        Code: sdkgo.Pointer("453"),
                        ParentID: sdkgo.Pointer("123456"),
                        DisplayID: sdkgo.Pointer("123456"),
                    },
                    CustomFields: []components.CustomField{
                        components.CreateCustomFieldCustomField1(
                            components.CustomField1{
                                ID: sdkgo.Pointer("2389328923893298"),
                                Name: sdkgo.Pointer("employee_level"),
                                RefName: sdkgo.Pointer("Marketing"),
                                Description: sdkgo.Pointer("Employee Level"),
                                Value: sdkgo.Pointer(components.CreateCustomField1ValueStr(
                                    "Uses Salesforce and Marketo",
                                )),
                            },
                        ),
                    },
                    RowVersion: sdkgo.Pointer("1-12345"),
                },
            },
            BillingAddress: &components.Address{
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
            ShippingAddress: &components.Address{
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
            TrackingCategories: []*components.LinkedTrackingCategory{
                &components.LinkedTrackingCategory{
                    ID: sdkgo.Pointer("123456"),
                    Code: sdkgo.Pointer("100"),
                    Name: sdkgo.Pointer("New York"),
                    ParentID: sdkgo.Pointer("123456"),
                    ParentName: sdkgo.Pointer("New York"),
                },
            },
            TemplateID: sdkgo.Pointer("123456"),
            SourceDocumentURL: sdkgo.Pointer("https://www.ordersolution.com/order/123456"),
            CustomFields: []components.CustomField{
                components.CreateCustomFieldCustomField1(
                    components.CustomField1{
                        ID: sdkgo.Pointer("2389328923893298"),
                        Name: sdkgo.Pointer("employee_level"),
                        RefName: sdkgo.Pointer("Marketing"),
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
                    },
                },
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.UpdateSalesOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingSalesOrdersUpdateRequest](../../models/operations/accountingsalesordersupdaterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingSalesOrdersUpdateResponse](../../models/operations/accountingsalesordersupdateresponse.md), error**

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

Delete Sales Order

### Example Usage

<!-- UsageSnippet language="go" operationID="accounting.salesOrdersDelete" method="delete" path="/accounting/sales-orders/{id}" -->
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

    res, err := s.Accounting.SalesOrders.Delete(ctx, operations.AccountingSalesOrdersDeleteRequest{
        ID: "<id>",
        ServiceID: sdkgo.Pointer("salesforce"),
        CompanyID: sdkgo.Pointer("12345"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.DeleteSalesOrderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.AccountingSalesOrdersDeleteRequest](../../models/operations/accountingsalesordersdeleterequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.AccountingSalesOrdersDeleteResponse](../../models/operations/accountingsalesordersdeleteresponse.md), error**

### Errors

| Error Type                        | Status Code                       | Content Type                      |
| --------------------------------- | --------------------------------- | --------------------------------- |
| apierrors.BadRequestResponse      | 400                               | application/json                  |
| apierrors.UnauthorizedResponse    | 401                               | application/json                  |
| apierrors.PaymentRequiredResponse | 402                               | application/json                  |
| apierrors.NotFoundResponse        | 404                               | application/json                  |
| apierrors.UnprocessableResponse   | 422                               | application/json                  |
| apierrors.APIError                | 4XX, 5XX                          | \*/\*                             |