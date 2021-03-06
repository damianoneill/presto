// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NrpInterfaceCreateConnectivityServiceInputAugmentation1 nrp interface create connectivity service input augmentation1
// swagger:model nrp.interface.CreateConnectivityServiceInputAugmentation1
type NrpInterfaceCreateConnectivityServiceInputAugmentation1 struct {

	// none
	NrpCarrierEthConnectivityResource *NrpInterfaceNrpCarrierEthConnectivityResource `json:"nrp-carrier-eth-connectivity-resource,omitempty"`
}

// Validate validates this nrp interface create connectivity service input augmentation1
func (m *NrpInterfaceCreateConnectivityServiceInputAugmentation1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNrpCarrierEthConnectivityResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceCreateConnectivityServiceInputAugmentation1) validateNrpCarrierEthConnectivityResource(formats strfmt.Registry) error {

	if swag.IsZero(m.NrpCarrierEthConnectivityResource) { // not required
		return nil
	}

	if m.NrpCarrierEthConnectivityResource != nil {
		if err := m.NrpCarrierEthConnectivityResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nrp-carrier-eth-connectivity-resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceCreateConnectivityServiceInputAugmentation1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceCreateConnectivityServiceInputAugmentation1) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceCreateConnectivityServiceInputAugmentation1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
