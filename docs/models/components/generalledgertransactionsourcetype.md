# GeneralLedgerTransactionSourceType

The originating transaction type that produced this posting in the general ledger. Discriminates whether the entry came from an invoice, a bill, a payment, a manual journal, etc. This is the key field that distinguishes general-ledger-transactions from journal-entries (which only covers manually-captured entries). May be `null` when the upstream connector did not return an origin discriminator (e.g. Xero's single-record endpoint strips `SourceType` for every record; certain historical records also omit it). To recover a populated value, query the list endpoint.

## Example Usage

```go
import (
	"github.com/apideck-libraries/sdk-go/models/components"
)

value := components.GeneralLedgerTransactionSourceTypeOther

// Open enum: custom values can be created with a direct type cast
custom := components.GeneralLedgerTransactionSourceType("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `GeneralLedgerTransactionSourceTypeOther`        | other                                            |
| `GeneralLedgerTransactionSourceTypeInvoice`      | invoice                                          |
| `GeneralLedgerTransactionSourceTypeBill`         | bill                                             |
| `GeneralLedgerTransactionSourceTypeCreditNote`   | credit_note                                      |
| `GeneralLedgerTransactionSourceTypePayment`      | payment                                          |
| `GeneralLedgerTransactionSourceTypeRefund`       | refund                                           |
| `GeneralLedgerTransactionSourceTypeExpense`      | expense                                          |
| `GeneralLedgerTransactionSourceTypeJournalEntry` | journal_entry                                    |
| `GeneralLedgerTransactionSourceTypePayroll`      | payroll                                          |