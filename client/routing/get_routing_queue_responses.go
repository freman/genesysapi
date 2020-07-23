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

// GetRoutingQueueReader is a Reader for the GetRoutingQueue structure.
type GetRoutingQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueueRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueueUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueueTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueueOK creates a GetRoutingQueueOK with default headers values
func NewGetRoutingQueueOK() *GetRoutingQueueOK {
	return &GetRoutingQueueOK{}
}

/*GetRoutingQueueOK handles this case with default header values.

successful operation
*/
type GetRoutingQueueOK struct {
	Payload *models.Queue
}

func (o *GetRoutingQueueOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueOK) GetPayload() *models.Queue {
	return o.Payload
}

func (o *GetRoutingQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Queue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueBadRequest creates a GetRoutingQueueBadRequest with default headers values
func NewGetRoutingQueueBadRequest() *GetRoutingQueueBadRequest {
	return &GetRoutingQueueBadRequest{}
}

/*GetRoutingQueueBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueueBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUnauthorized creates a GetRoutingQueueUnauthorized with default headers values
func NewGetRoutingQueueUnauthorized() *GetRoutingQueueUnauthorized {
	return &GetRoutingQueueUnauthorized{}
}

/*GetRoutingQueueUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueueUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueForbidden creates a GetRoutingQueueForbidden with default headers values
func NewGetRoutingQueueForbidden() *GetRoutingQueueForbidden {
	return &GetRoutingQueueForbidden{}
}

/*GetRoutingQueueForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueueForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueNotFound creates a GetRoutingQueueNotFound with default headers values
func NewGetRoutingQueueNotFound() *GetRoutingQueueNotFound {
	return &GetRoutingQueueNotFound{}
}

/*GetRoutingQueueNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueueNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueRequestEntityTooLarge creates a GetRoutingQueueRequestEntityTooLarge with default headers values
func NewGetRoutingQueueRequestEntityTooLarge() *GetRoutingQueueRequestEntityTooLarge {
	return &GetRoutingQueueRequestEntityTooLarge{}
}

/*GetRoutingQueueRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingQueueRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueUnsupportedMediaType creates a GetRoutingQueueUnsupportedMediaType with default headers values
func NewGetRoutingQueueUnsupportedMediaType() *GetRoutingQueueUnsupportedMediaType {
	return &GetRoutingQueueUnsupportedMediaType{}
}

/*GetRoutingQueueUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueueUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueTooManyRequests creates a GetRoutingQueueTooManyRequests with default headers values
func NewGetRoutingQueueTooManyRequests() *GetRoutingQueueTooManyRequests {
	return &GetRoutingQueueTooManyRequests{}
}

/*GetRoutingQueueTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingQueueTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueInternalServerError creates a GetRoutingQueueInternalServerError with default headers values
func NewGetRoutingQueueInternalServerError() *GetRoutingQueueInternalServerError {
	return &GetRoutingQueueInternalServerError{}
}

/*GetRoutingQueueInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueueInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueServiceUnavailable creates a GetRoutingQueueServiceUnavailable with default headers values
func NewGetRoutingQueueServiceUnavailable() *GetRoutingQueueServiceUnavailable {
	return &GetRoutingQueueServiceUnavailable{}
}

/*GetRoutingQueueServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueueServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueGatewayTimeout creates a GetRoutingQueueGatewayTimeout with default headers values
func NewGetRoutingQueueGatewayTimeout() *GetRoutingQueueGatewayTimeout {
	return &GetRoutingQueueGatewayTimeout{}
}

/*GetRoutingQueueGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueueGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}][%d] getRoutingQueueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
