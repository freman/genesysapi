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

// DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDReader is a Reader for the DeleteConversationsMessagingIntegrationsWhatsappIntegrationID structure.
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK handles this case with default header values.

successful operation
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK struct {
	Payload *models.WhatsAppIntegration
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdOK  %+v", 200, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK) GetPayload() *models.WhatsAppIntegration {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhatsAppIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted handles this case with default header values.

Request Accepted
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted struct {
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdAccepted ", 202)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout creates a DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout() *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout {
	return &DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout{}
}

/*DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] deleteConversationsMessagingIntegrationsWhatsappIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
