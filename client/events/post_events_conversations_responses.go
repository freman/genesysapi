// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostEventsConversationsReader is a Reader for the PostEventsConversations structure.
type PostEventsConversationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEventsConversationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEventsConversationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostEventsConversationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostEventsConversationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostEventsConversationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostEventsConversationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostEventsConversationsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostEventsConversationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostEventsConversationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostEventsConversationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostEventsConversationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostEventsConversationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostEventsConversationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostEventsConversationsOK creates a PostEventsConversationsOK with default headers values
func NewPostEventsConversationsOK() *PostEventsConversationsOK {
	return &PostEventsConversationsOK{}
}

/*PostEventsConversationsOK handles this case with default header values.

successful operation
*/
type PostEventsConversationsOK struct {
	Payload *models.BatchEventResponse
}

func (o *PostEventsConversationsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsOK  %+v", 200, o.Payload)
}

func (o *PostEventsConversationsOK) GetPayload() *models.BatchEventResponse {
	return o.Payload
}

func (o *PostEventsConversationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BatchEventResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsBadRequest creates a PostEventsConversationsBadRequest with default headers values
func NewPostEventsConversationsBadRequest() *PostEventsConversationsBadRequest {
	return &PostEventsConversationsBadRequest{}
}

/*PostEventsConversationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostEventsConversationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsBadRequest  %+v", 400, o.Payload)
}

func (o *PostEventsConversationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsUnauthorized creates a PostEventsConversationsUnauthorized with default headers values
func NewPostEventsConversationsUnauthorized() *PostEventsConversationsUnauthorized {
	return &PostEventsConversationsUnauthorized{}
}

/*PostEventsConversationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostEventsConversationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostEventsConversationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsForbidden creates a PostEventsConversationsForbidden with default headers values
func NewPostEventsConversationsForbidden() *PostEventsConversationsForbidden {
	return &PostEventsConversationsForbidden{}
}

/*PostEventsConversationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostEventsConversationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsForbidden  %+v", 403, o.Payload)
}

func (o *PostEventsConversationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsNotFound creates a PostEventsConversationsNotFound with default headers values
func NewPostEventsConversationsNotFound() *PostEventsConversationsNotFound {
	return &PostEventsConversationsNotFound{}
}

/*PostEventsConversationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostEventsConversationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsNotFound  %+v", 404, o.Payload)
}

func (o *PostEventsConversationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsRequestTimeout creates a PostEventsConversationsRequestTimeout with default headers values
func NewPostEventsConversationsRequestTimeout() *PostEventsConversationsRequestTimeout {
	return &PostEventsConversationsRequestTimeout{}
}

/*PostEventsConversationsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostEventsConversationsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostEventsConversationsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsRequestEntityTooLarge creates a PostEventsConversationsRequestEntityTooLarge with default headers values
func NewPostEventsConversationsRequestEntityTooLarge() *PostEventsConversationsRequestEntityTooLarge {
	return &PostEventsConversationsRequestEntityTooLarge{}
}

/*PostEventsConversationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostEventsConversationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostEventsConversationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsUnsupportedMediaType creates a PostEventsConversationsUnsupportedMediaType with default headers values
func NewPostEventsConversationsUnsupportedMediaType() *PostEventsConversationsUnsupportedMediaType {
	return &PostEventsConversationsUnsupportedMediaType{}
}

/*PostEventsConversationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostEventsConversationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostEventsConversationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsTooManyRequests creates a PostEventsConversationsTooManyRequests with default headers values
func NewPostEventsConversationsTooManyRequests() *PostEventsConversationsTooManyRequests {
	return &PostEventsConversationsTooManyRequests{}
}

/*PostEventsConversationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostEventsConversationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostEventsConversationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsInternalServerError creates a PostEventsConversationsInternalServerError with default headers values
func NewPostEventsConversationsInternalServerError() *PostEventsConversationsInternalServerError {
	return &PostEventsConversationsInternalServerError{}
}

/*PostEventsConversationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostEventsConversationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostEventsConversationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsServiceUnavailable creates a PostEventsConversationsServiceUnavailable with default headers values
func NewPostEventsConversationsServiceUnavailable() *PostEventsConversationsServiceUnavailable {
	return &PostEventsConversationsServiceUnavailable{}
}

/*PostEventsConversationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostEventsConversationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostEventsConversationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEventsConversationsGatewayTimeout creates a PostEventsConversationsGatewayTimeout with default headers values
func NewPostEventsConversationsGatewayTimeout() *PostEventsConversationsGatewayTimeout {
	return &PostEventsConversationsGatewayTimeout{}
}

/*PostEventsConversationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostEventsConversationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostEventsConversationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/events/conversations][%d] postEventsConversationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostEventsConversationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostEventsConversationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
