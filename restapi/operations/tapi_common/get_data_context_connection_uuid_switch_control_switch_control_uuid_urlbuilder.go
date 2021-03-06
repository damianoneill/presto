// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL generates an URL for the get data context connection UUID switch control switch control UUID operation
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL struct {
	SwitchControlUUID string
	UUID              string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) WithBasePath(bp string) *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/data/context/connection={uuid}/switch-control={switch-control-uuid}/"

	switchControlUUID := o.SwitchControlUUID
	if switchControlUUID != "" {
		_path = strings.Replace(_path, "{switch-control-uuid}", switchControlUUID, -1)
	} else {
		return nil, errors.New("switchControlUuid is required on GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL")
	}

	uuid := o.UUID
	if uuid != "" {
		_path = strings.Replace(_path, "{uuid}", uuid, -1)
	} else {
		return nil, errors.New("uuid is required on GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
