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

// PostRoutingQueueMembersReader is a Reader for the PostRoutingQueueMembers structure.
type PostRoutingQueueMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingQueueMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostRoutingQueueMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingQueueMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingQueueMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingQueueMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingQueueMembersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingQueueMembersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingQueueMembersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingQueueMembersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingQueueMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingQueueMembersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingQueueMembersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostRoutingQueueMembersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostRoutingQueueMembersBadRequest creates a PostRoutingQueueMembersBadRequest with default headers values
func NewPostRoutingQueueMembersBadRequest() *PostRoutingQueueMembersBadRequest {
	return &PostRoutingQueueMembersBadRequest{}
}

/*PostRoutingQueueMembersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingQueueMembersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingQueueMembersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersUnauthorized creates a PostRoutingQueueMembersUnauthorized with default headers values
func NewPostRoutingQueueMembersUnauthorized() *PostRoutingQueueMembersUnauthorized {
	return &PostRoutingQueueMembersUnauthorized{}
}

/*PostRoutingQueueMembersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingQueueMembersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingQueueMembersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersForbidden creates a PostRoutingQueueMembersForbidden with default headers values
func NewPostRoutingQueueMembersForbidden() *PostRoutingQueueMembersForbidden {
	return &PostRoutingQueueMembersForbidden{}
}

/*PostRoutingQueueMembersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingQueueMembersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingQueueMembersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersNotFound creates a PostRoutingQueueMembersNotFound with default headers values
func NewPostRoutingQueueMembersNotFound() *PostRoutingQueueMembersNotFound {
	return &PostRoutingQueueMembersNotFound{}
}

/*PostRoutingQueueMembersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRoutingQueueMembersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingQueueMembersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersRequestTimeout creates a PostRoutingQueueMembersRequestTimeout with default headers values
func NewPostRoutingQueueMembersRequestTimeout() *PostRoutingQueueMembersRequestTimeout {
	return &PostRoutingQueueMembersRequestTimeout{}
}

/*PostRoutingQueueMembersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingQueueMembersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingQueueMembersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersRequestEntityTooLarge creates a PostRoutingQueueMembersRequestEntityTooLarge with default headers values
func NewPostRoutingQueueMembersRequestEntityTooLarge() *PostRoutingQueueMembersRequestEntityTooLarge {
	return &PostRoutingQueueMembersRequestEntityTooLarge{}
}

/*PostRoutingQueueMembersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRoutingQueueMembersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingQueueMembersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersUnsupportedMediaType creates a PostRoutingQueueMembersUnsupportedMediaType with default headers values
func NewPostRoutingQueueMembersUnsupportedMediaType() *PostRoutingQueueMembersUnsupportedMediaType {
	return &PostRoutingQueueMembersUnsupportedMediaType{}
}

/*PostRoutingQueueMembersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingQueueMembersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingQueueMembersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersTooManyRequests creates a PostRoutingQueueMembersTooManyRequests with default headers values
func NewPostRoutingQueueMembersTooManyRequests() *PostRoutingQueueMembersTooManyRequests {
	return &PostRoutingQueueMembersTooManyRequests{}
}

/*PostRoutingQueueMembersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingQueueMembersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingQueueMembersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersInternalServerError creates a PostRoutingQueueMembersInternalServerError with default headers values
func NewPostRoutingQueueMembersInternalServerError() *PostRoutingQueueMembersInternalServerError {
	return &PostRoutingQueueMembersInternalServerError{}
}

/*PostRoutingQueueMembersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingQueueMembersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingQueueMembersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersServiceUnavailable creates a PostRoutingQueueMembersServiceUnavailable with default headers values
func NewPostRoutingQueueMembersServiceUnavailable() *PostRoutingQueueMembersServiceUnavailable {
	return &PostRoutingQueueMembersServiceUnavailable{}
}

/*PostRoutingQueueMembersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingQueueMembersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingQueueMembersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersGatewayTimeout creates a PostRoutingQueueMembersGatewayTimeout with default headers values
func NewPostRoutingQueueMembersGatewayTimeout() *PostRoutingQueueMembersGatewayTimeout {
	return &PostRoutingQueueMembersGatewayTimeout{}
}

/*PostRoutingQueueMembersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRoutingQueueMembersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueMembersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingQueueMembersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueMembersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueMembersDefault creates a PostRoutingQueueMembersDefault with default headers values
func NewPostRoutingQueueMembersDefault(code int) *PostRoutingQueueMembersDefault {
	return &PostRoutingQueueMembersDefault{
		_statusCode: code,
	}
}

/*PostRoutingQueueMembersDefault handles this case with default header values.

successful operation
*/
type PostRoutingQueueMembersDefault struct {
	_statusCode int
}

// Code gets the status code for the post routing queue members default response
func (o *PostRoutingQueueMembersDefault) Code() int {
	return o._statusCode
}

func (o *PostRoutingQueueMembersDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/members][%d] postRoutingQueueMembers default ", o._statusCode)
}

func (o *PostRoutingQueueMembersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
