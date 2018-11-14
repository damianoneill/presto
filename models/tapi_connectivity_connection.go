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

// TapiConnectivityConnection tapi connectivity connection
// swagger:model tapi.connectivity.Connection
type TapiConnectivityConnection struct {
	TapiCommonGlobalClass

	TapiCommonOperationalStatePac

	// none
	ConnectionEndPoint []*TapiConnectivityConnectionEndPointRef `json:"connection-end-point"`

	// none
	Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

	// none
	LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

	// An Connection object supports a recursive aggregation relationship such that the internal construction of an Connection can be exposed as multiple lower level Connection objects (partitioning).
	//                     Aggregation is used as for the Node/Topology  to allow changes in hierarchy.
	//                     Connection aggregation reflects Node/Topology aggregation.
	//                     The FC represents a Cross-Connection in an NE. The Cross-Connection in an NE is not necessarily the lowest level of FC partitioning.
	LowerConnection []string `json:"lower-connection"`

	// none
	Route []*TapiConnectivityRoute `json:"route"`

	// none
	SwitchControl []*TapiConnectivitySwitchControl `json:"switch-control"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *TapiConnectivityConnection) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TapiCommonGlobalClass
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TapiCommonGlobalClass = aO0

	// AO1
	var aO1 TapiCommonOperationalStatePac
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.TapiCommonOperationalStatePac = aO1

	// AO2
	var dataAO2 struct {
		ConnectionEndPoint []*TapiConnectivityConnectionEndPointRef `json:"connection-end-point"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		LowerConnection []string `json:"lower-connection"`

		Route []*TapiConnectivityRoute `json:"route"`

		SwitchControl []*TapiConnectivitySwitchControl `json:"switch-control"`
	}
	if err := swag.ReadJSON(raw, &dataAO2); err != nil {
		return err
	}

	m.ConnectionEndPoint = dataAO2.ConnectionEndPoint

	m.Direction = dataAO2.Direction

	m.LayerProtocolName = dataAO2.LayerProtocolName

	m.LowerConnection = dataAO2.LowerConnection

	m.Route = dataAO2.Route

	m.SwitchControl = dataAO2.SwitchControl

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m TapiConnectivityConnection) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.TapiCommonGlobalClass)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.TapiCommonOperationalStatePac)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	var dataAO2 struct {
		ConnectionEndPoint []*TapiConnectivityConnectionEndPointRef `json:"connection-end-point"`

		Direction TapiCommonForwardingDirection `json:"direction,omitempty"`

		LayerProtocolName TapiCommonLayerProtocolName `json:"layer-protocol-name,omitempty"`

		LowerConnection []string `json:"lower-connection"`

		Route []*TapiConnectivityRoute `json:"route"`

		SwitchControl []*TapiConnectivitySwitchControl `json:"switch-control"`
	}

	dataAO2.ConnectionEndPoint = m.ConnectionEndPoint

	dataAO2.Direction = m.Direction

	dataAO2.LayerProtocolName = m.LayerProtocolName

	dataAO2.LowerConnection = m.LowerConnection

	dataAO2.Route = m.Route

	dataAO2.SwitchControl = m.SwitchControl

	jsonDataAO2, errAO2 := swag.WriteJSON(dataAO2)
	if errAO2 != nil {
		return nil, errAO2
	}
	_parts = append(_parts, jsonDataAO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this tapi connectivity connection
func (m *TapiConnectivityConnection) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TapiCommonGlobalClass
	if err := m.TapiCommonGlobalClass.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with TapiCommonOperationalStatePac
	if err := m.TapiCommonOperationalStatePac.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectionEndPoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayerProtocolName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwitchControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiConnectivityConnection) validateConnectionEndPoint(formats strfmt.Registry) error {

	if swag.IsZero(m.ConnectionEndPoint) { // not required
		return nil
	}

	for i := 0; i < len(m.ConnectionEndPoint); i++ {
		if swag.IsZero(m.ConnectionEndPoint[i]) { // not required
			continue
		}

		if m.ConnectionEndPoint[i] != nil {
			if err := m.ConnectionEndPoint[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connection-end-point" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityConnection) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	if err := m.Direction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("direction")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityConnection) validateLayerProtocolName(formats strfmt.Registry) error {

	if swag.IsZero(m.LayerProtocolName) { // not required
		return nil
	}

	if err := m.LayerProtocolName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("layer-protocol-name")
		}
		return err
	}

	return nil
}

func (m *TapiConnectivityConnection) validateRoute(formats strfmt.Registry) error {

	if swag.IsZero(m.Route) { // not required
		return nil
	}

	for i := 0; i < len(m.Route); i++ {
		if swag.IsZero(m.Route[i]) { // not required
			continue
		}

		if m.Route[i] != nil {
			if err := m.Route[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("route" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *TapiConnectivityConnection) validateSwitchControl(formats strfmt.Registry) error {

	if swag.IsZero(m.SwitchControl) { // not required
		return nil
	}

	for i := 0; i < len(m.SwitchControl); i++ {
		if swag.IsZero(m.SwitchControl[i]) { // not required
			continue
		}

		if m.SwitchControl[i] != nil {
			if err := m.SwitchControl[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("switch-control" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiConnectivityConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiConnectivityConnection) UnmarshalBinary(b []byte) error {
	var res TapiConnectivityConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}