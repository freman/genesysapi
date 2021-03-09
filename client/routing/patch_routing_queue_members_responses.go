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

// PatchRoutingQueueMembersReader is a Reader for the PatchRoutingQueueMembers structure.
type PatchRoutingQueueMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingQueueMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingQueueMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingQueueMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingQueueMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingQueueMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingQueueMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingQueueMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingQueueMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingQueueMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingQueueMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingQueueMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingQueueMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingQueueMembersOK creates a PatchRoutingQueueMembersOK with default headers values
func NewPatchRoutingQueueMembersOK() *PatchRoutingQueueMembersOK {
	return &PatchRoutingQueueMembersOK{}
}

/*PatchRoutingQueueMembersOK handles this case with default header values.

successful operation
*/
type PatchRoutingQueueMembersOK struct {
	Payload *models.QueueMemberEntityListing
}

func (o *PatchRoutingQueueMembersOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingQueueMembersOK) GetPayload() *models.QueueMemberEntityListing {
	return o.Payload
}

func (o *PatchRoutingQueueMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueMemberEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersBadRequest creates a PatchRoutingQueueMembersBadRequest with default headers values
func NewPatchRoutingQueueMembersBadRequest() *PatchRoutingQueueMembersBadRequest {
	return &PatchRoutingQueueMembersBadRequest{}
}

/*PatchRoutingQueueMembersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingQueueMembersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingQueueMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersUnauthorized creates a PatchRoutingQueueMembersUnauthorized with default headers values
func NewPatchRoutingQueueMembersUnauthorized() *PatchRoutingQueueMembersUnauthorized {
	return &PatchRoutingQueueMembersUnauthorized{}
}

/*PatchRoutingQueueMembersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingQueueMembersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingQueueMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersForbidden creates a PatchRoutingQueueMembersForbidden with default headers values
func NewPatchRoutingQueueMembersForbidden() *PatchRoutingQueueMembersForbidden {
	return &PatchRoutingQueueMembersForbidden{}
}

/*PatchRoutingQueueMembersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingQueueMembersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingQueueMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersNotFound creates a PatchRoutingQueueMembersNotFound with default headers values
func NewPatchRoutingQueueMembersNotFound() *PatchRoutingQueueMembersNotFound {
	return &PatchRoutingQueueMembersNotFound{}
}

/*PatchRoutingQueueMembersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingQueueMembersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingQueueMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersRequestEntityTooLarge creates a PatchRoutingQueueMembersRequestEntityTooLarge with default headers values
func NewPatchRoutingQueueMembersRequestEntityTooLarge() *PatchRoutingQueueMembersRequestEntityTooLarge {
	return &PatchRoutingQueueMembersRequestEntityTooLarge{}
}

/*PatchRoutingQueueMembersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingQueueMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingQueueMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersUnsupportedMediaType creates a PatchRoutingQueueMembersUnsupportedMediaType with default headers values
func NewPatchRoutingQueueMembersUnsupportedMediaType() *PatchRoutingQueueMembersUnsupportedMediaType {
	return &PatchRoutingQueueMembersUnsupportedMediaType{}
}

/*PatchRoutingQueueMembersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingQueueMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingQueueMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersTooManyRequests creates a PatchRoutingQueueMembersTooManyRequests with default headers values
func NewPatchRoutingQueueMembersTooManyRequests() *PatchRoutingQueueMembersTooManyRequests {
	return &PatchRoutingQueueMembersTooManyRequests{}
}

/*PatchRoutingQueueMembersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchRoutingQueueMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingQueueMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersInternalServerError creates a PatchRoutingQueueMembersInternalServerError with default headers values
func NewPatchRoutingQueueMembersInternalServerError() *PatchRoutingQueueMembersInternalServerError {
	return &PatchRoutingQueueMembersInternalServerError{}
}

/*PatchRoutingQueueMembersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingQueueMembersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingQueueMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersServiceUnavailable creates a PatchRoutingQueueMembersServiceUnavailable with default headers values
func NewPatchRoutingQueueMembersServiceUnavailable() *PatchRoutingQueueMembersServiceUnavailable {
	return &PatchRoutingQueueMembersServiceUnavailable{}
}

/*PatchRoutingQueueMembersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingQueueMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingQueueMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMembersGatewayTimeout creates a PatchRoutingQueueMembersGatewayTimeout with default headers values
func NewPatchRoutingQueueMembersGatewayTimeout() *PatchRoutingQueueMembersGatewayTimeout {
	return &PatchRoutingQueueMembersGatewayTimeout{}
}

/*PatchRoutingQueueMembersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingQueueMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members][%d] patchRoutingQueueMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingQueueMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
