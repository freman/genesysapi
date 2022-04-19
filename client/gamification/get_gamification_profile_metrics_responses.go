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

// GetGamificationProfileMetricsReader is a Reader for the GetGamificationProfileMetrics structure.
type GetGamificationProfileMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationProfileMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationProfileMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationProfileMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationProfileMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationProfileMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationProfileMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationProfileMetricsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationProfileMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationProfileMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationProfileMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationProfileMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationProfileMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationProfileMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationProfileMetricsOK creates a GetGamificationProfileMetricsOK with default headers values
func NewGetGamificationProfileMetricsOK() *GetGamificationProfileMetricsOK {
	return &GetGamificationProfileMetricsOK{}
}

/*GetGamificationProfileMetricsOK handles this case with default header values.

successful operation
*/
type GetGamificationProfileMetricsOK struct {
	Payload *models.GetMetricResponse
}

func (o *GetGamificationProfileMetricsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsOK  %+v", 200, o.Payload)
}

func (o *GetGamificationProfileMetricsOK) GetPayload() *models.GetMetricResponse {
	return o.Payload
}

func (o *GetGamificationProfileMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsBadRequest creates a GetGamificationProfileMetricsBadRequest with default headers values
func NewGetGamificationProfileMetricsBadRequest() *GetGamificationProfileMetricsBadRequest {
	return &GetGamificationProfileMetricsBadRequest{}
}

/*GetGamificationProfileMetricsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationProfileMetricsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationProfileMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsUnauthorized creates a GetGamificationProfileMetricsUnauthorized with default headers values
func NewGetGamificationProfileMetricsUnauthorized() *GetGamificationProfileMetricsUnauthorized {
	return &GetGamificationProfileMetricsUnauthorized{}
}

/*GetGamificationProfileMetricsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationProfileMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationProfileMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsForbidden creates a GetGamificationProfileMetricsForbidden with default headers values
func NewGetGamificationProfileMetricsForbidden() *GetGamificationProfileMetricsForbidden {
	return &GetGamificationProfileMetricsForbidden{}
}

/*GetGamificationProfileMetricsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationProfileMetricsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationProfileMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsNotFound creates a GetGamificationProfileMetricsNotFound with default headers values
func NewGetGamificationProfileMetricsNotFound() *GetGamificationProfileMetricsNotFound {
	return &GetGamificationProfileMetricsNotFound{}
}

/*GetGamificationProfileMetricsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationProfileMetricsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationProfileMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsRequestTimeout creates a GetGamificationProfileMetricsRequestTimeout with default headers values
func NewGetGamificationProfileMetricsRequestTimeout() *GetGamificationProfileMetricsRequestTimeout {
	return &GetGamificationProfileMetricsRequestTimeout{}
}

/*GetGamificationProfileMetricsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationProfileMetricsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationProfileMetricsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsRequestEntityTooLarge creates a GetGamificationProfileMetricsRequestEntityTooLarge with default headers values
func NewGetGamificationProfileMetricsRequestEntityTooLarge() *GetGamificationProfileMetricsRequestEntityTooLarge {
	return &GetGamificationProfileMetricsRequestEntityTooLarge{}
}

/*GetGamificationProfileMetricsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetGamificationProfileMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationProfileMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsUnsupportedMediaType creates a GetGamificationProfileMetricsUnsupportedMediaType with default headers values
func NewGetGamificationProfileMetricsUnsupportedMediaType() *GetGamificationProfileMetricsUnsupportedMediaType {
	return &GetGamificationProfileMetricsUnsupportedMediaType{}
}

/*GetGamificationProfileMetricsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationProfileMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationProfileMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsTooManyRequests creates a GetGamificationProfileMetricsTooManyRequests with default headers values
func NewGetGamificationProfileMetricsTooManyRequests() *GetGamificationProfileMetricsTooManyRequests {
	return &GetGamificationProfileMetricsTooManyRequests{}
}

/*GetGamificationProfileMetricsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationProfileMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationProfileMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsInternalServerError creates a GetGamificationProfileMetricsInternalServerError with default headers values
func NewGetGamificationProfileMetricsInternalServerError() *GetGamificationProfileMetricsInternalServerError {
	return &GetGamificationProfileMetricsInternalServerError{}
}

/*GetGamificationProfileMetricsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationProfileMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationProfileMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsServiceUnavailable creates a GetGamificationProfileMetricsServiceUnavailable with default headers values
func NewGetGamificationProfileMetricsServiceUnavailable() *GetGamificationProfileMetricsServiceUnavailable {
	return &GetGamificationProfileMetricsServiceUnavailable{}
}

/*GetGamificationProfileMetricsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationProfileMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationProfileMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfileMetricsGatewayTimeout creates a GetGamificationProfileMetricsGatewayTimeout with default headers values
func NewGetGamificationProfileMetricsGatewayTimeout() *GetGamificationProfileMetricsGatewayTimeout {
	return &GetGamificationProfileMetricsGatewayTimeout{}
}

/*GetGamificationProfileMetricsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationProfileMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfileMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles/{profileId}/metrics][%d] getGamificationProfileMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationProfileMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfileMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
