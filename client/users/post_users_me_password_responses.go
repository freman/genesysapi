// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostUsersMePasswordReader is a Reader for the PostUsersMePassword structure.
type PostUsersMePasswordReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUsersMePasswordReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostUsersMePasswordNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUsersMePasswordBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUsersMePasswordUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUsersMePasswordForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUsersMePasswordNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostUsersMePasswordRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostUsersMePasswordRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostUsersMePasswordUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostUsersMePasswordTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUsersMePasswordInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostUsersMePasswordServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostUsersMePasswordGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostUsersMePasswordNoContent creates a PostUsersMePasswordNoContent with default headers values
func NewPostUsersMePasswordNoContent() *PostUsersMePasswordNoContent {
	return &PostUsersMePasswordNoContent{}
}

/*
PostUsersMePasswordNoContent describes a response with status code 204, with default header values.

Password changed
*/
type PostUsersMePasswordNoContent struct {
}

// IsSuccess returns true when this post users me password no content response has a 2xx status code
func (o *PostUsersMePasswordNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post users me password no content response has a 3xx status code
func (o *PostUsersMePasswordNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password no content response has a 4xx status code
func (o *PostUsersMePasswordNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users me password no content response has a 5xx status code
func (o *PostUsersMePasswordNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password no content response a status code equal to that given
func (o *PostUsersMePasswordNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PostUsersMePasswordNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordNoContent ", 204)
}

func (o *PostUsersMePasswordNoContent) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordNoContent ", 204)
}

func (o *PostUsersMePasswordNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostUsersMePasswordBadRequest creates a PostUsersMePasswordBadRequest with default headers values
func NewPostUsersMePasswordBadRequest() *PostUsersMePasswordBadRequest {
	return &PostUsersMePasswordBadRequest{}
}

/*
PostUsersMePasswordBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostUsersMePasswordBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password bad request response has a 2xx status code
func (o *PostUsersMePasswordBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password bad request response has a 3xx status code
func (o *PostUsersMePasswordBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password bad request response has a 4xx status code
func (o *PostUsersMePasswordBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password bad request response has a 5xx status code
func (o *PostUsersMePasswordBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password bad request response a status code equal to that given
func (o *PostUsersMePasswordBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostUsersMePasswordBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersMePasswordBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordBadRequest  %+v", 400, o.Payload)
}

func (o *PostUsersMePasswordBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordUnauthorized creates a PostUsersMePasswordUnauthorized with default headers values
func NewPostUsersMePasswordUnauthorized() *PostUsersMePasswordUnauthorized {
	return &PostUsersMePasswordUnauthorized{}
}

/*
PostUsersMePasswordUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostUsersMePasswordUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password unauthorized response has a 2xx status code
func (o *PostUsersMePasswordUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password unauthorized response has a 3xx status code
func (o *PostUsersMePasswordUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password unauthorized response has a 4xx status code
func (o *PostUsersMePasswordUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password unauthorized response has a 5xx status code
func (o *PostUsersMePasswordUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password unauthorized response a status code equal to that given
func (o *PostUsersMePasswordUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostUsersMePasswordUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersMePasswordUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUsersMePasswordUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordForbidden creates a PostUsersMePasswordForbidden with default headers values
func NewPostUsersMePasswordForbidden() *PostUsersMePasswordForbidden {
	return &PostUsersMePasswordForbidden{}
}

/*
PostUsersMePasswordForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostUsersMePasswordForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password forbidden response has a 2xx status code
func (o *PostUsersMePasswordForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password forbidden response has a 3xx status code
func (o *PostUsersMePasswordForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password forbidden response has a 4xx status code
func (o *PostUsersMePasswordForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password forbidden response has a 5xx status code
func (o *PostUsersMePasswordForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password forbidden response a status code equal to that given
func (o *PostUsersMePasswordForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostUsersMePasswordForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersMePasswordForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordForbidden  %+v", 403, o.Payload)
}

func (o *PostUsersMePasswordForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordNotFound creates a PostUsersMePasswordNotFound with default headers values
func NewPostUsersMePasswordNotFound() *PostUsersMePasswordNotFound {
	return &PostUsersMePasswordNotFound{}
}

/*
PostUsersMePasswordNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostUsersMePasswordNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password not found response has a 2xx status code
func (o *PostUsersMePasswordNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password not found response has a 3xx status code
func (o *PostUsersMePasswordNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password not found response has a 4xx status code
func (o *PostUsersMePasswordNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password not found response has a 5xx status code
func (o *PostUsersMePasswordNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password not found response a status code equal to that given
func (o *PostUsersMePasswordNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostUsersMePasswordNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersMePasswordNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordNotFound  %+v", 404, o.Payload)
}

func (o *PostUsersMePasswordNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordRequestTimeout creates a PostUsersMePasswordRequestTimeout with default headers values
func NewPostUsersMePasswordRequestTimeout() *PostUsersMePasswordRequestTimeout {
	return &PostUsersMePasswordRequestTimeout{}
}

/*
PostUsersMePasswordRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostUsersMePasswordRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password request timeout response has a 2xx status code
func (o *PostUsersMePasswordRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password request timeout response has a 3xx status code
func (o *PostUsersMePasswordRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password request timeout response has a 4xx status code
func (o *PostUsersMePasswordRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password request timeout response has a 5xx status code
func (o *PostUsersMePasswordRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password request timeout response a status code equal to that given
func (o *PostUsersMePasswordRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostUsersMePasswordRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersMePasswordRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostUsersMePasswordRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordRequestEntityTooLarge creates a PostUsersMePasswordRequestEntityTooLarge with default headers values
func NewPostUsersMePasswordRequestEntityTooLarge() *PostUsersMePasswordRequestEntityTooLarge {
	return &PostUsersMePasswordRequestEntityTooLarge{}
}

/*
PostUsersMePasswordRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostUsersMePasswordRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password request entity too large response has a 2xx status code
func (o *PostUsersMePasswordRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password request entity too large response has a 3xx status code
func (o *PostUsersMePasswordRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password request entity too large response has a 4xx status code
func (o *PostUsersMePasswordRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password request entity too large response has a 5xx status code
func (o *PostUsersMePasswordRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password request entity too large response a status code equal to that given
func (o *PostUsersMePasswordRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostUsersMePasswordRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersMePasswordRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostUsersMePasswordRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordUnsupportedMediaType creates a PostUsersMePasswordUnsupportedMediaType with default headers values
func NewPostUsersMePasswordUnsupportedMediaType() *PostUsersMePasswordUnsupportedMediaType {
	return &PostUsersMePasswordUnsupportedMediaType{}
}

/*
PostUsersMePasswordUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostUsersMePasswordUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password unsupported media type response has a 2xx status code
func (o *PostUsersMePasswordUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password unsupported media type response has a 3xx status code
func (o *PostUsersMePasswordUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password unsupported media type response has a 4xx status code
func (o *PostUsersMePasswordUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password unsupported media type response has a 5xx status code
func (o *PostUsersMePasswordUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password unsupported media type response a status code equal to that given
func (o *PostUsersMePasswordUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostUsersMePasswordUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersMePasswordUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostUsersMePasswordUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordTooManyRequests creates a PostUsersMePasswordTooManyRequests with default headers values
func NewPostUsersMePasswordTooManyRequests() *PostUsersMePasswordTooManyRequests {
	return &PostUsersMePasswordTooManyRequests{}
}

/*
PostUsersMePasswordTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostUsersMePasswordTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password too many requests response has a 2xx status code
func (o *PostUsersMePasswordTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password too many requests response has a 3xx status code
func (o *PostUsersMePasswordTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password too many requests response has a 4xx status code
func (o *PostUsersMePasswordTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post users me password too many requests response has a 5xx status code
func (o *PostUsersMePasswordTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post users me password too many requests response a status code equal to that given
func (o *PostUsersMePasswordTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostUsersMePasswordTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersMePasswordTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostUsersMePasswordTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordInternalServerError creates a PostUsersMePasswordInternalServerError with default headers values
func NewPostUsersMePasswordInternalServerError() *PostUsersMePasswordInternalServerError {
	return &PostUsersMePasswordInternalServerError{}
}

/*
PostUsersMePasswordInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostUsersMePasswordInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password internal server error response has a 2xx status code
func (o *PostUsersMePasswordInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password internal server error response has a 3xx status code
func (o *PostUsersMePasswordInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password internal server error response has a 4xx status code
func (o *PostUsersMePasswordInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users me password internal server error response has a 5xx status code
func (o *PostUsersMePasswordInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post users me password internal server error response a status code equal to that given
func (o *PostUsersMePasswordInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostUsersMePasswordInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersMePasswordInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUsersMePasswordInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordServiceUnavailable creates a PostUsersMePasswordServiceUnavailable with default headers values
func NewPostUsersMePasswordServiceUnavailable() *PostUsersMePasswordServiceUnavailable {
	return &PostUsersMePasswordServiceUnavailable{}
}

/*
PostUsersMePasswordServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostUsersMePasswordServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password service unavailable response has a 2xx status code
func (o *PostUsersMePasswordServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password service unavailable response has a 3xx status code
func (o *PostUsersMePasswordServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password service unavailable response has a 4xx status code
func (o *PostUsersMePasswordServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users me password service unavailable response has a 5xx status code
func (o *PostUsersMePasswordServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post users me password service unavailable response a status code equal to that given
func (o *PostUsersMePasswordServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostUsersMePasswordServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersMePasswordServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostUsersMePasswordServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUsersMePasswordGatewayTimeout creates a PostUsersMePasswordGatewayTimeout with default headers values
func NewPostUsersMePasswordGatewayTimeout() *PostUsersMePasswordGatewayTimeout {
	return &PostUsersMePasswordGatewayTimeout{}
}

/*
PostUsersMePasswordGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostUsersMePasswordGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post users me password gateway timeout response has a 2xx status code
func (o *PostUsersMePasswordGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post users me password gateway timeout response has a 3xx status code
func (o *PostUsersMePasswordGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post users me password gateway timeout response has a 4xx status code
func (o *PostUsersMePasswordGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post users me password gateway timeout response has a 5xx status code
func (o *PostUsersMePasswordGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post users me password gateway timeout response a status code equal to that given
func (o *PostUsersMePasswordGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostUsersMePasswordGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersMePasswordGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/users/me/password][%d] postUsersMePasswordGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostUsersMePasswordGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostUsersMePasswordGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
