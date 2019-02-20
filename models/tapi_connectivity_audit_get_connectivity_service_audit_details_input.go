// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput tapi connectivity audit get connectivity service audit details input
// swagger:model tapi.connectivity.audit.GetConnectivityServiceAuditDetailsInput
type TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput struct {

	// input
	Input *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput `json:"input,omitempty"`
}

// Validate validates this tapi connectivity audit get connectivity service audit details input
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput) validateInput(formats strfmt.Registry) error {

	if swag.IsZero(m.Input) { // not required
		return nil
	}

	if m.Input != nil {
		if err := m.Input.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("input")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityAuditGetConnectivityServiceAuditDetailsInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput tapi connectivity audit get connectivity service audit details input input
// swagger:model TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput
type TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput struct {

	// none
	ServiceIDOrName string `json:"service-id-or-name,omitempty"`
}

// Validate validates this tapi connectivity audit get connectivity service audit details input input
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityAuditGetConnectivityServiceAuditDetailsInputInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
