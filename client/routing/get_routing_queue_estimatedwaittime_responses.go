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

// GetRoutingQueueEstimatedwaittimeReader is a Reader for the GetRoutingQueueEstimatedwaittime structure.
type GetRoutingQueueEstimatedwaittimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueueEstimatedwaittimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueueEstimatedwaittimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueueEstimatedwaittimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueueEstimatedwaittimeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueueEstimatedwaittimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueueEstimatedwaittimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingQueueEstimatedwaittimeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueueEstimatedwaittimeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueueEstimatedwaittimeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueueEstimatedwaittimeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueueEstimatedwaittimeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueueEstimatedwaittimeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueueEstimatedwaittimeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueueEstimatedwaittimeOK creates a GetRoutingQueueEstimatedwaittimeOK with default headers values
func NewGetRoutingQueueEstimatedwaittimeOK() *GetRoutingQueueEstimatedwaittimeOK {
	return &GetRoutingQueueEstimatedwaittimeOK{}
}

/*GetRoutingQueueEstimatedwaittimeOK handles this case with default header values.

successful operation
*/
type GetRoutingQueueEstimatedwaittimeOK struct {
	Payload *models.EstimatedWaitTimePredictions
}

func (o *GetRoutingQueueEstimatedwaittimeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeOK) GetPayload() *models.EstimatedWaitTimePredictions {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimatedWaitTimePredictions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeBadRequest creates a GetRoutingQueueEstimatedwaittimeBadRequest with default headers values
func NewGetRoutingQueueEstimatedwaittimeBadRequest() *GetRoutingQueueEstimatedwaittimeBadRequest {
	return &GetRoutingQueueEstimatedwaittimeBadRequest{}
}

/*GetRoutingQueueEstimatedwaittimeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueueEstimatedwaittimeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeUnauthorized creates a GetRoutingQueueEstimatedwaittimeUnauthorized with default headers values
func NewGetRoutingQueueEstimatedwaittimeUnauthorized() *GetRoutingQueueEstimatedwaittimeUnauthorized {
	return &GetRoutingQueueEstimatedwaittimeUnauthorized{}
}

/*GetRoutingQueueEstimatedwaittimeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueueEstimatedwaittimeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeForbidden creates a GetRoutingQueueEstimatedwaittimeForbidden with default headers values
func NewGetRoutingQueueEstimatedwaittimeForbidden() *GetRoutingQueueEstimatedwaittimeForbidden {
	return &GetRoutingQueueEstimatedwaittimeForbidden{}
}

/*GetRoutingQueueEstimatedwaittimeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueueEstimatedwaittimeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeNotFound creates a GetRoutingQueueEstimatedwaittimeNotFound with default headers values
func NewGetRoutingQueueEstimatedwaittimeNotFound() *GetRoutingQueueEstimatedwaittimeNotFound {
	return &GetRoutingQueueEstimatedwaittimeNotFound{}
}

/*GetRoutingQueueEstimatedwaittimeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingQueueEstimatedwaittimeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeRequestTimeout creates a GetRoutingQueueEstimatedwaittimeRequestTimeout with default headers values
func NewGetRoutingQueueEstimatedwaittimeRequestTimeout() *GetRoutingQueueEstimatedwaittimeRequestTimeout {
	return &GetRoutingQueueEstimatedwaittimeRequestTimeout{}
}

/*GetRoutingQueueEstimatedwaittimeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingQueueEstimatedwaittimeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeRequestEntityTooLarge creates a GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge with default headers values
func NewGetRoutingQueueEstimatedwaittimeRequestEntityTooLarge() *GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge {
	return &GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge{}
}

/*GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeUnsupportedMediaType creates a GetRoutingQueueEstimatedwaittimeUnsupportedMediaType with default headers values
func NewGetRoutingQueueEstimatedwaittimeUnsupportedMediaType() *GetRoutingQueueEstimatedwaittimeUnsupportedMediaType {
	return &GetRoutingQueueEstimatedwaittimeUnsupportedMediaType{}
}

/*GetRoutingQueueEstimatedwaittimeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueueEstimatedwaittimeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeTooManyRequests creates a GetRoutingQueueEstimatedwaittimeTooManyRequests with default headers values
func NewGetRoutingQueueEstimatedwaittimeTooManyRequests() *GetRoutingQueueEstimatedwaittimeTooManyRequests {
	return &GetRoutingQueueEstimatedwaittimeTooManyRequests{}
}

/*GetRoutingQueueEstimatedwaittimeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingQueueEstimatedwaittimeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeInternalServerError creates a GetRoutingQueueEstimatedwaittimeInternalServerError with default headers values
func NewGetRoutingQueueEstimatedwaittimeInternalServerError() *GetRoutingQueueEstimatedwaittimeInternalServerError {
	return &GetRoutingQueueEstimatedwaittimeInternalServerError{}
}

/*GetRoutingQueueEstimatedwaittimeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueueEstimatedwaittimeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeServiceUnavailable creates a GetRoutingQueueEstimatedwaittimeServiceUnavailable with default headers values
func NewGetRoutingQueueEstimatedwaittimeServiceUnavailable() *GetRoutingQueueEstimatedwaittimeServiceUnavailable {
	return &GetRoutingQueueEstimatedwaittimeServiceUnavailable{}
}

/*GetRoutingQueueEstimatedwaittimeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueueEstimatedwaittimeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueEstimatedwaittimeGatewayTimeout creates a GetRoutingQueueEstimatedwaittimeGatewayTimeout with default headers values
func NewGetRoutingQueueEstimatedwaittimeGatewayTimeout() *GetRoutingQueueEstimatedwaittimeGatewayTimeout {
	return &GetRoutingQueueEstimatedwaittimeGatewayTimeout{}
}

/*GetRoutingQueueEstimatedwaittimeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingQueueEstimatedwaittimeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingQueueEstimatedwaittimeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/estimatedwaittime][%d] getRoutingQueueEstimatedwaittimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueEstimatedwaittimeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueEstimatedwaittimeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
