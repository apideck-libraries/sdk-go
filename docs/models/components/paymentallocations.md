# PaymentAllocations


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    | Example                                        |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `ID`                                           | **string*                                      | :heavy_minus_sign:                             | ID of the payment                              | 123456                                         |
| `AllocatedAmount`                              | **float64*                                     | :heavy_minus_sign:                             | Amount of the payment allocated to the invoice | 1000                                           |
| `Date`                                         | [*time.Time](https://pkg.go.dev/time#Time)     | :heavy_minus_sign:                             | Date of the payment                            | 2020-09-30T07:43:32.000Z                       |