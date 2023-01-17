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

// DeleteConversationParticipantCodeReader is a Reader for the DeleteConversationParticipantCode structure.
type DeleteConversationParticipantCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationParticipantCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewDeleteConversationParticipantCodeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationParticipantCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationParticipantCodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationParticipantCodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteConversationParticipantCodeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationParticipantCodeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationParticipantCodeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationParticipantCodeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationParticipantCodeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationParticipantCodeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationParticipantCodeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteConversationParticipantCodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteConversationParticipantCodeBadRequest creates a DeleteConversationParticipantCodeBadRequest with default headers values
func NewDeleteConversationParticipantCodeBadRequest() *DeleteConversationParticipantCodeBadRequest {
	return &DeleteConversationParticipantCodeBadRequest{}
}

/*
DeleteConversationParticipantCodeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationParticipantCodeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code bad request response has a 2xx status code
func (o *DeleteConversationParticipantCodeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code bad request response has a 3xx status code
func (o *DeleteConversationParticipantCodeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code bad request response has a 4xx status code
func (o *DeleteConversationParticipantCodeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code bad request response has a 5xx status code
func (o *DeleteConversationParticipantCodeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code bad request response a status code equal to that given
func (o *DeleteConversationParticipantCodeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteConversationParticipantCodeBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationParticipantCodeBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationParticipantCodeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeUnauthorized creates a DeleteConversationParticipantCodeUnauthorized with default headers values
func NewDeleteConversationParticipantCodeUnauthorized() *DeleteConversationParticipantCodeUnauthorized {
	return &DeleteConversationParticipantCodeUnauthorized{}
}

/*
DeleteConversationParticipantCodeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationParticipantCodeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code unauthorized response has a 2xx status code
func (o *DeleteConversationParticipantCodeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code unauthorized response has a 3xx status code
func (o *DeleteConversationParticipantCodeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code unauthorized response has a 4xx status code
func (o *DeleteConversationParticipantCodeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code unauthorized response has a 5xx status code
func (o *DeleteConversationParticipantCodeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code unauthorized response a status code equal to that given
func (o *DeleteConversationParticipantCodeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteConversationParticipantCodeUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationParticipantCodeUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationParticipantCodeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeForbidden creates a DeleteConversationParticipantCodeForbidden with default headers values
func NewDeleteConversationParticipantCodeForbidden() *DeleteConversationParticipantCodeForbidden {
	return &DeleteConversationParticipantCodeForbidden{}
}

/*
DeleteConversationParticipantCodeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationParticipantCodeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code forbidden response has a 2xx status code
func (o *DeleteConversationParticipantCodeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code forbidden response has a 3xx status code
func (o *DeleteConversationParticipantCodeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code forbidden response has a 4xx status code
func (o *DeleteConversationParticipantCodeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code forbidden response has a 5xx status code
func (o *DeleteConversationParticipantCodeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code forbidden response a status code equal to that given
func (o *DeleteConversationParticipantCodeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteConversationParticipantCodeForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationParticipantCodeForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationParticipantCodeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeNotFound creates a DeleteConversationParticipantCodeNotFound with default headers values
func NewDeleteConversationParticipantCodeNotFound() *DeleteConversationParticipantCodeNotFound {
	return &DeleteConversationParticipantCodeNotFound{}
}

/*
DeleteConversationParticipantCodeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteConversationParticipantCodeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code not found response has a 2xx status code
func (o *DeleteConversationParticipantCodeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code not found response has a 3xx status code
func (o *DeleteConversationParticipantCodeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code not found response has a 4xx status code
func (o *DeleteConversationParticipantCodeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code not found response has a 5xx status code
func (o *DeleteConversationParticipantCodeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code not found response a status code equal to that given
func (o *DeleteConversationParticipantCodeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteConversationParticipantCodeNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationParticipantCodeNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationParticipantCodeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeRequestTimeout creates a DeleteConversationParticipantCodeRequestTimeout with default headers values
func NewDeleteConversationParticipantCodeRequestTimeout() *DeleteConversationParticipantCodeRequestTimeout {
	return &DeleteConversationParticipantCodeRequestTimeout{}
}

/*
DeleteConversationParticipantCodeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteConversationParticipantCodeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code request timeout response has a 2xx status code
func (o *DeleteConversationParticipantCodeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code request timeout response has a 3xx status code
func (o *DeleteConversationParticipantCodeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code request timeout response has a 4xx status code
func (o *DeleteConversationParticipantCodeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code request timeout response has a 5xx status code
func (o *DeleteConversationParticipantCodeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code request timeout response a status code equal to that given
func (o *DeleteConversationParticipantCodeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteConversationParticipantCodeRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationParticipantCodeRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationParticipantCodeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeRequestEntityTooLarge creates a DeleteConversationParticipantCodeRequestEntityTooLarge with default headers values
func NewDeleteConversationParticipantCodeRequestEntityTooLarge() *DeleteConversationParticipantCodeRequestEntityTooLarge {
	return &DeleteConversationParticipantCodeRequestEntityTooLarge{}
}

/*
DeleteConversationParticipantCodeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteConversationParticipantCodeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code request entity too large response has a 2xx status code
func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code request entity too large response has a 3xx status code
func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code request entity too large response has a 4xx status code
func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code request entity too large response has a 5xx status code
func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code request entity too large response a status code equal to that given
func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeUnsupportedMediaType creates a DeleteConversationParticipantCodeUnsupportedMediaType with default headers values
func NewDeleteConversationParticipantCodeUnsupportedMediaType() *DeleteConversationParticipantCodeUnsupportedMediaType {
	return &DeleteConversationParticipantCodeUnsupportedMediaType{}
}

/*
DeleteConversationParticipantCodeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationParticipantCodeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code unsupported media type response has a 2xx status code
func (o *DeleteConversationParticipantCodeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code unsupported media type response has a 3xx status code
func (o *DeleteConversationParticipantCodeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code unsupported media type response has a 4xx status code
func (o *DeleteConversationParticipantCodeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code unsupported media type response has a 5xx status code
func (o *DeleteConversationParticipantCodeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code unsupported media type response a status code equal to that given
func (o *DeleteConversationParticipantCodeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteConversationParticipantCodeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationParticipantCodeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationParticipantCodeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeTooManyRequests creates a DeleteConversationParticipantCodeTooManyRequests with default headers values
func NewDeleteConversationParticipantCodeTooManyRequests() *DeleteConversationParticipantCodeTooManyRequests {
	return &DeleteConversationParticipantCodeTooManyRequests{}
}

/*
DeleteConversationParticipantCodeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteConversationParticipantCodeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code too many requests response has a 2xx status code
func (o *DeleteConversationParticipantCodeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code too many requests response has a 3xx status code
func (o *DeleteConversationParticipantCodeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code too many requests response has a 4xx status code
func (o *DeleteConversationParticipantCodeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversation participant code too many requests response has a 5xx status code
func (o *DeleteConversationParticipantCodeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversation participant code too many requests response a status code equal to that given
func (o *DeleteConversationParticipantCodeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteConversationParticipantCodeTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationParticipantCodeTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationParticipantCodeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeInternalServerError creates a DeleteConversationParticipantCodeInternalServerError with default headers values
func NewDeleteConversationParticipantCodeInternalServerError() *DeleteConversationParticipantCodeInternalServerError {
	return &DeleteConversationParticipantCodeInternalServerError{}
}

/*
DeleteConversationParticipantCodeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationParticipantCodeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code internal server error response has a 2xx status code
func (o *DeleteConversationParticipantCodeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code internal server error response has a 3xx status code
func (o *DeleteConversationParticipantCodeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code internal server error response has a 4xx status code
func (o *DeleteConversationParticipantCodeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversation participant code internal server error response has a 5xx status code
func (o *DeleteConversationParticipantCodeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversation participant code internal server error response a status code equal to that given
func (o *DeleteConversationParticipantCodeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteConversationParticipantCodeInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationParticipantCodeInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationParticipantCodeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeServiceUnavailable creates a DeleteConversationParticipantCodeServiceUnavailable with default headers values
func NewDeleteConversationParticipantCodeServiceUnavailable() *DeleteConversationParticipantCodeServiceUnavailable {
	return &DeleteConversationParticipantCodeServiceUnavailable{}
}

/*
DeleteConversationParticipantCodeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationParticipantCodeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code service unavailable response has a 2xx status code
func (o *DeleteConversationParticipantCodeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code service unavailable response has a 3xx status code
func (o *DeleteConversationParticipantCodeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code service unavailable response has a 4xx status code
func (o *DeleteConversationParticipantCodeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversation participant code service unavailable response has a 5xx status code
func (o *DeleteConversationParticipantCodeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversation participant code service unavailable response a status code equal to that given
func (o *DeleteConversationParticipantCodeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteConversationParticipantCodeServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationParticipantCodeServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationParticipantCodeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeGatewayTimeout creates a DeleteConversationParticipantCodeGatewayTimeout with default headers values
func NewDeleteConversationParticipantCodeGatewayTimeout() *DeleteConversationParticipantCodeGatewayTimeout {
	return &DeleteConversationParticipantCodeGatewayTimeout{}
}

/*
DeleteConversationParticipantCodeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteConversationParticipantCodeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversation participant code gateway timeout response has a 2xx status code
func (o *DeleteConversationParticipantCodeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversation participant code gateway timeout response has a 3xx status code
func (o *DeleteConversationParticipantCodeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversation participant code gateway timeout response has a 4xx status code
func (o *DeleteConversationParticipantCodeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversation participant code gateway timeout response has a 5xx status code
func (o *DeleteConversationParticipantCodeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversation participant code gateway timeout response a status code equal to that given
func (o *DeleteConversationParticipantCodeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteConversationParticipantCodeGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationParticipantCodeGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCodeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationParticipantCodeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantCodeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantCodeDefault creates a DeleteConversationParticipantCodeDefault with default headers values
func NewDeleteConversationParticipantCodeDefault(code int) *DeleteConversationParticipantCodeDefault {
	return &DeleteConversationParticipantCodeDefault{
		_statusCode: code,
	}
}

/*
DeleteConversationParticipantCodeDefault describes a response with status code -1, with default header values.

successful operation
*/
type DeleteConversationParticipantCodeDefault struct {
	_statusCode int
}

// Code gets the status code for the delete conversation participant code default response
func (o *DeleteConversationParticipantCodeDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this delete conversation participant code default response has a 2xx status code
func (o *DeleteConversationParticipantCodeDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this delete conversation participant code default response has a 3xx status code
func (o *DeleteConversationParticipantCodeDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this delete conversation participant code default response has a 4xx status code
func (o *DeleteConversationParticipantCodeDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this delete conversation participant code default response has a 5xx status code
func (o *DeleteConversationParticipantCodeDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this delete conversation participant code default response a status code equal to that given
func (o *DeleteConversationParticipantCodeDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DeleteConversationParticipantCodeDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCode default ", o._statusCode)
}

func (o *DeleteConversationParticipantCodeDefault) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/codes/{addCommunicationCode}][%d] deleteConversationParticipantCode default ", o._statusCode)
}

func (o *DeleteConversationParticipantCodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
