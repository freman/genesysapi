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

// GetConversationReader is a Reader for the GetConversation structure.
type GetConversationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationOK creates a GetConversationOK with default headers values
func NewGetConversationOK() *GetConversationOK {
	return &GetConversationOK{}
}

/*GetConversationOK handles this case with default header values.

successful operation
*/
type GetConversationOK struct {
	Payload *models.Conversation
}

func (o *GetConversationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationOK  %+v", 200, o.Payload)
}

func (o *GetConversationOK) GetPayload() *models.Conversation {
	return o.Payload
}

func (o *GetConversationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Conversation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationBadRequest creates a GetConversationBadRequest with default headers values
func NewGetConversationBadRequest() *GetConversationBadRequest {
	return &GetConversationBadRequest{}
}

/*GetConversationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationUnauthorized creates a GetConversationUnauthorized with default headers values
func NewGetConversationUnauthorized() *GetConversationUnauthorized {
	return &GetConversationUnauthorized{}
}

/*GetConversationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationForbidden creates a GetConversationForbidden with default headers values
func NewGetConversationForbidden() *GetConversationForbidden {
	return &GetConversationForbidden{}
}

/*GetConversationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationNotFound creates a GetConversationNotFound with default headers values
func NewGetConversationNotFound() *GetConversationNotFound {
	return &GetConversationNotFound{}
}

/*GetConversationNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRequestTimeout creates a GetConversationRequestTimeout with default headers values
func NewGetConversationRequestTimeout() *GetConversationRequestTimeout {
	return &GetConversationRequestTimeout{}
}

/*GetConversationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRequestEntityTooLarge creates a GetConversationRequestEntityTooLarge with default headers values
func NewGetConversationRequestEntityTooLarge() *GetConversationRequestEntityTooLarge {
	return &GetConversationRequestEntityTooLarge{}
}

/*GetConversationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationUnsupportedMediaType creates a GetConversationUnsupportedMediaType with default headers values
func NewGetConversationUnsupportedMediaType() *GetConversationUnsupportedMediaType {
	return &GetConversationUnsupportedMediaType{}
}

/*GetConversationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationTooManyRequests creates a GetConversationTooManyRequests with default headers values
func NewGetConversationTooManyRequests() *GetConversationTooManyRequests {
	return &GetConversationTooManyRequests{}
}

/*GetConversationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationInternalServerError creates a GetConversationInternalServerError with default headers values
func NewGetConversationInternalServerError() *GetConversationInternalServerError {
	return &GetConversationInternalServerError{}
}

/*GetConversationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationServiceUnavailable creates a GetConversationServiceUnavailable with default headers values
func NewGetConversationServiceUnavailable() *GetConversationServiceUnavailable {
	return &GetConversationServiceUnavailable{}
}

/*GetConversationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationGatewayTimeout creates a GetConversationGatewayTimeout with default headers values
func NewGetConversationGatewayTimeout() *GetConversationGatewayTimeout {
	return &GetConversationGatewayTimeout{}
}

/*GetConversationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}][%d] getConversationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
