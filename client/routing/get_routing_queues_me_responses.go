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

// GetRoutingQueuesMeReader is a Reader for the GetRoutingQueuesMe structure.
type GetRoutingQueuesMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueuesMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueuesMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueuesMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueuesMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueuesMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueuesMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueuesMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueuesMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueuesMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueuesMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueuesMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueuesMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueuesMeOK creates a GetRoutingQueuesMeOK with default headers values
func NewGetRoutingQueuesMeOK() *GetRoutingQueuesMeOK {
	return &GetRoutingQueuesMeOK{}
}

/*GetRoutingQueuesMeOK handles this case with default header values.

successful operation
*/
type GetRoutingQueuesMeOK struct {
	Payload *models.UserQueueEntityListing
}

func (o *GetRoutingQueuesMeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueuesMeOK) GetPayload() *models.UserQueueEntityListing {
	return o.Payload
}

func (o *GetRoutingQueuesMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeBadRequest creates a GetRoutingQueuesMeBadRequest with default headers values
func NewGetRoutingQueuesMeBadRequest() *GetRoutingQueuesMeBadRequest {
	return &GetRoutingQueuesMeBadRequest{}
}

/*GetRoutingQueuesMeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueuesMeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueuesMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeUnauthorized creates a GetRoutingQueuesMeUnauthorized with default headers values
func NewGetRoutingQueuesMeUnauthorized() *GetRoutingQueuesMeUnauthorized {
	return &GetRoutingQueuesMeUnauthorized{}
}

/*GetRoutingQueuesMeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueuesMeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueuesMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeForbidden creates a GetRoutingQueuesMeForbidden with default headers values
func NewGetRoutingQueuesMeForbidden() *GetRoutingQueuesMeForbidden {
	return &GetRoutingQueuesMeForbidden{}
}

/*GetRoutingQueuesMeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueuesMeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueuesMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeNotFound creates a GetRoutingQueuesMeNotFound with default headers values
func NewGetRoutingQueuesMeNotFound() *GetRoutingQueuesMeNotFound {
	return &GetRoutingQueuesMeNotFound{}
}

/*GetRoutingQueuesMeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueuesMeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueuesMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeRequestEntityTooLarge creates a GetRoutingQueuesMeRequestEntityTooLarge with default headers values
func NewGetRoutingQueuesMeRequestEntityTooLarge() *GetRoutingQueuesMeRequestEntityTooLarge {
	return &GetRoutingQueuesMeRequestEntityTooLarge{}
}

/*GetRoutingQueuesMeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingQueuesMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueuesMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeUnsupportedMediaType creates a GetRoutingQueuesMeUnsupportedMediaType with default headers values
func NewGetRoutingQueuesMeUnsupportedMediaType() *GetRoutingQueuesMeUnsupportedMediaType {
	return &GetRoutingQueuesMeUnsupportedMediaType{}
}

/*GetRoutingQueuesMeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueuesMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueuesMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeTooManyRequests creates a GetRoutingQueuesMeTooManyRequests with default headers values
func NewGetRoutingQueuesMeTooManyRequests() *GetRoutingQueuesMeTooManyRequests {
	return &GetRoutingQueuesMeTooManyRequests{}
}

/*GetRoutingQueuesMeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingQueuesMeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueuesMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeInternalServerError creates a GetRoutingQueuesMeInternalServerError with default headers values
func NewGetRoutingQueuesMeInternalServerError() *GetRoutingQueuesMeInternalServerError {
	return &GetRoutingQueuesMeInternalServerError{}
}

/*GetRoutingQueuesMeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueuesMeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueuesMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeServiceUnavailable creates a GetRoutingQueuesMeServiceUnavailable with default headers values
func NewGetRoutingQueuesMeServiceUnavailable() *GetRoutingQueuesMeServiceUnavailable {
	return &GetRoutingQueuesMeServiceUnavailable{}
}

/*GetRoutingQueuesMeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueuesMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueuesMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesMeGatewayTimeout creates a GetRoutingQueuesMeGatewayTimeout with default headers values
func NewGetRoutingQueuesMeGatewayTimeout() *GetRoutingQueuesMeGatewayTimeout {
	return &GetRoutingQueuesMeGatewayTimeout{}
}

/*GetRoutingQueuesMeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueuesMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/me][%d] getRoutingQueuesMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueuesMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
