// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuthorizationSubjectBulkreplaceReader is a Reader for the PostAuthorizationSubjectBulkreplace structure.
type PostAuthorizationSubjectBulkreplaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuthorizationSubjectBulkreplaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostAuthorizationSubjectBulkreplaceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuthorizationSubjectBulkreplaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuthorizationSubjectBulkreplaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuthorizationSubjectBulkreplaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuthorizationSubjectBulkreplaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAuthorizationSubjectBulkreplaceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuthorizationSubjectBulkreplaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuthorizationSubjectBulkreplaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuthorizationSubjectBulkreplaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuthorizationSubjectBulkreplaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuthorizationSubjectBulkreplaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuthorizationSubjectBulkreplaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuthorizationSubjectBulkreplaceNoContent creates a PostAuthorizationSubjectBulkreplaceNoContent with default headers values
func NewPostAuthorizationSubjectBulkreplaceNoContent() *PostAuthorizationSubjectBulkreplaceNoContent {
	return &PostAuthorizationSubjectBulkreplaceNoContent{}
}

/*
PostAuthorizationSubjectBulkreplaceNoContent describes a response with status code 204, with default header values.

Bulk Grants Replaced
*/
type PostAuthorizationSubjectBulkreplaceNoContent struct {
}

// IsSuccess returns true when this post authorization subject bulkreplace no content response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post authorization subject bulkreplace no content response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace no content response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post authorization subject bulkreplace no content response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace no content response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PostAuthorizationSubjectBulkreplaceNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceNoContent ", 204)
}

func (o *PostAuthorizationSubjectBulkreplaceNoContent) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceNoContent ", 204)
}

