// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TapiTopologyOwnedNodeEdgePointRef tapi topology owned node edge point ref
// swagger:model tapi.topology.OwnedNodeEdgePointRef
type TapiTopologyOwnedNodeEdgePointRef struct {
	TapiTopologyNodeRef

	// owned node edge point id
	OwnedNodeEdgePointID string `json:"owned-node-edge-point-id,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiTopologyOwnedNodeEdgePointRef) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiTopologyNodeRef
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiTopologyNodeRef = aO0

	// AO1
	var dataAO1 struct {
		OwnedNodeEdgePointID string `json:"owned-node-edge-point-id,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.OwnedNodeEdgePointID = dataAO1.OwnedNodeEdgePointID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiTopologyOwnedNodeEdgePointRef) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.TapiTopologyNodeRef)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		OwnedNodeEdgePointID string `json:"owned-node-edge-point-id,omitempty"`
	}

	dataAO1.OwnedNodeEdgePointID = m.OwnedNodeEdgePointID

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi topology owned node edge point ref
func (m *TapiTopologyOwnedNodeEdgePointRef) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiTopologyNodeRef
	if err := m.TapiTopologyNodeRef.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TapiTopologyOwnedNodeEdgePointRef) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiTopologyOwnedNodeEdgePointRef) UnmarshalBinary(b []byte) error {
	var res TapiTopologyOwnedNodeEdgePointRef
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
