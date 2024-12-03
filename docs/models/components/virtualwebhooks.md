# VirtualWebhooks

Virtual webhook config for the connector.


## Fields

| Field                                                                                               | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `RequestRate`                                                                                       | [components.RequestRate](../../models/components/requestrate.md)                                    | :heavy_check_mark:                                                                                  | The rate at which requests for resources will be made to downstream.                                |
| `Resources`                                                                                         | map[string][components.WebhookSupportResources](../../models/components/webhooksupportresources.md) | :heavy_minus_sign:                                                                                  | The resources that will be requested from downstream.                                               |