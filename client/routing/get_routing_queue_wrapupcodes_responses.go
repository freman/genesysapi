// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingQueueWrapupcodesReader is a Reader for the GetRoutingQueueWrapupcodes structure.
type GetRoutingQueueWrapupcodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueueWrapupcodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueueWrapupcodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueueWrapupcodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueueWrapupcodesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueueWrapupcodesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueueWrapupcodesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingQueueWrapupcodesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueueWrapupcodesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueueWrapupcodesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueueWrapupcodesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueueWrapupcodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueueWrapupcodesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueueWrapupcodesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueueWrapupcodesOK creates a GetRoutingQueueWrapupcodesOK with default headers values
func NewGetRoutingQueueWrapupcodesOK() *GetRoutingQueueWrapupcodesOK {
	return &GetRoutingQueueWrapupcodesOK{}
}

/*GetRoutingQueueWrapupcodesOK handles this case with default header values.

successful operation
*/
type GetRoutingQueueWrapupcodesOK struct {
	Payload *models.WrapupCodeEntityListing
}

func (o *GetRoutingQueueWrapupcodesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesOK) GetPayload() *models.WrapupCodeEntityListing {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WrapupCodeEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesBadRequest creates a GetRoutingQueueWrapupcodesBadRequest with default headers values
func NewGetRoutingQueueWrapupcodesBadRequest() *GetRoutingQueueWrapupcodesBadRequest {
	return &GetRoutingQueueWrapupcodesBadRequest{}
}

/*GetRoutingQueueWrapupcodesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueueWrapupcodesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesUnauthorized creates a GetRoutingQueueWrapupcodesUnauthorized with default headers values
func NewGetRoutingQueueWrapupcodesUnauthorized() *GetRoutingQueueWrapupcodesUnauthorized {
	return &GetRoutingQueueWrapupcodesUnauthorized{}
}

/*GetRoutingQueueWrapupcodesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueueWrapupcodesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesForbidden creates a GetRoutingQueueWrapupcodesForbidden with default headers values
func NewGetRoutingQueueWrapupcodesForbidden() *GetRoutingQueueWrapupcodesForbidden {
	return &GetRoutingQueueWrapupcodesForbidden{}
}

/*GetRoutingQueueWrapupcodesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueueWrapupcodesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesNotFound creates a GetRoutingQueueWrapupcodesNotFound with default headers values
func NewGetRoutingQueueWrapupcodesNotFound() *GetRoutingQueueWrapupcodesNotFound {
	return &GetRoutingQueueWrapupcodesNotFound{}
}

/*GetRoutingQueueWrapupcodesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueueWrapupcodesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesRequestTimeout creates a GetRoutingQueueWrapupcodesRequestTimeout with default headers values
func NewGetRoutingQueueWrapupcodesRequestTimeout() *GetRoutingQueueWrapupcodesRequestTimeout {
	return &GetRoutingQueueWrapupcodesRequestTimeout{}
}

/*GetRoutingQueueWrapupcodesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingQueueWrapupcodesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesRequestEntityTooLarge creates a GetRoutingQueueWrapupcodesRequestEntityTooLarge with default headers values
func NewGetRoutingQueueWrapupcodesRequestEntityTooLarge() *GetRoutingQueueWrapupcodesRequestEntityTooLarge {
	return &GetRoutingQueueWrapupcodesRequestEntityTooLarge{}
}

/*GetRoutingQueueWrapupcodesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingQueueWrapupcodesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesUnsupportedMediaType creates a GetRoutingQueueWrapupcodesUnsupportedMediaType with default headers values
func NewGetRoutingQueueWrapupcodesUnsupportedMediaType() *GetRoutingQueueWrapupcodesUnsupportedMediaType {
	return &GetRoutingQueueWrapupcodesUnsupportedMediaType{}
}

/*GetRoutingQueueWrapupcodesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueueWrapupcodesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesTooManyRequests creates a GetRoutingQueueWrapupcodesTooManyRequests with default headers values
func NewGetRoutingQueueWrapupcodesTooManyRequests() *GetRoutingQueueWrapupcodesTooManyRequests {
	return &GetRoutingQueueWrapupcodesTooManyRequests{}
}

/*GetRoutingQueueWrapupcodesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingQueueWrapupcodesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesInternalServerError creates a GetRoutingQueueWrapupcodesInternalServerError with default headers values
func NewGetRoutingQueueWrapupcodesInternalServerError() *GetRoutingQueueWrapupcodesInternalServerError {
	return &GetRoutingQueueWrapupcodesInternalServerError{}
}

/*GetRoutingQueueWrapupcodesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueueWrapupcodesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesServiceUnavailable creates a GetRoutingQueueWrapupcodesServiceUnavailable with default headers values
func NewGetRoutingQueueWrapupcodesServiceUnavailable() *GetRoutingQueueWrapupcodesServiceUnavailable {
	return &GetRoutingQueueWrapupcodesServiceUnavailable{}
}

/*GetRoutingQueueWrapupcodesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueueWrapupcodesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueWrapupcodesGatewayTimeout creates a GetRoutingQueueWrapupcodesGatewayTimeout with default headers values
func NewGetRoutingQueueWrapupcodesGatewayTimeout() *GetRoutingQueueWrapupcodesGatewayTimeout {
	return &GetRoutingQueueWrapupcodesGatewayTimeout{}
}

/*GetRoutingQueueWrapupcodesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueueWrapupcodesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueWrapupcodesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/wrapupcodes][%d] getRoutingQueueWrapupcodesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueWrapupcodesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueWrapupcodesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
