// Code generated by go-swagger; DO NOT EDIT.

package mantan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"ngajarinlinda/gen/models"
)

// GetAllMantanOKCode is the HTTP code returned for type GetAllMantanOK
const GetAllMantanOKCode int = 200

/*
GetAllMantanOK return list of mantans

swagger:response getAllMantanOK
*/
type GetAllMantanOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Mantan `json:"body,omitempty"`
}

// NewGetAllMantanOK creates GetAllMantanOK with default headers values
func NewGetAllMantanOK() *GetAllMantanOK {

	return &GetAllMantanOK{}
}

// WithPayload adds the payload to the get all mantan o k response
func (o *GetAllMantanOK) WithPayload(payload []*models.Mantan) *GetAllMantanOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all mantan o k response
func (o *GetAllMantanOK) SetPayload(payload []*models.Mantan) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllMantanOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Mantan, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAllMantanInternalServerErrorCode is the HTTP code returned for type GetAllMantanInternalServerError
const GetAllMantanInternalServerErrorCode int = 500

/*
GetAllMantanInternalServerError server error

swagger:response getAllMantanInternalServerError
*/
type GetAllMantanInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetAllMantanInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetAllMantanInternalServerError creates GetAllMantanInternalServerError with default headers values
func NewGetAllMantanInternalServerError() *GetAllMantanInternalServerError {

	return &GetAllMantanInternalServerError{}
}

// WithPayload adds the payload to the get all mantan internal server error response
func (o *GetAllMantanInternalServerError) WithPayload(payload *GetAllMantanInternalServerErrorBody) *GetAllMantanInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get all mantan internal server error response
func (o *GetAllMantanInternalServerError) SetPayload(payload *GetAllMantanInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAllMantanInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}