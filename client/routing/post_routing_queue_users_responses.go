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

// PostRoutingQueueUsersReader is a Reader for the PostRoutingQueueUsers structure.
type PostRoutingQueueUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingQueueUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostRoutingQueueUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingQueueUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingQueueUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingQueueUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingQueueUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingQueueUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingQueueUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingQueueUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingQueueUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingQueueUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingQueueUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostRoutingQueueUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostRoutingQueueUsersBadRequest creates a PostRoutingQueueUsersBadRequest with default headers values
func NewPostRoutingQueueUsersBadRequest() *PostRoutingQueueUsersBadRequest {
	return &PostRoutingQueueUsersBadRequest{}
}

/*PostRoutingQueueUsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingQueueUsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingQueueUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersUnauthorized creates a PostRoutingQueueUsersUnauthorized with default headers values
func NewPostRoutingQueueUsersUnauthorized() *PostRoutingQueueUsersUnauthorized {
	return &PostRoutingQueueUsersUnauthorized{}
}

/*PostRoutingQueueUsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingQueueUsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingQueueUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersForbidden creates a PostRoutingQueueUsersForbidden with default headers values
func NewPostRoutingQueueUsersForbidden() *PostRoutingQueueUsersForbidden {
	return &PostRoutingQueueUsersForbidden{}
}

/*PostRoutingQueueUsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingQueueUsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingQueueUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersNotFound creates a PostRoutingQueueUsersNotFound with default headers values
func NewPostRoutingQueueUsersNotFound() *PostRoutingQueueUsersNotFound {
	return &PostRoutingQueueUsersNotFound{}
}

/*PostRoutingQueueUsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostRoutingQueueUsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingQueueUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersRequestTimeout creates a PostRoutingQueueUsersRequestTimeout with default headers values
func NewPostRoutingQueueUsersRequestTimeout() *PostRoutingQueueUsersRequestTimeout {
	return &PostRoutingQueueUsersRequestTimeout{}
}

/*PostRoutingQueueUsersRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingQueueUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingQueueUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersRequestEntityTooLarge creates a PostRoutingQueueUsersRequestEntityTooLarge with default headers values
func NewPostRoutingQueueUsersRequestEntityTooLarge() *PostRoutingQueueUsersRequestEntityTooLarge {
	return &PostRoutingQueueUsersRequestEntityTooLarge{}
}

/*PostRoutingQueueUsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostRoutingQueueUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingQueueUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersUnsupportedMediaType creates a PostRoutingQueueUsersUnsupportedMediaType with default headers values
func NewPostRoutingQueueUsersUnsupportedMediaType() *PostRoutingQueueUsersUnsupportedMediaType {
	return &PostRoutingQueueUsersUnsupportedMediaType{}
}

/*PostRoutingQueueUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingQueueUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingQueueUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersTooManyRequests creates a PostRoutingQueueUsersTooManyRequests with default headers values
func NewPostRoutingQueueUsersTooManyRequests() *PostRoutingQueueUsersTooManyRequests {
	return &PostRoutingQueueUsersTooManyRequests{}
}

/*PostRoutingQueueUsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingQueueUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingQueueUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersInternalServerError creates a PostRoutingQueueUsersInternalServerError with default headers values
func NewPostRoutingQueueUsersInternalServerError() *PostRoutingQueueUsersInternalServerError {
	return &PostRoutingQueueUsersInternalServerError{}
}

/*PostRoutingQueueUsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingQueueUsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingQueueUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersServiceUnavailable creates a PostRoutingQueueUsersServiceUnavailable with default headers values
func NewPostRoutingQueueUsersServiceUnavailable() *PostRoutingQueueUsersServiceUnavailable {
	return &PostRoutingQueueUsersServiceUnavailable{}
}

/*PostRoutingQueueUsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingQueueUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingQueueUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersGatewayTimeout creates a PostRoutingQueueUsersGatewayTimeout with default headers values
func NewPostRoutingQueueUsersGatewayTimeout() *PostRoutingQueueUsersGatewayTimeout {
	return &PostRoutingQueueUsersGatewayTimeout{}
}

/*PostRoutingQueueUsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostRoutingQueueUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostRoutingQueueUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingQueueUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingQueueUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingQueueUsersDefault creates a PostRoutingQueueUsersDefault with default headers values
func NewPostRoutingQueueUsersDefault(code int) *PostRoutingQueueUsersDefault {
	return &PostRoutingQueueUsersDefault{
		_statusCode: code,
	}
}

/*PostRoutingQueueUsersDefault handles this case with default header values.

successful operation
*/
type PostRoutingQueueUsersDefault struct {
	_statusCode int
}

// Code gets the status code for the post routing queue users default response
func (o *PostRoutingQueueUsersDefault) Code() int {
	return o._statusCode
}

func (o *PostRoutingQueueUsersDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/queues/{queueId}/users][%d] postRoutingQueueUsers default ", o._statusCode)
}

func (o *PostRoutingQueueUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
