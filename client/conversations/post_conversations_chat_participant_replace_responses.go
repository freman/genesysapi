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

// PostConversationsChatParticipantReplaceReader is a Reader for the PostConversationsChatParticipantReplace structure.
type PostConversationsChatParticipantReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsChatParticipantReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostConversationsChatParticipantReplaceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsChatParticipantReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsChatParticipantReplaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsChatParticipantReplaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsChatParticipantReplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsChatParticipantReplaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsChatParticipantReplaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsChatParticipantReplaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsChatParticipantReplaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsChatParticipantReplaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsChatParticipantReplaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsChatParticipantReplaceAccepted creates a PostConversationsChatParticipantReplaceAccepted with default headers values
func NewPostConversationsChatParticipantReplaceAccepted() *PostConversationsChatParticipantReplaceAccepted {
	return &PostConversationsChatParticipantReplaceAccepted{}
}

/*PostConversationsChatParticipantReplaceAccepted handles this case with default header values.

Accepted
*/
type PostConversationsChatParticipantReplaceAccepted struct {
}

func (o *PostConversationsChatParticipantReplaceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsChatParticipantReplaceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsChatParticipantReplaceBadRequest creates a PostConversationsChatParticipantReplaceBadRequest with default headers values
func NewPostConversationsChatParticipantReplaceBadRequest() *PostConversationsChatParticipantReplaceBadRequest {
	return &PostConversationsChatParticipantReplaceBadRequest{}
}

/*PostConversationsChatParticipantReplaceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsChatParticipantReplaceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceUnauthorized creates a PostConversationsChatParticipantReplaceUnauthorized with default headers values
func NewPostConversationsChatParticipantReplaceUnauthorized() *PostConversationsChatParticipantReplaceUnauthorized {
	return &PostConversationsChatParticipantReplaceUnauthorized{}
}

/*PostConversationsChatParticipantReplaceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsChatParticipantReplaceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceForbidden creates a PostConversationsChatParticipantReplaceForbidden with default headers values
func NewPostConversationsChatParticipantReplaceForbidden() *PostConversationsChatParticipantReplaceForbidden {
	return &PostConversationsChatParticipantReplaceForbidden{}
}

/*PostConversationsChatParticipantReplaceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsChatParticipantReplaceForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceNotFound creates a PostConversationsChatParticipantReplaceNotFound with default headers values
func NewPostConversationsChatParticipantReplaceNotFound() *PostConversationsChatParticipantReplaceNotFound {
	return &PostConversationsChatParticipantReplaceNotFound{}
}

/*PostConversationsChatParticipantReplaceNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsChatParticipantReplaceNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceRequestEntityTooLarge creates a PostConversationsChatParticipantReplaceRequestEntityTooLarge with default headers values
func NewPostConversationsChatParticipantReplaceRequestEntityTooLarge() *PostConversationsChatParticipantReplaceRequestEntityTooLarge {
	return &PostConversationsChatParticipantReplaceRequestEntityTooLarge{}
}

/*PostConversationsChatParticipantReplaceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsChatParticipantReplaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceUnsupportedMediaType creates a PostConversationsChatParticipantReplaceUnsupportedMediaType with default headers values
func NewPostConversationsChatParticipantReplaceUnsupportedMediaType() *PostConversationsChatParticipantReplaceUnsupportedMediaType {
	return &PostConversationsChatParticipantReplaceUnsupportedMediaType{}
}

/*PostConversationsChatParticipantReplaceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsChatParticipantReplaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceTooManyRequests creates a PostConversationsChatParticipantReplaceTooManyRequests with default headers values
func NewPostConversationsChatParticipantReplaceTooManyRequests() *PostConversationsChatParticipantReplaceTooManyRequests {
	return &PostConversationsChatParticipantReplaceTooManyRequests{}
}

/*PostConversationsChatParticipantReplaceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsChatParticipantReplaceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceInternalServerError creates a PostConversationsChatParticipantReplaceInternalServerError with default headers values
func NewPostConversationsChatParticipantReplaceInternalServerError() *PostConversationsChatParticipantReplaceInternalServerError {
	return &PostConversationsChatParticipantReplaceInternalServerError{}
}

/*PostConversationsChatParticipantReplaceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsChatParticipantReplaceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceServiceUnavailable creates a PostConversationsChatParticipantReplaceServiceUnavailable with default headers values
func NewPostConversationsChatParticipantReplaceServiceUnavailable() *PostConversationsChatParticipantReplaceServiceUnavailable {
	return &PostConversationsChatParticipantReplaceServiceUnavailable{}
}

/*PostConversationsChatParticipantReplaceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsChatParticipantReplaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsChatParticipantReplaceGatewayTimeout creates a PostConversationsChatParticipantReplaceGatewayTimeout with default headers values
func NewPostConversationsChatParticipantReplaceGatewayTimeout() *PostConversationsChatParticipantReplaceGatewayTimeout {
	return &PostConversationsChatParticipantReplaceGatewayTimeout{}
}

/*PostConversationsChatParticipantReplaceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsChatParticipantReplaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsChatParticipantReplaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}