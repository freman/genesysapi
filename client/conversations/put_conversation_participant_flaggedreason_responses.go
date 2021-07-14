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

/*PutConversationParticipantFlaggedreasonNoContent handles this case with default header values.

The flagged reason was set successfully.
*/
type PutConversationParticipantFlaggedreasonNoContent struct {
}

func (o *PutConversationParticipantFlaggedreasonNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v2/conversations/{conversationId}/participants/{participantId}/flaggedreason][%d] putConversationParticipantFlaggedreasonNoContent ", 204)
}

func (o *PutConversationParticipantFlaggedreasonNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutConversationParticipantFlaggedreasonBadRequest creates a PutConversationParticipantFlaggedreasonBadRequest with default headers values
func NewPutConversationParticipantFlaggedreasonBadRequest() *PutConversationParticipantFlaggedreasonBadRequest {
	return &PutConversationParticipantFlaggedreasonBadRequest{}
}

/*PutConversationParticipantFlaggedreasonBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutConversationParticipantFlaggedreasonBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonBadRequest) Error() string {
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

/*PutConversationParticipantFlaggedreasonUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutConversationParticipantFlaggedreasonUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonUnauthorized) Error() string {
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

/*PutConversationParticipantFlaggedreasonForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutConversationParticipantFlaggedreasonForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonForbidden) Error() string {
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

/*PutConversationParticipantFlaggedreasonNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutConversationParticipantFlaggedreasonNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonNotFound) Error() string {
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

/*PutConversationParticipantFlaggedreasonRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutConversationParticipantFlaggedreasonRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonRequestTimeout) Error() string {
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

/*PutConversationParticipantFlaggedreasonRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutConversationParticipantFlaggedreasonRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonRequestEntityTooLarge) Error() string {
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

/*PutConversationParticipantFlaggedreasonUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutConversationParticipantFlaggedreasonUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonUnsupportedMediaType) Error() string {
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

/*PutConversationParticipantFlaggedreasonTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutConversationParticipantFlaggedreasonTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonTooManyRequests) Error() string {
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

/*PutConversationParticipantFlaggedreasonInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutConversationParticipantFlaggedreasonInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonInternalServerError) Error() string {
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

/*PutConversationParticipantFlaggedreasonServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutConversationParticipantFlaggedreasonServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonServiceUnavailable) Error() string {
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

/*PutConversationParticipantFlaggedreasonGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutConversationParticipantFlaggedreasonGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutConversationParticipantFlaggedreasonGatewayTimeout) Error() string {
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
