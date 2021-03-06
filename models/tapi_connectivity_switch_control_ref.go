// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiConnectivitySwitchControlRef tapi connectivity switch control ref
// swagger:model tapi.connectivity.SwitchControlRef
type TapiConnectivitySwitchControlRef struct {
	TapiConnectivityConnectionRef

	// switch control id
	SwitchControlID string `json:"switch-control-id,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivitySwitchControlRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiConnectivityConnectionRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiConnectivityConnectionRef = aO0

	// AO1
	var dataAO1 struct {
		SwitchControlID string `json:"switch-control-id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.SwitchControlID = dataAO1.SwitchControlID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivitySwitchControlRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiConnectivityConnectionRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		SwitchControlID string `json:"switch-control-id,omitempty"`
	}

	dataAO1.SwitchControlID = m.SwitchControlID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity switch control ref
func (m *TapiConnectivitySwitchControlRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiConnectivityConnectionRef
	if err := m.TapiConnectivityConnectionRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivitySwitchControlRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivitySwitchControlRef) UnmarshalBinary(b []byte) error {
	var res TapiConnectivitySwitchControlRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
