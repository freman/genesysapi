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

// PatchConversationsMessagingIntegrationsWhatsappIntegrationIDReader is a Reader for the PatchConversationsMessagingIntegrationsWhatsappIntegrationID structure.
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK handles this case with default header values.

successful operation
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK struct {
	Payload *models.WhatsAppIntegration
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdOK  %+v", 200, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK) GetPayload() *models.WhatsAppIntegration {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhatsAppIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted handles this case with default header values.

Processing Request - If request was to Activate, do a GET checking for activationStatus set to CODE_SENT.
If request was to Confirm, do a GET checking for the integration status set to Active
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted struct {
	Payload *models.WhatsAppIntegration
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdAccepted  %+v", 202, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted) GetPayload() *models.WhatsAppIntegration {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WhatsAppIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout creates a PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout with default headers values
func NewPatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout() *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout {
	return &PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout{}
}

/*PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/conversations/messaging/integrations/whatsapp/{integrationId}][%d] patchConversationsMessagingIntegrationsWhatsappIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchConversationsMessagingIntegrationsWhatsappIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
