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

// PostConversationsMessagingIntegrationsWhatsappReader is a Reader for the PostConversationsMessagingIntegrationsWhatsapp structure.
type PostConversationsMessagingIntegrationsWhatsappReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostConversationsMessagingIntegrationsWhatsappReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostConversationsMessagingIntegrationsWhatsappOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostConversationsMessagingIntegrationsWhatsappBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostConversationsMessagingIntegrationsWhatsappUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostConversationsMessagingIntegrationsWhatsappForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostConversationsMessagingIntegrationsWhatsappNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostConversationsMessagingIntegrationsWhatsappTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostConversationsMessagingIntegrationsWhatsappInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostConversationsMessagingIntegrationsWhatsappServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostConversationsMessagingIntegrationsWhatsappGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostConversationsMessagingIntegrationsWhatsappOK creates a PostConversationsMessagingIntegrationsWhatsappOK with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappOK() *PostConversationsMessagingIntegrationsWhatsappOK {
	return &PostConversationsMessagingIntegrationsWhatsappOK{}
}

/*PostConversationsMessagingIntegrationsWhatsappOK handles this case with default header values.

successful operation
*/
type PostConversationsMessagingIntegrationsWhatsappOK struct {
	Payload *models.WhatsAppIntegration
}

func (o *PostConversationsMessagingIntegrationsWhatsappOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappOK  %+v", 200, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappOK) GetPayload() *models.WhatsAppIntegration {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhatsAppIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappBadRequest creates a PostConversationsMessagingIntegrationsWhatsappBadRequest with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappBadRequest() *PostConversationsMessagingIntegrationsWhatsappBadRequest {
	return &PostConversationsMessagingIntegrationsWhatsappBadRequest{}
}

/*PostConversationsMessagingIntegrationsWhatsappBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostConversationsMessagingIntegrationsWhatsappBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappBadRequest  %+v", 400, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappUnauthorized creates a PostConversationsMessagingIntegrationsWhatsappUnauthorized with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappUnauthorized() *PostConversationsMessagingIntegrationsWhatsappUnauthorized {
	return &PostConversationsMessagingIntegrationsWhatsappUnauthorized{}
}

/*PostConversationsMessagingIntegrationsWhatsappUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostConversationsMessagingIntegrationsWhatsappUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappUnauthorized  %+v", 401, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappForbidden creates a PostConversationsMessagingIntegrationsWhatsappForbidden with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappForbidden() *PostConversationsMessagingIntegrationsWhatsappForbidden {
	return &PostConversationsMessagingIntegrationsWhatsappForbidden{}
}

/*PostConversationsMessagingIntegrationsWhatsappForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostConversationsMessagingIntegrationsWhatsappForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappForbidden  %+v", 403, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappNotFound creates a PostConversationsMessagingIntegrationsWhatsappNotFound with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappNotFound() *PostConversationsMessagingIntegrationsWhatsappNotFound {
	return &PostConversationsMessagingIntegrationsWhatsappNotFound{}
}

/*PostConversationsMessagingIntegrationsWhatsappNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostConversationsMessagingIntegrationsWhatsappNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappNotFound  %+v", 404, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge creates a PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge() *PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge {
	return &PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge{}
}

/*PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType creates a PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType() *PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType {
	return &PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType{}
}

/*PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappTooManyRequests creates a PostConversationsMessagingIntegrationsWhatsappTooManyRequests with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappTooManyRequests() *PostConversationsMessagingIntegrationsWhatsappTooManyRequests {
	return &PostConversationsMessagingIntegrationsWhatsappTooManyRequests{}
}

/*PostConversationsMessagingIntegrationsWhatsappTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostConversationsMessagingIntegrationsWhatsappTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappInternalServerError creates a PostConversationsMessagingIntegrationsWhatsappInternalServerError with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappInternalServerError() *PostConversationsMessagingIntegrationsWhatsappInternalServerError {
	return &PostConversationsMessagingIntegrationsWhatsappInternalServerError{}
}

/*PostConversationsMessagingIntegrationsWhatsappInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostConversationsMessagingIntegrationsWhatsappInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappInternalServerError  %+v", 500, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappServiceUnavailable creates a PostConversationsMessagingIntegrationsWhatsappServiceUnavailable with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappServiceUnavailable() *PostConversationsMessagingIntegrationsWhatsappServiceUnavailable {
	return &PostConversationsMessagingIntegrationsWhatsappServiceUnavailable{}
}

/*PostConversationsMessagingIntegrationsWhatsappServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostConversationsMessagingIntegrationsWhatsappServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostConversationsMessagingIntegrationsWhatsappGatewayTimeout creates a PostConversationsMessagingIntegrationsWhatsappGatewayTimeout with default headers values
func NewPostConversationsMessagingIntegrationsWhatsappGatewayTimeout() *PostConversationsMessagingIntegrationsWhatsappGatewayTimeout {
	return &PostConversationsMessagingIntegrationsWhatsappGatewayTimeout{}
}

/*PostConversationsMessagingIntegrationsWhatsappGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostConversationsMessagingIntegrationsWhatsappGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostConversationsMessagingIntegrationsWhatsappGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/conversations/messaging/integrations/whatsapp][%d] postConversationsMessagingIntegrationsWhatsappGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostConversationsMessagingIntegrationsWhatsappGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostConversationsMessagingIntegrationsWhatsappGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
