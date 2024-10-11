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

// UploadAppImageReader is a Reader for the UploadAppImage structure.
type UploadAppImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadAppImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadAppImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadAppImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadAppImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadAppImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadAppImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadAppImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /application/{id}/image] uploadAppImage", response, response.Code())
	}
}

// NewUploadAppImageOK creates a UploadAppImageOK with default headers values
func NewUploadAppImageOK() *UploadAppImageOK {
	return &UploadAppImageOK{}
}

/*
UploadAppImageOK describes a response with status code 200, with default header values.

Ok
*/
type UploadAppImageOK struct {
	Payload *models.Application
}

// IsSuccess returns true when this upload app image o k response has a 2xx status code
func (o *UploadAppImageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload app image o k response has a 3xx status code
func (o *UploadAppImageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image o k response has a 4xx status code
func (o *UploadAppImageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload app image o k response has a 5xx status code
func (o *UploadAppImageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload app image o k response a status code equal to that given
func (o *UploadAppImageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload app image o k response
func (o *UploadAppImageOK) Code() int {
	return 200
}

func (o *UploadAppImageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageOK %s", 200, payload)
}

func (o *UploadAppImageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageOK %s", 200, payload)
}

func (o *UploadAppImageOK) GetPayload() *models.Application {
	return o.Payload
}

func (o *UploadAppImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Application)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAppImageBadRequest creates a UploadAppImageBadRequest with default headers values
func NewUploadAppImageBadRequest() *UploadAppImageBadRequest {
	return &UploadAppImageBadRequest{}
}

/*
UploadAppImageBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadAppImageBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload app image bad request response has a 2xx status code
func (o *UploadAppImageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload app image bad request response has a 3xx status code
func (o *UploadAppImageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image bad request response has a 4xx status code
func (o *UploadAppImageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload app image bad request response has a 5xx status code
func (o *UploadAppImageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this upload app image bad request response a status code equal to that given
func (o *UploadAppImageBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the upload app image bad request response
func (o *UploadAppImageBadRequest) Code() int {
	return 400
}

func (o *UploadAppImageBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageBadRequest %s", 400, payload)
}

func (o *UploadAppImageBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageBadRequest %s", 400, payload)
}

func (o *UploadAppImageBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadAppImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAppImageUnauthorized creates a UploadAppImageUnauthorized with default headers values
func NewUploadAppImageUnauthorized() *UploadAppImageUnauthorized {
	return &UploadAppImageUnauthorized{}
}

/*
UploadAppImageUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UploadAppImageUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload app image unauthorized response has a 2xx status code
func (o *UploadAppImageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload app image unauthorized response has a 3xx status code
func (o *UploadAppImageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image unauthorized response has a 4xx status code
func (o *UploadAppImageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload app image unauthorized response has a 5xx status code
func (o *UploadAppImageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this upload app image unauthorized response a status code equal to that given
func (o *UploadAppImageUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the upload app image unauthorized response
func (o *UploadAppImageUnauthorized) Code() int {
	return 401
}

func (o *UploadAppImageUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageUnauthorized %s", 401, payload)
}

func (o *UploadAppImageUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageUnauthorized %s", 401, payload)
}

func (o *UploadAppImageUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadAppImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAppImageForbidden creates a UploadAppImageForbidden with default headers values
func NewUploadAppImageForbidden() *UploadAppImageForbidden {
	return &UploadAppImageForbidden{}
}

/*
UploadAppImageForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UploadAppImageForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload app image forbidden response has a 2xx status code
func (o *UploadAppImageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload app image forbidden response has a 3xx status code
func (o *UploadAppImageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image forbidden response has a 4xx status code
func (o *UploadAppImageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload app image forbidden response has a 5xx status code
func (o *UploadAppImageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload app image forbidden response a status code equal to that given
func (o *UploadAppImageForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload app image forbidden response
func (o *UploadAppImageForbidden) Code() int {
	return 403
}

func (o *UploadAppImageForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageForbidden %s", 403, payload)
}

func (o *UploadAppImageForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageForbidden %s", 403, payload)
}

func (o *UploadAppImageForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadAppImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAppImageNotFound creates a UploadAppImageNotFound with default headers values
func NewUploadAppImageNotFound() *UploadAppImageNotFound {
	return &UploadAppImageNotFound{}
}

/*
UploadAppImageNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UploadAppImageNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload app image not found response has a 2xx status code
func (o *UploadAppImageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload app image not found response has a 3xx status code
func (o *UploadAppImageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image not found response has a 4xx status code
func (o *UploadAppImageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload app image not found response has a 5xx status code
func (o *UploadAppImageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this upload app image not found response a status code equal to that given
func (o *UploadAppImageNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the upload app image not found response
func (o *UploadAppImageNotFound) Code() int {
	return 404
}

func (o *UploadAppImageNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageNotFound %s", 404, payload)
}

func (o *UploadAppImageNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageNotFound %s", 404, payload)
}

func (o *UploadAppImageNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadAppImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadAppImageInternalServerError creates a UploadAppImageInternalServerError with default headers values
func NewUploadAppImageInternalServerError() *UploadAppImageInternalServerError {
	return &UploadAppImageInternalServerError{}
}

/*
UploadAppImageInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UploadAppImageInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this upload app image internal server error response has a 2xx status code
func (o *UploadAppImageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload app image internal server error response has a 3xx status code
func (o *UploadAppImageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload app image internal server error response has a 4xx status code
func (o *UploadAppImageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload app image internal server error response has a 5xx status code
func (o *UploadAppImageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload app image internal server error response a status code equal to that given
func (o *UploadAppImageInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the upload app image internal server error response
func (o *UploadAppImageInternalServerError) Code() int {
	return 500
}

func (o *UploadAppImageInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageInternalServerError %s", 500, payload)
}

func (o *UploadAppImageInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /application/{id}/image][%d] uploadAppImageInternalServerError %s", 500, payload)
}

func (o *UploadAppImageInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadAppImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
