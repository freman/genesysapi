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

// GetConversationParticipantSecureivrsessionReader is a Reader for the GetConversationParticipantSecureivrsession structure.
type GetConversationParticipantSecureivrsessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationParticipantSecureivrsessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationParticipantSecureivrsessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationParticipantSecureivrsessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationParticipantSecureivrsessionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationParticipantSecureivrsessionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationParticipantSecureivrsessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationParticipantSecureivrsessionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationParticipantSecureivrsessionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationParticipantSecureivrsessionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationParticipantSecureivrsessionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationParticipantSecureivrsessionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationParticipantSecureivrsessionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationParticipantSecureivrsessionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationParticipantSecureivrsessionOK creates a GetConversationParticipantSecureivrsessionOK with default headers values
func NewGetConversationParticipantSecureivrsessionOK() *GetConversationParticipantSecureivrsessionOK {
	return &GetConversationParticipantSecureivrsessionOK{}
}

/*
GetConversationParticipantSecureivrsessionOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationParticipantSecureivrsessionOK struct {
	Payload *models.SecureSession
}

// IsSuccess returns true when this get conversation participant secureivrsession o k response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation participant secureivrsession o k response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession o k response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant secureivrsession o k response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession o k response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationParticipantSecureivrsessionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionOK  %+v", 200, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionOK  %+v", 200, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionOK) GetPayload() *models.SecureSession {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecureSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionBadRequest creates a GetConversationParticipantSecureivrsessionBadRequest with default headers values
func NewGetConversationParticipantSecureivrsessionBadRequest() *GetConversationParticipantSecureivrsessionBadRequest {
	return &GetConversationParticipantSecureivrsessionBadRequest{}
}

/*
GetConversationParticipantSecureivrsessionBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationParticipantSecureivrsessionBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession bad request response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession bad request response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession bad request response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession bad request response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession bad request response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationParticipantSecureivrsessionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionUnauthorized creates a GetConversationParticipantSecureivrsessionUnauthorized with default headers values
func NewGetConversationParticipantSecureivrsessionUnauthorized() *GetConversationParticipantSecureivrsessionUnauthorized {
	return &GetConversationParticipantSecureivrsessionUnauthorized{}
}

/*
GetConversationParticipantSecureivrsessionUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationParticipantSecureivrsessionUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession unauthorized response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession unauthorized response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession unauthorized response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession unauthorized response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession unauthorized response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationParticipantSecureivrsessionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionForbidden creates a GetConversationParticipantSecureivrsessionForbidden with default headers values
func NewGetConversationParticipantSecureivrsessionForbidden() *GetConversationParticipantSecureivrsessionForbidden {
	return &GetConversationParticipantSecureivrsessionForbidden{}
}

/*
GetConversationParticipantSecureivrsessionForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationParticipantSecureivrsessionForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession forbidden response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession forbidden response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession forbidden response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession forbidden response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession forbidden response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationParticipantSecureivrsessionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionNotFound creates a GetConversationParticipantSecureivrsessionNotFound with default headers values
func NewGetConversationParticipantSecureivrsessionNotFound() *GetConversationParticipantSecureivrsessionNotFound {
	return &GetConversationParticipantSecureivrsessionNotFound{}
}

/*
GetConversationParticipantSecureivrsessionNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationParticipantSecureivrsessionNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession not found response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession not found response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession not found response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession not found response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession not found response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationParticipantSecureivrsessionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionRequestTimeout creates a GetConversationParticipantSecureivrsessionRequestTimeout with default headers values
func NewGetConversationParticipantSecureivrsessionRequestTimeout() *GetConversationParticipantSecureivrsessionRequestTimeout {
	return &GetConversationParticipantSecureivrsessionRequestTimeout{}
}

/*
GetConversationParticipantSecureivrsessionRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationParticipantSecureivrsessionRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession request timeout response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession request timeout response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession request timeout response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession request timeout response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession request timeout response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationParticipantSecureivrsessionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionRequestEntityTooLarge creates a GetConversationParticipantSecureivrsessionRequestEntityTooLarge with default headers values
func NewGetConversationParticipantSecureivrsessionRequestEntityTooLarge() *GetConversationParticipantSecureivrsessionRequestEntityTooLarge {
	return &GetConversationParticipantSecureivrsessionRequestEntityTooLarge{}
}

/*
GetConversationParticipantSecureivrsessionRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationParticipantSecureivrsessionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession request entity too large response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession request entity too large response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession request entity too large response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession request entity too large response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession request entity too large response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionUnsupportedMediaType creates a GetConversationParticipantSecureivrsessionUnsupportedMediaType with default headers values
func NewGetConversationParticipantSecureivrsessionUnsupportedMediaType() *GetConversationParticipantSecureivrsessionUnsupportedMediaType {
	return &GetConversationParticipantSecureivrsessionUnsupportedMediaType{}
}

/*
GetConversationParticipantSecureivrsessionUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationParticipantSecureivrsessionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession unsupported media type response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession unsupported media type response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession unsupported media type response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession unsupported media type response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession unsupported media type response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionTooManyRequests creates a GetConversationParticipantSecureivrsessionTooManyRequests with default headers values
func NewGetConversationParticipantSecureivrsessionTooManyRequests() *GetConversationParticipantSecureivrsessionTooManyRequests {
	return &GetConversationParticipantSecureivrsessionTooManyRequests{}
}

/*
GetConversationParticipantSecureivrsessionTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationParticipantSecureivrsessionTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession too many requests response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession too many requests response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession too many requests response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation participant secureivrsession too many requests response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation participant secureivrsession too many requests response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationParticipantSecureivrsessionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionInternalServerError creates a GetConversationParticipantSecureivrsessionInternalServerError with default headers values
func NewGetConversationParticipantSecureivrsessionInternalServerError() *GetConversationParticipantSecureivrsessionInternalServerError {
	return &GetConversationParticipantSecureivrsessionInternalServerError{}
}

/*
GetConversationParticipantSecureivrsessionInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationParticipantSecureivrsessionInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession internal server error response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession internal server error response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession internal server error response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant secureivrsession internal server error response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant secureivrsession internal server error response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationParticipantSecureivrsessionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionServiceUnavailable creates a GetConversationParticipantSecureivrsessionServiceUnavailable with default headers values
func NewGetConversationParticipantSecureivrsessionServiceUnavailable() *GetConversationParticipantSecureivrsessionServiceUnavailable {
	return &GetConversationParticipantSecureivrsessionServiceUnavailable{}
}

/*
GetConversationParticipantSecureivrsessionServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationParticipantSecureivrsessionServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession service unavailable response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession service unavailable response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession service unavailable response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant secureivrsession service unavailable response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant secureivrsession service unavailable response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantSecureivrsessionGatewayTimeout creates a GetConversationParticipantSecureivrsessionGatewayTimeout with default headers values
func NewGetConversationParticipantSecureivrsessionGatewayTimeout() *GetConversationParticipantSecureivrsessionGatewayTimeout {
	return &GetConversationParticipantSecureivrsessionGatewayTimeout{}
}

/*
GetConversationParticipantSecureivrsessionGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationParticipantSecureivrsessionGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation participant secureivrsession gateway timeout response has a 2xx status code
func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation participant secureivrsession gateway timeout response has a 3xx status code
func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation participant secureivrsession gateway timeout response has a 4xx status code
func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation participant secureivrsession gateway timeout response has a 5xx status code
func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation participant secureivrsession gateway timeout response a status code equal to that given
func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/secureivrsessions/{secureSessionId}][%d] getConversationParticipantSecureivrsessionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantSecureivrsessionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
