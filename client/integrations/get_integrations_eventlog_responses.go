// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIntegrationsEventlogReader is a Reader for the GetIntegrationsEventlog structure.
type GetIntegrationsEventlogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsEventlogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsEventlogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsEventlogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsEventlogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsEventlogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsEventlogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsEventlogRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsEventlogUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsEventlogTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsEventlogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsEventlogServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsEventlogGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsEventlogOK creates a GetIntegrationsEventlogOK with default headers values
func NewGetIntegrationsEventlogOK() *GetIntegrationsEventlogOK {
	return &GetIntegrationsEventlogOK{}
}

/*GetIntegrationsEventlogOK handles this case with default header values.

successful operation
*/
type GetIntegrationsEventlogOK struct {
	Payload *models.IntegrationEventEntityListing
}

func (o *GetIntegrationsEventlogOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsEventlogOK) GetPayload() *models.IntegrationEventEntityListing {
	return o.Payload
}

func (o *GetIntegrationsEventlogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationEventEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogBadRequest creates a GetIntegrationsEventlogBadRequest with default headers values
func NewGetIntegrationsEventlogBadRequest() *GetIntegrationsEventlogBadRequest {
	return &GetIntegrationsEventlogBadRequest{}
}

/*GetIntegrationsEventlogBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsEventlogBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsEventlogBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogUnauthorized creates a GetIntegrationsEventlogUnauthorized with default headers values
func NewGetIntegrationsEventlogUnauthorized() *GetIntegrationsEventlogUnauthorized {
	return &GetIntegrationsEventlogUnauthorized{}
}

/*GetIntegrationsEventlogUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsEventlogUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsEventlogUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogForbidden creates a GetIntegrationsEventlogForbidden with default headers values
func NewGetIntegrationsEventlogForbidden() *GetIntegrationsEventlogForbidden {
	return &GetIntegrationsEventlogForbidden{}
}

/*GetIntegrationsEventlogForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsEventlogForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsEventlogForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogNotFound creates a GetIntegrationsEventlogNotFound with default headers values
func NewGetIntegrationsEventlogNotFound() *GetIntegrationsEventlogNotFound {
	return &GetIntegrationsEventlogNotFound{}
}

/*GetIntegrationsEventlogNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsEventlogNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsEventlogNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogRequestEntityTooLarge creates a GetIntegrationsEventlogRequestEntityTooLarge with default headers values
func NewGetIntegrationsEventlogRequestEntityTooLarge() *GetIntegrationsEventlogRequestEntityTooLarge {
	return &GetIntegrationsEventlogRequestEntityTooLarge{}
}

/*GetIntegrationsEventlogRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsEventlogRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsEventlogRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogUnsupportedMediaType creates a GetIntegrationsEventlogUnsupportedMediaType with default headers values
func NewGetIntegrationsEventlogUnsupportedMediaType() *GetIntegrationsEventlogUnsupportedMediaType {
	return &GetIntegrationsEventlogUnsupportedMediaType{}
}

/*GetIntegrationsEventlogUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsEventlogUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsEventlogUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogTooManyRequests creates a GetIntegrationsEventlogTooManyRequests with default headers values
func NewGetIntegrationsEventlogTooManyRequests() *GetIntegrationsEventlogTooManyRequests {
	return &GetIntegrationsEventlogTooManyRequests{}
}

/*GetIntegrationsEventlogTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsEventlogTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsEventlogTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogInternalServerError creates a GetIntegrationsEventlogInternalServerError with default headers values
func NewGetIntegrationsEventlogInternalServerError() *GetIntegrationsEventlogInternalServerError {
	return &GetIntegrationsEventlogInternalServerError{}
}

/*GetIntegrationsEventlogInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsEventlogInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsEventlogInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogServiceUnavailable creates a GetIntegrationsEventlogServiceUnavailable with default headers values
func NewGetIntegrationsEventlogServiceUnavailable() *GetIntegrationsEventlogServiceUnavailable {
	return &GetIntegrationsEventlogServiceUnavailable{}
}

/*GetIntegrationsEventlogServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsEventlogServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsEventlogServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogGatewayTimeout creates a GetIntegrationsEventlogGatewayTimeout with default headers values
func NewGetIntegrationsEventlogGatewayTimeout() *GetIntegrationsEventlogGatewayTimeout {
	return &GetIntegrationsEventlogGatewayTimeout{}
}

/*GetIntegrationsEventlogGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsEventlogGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog][%d] getIntegrationsEventlogGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsEventlogGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}