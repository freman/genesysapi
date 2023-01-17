// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostQualityEvaluationsScoringReader is a Reader for the PostQualityEvaluationsScoring structure.
type PostQualityEvaluationsScoringReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityEvaluationsScoringReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityEvaluationsScoringOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityEvaluationsScoringBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityEvaluationsScoringUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityEvaluationsScoringForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityEvaluationsScoringNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityEvaluationsScoringRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityEvaluationsScoringRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityEvaluationsScoringUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityEvaluationsScoringTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityEvaluationsScoringInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityEvaluationsScoringServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityEvaluationsScoringGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityEvaluationsScoringOK creates a PostQualityEvaluationsScoringOK with default headers values
func NewPostQualityEvaluationsScoringOK() *PostQualityEvaluationsScoringOK {
	return &PostQualityEvaluationsScoringOK{}
}

/*
PostQualityEvaluationsScoringOK describes a response with status code 200, with default header values.

successful operation
*/
type PostQualityEvaluationsScoringOK struct {
	Payload *models.EvaluationScoringSet
}

// IsSuccess returns true when this post quality evaluations scoring o k response has a 2xx status code
func (o *PostQualityEvaluationsScoringOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality evaluations scoring o k response has a 3xx status code
func (o *PostQualityEvaluationsScoringOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring o k response has a 4xx status code
func (o *PostQualityEvaluationsScoringOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations scoring o k response has a 5xx status code
func (o *PostQualityEvaluationsScoringOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring o k response a status code equal to that given
func (o *PostQualityEvaluationsScoringOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQualityEvaluationsScoringOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringOK  %+v", 200, o.Payload)
}

func (o *PostQualityEvaluationsScoringOK) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringOK  %+v", 200, o.Payload)
}

func (o *PostQualityEvaluationsScoringOK) GetPayload() *models.EvaluationScoringSet {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationScoringSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringBadRequest creates a PostQualityEvaluationsScoringBadRequest with default headers values
func NewPostQualityEvaluationsScoringBadRequest() *PostQualityEvaluationsScoringBadRequest {
	return &PostQualityEvaluationsScoringBadRequest{}
}

/*
PostQualityEvaluationsScoringBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityEvaluationsScoringBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring bad request response has a 2xx status code
func (o *PostQualityEvaluationsScoringBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring bad request response has a 3xx status code
func (o *PostQualityEvaluationsScoringBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring bad request response has a 4xx status code
func (o *PostQualityEvaluationsScoringBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring bad request response has a 5xx status code
func (o *PostQualityEvaluationsScoringBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring bad request response a status code equal to that given
func (o *PostQualityEvaluationsScoringBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQualityEvaluationsScoringBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityEvaluationsScoringBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityEvaluationsScoringBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringUnauthorized creates a PostQualityEvaluationsScoringUnauthorized with default headers values
func NewPostQualityEvaluationsScoringUnauthorized() *PostQualityEvaluationsScoringUnauthorized {
	return &PostQualityEvaluationsScoringUnauthorized{}
}

/*
PostQualityEvaluationsScoringUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityEvaluationsScoringUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring unauthorized response has a 2xx status code
func (o *PostQualityEvaluationsScoringUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring unauthorized response has a 3xx status code
func (o *PostQualityEvaluationsScoringUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring unauthorized response has a 4xx status code
func (o *PostQualityEvaluationsScoringUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring unauthorized response has a 5xx status code
func (o *PostQualityEvaluationsScoringUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring unauthorized response a status code equal to that given
func (o *PostQualityEvaluationsScoringUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostQualityEvaluationsScoringUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityEvaluationsScoringUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityEvaluationsScoringUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringForbidden creates a PostQualityEvaluationsScoringForbidden with default headers values
func NewPostQualityEvaluationsScoringForbidden() *PostQualityEvaluationsScoringForbidden {
	return &PostQualityEvaluationsScoringForbidden{}
}

/*
PostQualityEvaluationsScoringForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityEvaluationsScoringForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring forbidden response has a 2xx status code
func (o *PostQualityEvaluationsScoringForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring forbidden response has a 3xx status code
func (o *PostQualityEvaluationsScoringForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring forbidden response has a 4xx status code
func (o *PostQualityEvaluationsScoringForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring forbidden response has a 5xx status code
func (o *PostQualityEvaluationsScoringForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring forbidden response a status code equal to that given
func (o *PostQualityEvaluationsScoringForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostQualityEvaluationsScoringForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityEvaluationsScoringForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityEvaluationsScoringForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringNotFound creates a PostQualityEvaluationsScoringNotFound with default headers values
func NewPostQualityEvaluationsScoringNotFound() *PostQualityEvaluationsScoringNotFound {
	return &PostQualityEvaluationsScoringNotFound{}
}

/*
PostQualityEvaluationsScoringNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostQualityEvaluationsScoringNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring not found response has a 2xx status code
func (o *PostQualityEvaluationsScoringNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring not found response has a 3xx status code
func (o *PostQualityEvaluationsScoringNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring not found response has a 4xx status code
func (o *PostQualityEvaluationsScoringNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring not found response has a 5xx status code
func (o *PostQualityEvaluationsScoringNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring not found response a status code equal to that given
func (o *PostQualityEvaluationsScoringNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostQualityEvaluationsScoringNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityEvaluationsScoringNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityEvaluationsScoringNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringRequestTimeout creates a PostQualityEvaluationsScoringRequestTimeout with default headers values
func NewPostQualityEvaluationsScoringRequestTimeout() *PostQualityEvaluationsScoringRequestTimeout {
	return &PostQualityEvaluationsScoringRequestTimeout{}
}

/*
PostQualityEvaluationsScoringRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityEvaluationsScoringRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring request timeout response has a 2xx status code
func (o *PostQualityEvaluationsScoringRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring request timeout response has a 3xx status code
func (o *PostQualityEvaluationsScoringRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring request timeout response has a 4xx status code
func (o *PostQualityEvaluationsScoringRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring request timeout response has a 5xx status code
func (o *PostQualityEvaluationsScoringRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring request timeout response a status code equal to that given
func (o *PostQualityEvaluationsScoringRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostQualityEvaluationsScoringRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityEvaluationsScoringRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityEvaluationsScoringRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringRequestEntityTooLarge creates a PostQualityEvaluationsScoringRequestEntityTooLarge with default headers values
func NewPostQualityEvaluationsScoringRequestEntityTooLarge() *PostQualityEvaluationsScoringRequestEntityTooLarge {
	return &PostQualityEvaluationsScoringRequestEntityTooLarge{}
}

/*
PostQualityEvaluationsScoringRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostQualityEvaluationsScoringRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring request entity too large response has a 2xx status code
func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring request entity too large response has a 3xx status code
func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring request entity too large response has a 4xx status code
func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring request entity too large response has a 5xx status code
func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring request entity too large response a status code equal to that given
func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringUnsupportedMediaType creates a PostQualityEvaluationsScoringUnsupportedMediaType with default headers values
func NewPostQualityEvaluationsScoringUnsupportedMediaType() *PostQualityEvaluationsScoringUnsupportedMediaType {
	return &PostQualityEvaluationsScoringUnsupportedMediaType{}
}

/*
PostQualityEvaluationsScoringUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityEvaluationsScoringUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring unsupported media type response has a 2xx status code
func (o *PostQualityEvaluationsScoringUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring unsupported media type response has a 3xx status code
func (o *PostQualityEvaluationsScoringUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring unsupported media type response has a 4xx status code
func (o *PostQualityEvaluationsScoringUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring unsupported media type response has a 5xx status code
func (o *PostQualityEvaluationsScoringUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring unsupported media type response a status code equal to that given
func (o *PostQualityEvaluationsScoringUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostQualityEvaluationsScoringUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityEvaluationsScoringUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityEvaluationsScoringUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringTooManyRequests creates a PostQualityEvaluationsScoringTooManyRequests with default headers values
func NewPostQualityEvaluationsScoringTooManyRequests() *PostQualityEvaluationsScoringTooManyRequests {
	return &PostQualityEvaluationsScoringTooManyRequests{}
}

/*
PostQualityEvaluationsScoringTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityEvaluationsScoringTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring too many requests response has a 2xx status code
func (o *PostQualityEvaluationsScoringTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring too many requests response has a 3xx status code
func (o *PostQualityEvaluationsScoringTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring too many requests response has a 4xx status code
func (o *PostQualityEvaluationsScoringTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations scoring too many requests response has a 5xx status code
func (o *PostQualityEvaluationsScoringTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations scoring too many requests response a status code equal to that given
func (o *PostQualityEvaluationsScoringTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostQualityEvaluationsScoringTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityEvaluationsScoringTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityEvaluationsScoringTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringInternalServerError creates a PostQualityEvaluationsScoringInternalServerError with default headers values
func NewPostQualityEvaluationsScoringInternalServerError() *PostQualityEvaluationsScoringInternalServerError {
	return &PostQualityEvaluationsScoringInternalServerError{}
}

/*
PostQualityEvaluationsScoringInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityEvaluationsScoringInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring internal server error response has a 2xx status code
func (o *PostQualityEvaluationsScoringInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring internal server error response has a 3xx status code
func (o *PostQualityEvaluationsScoringInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring internal server error response has a 4xx status code
func (o *PostQualityEvaluationsScoringInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations scoring internal server error response has a 5xx status code
func (o *PostQualityEvaluationsScoringInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations scoring internal server error response a status code equal to that given
func (o *PostQualityEvaluationsScoringInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQualityEvaluationsScoringInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityEvaluationsScoringInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityEvaluationsScoringInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringServiceUnavailable creates a PostQualityEvaluationsScoringServiceUnavailable with default headers values
func NewPostQualityEvaluationsScoringServiceUnavailable() *PostQualityEvaluationsScoringServiceUnavailable {
	return &PostQualityEvaluationsScoringServiceUnavailable{}
}

/*
PostQualityEvaluationsScoringServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityEvaluationsScoringServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring service unavailable response has a 2xx status code
func (o *PostQualityEvaluationsScoringServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring service unavailable response has a 3xx status code
func (o *PostQualityEvaluationsScoringServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring service unavailable response has a 4xx status code
func (o *PostQualityEvaluationsScoringServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations scoring service unavailable response has a 5xx status code
func (o *PostQualityEvaluationsScoringServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations scoring service unavailable response a status code equal to that given
func (o *PostQualityEvaluationsScoringServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostQualityEvaluationsScoringServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityEvaluationsScoringServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityEvaluationsScoringServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsScoringGatewayTimeout creates a PostQualityEvaluationsScoringGatewayTimeout with default headers values
func NewPostQualityEvaluationsScoringGatewayTimeout() *PostQualityEvaluationsScoringGatewayTimeout {
	return &PostQualityEvaluationsScoringGatewayTimeout{}
}

/*
PostQualityEvaluationsScoringGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostQualityEvaluationsScoringGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations scoring gateway timeout response has a 2xx status code
func (o *PostQualityEvaluationsScoringGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations scoring gateway timeout response has a 3xx status code
func (o *PostQualityEvaluationsScoringGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations scoring gateway timeout response has a 4xx status code
func (o *PostQualityEvaluationsScoringGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations scoring gateway timeout response has a 5xx status code
func (o *PostQualityEvaluationsScoringGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations scoring gateway timeout response a status code equal to that given
func (o *PostQualityEvaluationsScoringGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostQualityEvaluationsScoringGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityEvaluationsScoringGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/scoring][%d] postQualityEvaluationsScoringGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityEvaluationsScoringGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsScoringGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
