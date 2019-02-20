// Code generated by go-swagger; DO NOT EDIT.

package tapi_connectivity_audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/presto/models"
)

// PostOperationsGetConnectivityServiceAuditDetailsOKCode is the HTTP code returned for type PostOperationsGetConnectivityServiceAuditDetailsOK
const PostOperationsGetConnectivityServiceAuditDetailsOKCode int = 200

/*PostOperationsGetConnectivityServiceAuditDetailsOK Correct response

swagger:response postOperationsGetConnectivityServiceAuditDetailsOK
*/
type PostOperationsGetConnectivityServiceAuditDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiConnectivityAuditGetConnectivityServiceAuditDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetConnectivityServiceAuditDetailsOK creates PostOperationsGetConnectivityServiceAuditDetailsOK with default headers values
func NewPostOperationsGetConnectivityServiceAuditDetailsOK() *PostOperationsGetConnectivityServiceAuditDetailsOK {

	return &PostOperationsGetConnectivityServiceAuditDetailsOK{}
}

// WithPayload adds the payload to the post operations get connectivity service audit details o k response
func (o *PostOperationsGetConnectivityServiceAuditDetailsOK) WithPayload(payload *models.TapiConnectivityAuditGetConnectivityServiceAuditDetailsOutput) *PostOperationsGetConnectivityServiceAuditDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get connectivity service audit details o k response
func (o *PostOperationsGetConnectivityServiceAuditDetailsOK) SetPayload(payload *models.TapiConnectivityAuditGetConnectivityServiceAuditDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetConnectivityServiceAuditDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetConnectivityServiceAuditDetailsCreatedCode is the HTTP code returned for type PostOperationsGetConnectivityServiceAuditDetailsCreated
const PostOperationsGetConnectivityServiceAuditDetailsCreatedCode int = 201

/*PostOperationsGetConnectivityServiceAuditDetailsCreated No response

swagger:response postOperationsGetConnectivityServiceAuditDetailsCreated
*/
type PostOperationsGetConnectivityServiceAuditDetailsCreated struct {
}

// NewPostOperationsGetConnectivityServiceAuditDetailsCreated creates PostOperationsGetConnectivityServiceAuditDetailsCreated with default headers values
func NewPostOperationsGetConnectivityServiceAuditDetailsCreated() *PostOperationsGetConnectivityServiceAuditDetailsCreated {

	return &PostOperationsGetConnectivityServiceAuditDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetConnectivityServiceAuditDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetConnectivityServiceAuditDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetConnectivityServiceAuditDetailsBadRequest
const PostOperationsGetConnectivityServiceAuditDetailsBadRequestCode int = 400

/*PostOperationsGetConnectivityServiceAuditDetailsBadRequest Internal error

swagger:response postOperationsGetConnectivityServiceAuditDetailsBadRequest
*/
type PostOperationsGetConnectivityServiceAuditDetailsBadRequest struct {
}

// NewPostOperationsGetConnectivityServiceAuditDetailsBadRequest creates PostOperationsGetConnectivityServiceAuditDetailsBadRequest with default headers values
func NewPostOperationsGetConnectivityServiceAuditDetailsBadRequest() *PostOperationsGetConnectivityServiceAuditDetailsBadRequest {

	return &PostOperationsGetConnectivityServiceAuditDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetConnectivityServiceAuditDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}
