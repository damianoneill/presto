// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MefCommonPcpColorIDPac mef common pcp color Id pac
// swagger:model mef.common.PcpColorIdPac
type MefCommonPcpColorIDPac struct {

	// This attribute provides a list PCP values map to the green ingress EI frames. The pcpValueForGreenList and the pcpValueForYellowList must disjoint and the union of the two lists must include all possible PCP values.
	PcpValueForGreenList []int64 `json:"pcp-value-for-green-list"`

	// This attribute provides a list PCP values map to the yellow ingress EI frames. The pcpValueForGreenList and the pcpValueForYellowList must disjoint and the union of the two lists must include all possible PCP values.
	PcpValueForYellowList []int64 `json:"pcp-value-for-yellow-list"`
}

// Validate validates this mef common pcp color Id pac
func (m *MefCommonPcpColorIDPac) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MefCommonPcpColorIDPac) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MefCommonPcpColorIDPac) UnmarshalBinary(b []byte) error {
	var res MefCommonPcpColorIDPac
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}