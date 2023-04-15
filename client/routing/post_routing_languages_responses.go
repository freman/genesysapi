// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostRoutingLanguagesReader is a Reader for the PostRoutingLanguages structure.
type PostRoutingLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRoutingLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRoutingLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRoutingLanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRoutingLanguagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRoutingLanguagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRoutingLanguagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostRoutingLanguagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostRoutingLanguagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostRoutingLanguagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostRoutingLanguagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostRoutingLanguagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostRoutingLanguagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostRoutingLanguagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostRoutingLanguagesOK creates a PostRoutingLanguagesOK with default headers values
func NewPostRoutingLanguagesOK() *PostRoutingLanguagesOK {
	return &PostRoutingLanguagesOK{}
}

/*
PostRoutingLanguagesOK describes a response with status code 200, with default header values.

successful operation
*/
type PostRoutingLanguagesOK struct {
	Payload *models.Language
}

// IsSuccess returns true when this post routing languages o k response has a 2xx status code
func (o *PostRoutingLanguagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post routing languages o k response has a 3xx status code
func (o *PostRoutingLanguagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages o k response has a 4xx status code
func (o *PostRoutingLanguagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing languages o k response has a 5xx status code
func (o *PostRoutingLanguagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages o k response a status code equal to that given
func (o *PostRoutingLanguagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostRoutingLanguagesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesOK  %+v", 200, o.Payload)
}

func (o *PostRoutingLanguagesOK) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesOK  %+v", 200, o.Payload)
}

func (o *PostRoutingLanguagesOK) GetPayload() *models.Language {
	return o.Payload
}

func (o *PostRoutingLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Language)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesBadRequest creates a PostRoutingLanguagesBadRequest with default headers values
func NewPostRoutingLanguagesBadRequest() *PostRoutingLanguagesBadRequest {
	return &PostRoutingLanguagesBadRequest{}
}

/*
PostRoutingLanguagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostRoutingLanguagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages bad request response has a 2xx status code
func (o *PostRoutingLanguagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages bad request response has a 3xx status code
func (o *PostRoutingLanguagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages bad request response has a 4xx status code
func (o *PostRoutingLanguagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages bad request response has a 5xx status code
func (o *PostRoutingLanguagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages bad request response a status code equal to that given
func (o *PostRoutingLanguagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostRoutingLanguagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingLanguagesBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *PostRoutingLanguagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesUnauthorized creates a PostRoutingLanguagesUnauthorized with default headers values
func NewPostRoutingLanguagesUnauthorized() *PostRoutingLanguagesUnauthorized {
	return &PostRoutingLanguagesUnauthorized{}
}

/*
PostRoutingLanguagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostRoutingLanguagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages unauthorized response has a 2xx status code
func (o *PostRoutingLanguagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages unauthorized response has a 3xx status code
func (o *PostRoutingLanguagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages unauthorized response has a 4xx status code
func (o *PostRoutingLanguagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages unauthorized response has a 5xx status code
func (o *PostRoutingLanguagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages unauthorized response a status code equal to that given
func (o *PostRoutingLanguagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostRoutingLanguagesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingLanguagesUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRoutingLanguagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesForbidden creates a PostRoutingLanguagesForbidden with default headers values
func NewPostRoutingLanguagesForbidden() *PostRoutingLanguagesForbidden {
	return &PostRoutingLanguagesForbidden{}
}

/*
PostRoutingLanguagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostRoutingLanguagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages forbidden response has a 2xx status code
func (o *PostRoutingLanguagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages forbidden response has a 3xx status code
func (o *PostRoutingLanguagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages forbidden response has a 4xx status code
func (o *PostRoutingLanguagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages forbidden response has a 5xx status code
func (o *PostRoutingLanguagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages forbidden response a status code equal to that given
func (o *PostRoutingLanguagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostRoutingLanguagesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingLanguagesForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *PostRoutingLanguagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesNotFound creates a PostRoutingLanguagesNotFound with default headers values
func NewPostRoutingLanguagesNotFound() *PostRoutingLanguagesNotFound {
	return &PostRoutingLanguagesNotFound{}
}

/*
PostRoutingLanguagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostRoutingLanguagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages not found response has a 2xx status code
func (o *PostRoutingLanguagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages not found response has a 3xx status code
func (o *PostRoutingLanguagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages not found response has a 4xx status code
func (o *PostRoutingLanguagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages not found response has a 5xx status code
func (o *PostRoutingLanguagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages not found response a status code equal to that given
func (o *PostRoutingLanguagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostRoutingLanguagesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingLanguagesNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *PostRoutingLanguagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesRequestTimeout creates a PostRoutingLanguagesRequestTimeout with default headers values
func NewPostRoutingLanguagesRequestTimeout() *PostRoutingLanguagesRequestTimeout {
	return &PostRoutingLanguagesRequestTimeout{}
}

/*
PostRoutingLanguagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostRoutingLanguagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages request timeout response has a 2xx status code
func (o *PostRoutingLanguagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages request timeout response has a 3xx status code
func (o *PostRoutingLanguagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages request timeout response has a 4xx status code
func (o *PostRoutingLanguagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages request timeout response has a 5xx status code
func (o *PostRoutingLanguagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages request timeout response a status code equal to that given
func (o *PostRoutingLanguagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostRoutingLanguagesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingLanguagesRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostRoutingLanguagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesRequestEntityTooLarge creates a PostRoutingLanguagesRequestEntityTooLarge with default headers values
func NewPostRoutingLanguagesRequestEntityTooLarge() *PostRoutingLanguagesRequestEntityTooLarge {
	return &PostRoutingLanguagesRequestEntityTooLarge{}
}

/*
PostRoutingLanguagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostRoutingLanguagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages request entity too large response has a 2xx status code
func (o *PostRoutingLanguagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages request entity too large response has a 3xx status code
func (o *PostRoutingLanguagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages request entity too large response has a 4xx status code
func (o *PostRoutingLanguagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages request entity too large response has a 5xx status code
func (o *PostRoutingLanguagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages request entity too large response a status code equal to that given
func (o *PostRoutingLanguagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesUnsupportedMediaType creates a PostRoutingLanguagesUnsupportedMediaType with default headers values
func NewPostRoutingLanguagesUnsupportedMediaType() *PostRoutingLanguagesUnsupportedMediaType {
	return &PostRoutingLanguagesUnsupportedMediaType{}
}

/*
PostRoutingLanguagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostRoutingLanguagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages unsupported media type response has a 2xx status code
func (o *PostRoutingLanguagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages unsupported media type response has a 3xx status code
func (o *PostRoutingLanguagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages unsupported media type response has a 4xx status code
func (o *PostRoutingLanguagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages unsupported media type response has a 5xx status code
func (o *PostRoutingLanguagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages unsupported media type response a status code equal to that given
func (o *PostRoutingLanguagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostRoutingLanguagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingLanguagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostRoutingLanguagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesTooManyRequests creates a PostRoutingLanguagesTooManyRequests with default headers values
func NewPostRoutingLanguagesTooManyRequests() *PostRoutingLanguagesTooManyRequests {
	return &PostRoutingLanguagesTooManyRequests{}
}

/*
PostRoutingLanguagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostRoutingLanguagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages too many requests response has a 2xx status code
func (o *PostRoutingLanguagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages too many requests response has a 3xx status code
func (o *PostRoutingLanguagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages too many requests response has a 4xx status code
func (o *PostRoutingLanguagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post routing languages too many requests response has a 5xx status code
func (o *PostRoutingLanguagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post routing languages too many requests response a status code equal to that given
func (o *PostRoutingLanguagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostRoutingLanguagesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingLanguagesTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostRoutingLanguagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesInternalServerError creates a PostRoutingLanguagesInternalServerError with default headers values
func NewPostRoutingLanguagesInternalServerError() *PostRoutingLanguagesInternalServerError {
	return &PostRoutingLanguagesInternalServerError{}
}

/*
PostRoutingLanguagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostRoutingLanguagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages internal server error response has a 2xx status code
func (o *PostRoutingLanguagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages internal server error response has a 3xx status code
func (o *PostRoutingLanguagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages internal server error response has a 4xx status code
func (o *PostRoutingLanguagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing languages internal server error response has a 5xx status code
func (o *PostRoutingLanguagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing languages internal server error response a status code equal to that given
func (o *PostRoutingLanguagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostRoutingLanguagesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingLanguagesInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostRoutingLanguagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesServiceUnavailable creates a PostRoutingLanguagesServiceUnavailable with default headers values
func NewPostRoutingLanguagesServiceUnavailable() *PostRoutingLanguagesServiceUnavailable {
	return &PostRoutingLanguagesServiceUnavailable{}
}

/*
PostRoutingLanguagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostRoutingLanguagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages service unavailable response has a 2xx status code
func (o *PostRoutingLanguagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages service unavailable response has a 3xx status code
func (o *PostRoutingLanguagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages service unavailable response has a 4xx status code
func (o *PostRoutingLanguagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing languages service unavailable response has a 5xx status code
func (o *PostRoutingLanguagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing languages service unavailable response a status code equal to that given
func (o *PostRoutingLanguagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostRoutingLanguagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingLanguagesServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostRoutingLanguagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRoutingLanguagesGatewayTimeout creates a PostRoutingLanguagesGatewayTimeout with default headers values
func NewPostRoutingLanguagesGatewayTimeout() *PostRoutingLanguagesGatewayTimeout {
	return &PostRoutingLanguagesGatewayTimeout{}
}

/*
PostRoutingLanguagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostRoutingLanguagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post routing languages gateway timeout response has a 2xx status code
func (o *PostRoutingLanguagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post routing languages gateway timeout response has a 3xx status code
func (o *PostRoutingLanguagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post routing languages gateway timeout response has a 4xx status code
func (o *PostRoutingLanguagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post routing languages gateway timeout response has a 5xx status code
func (o *PostRoutingLanguagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post routing languages gateway timeout response a status code equal to that given
func (o *PostRoutingLanguagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostRoutingLanguagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingLanguagesGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/routing/languages][%d] postRoutingLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostRoutingLanguagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostRoutingLanguagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
