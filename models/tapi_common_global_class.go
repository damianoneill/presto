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

// TapiCommonGlobalClass tapi common global class
// swagger:model tapi.common.GlobalClass
type TapiCommonGlobalClass struct {

	// List of names. A property of an entity with a value that is unique in some namespace but may change during the life of the entity. A name carries no semantics with respect to the purpose of the entity.
	Name []*TapiCommonNameAndValue `json:"name"`

	// UUID: An identifier that is universally unique within an identifier space, where the identifier space is itself globally unique, and immutable. An UUID carries no semantics with respect to the purpose or state of the entity.
	//                     UUID here uses string representation as defined in RFC 4122.  The canonical representation uses lowercase characters.
	//                     Pattern: [0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-' + '[0-9a-fA-F]{4}-[0-9a-fA-F]{12}
	//                     Example of a UUID in string representation: f81d4fae-7dec-11d0-a765-00a0c91e6bf6
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this tapi common global class
func (m *TapiCommonGlobalClass) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TapiCommonGlobalClass) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	for i := 0; i < len(m.Name); i++ {
		if swag.IsZero(m.Name[i]) { // not required
			continue
		}

		if m.Name[i] != nil {
			if err := m.Name[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("name" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TapiCommonGlobalClass) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TapiCommonGlobalClass) UnmarshalBinary(b []byte) error {
	var res TapiCommonGlobalClass
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}