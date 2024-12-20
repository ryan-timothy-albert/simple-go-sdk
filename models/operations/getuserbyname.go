// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetUserByNameRequest struct {
	// The name that needs to be fetched. Use user1 for testing.
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserByNameRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserByNameResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// successful operation
	User *components.User
}

func (o *GetUserByNameResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserByNameResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
