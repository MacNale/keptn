// Code generated by go-swagger; DO NOT EDIT.

package event

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/api/models"
)

// GetEventOKCode is the HTTP code returned for type GetEventOK
const GetEventOKCode int = 200

/*GetEventOK Success

swagger:response getEventOK
*/
type GetEventOK struct {

	/*
	  In: Body
	*/
	Payload *models.KeptnContextExtendedCE `json:"body,omitempty"`
}

// NewGetEventOK creates GetEventOK with default headers values
func NewGetEventOK() *GetEventOK {

	return &GetEventOK{}
}

// WithPayload adds the payload to the get event o k response
func (o *GetEventOK) WithPayload(payload *models.KeptnContextExtendedCE) *GetEventOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event o k response
func (o *GetEventOK) SetPayload(payload *models.KeptnContextExtendedCE) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEventNotFoundCode is the HTTP code returned for type GetEventNotFound
const GetEventNotFoundCode int = 404

/*GetEventNotFound Failed. Event could not be found.

swagger:response getEventNotFound
*/
type GetEventNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventNotFound creates GetEventNotFound with default headers values
func NewGetEventNotFound() *GetEventNotFound {

	return &GetEventNotFound{}
}

// WithPayload adds the payload to the get event not found response
func (o *GetEventNotFound) WithPayload(payload *models.Error) *GetEventNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event not found response
func (o *GetEventNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetEventDefault Error

swagger:response getEventDefault
*/
type GetEventDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetEventDefault creates GetEventDefault with default headers values
func NewGetEventDefault(code int) *GetEventDefault {
	if code <= 0 {
		code = 500
	}

	return &GetEventDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get event default response
func (o *GetEventDefault) WithStatusCode(code int) *GetEventDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get event default response
func (o *GetEventDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get event default response
func (o *GetEventDefault) WithPayload(payload *models.Error) *GetEventDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event default response
func (o *GetEventDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
