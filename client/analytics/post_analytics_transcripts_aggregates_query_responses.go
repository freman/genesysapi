// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAnalyticsTranscriptsAggregatesQueryReader is a Reader for the PostAnalyticsTranscriptsAggregatesQuery structure.
type PostAnalyticsTranscriptsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsTranscriptsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsTranscriptsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsTranscriptsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsTranscriptsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsTranscriptsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsTranscriptsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAnalyticsTranscriptsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsTranscriptsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsTranscriptsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsTranscriptsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsTranscriptsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsTranscriptsAggregatesQueryOK creates a PostAnalyticsTranscriptsAggregatesQueryOK with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryOK() *PostAnalyticsTranscriptsAggregatesQueryOK {
	return &PostAnalyticsTranscriptsAggregatesQueryOK{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostAnalyticsTranscriptsAggregatesQueryOK struct {
	Payload *models.TranscriptAggregateQueryResponse
}

// IsSuccess returns true when this post analytics transcripts aggregates query o k response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post analytics transcripts aggregates query o k response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query o k response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics transcripts aggregates query o k response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query o k response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAnalyticsTranscriptsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryOK) GetPayload() *models.TranscriptAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranscriptAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryBadRequest creates a PostAnalyticsTranscriptsAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryBadRequest() *PostAnalyticsTranscriptsAggregatesQueryBadRequest {
	return &PostAnalyticsTranscriptsAggregatesQueryBadRequest{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsTranscriptsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query bad request response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query bad request response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query bad request response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query bad request response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query bad request response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryUnauthorized creates a PostAnalyticsTranscriptsAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryUnauthorized() *PostAnalyticsTranscriptsAggregatesQueryUnauthorized {
	return &PostAnalyticsTranscriptsAggregatesQueryUnauthorized{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsTranscriptsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query unauthorized response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query unauthorized response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query unauthorized response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query unauthorized response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query unauthorized response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryForbidden creates a PostAnalyticsTranscriptsAggregatesQueryForbidden with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryForbidden() *PostAnalyticsTranscriptsAggregatesQueryForbidden {
	return &PostAnalyticsTranscriptsAggregatesQueryForbidden{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsTranscriptsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query forbidden response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query forbidden response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query forbidden response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query forbidden response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query forbidden response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryNotFound creates a PostAnalyticsTranscriptsAggregatesQueryNotFound with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryNotFound() *PostAnalyticsTranscriptsAggregatesQueryNotFound {
	return &PostAnalyticsTranscriptsAggregatesQueryNotFound{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostAnalyticsTranscriptsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query not found response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query not found response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query not found response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query not found response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query not found response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryRequestTimeout creates a PostAnalyticsTranscriptsAggregatesQueryRequestTimeout with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryRequestTimeout() *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout {
	return &PostAnalyticsTranscriptsAggregatesQueryRequestTimeout{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAnalyticsTranscriptsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query request timeout response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query request timeout response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query request timeout response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query request timeout response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query request timeout response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge() *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query request entity too large response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query request entity too large response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query request entity too large response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query request entity too large response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query request entity too large response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType creates a PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType() *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query unsupported media type response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query unsupported media type response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query unsupported media type response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query unsupported media type response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query unsupported media type response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryTooManyRequests creates a PostAnalyticsTranscriptsAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryTooManyRequests() *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests {
	return &PostAnalyticsTranscriptsAggregatesQueryTooManyRequests{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAnalyticsTranscriptsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query too many requests response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query too many requests response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query too many requests response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics transcripts aggregates query too many requests response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics transcripts aggregates query too many requests response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryInternalServerError creates a PostAnalyticsTranscriptsAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryInternalServerError() *PostAnalyticsTranscriptsAggregatesQueryInternalServerError {
	return &PostAnalyticsTranscriptsAggregatesQueryInternalServerError{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsTranscriptsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query internal server error response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query internal server error response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query internal server error response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics transcripts aggregates query internal server error response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics transcripts aggregates query internal server error response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryServiceUnavailable creates a PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryServiceUnavailable() *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable {
	return &PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query service unavailable response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query service unavailable response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query service unavailable response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics transcripts aggregates query service unavailable response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics transcripts aggregates query service unavailable response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsTranscriptsAggregatesQueryGatewayTimeout creates a PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsTranscriptsAggregatesQueryGatewayTimeout() *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout {
	return &PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout{}
}

/*
PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics transcripts aggregates query gateway timeout response has a 2xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics transcripts aggregates query gateway timeout response has a 3xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics transcripts aggregates query gateway timeout response has a 4xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics transcripts aggregates query gateway timeout response has a 5xx status code
func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics transcripts aggregates query gateway timeout response a status code equal to that given
func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/transcripts/aggregates/query][%d] postAnalyticsTranscriptsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsTranscriptsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
