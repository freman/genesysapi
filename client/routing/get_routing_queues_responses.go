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

// GetRoutingQueuesReader is a Reader for the GetRoutingQueues structure.
type GetRoutingQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueuesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueuesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueuesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueuesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueuesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueuesOK creates a GetRoutingQueuesOK with default headers values
func NewGetRoutingQueuesOK() *GetRoutingQueuesOK {
	return &GetRoutingQueuesOK{}
}

/*GetRoutingQueuesOK handles this case with default header values.

successful operation
*/
type GetRoutingQueuesOK struct {
	Payload *models.QueueEntityListing
}

func (o *GetRoutingQueuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueuesOK) GetPayload() *models.QueueEntityListing {
	return o.Payload
}

func (o *GetRoutingQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesBadRequest creates a GetRoutingQueuesBadRequest with default headers values
func NewGetRoutingQueuesBadRequest() *GetRoutingQueuesBadRequest {
	return &GetRoutingQueuesBadRequest{}
}

/*GetRoutingQueuesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueuesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueuesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesUnauthorized creates a GetRoutingQueuesUnauthorized with default headers values
func NewGetRoutingQueuesUnauthorized() *GetRoutingQueuesUnauthorized {
	return &GetRoutingQueuesUnauthorized{}
}

/*GetRoutingQueuesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueuesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueuesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesForbidden creates a GetRoutingQueuesForbidden with default headers values
func NewGetRoutingQueuesForbidden() *GetRoutingQueuesForbidden {
	return &GetRoutingQueuesForbidden{}
}

/*GetRoutingQueuesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueuesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueuesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesNotFound creates a GetRoutingQueuesNotFound with default headers values
func NewGetRoutingQueuesNotFound() *GetRoutingQueuesNotFound {
	return &GetRoutingQueuesNotFound{}
}

/*GetRoutingQueuesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueuesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueuesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesRequestEntityTooLarge creates a GetRoutingQueuesRequestEntityTooLarge with default headers values
func NewGetRoutingQueuesRequestEntityTooLarge() *GetRoutingQueuesRequestEntityTooLarge {
	return &GetRoutingQueuesRequestEntityTooLarge{}
}

/*GetRoutingQueuesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingQueuesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueuesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesUnsupportedMediaType creates a GetRoutingQueuesUnsupportedMediaType with default headers values
func NewGetRoutingQueuesUnsupportedMediaType() *GetRoutingQueuesUnsupportedMediaType {
	return &GetRoutingQueuesUnsupportedMediaType{}
}

/*GetRoutingQueuesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueuesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueuesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesTooManyRequests creates a GetRoutingQueuesTooManyRequests with default headers values
func NewGetRoutingQueuesTooManyRequests() *GetRoutingQueuesTooManyRequests {
	return &GetRoutingQueuesTooManyRequests{}
}

/*GetRoutingQueuesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingQueuesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueuesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesInternalServerError creates a GetRoutingQueuesInternalServerError with default headers values
func NewGetRoutingQueuesInternalServerError() *GetRoutingQueuesInternalServerError {
	return &GetRoutingQueuesInternalServerError{}
}

/*GetRoutingQueuesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueuesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueuesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesServiceUnavailable creates a GetRoutingQueuesServiceUnavailable with default headers values
func NewGetRoutingQueuesServiceUnavailable() *GetRoutingQueuesServiceUnavailable {
	return &GetRoutingQueuesServiceUnavailable{}
}

/*GetRoutingQueuesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueuesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueuesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesGatewayTimeout creates a GetRoutingQueuesGatewayTimeout with default headers values
func NewGetRoutingQueuesGatewayTimeout() *GetRoutingQueuesGatewayTimeout {
	return &GetRoutingQueuesGatewayTimeout{}
}

/*GetRoutingQueuesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueuesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues][%d] getRoutingQueuesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueuesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
