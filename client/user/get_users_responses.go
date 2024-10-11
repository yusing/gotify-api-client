// Code generated by go-swagger; DO NOT EDIT.

package user

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

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /user] getUsers", response, response.Code())
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*
GetUsersOK describes a response with status code 200, with default header values.

Ok
*/
type GetUsersOK struct {
	Payload []*models.UserExternal
}

// IsSuccess returns true when this get users o k response has a 2xx status code
func (o *GetUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get users o k response has a 3xx status code
func (o *GetUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users o k response has a 4xx status code
func (o *GetUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get users o k response has a 5xx status code
func (o *GetUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get users o k response a status code equal to that given
func (o *GetUsersOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get users o k response
func (o *GetUsersOK) Code() int {
	return 200
}

func (o *GetUsersOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersOK %s", 200, payload)
}

func (o *GetUsersOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersOK %s", 200, payload)
}

func (o *GetUsersOK) GetPayload() []*models.UserExternal {
	return o.Payload
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*
GetUsersUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetUsersUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this get users unauthorized response has a 2xx status code
func (o *GetUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users unauthorized response has a 3xx status code
func (o *GetUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users unauthorized response has a 4xx status code
func (o *GetUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users unauthorized response has a 5xx status code
func (o *GetUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get users unauthorized response a status code equal to that given
func (o *GetUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get users unauthorized response
func (o *GetUsersUnauthorized) Code() int {
	return 401
}

func (o *GetUsersUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersUnauthorized %s", 401, payload)
}

func (o *GetUsersUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersUnauthorized %s", 401, payload)
}

func (o *GetUsersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersForbidden creates a GetUsersForbidden with default headers values
func NewGetUsersForbidden() *GetUsersForbidden {
	return &GetUsersForbidden{}
}

/*
GetUsersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsersForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get users forbidden response has a 2xx status code
func (o *GetUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get users forbidden response has a 3xx status code
func (o *GetUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get users forbidden response has a 4xx status code
func (o *GetUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get users forbidden response has a 5xx status code
func (o *GetUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get users forbidden response a status code equal to that given
func (o *GetUsersForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get users forbidden response
func (o *GetUsersForbidden) Code() int {
	return 403
}

func (o *GetUsersForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersForbidden %s", 403, payload)
}

func (o *GetUsersForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /user][%d] getUsersForbidden %s", 403, payload)
}

func (o *GetUsersForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
