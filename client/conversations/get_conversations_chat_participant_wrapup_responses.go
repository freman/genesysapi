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

// GetConversationsChatParticipantWrapupReader is a Reader for the GetConversationsChatParticipantWrapup structure.
type GetConversationsChatParticipantWrapupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsChatParticipantWrapupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsChatParticipantWrapupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsChatParticipantWrapupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsChatParticipantWrapupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsChatParticipantWrapupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsChatParticipantWrapupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsChatParticipantWrapupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsChatParticipantWrapupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsChatParticipantWrapupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsChatParticipantWrapupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsChatParticipantWrapupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsChatParticipantWrapupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsChatParticipantWrapupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsChatParticipantWrapupOK creates a GetConversationsChatParticipantWrapupOK with default headers values
func NewGetConversationsChatParticipantWrapupOK() *GetConversationsChatParticipantWrapupOK {
	return &GetConversationsChatParticipantWrapupOK{}
}

/*
GetConversationsChatParticipantWrapupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsChatParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

// IsSuccess returns true when this get conversations chat participant wrapup o k response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations chat participant wrapup o k response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup o k response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapup o k response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup o k response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsChatParticipantWrapupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupOK) GetPayload() *models.AssignedWrapupCode {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedWrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupBadRequest creates a GetConversationsChatParticipantWrapupBadRequest with default headers values
func NewGetConversationsChatParticipantWrapupBadRequest() *GetConversationsChatParticipantWrapupBadRequest {
	return &GetConversationsChatParticipantWrapupBadRequest{}
}

/*
GetConversationsChatParticipantWrapupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsChatParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup bad request response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup bad request response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup bad request response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup bad request response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup bad request response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsChatParticipantWrapupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupUnauthorized creates a GetConversationsChatParticipantWrapupUnauthorized with default headers values
func NewGetConversationsChatParticipantWrapupUnauthorized() *GetConversationsChatParticipantWrapupUnauthorized {
	return &GetConversationsChatParticipantWrapupUnauthorized{}
}

/*
GetConversationsChatParticipantWrapupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsChatParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup unauthorized response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup unauthorized response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup unauthorized response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup unauthorized response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup unauthorized response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsChatParticipantWrapupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupForbidden creates a GetConversationsChatParticipantWrapupForbidden with default headers values
func NewGetConversationsChatParticipantWrapupForbidden() *GetConversationsChatParticipantWrapupForbidden {
	return &GetConversationsChatParticipantWrapupForbidden{}
}

/*
GetConversationsChatParticipantWrapupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsChatParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup forbidden response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup forbidden response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup forbidden response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup forbidden response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup forbidden response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsChatParticipantWrapupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupNotFound creates a GetConversationsChatParticipantWrapupNotFound with default headers values
func NewGetConversationsChatParticipantWrapupNotFound() *GetConversationsChatParticipantWrapupNotFound {
	return &GetConversationsChatParticipantWrapupNotFound{}
}

/*
GetConversationsChatParticipantWrapupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsChatParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup not found response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup not found response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup not found response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup not found response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup not found response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsChatParticipantWrapupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupRequestTimeout creates a GetConversationsChatParticipantWrapupRequestTimeout with default headers values
func NewGetConversationsChatParticipantWrapupRequestTimeout() *GetConversationsChatParticipantWrapupRequestTimeout {
	return &GetConversationsChatParticipantWrapupRequestTimeout{}
}

/*
GetConversationsChatParticipantWrapupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsChatParticipantWrapupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup request timeout response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup request timeout response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup request timeout response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup request timeout response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup request timeout response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsChatParticipantWrapupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupRequestEntityTooLarge creates a GetConversationsChatParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationsChatParticipantWrapupRequestEntityTooLarge() *GetConversationsChatParticipantWrapupRequestEntityTooLarge {
	return &GetConversationsChatParticipantWrapupRequestEntityTooLarge{}
}

/*
GetConversationsChatParticipantWrapupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsChatParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup request entity too large response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup request entity too large response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup request entity too large response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup request entity too large response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup request entity too large response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupUnsupportedMediaType creates a GetConversationsChatParticipantWrapupUnsupportedMediaType with default headers values
func NewGetConversationsChatParticipantWrapupUnsupportedMediaType() *GetConversationsChatParticipantWrapupUnsupportedMediaType {
	return &GetConversationsChatParticipantWrapupUnsupportedMediaType{}
}

/*
GetConversationsChatParticipantWrapupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsChatParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup unsupported media type response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup unsupported media type response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup unsupported media type response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup unsupported media type response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup unsupported media type response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupTooManyRequests creates a GetConversationsChatParticipantWrapupTooManyRequests with default headers values
func NewGetConversationsChatParticipantWrapupTooManyRequests() *GetConversationsChatParticipantWrapupTooManyRequests {
	return &GetConversationsChatParticipantWrapupTooManyRequests{}
}

/*
GetConversationsChatParticipantWrapupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsChatParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup too many requests response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup too many requests response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup too many requests response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations chat participant wrapup too many requests response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations chat participant wrapup too many requests response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsChatParticipantWrapupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupInternalServerError creates a GetConversationsChatParticipantWrapupInternalServerError with default headers values
func NewGetConversationsChatParticipantWrapupInternalServerError() *GetConversationsChatParticipantWrapupInternalServerError {
	return &GetConversationsChatParticipantWrapupInternalServerError{}
}

/*
GetConversationsChatParticipantWrapupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsChatParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup internal server error response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup internal server error response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup internal server error response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapup internal server error response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapup internal server error response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsChatParticipantWrapupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupServiceUnavailable creates a GetConversationsChatParticipantWrapupServiceUnavailable with default headers values
func NewGetConversationsChatParticipantWrapupServiceUnavailable() *GetConversationsChatParticipantWrapupServiceUnavailable {
	return &GetConversationsChatParticipantWrapupServiceUnavailable{}
}

/*
GetConversationsChatParticipantWrapupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsChatParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup service unavailable response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup service unavailable response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup service unavailable response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapup service unavailable response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapup service unavailable response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsChatParticipantWrapupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsChatParticipantWrapupGatewayTimeout creates a GetConversationsChatParticipantWrapupGatewayTimeout with default headers values
func NewGetConversationsChatParticipantWrapupGatewayTimeout() *GetConversationsChatParticipantWrapupGatewayTimeout {
	return &GetConversationsChatParticipantWrapupGatewayTimeout{}
}

/*
GetConversationsChatParticipantWrapupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsChatParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations chat participant wrapup gateway timeout response has a 2xx status code
func (o *GetConversationsChatParticipantWrapupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations chat participant wrapup gateway timeout response has a 3xx status code
func (o *GetConversationsChatParticipantWrapupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations chat participant wrapup gateway timeout response has a 4xx status code
func (o *GetConversationsChatParticipantWrapupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations chat participant wrapup gateway timeout response has a 5xx status code
func (o *GetConversationsChatParticipantWrapupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations chat participant wrapup gateway timeout response a status code equal to that given
func (o *GetConversationsChatParticipantWrapupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsChatParticipantWrapupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/chats/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsChatParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsChatParticipantWrapupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsChatParticipantWrapupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
