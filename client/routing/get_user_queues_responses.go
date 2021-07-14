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

// GetUserQueuesReader is a Reader for the GetUserQueues structure.
type GetUserQueuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserQueuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserQueuesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserQueuesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserQueuesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserQueuesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserQueuesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserQueuesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserQueuesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserQueuesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserQueuesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserQueuesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserQueuesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserQueuesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserQueuesOK creates a GetUserQueuesOK with default headers values
func NewGetUserQueuesOK() *GetUserQueuesOK {
	return &GetUserQueuesOK{}
}

/*GetUserQueuesOK handles this case with default header values.

successful operation
*/
type GetUserQueuesOK struct {
	Payload *models.UserQueueEntityListing
}

func (o *GetUserQueuesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesOK  %+v", 200, o.Payload)
}

func (o *GetUserQueuesOK) GetPayload() *models.UserQueueEntityListing {
	return o.Payload
}

func (o *GetUserQueuesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserQueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesBadRequest creates a GetUserQueuesBadRequest with default headers values
func NewGetUserQueuesBadRequest() *GetUserQueuesBadRequest {
	return &GetUserQueuesBadRequest{}
}

/*GetUserQueuesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserQueuesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserQueuesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesUnauthorized creates a GetUserQueuesUnauthorized with default headers values
func NewGetUserQueuesUnauthorized() *GetUserQueuesUnauthorized {
	return &GetUserQueuesUnauthorized{}
}

/*GetUserQueuesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserQueuesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserQueuesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesForbidden creates a GetUserQueuesForbidden with default headers values
func NewGetUserQueuesForbidden() *GetUserQueuesForbidden {
	return &GetUserQueuesForbidden{}
}

/*GetUserQueuesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserQueuesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesForbidden  %+v", 403, o.Payload)
}

func (o *GetUserQueuesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesNotFound creates a GetUserQueuesNotFound with default headers values
func NewGetUserQueuesNotFound() *GetUserQueuesNotFound {
	return &GetUserQueuesNotFound{}
}

/*GetUserQueuesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserQueuesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesNotFound  %+v", 404, o.Payload)
}

func (o *GetUserQueuesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesRequestTimeout creates a GetUserQueuesRequestTimeout with default headers values
func NewGetUserQueuesRequestTimeout() *GetUserQueuesRequestTimeout {
	return &GetUserQueuesRequestTimeout{}
}

/*GetUserQueuesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserQueuesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserQueuesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesRequestEntityTooLarge creates a GetUserQueuesRequestEntityTooLarge with default headers values
func NewGetUserQueuesRequestEntityTooLarge() *GetUserQueuesRequestEntityTooLarge {
	return &GetUserQueuesRequestEntityTooLarge{}
}

/*GetUserQueuesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserQueuesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserQueuesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesUnsupportedMediaType creates a GetUserQueuesUnsupportedMediaType with default headers values
func NewGetUserQueuesUnsupportedMediaType() *GetUserQueuesUnsupportedMediaType {
	return &GetUserQueuesUnsupportedMediaType{}
}

/*GetUserQueuesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserQueuesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserQueuesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesTooManyRequests creates a GetUserQueuesTooManyRequests with default headers values
func NewGetUserQueuesTooManyRequests() *GetUserQueuesTooManyRequests {
	return &GetUserQueuesTooManyRequests{}
}

/*GetUserQueuesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserQueuesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserQueuesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesInternalServerError creates a GetUserQueuesInternalServerError with default headers values
func NewGetUserQueuesInternalServerError() *GetUserQueuesInternalServerError {
	return &GetUserQueuesInternalServerError{}
}

/*GetUserQueuesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserQueuesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserQueuesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesServiceUnavailable creates a GetUserQueuesServiceUnavailable with default headers values
func NewGetUserQueuesServiceUnavailable() *GetUserQueuesServiceUnavailable {
	return &GetUserQueuesServiceUnavailable{}
}

/*GetUserQueuesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserQueuesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserQueuesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserQueuesGatewayTimeout creates a GetUserQueuesGatewayTimeout with default headers values
func NewGetUserQueuesGatewayTimeout() *GetUserQueuesGatewayTimeout {
	return &GetUserQueuesGatewayTimeout{}
}

/*GetUserQueuesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserQueuesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserQueuesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/queues][%d] getUserQueuesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserQueuesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserQueuesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
