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

// GetOauthScopesReader is a Reader for the GetOauthScopes structure.
type GetOauthScopesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOauthScopesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOauthScopesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOauthScopesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOauthScopesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOauthScopesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOauthScopesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOauthScopesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOauthScopesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOauthScopesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOauthScopesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOauthScopesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOauthScopesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOauthScopesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOauthScopesOK creates a GetOauthScopesOK with default headers values
func NewGetOauthScopesOK() *GetOauthScopesOK {
	return &GetOauthScopesOK{}
}

/*
GetOauthScopesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOauthScopesOK struct {
	Payload *models.OAuthScopeListing
}

// IsSuccess returns true when this get oauth scopes o k response has a 2xx status code
func (o *GetOauthScopesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get oauth scopes o k response has a 3xx status code
func (o *GetOauthScopesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes o k response has a 4xx status code
func (o *GetOauthScopesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get oauth scopes o k response has a 5xx status code
func (o *GetOauthScopesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes o k response a status code equal to that given
func (o *GetOauthScopesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOauthScopesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesOK  %+v", 200, o.Payload)
}

func (o *GetOauthScopesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesOK  %+v", 200, o.Payload)
}

func (o *GetOauthScopesOK) GetPayload() *models.OAuthScopeListing {
	return o.Payload
}

func (o *GetOauthScopesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuthScopeListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesBadRequest creates a GetOauthScopesBadRequest with default headers values
func NewGetOauthScopesBadRequest() *GetOauthScopesBadRequest {
	return &GetOauthScopesBadRequest{}
}

/*
GetOauthScopesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOauthScopesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes bad request response has a 2xx status code
func (o *GetOauthScopesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes bad request response has a 3xx status code
func (o *GetOauthScopesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes bad request response has a 4xx status code
func (o *GetOauthScopesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes bad request response has a 5xx status code
func (o *GetOauthScopesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes bad request response a status code equal to that given
func (o *GetOauthScopesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOauthScopesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOauthScopesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOauthScopesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesUnauthorized creates a GetOauthScopesUnauthorized with default headers values
func NewGetOauthScopesUnauthorized() *GetOauthScopesUnauthorized {
	return &GetOauthScopesUnauthorized{}
}

/*
GetOauthScopesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOauthScopesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes unauthorized response has a 2xx status code
func (o *GetOauthScopesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes unauthorized response has a 3xx status code
func (o *GetOauthScopesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes unauthorized response has a 4xx status code
func (o *GetOauthScopesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes unauthorized response has a 5xx status code
func (o *GetOauthScopesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes unauthorized response a status code equal to that given
func (o *GetOauthScopesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOauthScopesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOauthScopesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOauthScopesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesForbidden creates a GetOauthScopesForbidden with default headers values
func NewGetOauthScopesForbidden() *GetOauthScopesForbidden {
	return &GetOauthScopesForbidden{}
}

/*
GetOauthScopesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOauthScopesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes forbidden response has a 2xx status code
func (o *GetOauthScopesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes forbidden response has a 3xx status code
func (o *GetOauthScopesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes forbidden response has a 4xx status code
func (o *GetOauthScopesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes forbidden response has a 5xx status code
func (o *GetOauthScopesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes forbidden response a status code equal to that given
func (o *GetOauthScopesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOauthScopesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesForbidden  %+v", 403, o.Payload)
}

func (o *GetOauthScopesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesForbidden  %+v", 403, o.Payload)
}

func (o *GetOauthScopesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesNotFound creates a GetOauthScopesNotFound with default headers values
func NewGetOauthScopesNotFound() *GetOauthScopesNotFound {
	return &GetOauthScopesNotFound{}
}

/*
GetOauthScopesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOauthScopesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes not found response has a 2xx status code
func (o *GetOauthScopesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes not found response has a 3xx status code
func (o *GetOauthScopesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes not found response has a 4xx status code
func (o *GetOauthScopesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes not found response has a 5xx status code
func (o *GetOauthScopesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes not found response a status code equal to that given
func (o *GetOauthScopesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOauthScopesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesNotFound  %+v", 404, o.Payload)
}

func (o *GetOauthScopesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesNotFound  %+v", 404, o.Payload)
}

func (o *GetOauthScopesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesRequestTimeout creates a GetOauthScopesRequestTimeout with default headers values
func NewGetOauthScopesRequestTimeout() *GetOauthScopesRequestTimeout {
	return &GetOauthScopesRequestTimeout{}
}

/*
GetOauthScopesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOauthScopesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes request timeout response has a 2xx status code
func (o *GetOauthScopesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes request timeout response has a 3xx status code
func (o *GetOauthScopesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes request timeout response has a 4xx status code
func (o *GetOauthScopesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes request timeout response has a 5xx status code
func (o *GetOauthScopesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes request timeout response a status code equal to that given
func (o *GetOauthScopesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOauthScopesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOauthScopesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOauthScopesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesRequestEntityTooLarge creates a GetOauthScopesRequestEntityTooLarge with default headers values
func NewGetOauthScopesRequestEntityTooLarge() *GetOauthScopesRequestEntityTooLarge {
	return &GetOauthScopesRequestEntityTooLarge{}
}

/*
GetOauthScopesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOauthScopesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes request entity too large response has a 2xx status code
func (o *GetOauthScopesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes request entity too large response has a 3xx status code
func (o *GetOauthScopesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes request entity too large response has a 4xx status code
func (o *GetOauthScopesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes request entity too large response has a 5xx status code
func (o *GetOauthScopesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes request entity too large response a status code equal to that given
func (o *GetOauthScopesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOauthScopesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOauthScopesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOauthScopesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesUnsupportedMediaType creates a GetOauthScopesUnsupportedMediaType with default headers values
func NewGetOauthScopesUnsupportedMediaType() *GetOauthScopesUnsupportedMediaType {
	return &GetOauthScopesUnsupportedMediaType{}
}

/*
GetOauthScopesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOauthScopesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes unsupported media type response has a 2xx status code
func (o *GetOauthScopesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes unsupported media type response has a 3xx status code
func (o *GetOauthScopesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes unsupported media type response has a 4xx status code
func (o *GetOauthScopesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes unsupported media type response has a 5xx status code
func (o *GetOauthScopesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes unsupported media type response a status code equal to that given
func (o *GetOauthScopesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOauthScopesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOauthScopesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOauthScopesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesTooManyRequests creates a GetOauthScopesTooManyRequests with default headers values
func NewGetOauthScopesTooManyRequests() *GetOauthScopesTooManyRequests {
	return &GetOauthScopesTooManyRequests{}
}

/*
GetOauthScopesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOauthScopesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes too many requests response has a 2xx status code
func (o *GetOauthScopesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes too many requests response has a 3xx status code
func (o *GetOauthScopesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes too many requests response has a 4xx status code
func (o *GetOauthScopesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get oauth scopes too many requests response has a 5xx status code
func (o *GetOauthScopesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get oauth scopes too many requests response a status code equal to that given
func (o *GetOauthScopesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOauthScopesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOauthScopesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOauthScopesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesInternalServerError creates a GetOauthScopesInternalServerError with default headers values
func NewGetOauthScopesInternalServerError() *GetOauthScopesInternalServerError {
	return &GetOauthScopesInternalServerError{}
}

/*
GetOauthScopesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOauthScopesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes internal server error response has a 2xx status code
func (o *GetOauthScopesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes internal server error response has a 3xx status code
func (o *GetOauthScopesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes internal server error response has a 4xx status code
func (o *GetOauthScopesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get oauth scopes internal server error response has a 5xx status code
func (o *GetOauthScopesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get oauth scopes internal server error response a status code equal to that given
func (o *GetOauthScopesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOauthScopesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOauthScopesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOauthScopesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesServiceUnavailable creates a GetOauthScopesServiceUnavailable with default headers values
func NewGetOauthScopesServiceUnavailable() *GetOauthScopesServiceUnavailable {
	return &GetOauthScopesServiceUnavailable{}
}

/*
GetOauthScopesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOauthScopesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes service unavailable response has a 2xx status code
func (o *GetOauthScopesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes service unavailable response has a 3xx status code
func (o *GetOauthScopesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes service unavailable response has a 4xx status code
func (o *GetOauthScopesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get oauth scopes service unavailable response has a 5xx status code
func (o *GetOauthScopesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get oauth scopes service unavailable response a status code equal to that given
func (o *GetOauthScopesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOauthScopesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOauthScopesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOauthScopesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOauthScopesGatewayTimeout creates a GetOauthScopesGatewayTimeout with default headers values
func NewGetOauthScopesGatewayTimeout() *GetOauthScopesGatewayTimeout {
	return &GetOauthScopesGatewayTimeout{}
}

/*
GetOauthScopesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOauthScopesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get oauth scopes gateway timeout response has a 2xx status code
func (o *GetOauthScopesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get oauth scopes gateway timeout response has a 3xx status code
func (o *GetOauthScopesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get oauth scopes gateway timeout response has a 4xx status code
func (o *GetOauthScopesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get oauth scopes gateway timeout response has a 5xx status code
func (o *GetOauthScopesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get oauth scopes gateway timeout response a status code equal to that given
func (o *GetOauthScopesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOauthScopesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOauthScopesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/oauth/scopes][%d] getOauthScopesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOauthScopesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOauthScopesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
