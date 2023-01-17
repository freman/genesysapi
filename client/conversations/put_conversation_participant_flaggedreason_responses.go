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

// PutConversationParticipantFlaggedreasonReader is a Reader for the PutConversationParticipantFlaggedreason structure.
type PutConversationParticipantFlaggedreasonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutConversationParticipantFlaggedreasonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutConversationParticipantFlaggedreasonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutConversationParticipantFlaggedreasonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutConversationParticipantFlaggedreasonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutConversationParticipantFlaggedreasonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutConversationParticipantFlaggedreasonNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutConversationParticipantFlaggedreasonRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutConversationParticipantFlaggedreasonRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutConversationParticipantFlaggedreasonUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutConversationParticipantFlaggedreasonTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutConversationParticipantFlaggedreasonInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutConversationParticipantFlaggedreasonServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutConversationParticipantFlaggedreasonGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutConversationParticipantFlaggedreasonNoContent creates a PutConversationParticipantFlaggedreasonNoContent with default headers values
func NewPutConversationParticipantFlaggedreasonNoContent() *PutConversationParticipantFlaggedreasonNoContent {
	return &PutConversationParticipantFlaggedreasonNoContent{}
}

/*
PutConversationParticipantFlaggedreasonNoContent describes a response with status code 204, with default header values.

The flagged reason was set successfully.
*/
type PutConversationParticipantFlaggedreasonNoContent struct {
}

// IsSuccess returns true when this put conversation participant flaggedreason no content response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put conversation participant flaggedreason no content response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason no content response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation participant flaggedreason no content response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason no content response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PutConversationParticipantFlaggedreasonNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonNoContent ", 204)
}

func (o *PutConversationParticipantFlaggedreasonNoContent) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonNoContent ", 204)
}

func (o *PutConversationParticipantFlaggedreasonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutConversationParticipantFlaggedreasonBadRequest creates a PutConversationParticipantFlaggedreasonBadRequest with default headers values
func NewPutConversationParticipantFlaggedreasonBadRequest() *PutConversationParticipantFlaggedreasonBadRequest {
	return &PutConversationParticipantFlaggedreasonBadRequest{}
}

/*
PutConversationParticipantFlaggedreasonBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationParticipantFlaggedreasonBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason bad request response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason bad request response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason bad request response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason bad request response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason bad request response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutConversationParticipantFlaggedreasonBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonBadRequest  %+v", 400, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonUnauthorized creates a PutConversationParticipantFlaggedreasonUnauthorized with default headers values
func NewPutConversationParticipantFlaggedreasonUnauthorized() *PutConversationParticipantFlaggedreasonUnauthorized {
	return &PutConversationParticipantFlaggedreasonUnauthorized{}
}

/*
PutConversationParticipantFlaggedreasonUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationParticipantFlaggedreasonUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason unauthorized response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason unauthorized response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason unauthorized response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason unauthorized response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason unauthorized response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutConversationParticipantFlaggedreasonUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonUnauthorized  %+v", 401, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonForbidden creates a PutConversationParticipantFlaggedreasonForbidden with default headers values
func NewPutConversationParticipantFlaggedreasonForbidden() *PutConversationParticipantFlaggedreasonForbidden {
	return &PutConversationParticipantFlaggedreasonForbidden{}
}

/*
PutConversationParticipantFlaggedreasonForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationParticipantFlaggedreasonForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason forbidden response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason forbidden response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason forbidden response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason forbidden response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason forbidden response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutConversationParticipantFlaggedreasonForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonForbidden  %+v", 403, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonNotFound creates a PutConversationParticipantFlaggedreasonNotFound with default headers values
func NewPutConversationParticipantFlaggedreasonNotFound() *PutConversationParticipantFlaggedreasonNotFound {
	return &PutConversationParticipantFlaggedreasonNotFound{}
}

/*
PutConversationParticipantFlaggedreasonNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutConversationParticipantFlaggedreasonNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason not found response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason not found response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason not found response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason not found response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason not found response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutConversationParticipantFlaggedreasonNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonNotFound  %+v", 404, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonRequestTimeout creates a PutConversationParticipantFlaggedreasonRequestTimeout with default headers values
func NewPutConversationParticipantFlaggedreasonRequestTimeout() *PutConversationParticipantFlaggedreasonRequestTimeout {
	return &PutConversationParticipantFlaggedreasonRequestTimeout{}
}

/*
PutConversationParticipantFlaggedreasonRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationParticipantFlaggedreasonRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason request timeout response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason request timeout response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason request timeout response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason request timeout response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason request timeout response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutConversationParticipantFlaggedreasonRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonRequestEntityTooLarge creates a PutConversationParticipantFlaggedreasonRequestEntityTooLarge with default headers values
func NewPutConversationParticipantFlaggedreasonRequestEntityTooLarge() *PutConversationParticipantFlaggedreasonRequestEntityTooLarge {
	return &PutConversationParticipantFlaggedreasonRequestEntityTooLarge{}
}

/*
PutConversationParticipantFlaggedreasonRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutConversationParticipantFlaggedreasonRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason request entity too large response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason request entity too large response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason request entity too large response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason request entity too large response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason request entity too large response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonUnsupportedMediaType creates a PutConversationParticipantFlaggedreasonUnsupportedMediaType with default headers values
func NewPutConversationParticipantFlaggedreasonUnsupportedMediaType() *PutConversationParticipantFlaggedreasonUnsupportedMediaType {
	return &PutConversationParticipantFlaggedreasonUnsupportedMediaType{}
}

/*
PutConversationParticipantFlaggedreasonUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationParticipantFlaggedreasonUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason unsupported media type response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason unsupported media type response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason unsupported media type response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason unsupported media type response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason unsupported media type response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonTooManyRequests creates a PutConversationParticipantFlaggedreasonTooManyRequests with default headers values
func NewPutConversationParticipantFlaggedreasonTooManyRequests() *PutConversationParticipantFlaggedreasonTooManyRequests {
	return &PutConversationParticipantFlaggedreasonTooManyRequests{}
}

/*
PutConversationParticipantFlaggedreasonTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationParticipantFlaggedreasonTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason too many requests response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason too many requests response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason too many requests response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put conversation participant flaggedreason too many requests response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put conversation participant flaggedreason too many requests response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutConversationParticipantFlaggedreasonTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonInternalServerError creates a PutConversationParticipantFlaggedreasonInternalServerError with default headers values
func NewPutConversationParticipantFlaggedreasonInternalServerError() *PutConversationParticipantFlaggedreasonInternalServerError {
	return &PutConversationParticipantFlaggedreasonInternalServerError{}
}

/*
PutConversationParticipantFlaggedreasonInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationParticipantFlaggedreasonInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason internal server error response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason internal server error response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason internal server error response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation participant flaggedreason internal server error response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation participant flaggedreason internal server error response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutConversationParticipantFlaggedreasonInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonInternalServerError  %+v", 500, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonServiceUnavailable creates a PutConversationParticipantFlaggedreasonServiceUnavailable with default headers values
func NewPutConversationParticipantFlaggedreasonServiceUnavailable() *PutConversationParticipantFlaggedreasonServiceUnavailable {
	return &PutConversationParticipantFlaggedreasonServiceUnavailable{}
}

/*
PutConversationParticipantFlaggedreasonServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationParticipantFlaggedreasonServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason service unavailable response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason service unavailable response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason service unavailable response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation participant flaggedreason service unavailable response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation participant flaggedreason service unavailable response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutConversationParticipantFlaggedreasonGatewayTimeout creates a PutConversationParticipantFlaggedreasonGatewayTimeout with default headers values
func NewPutConversationParticipantFlaggedreasonGatewayTimeout() *PutConversationParticipantFlaggedreasonGatewayTimeout {
	return &PutConversationParticipantFlaggedreasonGatewayTimeout{}
}

/*
PutConversationParticipantFlaggedreasonGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutConversationParticipantFlaggedreasonGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put conversation participant flaggedreason gateway timeout response has a 2xx status code
func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put conversation participant flaggedreason gateway timeout response has a 3xx status code
func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put conversation participant flaggedreason gateway timeout response has a 4xx status code
func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put conversation participant flaggedreason gateway timeout response has a 5xx status code
func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put conversation participant flaggedreason gateway timeout response a status code equal to that given
func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
