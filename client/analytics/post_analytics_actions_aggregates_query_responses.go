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

// PostAnalyticsActionsAggregatesQueryReader is a Reader for the PostAnalyticsActionsAggregatesQuery structure.
type PostAnalyticsActionsAggregatesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsActionsAggregatesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsActionsAggregatesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsActionsAggregatesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsActionsAggregatesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsActionsAggregatesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsActionsAggregatesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAnalyticsActionsAggregatesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsActionsAggregatesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsActionsAggregatesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsActionsAggregatesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsActionsAggregatesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsActionsAggregatesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsActionsAggregatesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsActionsAggregatesQueryOK creates a PostAnalyticsActionsAggregatesQueryOK with default headers values
func NewPostAnalyticsActionsAggregatesQueryOK() *PostAnalyticsActionsAggregatesQueryOK {
	return &PostAnalyticsActionsAggregatesQueryOK{}
}

/*
PostAnalyticsActionsAggregatesQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostAnalyticsActionsAggregatesQueryOK struct {
	Payload *models.ActionAggregateQueryResponse
}

// IsSuccess returns true when this post analytics actions aggregates query o k response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post analytics actions aggregates query o k response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query o k response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics actions aggregates query o k response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query o k response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAnalyticsActionsAggregatesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryOK) GetPayload() *models.ActionAggregateQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionAggregateQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryBadRequest creates a PostAnalyticsActionsAggregatesQueryBadRequest with default headers values
func NewPostAnalyticsActionsAggregatesQueryBadRequest() *PostAnalyticsActionsAggregatesQueryBadRequest {
	return &PostAnalyticsActionsAggregatesQueryBadRequest{}
}

/*
PostAnalyticsActionsAggregatesQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsActionsAggregatesQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query bad request response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query bad request response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query bad request response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query bad request response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query bad request response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAnalyticsActionsAggregatesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryUnauthorized creates a PostAnalyticsActionsAggregatesQueryUnauthorized with default headers values
func NewPostAnalyticsActionsAggregatesQueryUnauthorized() *PostAnalyticsActionsAggregatesQueryUnauthorized {
	return &PostAnalyticsActionsAggregatesQueryUnauthorized{}
}

/*
PostAnalyticsActionsAggregatesQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsActionsAggregatesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query unauthorized response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query unauthorized response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query unauthorized response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query unauthorized response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query unauthorized response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryForbidden creates a PostAnalyticsActionsAggregatesQueryForbidden with default headers values
func NewPostAnalyticsActionsAggregatesQueryForbidden() *PostAnalyticsActionsAggregatesQueryForbidden {
	return &PostAnalyticsActionsAggregatesQueryForbidden{}
}

/*
PostAnalyticsActionsAggregatesQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsActionsAggregatesQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query forbidden response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query forbidden response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query forbidden response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query forbidden response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query forbidden response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAnalyticsActionsAggregatesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryNotFound creates a PostAnalyticsActionsAggregatesQueryNotFound with default headers values
func NewPostAnalyticsActionsAggregatesQueryNotFound() *PostAnalyticsActionsAggregatesQueryNotFound {
	return &PostAnalyticsActionsAggregatesQueryNotFound{}
}

/*
PostAnalyticsActionsAggregatesQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostAnalyticsActionsAggregatesQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query not found response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query not found response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query not found response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query not found response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query not found response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAnalyticsActionsAggregatesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryRequestTimeout creates a PostAnalyticsActionsAggregatesQueryRequestTimeout with default headers values
func NewPostAnalyticsActionsAggregatesQueryRequestTimeout() *PostAnalyticsActionsAggregatesQueryRequestTimeout {
	return &PostAnalyticsActionsAggregatesQueryRequestTimeout{}
}

/*
PostAnalyticsActionsAggregatesQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAnalyticsActionsAggregatesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query request timeout response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query request timeout response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query request timeout response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query request timeout response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query request timeout response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryRequestEntityTooLarge creates a PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsActionsAggregatesQueryRequestEntityTooLarge() *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge {
	return &PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge{}
}

/*
PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query request entity too large response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query request entity too large response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query request entity too large response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query request entity too large response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query request entity too large response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryUnsupportedMediaType creates a PostAnalyticsActionsAggregatesQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsActionsAggregatesQueryUnsupportedMediaType() *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType {
	return &PostAnalyticsActionsAggregatesQueryUnsupportedMediaType{}
}

/*
PostAnalyticsActionsAggregatesQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsActionsAggregatesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query unsupported media type response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query unsupported media type response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query unsupported media type response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query unsupported media type response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query unsupported media type response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryTooManyRequests creates a PostAnalyticsActionsAggregatesQueryTooManyRequests with default headers values
func NewPostAnalyticsActionsAggregatesQueryTooManyRequests() *PostAnalyticsActionsAggregatesQueryTooManyRequests {
	return &PostAnalyticsActionsAggregatesQueryTooManyRequests{}
}

/*
PostAnalyticsActionsAggregatesQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAnalyticsActionsAggregatesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query too many requests response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query too many requests response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query too many requests response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics actions aggregates query too many requests response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics actions aggregates query too many requests response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryInternalServerError creates a PostAnalyticsActionsAggregatesQueryInternalServerError with default headers values
func NewPostAnalyticsActionsAggregatesQueryInternalServerError() *PostAnalyticsActionsAggregatesQueryInternalServerError {
	return &PostAnalyticsActionsAggregatesQueryInternalServerError{}
}

/*
PostAnalyticsActionsAggregatesQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsActionsAggregatesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query internal server error response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query internal server error response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query internal server error response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics actions aggregates query internal server error response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics actions aggregates query internal server error response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryServiceUnavailable creates a PostAnalyticsActionsAggregatesQueryServiceUnavailable with default headers values
func NewPostAnalyticsActionsAggregatesQueryServiceUnavailable() *PostAnalyticsActionsAggregatesQueryServiceUnavailable {
	return &PostAnalyticsActionsAggregatesQueryServiceUnavailable{}
}

/*
PostAnalyticsActionsAggregatesQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsActionsAggregatesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query service unavailable response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query service unavailable response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query service unavailable response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics actions aggregates query service unavailable response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics actions aggregates query service unavailable response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsActionsAggregatesQueryGatewayTimeout creates a PostAnalyticsActionsAggregatesQueryGatewayTimeout with default headers values
func NewPostAnalyticsActionsAggregatesQueryGatewayTimeout() *PostAnalyticsActionsAggregatesQueryGatewayTimeout {
	return &PostAnalyticsActionsAggregatesQueryGatewayTimeout{}
}

/*
PostAnalyticsActionsAggregatesQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostAnalyticsActionsAggregatesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics actions aggregates query gateway timeout response has a 2xx status code
func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics actions aggregates query gateway timeout response has a 3xx status code
func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics actions aggregates query gateway timeout response has a 4xx status code
func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics actions aggregates query gateway timeout response has a 5xx status code
func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics actions aggregates query gateway timeout response a status code equal to that given
func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/actions/aggregates/query][%d] postAnalyticsActionsAggregatesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsActionsAggregatesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
