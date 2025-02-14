// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type GetInventoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	Object map[string]int
}

func (o *GetInventoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetInventoryResponse) GetObject() map[string]int {
	if o == nil {
		return nil
	}
	return o.Object
}
