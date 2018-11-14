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

// TapiTopologyGetTopologyListOutput tapi topology get topology list output
// swagger:model tapi.topology.GetTopologyListOutput
type TapiTopologyGetTopologyListOutput struct {

	// output
	Output *TapiTopologyGetTopologyListOutputOutput `json:"output,omitempty"`
}

// Validate validates this tapi topology get topology list output
func (m *TapiTopologyGetTopologyListOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutput(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyGetTopologyListOutput) validateOutput(formats strfmt.Registry) error {

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
func (m *TapiTopologyGetTopologyListOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyGetTopologyListOutput) UnmarshalBinary(b []byte) error {
	var res TapiTopologyGetTopologyListOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TapiTopologyGetTopologyListOutputOutput tapi topology get topology list output output
// swagger:model TapiTopologyGetTopologyListOutputOutput
type TapiTopologyGetTopologyListOutputOutput struct {

	// none
	Topology []*TapiTopologyTopology `json:"topology"`
}

// Validate validates this tapi topology get topology list output output
func (m *TapiTopologyGetTopologyListOutputOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyGetTopologyListOutputOutput) validateTopology(formats strfmt.Registry) error {

	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	for i := 0; i < len(m.Topology); i++ {
		if swag.IsZero(m.Topology[i]) { // not required
			continue
		}

		if m.Topology[i] != nil {
			if err := m.Topology[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("output" + "." + "topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyGetTopologyListOutputOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyGetTopologyListOutputOutput) UnmarshalBinary(b []byte) error {
	var res TapiTopologyGetTopologyListOutputOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}