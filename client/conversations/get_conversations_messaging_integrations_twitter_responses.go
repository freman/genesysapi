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

// GetConversationsMessagingIntegrationsTwitterReader is a Reader for the GetConversationsMessagingIntegrationsTwitter structure.
type GetConversationsMessagingIntegrationsTwitterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsTwitterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsTwitterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsTwitterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsTwitterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsTwitterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsTwitterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsTwitterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsTwitterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsTwitterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsTwitterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsTwitterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsTwitterOK creates a GetConversationsMessagingIntegrationsTwitterOK with default headers values
func NewGetConversationsMessagingIntegrationsTwitterOK() *GetConversationsMessagingIntegrationsTwitterOK {
	return &GetConversationsMessagingIntegrationsTwitterOK{}
}

/*GetConversationsMessagingIntegrationsTwitterOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsTwitterOK struct {
	Payload *models.TwitterIntegrationEntityListing
}

func (o *GetConversationsMessagingIntegrationsTwitterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterOK) GetPayload() *models.TwitterIntegrationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TwitterIntegrationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterBadRequest creates a GetConversationsMessagingIntegrationsTwitterBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsTwitterBadRequest() *GetConversationsMessagingIntegrationsTwitterBadRequest {
	return &GetConversationsMessagingIntegrationsTwitterBadRequest{}
}

/*GetConversationsMessagingIntegrationsTwitterBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsTwitterBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterUnauthorized creates a GetConversationsMessagingIntegrationsTwitterUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsTwitterUnauthorized() *GetConversationsMessagingIntegrationsTwitterUnauthorized {
	return &GetConversationsMessagingIntegrationsTwitterUnauthorized{}
}

/*GetConversationsMessagingIntegrationsTwitterUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsTwitterUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterForbidden creates a GetConversationsMessagingIntegrationsTwitterForbidden with default headers values
func NewGetConversationsMessagingIntegrationsTwitterForbidden() *GetConversationsMessagingIntegrationsTwitterForbidden {
	return &GetConversationsMessagingIntegrationsTwitterForbidden{}
}

/*GetConversationsMessagingIntegrationsTwitterForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsTwitterForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterNotFound creates a GetConversationsMessagingIntegrationsTwitterNotFound with default headers values
func NewGetConversationsMessagingIntegrationsTwitterNotFound() *GetConversationsMessagingIntegrationsTwitterNotFound {
	return &GetConversationsMessagingIntegrationsTwitterNotFound{}
}

/*GetConversationsMessagingIntegrationsTwitterNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsTwitterNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge() *GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge{}
}

/*GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterUnsupportedMediaType creates a GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsTwitterUnsupportedMediaType() *GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType{}
}

/*GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterTooManyRequests creates a GetConversationsMessagingIntegrationsTwitterTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsTwitterTooManyRequests() *GetConversationsMessagingIntegrationsTwitterTooManyRequests {
	return &GetConversationsMessagingIntegrationsTwitterTooManyRequests{}
}

/*GetConversationsMessagingIntegrationsTwitterTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsMessagingIntegrationsTwitterTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterInternalServerError creates a GetConversationsMessagingIntegrationsTwitterInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsTwitterInternalServerError() *GetConversationsMessagingIntegrationsTwitterInternalServerError {
	return &GetConversationsMessagingIntegrationsTwitterInternalServerError{}
}

/*GetConversationsMessagingIntegrationsTwitterInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsTwitterInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterServiceUnavailable creates a GetConversationsMessagingIntegrationsTwitterServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsTwitterServiceUnavailable() *GetConversationsMessagingIntegrationsTwitterServiceUnavailable {
	return &GetConversationsMessagingIntegrationsTwitterServiceUnavailable{}
}

/*GetConversationsMessagingIntegrationsTwitterServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsTwitterServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsTwitterGatewayTimeout creates a GetConversationsMessagingIntegrationsTwitterGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsTwitterGatewayTimeout() *GetConversationsMessagingIntegrationsTwitterGatewayTimeout {
	return &GetConversationsMessagingIntegrationsTwitterGatewayTimeout{}
}

/*GetConversationsMessagingIntegrationsTwitterGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsTwitterGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsTwitterGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/twitter][%d] getConversationsMessagingIntegrationsTwitterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsTwitterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsTwitterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
