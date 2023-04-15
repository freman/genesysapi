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
	case 408:
		result := NewPostConversationsChatParticipantReplaceRequestTimeout()
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

/*
PostConversationsChatParticipantReplaceAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostConversationsChatParticipantReplaceAccepted struct {
}

// IsSuccess returns true when this post conversations chat participant replace accepted response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations chat participant replace accepted response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace accepted response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chat participant replace accepted response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace accepted response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostConversationsChatParticipantReplaceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsChatParticipantReplaceAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsChatParticipantReplaceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsChatParticipantReplaceBadRequest creates a PostConversationsChatParticipantReplaceBadRequest with default headers values
func NewPostConversationsChatParticipantReplaceBadRequest() *PostConversationsChatParticipantReplaceBadRequest {
	return &PostConversationsChatParticipantReplaceBadRequest{}
}

/*
PostConversationsChatParticipantReplaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsChatParticipantReplaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace bad request response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace bad request response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace bad request response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace bad request response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace bad request response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsChatParticipantReplaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceBadRequest) String() string {
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

/*
PostConversationsChatParticipantReplaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsChatParticipantReplaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace unauthorized response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace unauthorized response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace unauthorized response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace unauthorized response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace unauthorized response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsChatParticipantReplaceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceUnauthorized) String() string {
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

/*
PostConversationsChatParticipantReplaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsChatParticipantReplaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace forbidden response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace forbidden response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace forbidden response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace forbidden response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace forbidden response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsChatParticipantReplaceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceForbidden) String() string {
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

/*
PostConversationsChatParticipantReplaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsChatParticipantReplaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace not found response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace not found response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace not found response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace not found response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace not found response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsChatParticipantReplaceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceNotFound) String() string {
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

// NewPostConversationsChatParticipantReplaceRequestTimeout creates a PostConversationsChatParticipantReplaceRequestTimeout with default headers values
func NewPostConversationsChatParticipantReplaceRequestTimeout() *PostConversationsChatParticipantReplaceRequestTimeout {
	return &PostConversationsChatParticipantReplaceRequestTimeout{}
}

/*
PostConversationsChatParticipantReplaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsChatParticipantReplaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace request timeout response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace request timeout response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace request timeout response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace request timeout response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace request timeout response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsChatParticipantReplaceRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsChatParticipantReplaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PostConversationsChatParticipantReplaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsChatParticipantReplaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace request entity too large response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace request entity too large response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace request entity too large response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace request entity too large response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace request entity too large response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceRequestEntityTooLarge) String() string {
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

/*
PostConversationsChatParticipantReplaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsChatParticipantReplaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace unsupported media type response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace unsupported media type response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace unsupported media type response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace unsupported media type response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace unsupported media type response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceUnsupportedMediaType) String() string {
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

/*
PostConversationsChatParticipantReplaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsChatParticipantReplaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace too many requests response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace too many requests response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace too many requests response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations chat participant replace too many requests response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations chat participant replace too many requests response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsChatParticipantReplaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceTooManyRequests) String() string {
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

/*
PostConversationsChatParticipantReplaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsChatParticipantReplaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace internal server error response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace internal server error response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace internal server error response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chat participant replace internal server error response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chat participant replace internal server error response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsChatParticipantReplaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceInternalServerError) String() string {
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

/*
PostConversationsChatParticipantReplaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsChatParticipantReplaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace service unavailable response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace service unavailable response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace service unavailable response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chat participant replace service unavailable response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chat participant replace service unavailable response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsChatParticipantReplaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceServiceUnavailable) String() string {
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

/*
PostConversationsChatParticipantReplaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsChatParticipantReplaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations chat participant replace gateway timeout response has a 2xx status code
func (o *PostConversationsChatParticipantReplaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations chat participant replace gateway timeout response has a 3xx status code
func (o *PostConversationsChatParticipantReplaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations chat participant replace gateway timeout response has a 4xx status code
func (o *PostConversationsChatParticipantReplaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations chat participant replace gateway timeout response has a 5xx status code
func (o *PostConversationsChatParticipantReplaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations chat participant replace gateway timeout response a status code equal to that given
func (o *PostConversationsChatParticipantReplaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsChatParticipantReplaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/chats/{conversationId}/participants/{participantId}/replace][%d] postConversationsChatParticipantReplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsChatParticipantReplaceGatewayTimeout) String() string {
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
