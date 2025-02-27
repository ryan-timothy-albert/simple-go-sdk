# User
(*User*)

## Overview

Operations about user

### Available Operations

* [CreateUser](#createuser) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [GetUserByName](#getuserbyname) - Get user by user name
* [UpdateUser](#updateuser) - Update user
* [DeleteUser](#deleteuser) - Delete user

## CreateUser

This can only be done by the logged in user.

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

    res, err := s.User.CreateUser(ctx, &components.User{
        ID: openapi.Int64(10),
        Username: openapi.String("theUser"),
        FirstName: openapi.String("John"),
        LastName: openapi.String("James"),
        Email: openapi.String("john@email.com"),
        Password: openapi.String("12345"),
        Phone: openapi.String("12345"),
        UserStatus: openapi.Int(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [components.User](../../models/components/user.md)       | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateUsersWithListInput

Creates list of users with given input array

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

    res, err := s.User.CreateUsersWithListInput(ctx, []components.User{
        components.User{
            ID: openapi.Int64(10),
            Username: openapi.String("theUser"),
            FirstName: openapi.String("John"),
            LastName: openapi.String("James"),
            Email: openapi.String("john@email.com"),
            Password: openapi.String("12345"),
            Phone: openapi.String("12345"),
            UserStatus: openapi.Int(1),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `request`                                                | [[]components.User](../../.md)                           | :heavy_check_mark:                                       | The request object to use for the request.               |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.CreateUsersWithListInputResponse](../../models/operations/createuserswithlistinputresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## LoginUser

Logs user into the system

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

    res, err := s.User.LoginUser(ctx, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.String != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | **string*                                                | :heavy_minus_sign:                                       | The user name for login                                  |
| `password`                                               | **string*                                                | :heavy_minus_sign:                                       | The password for login in clear text                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.LoginUserResponse](../../models/operations/loginuserresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## LogoutUser

Logs out current logged in user session

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

    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
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

**[*operations.LogoutUserResponse](../../models/operations/logoutuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetUserByName

Get user by user name

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

    res, err := s.User.GetUserByName(ctx, "Zachery_Lubowitz15")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `username`                                                 | *string*                                                   | :heavy_check_mark:                                         | The name that needs to be fetched. Use user1 for testing.  |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.GetUserByNameResponse](../../models/operations/getuserbynameresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |

## UpdateUser

This can only be done by the logged in user.

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

    res, err := s.User.UpdateUser(ctx, "Dandre_Hand41", &components.User{
        ID: openapi.Int64(10),
        Username: openapi.String("theUser"),
        FirstName: openapi.String("John"),
        LastName: openapi.String("James"),
        Email: openapi.String("john@email.com"),
        Password: openapi.String("12345"),
        Phone: openapi.String("12345"),
        UserStatus: openapi.Int(1),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | *string*                                                 | :heavy_check_mark:                                       | name that needs to be updated                            |
| `user`                                                   | [*components.User](../../models/components/user.md)      | :heavy_minus_sign:                                       | Update an existent user in the store                     |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteUser

This can only be done by the logged in user.

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

    res, err := s.User.DeleteUser(ctx, "Demetris_Schmitt")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `username`                                               | *string*                                                 | :heavy_check_mark:                                       | The name that needs to be deleted                        |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type                     | Status Code                    | Content Type                   |
| ------------------------------ | ------------------------------ | ------------------------------ |
| apierrors.APIErrorInvalidInput | 400                            | application/json               |
| apierrors.APIErrorUnauthorized | 401                            | application/json               |
| apierrors.APIErrorNotFound     | 404                            | application/json               |
| apierrors.APIError             | 4XX, 5XX                       | \*/\*                          |