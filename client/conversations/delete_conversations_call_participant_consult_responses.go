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

// DeleteConversationsCallParticipantConsultReader is a Reader for the DeleteConversationsCallParticipantConsult structure.
type DeleteConversationsCallParticipantConsultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsCallParticipantConsultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConversationsCallParticipantConsultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsCallParticipantConsultBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsCallParticipantConsultUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsCallParticipantConsultForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsCallParticipantConsultNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteConversationsCallParticipantConsultRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsCallParticipantConsultRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsCallParticipantConsultUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsCallParticipantConsultTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsCallParticipantConsultInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsCallParticipantConsultServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsCallParticipantConsultGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsCallParticipantConsultOK creates a DeleteConversationsCallParticipantConsultOK with default headers values
func NewDeleteConversationsCallParticipantConsultOK() *DeleteConversationsCallParticipantConsultOK {
	return &DeleteConversationsCallParticipantConsultOK{}
}

/*
DeleteConversationsCallParticipantConsultOK describes a response with status code 200, with default header values.

Operation was successful.
*/
type DeleteConversationsCallParticipantConsultOK struct {
}

// IsSuccess returns true when this delete conversations call participant consult o k response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete conversations call participant consult o k response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult o k response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversations call participant consult o k response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult o k response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteConversationsCallParticipantConsultOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultOK ", 200)
}

func (o *DeleteConversationsCallParticipantConsultOK) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultOK ", 200)
}

func (o *DeleteConversationsCallParticipantConsultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsCallParticipantConsultBadRequest creates a DeleteConversationsCallParticipantConsultBadRequest with default headers values
func NewDeleteConversationsCallParticipantConsultBadRequest() *DeleteConversationsCallParticipantConsultBadRequest {
	return &DeleteConversationsCallParticipantConsultBadRequest{}
}

/*
DeleteConversationsCallParticipantConsultBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsCallParticipantConsultBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult bad request response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult bad request response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult bad request response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult bad request response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult bad request response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteConversationsCallParticipantConsultBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultUnauthorized creates a DeleteConversationsCallParticipantConsultUnauthorized with default headers values
func NewDeleteConversationsCallParticipantConsultUnauthorized() *DeleteConversationsCallParticipantConsultUnauthorized {
	return &DeleteConversationsCallParticipantConsultUnauthorized{}
}

/*
DeleteConversationsCallParticipantConsultUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsCallParticipantConsultUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult unauthorized response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult unauthorized response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult unauthorized response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult unauthorized response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult unauthorized response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteConversationsCallParticipantConsultUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultForbidden creates a DeleteConversationsCallParticipantConsultForbidden with default headers values
func NewDeleteConversationsCallParticipantConsultForbidden() *DeleteConversationsCallParticipantConsultForbidden {
	return &DeleteConversationsCallParticipantConsultForbidden{}
}

/*
DeleteConversationsCallParticipantConsultForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsCallParticipantConsultForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult forbidden response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult forbidden response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult forbidden response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult forbidden response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult forbidden response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteConversationsCallParticipantConsultForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultNotFound creates a DeleteConversationsCallParticipantConsultNotFound with default headers values
func NewDeleteConversationsCallParticipantConsultNotFound() *DeleteConversationsCallParticipantConsultNotFound {
	return &DeleteConversationsCallParticipantConsultNotFound{}
}

/*
DeleteConversationsCallParticipantConsultNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteConversationsCallParticipantConsultNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult not found response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult not found response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult not found response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult not found response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult not found response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteConversationsCallParticipantConsultNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultRequestTimeout creates a DeleteConversationsCallParticipantConsultRequestTimeout with default headers values
func NewDeleteConversationsCallParticipantConsultRequestTimeout() *DeleteConversationsCallParticipantConsultRequestTimeout {
	return &DeleteConversationsCallParticipantConsultRequestTimeout{}
}

/*
DeleteConversationsCallParticipantConsultRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteConversationsCallParticipantConsultRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult request timeout response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult request timeout response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult request timeout response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult request timeout response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult request timeout response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteConversationsCallParticipantConsultRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultRequestEntityTooLarge creates a DeleteConversationsCallParticipantConsultRequestEntityTooLarge with default headers values
func NewDeleteConversationsCallParticipantConsultRequestEntityTooLarge() *DeleteConversationsCallParticipantConsultRequestEntityTooLarge {
	return &DeleteConversationsCallParticipantConsultRequestEntityTooLarge{}
}

/*
DeleteConversationsCallParticipantConsultRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteConversationsCallParticipantConsultRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult request entity too large response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult request entity too large response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult request entity too large response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult request entity too large response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult request entity too large response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultUnsupportedMediaType creates a DeleteConversationsCallParticipantConsultUnsupportedMediaType with default headers values
func NewDeleteConversationsCallParticipantConsultUnsupportedMediaType() *DeleteConversationsCallParticipantConsultUnsupportedMediaType {
	return &DeleteConversationsCallParticipantConsultUnsupportedMediaType{}
}

/*
DeleteConversationsCallParticipantConsultUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsCallParticipantConsultUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult unsupported media type response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult unsupported media type response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult unsupported media type response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult unsupported media type response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult unsupported media type response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultTooManyRequests creates a DeleteConversationsCallParticipantConsultTooManyRequests with default headers values
func NewDeleteConversationsCallParticipantConsultTooManyRequests() *DeleteConversationsCallParticipantConsultTooManyRequests {
	return &DeleteConversationsCallParticipantConsultTooManyRequests{}
}

/*
DeleteConversationsCallParticipantConsultTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteConversationsCallParticipantConsultTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult too many requests response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult too many requests response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult too many requests response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete conversations call participant consult too many requests response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete conversations call participant consult too many requests response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteConversationsCallParticipantConsultTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultInternalServerError creates a DeleteConversationsCallParticipantConsultInternalServerError with default headers values
func NewDeleteConversationsCallParticipantConsultInternalServerError() *DeleteConversationsCallParticipantConsultInternalServerError {
	return &DeleteConversationsCallParticipantConsultInternalServerError{}
}

/*
DeleteConversationsCallParticipantConsultInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsCallParticipantConsultInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult internal server error response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult internal server error response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult internal server error response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversations call participant consult internal server error response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversations call participant consult internal server error response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteConversationsCallParticipantConsultInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultServiceUnavailable creates a DeleteConversationsCallParticipantConsultServiceUnavailable with default headers values
func NewDeleteConversationsCallParticipantConsultServiceUnavailable() *DeleteConversationsCallParticipantConsultServiceUnavailable {
	return &DeleteConversationsCallParticipantConsultServiceUnavailable{}
}

/*
DeleteConversationsCallParticipantConsultServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsCallParticipantConsultServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult service unavailable response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult service unavailable response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult service unavailable response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversations call participant consult service unavailable response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversations call participant consult service unavailable response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsCallParticipantConsultGatewayTimeout creates a DeleteConversationsCallParticipantConsultGatewayTimeout with default headers values
func NewDeleteConversationsCallParticipantConsultGatewayTimeout() *DeleteConversationsCallParticipantConsultGatewayTimeout {
	return &DeleteConversationsCallParticipantConsultGatewayTimeout{}
}

/*
DeleteConversationsCallParticipantConsultGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteConversationsCallParticipantConsultGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete conversations call participant consult gateway timeout response has a 2xx status code
func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete conversations call participant consult gateway timeout response has a 3xx status code
func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete conversations call participant consult gateway timeout response has a 4xx status code
func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete conversations call participant consult gateway timeout response has a 5xx status code
func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete conversations call participant consult gateway timeout response a status code equal to that given
func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/calls/{conversationId}/participants/{participantId}/consult][%d] deleteConversationsCallParticipantConsultGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsCallParticipantConsultGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
