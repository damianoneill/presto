// Code generated by go-swagger; DO NOT EDIT.

package tapi_common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/presto/models"
)

// GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOKCode is the HTTP code returned for type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK
const GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOKCode int = 200

/*GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK tapi.connectivity.SwitchControl

swagger:response getDataContextConnectionUuidSwitchControlSwitchControlUuidOK
*/
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiConnectivitySwitchControl `json:"body,omitempty"`
}

// NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK creates GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK with default headers values
func NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK() *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK {

	return &GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK{}
}

// WithPayload adds the payload to the get data context connection Uuid switch control switch control Uuid o k response
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK) WithPayload(payload *models.TapiConnectivitySwitchControl) *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get data context connection Uuid switch control switch control Uuid o k response
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK) SetPayload(payload *models.TapiConnectivitySwitchControl) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequestCode is the HTTP code returned for type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest
const GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequestCode int = 400

/*GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest Internal error

swagger:response getDataContextConnectionUuidSwitchControlSwitchControlUuidBadRequest
*/
type GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest struct {
}

// NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest creates GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest with default headers values
func NewGetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest() *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest {

	return &GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest{}
}

// WriteResponse to the client
func (o *GetDataContextConnectionUUIDSwitchControlSwitchControlUUIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}