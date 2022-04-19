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

// GetConversationsMessagingIntegrationsOpenIntegrationIDReader is a Reader for the GetConversationsMessagingIntegrationsOpenIntegrationID structure.
type GetConversationsMessagingIntegrationsOpenIntegrationIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDOK creates a GetConversationsMessagingIntegrationsOpenIntegrationIDOK with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDOK() *GetConversationsMessagingIntegrationsOpenIntegrationIDOK {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDOK{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDOK struct {
	Payload *models.OpenIntegration
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDOK) GetPayload() *models.OpenIntegration {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenIntegration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest creates a GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest() *GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized creates a GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized() *GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDForbidden creates a GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDForbidden() *GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDNotFound creates a GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDNotFound() *GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout creates a GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout() *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge() *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType creates a GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType() *GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests creates a GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests() *GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError creates a GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError() *GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable creates a GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable() *GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout creates a GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout() *GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout {
	return &GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout{}
}

/*GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/open/{integrationId}][%d] getConversationsMessagingIntegrationsOpenIntegrationIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsOpenIntegrationIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
