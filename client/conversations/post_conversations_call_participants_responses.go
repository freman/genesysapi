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

// PostConversationsCallParticipantsReader is a Reader for the PostConversationsCallParticipants structure.
type PostConversationsCallParticipantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsCallParticipantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsCallParticipantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsCallParticipantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsCallParticipantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsCallParticipantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsCallParticipantsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostConversationsCallParticipantsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsCallParticipantsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsCallParticipantsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsCallParticipantsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsCallParticipantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsCallParticipantsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsCallParticipantsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsCallParticipantsOK creates a PostConversationsCallParticipantsOK with default headers values
func NewPostConversationsCallParticipantsOK() *PostConversationsCallParticipantsOK {
	return &PostConversationsCallParticipantsOK{}
}

/*
PostConversationsCallParticipantsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostConversationsCallParticipantsOK struct {
	Payload *models.Conversation
}

// IsSuccess returns true when this post conversations call participants o k response has a 2xx status code
func (o *PostConversationsCallParticipantsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post conversations call participants o k response has a 3xx status code
func (o *PostConversationsCallParticipantsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants o k response has a 4xx status code
func (o *PostConversationsCallParticipantsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participants o k response has a 5xx status code
func (o *PostConversationsCallParticipantsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants o k response a status code equal to that given
func (o *PostConversationsCallParticipantsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostConversationsCallParticipantsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsCallParticipantsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsOK  %+v", 200, o.Payload)
}

func (o *PostConversationsCallParticipantsOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *PostConversationsCallParticipantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsBadRequest creates a PostConversationsCallParticipantsBadRequest with default headers values
func NewPostConversationsCallParticipantsBadRequest() *PostConversationsCallParticipantsBadRequest {
	return &PostConversationsCallParticipantsBadRequest{}
}

/*
PostConversationsCallParticipantsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsCallParticipantsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants bad request response has a 2xx status code
func (o *PostConversationsCallParticipantsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants bad request response has a 3xx status code
func (o *PostConversationsCallParticipantsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants bad request response has a 4xx status code
func (o *PostConversationsCallParticipantsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants bad request response has a 5xx status code
func (o *PostConversationsCallParticipantsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants bad request response a status code equal to that given
func (o *PostConversationsCallParticipantsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostConversationsCallParticipantsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallParticipantsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsCallParticipantsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsUnauthorized creates a PostConversationsCallParticipantsUnauthorized with default headers values
func NewPostConversationsCallParticipantsUnauthorized() *PostConversationsCallParticipantsUnauthorized {
	return &PostConversationsCallParticipantsUnauthorized{}
}

/*
PostConversationsCallParticipantsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsCallParticipantsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants unauthorized response has a 2xx status code
func (o *PostConversationsCallParticipantsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants unauthorized response has a 3xx status code
func (o *PostConversationsCallParticipantsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants unauthorized response has a 4xx status code
func (o *PostConversationsCallParticipantsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants unauthorized response has a 5xx status code
func (o *PostConversationsCallParticipantsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants unauthorized response a status code equal to that given
func (o *PostConversationsCallParticipantsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostConversationsCallParticipantsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallParticipantsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsCallParticipantsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsForbidden creates a PostConversationsCallParticipantsForbidden with default headers values
func NewPostConversationsCallParticipantsForbidden() *PostConversationsCallParticipantsForbidden {
	return &PostConversationsCallParticipantsForbidden{}
}

/*
PostConversationsCallParticipantsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsCallParticipantsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants forbidden response has a 2xx status code
func (o *PostConversationsCallParticipantsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants forbidden response has a 3xx status code
func (o *PostConversationsCallParticipantsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants forbidden response has a 4xx status code
func (o *PostConversationsCallParticipantsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants forbidden response has a 5xx status code
func (o *PostConversationsCallParticipantsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants forbidden response a status code equal to that given
func (o *PostConversationsCallParticipantsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostConversationsCallParticipantsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallParticipantsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsCallParticipantsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsNotFound creates a PostConversationsCallParticipantsNotFound with default headers values
func NewPostConversationsCallParticipantsNotFound() *PostConversationsCallParticipantsNotFound {
	return &PostConversationsCallParticipantsNotFound{}
}

/*
PostConversationsCallParticipantsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostConversationsCallParticipantsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants not found response has a 2xx status code
func (o *PostConversationsCallParticipantsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants not found response has a 3xx status code
func (o *PostConversationsCallParticipantsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants not found response has a 4xx status code
func (o *PostConversationsCallParticipantsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants not found response has a 5xx status code
func (o *PostConversationsCallParticipantsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants not found response a status code equal to that given
func (o *PostConversationsCallParticipantsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostConversationsCallParticipantsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallParticipantsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsCallParticipantsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsRequestTimeout creates a PostConversationsCallParticipantsRequestTimeout with default headers values
func NewPostConversationsCallParticipantsRequestTimeout() *PostConversationsCallParticipantsRequestTimeout {
	return &PostConversationsCallParticipantsRequestTimeout{}
}

/*
PostConversationsCallParticipantsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostConversationsCallParticipantsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants request timeout response has a 2xx status code
func (o *PostConversationsCallParticipantsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants request timeout response has a 3xx status code
func (o *PostConversationsCallParticipantsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants request timeout response has a 4xx status code
func (o *PostConversationsCallParticipantsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants request timeout response has a 5xx status code
func (o *PostConversationsCallParticipantsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants request timeout response a status code equal to that given
func (o *PostConversationsCallParticipantsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostConversationsCallParticipantsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsCallParticipantsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostConversationsCallParticipantsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsRequestEntityTooLarge creates a PostConversationsCallParticipantsRequestEntityTooLarge with default headers values
func NewPostConversationsCallParticipantsRequestEntityTooLarge() *PostConversationsCallParticipantsRequestEntityTooLarge {
	return &PostConversationsCallParticipantsRequestEntityTooLarge{}
}

/*
PostConversationsCallParticipantsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostConversationsCallParticipantsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants request entity too large response has a 2xx status code
func (o *PostConversationsCallParticipantsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants request entity too large response has a 3xx status code
func (o *PostConversationsCallParticipantsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants request entity too large response has a 4xx status code
func (o *PostConversationsCallParticipantsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants request entity too large response has a 5xx status code
func (o *PostConversationsCallParticipantsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants request entity too large response a status code equal to that given
func (o *PostConversationsCallParticipantsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostConversationsCallParticipantsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallParticipantsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsCallParticipantsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsUnsupportedMediaType creates a PostConversationsCallParticipantsUnsupportedMediaType with default headers values
func NewPostConversationsCallParticipantsUnsupportedMediaType() *PostConversationsCallParticipantsUnsupportedMediaType {
	return &PostConversationsCallParticipantsUnsupportedMediaType{}
}

/*
PostConversationsCallParticipantsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsCallParticipantsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants unsupported media type response has a 2xx status code
func (o *PostConversationsCallParticipantsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants unsupported media type response has a 3xx status code
func (o *PostConversationsCallParticipantsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants unsupported media type response has a 4xx status code
func (o *PostConversationsCallParticipantsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants unsupported media type response has a 5xx status code
func (o *PostConversationsCallParticipantsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants unsupported media type response a status code equal to that given
func (o *PostConversationsCallParticipantsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostConversationsCallParticipantsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallParticipantsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsCallParticipantsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsTooManyRequests creates a PostConversationsCallParticipantsTooManyRequests with default headers values
func NewPostConversationsCallParticipantsTooManyRequests() *PostConversationsCallParticipantsTooManyRequests {
	return &PostConversationsCallParticipantsTooManyRequests{}
}

/*
PostConversationsCallParticipantsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostConversationsCallParticipantsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants too many requests response has a 2xx status code
func (o *PostConversationsCallParticipantsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants too many requests response has a 3xx status code
func (o *PostConversationsCallParticipantsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants too many requests response has a 4xx status code
func (o *PostConversationsCallParticipantsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post conversations call participants too many requests response has a 5xx status code
func (o *PostConversationsCallParticipantsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post conversations call participants too many requests response a status code equal to that given
func (o *PostConversationsCallParticipantsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostConversationsCallParticipantsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallParticipantsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsCallParticipantsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsInternalServerError creates a PostConversationsCallParticipantsInternalServerError with default headers values
func NewPostConversationsCallParticipantsInternalServerError() *PostConversationsCallParticipantsInternalServerError {
	return &PostConversationsCallParticipantsInternalServerError{}
}

/*
PostConversationsCallParticipantsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsCallParticipantsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants internal server error response has a 2xx status code
func (o *PostConversationsCallParticipantsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants internal server error response has a 3xx status code
func (o *PostConversationsCallParticipantsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants internal server error response has a 4xx status code
func (o *PostConversationsCallParticipantsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participants internal server error response has a 5xx status code
func (o *PostConversationsCallParticipantsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participants internal server error response a status code equal to that given
func (o *PostConversationsCallParticipantsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostConversationsCallParticipantsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallParticipantsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsCallParticipantsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsServiceUnavailable creates a PostConversationsCallParticipantsServiceUnavailable with default headers values
func NewPostConversationsCallParticipantsServiceUnavailable() *PostConversationsCallParticipantsServiceUnavailable {
	return &PostConversationsCallParticipantsServiceUnavailable{}
}

/*
PostConversationsCallParticipantsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsCallParticipantsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants service unavailable response has a 2xx status code
func (o *PostConversationsCallParticipantsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants service unavailable response has a 3xx status code
func (o *PostConversationsCallParticipantsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants service unavailable response has a 4xx status code
func (o *PostConversationsCallParticipantsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participants service unavailable response has a 5xx status code
func (o *PostConversationsCallParticipantsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participants service unavailable response a status code equal to that given
func (o *PostConversationsCallParticipantsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostConversationsCallParticipantsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallParticipantsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsCallParticipantsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsCallParticipantsGatewayTimeout creates a PostConversationsCallParticipantsGatewayTimeout with default headers values
func NewPostConversationsCallParticipantsGatewayTimeout() *PostConversationsCallParticipantsGatewayTimeout {
	return &PostConversationsCallParticipantsGatewayTimeout{}
}

/*
PostConversationsCallParticipantsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostConversationsCallParticipantsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post conversations call participants gateway timeout response has a 2xx status code
func (o *PostConversationsCallParticipantsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post conversations call participants gateway timeout response has a 3xx status code
func (o *PostConversationsCallParticipantsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post conversations call participants gateway timeout response has a 4xx status code
func (o *PostConversationsCallParticipantsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post conversations call participants gateway timeout response has a 5xx status code
func (o *PostConversationsCallParticipantsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post conversations call participants gateway timeout response a status code equal to that given
func (o *PostConversationsCallParticipantsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostConversationsCallParticipantsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallParticipantsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/conversations/calls/{conversationId}/participants][%d] postConversationsCallParticipantsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsCallParticipantsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsCallParticipantsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
