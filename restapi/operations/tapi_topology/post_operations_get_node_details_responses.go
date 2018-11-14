// Code generated by go-swagger; DO NOT EDIT.

package tapi_topology

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/damianoneill/presto/models"
)

// PostOperationsGetNodeDetailsOKCode is the HTTP code returned for type PostOperationsGetNodeDetailsOK
const PostOperationsGetNodeDetailsOKCode int = 200

/*PostOperationsGetNodeDetailsOK Correct response

swagger:response postOperationsGetNodeDetailsOK
*/
type PostOperationsGetNodeDetailsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TapiTopologyGetNodeDetailsOutput `json:"body,omitempty"`
}

// NewPostOperationsGetNodeDetailsOK creates PostOperationsGetNodeDetailsOK with default headers values
func NewPostOperationsGetNodeDetailsOK() *PostOperationsGetNodeDetailsOK {

	return &PostOperationsGetNodeDetailsOK{}
}

// WithPayload adds the payload to the post operations get node details o k response
func (o *PostOperationsGetNodeDetailsOK) WithPayload(payload *models.TapiTopologyGetNodeDetailsOutput) *PostOperationsGetNodeDetailsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post operations get node details o k response
func (o *PostOperationsGetNodeDetailsOK) SetPayload(payload *models.TapiTopologyGetNodeDetailsOutput) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostOperationsGetNodeDetailsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostOperationsGetNodeDetailsCreatedCode is the HTTP code returned for type PostOperationsGetNodeDetailsCreated
const PostOperationsGetNodeDetailsCreatedCode int = 201

/*PostOperationsGetNodeDetailsCreated No response

swagger:response postOperationsGetNodeDetailsCreated
*/
type PostOperationsGetNodeDetailsCreated struct {
}

// NewPostOperationsGetNodeDetailsCreated creates PostOperationsGetNodeDetailsCreated with default headers values
func NewPostOperationsGetNodeDetailsCreated() *PostOperationsGetNodeDetailsCreated {

	return &PostOperationsGetNodeDetailsCreated{}
}

// WriteResponse to the client
func (o *PostOperationsGetNodeDetailsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// PostOperationsGetNodeDetailsBadRequestCode is the HTTP code returned for type PostOperationsGetNodeDetailsBadRequest
const PostOperationsGetNodeDetailsBadRequestCode int = 400

/*PostOperationsGetNodeDetailsBadRequest Internal error

swagger:response postOperationsGetNodeDetailsBadRequest
*/
type PostOperationsGetNodeDetailsBadRequest struct {
}

// NewPostOperationsGetNodeDetailsBadRequest creates PostOperationsGetNodeDetailsBadRequest with default headers values
func NewPostOperationsGetNodeDetailsBadRequest() *PostOperationsGetNodeDetailsBadRequest {

	return &PostOperationsGetNodeDetailsBadRequest{}
}

// WriteResponse to the client
func (o *PostOperationsGetNodeDetailsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}