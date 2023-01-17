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

// GetConversationsCallbackParticipantWrapupReader is a Reader for the GetConversationsCallbackParticipantWrapup structure.
type GetConversationsCallbackParticipantWrapupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallbackParticipantWrapupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallbackParticipantWrapupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallbackParticipantWrapupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallbackParticipantWrapupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallbackParticipantWrapupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallbackParticipantWrapupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallbackParticipantWrapupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallbackParticipantWrapupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallbackParticipantWrapupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallbackParticipantWrapupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallbackParticipantWrapupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallbackParticipantWrapupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallbackParticipantWrapupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallbackParticipantWrapupOK creates a GetConversationsCallbackParticipantWrapupOK with default headers values
func NewGetConversationsCallbackParticipantWrapupOK() *GetConversationsCallbackParticipantWrapupOK {
	return &GetConversationsCallbackParticipantWrapupOK{}
}

/*
GetConversationsCallbackParticipantWrapupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsCallbackParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

// IsSuccess returns true when this get conversations callback participant wrapup o k response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations callback participant wrapup o k response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup o k response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback participant wrapup o k response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup o k response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsCallbackParticipantWrapupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupOK) GetPayload() *models.AssignedWrapupCode {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedWrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupBadRequest creates a GetConversationsCallbackParticipantWrapupBadRequest with default headers values
func NewGetConversationsCallbackParticipantWrapupBadRequest() *GetConversationsCallbackParticipantWrapupBadRequest {
	return &GetConversationsCallbackParticipantWrapupBadRequest{}
}

/*
GetConversationsCallbackParticipantWrapupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallbackParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup bad request response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup bad request response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup bad request response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup bad request response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup bad request response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsCallbackParticipantWrapupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupUnauthorized creates a GetConversationsCallbackParticipantWrapupUnauthorized with default headers values
func NewGetConversationsCallbackParticipantWrapupUnauthorized() *GetConversationsCallbackParticipantWrapupUnauthorized {
	return &GetConversationsCallbackParticipantWrapupUnauthorized{}
}

/*
GetConversationsCallbackParticipantWrapupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallbackParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup unauthorized response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup unauthorized response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup unauthorized response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup unauthorized response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup unauthorized response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsCallbackParticipantWrapupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupForbidden creates a GetConversationsCallbackParticipantWrapupForbidden with default headers values
func NewGetConversationsCallbackParticipantWrapupForbidden() *GetConversationsCallbackParticipantWrapupForbidden {
	return &GetConversationsCallbackParticipantWrapupForbidden{}
}

/*
GetConversationsCallbackParticipantWrapupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallbackParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup forbidden response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup forbidden response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup forbidden response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup forbidden response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup forbidden response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsCallbackParticipantWrapupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupNotFound creates a GetConversationsCallbackParticipantWrapupNotFound with default headers values
func NewGetConversationsCallbackParticipantWrapupNotFound() *GetConversationsCallbackParticipantWrapupNotFound {
	return &GetConversationsCallbackParticipantWrapupNotFound{}
}

/*
GetConversationsCallbackParticipantWrapupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsCallbackParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup not found response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup not found response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup not found response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup not found response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup not found response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsCallbackParticipantWrapupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupRequestTimeout creates a GetConversationsCallbackParticipantWrapupRequestTimeout with default headers values
func NewGetConversationsCallbackParticipantWrapupRequestTimeout() *GetConversationsCallbackParticipantWrapupRequestTimeout {
	return &GetConversationsCallbackParticipantWrapupRequestTimeout{}
}

/*
GetConversationsCallbackParticipantWrapupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallbackParticipantWrapupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup request timeout response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup request timeout response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup request timeout response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup request timeout response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup request timeout response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupRequestEntityTooLarge creates a GetConversationsCallbackParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationsCallbackParticipantWrapupRequestEntityTooLarge() *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge {
	return &GetConversationsCallbackParticipantWrapupRequestEntityTooLarge{}
}

/*
GetConversationsCallbackParticipantWrapupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsCallbackParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup request entity too large response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup request entity too large response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup request entity too large response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup request entity too large response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup request entity too large response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupUnsupportedMediaType creates a GetConversationsCallbackParticipantWrapupUnsupportedMediaType with default headers values
func NewGetConversationsCallbackParticipantWrapupUnsupportedMediaType() *GetConversationsCallbackParticipantWrapupUnsupportedMediaType {
	return &GetConversationsCallbackParticipantWrapupUnsupportedMediaType{}
}

/*
GetConversationsCallbackParticipantWrapupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallbackParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup unsupported media type response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup unsupported media type response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup unsupported media type response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup unsupported media type response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup unsupported media type response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupTooManyRequests creates a GetConversationsCallbackParticipantWrapupTooManyRequests with default headers values
func NewGetConversationsCallbackParticipantWrapupTooManyRequests() *GetConversationsCallbackParticipantWrapupTooManyRequests {
	return &GetConversationsCallbackParticipantWrapupTooManyRequests{}
}

/*
GetConversationsCallbackParticipantWrapupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallbackParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup too many requests response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup too many requests response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup too many requests response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations callback participant wrapup too many requests response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations callback participant wrapup too many requests response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupInternalServerError creates a GetConversationsCallbackParticipantWrapupInternalServerError with default headers values
func NewGetConversationsCallbackParticipantWrapupInternalServerError() *GetConversationsCallbackParticipantWrapupInternalServerError {
	return &GetConversationsCallbackParticipantWrapupInternalServerError{}
}

/*
GetConversationsCallbackParticipantWrapupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallbackParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup internal server error response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup internal server error response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup internal server error response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback participant wrapup internal server error response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback participant wrapup internal server error response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsCallbackParticipantWrapupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupServiceUnavailable creates a GetConversationsCallbackParticipantWrapupServiceUnavailable with default headers values
func NewGetConversationsCallbackParticipantWrapupServiceUnavailable() *GetConversationsCallbackParticipantWrapupServiceUnavailable {
	return &GetConversationsCallbackParticipantWrapupServiceUnavailable{}
}

/*
GetConversationsCallbackParticipantWrapupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallbackParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup service unavailable response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup service unavailable response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup service unavailable response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback participant wrapup service unavailable response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback participant wrapup service unavailable response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupGatewayTimeout creates a GetConversationsCallbackParticipantWrapupGatewayTimeout with default headers values
func NewGetConversationsCallbackParticipantWrapupGatewayTimeout() *GetConversationsCallbackParticipantWrapupGatewayTimeout {
	return &GetConversationsCallbackParticipantWrapupGatewayTimeout{}
}

/*
GetConversationsCallbackParticipantWrapupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsCallbackParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations callback participant wrapup gateway timeout response has a 2xx status code
func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations callback participant wrapup gateway timeout response has a 3xx status code
func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations callback participant wrapup gateway timeout response has a 4xx status code
func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations callback participant wrapup gateway timeout response has a 5xx status code
func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations callback participant wrapup gateway timeout response a status code equal to that given
func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallbackParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
