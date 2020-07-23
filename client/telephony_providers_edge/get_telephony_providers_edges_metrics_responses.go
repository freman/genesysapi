// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetTelephonyProvidersEdgesMetricsReader is a Reader for the GetTelephonyProvidersEdgesMetrics structure.
type GetTelephonyProvidersEdgesMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesMetricsOK creates a GetTelephonyProvidersEdgesMetricsOK with default headers values
func NewGetTelephonyProvidersEdgesMetricsOK() *GetTelephonyProvidersEdgesMetricsOK {
	return &GetTelephonyProvidersEdgesMetricsOK{}
}

/*GetTelephonyProvidersEdgesMetricsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesMetricsOK struct {
	Payload []*models.EdgeMetrics
}

func (o *GetTelephonyProvidersEdgesMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsOK) GetPayload() []*models.EdgeMetrics {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsBadRequest creates a GetTelephonyProvidersEdgesMetricsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesMetricsBadRequest() *GetTelephonyProvidersEdgesMetricsBadRequest {
	return &GetTelephonyProvidersEdgesMetricsBadRequest{}
}

/*GetTelephonyProvidersEdgesMetricsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesMetricsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsUnauthorized creates a GetTelephonyProvidersEdgesMetricsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesMetricsUnauthorized() *GetTelephonyProvidersEdgesMetricsUnauthorized {
	return &GetTelephonyProvidersEdgesMetricsUnauthorized{}
}

/*GetTelephonyProvidersEdgesMetricsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsForbidden creates a GetTelephonyProvidersEdgesMetricsForbidden with default headers values
func NewGetTelephonyProvidersEdgesMetricsForbidden() *GetTelephonyProvidersEdgesMetricsForbidden {
	return &GetTelephonyProvidersEdgesMetricsForbidden{}
}

/*GetTelephonyProvidersEdgesMetricsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesMetricsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsNotFound creates a GetTelephonyProvidersEdgesMetricsNotFound with default headers values
func NewGetTelephonyProvidersEdgesMetricsNotFound() *GetTelephonyProvidersEdgesMetricsNotFound {
	return &GetTelephonyProvidersEdgesMetricsNotFound{}
}

/*GetTelephonyProvidersEdgesMetricsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesMetricsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesMetricsRequestEntityTooLarge() *GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsUnsupportedMediaType creates a GetTelephonyProvidersEdgesMetricsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesMetricsUnsupportedMediaType() *GetTelephonyProvidersEdgesMetricsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesMetricsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesMetricsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsTooManyRequests creates a GetTelephonyProvidersEdgesMetricsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesMetricsTooManyRequests() *GetTelephonyProvidersEdgesMetricsTooManyRequests {
	return &GetTelephonyProvidersEdgesMetricsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesMetricsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsInternalServerError creates a GetTelephonyProvidersEdgesMetricsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesMetricsInternalServerError() *GetTelephonyProvidersEdgesMetricsInternalServerError {
	return &GetTelephonyProvidersEdgesMetricsInternalServerError{}
}

/*GetTelephonyProvidersEdgesMetricsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsServiceUnavailable creates a GetTelephonyProvidersEdgesMetricsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesMetricsServiceUnavailable() *GetTelephonyProvidersEdgesMetricsServiceUnavailable {
	return &GetTelephonyProvidersEdgesMetricsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesMetricsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesMetricsGatewayTimeout creates a GetTelephonyProvidersEdgesMetricsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesMetricsGatewayTimeout() *GetTelephonyProvidersEdgesMetricsGatewayTimeout {
	return &GetTelephonyProvidersEdgesMetricsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesMetricsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/metrics][%d] getTelephonyProvidersEdgesMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
