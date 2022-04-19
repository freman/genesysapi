// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOutboundConversationDncReader is a Reader for the PostOutboundConversationDnc structure.
type PostOutboundConversationDncReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOutboundConversationDncReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 400:
		result := NewPostOutboundConversationDncBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOutboundConversationDncUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOutboundConversationDncForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOutboundConversationDncNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOutboundConversationDncRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOutboundConversationDncRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOutboundConversationDncUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOutboundConversationDncTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOutboundConversationDncInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOutboundConversationDncServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOutboundConversationDncGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostOutboundConversationDncDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostOutboundConversationDncBadRequest creates a PostOutboundConversationDncBadRequest with default headers values
func NewPostOutboundConversationDncBadRequest() *PostOutboundConversationDncBadRequest {
	return &PostOutboundConversationDncBadRequest{}
}

/*PostOutboundConversationDncBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOutboundConversationDncBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncBadRequest  %+v", 400, o.Payload)
}

func (o *PostOutboundConversationDncBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncUnauthorized creates a PostOutboundConversationDncUnauthorized with default headers values
func NewPostOutboundConversationDncUnauthorized() *PostOutboundConversationDncUnauthorized {
	return &PostOutboundConversationDncUnauthorized{}
}

/*PostOutboundConversationDncUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOutboundConversationDncUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOutboundConversationDncUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncForbidden creates a PostOutboundConversationDncForbidden with default headers values
func NewPostOutboundConversationDncForbidden() *PostOutboundConversationDncForbidden {
	return &PostOutboundConversationDncForbidden{}
}

/*PostOutboundConversationDncForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostOutboundConversationDncForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncForbidden  %+v", 403, o.Payload)
}

func (o *PostOutboundConversationDncForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncNotFound creates a PostOutboundConversationDncNotFound with default headers values
func NewPostOutboundConversationDncNotFound() *PostOutboundConversationDncNotFound {
	return &PostOutboundConversationDncNotFound{}
}

/*PostOutboundConversationDncNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostOutboundConversationDncNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncNotFound  %+v", 404, o.Payload)
}

func (o *PostOutboundConversationDncNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncRequestTimeout creates a PostOutboundConversationDncRequestTimeout with default headers values
func NewPostOutboundConversationDncRequestTimeout() *PostOutboundConversationDncRequestTimeout {
	return &PostOutboundConversationDncRequestTimeout{}
}

/*PostOutboundConversationDncRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOutboundConversationDncRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOutboundConversationDncRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncRequestEntityTooLarge creates a PostOutboundConversationDncRequestEntityTooLarge with default headers values
func NewPostOutboundConversationDncRequestEntityTooLarge() *PostOutboundConversationDncRequestEntityTooLarge {
	return &PostOutboundConversationDncRequestEntityTooLarge{}
}

/*PostOutboundConversationDncRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostOutboundConversationDncRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOutboundConversationDncRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncUnsupportedMediaType creates a PostOutboundConversationDncUnsupportedMediaType with default headers values
func NewPostOutboundConversationDncUnsupportedMediaType() *PostOutboundConversationDncUnsupportedMediaType {
	return &PostOutboundConversationDncUnsupportedMediaType{}
}

/*PostOutboundConversationDncUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOutboundConversationDncUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOutboundConversationDncUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncTooManyRequests creates a PostOutboundConversationDncTooManyRequests with default headers values
func NewPostOutboundConversationDncTooManyRequests() *PostOutboundConversationDncTooManyRequests {
	return &PostOutboundConversationDncTooManyRequests{}
}

/*PostOutboundConversationDncTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOutboundConversationDncTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOutboundConversationDncTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncInternalServerError creates a PostOutboundConversationDncInternalServerError with default headers values
func NewPostOutboundConversationDncInternalServerError() *PostOutboundConversationDncInternalServerError {
	return &PostOutboundConversationDncInternalServerError{}
}

/*PostOutboundConversationDncInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOutboundConversationDncInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOutboundConversationDncInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncServiceUnavailable creates a PostOutboundConversationDncServiceUnavailable with default headers values
func NewPostOutboundConversationDncServiceUnavailable() *PostOutboundConversationDncServiceUnavailable {
	return &PostOutboundConversationDncServiceUnavailable{}
}

/*PostOutboundConversationDncServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOutboundConversationDncServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOutboundConversationDncServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncGatewayTimeout creates a PostOutboundConversationDncGatewayTimeout with default headers values
func NewPostOutboundConversationDncGatewayTimeout() *PostOutboundConversationDncGatewayTimeout {
	return &PostOutboundConversationDncGatewayTimeout{}
}

/*PostOutboundConversationDncGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostOutboundConversationDncGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostOutboundConversationDncGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDncGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOutboundConversationDncGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOutboundConversationDncGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOutboundConversationDncDefault creates a PostOutboundConversationDncDefault with default headers values
func NewPostOutboundConversationDncDefault(code int) *PostOutboundConversationDncDefault {
	return &PostOutboundConversationDncDefault{
		_statusCode: code,
	}
}

/*PostOutboundConversationDncDefault handles this case with default header values.

successful operation
*/
type PostOutboundConversationDncDefault struct {
	_statusCode int
}

// Code gets the status code for the post outbound conversation dnc default response
func (o *PostOutboundConversationDncDefault) Code() int {
	return o._statusCode
}

func (o *PostOutboundConversationDncDefault) Error() string {
	return fmt.Sprintf("[POST /api/v2/outbound/conversations/{conversationId}/dnc][%d] postOutboundConversationDnc default ", o._statusCode)
}

func (o *PostOutboundConversationDncDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
