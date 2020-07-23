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

// PostConversationsChatCommunicationTypingReader is a Reader for the PostConversationsChatCommunicationTyping structure.
type PostConversationsChatCommunicationTypingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsChatCommunicationTypingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsChatCommunicationTypingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsChatCommunicationTypingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsChatCommunicationTypingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsChatCommunicationTypingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsChatCommunicationTypingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsChatCommunicationTypingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsChatCommunicationTypingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsChatCommunicationTypingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsChatCommunicationTypingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsChatCommunicationTypingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsChatCommunicationTypingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsChatCommunicationTypingOK creates a PostConversationsChatCommunicationTypingOK with default headers values
func NewPostConversationsChatCommunicationTypingOK() *PostConversationsChatCommunicationTypingOK {
	return &PostConversationsChatCommunicationTypingOK{}
}

/*PostConversationsChatCommunicationTypingOK handles this case with default header values.

successful operation
*/
type PostConversationsChatCommunicationTypingOK struct {
	Payload *models.WebChatTyping
}

func (o *PostConversationsChatCommunicationTypingOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingOK  %+v", 200, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingOK) GetPayload() *models.WebChatTyping {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebChatTyping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingBadRequest creates a PostConversationsChatCommunicationTypingBadRequest with default headers values
func NewPostConversationsChatCommunicationTypingBadRequest() *PostConversationsChatCommunicationTypingBadRequest {
	return &PostConversationsChatCommunicationTypingBadRequest{}
}

/*PostConversationsChatCommunicationTypingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsChatCommunicationTypingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingUnauthorized creates a PostConversationsChatCommunicationTypingUnauthorized with default headers values
func NewPostConversationsChatCommunicationTypingUnauthorized() *PostConversationsChatCommunicationTypingUnauthorized {
	return &PostConversationsChatCommunicationTypingUnauthorized{}
}

/*PostConversationsChatCommunicationTypingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsChatCommunicationTypingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingForbidden creates a PostConversationsChatCommunicationTypingForbidden with default headers values
func NewPostConversationsChatCommunicationTypingForbidden() *PostConversationsChatCommunicationTypingForbidden {
	return &PostConversationsChatCommunicationTypingForbidden{}
}

/*PostConversationsChatCommunicationTypingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsChatCommunicationTypingForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingNotFound creates a PostConversationsChatCommunicationTypingNotFound with default headers values
func NewPostConversationsChatCommunicationTypingNotFound() *PostConversationsChatCommunicationTypingNotFound {
	return &PostConversationsChatCommunicationTypingNotFound{}
}

/*PostConversationsChatCommunicationTypingNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsChatCommunicationTypingNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingRequestEntityTooLarge creates a PostConversationsChatCommunicationTypingRequestEntityTooLarge with default headers values
func NewPostConversationsChatCommunicationTypingRequestEntityTooLarge() *PostConversationsChatCommunicationTypingRequestEntityTooLarge {
	return &PostConversationsChatCommunicationTypingRequestEntityTooLarge{}
}

/*PostConversationsChatCommunicationTypingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsChatCommunicationTypingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingUnsupportedMediaType creates a PostConversationsChatCommunicationTypingUnsupportedMediaType with default headers values
func NewPostConversationsChatCommunicationTypingUnsupportedMediaType() *PostConversationsChatCommunicationTypingUnsupportedMediaType {
	return &PostConversationsChatCommunicationTypingUnsupportedMediaType{}
}

/*PostConversationsChatCommunicationTypingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsChatCommunicationTypingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingTooManyRequests creates a PostConversationsChatCommunicationTypingTooManyRequests with default headers values
func NewPostConversationsChatCommunicationTypingTooManyRequests() *PostConversationsChatCommunicationTypingTooManyRequests {
	return &PostConversationsChatCommunicationTypingTooManyRequests{}
}

/*PostConversationsChatCommunicationTypingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsChatCommunicationTypingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingInternalServerError creates a PostConversationsChatCommunicationTypingInternalServerError with default headers values
func NewPostConversationsChatCommunicationTypingInternalServerError() *PostConversationsChatCommunicationTypingInternalServerError {
	return &PostConversationsChatCommunicationTypingInternalServerError{}
}

/*PostConversationsChatCommunicationTypingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsChatCommunicationTypingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingServiceUnavailable creates a PostConversationsChatCommunicationTypingServiceUnavailable with default headers values
func NewPostConversationsChatCommunicationTypingServiceUnavailable() *PostConversationsChatCommunicationTypingServiceUnavailable {
	return &PostConversationsChatCommunicationTypingServiceUnavailable{}
}

/*PostConversationsChatCommunicationTypingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsChatCommunicationTypingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatCommunicationTypingGatewayTimeout creates a PostConversationsChatCommunicationTypingGatewayTimeout with default headers values
func NewPostConversationsChatCommunicationTypingGatewayTimeout() *PostConversationsChatCommunicationTypingGatewayTimeout {
	return &PostConversationsChatCommunicationTypingGatewayTimeout{}
}

/*PostConversationsChatCommunicationTypingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsChatCommunicationTypingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatCommunicationTypingGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/communications/{communicationId}/typing][%d] postConversationsChatCommunicationTypingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatCommunicationTypingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatCommunicationTypingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
