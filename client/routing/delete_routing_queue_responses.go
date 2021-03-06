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

// DeleteRoutingQueueReader is a Reader for the DeleteRoutingQueue structure.
type DeleteRoutingQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRoutingQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingQueueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingQueueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingQueueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingQueueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingQueueRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingQueueUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingQueueTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingQueueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingQueueServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingQueueGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingQueueOK creates a DeleteRoutingQueueOK with default headers values
func NewDeleteRoutingQueueOK() *DeleteRoutingQueueOK {
	return &DeleteRoutingQueueOK{}
}

/*DeleteRoutingQueueOK handles this case with default header values.

Operation was successful.
*/
type DeleteRoutingQueueOK struct {
}

func (o *DeleteRoutingQueueOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueOK ", 200)
}

func (o *DeleteRoutingQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingQueueBadRequest creates a DeleteRoutingQueueBadRequest with default headers values
func NewDeleteRoutingQueueBadRequest() *DeleteRoutingQueueBadRequest {
	return &DeleteRoutingQueueBadRequest{}
}

/*DeleteRoutingQueueBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingQueueBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingQueueBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueUnauthorized creates a DeleteRoutingQueueUnauthorized with default headers values
func NewDeleteRoutingQueueUnauthorized() *DeleteRoutingQueueUnauthorized {
	return &DeleteRoutingQueueUnauthorized{}
}

/*DeleteRoutingQueueUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingQueueUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingQueueUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueForbidden creates a DeleteRoutingQueueForbidden with default headers values
func NewDeleteRoutingQueueForbidden() *DeleteRoutingQueueForbidden {
	return &DeleteRoutingQueueForbidden{}
}

/*DeleteRoutingQueueForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingQueueForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingQueueForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueNotFound creates a DeleteRoutingQueueNotFound with default headers values
func NewDeleteRoutingQueueNotFound() *DeleteRoutingQueueNotFound {
	return &DeleteRoutingQueueNotFound{}
}

/*DeleteRoutingQueueNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRoutingQueueNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingQueueNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueRequestEntityTooLarge creates a DeleteRoutingQueueRequestEntityTooLarge with default headers values
func NewDeleteRoutingQueueRequestEntityTooLarge() *DeleteRoutingQueueRequestEntityTooLarge {
	return &DeleteRoutingQueueRequestEntityTooLarge{}
}

/*DeleteRoutingQueueRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRoutingQueueRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingQueueRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueUnsupportedMediaType creates a DeleteRoutingQueueUnsupportedMediaType with default headers values
func NewDeleteRoutingQueueUnsupportedMediaType() *DeleteRoutingQueueUnsupportedMediaType {
	return &DeleteRoutingQueueUnsupportedMediaType{}
}

/*DeleteRoutingQueueUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingQueueUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingQueueUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueTooManyRequests creates a DeleteRoutingQueueTooManyRequests with default headers values
func NewDeleteRoutingQueueTooManyRequests() *DeleteRoutingQueueTooManyRequests {
	return &DeleteRoutingQueueTooManyRequests{}
}

/*DeleteRoutingQueueTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRoutingQueueTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingQueueTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueInternalServerError creates a DeleteRoutingQueueInternalServerError with default headers values
func NewDeleteRoutingQueueInternalServerError() *DeleteRoutingQueueInternalServerError {
	return &DeleteRoutingQueueInternalServerError{}
}

/*DeleteRoutingQueueInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingQueueInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingQueueInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueServiceUnavailable creates a DeleteRoutingQueueServiceUnavailable with default headers values
func NewDeleteRoutingQueueServiceUnavailable() *DeleteRoutingQueueServiceUnavailable {
	return &DeleteRoutingQueueServiceUnavailable{}
}

/*DeleteRoutingQueueServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingQueueServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingQueueServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingQueueGatewayTimeout creates a DeleteRoutingQueueGatewayTimeout with default headers values
func NewDeleteRoutingQueueGatewayTimeout() *DeleteRoutingQueueGatewayTimeout {
	return &DeleteRoutingQueueGatewayTimeout{}
}

/*DeleteRoutingQueueGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRoutingQueueGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingQueueGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/queues/{queueId}][%d] deleteRoutingQueueGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingQueueGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingQueueGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