func (o *PostAuthorizationSubjectBulkreplaceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceBadRequest creates a PostAuthorizationSubjectBulkreplaceBadRequest with default headers values
func NewPostAuthorizationSubjectBulkreplaceBadRequest() *PostAuthorizationSubjectBulkreplaceBadRequest {
	return &PostAuthorizationSubjectBulkreplaceBadRequest{}
}

/*
PostAuthorizationSubjectBulkreplaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuthorizationSubjectBulkreplaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace bad request response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace bad request response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace bad request response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace bad request response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace bad request response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostAuthorizationSubjectBulkreplaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceUnauthorized creates a PostAuthorizationSubjectBulkreplaceUnauthorized with default headers values
func NewPostAuthorizationSubjectBulkreplaceUnauthorized() *PostAuthorizationSubjectBulkreplaceUnauthorized {
	return &PostAuthorizationSubjectBulkreplaceUnauthorized{}
}

/*
PostAuthorizationSubjectBulkreplaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuthorizationSubjectBulkreplaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace unauthorized response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace unauthorized response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace unauthorized response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace unauthorized response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace unauthorized response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceForbidden creates a PostAuthorizationSubjectBulkreplaceForbidden with default headers values
func NewPostAuthorizationSubjectBulkreplaceForbidden() *PostAuthorizationSubjectBulkreplaceForbidden {
	return &PostAuthorizationSubjectBulkreplaceForbidden{}
}

/*
PostAuthorizationSubjectBulkreplaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostAuthorizationSubjectBulkreplaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace forbidden response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace forbidden response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace forbidden response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace forbidden response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace forbidden response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostAuthorizationSubjectBulkreplaceForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceForbidden  %+v", 403, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceNotFound creates a PostAuthorizationSubjectBulkreplaceNotFound with default headers values
func NewPostAuthorizationSubjectBulkreplaceNotFound() *PostAuthorizationSubjectBulkreplaceNotFound {
	return &PostAuthorizationSubjectBulkreplaceNotFound{}
}

/*
PostAuthorizationSubjectBulkreplaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostAuthorizationSubjectBulkreplaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace not found response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace not found response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace not found response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace not found response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace not found response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostAuthorizationSubjectBulkreplaceNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceNotFound  %+v", 404, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceRequestTimeout creates a PostAuthorizationSubjectBulkreplaceRequestTimeout with default headers values
func NewPostAuthorizationSubjectBulkreplaceRequestTimeout() *PostAuthorizationSubjectBulkreplaceRequestTimeout {
	return &PostAuthorizationSubjectBulkreplaceRequestTimeout{}
}

/*
PostAuthorizationSubjectBulkreplaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAuthorizationSubjectBulkreplaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace request timeout response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace request timeout response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace request timeout response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace request timeout response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace request timeout response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceRequestEntityTooLarge creates a PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge with default headers values
func NewPostAuthorizationSubjectBulkreplaceRequestEntityTooLarge() *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge {
	return &PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge{}
}

/*
PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace request entity too large response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace request entity too large response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace request entity too large response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace request entity too large response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace request entity too large response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceUnsupportedMediaType creates a PostAuthorizationSubjectBulkreplaceUnsupportedMediaType with default headers values
func NewPostAuthorizationSubjectBulkreplaceUnsupportedMediaType() *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType {
	return &PostAuthorizationSubjectBulkreplaceUnsupportedMediaType{}
}

/*
PostAuthorizationSubjectBulkreplaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuthorizationSubjectBulkreplaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace unsupported media type response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace unsupported media type response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace unsupported media type response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace unsupported media type response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace unsupported media type response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceTooManyRequests creates a PostAuthorizationSubjectBulkreplaceTooManyRequests with default headers values
func NewPostAuthorizationSubjectBulkreplaceTooManyRequests() *PostAuthorizationSubjectBulkreplaceTooManyRequests {
	return &PostAuthorizationSubjectBulkreplaceTooManyRequests{}
}

/*
PostAuthorizationSubjectBulkreplaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAuthorizationSubjectBulkreplaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace too many requests response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace too many requests response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace too many requests response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post authorization subject bulkreplace too many requests response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post authorization subject bulkreplace too many requests response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceInternalServerError creates a PostAuthorizationSubjectBulkreplaceInternalServerError with default headers values
func NewPostAuthorizationSubjectBulkreplaceInternalServerError() *PostAuthorizationSubjectBulkreplaceInternalServerError {
	return &PostAuthorizationSubjectBulkreplaceInternalServerError{}
}

/*
PostAuthorizationSubjectBulkreplaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuthorizationSubjectBulkreplaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace internal server error response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace internal server error response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace internal server error response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post authorization subject bulkreplace internal server error response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post authorization subject bulkreplace internal server error response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceServiceUnavailable creates a PostAuthorizationSubjectBulkreplaceServiceUnavailable with default headers values
func NewPostAuthorizationSubjectBulkreplaceServiceUnavailable() *PostAuthorizationSubjectBulkreplaceServiceUnavailable {
	return &PostAuthorizationSubjectBulkreplaceServiceUnavailable{}
}

/*
PostAuthorizationSubjectBulkreplaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuthorizationSubjectBulkreplaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace service unavailable response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace service unavailable response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace service unavailable response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post authorization subject bulkreplace service unavailable response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post authorization subject bulkreplace service unavailable response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuthorizationSubjectBulkreplaceGatewayTimeout creates a PostAuthorizationSubjectBulkreplaceGatewayTimeout with default headers values
func NewPostAuthorizationSubjectBulkreplaceGatewayTimeout() *PostAuthorizationSubjectBulkreplaceGatewayTimeout {
	return &PostAuthorizationSubjectBulkreplaceGatewayTimeout{}
}

/*
PostAuthorizationSubjectBulkreplaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostAuthorizationSubjectBulkreplaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post authorization subject bulkreplace gateway timeout response has a 2xx status code
func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post authorization subject bulkreplace gateway timeout response has a 3xx status code
func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post authorization subject bulkreplace gateway timeout response has a 4xx status code
func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post authorization subject bulkreplace gateway timeout response has a 5xx status code
func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post authorization subject bulkreplace gateway timeout response a status code equal to that given
func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/authorization/subjects/{subjectId}/bulkreplace][%d] postAuthorizationSubjectBulkreplaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuthorizationSubjectBulkreplaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
