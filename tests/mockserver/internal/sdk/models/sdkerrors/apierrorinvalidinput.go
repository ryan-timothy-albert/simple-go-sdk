// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type APIErrorInvalidInput struct {
	Status int    `json:"status"`
	Error_ string `json:"error"`
}

var _ error = &APIErrorInvalidInput{}

func (e *APIErrorInvalidInput) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
