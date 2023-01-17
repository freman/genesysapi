// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostResponsemanagementResponsesReader is a Reader for the PostResponsemanagementResponses structure.
type PostResponsemanagementResponsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostResponsemanagementResponsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostResponsemanagementResponsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostResponsemanagementResponsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostResponsemanagementResponsesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostResponsemanagementResponsesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostResponsemanagementResponsesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostResponsemanagementResponsesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewPostResponsemanagementResponsesPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostResponsemanagementResponsesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostResponsemanagementResponsesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostResponsemanagementResponsesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostResponsemanagementResponsesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostResponsemanagementResponsesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostResponsemanagementResponsesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostResponsemanagementResponsesOK creates a PostResponsemanagementResponsesOK with default headers values
func NewPostResponsemanagementResponsesOK() *PostResponsemanagementResponsesOK {
	return &PostResponsemanagementResponsesOK{}
}

/*
PostResponsemanagementResponsesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostResponsemanagementResponsesOK struct {
	Payload *models.Response
}

// IsSuccess returns true when this post responsemanagement responses o k response has a 2xx status code
func (o *PostResponsemanagementResponsesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post responsemanagement responses o k response has a 3xx status code
func (o *PostResponsemanagementResponsesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses o k response has a 4xx status code
func (o *PostResponsemanagementResponsesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responses o k response has a 5xx status code
func (o *PostResponsemanagementResponsesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses o k response a status code equal to that given
func (o *PostResponsemanagementResponsesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostResponsemanagementResponsesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesOK  %+v", 200, o.Payload)
}

func (o *PostResponsemanagementResponsesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesOK  %+v", 200, o.Payload)
}

func (o *PostResponsemanagementResponsesOK) GetPayload() *models.Response {
	return o.Payload
}

func (o *PostResponsemanagementResponsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesBadRequest creates a PostResponsemanagementResponsesBadRequest with default headers values
func NewPostResponsemanagementResponsesBadRequest() *PostResponsemanagementResponsesBadRequest {
	return &PostResponsemanagementResponsesBadRequest{}
}

/*
PostResponsemanagementResponsesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostResponsemanagementResponsesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses bad request response has a 2xx status code
func (o *PostResponsemanagementResponsesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses bad request response has a 3xx status code
func (o *PostResponsemanagementResponsesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses bad request response has a 4xx status code
func (o *PostResponsemanagementResponsesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses bad request response has a 5xx status code
func (o *PostResponsemanagementResponsesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses bad request response a status code equal to that given
func (o *PostResponsemanagementResponsesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostResponsemanagementResponsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesBadRequest  %+v", 400, o.Payload)
}

func (o *PostResponsemanagementResponsesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesBadRequest  %+v", 400, o.Payload)
}

func (o *PostResponsemanagementResponsesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesUnauthorized creates a PostResponsemanagementResponsesUnauthorized with default headers values
func NewPostResponsemanagementResponsesUnauthorized() *PostResponsemanagementResponsesUnauthorized {
	return &PostResponsemanagementResponsesUnauthorized{}
}

/*
PostResponsemanagementResponsesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostResponsemanagementResponsesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses unauthorized response has a 2xx status code
func (o *PostResponsemanagementResponsesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses unauthorized response has a 3xx status code
func (o *PostResponsemanagementResponsesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses unauthorized response has a 4xx status code
func (o *PostResponsemanagementResponsesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses unauthorized response has a 5xx status code
func (o *PostResponsemanagementResponsesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses unauthorized response a status code equal to that given
func (o *PostResponsemanagementResponsesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostResponsemanagementResponsesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostResponsemanagementResponsesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostResponsemanagementResponsesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesForbidden creates a PostResponsemanagementResponsesForbidden with default headers values
func NewPostResponsemanagementResponsesForbidden() *PostResponsemanagementResponsesForbidden {
	return &PostResponsemanagementResponsesForbidden{}
}

/*
PostResponsemanagementResponsesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostResponsemanagementResponsesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses forbidden response has a 2xx status code
func (o *PostResponsemanagementResponsesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses forbidden response has a 3xx status code
func (o *PostResponsemanagementResponsesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses forbidden response has a 4xx status code
func (o *PostResponsemanagementResponsesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses forbidden response has a 5xx status code
func (o *PostResponsemanagementResponsesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses forbidden response a status code equal to that given
func (o *PostResponsemanagementResponsesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostResponsemanagementResponsesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesForbidden  %+v", 403, o.Payload)
}

func (o *PostResponsemanagementResponsesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesForbidden  %+v", 403, o.Payload)
}

func (o *PostResponsemanagementResponsesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesNotFound creates a PostResponsemanagementResponsesNotFound with default headers values
func NewPostResponsemanagementResponsesNotFound() *PostResponsemanagementResponsesNotFound {
	return &PostResponsemanagementResponsesNotFound{}
}

/*
PostResponsemanagementResponsesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostResponsemanagementResponsesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses not found response has a 2xx status code
func (o *PostResponsemanagementResponsesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses not found response has a 3xx status code
func (o *PostResponsemanagementResponsesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses not found response has a 4xx status code
func (o *PostResponsemanagementResponsesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses not found response has a 5xx status code
func (o *PostResponsemanagementResponsesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses not found response a status code equal to that given
func (o *PostResponsemanagementResponsesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostResponsemanagementResponsesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesNotFound  %+v", 404, o.Payload)
}

func (o *PostResponsemanagementResponsesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesNotFound  %+v", 404, o.Payload)
}

func (o *PostResponsemanagementResponsesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesRequestTimeout creates a PostResponsemanagementResponsesRequestTimeout with default headers values
func NewPostResponsemanagementResponsesRequestTimeout() *PostResponsemanagementResponsesRequestTimeout {
	return &PostResponsemanagementResponsesRequestTimeout{}
}

/*
PostResponsemanagementResponsesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostResponsemanagementResponsesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses request timeout response has a 2xx status code
func (o *PostResponsemanagementResponsesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses request timeout response has a 3xx status code
func (o *PostResponsemanagementResponsesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses request timeout response has a 4xx status code
func (o *PostResponsemanagementResponsesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses request timeout response has a 5xx status code
func (o *PostResponsemanagementResponsesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses request timeout response a status code equal to that given
func (o *PostResponsemanagementResponsesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostResponsemanagementResponsesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostResponsemanagementResponsesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostResponsemanagementResponsesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesPreconditionFailed creates a PostResponsemanagementResponsesPreconditionFailed with default headers values
func NewPostResponsemanagementResponsesPreconditionFailed() *PostResponsemanagementResponsesPreconditionFailed {
	return &PostResponsemanagementResponsesPreconditionFailed{}
}

/*
PostResponsemanagementResponsesPreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type PostResponsemanagementResponsesPreconditionFailed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses precondition failed response has a 2xx status code
func (o *PostResponsemanagementResponsesPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses precondition failed response has a 3xx status code
func (o *PostResponsemanagementResponsesPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses precondition failed response has a 4xx status code
func (o *PostResponsemanagementResponsesPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses precondition failed response has a 5xx status code
func (o *PostResponsemanagementResponsesPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses precondition failed response a status code equal to that given
func (o *PostResponsemanagementResponsesPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

func (o *PostResponsemanagementResponsesPreconditionFailed) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PostResponsemanagementResponsesPreconditionFailed) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesPreconditionFailed  %+v", 412, o.Payload)
}

func (o *PostResponsemanagementResponsesPreconditionFailed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesRequestEntityTooLarge creates a PostResponsemanagementResponsesRequestEntityTooLarge with default headers values
func NewPostResponsemanagementResponsesRequestEntityTooLarge() *PostResponsemanagementResponsesRequestEntityTooLarge {
	return &PostResponsemanagementResponsesRequestEntityTooLarge{}
}

/*
PostResponsemanagementResponsesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostResponsemanagementResponsesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses request entity too large response has a 2xx status code
func (o *PostResponsemanagementResponsesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses request entity too large response has a 3xx status code
func (o *PostResponsemanagementResponsesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses request entity too large response has a 4xx status code
func (o *PostResponsemanagementResponsesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses request entity too large response has a 5xx status code
func (o *PostResponsemanagementResponsesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses request entity too large response a status code equal to that given
func (o *PostResponsemanagementResponsesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostResponsemanagementResponsesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostResponsemanagementResponsesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostResponsemanagementResponsesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesUnsupportedMediaType creates a PostResponsemanagementResponsesUnsupportedMediaType with default headers values
func NewPostResponsemanagementResponsesUnsupportedMediaType() *PostResponsemanagementResponsesUnsupportedMediaType {
	return &PostResponsemanagementResponsesUnsupportedMediaType{}
}

/*
PostResponsemanagementResponsesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostResponsemanagementResponsesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses unsupported media type response has a 2xx status code
func (o *PostResponsemanagementResponsesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses unsupported media type response has a 3xx status code
func (o *PostResponsemanagementResponsesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses unsupported media type response has a 4xx status code
func (o *PostResponsemanagementResponsesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses unsupported media type response has a 5xx status code
func (o *PostResponsemanagementResponsesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses unsupported media type response a status code equal to that given
func (o *PostResponsemanagementResponsesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostResponsemanagementResponsesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostResponsemanagementResponsesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostResponsemanagementResponsesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesTooManyRequests creates a PostResponsemanagementResponsesTooManyRequests with default headers values
func NewPostResponsemanagementResponsesTooManyRequests() *PostResponsemanagementResponsesTooManyRequests {
	return &PostResponsemanagementResponsesTooManyRequests{}
}

/*
PostResponsemanagementResponsesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostResponsemanagementResponsesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses too many requests response has a 2xx status code
func (o *PostResponsemanagementResponsesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses too many requests response has a 3xx status code
func (o *PostResponsemanagementResponsesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses too many requests response has a 4xx status code
func (o *PostResponsemanagementResponsesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post responsemanagement responses too many requests response has a 5xx status code
func (o *PostResponsemanagementResponsesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post responsemanagement responses too many requests response a status code equal to that given
func (o *PostResponsemanagementResponsesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostResponsemanagementResponsesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostResponsemanagementResponsesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostResponsemanagementResponsesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesInternalServerError creates a PostResponsemanagementResponsesInternalServerError with default headers values
func NewPostResponsemanagementResponsesInternalServerError() *PostResponsemanagementResponsesInternalServerError {
	return &PostResponsemanagementResponsesInternalServerError{}
}

/*
PostResponsemanagementResponsesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostResponsemanagementResponsesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses internal server error response has a 2xx status code
func (o *PostResponsemanagementResponsesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses internal server error response has a 3xx status code
func (o *PostResponsemanagementResponsesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses internal server error response has a 4xx status code
func (o *PostResponsemanagementResponsesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responses internal server error response has a 5xx status code
func (o *PostResponsemanagementResponsesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responses internal server error response a status code equal to that given
func (o *PostResponsemanagementResponsesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostResponsemanagementResponsesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostResponsemanagementResponsesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostResponsemanagementResponsesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesServiceUnavailable creates a PostResponsemanagementResponsesServiceUnavailable with default headers values
func NewPostResponsemanagementResponsesServiceUnavailable() *PostResponsemanagementResponsesServiceUnavailable {
	return &PostResponsemanagementResponsesServiceUnavailable{}
}

/*
PostResponsemanagementResponsesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostResponsemanagementResponsesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses service unavailable response has a 2xx status code
func (o *PostResponsemanagementResponsesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses service unavailable response has a 3xx status code
func (o *PostResponsemanagementResponsesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses service unavailable response has a 4xx status code
func (o *PostResponsemanagementResponsesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responses service unavailable response has a 5xx status code
func (o *PostResponsemanagementResponsesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responses service unavailable response a status code equal to that given
func (o *PostResponsemanagementResponsesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostResponsemanagementResponsesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostResponsemanagementResponsesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostResponsemanagementResponsesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostResponsemanagementResponsesGatewayTimeout creates a PostResponsemanagementResponsesGatewayTimeout with default headers values
func NewPostResponsemanagementResponsesGatewayTimeout() *PostResponsemanagementResponsesGatewayTimeout {
	return &PostResponsemanagementResponsesGatewayTimeout{}
}

/*
PostResponsemanagementResponsesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostResponsemanagementResponsesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post responsemanagement responses gateway timeout response has a 2xx status code
func (o *PostResponsemanagementResponsesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post responsemanagement responses gateway timeout response has a 3xx status code
func (o *PostResponsemanagementResponsesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post responsemanagement responses gateway timeout response has a 4xx status code
func (o *PostResponsemanagementResponsesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post responsemanagement responses gateway timeout response has a 5xx status code
func (o *PostResponsemanagementResponsesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post responsemanagement responses gateway timeout response a status code equal to that given
func (o *PostResponsemanagementResponsesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostResponsemanagementResponsesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostResponsemanagementResponsesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/responsemanagement/responses][%d] postResponsemanagementResponsesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostResponsemanagementResponsesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostResponsemanagementResponsesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
