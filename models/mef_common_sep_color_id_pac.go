// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MefCommonSepColorIDPac mef common sep color Id pac
// swagger:model mef.common.SepColorIdPac
type MefCommonSepColorIDPac struct {

	// This attribute denotes the color of the EI frame, green or yellow.
	Color MefCommonTypesFrameColor `json:"color,omitempty"`
}

// Validate validates this mef common sep color Id pac
func (m *MefCommonSepColorIDPac) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MefCommonSepColorIDPac) validateColor(formats strfmt.Registry) error {

	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := m.Color.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("color")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MefCommonSepColorIDPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MefCommonSepColorIDPac) UnmarshalBinary(b []byte) error {
	var res MefCommonSepColorIDPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}