# Proxy

## Overview

### Available Operations

* [Get](#get) - GET
* [Options](#options) - OPTIONS
* [Post](#post) - POST
* [Put](#put) - PUT
* [Patch](#patch) - PATCH
* [Delete](#delete) - DELETE

## Get

Proxies a downstream GET request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.getProxy" method="get" path="/proxy" -->
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

    res, err := s.Proxy.Get(ctx, operations.ProxyGetProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ProxyGetProxyRequest](../../models/operations/proxygetproxyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ProxyGetProxyResponse](../../models/operations/proxygetproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## Options

Proxies a downstream OPTION request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.optionsProxy" method="options" path="/proxy" -->
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

    res, err := s.Proxy.Options(ctx, operations.ProxyOptionsProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ProxyOptionsProxyRequest](../../models/operations/proxyoptionsproxyrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `opts`                                                                                     | [][operations.Option](../../models/operations/option.md)                                   | :heavy_minus_sign:                                                                         | The options for this request.                                                              |

### Response

**[*operations.ProxyOptionsProxyResponse](../../models/operations/proxyoptionsproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## Post

Proxies a downstream POST request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.postProxy" method="post" path="/proxy" -->
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

    res, err := s.Proxy.Post(ctx, operations.ProxyPostProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ProxyPostProxyRequest](../../models/operations/proxypostproxyrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `opts`                                                                               | [][operations.Option](../../models/operations/option.md)                             | :heavy_minus_sign:                                                                   | The options for this request.                                                        |

### Response

**[*operations.ProxyPostProxyResponse](../../models/operations/proxypostproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## Put

Proxies a downstream PUT request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.putProxy" method="put" path="/proxy" -->
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

    res, err := s.Proxy.Put(ctx, operations.ProxyPutProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ProxyPutProxyRequest](../../models/operations/proxyputproxyrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ProxyPutProxyResponse](../../models/operations/proxyputproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## Patch

Proxies a downstream PATCH request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.patchProxy" method="patch" path="/proxy" -->
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

    res, err := s.Proxy.Patch(ctx, operations.ProxyPatchProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ProxyPatchProxyRequest](../../models/operations/proxypatchproxyrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ProxyPatchProxyResponse](../../models/operations/proxypatchproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |

## Delete

Proxies a downstream DELETE request to a service and injects the necessary credentials into a request stored in Vault. This allows you to have an additional security layer and logging without needing to rely on Unify for normalization.
**Note**: Vault will proxy all data to the downstream URL and method/verb in the headers.


### Example Usage

<!-- UsageSnippet language="go" operationID="proxy.deleteProxy" method="delete" path="/proxy" -->
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

    res, err := s.Proxy.Delete(ctx, operations.ProxyDeleteProxyRequest{
        ServiceID: "close",
        UnifiedAPI: sdkgo.Pointer("hris"),
        DownstreamURL: "https://api.close.com/api/v1/lead",
        DownstreamAuthorization: sdkgo.Pointer("Bearer <token>"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ResponseJSON != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ProxyDeleteProxyRequest](../../models/operations/proxydeleteproxyrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.ProxyDeleteProxyResponse](../../models/operations/proxydeleteproxyresponse.md), error**

### Errors

| Error Type             | Status Code            | Content Type           |
| ---------------------- | ---------------------- | ---------------------- |
| apierrors.Unauthorized | 401                    | application/json       |
| apierrors.APIError     | 4XX, 5XX               | \*/\*                  |