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

// GetRoutingQueuesDivisionviewsReader is a Reader for the GetRoutingQueuesDivisionviews structure.
type GetRoutingQueuesDivisionviewsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueuesDivisionviewsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueuesDivisionviewsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueuesDivisionviewsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueuesDivisionviewsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueuesDivisionviewsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueuesDivisionviewsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingQueuesDivisionviewsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueuesDivisionviewsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueuesDivisionviewsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueuesDivisionviewsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueuesDivisionviewsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueuesDivisionviewsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueuesDivisionviewsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueuesDivisionviewsOK creates a GetRoutingQueuesDivisionviewsOK with default headers values
func NewGetRoutingQueuesDivisionviewsOK() *GetRoutingQueuesDivisionviewsOK {
	return &GetRoutingQueuesDivisionviewsOK{}
}

/*GetRoutingQueuesDivisionviewsOK handles this case with default header values.

successful operation
*/
type GetRoutingQueuesDivisionviewsOK struct {
	Payload *models.QueueEntityListing
}

func (o *GetRoutingQueuesDivisionviewsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsOK) GetPayload() *models.QueueEntityListing {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsBadRequest creates a GetRoutingQueuesDivisionviewsBadRequest with default headers values
func NewGetRoutingQueuesDivisionviewsBadRequest() *GetRoutingQueuesDivisionviewsBadRequest {
	return &GetRoutingQueuesDivisionviewsBadRequest{}
}

/*GetRoutingQueuesDivisionviewsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueuesDivisionviewsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsUnauthorized creates a GetRoutingQueuesDivisionviewsUnauthorized with default headers values
func NewGetRoutingQueuesDivisionviewsUnauthorized() *GetRoutingQueuesDivisionviewsUnauthorized {
	return &GetRoutingQueuesDivisionviewsUnauthorized{}
}

/*GetRoutingQueuesDivisionviewsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueuesDivisionviewsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsForbidden creates a GetRoutingQueuesDivisionviewsForbidden with default headers values
func NewGetRoutingQueuesDivisionviewsForbidden() *GetRoutingQueuesDivisionviewsForbidden {
	return &GetRoutingQueuesDivisionviewsForbidden{}
}

/*GetRoutingQueuesDivisionviewsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueuesDivisionviewsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsNotFound creates a GetRoutingQueuesDivisionviewsNotFound with default headers values
func NewGetRoutingQueuesDivisionviewsNotFound() *GetRoutingQueuesDivisionviewsNotFound {
	return &GetRoutingQueuesDivisionviewsNotFound{}
}

/*GetRoutingQueuesDivisionviewsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueuesDivisionviewsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsRequestTimeout creates a GetRoutingQueuesDivisionviewsRequestTimeout with default headers values
func NewGetRoutingQueuesDivisionviewsRequestTimeout() *GetRoutingQueuesDivisionviewsRequestTimeout {
	return &GetRoutingQueuesDivisionviewsRequestTimeout{}
}

/*GetRoutingQueuesDivisionviewsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingQueuesDivisionviewsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsRequestEntityTooLarge creates a GetRoutingQueuesDivisionviewsRequestEntityTooLarge with default headers values
func NewGetRoutingQueuesDivisionviewsRequestEntityTooLarge() *GetRoutingQueuesDivisionviewsRequestEntityTooLarge {
	return &GetRoutingQueuesDivisionviewsRequestEntityTooLarge{}
}

/*GetRoutingQueuesDivisionviewsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRoutingQueuesDivisionviewsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsUnsupportedMediaType creates a GetRoutingQueuesDivisionviewsUnsupportedMediaType with default headers values
func NewGetRoutingQueuesDivisionviewsUnsupportedMediaType() *GetRoutingQueuesDivisionviewsUnsupportedMediaType {
	return &GetRoutingQueuesDivisionviewsUnsupportedMediaType{}
}

/*GetRoutingQueuesDivisionviewsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueuesDivisionviewsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsTooManyRequests creates a GetRoutingQueuesDivisionviewsTooManyRequests with default headers values
func NewGetRoutingQueuesDivisionviewsTooManyRequests() *GetRoutingQueuesDivisionviewsTooManyRequests {
	return &GetRoutingQueuesDivisionviewsTooManyRequests{}
}

/*GetRoutingQueuesDivisionviewsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingQueuesDivisionviewsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsInternalServerError creates a GetRoutingQueuesDivisionviewsInternalServerError with default headers values
func NewGetRoutingQueuesDivisionviewsInternalServerError() *GetRoutingQueuesDivisionviewsInternalServerError {
	return &GetRoutingQueuesDivisionviewsInternalServerError{}
}

/*GetRoutingQueuesDivisionviewsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueuesDivisionviewsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsServiceUnavailable creates a GetRoutingQueuesDivisionviewsServiceUnavailable with default headers values
func NewGetRoutingQueuesDivisionviewsServiceUnavailable() *GetRoutingQueuesDivisionviewsServiceUnavailable {
	return &GetRoutingQueuesDivisionviewsServiceUnavailable{}
}

/*GetRoutingQueuesDivisionviewsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueuesDivisionviewsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueuesDivisionviewsGatewayTimeout creates a GetRoutingQueuesDivisionviewsGatewayTimeout with default headers values
func NewGetRoutingQueuesDivisionviewsGatewayTimeout() *GetRoutingQueuesDivisionviewsGatewayTimeout {
	return &GetRoutingQueuesDivisionviewsGatewayTimeout{}
}

/*GetRoutingQueuesDivisionviewsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueuesDivisionviewsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueuesDivisionviewsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/divisionviews][%d] getRoutingQueuesDivisionviewsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueuesDivisionviewsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueuesDivisionviewsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
