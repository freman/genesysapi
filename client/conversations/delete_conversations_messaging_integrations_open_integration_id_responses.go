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

// DeleteConversationsMessagingIntegrationsOpenIntegrationIDReader is a Reader for the DeleteConversationsMessagingIntegrationsOpenIntegrationID structure.
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent handles this case with default header values.

Operation was successful
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent struct {
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdNoContent ", 204)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout creates a DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout() *DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout {
	return &DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout{}
}

/*DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] deleteConversationsMessagingIntegrationsOpenIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}