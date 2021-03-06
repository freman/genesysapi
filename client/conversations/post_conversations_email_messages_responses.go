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

// PostConversationsEmailMessagesReader is a Reader for the PostConversationsEmailMessages structure.
type PostConversationsEmailMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsEmailMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsEmailMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsEmailMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsEmailMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsEmailMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsEmailMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostConversationsEmailMessagesConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsEmailMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsEmailMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsEmailMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsEmailMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsEmailMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsEmailMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsEmailMessagesOK creates a PostConversationsEmailMessagesOK with default headers values
func NewPostConversationsEmailMessagesOK() *PostConversationsEmailMessagesOK {
	return &PostConversationsEmailMessagesOK{}
}

/*PostConversationsEmailMessagesOK handles this case with default header values.

successful operation
*/
type PostConversationsEmailMessagesOK struct {
	Payload *models.EmailMessage
}

func (o *PostConversationsEmailMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsEmailMessagesOK) GetPayload() *models.EmailMessage {
	return o.Payload
}

func (o *PostConversationsEmailMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesBadRequest creates a PostConversationsEmailMessagesBadRequest with default headers values
func NewPostConversationsEmailMessagesBadRequest() *PostConversationsEmailMessagesBadRequest {
	return &PostConversationsEmailMessagesBadRequest{}
}

/*PostConversationsEmailMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsEmailMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsEmailMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesUnauthorized creates a PostConversationsEmailMessagesUnauthorized with default headers values
func NewPostConversationsEmailMessagesUnauthorized() *PostConversationsEmailMessagesUnauthorized {
	return &PostConversationsEmailMessagesUnauthorized{}
}

/*PostConversationsEmailMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsEmailMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsEmailMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesForbidden creates a PostConversationsEmailMessagesForbidden with default headers values
func NewPostConversationsEmailMessagesForbidden() *PostConversationsEmailMessagesForbidden {
	return &PostConversationsEmailMessagesForbidden{}
}

/*PostConversationsEmailMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsEmailMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsEmailMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesNotFound creates a PostConversationsEmailMessagesNotFound with default headers values
func NewPostConversationsEmailMessagesNotFound() *PostConversationsEmailMessagesNotFound {
	return &PostConversationsEmailMessagesNotFound{}
}

/*PostConversationsEmailMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsEmailMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsEmailMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesConflict creates a PostConversationsEmailMessagesConflict with default headers values
func NewPostConversationsEmailMessagesConflict() *PostConversationsEmailMessagesConflict {
	return &PostConversationsEmailMessagesConflict{}
}

/*PostConversationsEmailMessagesConflict handles this case with default header values.

Conflict
*/
type PostConversationsEmailMessagesConflict struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesConflict  %+v", 409, o.Payload)
}

func (o *PostConversationsEmailMessagesConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesRequestEntityTooLarge creates a PostConversationsEmailMessagesRequestEntityTooLarge with default headers values
func NewPostConversationsEmailMessagesRequestEntityTooLarge() *PostConversationsEmailMessagesRequestEntityTooLarge {
	return &PostConversationsEmailMessagesRequestEntityTooLarge{}
}

/*PostConversationsEmailMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsEmailMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesUnsupportedMediaType creates a PostConversationsEmailMessagesUnsupportedMediaType with default headers values
func NewPostConversationsEmailMessagesUnsupportedMediaType() *PostConversationsEmailMessagesUnsupportedMediaType {
	return &PostConversationsEmailMessagesUnsupportedMediaType{}
}

/*PostConversationsEmailMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsEmailMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesTooManyRequests creates a PostConversationsEmailMessagesTooManyRequests with default headers values
func NewPostConversationsEmailMessagesTooManyRequests() *PostConversationsEmailMessagesTooManyRequests {
	return &PostConversationsEmailMessagesTooManyRequests{}
}

/*PostConversationsEmailMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsEmailMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsEmailMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesInternalServerError creates a PostConversationsEmailMessagesInternalServerError with default headers values
func NewPostConversationsEmailMessagesInternalServerError() *PostConversationsEmailMessagesInternalServerError {
	return &PostConversationsEmailMessagesInternalServerError{}
}

/*PostConversationsEmailMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsEmailMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsEmailMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesServiceUnavailable creates a PostConversationsEmailMessagesServiceUnavailable with default headers values
func NewPostConversationsEmailMessagesServiceUnavailable() *PostConversationsEmailMessagesServiceUnavailable {
	return &PostConversationsEmailMessagesServiceUnavailable{}
}

/*PostConversationsEmailMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsEmailMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsEmailMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsEmailMessagesGatewayTimeout creates a PostConversationsEmailMessagesGatewayTimeout with default headers values
func NewPostConversationsEmailMessagesGatewayTimeout() *PostConversationsEmailMessagesGatewayTimeout {
	return &PostConversationsEmailMessagesGatewayTimeout{}
}

/*PostConversationsEmailMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsEmailMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsEmailMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/emails/{conversationId}/messages][%d] postConversationsEmailMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsEmailMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsEmailMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
