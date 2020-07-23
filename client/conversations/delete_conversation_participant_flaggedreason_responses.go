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

// DeleteConversationParticipantFlaggedreasonReader is a Reader for the DeleteConversationParticipantFlaggedreason structure.
type DeleteConversationParticipantFlaggedreasonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationParticipantFlaggedreasonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConversationParticipantFlaggedreasonNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationParticipantFlaggedreasonBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationParticipantFlaggedreasonUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationParticipantFlaggedreasonForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationParticipantFlaggedreasonNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationParticipantFlaggedreasonRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationParticipantFlaggedreasonUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationParticipantFlaggedreasonTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationParticipantFlaggedreasonInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationParticipantFlaggedreasonServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationParticipantFlaggedreasonGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationParticipantFlaggedreasonNoContent creates a DeleteConversationParticipantFlaggedreasonNoContent with default headers values
func NewDeleteConversationParticipantFlaggedreasonNoContent() *DeleteConversationParticipantFlaggedreasonNoContent {
	return &DeleteConversationParticipantFlaggedreasonNoContent{}
}

/*DeleteConversationParticipantFlaggedreasonNoContent handles this case with default header values.

The flagged reason was removed successfully.
*/
type DeleteConversationParticipantFlaggedreasonNoContent struct {
}

func (o *DeleteConversationParticipantFlaggedreasonNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonNoContent ", 204)
}

func (o *DeleteConversationParticipantFlaggedreasonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonBadRequest creates a DeleteConversationParticipantFlaggedreasonBadRequest with default headers values
func NewDeleteConversationParticipantFlaggedreasonBadRequest() *DeleteConversationParticipantFlaggedreasonBadRequest {
	return &DeleteConversationParticipantFlaggedreasonBadRequest{}
}

/*DeleteConversationParticipantFlaggedreasonBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationParticipantFlaggedreasonBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonUnauthorized creates a DeleteConversationParticipantFlaggedreasonUnauthorized with default headers values
func NewDeleteConversationParticipantFlaggedreasonUnauthorized() *DeleteConversationParticipantFlaggedreasonUnauthorized {
	return &DeleteConversationParticipantFlaggedreasonUnauthorized{}
}

/*DeleteConversationParticipantFlaggedreasonUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationParticipantFlaggedreasonUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonForbidden creates a DeleteConversationParticipantFlaggedreasonForbidden with default headers values
func NewDeleteConversationParticipantFlaggedreasonForbidden() *DeleteConversationParticipantFlaggedreasonForbidden {
	return &DeleteConversationParticipantFlaggedreasonForbidden{}
}

/*DeleteConversationParticipantFlaggedreasonForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationParticipantFlaggedreasonForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonNotFound creates a DeleteConversationParticipantFlaggedreasonNotFound with default headers values
func NewDeleteConversationParticipantFlaggedreasonNotFound() *DeleteConversationParticipantFlaggedreasonNotFound {
	return &DeleteConversationParticipantFlaggedreasonNotFound{}
}

/*DeleteConversationParticipantFlaggedreasonNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationParticipantFlaggedreasonNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonRequestEntityTooLarge creates a DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge with default headers values
func NewDeleteConversationParticipantFlaggedreasonRequestEntityTooLarge() *DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge {
	return &DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge{}
}

/*DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonUnsupportedMediaType creates a DeleteConversationParticipantFlaggedreasonUnsupportedMediaType with default headers values
func NewDeleteConversationParticipantFlaggedreasonUnsupportedMediaType() *DeleteConversationParticipantFlaggedreasonUnsupportedMediaType {
	return &DeleteConversationParticipantFlaggedreasonUnsupportedMediaType{}
}

/*DeleteConversationParticipantFlaggedreasonUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationParticipantFlaggedreasonUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonTooManyRequests creates a DeleteConversationParticipantFlaggedreasonTooManyRequests with default headers values
func NewDeleteConversationParticipantFlaggedreasonTooManyRequests() *DeleteConversationParticipantFlaggedreasonTooManyRequests {
	return &DeleteConversationParticipantFlaggedreasonTooManyRequests{}
}

/*DeleteConversationParticipantFlaggedreasonTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteConversationParticipantFlaggedreasonTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonInternalServerError creates a DeleteConversationParticipantFlaggedreasonInternalServerError with default headers values
func NewDeleteConversationParticipantFlaggedreasonInternalServerError() *DeleteConversationParticipantFlaggedreasonInternalServerError {
	return &DeleteConversationParticipantFlaggedreasonInternalServerError{}
}

/*DeleteConversationParticipantFlaggedreasonInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationParticipantFlaggedreasonInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonServiceUnavailable creates a DeleteConversationParticipantFlaggedreasonServiceUnavailable with default headers values
func NewDeleteConversationParticipantFlaggedreasonServiceUnavailable() *DeleteConversationParticipantFlaggedreasonServiceUnavailable {
	return &DeleteConversationParticipantFlaggedreasonServiceUnavailable{}
}

/*DeleteConversationParticipantFlaggedreasonServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationParticipantFlaggedreasonServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationParticipantFlaggedreasonGatewayTimeout creates a DeleteConversationParticipantFlaggedreasonGatewayTimeout with default headers values
func NewDeleteConversationParticipantFlaggedreasonGatewayTimeout() *DeleteConversationParticipantFlaggedreasonGatewayTimeout {
	return &DeleteConversationParticipantFlaggedreasonGatewayTimeout{}
}

/*DeleteConversationParticipantFlaggedreasonGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationParticipantFlaggedreasonGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationParticipantFlaggedreasonGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] deleteConversationParticipantFlaggedreasonGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationParticipantFlaggedreasonGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationParticipantFlaggedreasonGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
