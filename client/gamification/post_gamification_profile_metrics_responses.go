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

// PostGamificationProfileMetricsReader is a Reader for the PostGamificationProfileMetrics structure.
type PostGamificationProfileMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGamificationProfileMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostGamificationProfileMetricsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGamificationProfileMetricsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGamificationProfileMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGamificationProfileMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGamificationProfileMetricsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGamificationProfileMetricsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGamificationProfileMetricsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGamificationProfileMetricsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGamificationProfileMetricsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGamificationProfileMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGamificationProfileMetricsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGamificationProfileMetricsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGamificationProfileMetricsCreated creates a PostGamificationProfileMetricsCreated with default headers values
func NewPostGamificationProfileMetricsCreated() *PostGamificationProfileMetricsCreated {
	return &PostGamificationProfileMetricsCreated{}
}

/*
PostGamificationProfileMetricsCreated describes a response with status code 201, with default header values.

Metric successfully created
*/
type PostGamificationProfileMetricsCreated struct {
	Payload *models.Metric
}

// IsSuccess returns true when this post gamification profile metrics created response has a 2xx status code
func (o *PostGamificationProfileMetricsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post gamification profile metrics created response has a 3xx status code
func (o *PostGamificationProfileMetricsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics created response has a 4xx status code
func (o *PostGamificationProfileMetricsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile metrics created response has a 5xx status code
func (o *PostGamificationProfileMetricsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics created response a status code equal to that given
func (o *PostGamificationProfileMetricsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PostGamificationProfileMetricsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsCreated  %+v", 201, o.Payload)
}

func (o *PostGamificationProfileMetricsCreated) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsCreated  %+v", 201, o.Payload)
}

func (o *PostGamificationProfileMetricsCreated) GetPayload() *models.Metric {
	return o.Payload
}

func (o *PostGamificationProfileMetricsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Metric)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsBadRequest creates a PostGamificationProfileMetricsBadRequest with default headers values
func NewPostGamificationProfileMetricsBadRequest() *PostGamificationProfileMetricsBadRequest {
	return &PostGamificationProfileMetricsBadRequest{}
}

/*
PostGamificationProfileMetricsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGamificationProfileMetricsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics bad request response has a 2xx status code
func (o *PostGamificationProfileMetricsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics bad request response has a 3xx status code
func (o *PostGamificationProfileMetricsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics bad request response has a 4xx status code
func (o *PostGamificationProfileMetricsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics bad request response has a 5xx status code
func (o *PostGamificationProfileMetricsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics bad request response a status code equal to that given
func (o *PostGamificationProfileMetricsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGamificationProfileMetricsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileMetricsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileMetricsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsUnauthorized creates a PostGamificationProfileMetricsUnauthorized with default headers values
func NewPostGamificationProfileMetricsUnauthorized() *PostGamificationProfileMetricsUnauthorized {
	return &PostGamificationProfileMetricsUnauthorized{}
}

/*
PostGamificationProfileMetricsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGamificationProfileMetricsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics unauthorized response has a 2xx status code
func (o *PostGamificationProfileMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics unauthorized response has a 3xx status code
func (o *PostGamificationProfileMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics unauthorized response has a 4xx status code
func (o *PostGamificationProfileMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics unauthorized response has a 5xx status code
func (o *PostGamificationProfileMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics unauthorized response a status code equal to that given
func (o *PostGamificationProfileMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGamificationProfileMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileMetricsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileMetricsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsForbidden creates a PostGamificationProfileMetricsForbidden with default headers values
func NewPostGamificationProfileMetricsForbidden() *PostGamificationProfileMetricsForbidden {
	return &PostGamificationProfileMetricsForbidden{}
}

/*
PostGamificationProfileMetricsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGamificationProfileMetricsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics forbidden response has a 2xx status code
func (o *PostGamificationProfileMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics forbidden response has a 3xx status code
func (o *PostGamificationProfileMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics forbidden response has a 4xx status code
func (o *PostGamificationProfileMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics forbidden response has a 5xx status code
func (o *PostGamificationProfileMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics forbidden response a status code equal to that given
func (o *PostGamificationProfileMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGamificationProfileMetricsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileMetricsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileMetricsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsNotFound creates a PostGamificationProfileMetricsNotFound with default headers values
func NewPostGamificationProfileMetricsNotFound() *PostGamificationProfileMetricsNotFound {
	return &PostGamificationProfileMetricsNotFound{}
}

/*
PostGamificationProfileMetricsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGamificationProfileMetricsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics not found response has a 2xx status code
func (o *PostGamificationProfileMetricsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics not found response has a 3xx status code
func (o *PostGamificationProfileMetricsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics not found response has a 4xx status code
func (o *PostGamificationProfileMetricsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics not found response has a 5xx status code
func (o *PostGamificationProfileMetricsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics not found response a status code equal to that given
func (o *PostGamificationProfileMetricsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGamificationProfileMetricsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileMetricsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileMetricsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsRequestTimeout creates a PostGamificationProfileMetricsRequestTimeout with default headers values
func NewPostGamificationProfileMetricsRequestTimeout() *PostGamificationProfileMetricsRequestTimeout {
	return &PostGamificationProfileMetricsRequestTimeout{}
}

/*
PostGamificationProfileMetricsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGamificationProfileMetricsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics request timeout response has a 2xx status code
func (o *PostGamificationProfileMetricsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics request timeout response has a 3xx status code
func (o *PostGamificationProfileMetricsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics request timeout response has a 4xx status code
func (o *PostGamificationProfileMetricsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics request timeout response has a 5xx status code
func (o *PostGamificationProfileMetricsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics request timeout response a status code equal to that given
func (o *PostGamificationProfileMetricsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGamificationProfileMetricsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileMetricsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileMetricsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsRequestEntityTooLarge creates a PostGamificationProfileMetricsRequestEntityTooLarge with default headers values
func NewPostGamificationProfileMetricsRequestEntityTooLarge() *PostGamificationProfileMetricsRequestEntityTooLarge {
	return &PostGamificationProfileMetricsRequestEntityTooLarge{}
}

/*
PostGamificationProfileMetricsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostGamificationProfileMetricsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics request entity too large response has a 2xx status code
func (o *PostGamificationProfileMetricsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics request entity too large response has a 3xx status code
func (o *PostGamificationProfileMetricsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics request entity too large response has a 4xx status code
func (o *PostGamificationProfileMetricsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics request entity too large response has a 5xx status code
func (o *PostGamificationProfileMetricsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics request entity too large response a status code equal to that given
func (o *PostGamificationProfileMetricsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGamificationProfileMetricsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileMetricsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileMetricsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsUnsupportedMediaType creates a PostGamificationProfileMetricsUnsupportedMediaType with default headers values
func NewPostGamificationProfileMetricsUnsupportedMediaType() *PostGamificationProfileMetricsUnsupportedMediaType {
	return &PostGamificationProfileMetricsUnsupportedMediaType{}
}

/*
PostGamificationProfileMetricsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGamificationProfileMetricsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics unsupported media type response has a 2xx status code
func (o *PostGamificationProfileMetricsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics unsupported media type response has a 3xx status code
func (o *PostGamificationProfileMetricsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics unsupported media type response has a 4xx status code
func (o *PostGamificationProfileMetricsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics unsupported media type response has a 5xx status code
func (o *PostGamificationProfileMetricsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics unsupported media type response a status code equal to that given
func (o *PostGamificationProfileMetricsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGamificationProfileMetricsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileMetricsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileMetricsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsTooManyRequests creates a PostGamificationProfileMetricsTooManyRequests with default headers values
func NewPostGamificationProfileMetricsTooManyRequests() *PostGamificationProfileMetricsTooManyRequests {
	return &PostGamificationProfileMetricsTooManyRequests{}
}

/*
PostGamificationProfileMetricsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGamificationProfileMetricsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics too many requests response has a 2xx status code
func (o *PostGamificationProfileMetricsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics too many requests response has a 3xx status code
func (o *PostGamificationProfileMetricsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics too many requests response has a 4xx status code
func (o *PostGamificationProfileMetricsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile metrics too many requests response has a 5xx status code
func (o *PostGamificationProfileMetricsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile metrics too many requests response a status code equal to that given
func (o *PostGamificationProfileMetricsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGamificationProfileMetricsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileMetricsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileMetricsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsInternalServerError creates a PostGamificationProfileMetricsInternalServerError with default headers values
func NewPostGamificationProfileMetricsInternalServerError() *PostGamificationProfileMetricsInternalServerError {
	return &PostGamificationProfileMetricsInternalServerError{}
}

/*
PostGamificationProfileMetricsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGamificationProfileMetricsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics internal server error response has a 2xx status code
func (o *PostGamificationProfileMetricsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics internal server error response has a 3xx status code
func (o *PostGamificationProfileMetricsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics internal server error response has a 4xx status code
func (o *PostGamificationProfileMetricsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile metrics internal server error response has a 5xx status code
func (o *PostGamificationProfileMetricsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile metrics internal server error response a status code equal to that given
func (o *PostGamificationProfileMetricsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGamificationProfileMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileMetricsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileMetricsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsServiceUnavailable creates a PostGamificationProfileMetricsServiceUnavailable with default headers values
func NewPostGamificationProfileMetricsServiceUnavailable() *PostGamificationProfileMetricsServiceUnavailable {
	return &PostGamificationProfileMetricsServiceUnavailable{}
}

/*
PostGamificationProfileMetricsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGamificationProfileMetricsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics service unavailable response has a 2xx status code
func (o *PostGamificationProfileMetricsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics service unavailable response has a 3xx status code
func (o *PostGamificationProfileMetricsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics service unavailable response has a 4xx status code
func (o *PostGamificationProfileMetricsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile metrics service unavailable response has a 5xx status code
func (o *PostGamificationProfileMetricsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile metrics service unavailable response a status code equal to that given
func (o *PostGamificationProfileMetricsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGamificationProfileMetricsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileMetricsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileMetricsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMetricsGatewayTimeout creates a PostGamificationProfileMetricsGatewayTimeout with default headers values
func NewPostGamificationProfileMetricsGatewayTimeout() *PostGamificationProfileMetricsGatewayTimeout {
	return &PostGamificationProfileMetricsGatewayTimeout{}
}

/*
PostGamificationProfileMetricsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGamificationProfileMetricsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile metrics gateway timeout response has a 2xx status code
func (o *PostGamificationProfileMetricsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile metrics gateway timeout response has a 3xx status code
func (o *PostGamificationProfileMetricsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile metrics gateway timeout response has a 4xx status code
func (o *PostGamificationProfileMetricsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile metrics gateway timeout response has a 5xx status code
func (o *PostGamificationProfileMetricsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile metrics gateway timeout response a status code equal to that given
func (o *PostGamificationProfileMetricsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGamificationProfileMetricsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileMetricsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/metrics][%d] postGamificationProfileMetricsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileMetricsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMetricsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
