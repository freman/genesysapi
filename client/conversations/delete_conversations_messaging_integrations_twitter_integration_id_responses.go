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

// DeleteConversationsMessagingIntegrationsTwitterIntegrationIDReader is a Reader for the DeleteConversationsMessagingIntegrationsTwitterIntegrationID structure.
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent handles this case with default header values.

Operation was successful
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent struct {
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdNoContent ", 204)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout creates a DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout() *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout {
	return &DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout{}
}

/*DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/twitter/{integrationId}][%d] deleteConversationsMessagingIntegrationsTwitterIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsTwitterIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
