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

// GetConversationParticipantWrapupcodesReader is a Reader for the GetConversationParticipantWrapupcodes structure.
type GetConversationParticipantWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationParticipantWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationParticipantWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationParticipantWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationParticipantWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationParticipantWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationParticipantWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationParticipantWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationParticipantWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationParticipantWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationParticipantWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationParticipantWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationParticipantWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationParticipantWrapupcodesOK creates a GetConversationParticipantWrapupcodesOK with default headers values
func NewGetConversationParticipantWrapupcodesOK() *GetConversationParticipantWrapupcodesOK {
	return &GetConversationParticipantWrapupcodesOK{}
}

/*GetConversationParticipantWrapupcodesOK handles this case with default header values.

successful operation
*/
type GetConversationParticipantWrapupcodesOK struct {
	Payload []*models.WrapupCode
}

func (o *GetConversationParticipantWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesOK) GetPayload() []*models.WrapupCode {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesBadRequest creates a GetConversationParticipantWrapupcodesBadRequest with default headers values
func NewGetConversationParticipantWrapupcodesBadRequest() *GetConversationParticipantWrapupcodesBadRequest {
	return &GetConversationParticipantWrapupcodesBadRequest{}
}

/*GetConversationParticipantWrapupcodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationParticipantWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesUnauthorized creates a GetConversationParticipantWrapupcodesUnauthorized with default headers values
func NewGetConversationParticipantWrapupcodesUnauthorized() *GetConversationParticipantWrapupcodesUnauthorized {
	return &GetConversationParticipantWrapupcodesUnauthorized{}
}

/*GetConversationParticipantWrapupcodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationParticipantWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesForbidden creates a GetConversationParticipantWrapupcodesForbidden with default headers values
func NewGetConversationParticipantWrapupcodesForbidden() *GetConversationParticipantWrapupcodesForbidden {
	return &GetConversationParticipantWrapupcodesForbidden{}
}

/*GetConversationParticipantWrapupcodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationParticipantWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesNotFound creates a GetConversationParticipantWrapupcodesNotFound with default headers values
func NewGetConversationParticipantWrapupcodesNotFound() *GetConversationParticipantWrapupcodesNotFound {
	return &GetConversationParticipantWrapupcodesNotFound{}
}

/*GetConversationParticipantWrapupcodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationParticipantWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesRequestEntityTooLarge creates a GetConversationParticipantWrapupcodesRequestEntityTooLarge with default headers values
func NewGetConversationParticipantWrapupcodesRequestEntityTooLarge() *GetConversationParticipantWrapupcodesRequestEntityTooLarge {
	return &GetConversationParticipantWrapupcodesRequestEntityTooLarge{}
}

/*GetConversationParticipantWrapupcodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationParticipantWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesUnsupportedMediaType creates a GetConversationParticipantWrapupcodesUnsupportedMediaType with default headers values
func NewGetConversationParticipantWrapupcodesUnsupportedMediaType() *GetConversationParticipantWrapupcodesUnsupportedMediaType {
	return &GetConversationParticipantWrapupcodesUnsupportedMediaType{}
}

/*GetConversationParticipantWrapupcodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationParticipantWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesTooManyRequests creates a GetConversationParticipantWrapupcodesTooManyRequests with default headers values
func NewGetConversationParticipantWrapupcodesTooManyRequests() *GetConversationParticipantWrapupcodesTooManyRequests {
	return &GetConversationParticipantWrapupcodesTooManyRequests{}
}

/*GetConversationParticipantWrapupcodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationParticipantWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesInternalServerError creates a GetConversationParticipantWrapupcodesInternalServerError with default headers values
func NewGetConversationParticipantWrapupcodesInternalServerError() *GetConversationParticipantWrapupcodesInternalServerError {
	return &GetConversationParticipantWrapupcodesInternalServerError{}
}

/*GetConversationParticipantWrapupcodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationParticipantWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesServiceUnavailable creates a GetConversationParticipantWrapupcodesServiceUnavailable with default headers values
func NewGetConversationParticipantWrapupcodesServiceUnavailable() *GetConversationParticipantWrapupcodesServiceUnavailable {
	return &GetConversationParticipantWrapupcodesServiceUnavailable{}
}

/*GetConversationParticipantWrapupcodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationParticipantWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationParticipantWrapupcodesGatewayTimeout creates a GetConversationParticipantWrapupcodesGatewayTimeout with default headers values
func NewGetConversationParticipantWrapupcodesGatewayTimeout() *GetConversationParticipantWrapupcodesGatewayTimeout {
	return &GetConversationParticipantWrapupcodesGatewayTimeout{}
}

/*GetConversationParticipantWrapupcodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationParticipantWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationParticipantWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationParticipantWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationParticipantWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
