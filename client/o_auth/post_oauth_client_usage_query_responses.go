// Code generated by go-swagger; DO NOT EDIT.

package o_auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostOauthClientUsageQueryReader is a Reader for the PostOauthClientUsageQuery structure.
type PostOauthClientUsageQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOauthClientUsageQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOauthClientUsageQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPostOauthClientUsageQueryAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOauthClientUsageQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOauthClientUsageQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOauthClientUsageQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOauthClientUsageQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOauthClientUsageQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOauthClientUsageQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOauthClientUsageQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOauthClientUsageQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOauthClientUsageQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOauthClientUsageQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOauthClientUsageQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOauthClientUsageQueryOK creates a PostOauthClientUsageQueryOK with default headers values
func NewPostOauthClientUsageQueryOK() *PostOauthClientUsageQueryOK {
	return &PostOauthClientUsageQueryOK{}
}

/*
PostOauthClientUsageQueryOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOauthClientUsageQueryOK struct {
	Payload *models.UsageExecutionResult
}

// IsSuccess returns true when this post oauth client usage query o k response has a 2xx status code
func (o *PostOauthClientUsageQueryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post oauth client usage query o k response has a 3xx status code
func (o *PostOauthClientUsageQueryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query o k response has a 4xx status code
func (o *PostOauthClientUsageQueryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth client usage query o k response has a 5xx status code
func (o *PostOauthClientUsageQueryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query o k response a status code equal to that given
func (o *PostOauthClientUsageQueryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOauthClientUsageQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryOK  %+v", 200, o.Payload)
}

func (o *PostOauthClientUsageQueryOK) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryOK  %+v", 200, o.Payload)
}

func (o *PostOauthClientUsageQueryOK) GetPayload() *models.UsageExecutionResult {
	return o.Payload
}

func (o *PostOauthClientUsageQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsageExecutionResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryAccepted creates a PostOauthClientUsageQueryAccepted with default headers values
func NewPostOauthClientUsageQueryAccepted() *PostOauthClientUsageQueryAccepted {
	return &PostOauthClientUsageQueryAccepted{}
}

/*
PostOauthClientUsageQueryAccepted describes a response with status code 202, with default header values.

Execution not completed, check back for results
*/
type PostOauthClientUsageQueryAccepted struct {
	Payload *models.UsageExecutionResult
}

