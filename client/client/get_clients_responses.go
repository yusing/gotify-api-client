// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/yusing/gotify-api-client/v2/models"
)

// GetClientsReader is a Reader for the GetClients structure.
type GetClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetClientsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetClientsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /client] getClients", response, response.Code())
	}
}

// NewGetClientsOK creates a GetClientsOK with default headers values
func NewGetClientsOK() *GetClientsOK {
	return &GetClientsOK{}
}

/*
GetClientsOK describes a response with status code 200, with default header values.

Ok
*/
type GetClientsOK struct {
	Payload []*models.Client
}

// IsSuccess returns true when this get clients o k response has a 2xx status code
func (o *GetClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get clients o k response has a 3xx status code
func (o *GetClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clients o k response has a 4xx status code
func (o *GetClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get clients o k response has a 5xx status code
func (o *GetClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get clients o k response a status code equal to that given
func (o *GetClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get clients o k response
func (o *GetClientsOK) Code() int {
	return 200
}

func (o *GetClientsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsOK %s", 200, payload)
}

func (o *GetClientsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsOK %s", 200, payload)
}

func (o *GetClientsOK) GetPayload() []*models.Client {
	return o.Payload
}

func (o *GetClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientsUnauthorized creates a GetClientsUnauthorized with default headers values
func NewGetClientsUnauthorized() *GetClientsUnauthorized {
	return &GetClientsUnauthorized{}
}

/*
GetClientsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetClientsUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clients unauthorized response has a 2xx status code
func (o *GetClientsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clients unauthorized response has a 3xx status code
func (o *GetClientsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clients unauthorized response has a 4xx status code
func (o *GetClientsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clients unauthorized response has a 5xx status code
func (o *GetClientsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get clients unauthorized response a status code equal to that given
func (o *GetClientsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get clients unauthorized response
func (o *GetClientsUnauthorized) Code() int {
	return 401
}

func (o *GetClientsUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsUnauthorized %s", 401, payload)
}

func (o *GetClientsUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsUnauthorized %s", 401, payload)
}

func (o *GetClientsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClientsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClientsForbidden creates a GetClientsForbidden with default headers values
func NewGetClientsForbidden() *GetClientsForbidden {
	return &GetClientsForbidden{}
}

/*
GetClientsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetClientsForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get clients forbidden response has a 2xx status code
func (o *GetClientsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get clients forbidden response has a 3xx status code
func (o *GetClientsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get clients forbidden response has a 4xx status code
func (o *GetClientsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get clients forbidden response has a 5xx status code
func (o *GetClientsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get clients forbidden response a status code equal to that given
func (o *GetClientsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get clients forbidden response
func (o *GetClientsForbidden) Code() int {
	return 403
}

func (o *GetClientsForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsForbidden %s", 403, payload)
}

func (o *GetClientsForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /client][%d] getClientsForbidden %s", 403, payload)
}

func (o *GetClientsForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetClientsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
