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

// GetConversationsCobrowsesessionParticipantWrapupReader is a Reader for the GetConversationsCobrowsesessionParticipantWrapup structure.
type GetConversationsCobrowsesessionParticipantWrapupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCobrowsesessionParticipantWrapupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCobrowsesessionParticipantWrapupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCobrowsesessionParticipantWrapupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCobrowsesessionParticipantWrapupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCobrowsesessionParticipantWrapupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCobrowsesessionParticipantWrapupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCobrowsesessionParticipantWrapupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCobrowsesessionParticipantWrapupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCobrowsesessionParticipantWrapupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCobrowsesessionParticipantWrapupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCobrowsesessionParticipantWrapupOK creates a GetConversationsCobrowsesessionParticipantWrapupOK with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupOK() *GetConversationsCobrowsesessionParticipantWrapupOK {
	return &GetConversationsCobrowsesessionParticipantWrapupOK{}
}

/*GetConversationsCobrowsesessionParticipantWrapupOK handles this case with default header values.

successful operation
*/
type GetConversationsCobrowsesessionParticipantWrapupOK struct {
	Payload *models.AssignedWrapupCode
}

func (o *GetConversationsCobrowsesessionParticipantWrapupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupOK) GetPayload() *models.AssignedWrapupCode {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignedWrapupCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupBadRequest creates a GetConversationsCobrowsesessionParticipantWrapupBadRequest with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupBadRequest() *GetConversationsCobrowsesessionParticipantWrapupBadRequest {
	return &GetConversationsCobrowsesessionParticipantWrapupBadRequest{}
}

/*GetConversationsCobrowsesessionParticipantWrapupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCobrowsesessionParticipantWrapupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupUnauthorized creates a GetConversationsCobrowsesessionParticipantWrapupUnauthorized with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupUnauthorized() *GetConversationsCobrowsesessionParticipantWrapupUnauthorized {
	return &GetConversationsCobrowsesessionParticipantWrapupUnauthorized{}
}

/*GetConversationsCobrowsesessionParticipantWrapupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCobrowsesessionParticipantWrapupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupForbidden creates a GetConversationsCobrowsesessionParticipantWrapupForbidden with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupForbidden() *GetConversationsCobrowsesessionParticipantWrapupForbidden {
	return &GetConversationsCobrowsesessionParticipantWrapupForbidden{}
}

/*GetConversationsCobrowsesessionParticipantWrapupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCobrowsesessionParticipantWrapupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupNotFound creates a GetConversationsCobrowsesessionParticipantWrapupNotFound with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupNotFound() *GetConversationsCobrowsesessionParticipantWrapupNotFound {
	return &GetConversationsCobrowsesessionParticipantWrapupNotFound{}
}

/*GetConversationsCobrowsesessionParticipantWrapupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsCobrowsesessionParticipantWrapupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge creates a GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge() *GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge {
	return &GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge{}
}

/*GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType creates a GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType() *GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType {
	return &GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType{}
}

/*GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupTooManyRequests creates a GetConversationsCobrowsesessionParticipantWrapupTooManyRequests with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupTooManyRequests() *GetConversationsCobrowsesessionParticipantWrapupTooManyRequests {
	return &GetConversationsCobrowsesessionParticipantWrapupTooManyRequests{}
}

/*GetConversationsCobrowsesessionParticipantWrapupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsCobrowsesessionParticipantWrapupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupInternalServerError creates a GetConversationsCobrowsesessionParticipantWrapupInternalServerError with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupInternalServerError() *GetConversationsCobrowsesessionParticipantWrapupInternalServerError {
	return &GetConversationsCobrowsesessionParticipantWrapupInternalServerError{}
}

/*GetConversationsCobrowsesessionParticipantWrapupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCobrowsesessionParticipantWrapupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupServiceUnavailable creates a GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupServiceUnavailable() *GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable {
	return &GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable{}
}

/*GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCobrowsesessionParticipantWrapupGatewayTimeout creates a GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout with default headers values
func NewGetConversationsCobrowsesessionParticipantWrapupGatewayTimeout() *GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout {
	return &GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout{}
}

/*GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/cobrowsesessions/{conversationId}/participants/{participantId}/wrapup][%d] getConversationsCobrowsesessionParticipantWrapupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCobrowsesessionParticipantWrapupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
