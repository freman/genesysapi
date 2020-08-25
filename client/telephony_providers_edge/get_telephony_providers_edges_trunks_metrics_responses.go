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

// GetTelephonyProvidersEdgesTrunksMetricsReader is a Reader for the GetTelephonyProvidersEdgesTrunksMetrics structure.
type GetTelephonyProvidersEdgesTrunksMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesTrunksMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesTrunksMetricsOK creates a GetTelephonyProvidersEdgesTrunksMetricsOK with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsOK() *GetTelephonyProvidersEdgesTrunksMetricsOK {
	return &GetTelephonyProvidersEdgesTrunksMetricsOK{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesTrunksMetricsOK struct {
	Payload []*models.TrunkMetrics
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsOK) GetPayload() []*models.TrunkMetrics {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsBadRequest creates a GetTelephonyProvidersEdgesTrunksMetricsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsBadRequest() *GetTelephonyProvidersEdgesTrunksMetricsBadRequest {
	return &GetTelephonyProvidersEdgesTrunksMetricsBadRequest{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesTrunksMetricsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsUnauthorized creates a GetTelephonyProvidersEdgesTrunksMetricsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsUnauthorized() *GetTelephonyProvidersEdgesTrunksMetricsUnauthorized {
	return &GetTelephonyProvidersEdgesTrunksMetricsUnauthorized{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesTrunksMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsForbidden creates a GetTelephonyProvidersEdgesTrunksMetricsForbidden with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsForbidden() *GetTelephonyProvidersEdgesTrunksMetricsForbidden {
	return &GetTelephonyProvidersEdgesTrunksMetricsForbidden{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesTrunksMetricsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsNotFound creates a GetTelephonyProvidersEdgesTrunksMetricsNotFound with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsNotFound() *GetTelephonyProvidersEdgesTrunksMetricsNotFound {
	return &GetTelephonyProvidersEdgesTrunksMetricsNotFound{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesTrunksMetricsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge() *GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType creates a GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType() *GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsTooManyRequests creates a GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsTooManyRequests() *GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests {
	return &GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsInternalServerError creates a GetTelephonyProvidersEdgesTrunksMetricsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsInternalServerError() *GetTelephonyProvidersEdgesTrunksMetricsInternalServerError {
	return &GetTelephonyProvidersEdgesTrunksMetricsInternalServerError{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesTrunksMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable creates a GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable() *GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable {
	return &GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout creates a GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout() *GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout {
	return &GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/metrics][%d] getTelephonyProvidersEdgesTrunksMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunksMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}