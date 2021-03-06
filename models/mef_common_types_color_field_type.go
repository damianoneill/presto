// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// MefCommonTypesColorFieldType mef common types color field type
// swagger:model mef.common.types.ColorFieldType
type MefCommonTypesColorFieldType string

const (

	// MefCommonTypesColorFieldTypePCP captures enum value "PCP"
	MefCommonTypesColorFieldTypePCP MefCommonTypesColorFieldType = "PCP"

	// MefCommonTypesColorFieldTypeDEI captures enum value "DEI"
	MefCommonTypesColorFieldTypeDEI MefCommonTypesColorFieldType = "DEI"

	// MefCommonTypesColorFieldTypeENDPOINT captures enum value "END_POINT"
	MefCommonTypesColorFieldTypeENDPOINT MefCommonTypesColorFieldType = "END_POINT"

	// MefCommonTypesColorFieldTypeDSCP captures enum value "DSCP"
	MefCommonTypesColorFieldTypeDSCP MefCommonTypesColorFieldType = "DSCP"
)

// for schema
var mefCommonTypesColorFieldTypeEnum []interface{}

func init() {
	var res []MefCommonTypesColorFieldType
	if err := json.Unmarshal([]byte(`["PCP","DEI","END_POINT","DSCP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mefCommonTypesColorFieldTypeEnum = append(mefCommonTypesColorFieldTypeEnum, v)
	}
}

func (m MefCommonTypesColorFieldType) validateMefCommonTypesColorFieldTypeEnum(path, location string, value MefCommonTypesColorFieldType) error {
	if err := validate.Enum(path, location, value, mefCommonTypesColorFieldTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this mef common types color field type
func (m MefCommonTypesColorFieldType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMefCommonTypesColorFieldTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
