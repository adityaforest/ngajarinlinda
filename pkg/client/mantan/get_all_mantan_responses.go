// Code generated by go-swagger; DO NOT EDIT.

package mantan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"ngajarinlinda/gen/models"
)

// GetAllMantanReader is a Reader for the GetAllMantan structure.
type GetAllMantanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllMantanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllMantanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetAllMantanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /api/v1/mantan] getAllMantan", response, response.Code())
	}
}

// NewGetAllMantanOK creates a GetAllMantanOK with default headers values
func NewGetAllMantanOK() *GetAllMantanOK {
	return &GetAllMantanOK{}
}

/*
GetAllMantanOK describes a response with status code 200, with default header values.

return list of mantans
*/
type GetAllMantanOK struct {
	Payload []*models.Mantan
}

// IsSuccess returns true when this get all mantan o k response has a 2xx status code
func (o *GetAllMantanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all mantan o k response has a 3xx status code
func (o *GetAllMantanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all mantan o k response has a 4xx status code
func (o *GetAllMantanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all mantan o k response has a 5xx status code
func (o *GetAllMantanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all mantan o k response a status code equal to that given
func (o *GetAllMantanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get all mantan o k response
func (o *GetAllMantanOK) Code() int {
	return 200
}

func (o *GetAllMantanOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/mantan][%d] getAllMantanOK  %+v", 200, o.Payload)
}

func (o *GetAllMantanOK) String() string {
	return fmt.Sprintf("[GET /api/v1/mantan][%d] getAllMantanOK  %+v", 200, o.Payload)
}

func (o *GetAllMantanOK) GetPayload() []*models.Mantan {
	return o.Payload
}

func (o *GetAllMantanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllMantanInternalServerError creates a GetAllMantanInternalServerError with default headers values
func NewGetAllMantanInternalServerError() *GetAllMantanInternalServerError {
	return &GetAllMantanInternalServerError{}
}

/*
GetAllMantanInternalServerError describes a response with status code 500, with default header values.

server error
*/
type GetAllMantanInternalServerError struct {
	Payload *GetAllMantanInternalServerErrorBody
}

// IsSuccess returns true when this get all mantan internal server error response has a 2xx status code
func (o *GetAllMantanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all mantan internal server error response has a 3xx status code
func (o *GetAllMantanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all mantan internal server error response has a 4xx status code
func (o *GetAllMantanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all mantan internal server error response has a 5xx status code
func (o *GetAllMantanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get all mantan internal server error response a status code equal to that given
func (o *GetAllMantanInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get all mantan internal server error response
func (o *GetAllMantanInternalServerError) Code() int {
	return 500
}

func (o *GetAllMantanInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1/mantan][%d] getAllMantanInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllMantanInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v1/mantan][%d] getAllMantanInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAllMantanInternalServerError) GetPayload() *GetAllMantanInternalServerErrorBody {
	return o.Payload
}

func (o *GetAllMantanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAllMantanInternalServerErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetAllMantanInternalServerErrorBody get all mantan internal server error body
swagger:model GetAllMantanInternalServerErrorBody
*/
type GetAllMantanInternalServerErrorBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get all mantan internal server error body
func (o *GetAllMantanInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get all mantan internal server error body based on context it is used
func (o *GetAllMantanInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAllMantanInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAllMantanInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetAllMantanInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
