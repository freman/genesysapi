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

// DeleteConversationsMessagingIntegrationsFacebookIntegrationIDReader is a Reader for the DeleteConversationsMessagingIntegrationsFacebookIntegrationID structure.
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent handles this case with default header values.

Operation was successful
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent struct {
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdNoContent ", 204)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout creates a DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout with default headers values
func NewDeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout() *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout {
	return &DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout{}
}

/*DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/conversations/messaging/integrations/facebook/{integrationId}][%d] deleteConversationsMessagingIntegrationsFacebookIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteConversationsMessagingIntegrationsFacebookIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
