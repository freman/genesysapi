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

// DeleteRoutingQueueMemberReader is a Reader for the DeleteRoutingQueueMember structure.
type DeleteRoutingQueueMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingQueueMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingQueueMemberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingQueueMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingQueueMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingQueueMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingQueueMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingQueueMemberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingQueueMemberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingQueueMemberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingQueueMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingQueueMemberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingQueueMemberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingQueueMemberNoContent creates a DeleteRoutingQueueMemberNoContent with default headers values
func NewDeleteRoutingQueueMemberNoContent() *DeleteRoutingQueueMemberNoContent {
	return &DeleteRoutingQueueMemberNoContent{}
}

/*DeleteRoutingQueueMemberNoContent handles this case with default header values.

Deleted.
*/
type DeleteRoutingQueueMemberNoContent struct {
}

func (o *DeleteRoutingQueueMemberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNoContent ", 204)
}

func (o *DeleteRoutingQueueMemberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingQueueMemberBadRequest creates a DeleteRoutingQueueMemberBadRequest with default headers values
func NewDeleteRoutingQueueMemberBadRequest() *DeleteRoutingQueueMemberBadRequest {
	return &DeleteRoutingQueueMemberBadRequest{}
}

/*DeleteRoutingQueueMemberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingQueueMemberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingQueueMemberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberUnauthorized creates a DeleteRoutingQueueMemberUnauthorized with default headers values
func NewDeleteRoutingQueueMemberUnauthorized() *DeleteRoutingQueueMemberUnauthorized {
	return &DeleteRoutingQueueMemberUnauthorized{}
}

/*DeleteRoutingQueueMemberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingQueueMemberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberForbidden creates a DeleteRoutingQueueMemberForbidden with default headers values
func NewDeleteRoutingQueueMemberForbidden() *DeleteRoutingQueueMemberForbidden {
	return &DeleteRoutingQueueMemberForbidden{}
}

/*DeleteRoutingQueueMemberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingQueueMemberForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingQueueMemberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberNotFound creates a DeleteRoutingQueueMemberNotFound with default headers values
func NewDeleteRoutingQueueMemberNotFound() *DeleteRoutingQueueMemberNotFound {
	return &DeleteRoutingQueueMemberNotFound{}
}

/*DeleteRoutingQueueMemberNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRoutingQueueMemberNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingQueueMemberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberRequestEntityTooLarge creates a DeleteRoutingQueueMemberRequestEntityTooLarge with default headers values
func NewDeleteRoutingQueueMemberRequestEntityTooLarge() *DeleteRoutingQueueMemberRequestEntityTooLarge {
	return &DeleteRoutingQueueMemberRequestEntityTooLarge{}
}

/*DeleteRoutingQueueMemberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRoutingQueueMemberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberUnsupportedMediaType creates a DeleteRoutingQueueMemberUnsupportedMediaType with default headers values
func NewDeleteRoutingQueueMemberUnsupportedMediaType() *DeleteRoutingQueueMemberUnsupportedMediaType {
	return &DeleteRoutingQueueMemberUnsupportedMediaType{}
}

/*DeleteRoutingQueueMemberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingQueueMemberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberTooManyRequests creates a DeleteRoutingQueueMemberTooManyRequests with default headers values
func NewDeleteRoutingQueueMemberTooManyRequests() *DeleteRoutingQueueMemberTooManyRequests {
	return &DeleteRoutingQueueMemberTooManyRequests{}
}

/*DeleteRoutingQueueMemberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRoutingQueueMemberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingQueueMemberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberInternalServerError creates a DeleteRoutingQueueMemberInternalServerError with default headers values
func NewDeleteRoutingQueueMemberInternalServerError() *DeleteRoutingQueueMemberInternalServerError {
	return &DeleteRoutingQueueMemberInternalServerError{}
}

/*DeleteRoutingQueueMemberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingQueueMemberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingQueueMemberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberServiceUnavailable creates a DeleteRoutingQueueMemberServiceUnavailable with default headers values
func NewDeleteRoutingQueueMemberServiceUnavailable() *DeleteRoutingQueueMemberServiceUnavailable {
	return &DeleteRoutingQueueMemberServiceUnavailable{}
}

/*DeleteRoutingQueueMemberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingQueueMemberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueMemberGatewayTimeout creates a DeleteRoutingQueueMemberGatewayTimeout with default headers values
func NewDeleteRoutingQueueMemberGatewayTimeout() *DeleteRoutingQueueMemberGatewayTimeout {
	return &DeleteRoutingQueueMemberGatewayTimeout{}
}

/*DeleteRoutingQueueMemberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRoutingQueueMemberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}/members/{memberId}][%d] deleteRoutingQueueMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueMemberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}