// IsSuccess returns true when this post oauth client usage query accepted response has a 2xx status code
func (o *PostOauthClientUsageQueryAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post oauth client usage query accepted response has a 3xx status code
func (o *PostOauthClientUsageQueryAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query accepted response has a 4xx status code
func (o *PostOauthClientUsageQueryAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth client usage query accepted response has a 5xx status code
func (o *PostOauthClientUsageQueryAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query accepted response a status code equal to that given
func (o *PostOauthClientUsageQueryAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *PostOauthClientUsageQueryAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryAccepted  %+v", 202, o.Payload)
}

func (o *PostOauthClientUsageQueryAccepted) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryAccepted  %+v", 202, o.Payload)
}

func (o *PostOauthClientUsageQueryAccepted) GetPayload() *models.UsageExecutionResult {
	return o.Payload
}

func (o *PostOauthClientUsageQueryAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsageExecutionResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryBadRequest creates a PostOauthClientUsageQueryBadRequest with default headers values
func NewPostOauthClientUsageQueryBadRequest() *PostOauthClientUsageQueryBadRequest {
	return &PostOauthClientUsageQueryBadRequest{}
}

/*
PostOauthClientUsageQueryBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOauthClientUsageQueryBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query bad request response has a 2xx status code
func (o *PostOauthClientUsageQueryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query bad request response has a 3xx status code
func (o *PostOauthClientUsageQueryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query bad request response has a 4xx status code
func (o *PostOauthClientUsageQueryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query bad request response has a 5xx status code
func (o *PostOauthClientUsageQueryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query bad request response a status code equal to that given
func (o *PostOauthClientUsageQueryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOauthClientUsageQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostOauthClientUsageQueryBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostOauthClientUsageQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryUnauthorized creates a PostOauthClientUsageQueryUnauthorized with default headers values
func NewPostOauthClientUsageQueryUnauthorized() *PostOauthClientUsageQueryUnauthorized {
	return &PostOauthClientUsageQueryUnauthorized{}
}

/*
PostOauthClientUsageQueryUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOauthClientUsageQueryUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query unauthorized response has a 2xx status code
func (o *PostOauthClientUsageQueryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query unauthorized response has a 3xx status code
func (o *PostOauthClientUsageQueryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query unauthorized response has a 4xx status code
func (o *PostOauthClientUsageQueryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query unauthorized response has a 5xx status code
func (o *PostOauthClientUsageQueryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query unauthorized response a status code equal to that given
func (o *PostOauthClientUsageQueryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOauthClientUsageQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOauthClientUsageQueryUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOauthClientUsageQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryForbidden creates a PostOauthClientUsageQueryForbidden with default headers values
func NewPostOauthClientUsageQueryForbidden() *PostOauthClientUsageQueryForbidden {
	return &PostOauthClientUsageQueryForbidden{}
}

/*
PostOauthClientUsageQueryForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOauthClientUsageQueryForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query forbidden response has a 2xx status code
func (o *PostOauthClientUsageQueryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query forbidden response has a 3xx status code
func (o *PostOauthClientUsageQueryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query forbidden response has a 4xx status code
func (o *PostOauthClientUsageQueryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query forbidden response has a 5xx status code
func (o *PostOauthClientUsageQueryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query forbidden response a status code equal to that given
func (o *PostOauthClientUsageQueryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOauthClientUsageQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostOauthClientUsageQueryForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostOauthClientUsageQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryNotFound creates a PostOauthClientUsageQueryNotFound with default headers values
func NewPostOauthClientUsageQueryNotFound() *PostOauthClientUsageQueryNotFound {
	return &PostOauthClientUsageQueryNotFound{}
}

/*
PostOauthClientUsageQueryNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOauthClientUsageQueryNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query not found response has a 2xx status code
func (o *PostOauthClientUsageQueryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query not found response has a 3xx status code
func (o *PostOauthClientUsageQueryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query not found response has a 4xx status code
func (o *PostOauthClientUsageQueryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query not found response has a 5xx status code
func (o *PostOauthClientUsageQueryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query not found response a status code equal to that given
func (o *PostOauthClientUsageQueryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOauthClientUsageQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostOauthClientUsageQueryNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostOauthClientUsageQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryRequestTimeout creates a PostOauthClientUsageQueryRequestTimeout with default headers values
func NewPostOauthClientUsageQueryRequestTimeout() *PostOauthClientUsageQueryRequestTimeout {
	return &PostOauthClientUsageQueryRequestTimeout{}
}

/*
PostOauthClientUsageQueryRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOauthClientUsageQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query request timeout response has a 2xx status code
func (o *PostOauthClientUsageQueryRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query request timeout response has a 3xx status code
func (o *PostOauthClientUsageQueryRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query request timeout response has a 4xx status code
func (o *PostOauthClientUsageQueryRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query request timeout response has a 5xx status code
func (o *PostOauthClientUsageQueryRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query request timeout response a status code equal to that given
func (o *PostOauthClientUsageQueryRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOauthClientUsageQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOauthClientUsageQueryRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOauthClientUsageQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryRequestEntityTooLarge creates a PostOauthClientUsageQueryRequestEntityTooLarge with default headers values
func NewPostOauthClientUsageQueryRequestEntityTooLarge() *PostOauthClientUsageQueryRequestEntityTooLarge {
	return &PostOauthClientUsageQueryRequestEntityTooLarge{}
}

/*
PostOauthClientUsageQueryRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOauthClientUsageQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query request entity too large response has a 2xx status code
func (o *PostOauthClientUsageQueryRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query request entity too large response has a 3xx status code
func (o *PostOauthClientUsageQueryRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query request entity too large response has a 4xx status code
func (o *PostOauthClientUsageQueryRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query request entity too large response has a 5xx status code
func (o *PostOauthClientUsageQueryRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query request entity too large response a status code equal to that given
func (o *PostOauthClientUsageQueryRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOauthClientUsageQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOauthClientUsageQueryRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOauthClientUsageQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryUnsupportedMediaType creates a PostOauthClientUsageQueryUnsupportedMediaType with default headers values
func NewPostOauthClientUsageQueryUnsupportedMediaType() *PostOauthClientUsageQueryUnsupportedMediaType {
	return &PostOauthClientUsageQueryUnsupportedMediaType{}
}

/*
PostOauthClientUsageQueryUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOauthClientUsageQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query unsupported media type response has a 2xx status code
func (o *PostOauthClientUsageQueryUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query unsupported media type response has a 3xx status code
func (o *PostOauthClientUsageQueryUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query unsupported media type response has a 4xx status code
func (o *PostOauthClientUsageQueryUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query unsupported media type response has a 5xx status code
func (o *PostOauthClientUsageQueryUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query unsupported media type response a status code equal to that given
func (o *PostOauthClientUsageQueryUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOauthClientUsageQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOauthClientUsageQueryUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOauthClientUsageQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryTooManyRequests creates a PostOauthClientUsageQueryTooManyRequests with default headers values
func NewPostOauthClientUsageQueryTooManyRequests() *PostOauthClientUsageQueryTooManyRequests {
	return &PostOauthClientUsageQueryTooManyRequests{}
}

/*
PostOauthClientUsageQueryTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOauthClientUsageQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query too many requests response has a 2xx status code
func (o *PostOauthClientUsageQueryTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query too many requests response has a 3xx status code
func (o *PostOauthClientUsageQueryTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query too many requests response has a 4xx status code
func (o *PostOauthClientUsageQueryTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post oauth client usage query too many requests response has a 5xx status code
func (o *PostOauthClientUsageQueryTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post oauth client usage query too many requests response a status code equal to that given
func (o *PostOauthClientUsageQueryTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOauthClientUsageQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOauthClientUsageQueryTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOauthClientUsageQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryInternalServerError creates a PostOauthClientUsageQueryInternalServerError with default headers values
func NewPostOauthClientUsageQueryInternalServerError() *PostOauthClientUsageQueryInternalServerError {
	return &PostOauthClientUsageQueryInternalServerError{}
}

/*
PostOauthClientUsageQueryInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOauthClientUsageQueryInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query internal server error response has a 2xx status code
func (o *PostOauthClientUsageQueryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query internal server error response has a 3xx status code
func (o *PostOauthClientUsageQueryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query internal server error response has a 4xx status code
func (o *PostOauthClientUsageQueryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth client usage query internal server error response has a 5xx status code
func (o *PostOauthClientUsageQueryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post oauth client usage query internal server error response a status code equal to that given
func (o *PostOauthClientUsageQueryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOauthClientUsageQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOauthClientUsageQueryInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOauthClientUsageQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryServiceUnavailable creates a PostOauthClientUsageQueryServiceUnavailable with default headers values
func NewPostOauthClientUsageQueryServiceUnavailable() *PostOauthClientUsageQueryServiceUnavailable {
	return &PostOauthClientUsageQueryServiceUnavailable{}
}

/*
PostOauthClientUsageQueryServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOauthClientUsageQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query service unavailable response has a 2xx status code
func (o *PostOauthClientUsageQueryServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query service unavailable response has a 3xx status code
func (o *PostOauthClientUsageQueryServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query service unavailable response has a 4xx status code
func (o *PostOauthClientUsageQueryServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth client usage query service unavailable response has a 5xx status code
func (o *PostOauthClientUsageQueryServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post oauth client usage query service unavailable response a status code equal to that given
func (o *PostOauthClientUsageQueryServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOauthClientUsageQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOauthClientUsageQueryServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOauthClientUsageQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOauthClientUsageQueryGatewayTimeout creates a PostOauthClientUsageQueryGatewayTimeout with default headers values
func NewPostOauthClientUsageQueryGatewayTimeout() *PostOauthClientUsageQueryGatewayTimeout {
	return &PostOauthClientUsageQueryGatewayTimeout{}
}

/*
PostOauthClientUsageQueryGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOauthClientUsageQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post oauth client usage query gateway timeout response has a 2xx status code
func (o *PostOauthClientUsageQueryGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post oauth client usage query gateway timeout response has a 3xx status code
func (o *PostOauthClientUsageQueryGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post oauth client usage query gateway timeout response has a 4xx status code
func (o *PostOauthClientUsageQueryGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post oauth client usage query gateway timeout response has a 5xx status code
func (o *PostOauthClientUsageQueryGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post oauth client usage query gateway timeout response a status code equal to that given
func (o *PostOauthClientUsageQueryGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOauthClientUsageQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOauthClientUsageQueryGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/oauth/clients/{clientId}/usage/query][%d] postOauthClientUsageQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOauthClientUsageQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOauthClientUsageQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
