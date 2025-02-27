# Store
(*Store*)

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>

### Available Operations

* [GetInventory](#getinventory) - Returns pet inventories by status
* [PlaceOrder](#placeorder) - Place an order for a pet
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [DeleteOrder](#deleteorder) - Delete purchase order by ID

## GetInventory

Returns a map of status codes to quantities

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

    res, err := s.Store.GetInventory(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetInventoryResponse](../../models/operations/getinventoryresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## PlaceOrder

Place a new order in the store

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

    res, err := s.Store.PlaceOrder(ctx, &components.Order{
        ID: openapi.Int64(10),
        PetID: openapi.Int64(198772),
        Quantity: openapi.Int(7),
        Status: components.OrderStatusApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.Order](../../models/components/order.md)     | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.PlaceOrderResponse](../../models/operations/placeorderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

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

    res, err := s.Store.GetOrderByID(ctx, 614993)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `orderID`                                                | *int64*                                                  | :heavy_check_mark:                                       | ID of order that needs to be fetched                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOrderByIDResponse](../../models/operations/getorderbyidresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

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

    res, err := s.Store.DeleteOrder(ctx, 127902)
    if err != nil {
        log.Fatal(err)
    }
    if res.Order != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `orderID`                                                | *int64*                                                  | :heavy_check_mark:                                       | ID of the order that needs to be deleted                 |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOrderResponse](../../models/operations/deleteorderresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |