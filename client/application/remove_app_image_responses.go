// Code generated by go-swagger; DO NOT EDIT.

package application

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

// RemoveAppImageReader is a Reader for the RemoveAppImage structure.
type RemoveAppImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAppImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveAppImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveAppImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRemoveAppImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveAppImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveAppImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRemoveAppImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /application/{id}/image] removeAppImage", response, response.Code())
	}
}

// NewRemoveAppImageOK creates a RemoveAppImageOK with default headers values
func NewRemoveAppImageOK() *RemoveAppImageOK {
	return &RemoveAppImageOK{}
}

/*
RemoveAppImageOK describes a response with status code 200, with default header values.

Ok
*/
type RemoveAppImageOK struct {
}

// IsSuccess returns true when this remove app image o k response has a 2xx status code
func (o *RemoveAppImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this remove app image o k response has a 3xx status code
func (o *RemoveAppImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image o k response has a 4xx status code
func (o *RemoveAppImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove app image o k response has a 5xx status code
func (o *RemoveAppImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this remove app image o k response a status code equal to that given
func (o *RemoveAppImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the remove app image o k response
func (o *RemoveAppImageOK) Code() int {
	return 200
}

func (o *RemoveAppImageOK) Error() string {
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageOK", 200)
}

func (o *RemoveAppImageOK) String() string {
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageOK", 200)
}

func (o *RemoveAppImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveAppImageBadRequest creates a RemoveAppImageBadRequest with default headers values
func NewRemoveAppImageBadRequest() *RemoveAppImageBadRequest {
	return &RemoveAppImageBadRequest{}
}

/*
RemoveAppImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RemoveAppImageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove app image bad request response has a 2xx status code
func (o *RemoveAppImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove app image bad request response has a 3xx status code
func (o *RemoveAppImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image bad request response has a 4xx status code
func (o *RemoveAppImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove app image bad request response has a 5xx status code
func (o *RemoveAppImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this remove app image bad request response a status code equal to that given
func (o *RemoveAppImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the remove app image bad request response
func (o *RemoveAppImageBadRequest) Code() int {
	return 400
}

func (o *RemoveAppImageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageBadRequest %s", 400, payload)
}

func (o *RemoveAppImageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageBadRequest %s", 400, payload)
}

func (o *RemoveAppImageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAppImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppImageUnauthorized creates a RemoveAppImageUnauthorized with default headers values
func NewRemoveAppImageUnauthorized() *RemoveAppImageUnauthorized {
	return &RemoveAppImageUnauthorized{}
}

/*
RemoveAppImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RemoveAppImageUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove app image unauthorized response has a 2xx status code
func (o *RemoveAppImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove app image unauthorized response has a 3xx status code
func (o *RemoveAppImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image unauthorized response has a 4xx status code
func (o *RemoveAppImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove app image unauthorized response has a 5xx status code
func (o *RemoveAppImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this remove app image unauthorized response a status code equal to that given
func (o *RemoveAppImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the remove app image unauthorized response
func (o *RemoveAppImageUnauthorized) Code() int {
	return 401
}

func (o *RemoveAppImageUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageUnauthorized %s", 401, payload)
}

func (o *RemoveAppImageUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageUnauthorized %s", 401, payload)
}

func (o *RemoveAppImageUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAppImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppImageForbidden creates a RemoveAppImageForbidden with default headers values
func NewRemoveAppImageForbidden() *RemoveAppImageForbidden {
	return &RemoveAppImageForbidden{}
}

/*
RemoveAppImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RemoveAppImageForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove app image forbidden response has a 2xx status code
func (o *RemoveAppImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove app image forbidden response has a 3xx status code
func (o *RemoveAppImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image forbidden response has a 4xx status code
func (o *RemoveAppImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove app image forbidden response has a 5xx status code
func (o *RemoveAppImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this remove app image forbidden response a status code equal to that given
func (o *RemoveAppImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the remove app image forbidden response
func (o *RemoveAppImageForbidden) Code() int {
	return 403
}

func (o *RemoveAppImageForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageForbidden %s", 403, payload)
}

func (o *RemoveAppImageForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageForbidden %s", 403, payload)
}

func (o *RemoveAppImageForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAppImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppImageNotFound creates a RemoveAppImageNotFound with default headers values
func NewRemoveAppImageNotFound() *RemoveAppImageNotFound {
	return &RemoveAppImageNotFound{}
}

/*
RemoveAppImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RemoveAppImageNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove app image not found response has a 2xx status code
func (o *RemoveAppImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove app image not found response has a 3xx status code
func (o *RemoveAppImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image not found response has a 4xx status code
func (o *RemoveAppImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this remove app image not found response has a 5xx status code
func (o *RemoveAppImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this remove app image not found response a status code equal to that given
func (o *RemoveAppImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the remove app image not found response
func (o *RemoveAppImageNotFound) Code() int {
	return 404
}

func (o *RemoveAppImageNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageNotFound %s", 404, payload)
}

func (o *RemoveAppImageNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageNotFound %s", 404, payload)
}

func (o *RemoveAppImageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAppImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveAppImageInternalServerError creates a RemoveAppImageInternalServerError with default headers values
func NewRemoveAppImageInternalServerError() *RemoveAppImageInternalServerError {
	return &RemoveAppImageInternalServerError{}
}

/*
RemoveAppImageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type RemoveAppImageInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this remove app image internal server error response has a 2xx status code
func (o *RemoveAppImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this remove app image internal server error response has a 3xx status code
func (o *RemoveAppImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this remove app image internal server error response has a 4xx status code
func (o *RemoveAppImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this remove app image internal server error response has a 5xx status code
func (o *RemoveAppImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this remove app image internal server error response a status code equal to that given
func (o *RemoveAppImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the remove app image internal server error response
func (o *RemoveAppImageInternalServerError) Code() int {
	return 500
}

func (o *RemoveAppImageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageInternalServerError %s", 500, payload)
}

func (o *RemoveAppImageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /application/{id}/image][%d] removeAppImageInternalServerError %s", 500, payload)
}

func (o *RemoveAppImageInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RemoveAppImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
