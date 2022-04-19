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

// PostConversationsMessageParticipantReplaceReader is a Reader for the PostConversationsMessageParticipantReplace structure.
type PostConversationsMessageParticipantReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessageParticipantReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostConversationsMessageParticipantReplaceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessageParticipantReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessageParticipantReplaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessageParticipantReplaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessageParticipantReplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsMessageParticipantReplaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessageParticipantReplaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessageParticipantReplaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessageParticipantReplaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessageParticipantReplaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessageParticipantReplaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessageParticipantReplaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessageParticipantReplaceAccepted creates a PostConversationsMessageParticipantReplaceAccepted with default headers values
func NewPostConversationsMessageParticipantReplaceAccepted() *PostConversationsMessageParticipantReplaceAccepted {
	return &PostConversationsMessageParticipantReplaceAccepted{}
}

/*PostConversationsMessageParticipantReplaceAccepted handles this case with default header values.

Accepted
*/
type PostConversationsMessageParticipantReplaceAccepted struct {
}

func (o *PostConversationsMessageParticipantReplaceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsMessageParticipantReplaceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsMessageParticipantReplaceBadRequest creates a PostConversationsMessageParticipantReplaceBadRequest with default headers values
func NewPostConversationsMessageParticipantReplaceBadRequest() *PostConversationsMessageParticipantReplaceBadRequest {
	return &PostConversationsMessageParticipantReplaceBadRequest{}
}

/*PostConversationsMessageParticipantReplaceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessageParticipantReplaceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceUnauthorized creates a PostConversationsMessageParticipantReplaceUnauthorized with default headers values
func NewPostConversationsMessageParticipantReplaceUnauthorized() *PostConversationsMessageParticipantReplaceUnauthorized {
	return &PostConversationsMessageParticipantReplaceUnauthorized{}
}

/*PostConversationsMessageParticipantReplaceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessageParticipantReplaceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceForbidden creates a PostConversationsMessageParticipantReplaceForbidden with default headers values
func NewPostConversationsMessageParticipantReplaceForbidden() *PostConversationsMessageParticipantReplaceForbidden {
	return &PostConversationsMessageParticipantReplaceForbidden{}
}

/*PostConversationsMessageParticipantReplaceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessageParticipantReplaceForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceNotFound creates a PostConversationsMessageParticipantReplaceNotFound with default headers values
func NewPostConversationsMessageParticipantReplaceNotFound() *PostConversationsMessageParticipantReplaceNotFound {
	return &PostConversationsMessageParticipantReplaceNotFound{}
}

/*PostConversationsMessageParticipantReplaceNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessageParticipantReplaceNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceRequestTimeout creates a PostConversationsMessageParticipantReplaceRequestTimeout with default headers values
func NewPostConversationsMessageParticipantReplaceRequestTimeout() *PostConversationsMessageParticipantReplaceRequestTimeout {
	return &PostConversationsMessageParticipantReplaceRequestTimeout{}
}

/*PostConversationsMessageParticipantReplaceRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsMessageParticipantReplaceRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceRequestEntityTooLarge creates a PostConversationsMessageParticipantReplaceRequestEntityTooLarge with default headers values
func NewPostConversationsMessageParticipantReplaceRequestEntityTooLarge() *PostConversationsMessageParticipantReplaceRequestEntityTooLarge {
	return &PostConversationsMessageParticipantReplaceRequestEntityTooLarge{}
}

/*PostConversationsMessageParticipantReplaceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsMessageParticipantReplaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceUnsupportedMediaType creates a PostConversationsMessageParticipantReplaceUnsupportedMediaType with default headers values
func NewPostConversationsMessageParticipantReplaceUnsupportedMediaType() *PostConversationsMessageParticipantReplaceUnsupportedMediaType {
	return &PostConversationsMessageParticipantReplaceUnsupportedMediaType{}
}

/*PostConversationsMessageParticipantReplaceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessageParticipantReplaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceTooManyRequests creates a PostConversationsMessageParticipantReplaceTooManyRequests with default headers values
func NewPostConversationsMessageParticipantReplaceTooManyRequests() *PostConversationsMessageParticipantReplaceTooManyRequests {
	return &PostConversationsMessageParticipantReplaceTooManyRequests{}
}

/*PostConversationsMessageParticipantReplaceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsMessageParticipantReplaceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceInternalServerError creates a PostConversationsMessageParticipantReplaceInternalServerError with default headers values
func NewPostConversationsMessageParticipantReplaceInternalServerError() *PostConversationsMessageParticipantReplaceInternalServerError {
	return &PostConversationsMessageParticipantReplaceInternalServerError{}
}

/*PostConversationsMessageParticipantReplaceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessageParticipantReplaceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceServiceUnavailable creates a PostConversationsMessageParticipantReplaceServiceUnavailable with default headers values
func NewPostConversationsMessageParticipantReplaceServiceUnavailable() *PostConversationsMessageParticipantReplaceServiceUnavailable {
	return &PostConversationsMessageParticipantReplaceServiceUnavailable{}
}

/*PostConversationsMessageParticipantReplaceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessageParticipantReplaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessageParticipantReplaceGatewayTimeout creates a PostConversationsMessageParticipantReplaceGatewayTimeout with default headers values
func NewPostConversationsMessageParticipantReplaceGatewayTimeout() *PostConversationsMessageParticipantReplaceGatewayTimeout {
	return &PostConversationsMessageParticipantReplaceGatewayTimeout{}
}

/*PostConversationsMessageParticipantReplaceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessageParticipantReplaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessageParticipantReplaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messages/{conversationId}/participants/{participantId}/replace][%d] postConversationsMessageParticipantReplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessageParticipantReplaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessageParticipantReplaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
