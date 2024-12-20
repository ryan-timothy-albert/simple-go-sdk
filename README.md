# openapi

Developer-friendly & type-safe Go SDK specifically catered to leverage *openapi* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/ryan-local/ryan-docs-demo). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

Petstore - OpenAPI 3.1: This is a sample Pet Store Server based on the OpenAPI 3.1 specification.

Some useful links:
- [OpenAPI Reference](https://www.speakeasyapi.dev/openapi)
- [The Pet Store repository](https://github.com/swagger-api/swagger-petstore)
- [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)

For more information about the API: [Find out more about Swagger](http://swagger.io)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [openapi](#openapi)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get openapi
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name     | Type   | Scheme  |
| -------- | ------ | ------- |
| `APIKey` | apiKey | API key |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Pet](docs/sdks/pet/README.md)

* [PetsStoreMonday](docs/sdks/pet/README.md#petsstoremonday) - Update an existing pet
* [MyTestPets](docs/sdks/pet/README.md#mytestpets) - Add a new pet to the store
* [FindPetsByStatusTypes](docs/sdks/pet/README.md#findpetsbystatustypes) - Finds Pets by status
* [FindPetsByTags](docs/sdks/pet/README.md#findpetsbytags) - Finds Pets by tags
* [GetPetByIDS](docs/sdks/pet/README.md#getpetbyids) - Find pet by ID
* [DeletePet](docs/sdks/pet/README.md#deletepet) - Deletes a pet
* [UploadFile](docs/sdks/pet/README.md#uploadfile) - uploads an image


### [Store](docs/sdks/store/README.md)

* [GetInventory](docs/sdks/store/README.md#getinventory) - Returns pet inventories by status
* [PlaceOrder](docs/sdks/store/README.md#placeorder) - Place an order for a pet
* [GetOrderByID](docs/sdks/store/README.md#getorderbyid) - Find purchase order by ID
* [DeleteOrder](docs/sdks/store/README.md#deleteorder) - Delete purchase order by ID

### [User](docs/sdks/user/README.md)

* [CreateUser](docs/sdks/user/README.md#createuser) - Create user
* [CreateUsersWithListInput](docs/sdks/user/README.md#createuserswithlistinput) - Creates list of users with given input array
* [LoginUser](docs/sdks/user/README.md#loginuser) - Logs user into the system
* [LogoutUser](docs/sdks/user/README.md#logoutuser) - Logs out current logged in user session
* [GetUserByName](docs/sdks/user/README.md#getuserbyname) - Get user by user name
* [UpdateUser](docs/sdks/user/README.md#updateuser) - Update user
* [DeleteUser](docs/sdks/user/README.md#deleteuser) - Delete user

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	"log"
	"models/operations"
	"openapi"
	"openapi/models/components"
	"openapi/retry"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
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
	if res.Pet != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
	"openapi/retry"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithRetryConfig(
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
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `PetsStoreMonday` function may return the following errors:

| Error Type                     | Status Code | Content Type     |
| ------------------------------ | ----------- | ---------------- |
| apierrors.APIErrorInvalidInput | 400         | application/json |
| apierrors.APIErrorUnauthorized | 401         | application/json |
| apierrors.APIErrorNotFound     | 404         | application/json |
| apierrors.APIError             | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	"log"
	"openapi"
	"openapi/models/apierrors"
	"openapi/models/components"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {

		var e *apierrors.APIErrorInvalidInput
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIErrorUnauthorized
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIErrorNotFound
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

### Server Variables

The default server `https://{environment}.petstore.io` contains variables and is set to `https://prod.petstore.io` by default. To override default values, the following options are available when initializing the SDK client instance:
 * `WithEnvironment(environment ServerEnvironment)`

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	"log"
	"openapi"
	"openapi/models/components"
)

func main() {
	ctx := context.Background()

	s := openapi.New(
		openapi.WithServerURL("https://prod.petstore.io"),
		openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
	)

	res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
		ID:   openapi.Int64(10),
		Name: "doggie",
		Category: &components.Category{
			ID:   openapi.Int64(1),
			Name: openapi.String("Dogs"),
		},
		PhotoUrls: []string{
			"<value>",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.Pet != nil {
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

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=openapi&utm_campaign=go)
