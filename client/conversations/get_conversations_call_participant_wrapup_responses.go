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

// GetConversationsCallParticipantWrapupReader is a Reader for the GetConversationsCallParticipantWrapup structure.
type GetConversationsCallParticipantWrapupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallParticipantWrapupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallParticipantWrapupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallParticipantWrapupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallParticipantWrapupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallParticipantWrapupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallParticipantWrapupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallParticipantWrapupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallParticipantWrapupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallParticipantWrapupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallParticipantWrapupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallParticipantWrapupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallParticipantWrapupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallParticipantWrapupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallParticipantWrapupOK creates a GetConversationsCallParticipantWrapupOK with default headers values
func NewGetConversationsCallParticipantWrapupOK() *GetConversationsCallParticipantWrapupOK {
	return &GetConversationsCallParticipantWrapupOK{}
}

/*
GetConversationsCallParticipantWrapupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationsCallParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

// IsSuccess returns true when this get conversations call participant wrapup o k response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversations call participant wrapup o k response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup o k response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapup o k response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup o k response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationsCallParticipantWrapupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupOK) GetPayload() *models.AssignedWrapupCode {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedWrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupBadRequest creates a GetConversationsCallParticipantWrapupBadRequest with default headers values
func NewGetConversationsCallParticipantWrapupBadRequest() *GetConversationsCallParticipantWrapupBadRequest {
	return &GetConversationsCallParticipantWrapupBadRequest{}
}

/*
GetConversationsCallParticipantWrapupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup bad request response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup bad request response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup bad request response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup bad request response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup bad request response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationsCallParticipantWrapupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupUnauthorized creates a GetConversationsCallParticipantWrapupUnauthorized with default headers values
func NewGetConversationsCallParticipantWrapupUnauthorized() *GetConversationsCallParticipantWrapupUnauthorized {
	return &GetConversationsCallParticipantWrapupUnauthorized{}
}

/*
GetConversationsCallParticipantWrapupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup unauthorized response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup unauthorized response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup unauthorized response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup unauthorized response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup unauthorized response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationsCallParticipantWrapupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupForbidden creates a GetConversationsCallParticipantWrapupForbidden with default headers values
func NewGetConversationsCallParticipantWrapupForbidden() *GetConversationsCallParticipantWrapupForbidden {
	return &GetConversationsCallParticipantWrapupForbidden{}
}

/*
GetConversationsCallParticipantWrapupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup forbidden response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup forbidden response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup forbidden response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup forbidden response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup forbidden response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationsCallParticipantWrapupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupNotFound creates a GetConversationsCallParticipantWrapupNotFound with default headers values
func NewGetConversationsCallParticipantWrapupNotFound() *GetConversationsCallParticipantWrapupNotFound {
	return &GetConversationsCallParticipantWrapupNotFound{}
}

/*
GetConversationsCallParticipantWrapupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationsCallParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup not found response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup not found response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup not found response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup not found response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup not found response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationsCallParticipantWrapupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupRequestTimeout creates a GetConversationsCallParticipantWrapupRequestTimeout with default headers values
func NewGetConversationsCallParticipantWrapupRequestTimeout() *GetConversationsCallParticipantWrapupRequestTimeout {
	return &GetConversationsCallParticipantWrapupRequestTimeout{}
}

/*
GetConversationsCallParticipantWrapupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallParticipantWrapupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup request timeout response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup request timeout response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup request timeout response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup request timeout response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup request timeout response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationsCallParticipantWrapupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupRequestEntityTooLarge creates a GetConversationsCallParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationsCallParticipantWrapupRequestEntityTooLarge() *GetConversationsCallParticipantWrapupRequestEntityTooLarge {
	return &GetConversationsCallParticipantWrapupRequestEntityTooLarge{}
}

/*
GetConversationsCallParticipantWrapupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationsCallParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup request entity too large response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup request entity too large response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup request entity too large response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup request entity too large response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup request entity too large response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupUnsupportedMediaType creates a GetConversationsCallParticipantWrapupUnsupportedMediaType with default headers values
func NewGetConversationsCallParticipantWrapupUnsupportedMediaType() *GetConversationsCallParticipantWrapupUnsupportedMediaType {
	return &GetConversationsCallParticipantWrapupUnsupportedMediaType{}
}

/*
GetConversationsCallParticipantWrapupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup unsupported media type response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup unsupported media type response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup unsupported media type response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup unsupported media type response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup unsupported media type response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupTooManyRequests creates a GetConversationsCallParticipantWrapupTooManyRequests with default headers values
func NewGetConversationsCallParticipantWrapupTooManyRequests() *GetConversationsCallParticipantWrapupTooManyRequests {
	return &GetConversationsCallParticipantWrapupTooManyRequests{}
}

/*
GetConversationsCallParticipantWrapupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup too many requests response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup too many requests response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup too many requests response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversations call participant wrapup too many requests response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversations call participant wrapup too many requests response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationsCallParticipantWrapupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupInternalServerError creates a GetConversationsCallParticipantWrapupInternalServerError with default headers values
func NewGetConversationsCallParticipantWrapupInternalServerError() *GetConversationsCallParticipantWrapupInternalServerError {
	return &GetConversationsCallParticipantWrapupInternalServerError{}
}

/*
GetConversationsCallParticipantWrapupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup internal server error response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup internal server error response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup internal server error response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapup internal server error response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapup internal server error response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationsCallParticipantWrapupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupServiceUnavailable creates a GetConversationsCallParticipantWrapupServiceUnavailable with default headers values
func NewGetConversationsCallParticipantWrapupServiceUnavailable() *GetConversationsCallParticipantWrapupServiceUnavailable {
	return &GetConversationsCallParticipantWrapupServiceUnavailable{}
}

/*
GetConversationsCallParticipantWrapupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup service unavailable response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup service unavailable response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup service unavailable response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapup service unavailable response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapup service unavailable response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationsCallParticipantWrapupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallParticipantWrapupGatewayTimeout creates a GetConversationsCallParticipantWrapupGatewayTimeout with default headers values
func NewGetConversationsCallParticipantWrapupGatewayTimeout() *GetConversationsCallParticipantWrapupGatewayTimeout {
	return &GetConversationsCallParticipantWrapupGatewayTimeout{}
}

/*
GetConversationsCallParticipantWrapupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationsCallParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversations call participant wrapup gateway timeout response has a 2xx status code
func (o *GetConversationsCallParticipantWrapupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversations call participant wrapup gateway timeout response has a 3xx status code
func (o *GetConversationsCallParticipantWrapupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversations call participant wrapup gateway timeout response has a 4xx status code
func (o *GetConversationsCallParticipantWrapupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversations call participant wrapup gateway timeout response has a 5xx status code
func (o *GetConversationsCallParticipantWrapupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversations call participant wrapup gateway timeout response a status code equal to that given
func (o *GetConversationsCallParticipantWrapupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationsCallParticipantWrapupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/calls/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCallParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallParticipantWrapupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallParticipantWrapupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
