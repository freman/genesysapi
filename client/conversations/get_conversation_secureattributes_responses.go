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

// GetConversationSecureattributesReader is a Reader for the GetConversationSecureattributes structure.
type GetConversationSecureattributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationSecureattributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationSecureattributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationSecureattributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationSecureattributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationSecureattributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationSecureattributesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationSecureattributesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationSecureattributesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationSecureattributesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationSecureattributesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationSecureattributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationSecureattributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationSecureattributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationSecureattributesOK creates a GetConversationSecureattributesOK with default headers values
func NewGetConversationSecureattributesOK() *GetConversationSecureattributesOK {
	return &GetConversationSecureattributesOK{}
}

/*GetConversationSecureattributesOK handles this case with default header values.

Secure attributes retrieved successfully
*/
type GetConversationSecureattributesOK struct {
	Payload *models.ConversationSecureAttributes
}

func (o *GetConversationSecureattributesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesOK  %+v", 200, o.Payload)
}

func (o *GetConversationSecureattributesOK) GetPayload() *models.ConversationSecureAttributes {
	return o.Payload
}

func (o *GetConversationSecureattributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConversationSecureAttributes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesBadRequest creates a GetConversationSecureattributesBadRequest with default headers values
func NewGetConversationSecureattributesBadRequest() *GetConversationSecureattributesBadRequest {
	return &GetConversationSecureattributesBadRequest{}
}

/*GetConversationSecureattributesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationSecureattributesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationSecureattributesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesUnauthorized creates a GetConversationSecureattributesUnauthorized with default headers values
func NewGetConversationSecureattributesUnauthorized() *GetConversationSecureattributesUnauthorized {
	return &GetConversationSecureattributesUnauthorized{}
}

/*GetConversationSecureattributesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationSecureattributesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationSecureattributesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesForbidden creates a GetConversationSecureattributesForbidden with default headers values
func NewGetConversationSecureattributesForbidden() *GetConversationSecureattributesForbidden {
	return &GetConversationSecureattributesForbidden{}
}

/*GetConversationSecureattributesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationSecureattributesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationSecureattributesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesNotFound creates a GetConversationSecureattributesNotFound with default headers values
func NewGetConversationSecureattributesNotFound() *GetConversationSecureattributesNotFound {
	return &GetConversationSecureattributesNotFound{}
}

/*GetConversationSecureattributesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationSecureattributesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationSecureattributesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesRequestTimeout creates a GetConversationSecureattributesRequestTimeout with default headers values
func NewGetConversationSecureattributesRequestTimeout() *GetConversationSecureattributesRequestTimeout {
	return &GetConversationSecureattributesRequestTimeout{}
}

/*GetConversationSecureattributesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationSecureattributesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationSecureattributesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesRequestEntityTooLarge creates a GetConversationSecureattributesRequestEntityTooLarge with default headers values
func NewGetConversationSecureattributesRequestEntityTooLarge() *GetConversationSecureattributesRequestEntityTooLarge {
	return &GetConversationSecureattributesRequestEntityTooLarge{}
}

/*GetConversationSecureattributesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationSecureattributesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationSecureattributesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesUnsupportedMediaType creates a GetConversationSecureattributesUnsupportedMediaType with default headers values
func NewGetConversationSecureattributesUnsupportedMediaType() *GetConversationSecureattributesUnsupportedMediaType {
	return &GetConversationSecureattributesUnsupportedMediaType{}
}

/*GetConversationSecureattributesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationSecureattributesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationSecureattributesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesTooManyRequests creates a GetConversationSecureattributesTooManyRequests with default headers values
func NewGetConversationSecureattributesTooManyRequests() *GetConversationSecureattributesTooManyRequests {
	return &GetConversationSecureattributesTooManyRequests{}
}

/*GetConversationSecureattributesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationSecureattributesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationSecureattributesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesInternalServerError creates a GetConversationSecureattributesInternalServerError with default headers values
func NewGetConversationSecureattributesInternalServerError() *GetConversationSecureattributesInternalServerError {
	return &GetConversationSecureattributesInternalServerError{}
}

/*GetConversationSecureattributesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationSecureattributesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationSecureattributesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesServiceUnavailable creates a GetConversationSecureattributesServiceUnavailable with default headers values
func NewGetConversationSecureattributesServiceUnavailable() *GetConversationSecureattributesServiceUnavailable {
	return &GetConversationSecureattributesServiceUnavailable{}
}

/*GetConversationSecureattributesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationSecureattributesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationSecureattributesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationSecureattributesGatewayTimeout creates a GetConversationSecureattributesGatewayTimeout with default headers values
func NewGetConversationSecureattributesGatewayTimeout() *GetConversationSecureattributesGatewayTimeout {
	return &GetConversationSecureattributesGatewayTimeout{}
}

/*GetConversationSecureattributesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationSecureattributesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationSecureattributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/secureattributes][%d] getConversationSecureattributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationSecureattributesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationSecureattributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}