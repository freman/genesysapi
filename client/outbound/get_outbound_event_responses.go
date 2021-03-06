// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundEventReader is a Reader for the GetOutboundEvent structure.
type GetOutboundEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundEventBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundEventUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundEventForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundEventNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundEventRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundEventUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundEventTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundEventInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundEventServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundEventGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundEventOK creates a GetOutboundEventOK with default headers values
func NewGetOutboundEventOK() *GetOutboundEventOK {
	return &GetOutboundEventOK{}
}

/*GetOutboundEventOK handles this case with default header values.

successful operation
*/
type GetOutboundEventOK struct {
	Payload *models.EventLog
}

func (o *GetOutboundEventOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventOK  %+v", 200, o.Payload)
}

func (o *GetOutboundEventOK) GetPayload() *models.EventLog {
	return o.Payload
}

func (o *GetOutboundEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EventLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventBadRequest creates a GetOutboundEventBadRequest with default headers values
func NewGetOutboundEventBadRequest() *GetOutboundEventBadRequest {
	return &GetOutboundEventBadRequest{}
}

/*GetOutboundEventBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundEventBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundEventBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventUnauthorized creates a GetOutboundEventUnauthorized with default headers values
func NewGetOutboundEventUnauthorized() *GetOutboundEventUnauthorized {
	return &GetOutboundEventUnauthorized{}
}

/*GetOutboundEventUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundEventUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundEventUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventForbidden creates a GetOutboundEventForbidden with default headers values
func NewGetOutboundEventForbidden() *GetOutboundEventForbidden {
	return &GetOutboundEventForbidden{}
}

/*GetOutboundEventForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundEventForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundEventForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventNotFound creates a GetOutboundEventNotFound with default headers values
func NewGetOutboundEventNotFound() *GetOutboundEventNotFound {
	return &GetOutboundEventNotFound{}
}

/*GetOutboundEventNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundEventNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundEventNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventRequestEntityTooLarge creates a GetOutboundEventRequestEntityTooLarge with default headers values
func NewGetOutboundEventRequestEntityTooLarge() *GetOutboundEventRequestEntityTooLarge {
	return &GetOutboundEventRequestEntityTooLarge{}
}

/*GetOutboundEventRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundEventRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundEventRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventUnsupportedMediaType creates a GetOutboundEventUnsupportedMediaType with default headers values
func NewGetOutboundEventUnsupportedMediaType() *GetOutboundEventUnsupportedMediaType {
	return &GetOutboundEventUnsupportedMediaType{}
}

/*GetOutboundEventUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundEventUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundEventUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventTooManyRequests creates a GetOutboundEventTooManyRequests with default headers values
func NewGetOutboundEventTooManyRequests() *GetOutboundEventTooManyRequests {
	return &GetOutboundEventTooManyRequests{}
}

/*GetOutboundEventTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetOutboundEventTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundEventTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventInternalServerError creates a GetOutboundEventInternalServerError with default headers values
func NewGetOutboundEventInternalServerError() *GetOutboundEventInternalServerError {
	return &GetOutboundEventInternalServerError{}
}

/*GetOutboundEventInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundEventInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundEventInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventServiceUnavailable creates a GetOutboundEventServiceUnavailable with default headers values
func NewGetOutboundEventServiceUnavailable() *GetOutboundEventServiceUnavailable {
	return &GetOutboundEventServiceUnavailable{}
}

/*GetOutboundEventServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundEventServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundEventServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundEventGatewayTimeout creates a GetOutboundEventGatewayTimeout with default headers values
func NewGetOutboundEventGatewayTimeout() *GetOutboundEventGatewayTimeout {
	return &GetOutboundEventGatewayTimeout{}
}

/*GetOutboundEventGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundEventGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundEventGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/events/{eventId}][%d] getOutboundEventGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundEventGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundEventGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
