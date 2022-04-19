// Code generated by go-swagger; DO NOT EDIT.

package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostConversationsMessagesReader is a Reader for the PostConversationsMessages structure.
type PostConversationsMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsMessagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessagesOK creates a PostConversationsMessagesOK with default headers values
func NewPostConversationsMessagesOK() *PostConversationsMessagesOK {
	return &PostConversationsMessagesOK{}
}

/*PostConversationsMessagesOK handles this case with default header values.

successful operation
*/
type PostConversationsMessagesOK struct {
	Payload *models.MessageConversation
}

func (o *PostConversationsMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessagesOK) GetPayload() *models.MessageConversation {
	return o.Payload
}

func (o *PostConversationsMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageConversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesBadRequest creates a PostConversationsMessagesBadRequest with default headers values
func NewPostConversationsMessagesBadRequest() *PostConversationsMessagesBadRequest {
	return &PostConversationsMessagesBadRequest{}
}

/*PostConversationsMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesUnauthorized creates a PostConversationsMessagesUnauthorized with default headers values
func NewPostConversationsMessagesUnauthorized() *PostConversationsMessagesUnauthorized {
	return &PostConversationsMessagesUnauthorized{}
}

/*PostConversationsMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesForbidden creates a PostConversationsMessagesForbidden with default headers values
func NewPostConversationsMessagesForbidden() *PostConversationsMessagesForbidden {
	return &PostConversationsMessagesForbidden{}
}

/*PostConversationsMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesNotFound creates a PostConversationsMessagesNotFound with default headers values
func NewPostConversationsMessagesNotFound() *PostConversationsMessagesNotFound {
	return &PostConversationsMessagesNotFound{}
}

/*PostConversationsMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesRequestTimeout creates a PostConversationsMessagesRequestTimeout with default headers values
func NewPostConversationsMessagesRequestTimeout() *PostConversationsMessagesRequestTimeout {
	return &PostConversationsMessagesRequestTimeout{}
}

/*PostConversationsMessagesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsMessagesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsMessagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesRequestEntityTooLarge creates a PostConversationsMessagesRequestEntityTooLarge with default headers values
func NewPostConversationsMessagesRequestEntityTooLarge() *PostConversationsMessagesRequestEntityTooLarge {
	return &PostConversationsMessagesRequestEntityTooLarge{}
}

/*PostConversationsMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesUnsupportedMediaType creates a PostConversationsMessagesUnsupportedMediaType with default headers values
func NewPostConversationsMessagesUnsupportedMediaType() *PostConversationsMessagesUnsupportedMediaType {
	return &PostConversationsMessagesUnsupportedMediaType{}
}

/*PostConversationsMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesTooManyRequests creates a PostConversationsMessagesTooManyRequests with default headers values
func NewPostConversationsMessagesTooManyRequests() *PostConversationsMessagesTooManyRequests {
	return &PostConversationsMessagesTooManyRequests{}
}

/*PostConversationsMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesInternalServerError creates a PostConversationsMessagesInternalServerError with default headers values
func NewPostConversationsMessagesInternalServerError() *PostConversationsMessagesInternalServerError {
	return &PostConversationsMessagesInternalServerError{}
}

/*PostConversationsMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesServiceUnavailable creates a PostConversationsMessagesServiceUnavailable with default headers values
func NewPostConversationsMessagesServiceUnavailable() *PostConversationsMessagesServiceUnavailable {
	return &PostConversationsMessagesServiceUnavailable{}
}

/*PostConversationsMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagesGatewayTimeout creates a PostConversationsMessagesGatewayTimeout with default headers values
func NewPostConversationsMessagesGatewayTimeout() *PostConversationsMessagesGatewayTimeout {
	return &PostConversationsMessagesGatewayTimeout{}
}

/*PostConversationsMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages][%d] postConversationsMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
