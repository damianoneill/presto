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

// TapiTopologyInterRuleGroup tapi topology inter rule group
// swagger:model tapi.topology.InterRuleGroup
type TapiTopologyInterRuleGroup struct {
	TapiCommonCapacityPac

	TapiCommonGlobalClass

	TapiTopologyRiskParameterPac

	TapiTopologyTransferCostPac

	TapiTopologyTransferTimingPac

	// none
	AssociatedNodeRuleGroup []*TapiTopologyNodeRuleGroupRef `json:"associated-node-rule-group"`

	// none
	Rule []*TapiTopologyRule `json:"rule"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiTopologyInterRuleGroup) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonCapacityPac
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonCapacityPac = aO0

	// AO1
	var aO1 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO1

	// AO2
	var aO2 TapiTopologyRiskParameterPac
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.TapiTopologyRiskParameterPac = aO2

	// AO3
	var aO3 TapiTopologyTransferCostPac
	if err := swag.ReadJSON(raw, &aO3); err != nil {
		return err
	}
	m.TapiTopologyTransferCostPac = aO3

	// AO4
	var aO4 TapiTopologyTransferTimingPac
	if err := swag.ReadJSON(raw, &aO4); err != nil {
		return err
	}
	m.TapiTopologyTransferTimingPac = aO4

	// AO5
	var dataAO5 struct {
		AssociatedNodeRuleGroup []*TapiTopologyNodeRuleGroupRef `json:"associated-node-rule-group"`

		Rule []*TapiTopologyRule `json:"rule"`
	}
	if err := swag.ReadJSON(raw, &dataAO5); err != nil {
		return err
	}

	m.AssociatedNodeRuleGroup = dataAO5.AssociatedNodeRuleGroup

	m.Rule = dataAO5.Rule

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiTopologyInterRuleGroup) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 6)

	aO0, err := swag.WriteJSON(m.TapiCommonCapacityPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.TapiTopologyRiskParameterPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	aO3, err := swag.WriteJSON(m.TapiTopologyTransferCostPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO3)

	aO4, err := swag.WriteJSON(m.TapiTopologyTransferTimingPac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO4)

	var dataAO5 struct {
		AssociatedNodeRuleGroup []*TapiTopologyNodeRuleGroupRef `json:"associated-node-rule-group"`

		Rule []*TapiTopologyRule `json:"rule"`
	}

	dataAO5.AssociatedNodeRuleGroup = m.AssociatedNodeRuleGroup

	dataAO5.Rule = m.Rule

	jsonDataAO5, errAO5 := swag.WriteJSON(dataAO5)
	if errAO5 != nil {
		return nil, errAO5
	}
	_parts = append(_parts, jsonDataAO5)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi topology inter rule group
func (m *TapiTopologyInterRuleGroup) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonCapacityPac
	if err := m.TapiCommonCapacityPac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiTopologyRiskParameterPac
	if err := m.TapiTopologyRiskParameterPac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiTopologyTransferCostPac
	if err := m.TapiTopologyTransferCostPac.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiTopologyTransferTimingPac
	if err := m.TapiTopologyTransferTimingPac.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssociatedNodeRuleGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiTopologyInterRuleGroup) validateAssociatedNodeRuleGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.AssociatedNodeRuleGroup) { // not required
		return nil
	}

	for i := 0; i < len(m.AssociatedNodeRuleGroup); i++ {
		if swag.IsZero(m.AssociatedNodeRuleGroup[i]) { // not required
			continue
		}

		if m.AssociatedNodeRuleGroup[i] != nil {
			if err := m.AssociatedNodeRuleGroup[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("associated-node-rule-group" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiTopologyInterRuleGroup) validateRule(formats strfmt.Registry) error {

	if swag.IsZero(m.Rule) { // not required
		return nil
	}

	for i := 0; i < len(m.Rule); i++ {
		if swag.IsZero(m.Rule[i]) { // not required
			continue
		}

		if m.Rule[i] != nil {
			if err := m.Rule[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rule" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyInterRuleGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyInterRuleGroup) UnmarshalBinary(b []byte) error {
	var res TapiTopologyInterRuleGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}