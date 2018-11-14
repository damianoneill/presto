// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewPostOperationsGetServiceInterfacePointListParams creates a new PostOperationsGetServiceInterfacePointListParams object
// no default values defined in spec.
func NewPostOperationsGetServiceInterfacePointListParams() PostOperationsGetServiceInterfacePointListParams {

	return PostOperationsGetServiceInterfacePointListParams{}
}

// PostOperationsGetServiceInterfacePointListParams contains all the bound params for the post operations get service interface point list operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostOperationsGetServiceInterfacePointList
type PostOperationsGetServiceInterfacePointListParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPostOperationsGetServiceInterfacePointListParams() beforehand.
func (o *PostOperationsGetServiceInterfacePointListParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}