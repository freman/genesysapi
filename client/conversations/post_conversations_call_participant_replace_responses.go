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

// PostConversationsCallParticipantReplaceReader is a Reader for the PostConversationsCallParticipantReplace structure.
type PostConversationsCallParticipantReplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsCallParticipantReplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostConversationsCallParticipantReplaceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsCallParticipantReplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsCallParticipantReplaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsCallParticipantReplaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsCallParticipantReplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsCallParticipantReplaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsCallParticipantReplaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsCallParticipantReplaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsCallParticipantReplaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsCallParticipantReplaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsCallParticipantReplaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsCallParticipantReplaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsCallParticipantReplaceAccepted creates a PostConversationsCallParticipantReplaceAccepted with default headers values
func NewPostConversationsCallParticipantReplaceAccepted() *PostConversationsCallParticipantReplaceAccepted {
	return &PostConversationsCallParticipantReplaceAccepted{}
}

/*
PostConversationsCallParticipantReplaceAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PostConversationsCallParticipantReplaceAccepted struct {
}

// IsSuccess returns true when this post conversations call participant replace accepted response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations call participant replace accepted response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace accepted response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participant replace accepted response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace accepted response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostConversationsCallParticipantReplaceAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsCallParticipantReplaceAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceAccepted ", 202)
}

func (o *PostConversationsCallParticipantReplaceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostConversationsCallParticipantReplaceBadRequest creates a PostConversationsCallParticipantReplaceBadRequest with default headers values
func NewPostConversationsCallParticipantReplaceBadRequest() *PostConversationsCallParticipantReplaceBadRequest {
	return &PostConversationsCallParticipantReplaceBadRequest{}
}

/*
PostConversationsCallParticipantReplaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsCallParticipantReplaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace bad request response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace bad request response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace bad request response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace bad request response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace bad request response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsCallParticipantReplaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceUnauthorized creates a PostConversationsCallParticipantReplaceUnauthorized with default headers values
func NewPostConversationsCallParticipantReplaceUnauthorized() *PostConversationsCallParticipantReplaceUnauthorized {
	return &PostConversationsCallParticipantReplaceUnauthorized{}
}

/*
PostConversationsCallParticipantReplaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsCallParticipantReplaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace unauthorized response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace unauthorized response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace unauthorized response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace unauthorized response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace unauthorized response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsCallParticipantReplaceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceForbidden creates a PostConversationsCallParticipantReplaceForbidden with default headers values
func NewPostConversationsCallParticipantReplaceForbidden() *PostConversationsCallParticipantReplaceForbidden {
	return &PostConversationsCallParticipantReplaceForbidden{}
}

/*
PostConversationsCallParticipantReplaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsCallParticipantReplaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace forbidden response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace forbidden response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace forbidden response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace forbidden response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace forbidden response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsCallParticipantReplaceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceNotFound creates a PostConversationsCallParticipantReplaceNotFound with default headers values
func NewPostConversationsCallParticipantReplaceNotFound() *PostConversationsCallParticipantReplaceNotFound {
	return &PostConversationsCallParticipantReplaceNotFound{}
}

/*
PostConversationsCallParticipantReplaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsCallParticipantReplaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace not found response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace not found response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace not found response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace not found response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace not found response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsCallParticipantReplaceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceRequestTimeout creates a PostConversationsCallParticipantReplaceRequestTimeout with default headers values
func NewPostConversationsCallParticipantReplaceRequestTimeout() *PostConversationsCallParticipantReplaceRequestTimeout {
	return &PostConversationsCallParticipantReplaceRequestTimeout{}
}

/*
PostConversationsCallParticipantReplaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsCallParticipantReplaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace request timeout response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace request timeout response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace request timeout response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace request timeout response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace request timeout response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsCallParticipantReplaceRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceRequestEntityTooLarge creates a PostConversationsCallParticipantReplaceRequestEntityTooLarge with default headers values
func NewPostConversationsCallParticipantReplaceRequestEntityTooLarge() *PostConversationsCallParticipantReplaceRequestEntityTooLarge {
	return &PostConversationsCallParticipantReplaceRequestEntityTooLarge{}
}

/*
PostConversationsCallParticipantReplaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostConversationsCallParticipantReplaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace request entity too large response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace request entity too large response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace request entity too large response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace request entity too large response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace request entity too large response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceUnsupportedMediaType creates a PostConversationsCallParticipantReplaceUnsupportedMediaType with default headers values
func NewPostConversationsCallParticipantReplaceUnsupportedMediaType() *PostConversationsCallParticipantReplaceUnsupportedMediaType {
	return &PostConversationsCallParticipantReplaceUnsupportedMediaType{}
}

/*
PostConversationsCallParticipantReplaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsCallParticipantReplaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace unsupported media type response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace unsupported media type response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace unsupported media type response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace unsupported media type response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace unsupported media type response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceTooManyRequests creates a PostConversationsCallParticipantReplaceTooManyRequests with default headers values
func NewPostConversationsCallParticipantReplaceTooManyRequests() *PostConversationsCallParticipantReplaceTooManyRequests {
	return &PostConversationsCallParticipantReplaceTooManyRequests{}
}

/*
PostConversationsCallParticipantReplaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsCallParticipantReplaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace too many requests response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace too many requests response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace too many requests response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participant replace too many requests response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participant replace too many requests response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsCallParticipantReplaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceInternalServerError creates a PostConversationsCallParticipantReplaceInternalServerError with default headers values
func NewPostConversationsCallParticipantReplaceInternalServerError() *PostConversationsCallParticipantReplaceInternalServerError {
	return &PostConversationsCallParticipantReplaceInternalServerError{}
}

/*
PostConversationsCallParticipantReplaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsCallParticipantReplaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace internal server error response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace internal server error response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace internal server error response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participant replace internal server error response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participant replace internal server error response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsCallParticipantReplaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceServiceUnavailable creates a PostConversationsCallParticipantReplaceServiceUnavailable with default headers values
func NewPostConversationsCallParticipantReplaceServiceUnavailable() *PostConversationsCallParticipantReplaceServiceUnavailable {
	return &PostConversationsCallParticipantReplaceServiceUnavailable{}
}

/*
PostConversationsCallParticipantReplaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsCallParticipantReplaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace service unavailable response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace service unavailable response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace service unavailable response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participant replace service unavailable response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participant replace service unavailable response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsCallParticipantReplaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantReplaceGatewayTimeout creates a PostConversationsCallParticipantReplaceGatewayTimeout with default headers values
func NewPostConversationsCallParticipantReplaceGatewayTimeout() *PostConversationsCallParticipantReplaceGatewayTimeout {
	return &PostConversationsCallParticipantReplaceGatewayTimeout{}
}

/*
PostConversationsCallParticipantReplaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsCallParticipantReplaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participant replace gateway timeout response has a 2xx status code
func (o *PostConversationsCallParticipantReplaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participant replace gateway timeout response has a 3xx status code
func (o *PostConversationsCallParticipantReplaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participant replace gateway timeout response has a 4xx status code
func (o *PostConversationsCallParticipantReplaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participant replace gateway timeout response has a 5xx status code
func (o *PostConversationsCallParticipantReplaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participant replace gateway timeout response a status code equal to that given
func (o *PostConversationsCallParticipantReplaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsCallParticipantReplaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants/{participantId}/replace][%d] postConversationsCallParticipantReplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallParticipantReplaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantReplaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
