// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchUserQueueReader is a Reader for the PatchUserQueue structure.
type PatchUserQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUserQueueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUserQueueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUserQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchUserQueueRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchUserQueueRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchUserQueueUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUserQueueTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUserQueueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUserQueueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchUserQueueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUserQueueOK creates a PatchUserQueueOK with default headers values
func NewPatchUserQueueOK() *PatchUserQueueOK {
	return &PatchUserQueueOK{}
}

/*PatchUserQueueOK handles this case with default header values.

successful operation
*/
type PatchUserQueueOK struct {
	Payload *models.UserQueue
}

func (o *PatchUserQueueOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueOK  %+v", 200, o.Payload)
}

func (o *PatchUserQueueOK) GetPayload() *models.UserQueue {
	return o.Payload
}

func (o *PatchUserQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueBadRequest creates a PatchUserQueueBadRequest with default headers values
func NewPatchUserQueueBadRequest() *PatchUserQueueBadRequest {
	return &PatchUserQueueBadRequest{}
}

/*PatchUserQueueBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchUserQueueBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUserQueueBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueUnauthorized creates a PatchUserQueueUnauthorized with default headers values
func NewPatchUserQueueUnauthorized() *PatchUserQueueUnauthorized {
	return &PatchUserQueueUnauthorized{}
}

/*PatchUserQueueUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchUserQueueUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUserQueueUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueForbidden creates a PatchUserQueueForbidden with default headers values
func NewPatchUserQueueForbidden() *PatchUserQueueForbidden {
	return &PatchUserQueueForbidden{}
}

/*PatchUserQueueForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchUserQueueForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserQueueForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueNotFound creates a PatchUserQueueNotFound with default headers values
func NewPatchUserQueueNotFound() *PatchUserQueueNotFound {
	return &PatchUserQueueNotFound{}
}

/*PatchUserQueueNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchUserQueueNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserQueueNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueRequestTimeout creates a PatchUserQueueRequestTimeout with default headers values
func NewPatchUserQueueRequestTimeout() *PatchUserQueueRequestTimeout {
	return &PatchUserQueueRequestTimeout{}
}

/*PatchUserQueueRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchUserQueueRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchUserQueueRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueRequestEntityTooLarge creates a PatchUserQueueRequestEntityTooLarge with default headers values
func NewPatchUserQueueRequestEntityTooLarge() *PatchUserQueueRequestEntityTooLarge {
	return &PatchUserQueueRequestEntityTooLarge{}
}

/*PatchUserQueueRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchUserQueueRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUserQueueRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueUnsupportedMediaType creates a PatchUserQueueUnsupportedMediaType with default headers values
func NewPatchUserQueueUnsupportedMediaType() *PatchUserQueueUnsupportedMediaType {
	return &PatchUserQueueUnsupportedMediaType{}
}

/*PatchUserQueueUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchUserQueueUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUserQueueUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueTooManyRequests creates a PatchUserQueueTooManyRequests with default headers values
func NewPatchUserQueueTooManyRequests() *PatchUserQueueTooManyRequests {
	return &PatchUserQueueTooManyRequests{}
}

/*PatchUserQueueTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchUserQueueTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUserQueueTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueInternalServerError creates a PatchUserQueueInternalServerError with default headers values
func NewPatchUserQueueInternalServerError() *PatchUserQueueInternalServerError {
	return &PatchUserQueueInternalServerError{}
}

/*PatchUserQueueInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchUserQueueInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUserQueueInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueServiceUnavailable creates a PatchUserQueueServiceUnavailable with default headers values
func NewPatchUserQueueServiceUnavailable() *PatchUserQueueServiceUnavailable {
	return &PatchUserQueueServiceUnavailable{}
}

/*PatchUserQueueServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchUserQueueServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUserQueueServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueueGatewayTimeout creates a PatchUserQueueGatewayTimeout with default headers values
func NewPatchUserQueueGatewayTimeout() *PatchUserQueueGatewayTimeout {
	return &PatchUserQueueGatewayTimeout{}
}

/*PatchUserQueueGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchUserQueueGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueueGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues/{queueId}][%d] patchUserQueueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUserQueueGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
