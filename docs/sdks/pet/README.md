# Pet
(*Pet*)

## Overview

Everything about your Pets

Find out more dfdsiojfidjs
<http://swagger.io>

### Available Operations

* [PetsStoreMonday](#petsstoremonday) - Update an existing pet
* [MyTestPets](#mytestpets) - Add a new pet to the store
* [PetsStoreTuesday](#petsstoretuesday) - Update an existing pet
* [MyTestPetsTuesday](#mytestpetstuesday) - Add a new pet to the store
* [FindPetsByStatusTypes](#findpetsbystatustypes) - Finds Pets by status
* [FindPetsByTags](#findpetsbytags) - Finds Pets by tags
* [GetPetByIDS](#getpetbyids) - Find pet by ID
* [DeletePet](#deletepet) - Deletes a pet
* [UploadFile](#uploadfile) - uploads an image

## PetsStoreMonday

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/components"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.PetsStoreMonday(ctx, components.Pet{
        ID: openapi.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: openapi.Int64(1),
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PetsStoreMondayResponse](../../models/operations/petsstoremondayresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## MyTestPets

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/components"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.MyTestPets(ctx, components.Pet{
        ID: openapi.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: openapi.Int64(1),
            Name: openapi.String("Dogs"),
        },
        PhotoUrls: []string{
            "<value>",
            "<value>",
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MyTestPetsResponse](../../models/operations/mytestpetsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PetsStoreTuesday

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/components"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.PetsStoreTuesday(ctx, components.Pet{
        ID: openapi.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: openapi.Int64(1),
            Name: openapi.String("Dogs"),
        },
        PhotoUrls: []string{
            "<value>",
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PetsStoreTuesdayResponse](../../models/operations/petsstoretuesdayresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## MyTestPetsTuesday

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"openapi/models/components"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.MyTestPetsTuesday(ctx, components.Pet{
        ID: openapi.Int64(10),
        Name: "doggie",
        Category: &components.Category{
            ID: openapi.Int64(1),
            Name: openapi.String("Dogs"),
        },
        PhotoUrls: []string{
            "<value>",
            "<value>",
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

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Pet](../../models/components/pet.md)         | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.MyTestPetsTuesdayResponse](../../models/operations/mytestpetstuesdayresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## FindPetsByStatusTypes

Multiple status values can be provided with comma separated strings

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.FindPetsByStatusTypes(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `status`                                                 | [*operations.Status](../../models/operations/status.md)  | :heavy_minus_sign:                                       | Status values that need to be considered for filter      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FindPetsByStatusTypesResponse](../../models/operations/findpetsbystatustypesresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.FindPetsByTags(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `tags`                                                   | []*string*                                               | :heavy_minus_sign:                                       | Tags to filter by                                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.FindPetsByTagsResponse](../../models/operations/findpetsbytagsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## GetPetByIDS

Returns a single pet

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.GetPetByIDS(ctx, 137396)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | ID of pet to return                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPetByIDSResponse](../../models/operations/getpetbyidsresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## DeletePet

Deletes a pet

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.DeletePet(ctx, 441876, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Pet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | Pet id to delete                                         |
| `apiKey`                                                 | **string*                                                | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeletePetResponse](../../models/operations/deletepetresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## UploadFile

uploads an image

### Example Usage

```go
package main

import(
	"context"
	"openapi"
	"log"
)

func main() {
ctx := context.Background()


    s := openapi.New(
        openapi.WithSecurity("<YOUR_API_KEY_HERE>"),
    )

    res, err := s.Pet.UploadFile(ctx, 565380, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.APIResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `petID`                                                  | *int64*                                                  | :heavy_check_mark:                                       | ID of pet to update                                      |
| `additionalMetadata`                                     | **string*                                                | :heavy_minus_sign:                                       | Additional Metadata                                      |
| `requestBody`                                            | **any*                                                   | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |