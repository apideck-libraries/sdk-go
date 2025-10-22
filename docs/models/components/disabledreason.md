# DisabledReason

Indicates why the webhook has been disabled. `retry_limit`: webhook reached its retry limit. `usage_limit`: account is over its usage limit. `delivery_url_validation_failed`: delivery URL failed validation during webhook creation or update.


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `DisabledReasonNone`                        | none                                        |
| `DisabledReasonRetryLimit`                  | retry_limit                                 |
| `DisabledReasonUsageLimit`                  | usage_limit                                 |
| `DisabledReasonDeliveryURLValidationFailed` | delivery_url_validation_failed              |