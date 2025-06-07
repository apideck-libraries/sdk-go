<!-- Start SDK Example Usage [usage] -->
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
		sdkgo.WithConsumerID("test-consumer"),
		sdkgo.WithAppID("dSBdXd2H6Mqwfg0atXHXYcysLJE9qyn1VwBtXHX"),
		sdkgo.WithSecurity(os.Getenv("APIDECK_API_KEY")),
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