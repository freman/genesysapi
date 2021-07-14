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

// GetConversationsCallbackParticipantWrapupcodesReader is a Reader for the GetConversationsCallbackParticipantWrapupcodes structure.
type GetConversationsCallbackParticipantWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsCallbackParticipantWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsCallbackParticipantWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsCallbackParticipantWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsCallbackParticipantWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsCallbackParticipantWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsCallbackParticipantWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsCallbackParticipantWrapupcodesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsCallbackParticipantWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsCallbackParticipantWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsCallbackParticipantWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsCallbackParticipantWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsCallbackParticipantWrapupcodesOK creates a GetConversationsCallbackParticipantWrapupcodesOK with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesOK() *GetConversationsCallbackParticipantWrapupcodesOK {
	return &GetConversationsCallbackParticipantWrapupcodesOK{}
}

/*GetConversationsCallbackParticipantWrapupcodesOK handles this case with default header values.

successful operation
*/
type GetConversationsCallbackParticipantWrapupcodesOK struct {
	Payload []*models.WrapupCode
}

func (o *GetConversationsCallbackParticipantWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesOK) GetPayload() []*models.WrapupCode {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesBadRequest creates a GetConversationsCallbackParticipantWrapupcodesBadRequest with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesBadRequest() *GetConversationsCallbackParticipantWrapupcodesBadRequest {
	return &GetConversationsCallbackParticipantWrapupcodesBadRequest{}
}

/*GetConversationsCallbackParticipantWrapupcodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsCallbackParticipantWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesUnauthorized creates a GetConversationsCallbackParticipantWrapupcodesUnauthorized with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesUnauthorized() *GetConversationsCallbackParticipantWrapupcodesUnauthorized {
	return &GetConversationsCallbackParticipantWrapupcodesUnauthorized{}
}

/*GetConversationsCallbackParticipantWrapupcodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsCallbackParticipantWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesForbidden creates a GetConversationsCallbackParticipantWrapupcodesForbidden with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesForbidden() *GetConversationsCallbackParticipantWrapupcodesForbidden {
	return &GetConversationsCallbackParticipantWrapupcodesForbidden{}
}

/*GetConversationsCallbackParticipantWrapupcodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsCallbackParticipantWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesNotFound creates a GetConversationsCallbackParticipantWrapupcodesNotFound with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesNotFound() *GetConversationsCallbackParticipantWrapupcodesNotFound {
	return &GetConversationsCallbackParticipantWrapupcodesNotFound{}
}

/*GetConversationsCallbackParticipantWrapupcodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsCallbackParticipantWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesRequestTimeout creates a GetConversationsCallbackParticipantWrapupcodesRequestTimeout with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesRequestTimeout() *GetConversationsCallbackParticipantWrapupcodesRequestTimeout {
	return &GetConversationsCallbackParticipantWrapupcodesRequestTimeout{}
}

/*GetConversationsCallbackParticipantWrapupcodesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsCallbackParticipantWrapupcodesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge creates a GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge() *GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge {
	return &GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge{}
}

/*GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType creates a GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType() *GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType {
	return &GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType{}
}

/*GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesTooManyRequests creates a GetConversationsCallbackParticipantWrapupcodesTooManyRequests with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesTooManyRequests() *GetConversationsCallbackParticipantWrapupcodesTooManyRequests {
	return &GetConversationsCallbackParticipantWrapupcodesTooManyRequests{}
}

/*GetConversationsCallbackParticipantWrapupcodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsCallbackParticipantWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesInternalServerError creates a GetConversationsCallbackParticipantWrapupcodesInternalServerError with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesInternalServerError() *GetConversationsCallbackParticipantWrapupcodesInternalServerError {
	return &GetConversationsCallbackParticipantWrapupcodesInternalServerError{}
}

/*GetConversationsCallbackParticipantWrapupcodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsCallbackParticipantWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesServiceUnavailable creates a GetConversationsCallbackParticipantWrapupcodesServiceUnavailable with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesServiceUnavailable() *GetConversationsCallbackParticipantWrapupcodesServiceUnavailable {
	return &GetConversationsCallbackParticipantWrapupcodesServiceUnavailable{}
}

/*GetConversationsCallbackParticipantWrapupcodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsCallbackParticipantWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsCallbackParticipantWrapupcodesGatewayTimeout creates a GetConversationsCallbackParticipantWrapupcodesGatewayTimeout with default headers values
func NewGetConversationsCallbackParticipantWrapupcodesGatewayTimeout() *GetConversationsCallbackParticipantWrapupcodesGatewayTimeout {
	return &GetConversationsCallbackParticipantWrapupcodesGatewayTimeout{}
}

/*GetConversationsCallbackParticipantWrapupcodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsCallbackParticipantWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsCallbackParticipantWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/callbacks/{conversationId}/participants/{participantId}/wrapupcodes][%d] getConversationsCallbackParticipantWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsCallbackParticipantWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsCallbackParticipantWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
