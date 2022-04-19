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

// PostGamificationMetricsReader is a Reader for the PostGamificationMetrics structure.
type PostGamificationMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGamificationMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostGamificationMetricsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGamificationMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGamificationMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGamificationMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGamificationMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGamificationMetricsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGamificationMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGamificationMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGamificationMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGamificationMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGamificationMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGamificationMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGamificationMetricsCreated creates a PostGamificationMetricsCreated with default headers values
func NewPostGamificationMetricsCreated() *PostGamificationMetricsCreated {
	return &PostGamificationMetricsCreated{}
}

/*PostGamificationMetricsCreated handles this case with default header values.

Metric successfully created
*/
type PostGamificationMetricsCreated struct {
	Payload *models.Metric
}

func (o *PostGamificationMetricsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsCreated  %+v", 201, o.Payload)
}

func (o *PostGamificationMetricsCreated) GetPayload() *models.Metric {
	return o.Payload
}

func (o *PostGamificationMetricsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsBadRequest creates a PostGamificationMetricsBadRequest with default headers values
func NewPostGamificationMetricsBadRequest() *PostGamificationMetricsBadRequest {
	return &PostGamificationMetricsBadRequest{}
}

/*PostGamificationMetricsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGamificationMetricsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsUnauthorized creates a PostGamificationMetricsUnauthorized with default headers values
func NewPostGamificationMetricsUnauthorized() *PostGamificationMetricsUnauthorized {
	return &PostGamificationMetricsUnauthorized{}
}

/*PostGamificationMetricsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGamificationMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsForbidden creates a PostGamificationMetricsForbidden with default headers values
func NewPostGamificationMetricsForbidden() *PostGamificationMetricsForbidden {
	return &PostGamificationMetricsForbidden{}
}

/*PostGamificationMetricsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGamificationMetricsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsNotFound creates a PostGamificationMetricsNotFound with default headers values
func NewPostGamificationMetricsNotFound() *PostGamificationMetricsNotFound {
	return &PostGamificationMetricsNotFound{}
}

/*PostGamificationMetricsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGamificationMetricsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsRequestTimeout creates a PostGamificationMetricsRequestTimeout with default headers values
func NewPostGamificationMetricsRequestTimeout() *PostGamificationMetricsRequestTimeout {
	return &PostGamificationMetricsRequestTimeout{}
}

/*PostGamificationMetricsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGamificationMetricsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationMetricsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsRequestEntityTooLarge creates a PostGamificationMetricsRequestEntityTooLarge with default headers values
func NewPostGamificationMetricsRequestEntityTooLarge() *PostGamificationMetricsRequestEntityTooLarge {
	return &PostGamificationMetricsRequestEntityTooLarge{}
}

/*PostGamificationMetricsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGamificationMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsUnsupportedMediaType creates a PostGamificationMetricsUnsupportedMediaType with default headers values
func NewPostGamificationMetricsUnsupportedMediaType() *PostGamificationMetricsUnsupportedMediaType {
	return &PostGamificationMetricsUnsupportedMediaType{}
}

/*PostGamificationMetricsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGamificationMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsTooManyRequests creates a PostGamificationMetricsTooManyRequests with default headers values
func NewPostGamificationMetricsTooManyRequests() *PostGamificationMetricsTooManyRequests {
	return &PostGamificationMetricsTooManyRequests{}
}

/*PostGamificationMetricsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGamificationMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsInternalServerError creates a PostGamificationMetricsInternalServerError with default headers values
func NewPostGamificationMetricsInternalServerError() *PostGamificationMetricsInternalServerError {
	return &PostGamificationMetricsInternalServerError{}
}

/*PostGamificationMetricsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGamificationMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsServiceUnavailable creates a PostGamificationMetricsServiceUnavailable with default headers values
func NewPostGamificationMetricsServiceUnavailable() *PostGamificationMetricsServiceUnavailable {
	return &PostGamificationMetricsServiceUnavailable{}
}

/*PostGamificationMetricsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGamificationMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationMetricsGatewayTimeout creates a PostGamificationMetricsGatewayTimeout with default headers values
func NewPostGamificationMetricsGatewayTimeout() *PostGamificationMetricsGatewayTimeout {
	return &PostGamificationMetricsGatewayTimeout{}
}

/*PostGamificationMetricsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGamificationMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/metrics][%d] postGamificationMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
