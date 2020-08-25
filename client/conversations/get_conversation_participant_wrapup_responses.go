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

/*GetConversationParticipantWrapupOK handles this case with default header values.

successful operation
*/
type GetConversationParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

func (o *GetConversationParticipantWrapupOK) Error() string {
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

/*GetConversationParticipantWrapupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupBadRequest) Error() string {
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

/*GetConversationParticipantWrapupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupUnauthorized) Error() string {
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

/*GetConversationParticipantWrapupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupForbidden) Error() string {
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

/*GetConversationParticipantWrapupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupNotFound) Error() string {
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

// NewGetConversationParticipantWrapupRequestEntityTooLarge creates a GetConversationParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationParticipantWrapupRequestEntityTooLarge() *GetConversationParticipantWrapupRequestEntityTooLarge {
	return &GetConversationParticipantWrapupRequestEntityTooLarge{}
}

/*GetConversationParticipantWrapupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupRequestEntityTooLarge) Error() string {
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

/*GetConversationParticipantWrapupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupUnsupportedMediaType) Error() string {
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

/*GetConversationParticipantWrapupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupTooManyRequests) Error() string {
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

/*GetConversationParticipantWrapupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupInternalServerError) Error() string {
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

/*GetConversationParticipantWrapupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupServiceUnavailable) Error() string {
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

/*GetConversationParticipantWrapupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupGatewayTimeout) Error() string {
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