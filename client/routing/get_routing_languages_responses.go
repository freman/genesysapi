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

// GetRoutingLanguagesReader is a Reader for the GetRoutingLanguages structure.
type GetRoutingLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingLanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingLanguagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingLanguagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingLanguagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingLanguagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingLanguagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingLanguagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingLanguagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingLanguagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingLanguagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingLanguagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingLanguagesOK creates a GetRoutingLanguagesOK with default headers values
func NewGetRoutingLanguagesOK() *GetRoutingLanguagesOK {
	return &GetRoutingLanguagesOK{}
}

/*
GetRoutingLanguagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingLanguagesOK struct {
	Payload *models.LanguageEntityListing
}

// IsSuccess returns true when this get routing languages o k response has a 2xx status code
func (o *GetRoutingLanguagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing languages o k response has a 3xx status code
func (o *GetRoutingLanguagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages o k response has a 4xx status code
func (o *GetRoutingLanguagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing languages o k response has a 5xx status code
func (o *GetRoutingLanguagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages o k response a status code equal to that given
func (o *GetRoutingLanguagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingLanguagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingLanguagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesOK  %+v", 200, o.Payload)
}

func (o *GetRoutingLanguagesOK) GetPayload() *models.LanguageEntityListing {
	return o.Payload
}

func (o *GetRoutingLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LanguageEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesBadRequest creates a GetRoutingLanguagesBadRequest with default headers values
func NewGetRoutingLanguagesBadRequest() *GetRoutingLanguagesBadRequest {
	return &GetRoutingLanguagesBadRequest{}
}

/*
GetRoutingLanguagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingLanguagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages bad request response has a 2xx status code
func (o *GetRoutingLanguagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages bad request response has a 3xx status code
func (o *GetRoutingLanguagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages bad request response has a 4xx status code
func (o *GetRoutingLanguagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages bad request response has a 5xx status code
func (o *GetRoutingLanguagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages bad request response a status code equal to that given
func (o *GetRoutingLanguagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingLanguagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingLanguagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingLanguagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesUnauthorized creates a GetRoutingLanguagesUnauthorized with default headers values
func NewGetRoutingLanguagesUnauthorized() *GetRoutingLanguagesUnauthorized {
	return &GetRoutingLanguagesUnauthorized{}
}

/*
GetRoutingLanguagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingLanguagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages unauthorized response has a 2xx status code
func (o *GetRoutingLanguagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages unauthorized response has a 3xx status code
func (o *GetRoutingLanguagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages unauthorized response has a 4xx status code
func (o *GetRoutingLanguagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages unauthorized response has a 5xx status code
func (o *GetRoutingLanguagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages unauthorized response a status code equal to that given
func (o *GetRoutingLanguagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingLanguagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingLanguagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingLanguagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesForbidden creates a GetRoutingLanguagesForbidden with default headers values
func NewGetRoutingLanguagesForbidden() *GetRoutingLanguagesForbidden {
	return &GetRoutingLanguagesForbidden{}
}

/*
GetRoutingLanguagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingLanguagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages forbidden response has a 2xx status code
func (o *GetRoutingLanguagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages forbidden response has a 3xx status code
func (o *GetRoutingLanguagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages forbidden response has a 4xx status code
func (o *GetRoutingLanguagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages forbidden response has a 5xx status code
func (o *GetRoutingLanguagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages forbidden response a status code equal to that given
func (o *GetRoutingLanguagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingLanguagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingLanguagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingLanguagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesNotFound creates a GetRoutingLanguagesNotFound with default headers values
func NewGetRoutingLanguagesNotFound() *GetRoutingLanguagesNotFound {
	return &GetRoutingLanguagesNotFound{}
}

/*
GetRoutingLanguagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingLanguagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages not found response has a 2xx status code
func (o *GetRoutingLanguagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages not found response has a 3xx status code
func (o *GetRoutingLanguagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages not found response has a 4xx status code
func (o *GetRoutingLanguagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages not found response has a 5xx status code
func (o *GetRoutingLanguagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages not found response a status code equal to that given
func (o *GetRoutingLanguagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingLanguagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingLanguagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingLanguagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesRequestTimeout creates a GetRoutingLanguagesRequestTimeout with default headers values
func NewGetRoutingLanguagesRequestTimeout() *GetRoutingLanguagesRequestTimeout {
	return &GetRoutingLanguagesRequestTimeout{}
}

/*
GetRoutingLanguagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingLanguagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages request timeout response has a 2xx status code
func (o *GetRoutingLanguagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages request timeout response has a 3xx status code
func (o *GetRoutingLanguagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages request timeout response has a 4xx status code
func (o *GetRoutingLanguagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages request timeout response has a 5xx status code
func (o *GetRoutingLanguagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages request timeout response a status code equal to that given
func (o *GetRoutingLanguagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingLanguagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingLanguagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingLanguagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesRequestEntityTooLarge creates a GetRoutingLanguagesRequestEntityTooLarge with default headers values
func NewGetRoutingLanguagesRequestEntityTooLarge() *GetRoutingLanguagesRequestEntityTooLarge {
	return &GetRoutingLanguagesRequestEntityTooLarge{}
}

/*
GetRoutingLanguagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingLanguagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages request entity too large response has a 2xx status code
func (o *GetRoutingLanguagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages request entity too large response has a 3xx status code
func (o *GetRoutingLanguagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages request entity too large response has a 4xx status code
func (o *GetRoutingLanguagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages request entity too large response has a 5xx status code
func (o *GetRoutingLanguagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages request entity too large response a status code equal to that given
func (o *GetRoutingLanguagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingLanguagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingLanguagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingLanguagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesUnsupportedMediaType creates a GetRoutingLanguagesUnsupportedMediaType with default headers values
func NewGetRoutingLanguagesUnsupportedMediaType() *GetRoutingLanguagesUnsupportedMediaType {
	return &GetRoutingLanguagesUnsupportedMediaType{}
}

/*
GetRoutingLanguagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingLanguagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages unsupported media type response has a 2xx status code
func (o *GetRoutingLanguagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages unsupported media type response has a 3xx status code
func (o *GetRoutingLanguagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages unsupported media type response has a 4xx status code
func (o *GetRoutingLanguagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages unsupported media type response has a 5xx status code
func (o *GetRoutingLanguagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages unsupported media type response a status code equal to that given
func (o *GetRoutingLanguagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingLanguagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingLanguagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingLanguagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesTooManyRequests creates a GetRoutingLanguagesTooManyRequests with default headers values
func NewGetRoutingLanguagesTooManyRequests() *GetRoutingLanguagesTooManyRequests {
	return &GetRoutingLanguagesTooManyRequests{}
}

/*
GetRoutingLanguagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingLanguagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages too many requests response has a 2xx status code
func (o *GetRoutingLanguagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages too many requests response has a 3xx status code
func (o *GetRoutingLanguagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages too many requests response has a 4xx status code
func (o *GetRoutingLanguagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing languages too many requests response has a 5xx status code
func (o *GetRoutingLanguagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing languages too many requests response a status code equal to that given
func (o *GetRoutingLanguagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingLanguagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingLanguagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingLanguagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesInternalServerError creates a GetRoutingLanguagesInternalServerError with default headers values
func NewGetRoutingLanguagesInternalServerError() *GetRoutingLanguagesInternalServerError {
	return &GetRoutingLanguagesInternalServerError{}
}

/*
GetRoutingLanguagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingLanguagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages internal server error response has a 2xx status code
func (o *GetRoutingLanguagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages internal server error response has a 3xx status code
func (o *GetRoutingLanguagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages internal server error response has a 4xx status code
func (o *GetRoutingLanguagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing languages internal server error response has a 5xx status code
func (o *GetRoutingLanguagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing languages internal server error response a status code equal to that given
func (o *GetRoutingLanguagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingLanguagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingLanguagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingLanguagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesServiceUnavailable creates a GetRoutingLanguagesServiceUnavailable with default headers values
func NewGetRoutingLanguagesServiceUnavailable() *GetRoutingLanguagesServiceUnavailable {
	return &GetRoutingLanguagesServiceUnavailable{}
}

/*
GetRoutingLanguagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingLanguagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages service unavailable response has a 2xx status code
func (o *GetRoutingLanguagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages service unavailable response has a 3xx status code
func (o *GetRoutingLanguagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages service unavailable response has a 4xx status code
func (o *GetRoutingLanguagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing languages service unavailable response has a 5xx status code
func (o *GetRoutingLanguagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing languages service unavailable response a status code equal to that given
func (o *GetRoutingLanguagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingLanguagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingLanguagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingLanguagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingLanguagesGatewayTimeout creates a GetRoutingLanguagesGatewayTimeout with default headers values
func NewGetRoutingLanguagesGatewayTimeout() *GetRoutingLanguagesGatewayTimeout {
	return &GetRoutingLanguagesGatewayTimeout{}
}

/*
GetRoutingLanguagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingLanguagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing languages gateway timeout response has a 2xx status code
func (o *GetRoutingLanguagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing languages gateway timeout response has a 3xx status code
func (o *GetRoutingLanguagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing languages gateway timeout response has a 4xx status code
func (o *GetRoutingLanguagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing languages gateway timeout response has a 5xx status code
func (o *GetRoutingLanguagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing languages gateway timeout response a status code equal to that given
func (o *GetRoutingLanguagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingLanguagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingLanguagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/languages][%d] getRoutingLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingLanguagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingLanguagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
