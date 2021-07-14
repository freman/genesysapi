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

// GetTelephonyProvidersEdgesTrunkMetricsReader is a Reader for the GetTelephonyProvidersEdgesTrunkMetrics structure.
type GetTelephonyProvidersEdgesTrunkMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesTrunkMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesTrunkMetricsOK creates a GetTelephonyProvidersEdgesTrunkMetricsOK with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsOK() *GetTelephonyProvidersEdgesTrunkMetricsOK {
	return &GetTelephonyProvidersEdgesTrunkMetricsOK{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesTrunkMetricsOK struct {
	Payload *models.TrunkMetrics
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsOK) GetPayload() *models.TrunkMetrics {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrunkMetrics)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsBadRequest creates a GetTelephonyProvidersEdgesTrunkMetricsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsBadRequest() *GetTelephonyProvidersEdgesTrunkMetricsBadRequest {
	return &GetTelephonyProvidersEdgesTrunkMetricsBadRequest{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesTrunkMetricsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsUnauthorized creates a GetTelephonyProvidersEdgesTrunkMetricsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsUnauthorized() *GetTelephonyProvidersEdgesTrunkMetricsUnauthorized {
	return &GetTelephonyProvidersEdgesTrunkMetricsUnauthorized{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesTrunkMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsForbidden creates a GetTelephonyProvidersEdgesTrunkMetricsForbidden with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsForbidden() *GetTelephonyProvidersEdgesTrunkMetricsForbidden {
	return &GetTelephonyProvidersEdgesTrunkMetricsForbidden{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesTrunkMetricsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsNotFound creates a GetTelephonyProvidersEdgesTrunkMetricsNotFound with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsNotFound() *GetTelephonyProvidersEdgesTrunkMetricsNotFound {
	return &GetTelephonyProvidersEdgesTrunkMetricsNotFound{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesTrunkMetricsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsRequestTimeout creates a GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsRequestTimeout() *GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout {
	return &GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge() *GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType creates a GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType() *GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsTooManyRequests creates a GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsTooManyRequests() *GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests {
	return &GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsInternalServerError creates a GetTelephonyProvidersEdgesTrunkMetricsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsInternalServerError() *GetTelephonyProvidersEdgesTrunkMetricsInternalServerError {
	return &GetTelephonyProvidersEdgesTrunkMetricsInternalServerError{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesTrunkMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable creates a GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable() *GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable {
	return &GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout creates a GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout() *GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout {
	return &GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/trunks/{trunkId}/metrics][%d] getTelephonyProvidersEdgesTrunkMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTrunkMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
