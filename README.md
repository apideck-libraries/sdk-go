# github.com/apideck-libraries/sdk-go

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/apideck-libraries/sdk-go* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/apideck-libraries/sdk-go&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<!-- Start Summary [summary] -->
## Summary

Apideck: The Apideck OpenAPI Spec: SDK Optimized

For more information about the API: [Apideck Developer Docs](https://developers.apideck.com)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/apideck-libraries/sdk-go](#githubcomapideck-librariessdk-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/apideck-libraries/sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type | Scheme      | Environment Variable |
| -------- | ---- | ----------- | -------------------- |
| `APIKey` | http | HTTP Bearer | `APIDECK_API_KEY`    |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Accounting](docs/sdks/accounting/README.md)


#### [Accounting.AgedDebtors](docs/sdks/ageddebtors/README.md)

* [Get](docs/sdks/ageddebtors/README.md#get) - Get Aged Debtors

#### [Accounting.Attachments](docs/sdks/attachments/README.md)

* [List](docs/sdks/attachments/README.md#list) - List Attachments
* [Get](docs/sdks/attachments/README.md#get) - Get Attachment
* [Delete](docs/sdks/attachments/README.md#delete) - Delete Attachment
* [Download](docs/sdks/attachments/README.md#download) - Download Attachment

#### [Accounting.BalanceSheet](docs/sdks/balancesheet/README.md)

* [Get](docs/sdks/balancesheet/README.md#get) - Get BalanceSheet

#### [Accounting.BillPayments](docs/sdks/billpayments/README.md)

* [List](docs/sdks/billpayments/README.md#list) - List Bill Payments
* [Create](docs/sdks/billpayments/README.md#create) - Create Bill Payment
* [Get](docs/sdks/billpayments/README.md#get) - Get Bill Payment
* [Update](docs/sdks/billpayments/README.md#update) - Update Bill Payment
* [Delete](docs/sdks/billpayments/README.md#delete) - Delete Bill Payment

#### [Accounting.Bills](docs/sdks/bills/README.md)

* [List](docs/sdks/bills/README.md#list) - List Bills
* [Create](docs/sdks/bills/README.md#create) - Create Bill
* [Get](docs/sdks/bills/README.md#get) - Get Bill
* [Update](docs/sdks/bills/README.md#update) - Update Bill
* [Delete](docs/sdks/bills/README.md#delete) - Delete Bill

#### [Accounting.CompanyInfo](docs/sdks/companyinfo/README.md)

* [Get](docs/sdks/companyinfo/README.md#get) - Get company info

#### [Accounting.CreditNotes](docs/sdks/creditnotes/README.md)

* [List](docs/sdks/creditnotes/README.md#list) - List Credit Notes
* [Create](docs/sdks/creditnotes/README.md#create) - Create Credit Note
* [Get](docs/sdks/creditnotes/README.md#get) - Get Credit Note
* [Update](docs/sdks/creditnotes/README.md#update) - Update Credit Note
* [Delete](docs/sdks/creditnotes/README.md#delete) - Delete Credit Note

#### [Accounting.Customers](docs/sdks/customers/README.md)

* [List](docs/sdks/customers/README.md#list) - List Customers
* [Create](docs/sdks/customers/README.md#create) - Create Customer
* [Get](docs/sdks/customers/README.md#get) - Get Customer
* [Update](docs/sdks/customers/README.md#update) - Update Customer
* [Delete](docs/sdks/customers/README.md#delete) - Delete Customer

#### [Accounting.Departments](docs/sdks/departments/README.md)

* [List](docs/sdks/departments/README.md#list) - List Departments
* [Create](docs/sdks/departments/README.md#create) - Create Department
* [Get](docs/sdks/departments/README.md#get) - Get Department
* [Update](docs/sdks/departments/README.md#update) - Update Department
* [Delete](docs/sdks/departments/README.md#delete) - Delete Department

#### [Accounting.Expenses](docs/sdks/expenses/README.md)

* [List](docs/sdks/expenses/README.md#list) - List Expenses
* [Create](docs/sdks/expenses/README.md#create) - Create Expense
* [Get](docs/sdks/expenses/README.md#get) - Get Expense
* [Update](docs/sdks/expenses/README.md#update) - Update Expense
* [Delete](docs/sdks/expenses/README.md#delete) - Delete Expense

#### [Accounting.InvoiceItems](docs/sdks/invoiceitems/README.md)

* [List](docs/sdks/invoiceitems/README.md#list) - List Invoice Items
* [Create](docs/sdks/invoiceitems/README.md#create) - Create Invoice Item
* [Get](docs/sdks/invoiceitems/README.md#get) - Get Invoice Item
* [Update](docs/sdks/invoiceitems/README.md#update) - Update Invoice Item
* [Delete](docs/sdks/invoiceitems/README.md#delete) - Delete Invoice Item

#### [Accounting.Invoices](docs/sdks/invoices/README.md)

* [List](docs/sdks/invoices/README.md#list) - List Invoices
* [Create](docs/sdks/invoices/README.md#create) - Create Invoice
* [Get](docs/sdks/invoices/README.md#get) - Get Invoice
* [Update](docs/sdks/invoices/README.md#update) - Update Invoice
* [Delete](docs/sdks/invoices/README.md#delete) - Delete Invoice

#### [Accounting.JournalEntries](docs/sdks/journalentries/README.md)

* [List](docs/sdks/journalentries/README.md#list) - List Journal Entries
* [Create](docs/sdks/journalentries/README.md#create) - Create Journal Entry
* [Get](docs/sdks/journalentries/README.md#get) - Get Journal Entry
* [Update](docs/sdks/journalentries/README.md#update) - Update Journal Entry
* [Delete](docs/sdks/journalentries/README.md#delete) - Delete Journal Entry

#### [Accounting.LedgerAccounts](docs/sdks/ledgeraccounts/README.md)

* [List](docs/sdks/ledgeraccounts/README.md#list) - List Ledger Accounts
* [Create](docs/sdks/ledgeraccounts/README.md#create) - Create Ledger Account
* [Get](docs/sdks/ledgeraccounts/README.md#get) - Get Ledger Account
* [Update](docs/sdks/ledgeraccounts/README.md#update) - Update Ledger Account
* [Delete](docs/sdks/ledgeraccounts/README.md#delete) - Delete Ledger Account

#### [Accounting.Locations](docs/sdks/locations/README.md)

* [List](docs/sdks/locations/README.md#list) - List Locations
* [Create](docs/sdks/locations/README.md#create) - Create Location
* [Get](docs/sdks/locations/README.md#get) - Get Location
* [Update](docs/sdks/locations/README.md#update) - Update Location
* [Delete](docs/sdks/locations/README.md#delete) - Delete Location

#### [Accounting.Payments](docs/sdks/payments/README.md)

* [List](docs/sdks/payments/README.md#list) - List Payments
* [Create](docs/sdks/payments/README.md#create) - Create Payment
* [Get](docs/sdks/payments/README.md#get) - Get Payment
* [Update](docs/sdks/payments/README.md#update) - Update Payment
* [Delete](docs/sdks/payments/README.md#delete) - Delete Payment

#### [Accounting.ProfitAndLoss](docs/sdks/profitandloss/README.md)

* [Get](docs/sdks/profitandloss/README.md#get) - Get Profit and Loss

#### [Accounting.PurchaseOrders](docs/sdks/purchaseorders/README.md)

* [List](docs/sdks/purchaseorders/README.md#list) - List Purchase Orders
* [Create](docs/sdks/purchaseorders/README.md#create) - Create Purchase Order
* [Get](docs/sdks/purchaseorders/README.md#get) - Get Purchase Order
* [Update](docs/sdks/purchaseorders/README.md#update) - Update Purchase Order
* [Delete](docs/sdks/purchaseorders/README.md#delete) - Delete Purchase Order

#### [Accounting.Subsidiaries](docs/sdks/subsidiaries/README.md)

* [List](docs/sdks/subsidiaries/README.md#list) - List Subsidiaries
* [Create](docs/sdks/subsidiaries/README.md#create) - Create Subsidiary
* [Get](docs/sdks/subsidiaries/README.md#get) - Get Subsidiary
* [Update](docs/sdks/subsidiaries/README.md#update) - Update Subsidiary
* [Delete](docs/sdks/subsidiaries/README.md#delete) - Delete Subsidiary

#### [Accounting.Suppliers](docs/sdks/suppliers/README.md)

* [List](docs/sdks/suppliers/README.md#list) - List Suppliers
* [Create](docs/sdks/suppliers/README.md#create) - Create Supplier
* [Get](docs/sdks/suppliers/README.md#get) - Get Supplier
* [Update](docs/sdks/suppliers/README.md#update) - Update Supplier
* [Delete](docs/sdks/suppliers/README.md#delete) - Delete Supplier

#### [Accounting.TaxRates](docs/sdks/taxrates/README.md)

* [List](docs/sdks/taxrates/README.md#list) - List Tax Rates
* [Create](docs/sdks/taxrates/README.md#create) - Create Tax Rate
* [Get](docs/sdks/taxrates/README.md#get) - Get Tax Rate
* [Update](docs/sdks/taxrates/README.md#update) - Update Tax Rate
* [Delete](docs/sdks/taxrates/README.md#delete) - Delete Tax Rate

#### [Accounting.TrackingCategories](docs/sdks/trackingcategories/README.md)

* [List](docs/sdks/trackingcategories/README.md#list) - List Tracking Categories
* [Create](docs/sdks/trackingcategories/README.md#create) - Create Tracking Category
* [Get](docs/sdks/trackingcategories/README.md#get) - Get Tracking Category
* [Update](docs/sdks/trackingcategories/README.md#update) - Update Tracking Category
* [Delete](docs/sdks/trackingcategories/README.md#delete) - Delete Tracking Category


### [Ats](docs/sdks/ats/README.md)


#### [Ats.Applicants](docs/sdks/applicants/README.md)

* [List](docs/sdks/applicants/README.md#list) - List Applicants
* [Create](docs/sdks/applicants/README.md#create) - Create Applicant
* [Get](docs/sdks/applicants/README.md#get) - Get Applicant
* [Update](docs/sdks/applicants/README.md#update) - Update Applicant
* [Delete](docs/sdks/applicants/README.md#delete) - Delete Applicant

#### [Ats.Applications](docs/sdks/applications/README.md)

* [List](docs/sdks/applications/README.md#list) - List Applications
* [Create](docs/sdks/applications/README.md#create) - Create Application
* [Get](docs/sdks/applications/README.md#get) - Get Application
* [Update](docs/sdks/applications/README.md#update) - Update Application
* [Delete](docs/sdks/applications/README.md#delete) - Delete Application

#### [Ats.Jobs](docs/sdks/jobs/README.md)

* [List](docs/sdks/jobs/README.md#list) - List Jobs
* [Get](docs/sdks/jobs/README.md#get) - Get Job

### [Connector](docs/sdks/connector/README.md)


#### [Connector.APIResourceCoverage](docs/sdks/apiresourcecoverage/README.md)

* [Get](docs/sdks/apiresourcecoverage/README.md#get) - Get API Resource Coverage

#### [Connector.APIResources](docs/sdks/apiresources/README.md)

* [Get](docs/sdks/apiresources/README.md#get) - Get API Resource

#### [Connector.Apis](docs/sdks/apis/README.md)

* [List](docs/sdks/apis/README.md#list) - List APIs
* [Get](docs/sdks/apis/README.md#get) - Get API

#### [Connector.ConnectorDocs](docs/sdks/connectordocs/README.md)

* [Get](docs/sdks/connectordocs/README.md#get) - Get Connector Doc content

#### [Connector.ConnectorResources](docs/sdks/connectorresources/README.md)

* [Get](docs/sdks/connectorresources/README.md#get) - Get Connector Resource

#### [Connector.Connectors](docs/sdks/connectors/README.md)

* [List](docs/sdks/connectors/README.md#list) - List Connectors
* [Get](docs/sdks/connectors/README.md#get) - Get Connector

### [Crm](docs/sdks/crm/README.md)


#### [Crm.Activities](docs/sdks/activities/README.md)

* [List](docs/sdks/activities/README.md#list) - List activities
* [Create](docs/sdks/activities/README.md#create) - Create activity
* [Get](docs/sdks/activities/README.md#get) - Get activity
* [Update](docs/sdks/activities/README.md#update) - Update activity
* [Delete](docs/sdks/activities/README.md#delete) - Delete activity

#### [Crm.Companies](docs/sdks/companies/README.md)

* [List](docs/sdks/companies/README.md#list) - List companies
* [Create](docs/sdks/companies/README.md#create) - Create company
* [Get](docs/sdks/companies/README.md#get) - Get company
* [Update](docs/sdks/companies/README.md#update) - Update company
* [Delete](docs/sdks/companies/README.md#delete) - Delete company

#### [Crm.Contacts](docs/sdks/contacts/README.md)

* [List](docs/sdks/contacts/README.md#list) - List contacts
* [Create](docs/sdks/contacts/README.md#create) - Create contact
* [Get](docs/sdks/contacts/README.md#get) - Get contact
* [Update](docs/sdks/contacts/README.md#update) - Update contact
* [Delete](docs/sdks/contacts/README.md#delete) - Delete contact

#### [Crm.Leads](docs/sdks/leads/README.md)

* [List](docs/sdks/leads/README.md#list) - List leads
* [Create](docs/sdks/leads/README.md#create) - Create lead
* [Get](docs/sdks/leads/README.md#get) - Get lead
* [Update](docs/sdks/leads/README.md#update) - Update lead
* [Delete](docs/sdks/leads/README.md#delete) - Delete lead

#### [Crm.Notes](docs/sdks/notes/README.md)

* [List](docs/sdks/notes/README.md#list) - List notes
* [Create](docs/sdks/notes/README.md#create) - Create note
* [Get](docs/sdks/notes/README.md#get) - Get note
* [Update](docs/sdks/notes/README.md#update) - Update note
* [Delete](docs/sdks/notes/README.md#delete) - Delete note

#### [Crm.Opportunities](docs/sdks/opportunities/README.md)

* [List](docs/sdks/opportunities/README.md#list) - List opportunities
* [Create](docs/sdks/opportunities/README.md#create) - Create opportunity
* [Get](docs/sdks/opportunities/README.md#get) - Get opportunity
* [Update](docs/sdks/opportunities/README.md#update) - Update opportunity
* [Delete](docs/sdks/opportunities/README.md#delete) - Delete opportunity

#### [Crm.Pipelines](docs/sdks/pipelines/README.md)

* [List](docs/sdks/pipelines/README.md#list) - List pipelines

#### [Crm.Users](docs/sdks/users/README.md)

* [List](docs/sdks/users/README.md#list) - List users
* [Create](docs/sdks/users/README.md#create) - Create user
* [Get](docs/sdks/users/README.md#get) - Get user
* [Update](docs/sdks/users/README.md#update) - Update user
* [Delete](docs/sdks/users/README.md#delete) - Delete user

### [Ecommerce](docs/sdks/ecommerce/README.md)


#### [Ecommerce.Customers](docs/sdks/apideckcustomers/README.md)

* [List](docs/sdks/apideckcustomers/README.md#list) - List Customers
* [Get](docs/sdks/apideckcustomers/README.md#get) - Get Customer

#### [Ecommerce.Orders](docs/sdks/orders/README.md)

* [List](docs/sdks/orders/README.md#list) - List Orders
* [Get](docs/sdks/orders/README.md#get) - Get Order

#### [Ecommerce.Products](docs/sdks/products/README.md)

* [List](docs/sdks/products/README.md#list) - List Products
* [Get](docs/sdks/products/README.md#get) - Get Product

#### [Ecommerce.Stores](docs/sdks/stores/README.md)

* [Get](docs/sdks/stores/README.md#get) - Get Store

### [FileStorage](docs/sdks/filestorage/README.md)


#### [FileStorage.DriveGroups](docs/sdks/drivegroups/README.md)

* [List](docs/sdks/drivegroups/README.md#list) - List DriveGroups
* [Create](docs/sdks/drivegroups/README.md#create) - Create DriveGroup
* [Get](docs/sdks/drivegroups/README.md#get) - Get DriveGroup
* [Update](docs/sdks/drivegroups/README.md#update) - Update DriveGroup
* [Delete](docs/sdks/drivegroups/README.md#delete) - Delete DriveGroup

#### [FileStorage.Drives](docs/sdks/drives/README.md)

* [List](docs/sdks/drives/README.md#list) - List Drives
* [Create](docs/sdks/drives/README.md#create) - Create Drive
* [Get](docs/sdks/drives/README.md#get) - Get Drive
* [Update](docs/sdks/drives/README.md#update) - Update Drive
* [Delete](docs/sdks/drives/README.md#delete) - Delete Drive

#### [FileStorage.Files](docs/sdks/files/README.md)

* [List](docs/sdks/files/README.md#list) - List Files
* [Search](docs/sdks/files/README.md#search) - Search Files
* [Get](docs/sdks/files/README.md#get) - Get File
* [Update](docs/sdks/files/README.md#update) - Rename or move File
* [Delete](docs/sdks/files/README.md#delete) - Delete File
* [Download](docs/sdks/files/README.md#download) - Download File
* [Export](docs/sdks/files/README.md#export) - Export File

#### [FileStorage.Folders](docs/sdks/folders/README.md)

* [Create](docs/sdks/folders/README.md#create) - Create Folder
* [Get](docs/sdks/folders/README.md#get) - Get Folder
* [Update](docs/sdks/folders/README.md#update) - Rename or move Folder
* [Delete](docs/sdks/folders/README.md#delete) - Delete Folder
* [Copy](docs/sdks/folders/README.md#copy) - Copy Folder

#### [FileStorage.SharedLinks](docs/sdks/sharedlinks/README.md)

* [List](docs/sdks/sharedlinks/README.md#list) - List SharedLinks
* [Create](docs/sdks/sharedlinks/README.md#create) - Create Shared Link
* [Get](docs/sdks/sharedlinks/README.md#get) - Get Shared Link
* [Update](docs/sdks/sharedlinks/README.md#update) - Update Shared Link
* [Delete](docs/sdks/sharedlinks/README.md#delete) - Delete Shared Link

#### [FileStorage.UploadSessions](docs/sdks/uploadsessions/README.md)

* [Create](docs/sdks/uploadsessions/README.md#create) - Start Upload Session
* [Get](docs/sdks/uploadsessions/README.md#get) - Get Upload Session
* [Delete](docs/sdks/uploadsessions/README.md#delete) - Abort Upload Session
* [Finish](docs/sdks/uploadsessions/README.md#finish) - Finish Upload Session

### [Hris](docs/sdks/hris/README.md)


#### [Hris.Companies](docs/sdks/apideckcompanies/README.md)

* [List](docs/sdks/apideckcompanies/README.md#list) - List Companies
* [Create](docs/sdks/apideckcompanies/README.md#create) - Create Company
* [Get](docs/sdks/apideckcompanies/README.md#get) - Get Company
* [Update](docs/sdks/apideckcompanies/README.md#update) - Update Company
* [Delete](docs/sdks/apideckcompanies/README.md#delete) - Delete Company

#### [Hris.Departments](docs/sdks/apideckdepartments/README.md)

* [List](docs/sdks/apideckdepartments/README.md#list) - List Departments
* [Create](docs/sdks/apideckdepartments/README.md#create) - Create Department
* [Get](docs/sdks/apideckdepartments/README.md#get) - Get Department
* [Update](docs/sdks/apideckdepartments/README.md#update) - Update Department
* [Delete](docs/sdks/apideckdepartments/README.md#delete) - Delete Department

#### [Hris.EmployeePayrolls](docs/sdks/employeepayrolls/README.md)

* [List](docs/sdks/employeepayrolls/README.md#list) - List Employee Payrolls
* [Get](docs/sdks/employeepayrolls/README.md#get) - Get Employee Payroll

#### [Hris.Employees](docs/sdks/employees/README.md)

* [List](docs/sdks/employees/README.md#list) - List Employees
* [Create](docs/sdks/employees/README.md#create) - Create Employee
* [Get](docs/sdks/employees/README.md#get) - Get Employee
* [Update](docs/sdks/employees/README.md#update) - Update Employee
* [Delete](docs/sdks/employees/README.md#delete) - Delete Employee

#### [Hris.EmployeeSchedules](docs/sdks/employeeschedules/README.md)

* [List](docs/sdks/employeeschedules/README.md#list) - List Employee Schedules

#### [Hris.Payrolls](docs/sdks/payrolls/README.md)

* [List](docs/sdks/payrolls/README.md#list) - List Payroll
* [Get](docs/sdks/payrolls/README.md#get) - Get Payroll

#### [Hris.TimeOffRequests](docs/sdks/timeoffrequests/README.md)

* [List](docs/sdks/timeoffrequests/README.md#list) - List Time Off Requests
* [Create](docs/sdks/timeoffrequests/README.md#create) - Create Time Off Request
* [Get](docs/sdks/timeoffrequests/README.md#get) - Get Time Off Request
* [Update](docs/sdks/timeoffrequests/README.md#update) - Update Time Off Request
* [Delete](docs/sdks/timeoffrequests/README.md#delete) - Delete Time Off Request

### [IssueTracking](docs/sdks/issuetracking/README.md)


#### [IssueTracking.Collections](docs/sdks/collections/README.md)

* [List](docs/sdks/collections/README.md#list) - List Collections
* [Get](docs/sdks/collections/README.md#get) - Get Collection

#### [IssueTracking.CollectionTags](docs/sdks/collectiontags/README.md)

* [List](docs/sdks/collectiontags/README.md#list) - List Tags

#### [IssueTracking.CollectionTicketComments](docs/sdks/collectionticketcomments/README.md)

* [List](docs/sdks/collectionticketcomments/README.md#list) - List Comments
* [Create](docs/sdks/collectionticketcomments/README.md#create) - Create Comment
* [Get](docs/sdks/collectionticketcomments/README.md#get) - Get Comment
* [Update](docs/sdks/collectionticketcomments/README.md#update) - Update Comment
* [Delete](docs/sdks/collectionticketcomments/README.md#delete) - Delete Comment

#### [IssueTracking.CollectionTickets](docs/sdks/collectiontickets/README.md)

* [List](docs/sdks/collectiontickets/README.md#list) - List Tickets
* [Create](docs/sdks/collectiontickets/README.md#create) - Create Ticket
* [Get](docs/sdks/collectiontickets/README.md#get) - Get Ticket
* [Update](docs/sdks/collectiontickets/README.md#update) - Update Ticket
* [Delete](docs/sdks/collectiontickets/README.md#delete) - Delete Ticket

#### [IssueTracking.CollectionUsers](docs/sdks/collectionusers/README.md)

* [List](docs/sdks/collectionusers/README.md#list) - List Users
* [Get](docs/sdks/collectionusers/README.md#get) - Get user

### [Sms](docs/sdks/sms/README.md)


#### [Sms.Messages](docs/sdks/messages/README.md)

* [List](docs/sdks/messages/README.md#list) - List Messages
* [Create](docs/sdks/messages/README.md#create) - Create Message
* [Get](docs/sdks/messages/README.md#get) - Get Message
* [Update](docs/sdks/messages/README.md#update) - Update Message
* [Delete](docs/sdks/messages/README.md#delete) - Delete Message

### [Vault](docs/sdks/vault/README.md)


#### [Vault.ConnectionCustomMappings](docs/sdks/connectioncustommappings/README.md)

* [List](docs/sdks/connectioncustommappings/README.md#list) - List connection custom mappings

#### [Vault.Connections](docs/sdks/connections/README.md)

* [List](docs/sdks/connections/README.md#list) - Get all connections
* [Get](docs/sdks/connections/README.md#get) - Get connection
* [Update](docs/sdks/connections/README.md#update) - Update connection
* [Delete](docs/sdks/connections/README.md#delete) - Deletes a connection
* [Imports](docs/sdks/connections/README.md#imports) - Import connection
* [Token](docs/sdks/connections/README.md#token) - Authorize Access Token

#### [Vault.ConnectionSettings](docs/sdks/connectionsettings/README.md)

* [List](docs/sdks/connectionsettings/README.md#list) - Get resource settings
* [Update](docs/sdks/connectionsettings/README.md#update) - Update settings

#### [Vault.ConsumerRequestCounts](docs/sdks/consumerrequestcounts/README.md)

* [List](docs/sdks/consumerrequestcounts/README.md#list) - Consumer request counts

#### [Vault.Consumers](docs/sdks/consumers/README.md)

* [Create](docs/sdks/consumers/README.md#create) - Create consumer
* [List](docs/sdks/consumers/README.md#list) - Get all consumers
* [Get](docs/sdks/consumers/README.md#get) - Get consumer
* [Update](docs/sdks/consumers/README.md#update) - Update consumer
* [Delete](docs/sdks/consumers/README.md#delete) - Delete consumer

#### [Vault.CreateCallback](docs/sdks/createcallback/README.md)

* [State](docs/sdks/createcallback/README.md#state) - Create Callback State

#### [Vault.CustomFields](docs/sdks/customfields/README.md)

* [List](docs/sdks/customfields/README.md#list) - Get resource custom fields

#### [Vault.CustomMappings](docs/sdks/custommappings/README.md)

* [List](docs/sdks/custommappings/README.md#list) - List custom mappings

#### [Vault.Logs](docs/sdks/logs/README.md)

* [List](docs/sdks/logs/README.md#list) - Get all consumer request logs

#### [Vault.Sessions](docs/sdks/sessions/README.md)

* [Create](docs/sdks/sessions/README.md#create) - Create Session

#### [Vault.ValidateConnection](docs/sdks/validateconnection/README.md)

* [State](docs/sdks/validateconnection/README.md#state) - Validate Connection State

### [Webhook](docs/sdks/webhook/README.md)


#### [Webhook.EventLogs](docs/sdks/eventlogs/README.md)

* [List](docs/sdks/eventlogs/README.md#list) - List event logs

#### [Webhook.Webhooks](docs/sdks/webhooks/README.md)

* [List](docs/sdks/webhooks/README.md#list) - List webhook subscriptions
* [Create](docs/sdks/webhooks/README.md#create) - Create webhook subscription
* [Get](docs/sdks/webhooks/README.md#get) - Get webhook subscription
* [Update](docs/sdks/webhooks/README.md#update) - Update webhook subscription
* [Delete](docs/sdks/webhooks/README.md#delete) - Delete webhook subscription

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"github.com/apideck-libraries/sdk-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"github.com/apideck-libraries/sdk-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `List` function may return the following errors:

| Error Type                        | Status Code | Content Type     |
| --------------------------------- | ----------- | ---------------- |
| apierrors.BadRequestResponse      | 400         | application/json |
| apierrors.UnauthorizedResponse    | 401         | application/json |
| apierrors.PaymentRequiredResponse | 402         | application/json |
| apierrors.NotFoundResponse        | 404         | application/json |
| apierrors.UnprocessableResponse   | 422         | application/json |
| apierrors.APIError                | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/apierrors"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {

		var e *apierrors.BadRequestResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.UnauthorizedResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.PaymentRequiredResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.NotFoundResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.UnprocessableResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithServerURL("https://unify.apideck.com"),
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.Accounting.TaxRates.List(ctx, operations.AccountingTaxRatesAllRequest{
		ServiceID: sdkgo.String("salesforce"),
		Filter: &components.TaxRatesFilter{
			Assets:      sdkgo.Bool(true),
			Equity:      sdkgo.Bool(true),
			Expenses:    sdkgo.Bool(true),
			Liabilities: sdkgo.Bool(true),
			Revenue:     sdkgo.Bool(true),
		},
		PassThrough: map[string]any{
			"search": "San Francisco",
		},
		Fields: sdkgo.String("id,updated_at"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.GetTaxRatesResponse != nil {
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

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	sdkgo "github.com/apideck-libraries/sdk-go"
	"github.com/apideck-libraries/sdk-go/models/components"
	"github.com/apideck-libraries/sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkgo.New(
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
	)

	res, err := s.FileStorage.UploadSessions.Create(ctx, components.CreateUploadSessionRequest{
		Name:           "Documents",
		ParentFolderID: "1234",
		DriveID:        sdkgo.String("1234"),
		Size:           sdkgo.Int64(1810673),
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
	}, nil, sdkgo.String("salesforce"), operations.WithServerURL("https://upload.apideck.com"))
	if err != nil {
		log.Fatal(err)
	}
	if res.CreateUploadSessionResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/apideck-libraries/sdk-go&utm_campaign=go)
