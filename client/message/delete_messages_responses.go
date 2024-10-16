// Code generated by go-swagger; DO NOT EDIT.

package message

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

// DeleteMessagesReader is a Reader for the DeleteMessages structure.
type DeleteMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /message] deleteMessages", response, response.Code())
	}
}

// NewDeleteMessagesOK creates a DeleteMessagesOK with default headers values
func NewDeleteMessagesOK() *DeleteMessagesOK {
	return &DeleteMessagesOK{}
}

/*
DeleteMessagesOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteMessagesOK struct {
}

// IsSuccess returns true when this delete messages o k response has a 2xx status code
func (o *DeleteMessagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete messages o k response has a 3xx status code
func (o *DeleteMessagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete messages o k response has a 4xx status code
func (o *DeleteMessagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete messages o k response has a 5xx status code
func (o *DeleteMessagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete messages o k response a status code equal to that given
func (o *DeleteMessagesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete messages o k response
func (o *DeleteMessagesOK) Code() int {
	return 200
}

func (o *DeleteMessagesOK) Error() string {
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesOK", 200)
}

func (o *DeleteMessagesOK) String() string {
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesOK", 200)
}

func (o *DeleteMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMessagesUnauthorized creates a DeleteMessagesUnauthorized with default headers values
func NewDeleteMessagesUnauthorized() *DeleteMessagesUnauthorized {
	return &DeleteMessagesUnauthorized{}
}

/*
DeleteMessagesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteMessagesUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete messages unauthorized response has a 2xx status code
func (o *DeleteMessagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete messages unauthorized response has a 3xx status code
func (o *DeleteMessagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete messages unauthorized response has a 4xx status code
func (o *DeleteMessagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete messages unauthorized response has a 5xx status code
func (o *DeleteMessagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete messages unauthorized response a status code equal to that given
func (o *DeleteMessagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete messages unauthorized response
func (o *DeleteMessagesUnauthorized) Code() int {
	return 401
}

func (o *DeleteMessagesUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesUnauthorized %s", 401, payload)
}

func (o *DeleteMessagesUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesUnauthorized %s", 401, payload)
}

func (o *DeleteMessagesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMessagesForbidden creates a DeleteMessagesForbidden with default headers values
func NewDeleteMessagesForbidden() *DeleteMessagesForbidden {
	return &DeleteMessagesForbidden{}
}

/*
DeleteMessagesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteMessagesForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this delete messages forbidden response has a 2xx status code
func (o *DeleteMessagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete messages forbidden response has a 3xx status code
func (o *DeleteMessagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete messages forbidden response has a 4xx status code
func (o *DeleteMessagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete messages forbidden response has a 5xx status code
func (o *DeleteMessagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete messages forbidden response a status code equal to that given
func (o *DeleteMessagesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete messages forbidden response
func (o *DeleteMessagesForbidden) Code() int {
	return 403
}

func (o *DeleteMessagesForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesForbidden %s", 403, payload)
}

func (o *DeleteMessagesForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /message][%d] deleteMessagesForbidden %s", 403, payload)
}

func (o *DeleteMessagesForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
