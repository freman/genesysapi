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

// GetConversationParticipantWrapupReader is a Reader for the GetConversationParticipantWrapup structure.
type GetConversationParticipantWrapupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationParticipantWrapupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationParticipantWrapupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationParticipantWrapupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationParticipantWrapupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationParticipantWrapupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationParticipantWrapupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationParticipantWrapupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationParticipantWrapupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationParticipantWrapupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationParticipantWrapupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationParticipantWrapupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationParticipantWrapupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationParticipantWrapupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationParticipantWrapupOK creates a GetConversationParticipantWrapupOK with default headers values
func NewGetConversationParticipantWrapupOK() *GetConversationParticipantWrapupOK {
	return &GetConversationParticipantWrapupOK{}
}

/*
GetConversationParticipantWrapupOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

// IsSuccess returns true when this get conversation participant wrapup o k response has a 2xx status code
func (o *GetConversationParticipantWrapupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation participant wrapup o k response has a 3xx status code
func (o *GetConversationParticipantWrapupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup o k response has a 4xx status code
func (o *GetConversationParticipantWrapupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant wrapup o k response has a 5xx status code
func (o *GetConversationParticipantWrapupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup o k response a status code equal to that given
func (o *GetConversationParticipantWrapupOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationParticipantWrapupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationParticipantWrapupOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationParticipantWrapupOK) GetPayload() *models.AssignedWrapupCode {
	return o.Payload
}

func (o *GetConversationParticipantWrapupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedWrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupBadRequest creates a GetConversationParticipantWrapupBadRequest with default headers values
func NewGetConversationParticipantWrapupBadRequest() *GetConversationParticipantWrapupBadRequest {
	return &GetConversationParticipantWrapupBadRequest{}
}

/*
GetConversationParticipantWrapupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup bad request response has a 2xx status code
func (o *GetConversationParticipantWrapupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup bad request response has a 3xx status code
func (o *GetConversationParticipantWrapupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup bad request response has a 4xx status code
func (o *GetConversationParticipantWrapupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup bad request response has a 5xx status code
func (o *GetConversationParticipantWrapupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup bad request response a status code equal to that given
func (o *GetConversationParticipantWrapupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationParticipantWrapupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationParticipantWrapupBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationParticipantWrapupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupUnauthorized creates a GetConversationParticipantWrapupUnauthorized with default headers values
func NewGetConversationParticipantWrapupUnauthorized() *GetConversationParticipantWrapupUnauthorized {
	return &GetConversationParticipantWrapupUnauthorized{}
}

/*
GetConversationParticipantWrapupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup unauthorized response has a 2xx status code
func (o *GetConversationParticipantWrapupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup unauthorized response has a 3xx status code
func (o *GetConversationParticipantWrapupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup unauthorized response has a 4xx status code
func (o *GetConversationParticipantWrapupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup unauthorized response has a 5xx status code
func (o *GetConversationParticipantWrapupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup unauthorized response a status code equal to that given
func (o *GetConversationParticipantWrapupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationParticipantWrapupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationParticipantWrapupUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationParticipantWrapupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupForbidden creates a GetConversationParticipantWrapupForbidden with default headers values
func NewGetConversationParticipantWrapupForbidden() *GetConversationParticipantWrapupForbidden {
	return &GetConversationParticipantWrapupForbidden{}
}

/*
GetConversationParticipantWrapupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup forbidden response has a 2xx status code
func (o *GetConversationParticipantWrapupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup forbidden response has a 3xx status code
func (o *GetConversationParticipantWrapupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup forbidden response has a 4xx status code
func (o *GetConversationParticipantWrapupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup forbidden response has a 5xx status code
func (o *GetConversationParticipantWrapupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup forbidden response a status code equal to that given
func (o *GetConversationParticipantWrapupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationParticipantWrapupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationParticipantWrapupForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationParticipantWrapupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupNotFound creates a GetConversationParticipantWrapupNotFound with default headers values
func NewGetConversationParticipantWrapupNotFound() *GetConversationParticipantWrapupNotFound {
	return &GetConversationParticipantWrapupNotFound{}
}

/*
GetConversationParticipantWrapupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup not found response has a 2xx status code
func (o *GetConversationParticipantWrapupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup not found response has a 3xx status code
func (o *GetConversationParticipantWrapupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup not found response has a 4xx status code
func (o *GetConversationParticipantWrapupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup not found response has a 5xx status code
func (o *GetConversationParticipantWrapupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup not found response a status code equal to that given
func (o *GetConversationParticipantWrapupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationParticipantWrapupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationParticipantWrapupNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationParticipantWrapupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupRequestTimeout creates a GetConversationParticipantWrapupRequestTimeout with default headers values
func NewGetConversationParticipantWrapupRequestTimeout() *GetConversationParticipantWrapupRequestTimeout {
	return &GetConversationParticipantWrapupRequestTimeout{}
}

/*
GetConversationParticipantWrapupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationParticipantWrapupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup request timeout response has a 2xx status code
func (o *GetConversationParticipantWrapupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup request timeout response has a 3xx status code
func (o *GetConversationParticipantWrapupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup request timeout response has a 4xx status code
func (o *GetConversationParticipantWrapupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup request timeout response has a 5xx status code
func (o *GetConversationParticipantWrapupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup request timeout response a status code equal to that given
func (o *GetConversationParticipantWrapupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationParticipantWrapupRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationParticipantWrapupRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationParticipantWrapupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupRequestEntityTooLarge creates a GetConversationParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationParticipantWrapupRequestEntityTooLarge() *GetConversationParticipantWrapupRequestEntityTooLarge {
	return &GetConversationParticipantWrapupRequestEntityTooLarge{}
}

/*
GetConversationParticipantWrapupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup request entity too large response has a 2xx status code
func (o *GetConversationParticipantWrapupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup request entity too large response has a 3xx status code
func (o *GetConversationParticipantWrapupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup request entity too large response has a 4xx status code
func (o *GetConversationParticipantWrapupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup request entity too large response has a 5xx status code
func (o *GetConversationParticipantWrapupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup request entity too large response a status code equal to that given
func (o *GetConversationParticipantWrapupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationParticipantWrapupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationParticipantWrapupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationParticipantWrapupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupUnsupportedMediaType creates a GetConversationParticipantWrapupUnsupportedMediaType with default headers values
func NewGetConversationParticipantWrapupUnsupportedMediaType() *GetConversationParticipantWrapupUnsupportedMediaType {
	return &GetConversationParticipantWrapupUnsupportedMediaType{}
}

/*
GetConversationParticipantWrapupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup unsupported media type response has a 2xx status code
func (o *GetConversationParticipantWrapupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup unsupported media type response has a 3xx status code
func (o *GetConversationParticipantWrapupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup unsupported media type response has a 4xx status code
func (o *GetConversationParticipantWrapupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup unsupported media type response has a 5xx status code
func (o *GetConversationParticipantWrapupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup unsupported media type response a status code equal to that given
func (o *GetConversationParticipantWrapupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationParticipantWrapupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationParticipantWrapupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationParticipantWrapupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupTooManyRequests creates a GetConversationParticipantWrapupTooManyRequests with default headers values
func NewGetConversationParticipantWrapupTooManyRequests() *GetConversationParticipantWrapupTooManyRequests {
	return &GetConversationParticipantWrapupTooManyRequests{}
}

/*
GetConversationParticipantWrapupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup too many requests response has a 2xx status code
func (o *GetConversationParticipantWrapupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup too many requests response has a 3xx status code
func (o *GetConversationParticipantWrapupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup too many requests response has a 4xx status code
func (o *GetConversationParticipantWrapupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant wrapup too many requests response has a 5xx status code
func (o *GetConversationParticipantWrapupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant wrapup too many requests response a status code equal to that given
func (o *GetConversationParticipantWrapupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationParticipantWrapupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationParticipantWrapupTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationParticipantWrapupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupInternalServerError creates a GetConversationParticipantWrapupInternalServerError with default headers values
func NewGetConversationParticipantWrapupInternalServerError() *GetConversationParticipantWrapupInternalServerError {
	return &GetConversationParticipantWrapupInternalServerError{}
}

/*
GetConversationParticipantWrapupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup internal server error response has a 2xx status code
func (o *GetConversationParticipantWrapupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup internal server error response has a 3xx status code
func (o *GetConversationParticipantWrapupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup internal server error response has a 4xx status code
func (o *GetConversationParticipantWrapupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant wrapup internal server error response has a 5xx status code
func (o *GetConversationParticipantWrapupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant wrapup internal server error response a status code equal to that given
func (o *GetConversationParticipantWrapupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationParticipantWrapupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationParticipantWrapupInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationParticipantWrapupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupServiceUnavailable creates a GetConversationParticipantWrapupServiceUnavailable with default headers values
func NewGetConversationParticipantWrapupServiceUnavailable() *GetConversationParticipantWrapupServiceUnavailable {
	return &GetConversationParticipantWrapupServiceUnavailable{}
}

/*
GetConversationParticipantWrapupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup service unavailable response has a 2xx status code
func (o *GetConversationParticipantWrapupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup service unavailable response has a 3xx status code
func (o *GetConversationParticipantWrapupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup service unavailable response has a 4xx status code
func (o *GetConversationParticipantWrapupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant wrapup service unavailable response has a 5xx status code
func (o *GetConversationParticipantWrapupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant wrapup service unavailable response a status code equal to that given
func (o *GetConversationParticipantWrapupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationParticipantWrapupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationParticipantWrapupServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationParticipantWrapupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupGatewayTimeout creates a GetConversationParticipantWrapupGatewayTimeout with default headers values
func NewGetConversationParticipantWrapupGatewayTimeout() *GetConversationParticipantWrapupGatewayTimeout {
	return &GetConversationParticipantWrapupGatewayTimeout{}
}

/*
GetConversationParticipantWrapupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant wrapup gateway timeout response has a 2xx status code
func (o *GetConversationParticipantWrapupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant wrapup gateway timeout response has a 3xx status code
func (o *GetConversationParticipantWrapupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant wrapup gateway timeout response has a 4xx status code
func (o *GetConversationParticipantWrapupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant wrapup gateway timeout response has a 5xx status code
func (o *GetConversationParticipantWrapupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant wrapup gateway timeout response a status code equal to that given
func (o *GetConversationParticipantWrapupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationParticipantWrapupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationParticipantWrapupGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapup][%d] getConversationParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationParticipantWrapupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
