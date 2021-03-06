// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiCommonGetServiceInterfacePointListOutput tapi common get service interface point list output
// swagger:model tapi.common.GetServiceInterfacePointListOutput
type TapiCommonGetServiceInterfacePointListOutput struct {

	// output
	Output *TapiCommonGetServiceInterfacePointListOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi common get service interface point list output
func (m *TapiCommonGetServiceInterfacePointListOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGetServiceInterfacePointListOutput) validateOutput(formats strfmt.Registry) error {

	if swag.IsZero(m.Output) { // not required
		return nil
	}

	if m.Output != nil {
		if err := m.Output.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("output")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointListOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointListOutput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointListOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiCommonGetServiceInterfacePointListOutputOutput tapi common get service interface point list output output
// swagger:model TapiCommonGetServiceInterfacePointListOutputOutput
type TapiCommonGetServiceInterfacePointListOutputOutput struct {

	// none
	Sip []*TapiCommonGetserviceinterfacepointlistOutputSip `json:"sip"`
}

// Validate validates this tapi common get service interface point list output output
func (m *TapiCommonGetServiceInterfacePointListOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSip(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGetServiceInterfacePointListOutputOutput) validateSip(formats strfmt.Registry) error {

	if swag.IsZero(m.Sip) { // not required
		return nil
	}

	for i := 0; i < len(m.Sip); i++ {
		if swag.IsZero(m.Sip[i]) { // not required
			continue
		}

		if m.Sip[i] != nil {
			if err := m.Sip[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + "sip" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointListOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGetServiceInterfacePointListOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiCommonGetServiceInterfacePointListOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
