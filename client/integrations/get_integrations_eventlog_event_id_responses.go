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

// GetIntegrationsEventlogEventIDReader is a Reader for the GetIntegrationsEventlogEventID structure.
type GetIntegrationsEventlogEventIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsEventlogEventIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsEventlogEventIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsEventlogEventIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsEventlogEventIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsEventlogEventIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsEventlogEventIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsEventlogEventIDRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsEventlogEventIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsEventlogEventIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsEventlogEventIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsEventlogEventIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsEventlogEventIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsEventlogEventIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsEventlogEventIDOK creates a GetIntegrationsEventlogEventIDOK with default headers values
func NewGetIntegrationsEventlogEventIDOK() *GetIntegrationsEventlogEventIDOK {
	return &GetIntegrationsEventlogEventIDOK{}
}

/*GetIntegrationsEventlogEventIDOK handles this case with default header values.

successful operation
*/
type GetIntegrationsEventlogEventIDOK struct {
	Payload *models.IntegrationEvent
}

func (o *GetIntegrationsEventlogEventIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDOK) GetPayload() *models.IntegrationEvent {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDBadRequest creates a GetIntegrationsEventlogEventIDBadRequest with default headers values
func NewGetIntegrationsEventlogEventIDBadRequest() *GetIntegrationsEventlogEventIDBadRequest {
	return &GetIntegrationsEventlogEventIDBadRequest{}
}

/*GetIntegrationsEventlogEventIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsEventlogEventIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDUnauthorized creates a GetIntegrationsEventlogEventIDUnauthorized with default headers values
func NewGetIntegrationsEventlogEventIDUnauthorized() *GetIntegrationsEventlogEventIDUnauthorized {
	return &GetIntegrationsEventlogEventIDUnauthorized{}
}

/*GetIntegrationsEventlogEventIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsEventlogEventIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDForbidden creates a GetIntegrationsEventlogEventIDForbidden with default headers values
func NewGetIntegrationsEventlogEventIDForbidden() *GetIntegrationsEventlogEventIDForbidden {
	return &GetIntegrationsEventlogEventIDForbidden{}
}

/*GetIntegrationsEventlogEventIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsEventlogEventIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDNotFound creates a GetIntegrationsEventlogEventIDNotFound with default headers values
func NewGetIntegrationsEventlogEventIDNotFound() *GetIntegrationsEventlogEventIDNotFound {
	return &GetIntegrationsEventlogEventIDNotFound{}
}

/*GetIntegrationsEventlogEventIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsEventlogEventIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDRequestTimeout creates a GetIntegrationsEventlogEventIDRequestTimeout with default headers values
func NewGetIntegrationsEventlogEventIDRequestTimeout() *GetIntegrationsEventlogEventIDRequestTimeout {
	return &GetIntegrationsEventlogEventIDRequestTimeout{}
}

/*GetIntegrationsEventlogEventIDRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsEventlogEventIDRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDRequestEntityTooLarge creates a GetIntegrationsEventlogEventIDRequestEntityTooLarge with default headers values
func NewGetIntegrationsEventlogEventIDRequestEntityTooLarge() *GetIntegrationsEventlogEventIDRequestEntityTooLarge {
	return &GetIntegrationsEventlogEventIDRequestEntityTooLarge{}
}

/*GetIntegrationsEventlogEventIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsEventlogEventIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDUnsupportedMediaType creates a GetIntegrationsEventlogEventIDUnsupportedMediaType with default headers values
func NewGetIntegrationsEventlogEventIDUnsupportedMediaType() *GetIntegrationsEventlogEventIDUnsupportedMediaType {
	return &GetIntegrationsEventlogEventIDUnsupportedMediaType{}
}

/*GetIntegrationsEventlogEventIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsEventlogEventIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDTooManyRequests creates a GetIntegrationsEventlogEventIDTooManyRequests with default headers values
func NewGetIntegrationsEventlogEventIDTooManyRequests() *GetIntegrationsEventlogEventIDTooManyRequests {
	return &GetIntegrationsEventlogEventIDTooManyRequests{}
}

/*GetIntegrationsEventlogEventIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsEventlogEventIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDInternalServerError creates a GetIntegrationsEventlogEventIDInternalServerError with default headers values
func NewGetIntegrationsEventlogEventIDInternalServerError() *GetIntegrationsEventlogEventIDInternalServerError {
	return &GetIntegrationsEventlogEventIDInternalServerError{}
}

/*GetIntegrationsEventlogEventIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsEventlogEventIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDServiceUnavailable creates a GetIntegrationsEventlogEventIDServiceUnavailable with default headers values
func NewGetIntegrationsEventlogEventIDServiceUnavailable() *GetIntegrationsEventlogEventIDServiceUnavailable {
	return &GetIntegrationsEventlogEventIDServiceUnavailable{}
}

/*GetIntegrationsEventlogEventIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsEventlogEventIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsEventlogEventIDGatewayTimeout creates a GetIntegrationsEventlogEventIDGatewayTimeout with default headers values
func NewGetIntegrationsEventlogEventIDGatewayTimeout() *GetIntegrationsEventlogEventIDGatewayTimeout {
	return &GetIntegrationsEventlogEventIDGatewayTimeout{}
}

/*GetIntegrationsEventlogEventIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsEventlogEventIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsEventlogEventIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/eventlog/{eventId}][%d] getIntegrationsEventlogEventIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsEventlogEventIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsEventlogEventIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
