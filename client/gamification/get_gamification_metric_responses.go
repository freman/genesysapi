// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationMetricReader is a Reader for the GetGamificationMetric structure.
type GetGamificationMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationMetricBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationMetricUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationMetricForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationMetricNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationMetricRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationMetricRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationMetricUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationMetricTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationMetricInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationMetricServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationMetricGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationMetricOK creates a GetGamificationMetricOK with default headers values
func NewGetGamificationMetricOK() *GetGamificationMetricOK {
	return &GetGamificationMetricOK{}
}

/*GetGamificationMetricOK handles this case with default header values.

successful operation
*/
type GetGamificationMetricOK struct {
	Payload *models.Metric
}

func (o *GetGamificationMetricOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricOK  %+v", 200, o.Payload)
}

func (o *GetGamificationMetricOK) GetPayload() *models.Metric {
	return o.Payload
}

func (o *GetGamificationMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricBadRequest creates a GetGamificationMetricBadRequest with default headers values
func NewGetGamificationMetricBadRequest() *GetGamificationMetricBadRequest {
	return &GetGamificationMetricBadRequest{}
}

/*GetGamificationMetricBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationMetricBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationMetricBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricUnauthorized creates a GetGamificationMetricUnauthorized with default headers values
func NewGetGamificationMetricUnauthorized() *GetGamificationMetricUnauthorized {
	return &GetGamificationMetricUnauthorized{}
}

/*GetGamificationMetricUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationMetricUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationMetricUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricForbidden creates a GetGamificationMetricForbidden with default headers values
func NewGetGamificationMetricForbidden() *GetGamificationMetricForbidden {
	return &GetGamificationMetricForbidden{}
}

/*GetGamificationMetricForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationMetricForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationMetricForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricNotFound creates a GetGamificationMetricNotFound with default headers values
func NewGetGamificationMetricNotFound() *GetGamificationMetricNotFound {
	return &GetGamificationMetricNotFound{}
}

/*GetGamificationMetricNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationMetricNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationMetricNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricRequestTimeout creates a GetGamificationMetricRequestTimeout with default headers values
func NewGetGamificationMetricRequestTimeout() *GetGamificationMetricRequestTimeout {
	return &GetGamificationMetricRequestTimeout{}
}

/*GetGamificationMetricRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationMetricRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationMetricRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricRequestEntityTooLarge creates a GetGamificationMetricRequestEntityTooLarge with default headers values
func NewGetGamificationMetricRequestEntityTooLarge() *GetGamificationMetricRequestEntityTooLarge {
	return &GetGamificationMetricRequestEntityTooLarge{}
}

/*GetGamificationMetricRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationMetricRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationMetricRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricUnsupportedMediaType creates a GetGamificationMetricUnsupportedMediaType with default headers values
func NewGetGamificationMetricUnsupportedMediaType() *GetGamificationMetricUnsupportedMediaType {
	return &GetGamificationMetricUnsupportedMediaType{}
}

/*GetGamificationMetricUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationMetricUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationMetricUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricTooManyRequests creates a GetGamificationMetricTooManyRequests with default headers values
func NewGetGamificationMetricTooManyRequests() *GetGamificationMetricTooManyRequests {
	return &GetGamificationMetricTooManyRequests{}
}

/*GetGamificationMetricTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationMetricTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationMetricTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricInternalServerError creates a GetGamificationMetricInternalServerError with default headers values
func NewGetGamificationMetricInternalServerError() *GetGamificationMetricInternalServerError {
	return &GetGamificationMetricInternalServerError{}
}

/*GetGamificationMetricInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationMetricInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationMetricInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricServiceUnavailable creates a GetGamificationMetricServiceUnavailable with default headers values
func NewGetGamificationMetricServiceUnavailable() *GetGamificationMetricServiceUnavailable {
	return &GetGamificationMetricServiceUnavailable{}
}

/*GetGamificationMetricServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationMetricServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationMetricServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationMetricGatewayTimeout creates a GetGamificationMetricGatewayTimeout with default headers values
func NewGetGamificationMetricGatewayTimeout() *GetGamificationMetricGatewayTimeout {
	return &GetGamificationMetricGatewayTimeout{}
}

/*GetGamificationMetricGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationMetricGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationMetricGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/metrics/{metricId}][%d] getGamificationMetricGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationMetricGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationMetricGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
