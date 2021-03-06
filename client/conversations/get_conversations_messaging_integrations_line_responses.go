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

// GetConversationsMessagingIntegrationsLineReader is a Reader for the GetConversationsMessagingIntegrationsLine structure.
type GetConversationsMessagingIntegrationsLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationsMessagingIntegrationsLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationsMessagingIntegrationsLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationsMessagingIntegrationsLineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationsMessagingIntegrationsLineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationsMessagingIntegrationsLineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationsMessagingIntegrationsLineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationsMessagingIntegrationsLineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationsMessagingIntegrationsLineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationsMessagingIntegrationsLineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationsMessagingIntegrationsLineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationsMessagingIntegrationsLineOK creates a GetConversationsMessagingIntegrationsLineOK with default headers values
func NewGetConversationsMessagingIntegrationsLineOK() *GetConversationsMessagingIntegrationsLineOK {
	return &GetConversationsMessagingIntegrationsLineOK{}
}

/*GetConversationsMessagingIntegrationsLineOK handles this case with default header values.

successful operation
*/
type GetConversationsMessagingIntegrationsLineOK struct {
	Payload *models.LineIntegrationEntityListing
}

func (o *GetConversationsMessagingIntegrationsLineOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineOK  %+v", 200, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineOK) GetPayload() *models.LineIntegrationEntityListing {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LineIntegrationEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineBadRequest creates a GetConversationsMessagingIntegrationsLineBadRequest with default headers values
func NewGetConversationsMessagingIntegrationsLineBadRequest() *GetConversationsMessagingIntegrationsLineBadRequest {
	return &GetConversationsMessagingIntegrationsLineBadRequest{}
}

/*GetConversationsMessagingIntegrationsLineBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationsMessagingIntegrationsLineBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineUnauthorized creates a GetConversationsMessagingIntegrationsLineUnauthorized with default headers values
func NewGetConversationsMessagingIntegrationsLineUnauthorized() *GetConversationsMessagingIntegrationsLineUnauthorized {
	return &GetConversationsMessagingIntegrationsLineUnauthorized{}
}

/*GetConversationsMessagingIntegrationsLineUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationsMessagingIntegrationsLineUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineForbidden creates a GetConversationsMessagingIntegrationsLineForbidden with default headers values
func NewGetConversationsMessagingIntegrationsLineForbidden() *GetConversationsMessagingIntegrationsLineForbidden {
	return &GetConversationsMessagingIntegrationsLineForbidden{}
}

/*GetConversationsMessagingIntegrationsLineForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationsMessagingIntegrationsLineForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineNotFound creates a GetConversationsMessagingIntegrationsLineNotFound with default headers values
func NewGetConversationsMessagingIntegrationsLineNotFound() *GetConversationsMessagingIntegrationsLineNotFound {
	return &GetConversationsMessagingIntegrationsLineNotFound{}
}

/*GetConversationsMessagingIntegrationsLineNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationsMessagingIntegrationsLineNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge creates a GetConversationsMessagingIntegrationsLineRequestEntityTooLarge with default headers values
func NewGetConversationsMessagingIntegrationsLineRequestEntityTooLarge() *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge {
	return &GetConversationsMessagingIntegrationsLineRequestEntityTooLarge{}
}

/*GetConversationsMessagingIntegrationsLineRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationsMessagingIntegrationsLineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType creates a GetConversationsMessagingIntegrationsLineUnsupportedMediaType with default headers values
func NewGetConversationsMessagingIntegrationsLineUnsupportedMediaType() *GetConversationsMessagingIntegrationsLineUnsupportedMediaType {
	return &GetConversationsMessagingIntegrationsLineUnsupportedMediaType{}
}

/*GetConversationsMessagingIntegrationsLineUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationsMessagingIntegrationsLineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineTooManyRequests creates a GetConversationsMessagingIntegrationsLineTooManyRequests with default headers values
func NewGetConversationsMessagingIntegrationsLineTooManyRequests() *GetConversationsMessagingIntegrationsLineTooManyRequests {
	return &GetConversationsMessagingIntegrationsLineTooManyRequests{}
}

/*GetConversationsMessagingIntegrationsLineTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationsMessagingIntegrationsLineTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineInternalServerError creates a GetConversationsMessagingIntegrationsLineInternalServerError with default headers values
func NewGetConversationsMessagingIntegrationsLineInternalServerError() *GetConversationsMessagingIntegrationsLineInternalServerError {
	return &GetConversationsMessagingIntegrationsLineInternalServerError{}
}

/*GetConversationsMessagingIntegrationsLineInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationsMessagingIntegrationsLineInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineServiceUnavailable creates a GetConversationsMessagingIntegrationsLineServiceUnavailable with default headers values
func NewGetConversationsMessagingIntegrationsLineServiceUnavailable() *GetConversationsMessagingIntegrationsLineServiceUnavailable {
	return &GetConversationsMessagingIntegrationsLineServiceUnavailable{}
}

/*GetConversationsMessagingIntegrationsLineServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationsMessagingIntegrationsLineServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationsMessagingIntegrationsLineGatewayTimeout creates a GetConversationsMessagingIntegrationsLineGatewayTimeout with default headers values
func NewGetConversationsMessagingIntegrationsLineGatewayTimeout() *GetConversationsMessagingIntegrationsLineGatewayTimeout {
	return &GetConversationsMessagingIntegrationsLineGatewayTimeout{}
}

/*GetConversationsMessagingIntegrationsLineGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationsMessagingIntegrationsLineGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/messaging/integrations/line][%d] getConversationsMessagingIntegrationsLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationsMessagingIntegrationsLineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
