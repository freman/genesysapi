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

// PatchRoutingQueueUserReader is a Reader for the PatchRoutingQueueUser structure.
type PatchRoutingQueueUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingQueueUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingQueueUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPatchRoutingQueueUserAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingQueueUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingQueueUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingQueueUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingQueueUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingQueueUserRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingQueueUserUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingQueueUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingQueueUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingQueueUserServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingQueueUserGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingQueueUserOK creates a PatchRoutingQueueUserOK with default headers values
func NewPatchRoutingQueueUserOK() *PatchRoutingQueueUserOK {
	return &PatchRoutingQueueUserOK{}
}

/*PatchRoutingQueueUserOK handles this case with default header values.

successful operation
*/
type PatchRoutingQueueUserOK struct {
	Payload *models.QueueMember
}

func (o *PatchRoutingQueueUserOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingQueueUserOK) GetPayload() *models.QueueMember {
	return o.Payload
}

func (o *PatchRoutingQueueUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserAccepted creates a PatchRoutingQueueUserAccepted with default headers values
func NewPatchRoutingQueueUserAccepted() *PatchRoutingQueueUserAccepted {
	return &PatchRoutingQueueUserAccepted{}
}

/*PatchRoutingQueueUserAccepted handles this case with default header values.

User update has been accepted
*/
type PatchRoutingQueueUserAccepted struct {
}

func (o *PatchRoutingQueueUserAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserAccepted ", 202)
}

func (o *PatchRoutingQueueUserAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchRoutingQueueUserBadRequest creates a PatchRoutingQueueUserBadRequest with default headers values
func NewPatchRoutingQueueUserBadRequest() *PatchRoutingQueueUserBadRequest {
	return &PatchRoutingQueueUserBadRequest{}
}

/*PatchRoutingQueueUserBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingQueueUserBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingQueueUserBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserUnauthorized creates a PatchRoutingQueueUserUnauthorized with default headers values
func NewPatchRoutingQueueUserUnauthorized() *PatchRoutingQueueUserUnauthorized {
	return &PatchRoutingQueueUserUnauthorized{}
}

/*PatchRoutingQueueUserUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingQueueUserUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingQueueUserUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserForbidden creates a PatchRoutingQueueUserForbidden with default headers values
func NewPatchRoutingQueueUserForbidden() *PatchRoutingQueueUserForbidden {
	return &PatchRoutingQueueUserForbidden{}
}

/*PatchRoutingQueueUserForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingQueueUserForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingQueueUserForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserNotFound creates a PatchRoutingQueueUserNotFound with default headers values
func NewPatchRoutingQueueUserNotFound() *PatchRoutingQueueUserNotFound {
	return &PatchRoutingQueueUserNotFound{}
}

/*PatchRoutingQueueUserNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingQueueUserNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingQueueUserNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserRequestEntityTooLarge creates a PatchRoutingQueueUserRequestEntityTooLarge with default headers values
func NewPatchRoutingQueueUserRequestEntityTooLarge() *PatchRoutingQueueUserRequestEntityTooLarge {
	return &PatchRoutingQueueUserRequestEntityTooLarge{}
}

/*PatchRoutingQueueUserRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingQueueUserRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingQueueUserRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserUnsupportedMediaType creates a PatchRoutingQueueUserUnsupportedMediaType with default headers values
func NewPatchRoutingQueueUserUnsupportedMediaType() *PatchRoutingQueueUserUnsupportedMediaType {
	return &PatchRoutingQueueUserUnsupportedMediaType{}
}

/*PatchRoutingQueueUserUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingQueueUserUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingQueueUserUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserTooManyRequests creates a PatchRoutingQueueUserTooManyRequests with default headers values
func NewPatchRoutingQueueUserTooManyRequests() *PatchRoutingQueueUserTooManyRequests {
	return &PatchRoutingQueueUserTooManyRequests{}
}

/*PatchRoutingQueueUserTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchRoutingQueueUserTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingQueueUserTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserInternalServerError creates a PatchRoutingQueueUserInternalServerError with default headers values
func NewPatchRoutingQueueUserInternalServerError() *PatchRoutingQueueUserInternalServerError {
	return &PatchRoutingQueueUserInternalServerError{}
}

/*PatchRoutingQueueUserInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingQueueUserInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingQueueUserInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserServiceUnavailable creates a PatchRoutingQueueUserServiceUnavailable with default headers values
func NewPatchRoutingQueueUserServiceUnavailable() *PatchRoutingQueueUserServiceUnavailable {
	return &PatchRoutingQueueUserServiceUnavailable{}
}

/*PatchRoutingQueueUserServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingQueueUserServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingQueueUserServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingQueueUserGatewayTimeout creates a PatchRoutingQueueUserGatewayTimeout with default headers values
func NewPatchRoutingQueueUserGatewayTimeout() *PatchRoutingQueueUserGatewayTimeout {
	return &PatchRoutingQueueUserGatewayTimeout{}
}

/*PatchRoutingQueueUserGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingQueueUserGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingQueueUserGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/queues/{queueId}/users/{memberId}][%d] patchRoutingQueueUserGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingQueueUserGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingQueueUserGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
