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

// PostQualityEvaluationsAggregatesQueryMeReader is a Reader for the PostQualityEvaluationsAggregatesQueryMe structure.
type PostQualityEvaluationsAggregatesQueryMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostQualityEvaluationsAggregatesQueryMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostQualityEvaluationsAggregatesQueryMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostQualityEvaluationsAggregatesQueryMeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostQualityEvaluationsAggregatesQueryMeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostQualityEvaluationsAggregatesQueryMeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostQualityEvaluationsAggregatesQueryMeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostQualityEvaluationsAggregatesQueryMeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostQualityEvaluationsAggregatesQueryMeOK creates a PostQualityEvaluationsAggregatesQueryMeOK with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeOK() *PostQualityEvaluationsAggregatesQueryMeOK {
	return &PostQualityEvaluationsAggregatesQueryMeOK{}
}

/*
PostQualityEvaluationsAggregatesQueryMeOK describes a response with status code 200, with default header values.

successful operation
*/
type PostQualityEvaluationsAggregatesQueryMeOK struct {
	Payload *models.EvaluationAggregateQueryResponse
}

// IsSuccess returns true when this post quality evaluations aggregates query me o k response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post quality evaluations aggregates query me o k response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me o k response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations aggregates query me o k response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me o k response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeOK  %+v", 200, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeOK  %+v", 200, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) GetPayload() *models.EvaluationAggregateQueryResponse {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeBadRequest creates a PostQualityEvaluationsAggregatesQueryMeBadRequest with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeBadRequest() *PostQualityEvaluationsAggregatesQueryMeBadRequest {
	return &PostQualityEvaluationsAggregatesQueryMeBadRequest{}
}

/*
PostQualityEvaluationsAggregatesQueryMeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostQualityEvaluationsAggregatesQueryMeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me bad request response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me bad request response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me bad request response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me bad request response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me bad request response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeBadRequest  %+v", 400, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeUnauthorized creates a PostQualityEvaluationsAggregatesQueryMeUnauthorized with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeUnauthorized() *PostQualityEvaluationsAggregatesQueryMeUnauthorized {
	return &PostQualityEvaluationsAggregatesQueryMeUnauthorized{}
}

/*
PostQualityEvaluationsAggregatesQueryMeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostQualityEvaluationsAggregatesQueryMeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me unauthorized response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me unauthorized response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me unauthorized response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me unauthorized response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me unauthorized response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeForbidden creates a PostQualityEvaluationsAggregatesQueryMeForbidden with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeForbidden() *PostQualityEvaluationsAggregatesQueryMeForbidden {
	return &PostQualityEvaluationsAggregatesQueryMeForbidden{}
}

/*
PostQualityEvaluationsAggregatesQueryMeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostQualityEvaluationsAggregatesQueryMeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me forbidden response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me forbidden response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me forbidden response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me forbidden response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me forbidden response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeForbidden  %+v", 403, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeNotFound creates a PostQualityEvaluationsAggregatesQueryMeNotFound with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeNotFound() *PostQualityEvaluationsAggregatesQueryMeNotFound {
	return &PostQualityEvaluationsAggregatesQueryMeNotFound{}
}

/*
PostQualityEvaluationsAggregatesQueryMeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostQualityEvaluationsAggregatesQueryMeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me not found response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me not found response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me not found response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me not found response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me not found response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeNotFound  %+v", 404, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout creates a PostQualityEvaluationsAggregatesQueryMeRequestTimeout with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeRequestTimeout() *PostQualityEvaluationsAggregatesQueryMeRequestTimeout {
	return &PostQualityEvaluationsAggregatesQueryMeRequestTimeout{}
}

/*
PostQualityEvaluationsAggregatesQueryMeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostQualityEvaluationsAggregatesQueryMeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me request timeout response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me request timeout response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me request timeout response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me request timeout response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me request timeout response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge creates a PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge() *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge {
	return &PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge{}
}

/*
PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me request entity too large response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me request entity too large response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me request entity too large response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me request entity too large response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me request entity too large response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType creates a PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType() *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType {
	return &PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType{}
}

/*
PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me unsupported media type response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me unsupported media type response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me unsupported media type response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me unsupported media type response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me unsupported media type response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests creates a PostQualityEvaluationsAggregatesQueryMeTooManyRequests with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeTooManyRequests() *PostQualityEvaluationsAggregatesQueryMeTooManyRequests {
	return &PostQualityEvaluationsAggregatesQueryMeTooManyRequests{}
}

/*
PostQualityEvaluationsAggregatesQueryMeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostQualityEvaluationsAggregatesQueryMeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me too many requests response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me too many requests response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me too many requests response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post quality evaluations aggregates query me too many requests response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post quality evaluations aggregates query me too many requests response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeInternalServerError creates a PostQualityEvaluationsAggregatesQueryMeInternalServerError with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeInternalServerError() *PostQualityEvaluationsAggregatesQueryMeInternalServerError {
	return &PostQualityEvaluationsAggregatesQueryMeInternalServerError{}
}

/*
PostQualityEvaluationsAggregatesQueryMeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostQualityEvaluationsAggregatesQueryMeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me internal server error response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me internal server error response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me internal server error response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations aggregates query me internal server error response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations aggregates query me internal server error response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable creates a PostQualityEvaluationsAggregatesQueryMeServiceUnavailable with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeServiceUnavailable() *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable {
	return &PostQualityEvaluationsAggregatesQueryMeServiceUnavailable{}
}

/*
PostQualityEvaluationsAggregatesQueryMeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostQualityEvaluationsAggregatesQueryMeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me service unavailable response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me service unavailable response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me service unavailable response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations aggregates query me service unavailable response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations aggregates query me service unavailable response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout creates a PostQualityEvaluationsAggregatesQueryMeGatewayTimeout with default headers values
func NewPostQualityEvaluationsAggregatesQueryMeGatewayTimeout() *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout {
	return &PostQualityEvaluationsAggregatesQueryMeGatewayTimeout{}
}

/*
PostQualityEvaluationsAggregatesQueryMeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostQualityEvaluationsAggregatesQueryMeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post quality evaluations aggregates query me gateway timeout response has a 2xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post quality evaluations aggregates query me gateway timeout response has a 3xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post quality evaluations aggregates query me gateway timeout response has a 4xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post quality evaluations aggregates query me gateway timeout response has a 5xx status code
func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post quality evaluations aggregates query me gateway timeout response a status code equal to that given
func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/quality/evaluations/aggregates/query/me][%d] postQualityEvaluationsAggregatesQueryMeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostQualityEvaluationsAggregatesQueryMeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
