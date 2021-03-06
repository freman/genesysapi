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

// PatchRoutingQueueMemberReader is a Reader for the PatchRoutingQueueMember structure.
type PatchRoutingQueueMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingQueueMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingQueueMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPatchRoutingQueueMemberAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingQueueMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingQueueMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingQueueMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingQueueMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingQueueMemberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingQueueMemberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingQueueMemberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingQueueMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingQueueMemberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingQueueMemberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingQueueMemberOK creates a PatchRoutingQueueMemberOK with default headers values
func NewPatchRoutingQueueMemberOK() *PatchRoutingQueueMemberOK {
	return &PatchRoutingQueueMemberOK{}
}

/*PatchRoutingQueueMemberOK handles this case with default header values.

successful operation
*/
type PatchRoutingQueueMemberOK struct {
	Payload *models.QueueMember
}

func (o *PatchRoutingQueueMemberOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingQueueMemberOK) GetPayload() *models.QueueMember {
	return o.Payload
}

func (o *PatchRoutingQueueMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberAccepted creates a PatchRoutingQueueMemberAccepted with default headers values
func NewPatchRoutingQueueMemberAccepted() *PatchRoutingQueueMemberAccepted {
	return &PatchRoutingQueueMemberAccepted{}
}

/*PatchRoutingQueueMemberAccepted handles this case with default header values.

User update has been accepted
*/
type PatchRoutingQueueMemberAccepted struct {
}

func (o *PatchRoutingQueueMemberAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberAccepted ", 202)
}

func (o *PatchRoutingQueueMemberAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRoutingQueueMemberBadRequest creates a PatchRoutingQueueMemberBadRequest with default headers values
func NewPatchRoutingQueueMemberBadRequest() *PatchRoutingQueueMemberBadRequest {
	return &PatchRoutingQueueMemberBadRequest{}
}

/*PatchRoutingQueueMemberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingQueueMemberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingQueueMemberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberUnauthorized creates a PatchRoutingQueueMemberUnauthorized with default headers values
func NewPatchRoutingQueueMemberUnauthorized() *PatchRoutingQueueMemberUnauthorized {
	return &PatchRoutingQueueMemberUnauthorized{}
}

/*PatchRoutingQueueMemberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingQueueMemberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingQueueMemberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberForbidden creates a PatchRoutingQueueMemberForbidden with default headers values
func NewPatchRoutingQueueMemberForbidden() *PatchRoutingQueueMemberForbidden {
	return &PatchRoutingQueueMemberForbidden{}
}

/*PatchRoutingQueueMemberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingQueueMemberForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingQueueMemberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberNotFound creates a PatchRoutingQueueMemberNotFound with default headers values
func NewPatchRoutingQueueMemberNotFound() *PatchRoutingQueueMemberNotFound {
	return &PatchRoutingQueueMemberNotFound{}
}

/*PatchRoutingQueueMemberNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingQueueMemberNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingQueueMemberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberRequestEntityTooLarge creates a PatchRoutingQueueMemberRequestEntityTooLarge with default headers values
func NewPatchRoutingQueueMemberRequestEntityTooLarge() *PatchRoutingQueueMemberRequestEntityTooLarge {
	return &PatchRoutingQueueMemberRequestEntityTooLarge{}
}

/*PatchRoutingQueueMemberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingQueueMemberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingQueueMemberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberUnsupportedMediaType creates a PatchRoutingQueueMemberUnsupportedMediaType with default headers values
func NewPatchRoutingQueueMemberUnsupportedMediaType() *PatchRoutingQueueMemberUnsupportedMediaType {
	return &PatchRoutingQueueMemberUnsupportedMediaType{}
}

/*PatchRoutingQueueMemberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingQueueMemberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingQueueMemberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberTooManyRequests creates a PatchRoutingQueueMemberTooManyRequests with default headers values
func NewPatchRoutingQueueMemberTooManyRequests() *PatchRoutingQueueMemberTooManyRequests {
	return &PatchRoutingQueueMemberTooManyRequests{}
}

/*PatchRoutingQueueMemberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchRoutingQueueMemberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingQueueMemberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberInternalServerError creates a PatchRoutingQueueMemberInternalServerError with default headers values
func NewPatchRoutingQueueMemberInternalServerError() *PatchRoutingQueueMemberInternalServerError {
	return &PatchRoutingQueueMemberInternalServerError{}
}

/*PatchRoutingQueueMemberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingQueueMemberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingQueueMemberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberServiceUnavailable creates a PatchRoutingQueueMemberServiceUnavailable with default headers values
func NewPatchRoutingQueueMemberServiceUnavailable() *PatchRoutingQueueMemberServiceUnavailable {
	return &PatchRoutingQueueMemberServiceUnavailable{}
}

/*PatchRoutingQueueMemberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingQueueMemberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingQueueMemberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueMemberGatewayTimeout creates a PatchRoutingQueueMemberGatewayTimeout with default headers values
func NewPatchRoutingQueueMemberGatewayTimeout() *PatchRoutingQueueMemberGatewayTimeout {
	return &PatchRoutingQueueMemberGatewayTimeout{}
}

/*PatchRoutingQueueMemberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingQueueMemberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueMemberGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/members/{memberId}][%d] patchRoutingQueueMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingQueueMemberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueMemberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
