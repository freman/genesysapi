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

// GetConversationsMessagingIntegrationsWhatsappIntegrationIDReader is a Reader for the GetConversationsMessagingIntegrationsWhatsappIntegrationID structure.
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDOK creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDOK() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK struct {
	Payload *models.WhatsAppIntegration
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK) GetPayload() *models.WhatsAppIntegration {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhatsAppIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout creates a GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout() *GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout {
	return &GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout{}
}

/*GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] getConversationsMessagingIntegrationsWhatsappIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
