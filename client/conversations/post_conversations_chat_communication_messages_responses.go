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

// PostConversationsChatCommunicationMessagesReader is a Reader for the PostConversationsChatCommunicationMessages structure.
type PostConversationsChatCommunicationMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsChatCommunicationMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsChatCommunicationMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsChatCommunicationMessagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsChatCommunicationMessagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsChatCommunicationMessagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsChatCommunicationMessagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsChatCommunicationMessagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsChatCommunicationMessagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsChatCommunicationMessagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsChatCommunicationMessagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsChatCommunicationMessagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsChatCommunicationMessagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsChatCommunicationMessagesOK creates a PostConversationsChatCommunicationMessagesOK with default headers values
func NewPostConversationsChatCommunicationMessagesOK() *PostConversationsChatCommunicationMessagesOK {
	return &PostConversationsChatCommunicationMessagesOK{}
}

/*PostConversationsChatCommunicationMessagesOK handles this case with default header values.

successful operation
*/
type PostConversationsChatCommunicationMessagesOK struct {
	Payload *models.WebChatMessage
}

func (o *PostConversationsChatCommunicationMessagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesOK  %+v", 200, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesOK) GetPayload() *models.WebChatMessage {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesBadRequest creates a PostConversationsChatCommunicationMessagesBadRequest with default headers values
func NewPostConversationsChatCommunicationMessagesBadRequest() *PostConversationsChatCommunicationMessagesBadRequest {
	return &PostConversationsChatCommunicationMessagesBadRequest{}
}

/*PostConversationsChatCommunicationMessagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsChatCommunicationMessagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesUnauthorized creates a PostConversationsChatCommunicationMessagesUnauthorized with default headers values
func NewPostConversationsChatCommunicationMessagesUnauthorized() *PostConversationsChatCommunicationMessagesUnauthorized {
	return &PostConversationsChatCommunicationMessagesUnauthorized{}
}

/*PostConversationsChatCommunicationMessagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsChatCommunicationMessagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesForbidden creates a PostConversationsChatCommunicationMessagesForbidden with default headers values
func NewPostConversationsChatCommunicationMessagesForbidden() *PostConversationsChatCommunicationMessagesForbidden {
	return &PostConversationsChatCommunicationMessagesForbidden{}
}

/*PostConversationsChatCommunicationMessagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsChatCommunicationMessagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesNotFound creates a PostConversationsChatCommunicationMessagesNotFound with default headers values
func NewPostConversationsChatCommunicationMessagesNotFound() *PostConversationsChatCommunicationMessagesNotFound {
	return &PostConversationsChatCommunicationMessagesNotFound{}
}

/*PostConversationsChatCommunicationMessagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsChatCommunicationMessagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesRequestEntityTooLarge creates a PostConversationsChatCommunicationMessagesRequestEntityTooLarge with default headers values
func NewPostConversationsChatCommunicationMessagesRequestEntityTooLarge() *PostConversationsChatCommunicationMessagesRequestEntityTooLarge {
	return &PostConversationsChatCommunicationMessagesRequestEntityTooLarge{}
}

/*PostConversationsChatCommunicationMessagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsChatCommunicationMessagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesUnsupportedMediaType creates a PostConversationsChatCommunicationMessagesUnsupportedMediaType with default headers values
func NewPostConversationsChatCommunicationMessagesUnsupportedMediaType() *PostConversationsChatCommunicationMessagesUnsupportedMediaType {
	return &PostConversationsChatCommunicationMessagesUnsupportedMediaType{}
}

/*PostConversationsChatCommunicationMessagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsChatCommunicationMessagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesTooManyRequests creates a PostConversationsChatCommunicationMessagesTooManyRequests with default headers values
func NewPostConversationsChatCommunicationMessagesTooManyRequests() *PostConversationsChatCommunicationMessagesTooManyRequests {
	return &PostConversationsChatCommunicationMessagesTooManyRequests{}
}

/*PostConversationsChatCommunicationMessagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsChatCommunicationMessagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesInternalServerError creates a PostConversationsChatCommunicationMessagesInternalServerError with default headers values
func NewPostConversationsChatCommunicationMessagesInternalServerError() *PostConversationsChatCommunicationMessagesInternalServerError {
	return &PostConversationsChatCommunicationMessagesInternalServerError{}
}

/*PostConversationsChatCommunicationMessagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsChatCommunicationMessagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesServiceUnavailable creates a PostConversationsChatCommunicationMessagesServiceUnavailable with default headers values
func NewPostConversationsChatCommunicationMessagesServiceUnavailable() *PostConversationsChatCommunicationMessagesServiceUnavailable {
	return &PostConversationsChatCommunicationMessagesServiceUnavailable{}
}

/*PostConversationsChatCommunicationMessagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsChatCommunicationMessagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationMessagesGatewayTimeout creates a PostConversationsChatCommunicationMessagesGatewayTimeout with default headers values
func NewPostConversationsChatCommunicationMessagesGatewayTimeout() *PostConversationsChatCommunicationMessagesGatewayTimeout {
	return &PostConversationsChatCommunicationMessagesGatewayTimeout{}
}

/*PostConversationsChatCommunicationMessagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsChatCommunicationMessagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationMessagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/messages][%d] postConversationsChatCommunicationMessagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatCommunicationMessagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationMessagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
