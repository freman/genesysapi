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

// PostAnalyticsUsersDetailsQueryReader is a Reader for the PostAnalyticsUsersDetailsQuery structure.
type PostAnalyticsUsersDetailsQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAnalyticsUsersDetailsQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAnalyticsUsersDetailsQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAnalyticsUsersDetailsQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAnalyticsUsersDetailsQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAnalyticsUsersDetailsQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAnalyticsUsersDetailsQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAnalyticsUsersDetailsQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAnalyticsUsersDetailsQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAnalyticsUsersDetailsQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAnalyticsUsersDetailsQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAnalyticsUsersDetailsQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAnalyticsUsersDetailsQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAnalyticsUsersDetailsQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAnalyticsUsersDetailsQueryOK creates a PostAnalyticsUsersDetailsQueryOK with default headers values
func NewPostAnalyticsUsersDetailsQueryOK() *PostAnalyticsUsersDetailsQueryOK {
	return &PostAnalyticsUsersDetailsQueryOK{}
}

/*
PostAnalyticsUsersDetailsQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostAnalyticsUsersDetailsQueryOK struct {
	Payload *models.AnalyticsUserDetailsQueryResponse
}

// IsSuccess returns true when this post analytics users details query o k response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post analytics users details query o k response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query o k response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics users details query o k response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query o k response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostAnalyticsUsersDetailsQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryOK  %+v", 200, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryOK) GetPayload() *models.AnalyticsUserDetailsQueryResponse {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsUserDetailsQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryBadRequest creates a PostAnalyticsUsersDetailsQueryBadRequest with default headers values
func NewPostAnalyticsUsersDetailsQueryBadRequest() *PostAnalyticsUsersDetailsQueryBadRequest {
	return &PostAnalyticsUsersDetailsQueryBadRequest{}
}

/*
PostAnalyticsUsersDetailsQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAnalyticsUsersDetailsQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query bad request response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query bad request response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query bad request response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query bad request response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query bad request response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAnalyticsUsersDetailsQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryUnauthorized creates a PostAnalyticsUsersDetailsQueryUnauthorized with default headers values
func NewPostAnalyticsUsersDetailsQueryUnauthorized() *PostAnalyticsUsersDetailsQueryUnauthorized {
	return &PostAnalyticsUsersDetailsQueryUnauthorized{}
}

/*
PostAnalyticsUsersDetailsQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAnalyticsUsersDetailsQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query unauthorized response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query unauthorized response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query unauthorized response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query unauthorized response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query unauthorized response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAnalyticsUsersDetailsQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryForbidden creates a PostAnalyticsUsersDetailsQueryForbidden with default headers values
func NewPostAnalyticsUsersDetailsQueryForbidden() *PostAnalyticsUsersDetailsQueryForbidden {
	return &PostAnalyticsUsersDetailsQueryForbidden{}
}

/*
PostAnalyticsUsersDetailsQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostAnalyticsUsersDetailsQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query forbidden response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query forbidden response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query forbidden response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query forbidden response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query forbidden response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAnalyticsUsersDetailsQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryNotFound creates a PostAnalyticsUsersDetailsQueryNotFound with default headers values
func NewPostAnalyticsUsersDetailsQueryNotFound() *PostAnalyticsUsersDetailsQueryNotFound {
	return &PostAnalyticsUsersDetailsQueryNotFound{}
}

/*
PostAnalyticsUsersDetailsQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostAnalyticsUsersDetailsQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query not found response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query not found response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query not found response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query not found response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query not found response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAnalyticsUsersDetailsQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryRequestTimeout creates a PostAnalyticsUsersDetailsQueryRequestTimeout with default headers values
func NewPostAnalyticsUsersDetailsQueryRequestTimeout() *PostAnalyticsUsersDetailsQueryRequestTimeout {
	return &PostAnalyticsUsersDetailsQueryRequestTimeout{}
}

/*
PostAnalyticsUsersDetailsQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAnalyticsUsersDetailsQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query request timeout response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query request timeout response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query request timeout response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query request timeout response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query request timeout response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryRequestEntityTooLarge creates a PostAnalyticsUsersDetailsQueryRequestEntityTooLarge with default headers values
func NewPostAnalyticsUsersDetailsQueryRequestEntityTooLarge() *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge {
	return &PostAnalyticsUsersDetailsQueryRequestEntityTooLarge{}
}

/*
PostAnalyticsUsersDetailsQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostAnalyticsUsersDetailsQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query request entity too large response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query request entity too large response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query request entity too large response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query request entity too large response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query request entity too large response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryUnsupportedMediaType creates a PostAnalyticsUsersDetailsQueryUnsupportedMediaType with default headers values
func NewPostAnalyticsUsersDetailsQueryUnsupportedMediaType() *PostAnalyticsUsersDetailsQueryUnsupportedMediaType {
	return &PostAnalyticsUsersDetailsQueryUnsupportedMediaType{}
}

/*
PostAnalyticsUsersDetailsQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAnalyticsUsersDetailsQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query unsupported media type response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query unsupported media type response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query unsupported media type response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query unsupported media type response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query unsupported media type response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryTooManyRequests creates a PostAnalyticsUsersDetailsQueryTooManyRequests with default headers values
func NewPostAnalyticsUsersDetailsQueryTooManyRequests() *PostAnalyticsUsersDetailsQueryTooManyRequests {
	return &PostAnalyticsUsersDetailsQueryTooManyRequests{}
}

/*
PostAnalyticsUsersDetailsQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAnalyticsUsersDetailsQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query too many requests response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query too many requests response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query too many requests response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post analytics users details query too many requests response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post analytics users details query too many requests response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryInternalServerError creates a PostAnalyticsUsersDetailsQueryInternalServerError with default headers values
func NewPostAnalyticsUsersDetailsQueryInternalServerError() *PostAnalyticsUsersDetailsQueryInternalServerError {
	return &PostAnalyticsUsersDetailsQueryInternalServerError{}
}

/*
PostAnalyticsUsersDetailsQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAnalyticsUsersDetailsQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query internal server error response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query internal server error response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query internal server error response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics users details query internal server error response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics users details query internal server error response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAnalyticsUsersDetailsQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryServiceUnavailable creates a PostAnalyticsUsersDetailsQueryServiceUnavailable with default headers values
func NewPostAnalyticsUsersDetailsQueryServiceUnavailable() *PostAnalyticsUsersDetailsQueryServiceUnavailable {
	return &PostAnalyticsUsersDetailsQueryServiceUnavailable{}
}

/*
PostAnalyticsUsersDetailsQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAnalyticsUsersDetailsQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query service unavailable response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query service unavailable response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query service unavailable response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics users details query service unavailable response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics users details query service unavailable response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAnalyticsUsersDetailsQueryGatewayTimeout creates a PostAnalyticsUsersDetailsQueryGatewayTimeout with default headers values
func NewPostAnalyticsUsersDetailsQueryGatewayTimeout() *PostAnalyticsUsersDetailsQueryGatewayTimeout {
	return &PostAnalyticsUsersDetailsQueryGatewayTimeout{}
}

/*
PostAnalyticsUsersDetailsQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostAnalyticsUsersDetailsQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post analytics users details query gateway timeout response has a 2xx status code
func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post analytics users details query gateway timeout response has a 3xx status code
func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post analytics users details query gateway timeout response has a 4xx status code
func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post analytics users details query gateway timeout response has a 5xx status code
func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post analytics users details query gateway timeout response a status code equal to that given
func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/analytics/users/details/query][%d] postAnalyticsUsersDetailsQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAnalyticsUsersDetailsQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
