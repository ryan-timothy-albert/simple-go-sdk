// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type MyTestPetsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful operation
	Pet *components.Pet
}

func (o *MyTestPetsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *MyTestPetsResponse) GetPet() *components.Pet {
	if o == nil {
		return nil
	}
	return o.Pet
}