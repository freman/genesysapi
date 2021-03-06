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

// PatchUserQueuesReader is a Reader for the PatchUserQueues structure.
type PatchUserQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUserQueuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUserQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUserQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchUserQueuesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchUserQueuesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUserQueuesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUserQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUserQueuesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchUserQueuesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUserQueuesOK creates a PatchUserQueuesOK with default headers values
func NewPatchUserQueuesOK() *PatchUserQueuesOK {
	return &PatchUserQueuesOK{}
}

/*PatchUserQueuesOK handles this case with default header values.

successful operation
*/
type PatchUserQueuesOK struct {
	Payload *models.UserQueueEntityListing
}

func (o *PatchUserQueuesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesOK  %+v", 200, o.Payload)
}

func (o *PatchUserQueuesOK) GetPayload() *models.UserQueueEntityListing {
	return o.Payload
}

func (o *PatchUserQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesBadRequest creates a PatchUserQueuesBadRequest with default headers values
func NewPatchUserQueuesBadRequest() *PatchUserQueuesBadRequest {
	return &PatchUserQueuesBadRequest{}
}

/*PatchUserQueuesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchUserQueuesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUserQueuesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesUnauthorized creates a PatchUserQueuesUnauthorized with default headers values
func NewPatchUserQueuesUnauthorized() *PatchUserQueuesUnauthorized {
	return &PatchUserQueuesUnauthorized{}
}

/*PatchUserQueuesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchUserQueuesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUserQueuesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesForbidden creates a PatchUserQueuesForbidden with default headers values
func NewPatchUserQueuesForbidden() *PatchUserQueuesForbidden {
	return &PatchUserQueuesForbidden{}
}

/*PatchUserQueuesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchUserQueuesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserQueuesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesNotFound creates a PatchUserQueuesNotFound with default headers values
func NewPatchUserQueuesNotFound() *PatchUserQueuesNotFound {
	return &PatchUserQueuesNotFound{}
}

/*PatchUserQueuesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchUserQueuesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserQueuesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesRequestEntityTooLarge creates a PatchUserQueuesRequestEntityTooLarge with default headers values
func NewPatchUserQueuesRequestEntityTooLarge() *PatchUserQueuesRequestEntityTooLarge {
	return &PatchUserQueuesRequestEntityTooLarge{}
}

/*PatchUserQueuesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchUserQueuesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUserQueuesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesUnsupportedMediaType creates a PatchUserQueuesUnsupportedMediaType with default headers values
func NewPatchUserQueuesUnsupportedMediaType() *PatchUserQueuesUnsupportedMediaType {
	return &PatchUserQueuesUnsupportedMediaType{}
}

/*PatchUserQueuesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchUserQueuesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUserQueuesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesTooManyRequests creates a PatchUserQueuesTooManyRequests with default headers values
func NewPatchUserQueuesTooManyRequests() *PatchUserQueuesTooManyRequests {
	return &PatchUserQueuesTooManyRequests{}
}

/*PatchUserQueuesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchUserQueuesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUserQueuesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesInternalServerError creates a PatchUserQueuesInternalServerError with default headers values
func NewPatchUserQueuesInternalServerError() *PatchUserQueuesInternalServerError {
	return &PatchUserQueuesInternalServerError{}
}

/*PatchUserQueuesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchUserQueuesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUserQueuesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesServiceUnavailable creates a PatchUserQueuesServiceUnavailable with default headers values
func NewPatchUserQueuesServiceUnavailable() *PatchUserQueuesServiceUnavailable {
	return &PatchUserQueuesServiceUnavailable{}
}

/*PatchUserQueuesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchUserQueuesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUserQueuesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserQueuesGatewayTimeout creates a PatchUserQueuesGatewayTimeout with default headers values
func NewPatchUserQueuesGatewayTimeout() *PatchUserQueuesGatewayTimeout {
	return &PatchUserQueuesGatewayTimeout{}
}

/*PatchUserQueuesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchUserQueuesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchUserQueuesGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/queues][%d] patchUserQueuesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUserQueuesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserQueuesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
