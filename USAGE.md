<!-- Start SDK Example Usage [usage] -->
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