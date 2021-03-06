// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NrpInterfaceNrpCarrierEthUniNResource nrp interface nrp carrier eth uni n resource
// swagger:model nrp.interface.NrpCarrierEthUniNResource
type NrpInterfaceNrpCarrierEthUniNResource struct {
	NrmConnectivityCarrierEthInterfaceResource

	// This attribute represents the relationship between the UNI-N and the egress BwpFlow.
	EgressBwpFlow *MefCommonBwpFlow `json:"egress-bwp-flow,omitempty"`

	// This attribute denotes whether the ELMI is enabled or not. When the value is TRUE, the CEN MUST meet the mandatory requirements in MEF 16 that apply to the UNI-N.
	//                         Note: Ethernel Local Management Interface protocol contents are defined which clearly identify MEF Service/Resource constructs like UNI and EVC, hence the attribute cannot be placed in an ethernet generic class.
	ElmiPeModeEnabled *bool `json:"elmi-pe-mode-enabled,omitempty"`

	// This attribute represents the relationship between the UNI-N and the ingress BwpFlow.
	IngressBwpFlow *MefCommonBwpFlow `json:"ingress-bwp-flow,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NrpInterfaceNrpCarrierEthUniNResource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 NrmConnectivityCarrierEthInterfaceResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.NrmConnectivityCarrierEthInterfaceResource = aO0

	// AO1
	var dataAO1 struct {
		EgressBwpFlow *MefCommonBwpFlow `json:"egress-bwp-flow,omitempty"`

		ElmiPeModeEnabled *bool `json:"elmi-pe-mode-enabled,omitempty"`

		IngressBwpFlow *MefCommonBwpFlow `json:"ingress-bwp-flow,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.EgressBwpFlow = dataAO1.EgressBwpFlow

	m.ElmiPeModeEnabled = dataAO1.ElmiPeModeEnabled

	m.IngressBwpFlow = dataAO1.IngressBwpFlow

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NrpInterfaceNrpCarrierEthUniNResource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.NrmConnectivityCarrierEthInterfaceResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		EgressBwpFlow *MefCommonBwpFlow `json:"egress-bwp-flow,omitempty"`

		ElmiPeModeEnabled *bool `json:"elmi-pe-mode-enabled,omitempty"`

		IngressBwpFlow *MefCommonBwpFlow `json:"ingress-bwp-flow,omitempty"`
	}

	dataAO1.EgressBwpFlow = m.EgressBwpFlow

	dataAO1.ElmiPeModeEnabled = m.ElmiPeModeEnabled

	dataAO1.IngressBwpFlow = m.IngressBwpFlow

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this nrp interface nrp carrier eth uni n resource
func (m *NrpInterfaceNrpCarrierEthUniNResource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NrmConnectivityCarrierEthInterfaceResource
	if err := m.NrmConnectivityCarrierEthInterfaceResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEgressBwpFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngressBwpFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NrpInterfaceNrpCarrierEthUniNResource) validateEgressBwpFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.EgressBwpFlow) { // not required
		return nil
	}

	if m.EgressBwpFlow != nil {
		if err := m.EgressBwpFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("egress-bwp-flow")
			}
			return err
		}
	}

	return nil
}

func (m *NrpInterfaceNrpCarrierEthUniNResource) validateIngressBwpFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.IngressBwpFlow) { // not required
		return nil
	}

	if m.IngressBwpFlow != nil {
		if err := m.IngressBwpFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ingress-bwp-flow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthUniNResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NrpInterfaceNrpCarrierEthUniNResource) UnmarshalBinary(b []byte) error {
	var res NrpInterfaceNrpCarrierEthUniNResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
