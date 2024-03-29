// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetLanguagesReader is a Reader for the GetLanguages structure.
type GetLanguagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLanguagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLanguagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLanguagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLanguagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLanguagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetLanguagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetLanguagesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetLanguagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetLanguagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetLanguagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLanguagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetLanguagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetLanguagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLanguagesOK creates a GetLanguagesOK with default headers values
func NewGetLanguagesOK() *GetLanguagesOK {
	return &GetLanguagesOK{}
}

/*
GetLanguagesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetLanguagesOK struct {
	Payload *models.LanguageEntityListing
}

// IsSuccess returns true when this get languages o k response has a 2xx status code
func (o *GetLanguagesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get languages o k response has a 3xx status code
func (o *GetLanguagesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages o k response has a 4xx status code
func (o *GetLanguagesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages o k response has a 5xx status code
func (o *GetLanguagesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages o k response a status code equal to that given
func (o *GetLanguagesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLanguagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesOK  %+v", 200, o.Payload)
}

func (o *GetLanguagesOK) GetPayload() *models.LanguageEntityListing {
	return o.Payload
}

func (o *GetLanguagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LanguageEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesBadRequest creates a GetLanguagesBadRequest with default headers values
func NewGetLanguagesBadRequest() *GetLanguagesBadRequest {
	return &GetLanguagesBadRequest{}
}

/*
GetLanguagesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetLanguagesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages bad request response has a 2xx status code
func (o *GetLanguagesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages bad request response has a 3xx status code
func (o *GetLanguagesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages bad request response has a 4xx status code
func (o *GetLanguagesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages bad request response has a 5xx status code
func (o *GetLanguagesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages bad request response a status code equal to that given
func (o *GetLanguagesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetLanguagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguagesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetLanguagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesUnauthorized creates a GetLanguagesUnauthorized with default headers values
func NewGetLanguagesUnauthorized() *GetLanguagesUnauthorized {
	return &GetLanguagesUnauthorized{}
}

/*
GetLanguagesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetLanguagesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages unauthorized response has a 2xx status code
func (o *GetLanguagesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages unauthorized response has a 3xx status code
func (o *GetLanguagesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages unauthorized response has a 4xx status code
func (o *GetLanguagesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages unauthorized response has a 5xx status code
func (o *GetLanguagesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages unauthorized response a status code equal to that given
func (o *GetLanguagesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLanguagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLanguagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesForbidden creates a GetLanguagesForbidden with default headers values
func NewGetLanguagesForbidden() *GetLanguagesForbidden {
	return &GetLanguagesForbidden{}
}

/*
GetLanguagesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetLanguagesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages forbidden response has a 2xx status code
func (o *GetLanguagesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages forbidden response has a 3xx status code
func (o *GetLanguagesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages forbidden response has a 4xx status code
func (o *GetLanguagesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages forbidden response has a 5xx status code
func (o *GetLanguagesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages forbidden response a status code equal to that given
func (o *GetLanguagesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLanguagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguagesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesForbidden  %+v", 403, o.Payload)
}

func (o *GetLanguagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesNotFound creates a GetLanguagesNotFound with default headers values
func NewGetLanguagesNotFound() *GetLanguagesNotFound {
	return &GetLanguagesNotFound{}
}

/*
GetLanguagesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetLanguagesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages not found response has a 2xx status code
func (o *GetLanguagesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages not found response has a 3xx status code
func (o *GetLanguagesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages not found response has a 4xx status code
func (o *GetLanguagesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages not found response has a 5xx status code
func (o *GetLanguagesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages not found response a status code equal to that given
func (o *GetLanguagesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLanguagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesNotFound  %+v", 404, o.Payload)
}

func (o *GetLanguagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesRequestTimeout creates a GetLanguagesRequestTimeout with default headers values
func NewGetLanguagesRequestTimeout() *GetLanguagesRequestTimeout {
	return &GetLanguagesRequestTimeout{}
}

/*
GetLanguagesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetLanguagesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages request timeout response has a 2xx status code
func (o *GetLanguagesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages request timeout response has a 3xx status code
func (o *GetLanguagesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages request timeout response has a 4xx status code
func (o *GetLanguagesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages request timeout response has a 5xx status code
func (o *GetLanguagesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages request timeout response a status code equal to that given
func (o *GetLanguagesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetLanguagesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguagesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetLanguagesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesRequestEntityTooLarge creates a GetLanguagesRequestEntityTooLarge with default headers values
func NewGetLanguagesRequestEntityTooLarge() *GetLanguagesRequestEntityTooLarge {
	return &GetLanguagesRequestEntityTooLarge{}
}

/*
GetLanguagesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetLanguagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages request entity too large response has a 2xx status code
func (o *GetLanguagesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages request entity too large response has a 3xx status code
func (o *GetLanguagesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages request entity too large response has a 4xx status code
func (o *GetLanguagesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages request entity too large response has a 5xx status code
func (o *GetLanguagesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages request entity too large response a status code equal to that given
func (o *GetLanguagesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetLanguagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguagesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetLanguagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesUnsupportedMediaType creates a GetLanguagesUnsupportedMediaType with default headers values
func NewGetLanguagesUnsupportedMediaType() *GetLanguagesUnsupportedMediaType {
	return &GetLanguagesUnsupportedMediaType{}
}

/*
GetLanguagesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetLanguagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages unsupported media type response has a 2xx status code
func (o *GetLanguagesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages unsupported media type response has a 3xx status code
func (o *GetLanguagesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages unsupported media type response has a 4xx status code
func (o *GetLanguagesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages unsupported media type response has a 5xx status code
func (o *GetLanguagesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages unsupported media type response a status code equal to that given
func (o *GetLanguagesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetLanguagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguagesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetLanguagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesTooManyRequests creates a GetLanguagesTooManyRequests with default headers values
func NewGetLanguagesTooManyRequests() *GetLanguagesTooManyRequests {
	return &GetLanguagesTooManyRequests{}
}

/*
GetLanguagesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetLanguagesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages too many requests response has a 2xx status code
func (o *GetLanguagesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages too many requests response has a 3xx status code
func (o *GetLanguagesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages too many requests response has a 4xx status code
func (o *GetLanguagesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get languages too many requests response has a 5xx status code
func (o *GetLanguagesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get languages too many requests response a status code equal to that given
func (o *GetLanguagesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetLanguagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguagesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetLanguagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesInternalServerError creates a GetLanguagesInternalServerError with default headers values
func NewGetLanguagesInternalServerError() *GetLanguagesInternalServerError {
	return &GetLanguagesInternalServerError{}
}

/*
GetLanguagesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetLanguagesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages internal server error response has a 2xx status code
func (o *GetLanguagesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages internal server error response has a 3xx status code
func (o *GetLanguagesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages internal server error response has a 4xx status code
func (o *GetLanguagesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages internal server error response has a 5xx status code
func (o *GetLanguagesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages internal server error response a status code equal to that given
func (o *GetLanguagesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLanguagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguagesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLanguagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesServiceUnavailable creates a GetLanguagesServiceUnavailable with default headers values
func NewGetLanguagesServiceUnavailable() *GetLanguagesServiceUnavailable {
	return &GetLanguagesServiceUnavailable{}
}

/*
GetLanguagesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetLanguagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages service unavailable response has a 2xx status code
func (o *GetLanguagesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages service unavailable response has a 3xx status code
func (o *GetLanguagesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages service unavailable response has a 4xx status code
func (o *GetLanguagesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages service unavailable response has a 5xx status code
func (o *GetLanguagesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages service unavailable response a status code equal to that given
func (o *GetLanguagesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetLanguagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguagesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetLanguagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLanguagesGatewayTimeout creates a GetLanguagesGatewayTimeout with default headers values
func NewGetLanguagesGatewayTimeout() *GetLanguagesGatewayTimeout {
	return &GetLanguagesGatewayTimeout{}
}

/*
GetLanguagesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetLanguagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get languages gateway timeout response has a 2xx status code
func (o *GetLanguagesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get languages gateway timeout response has a 3xx status code
func (o *GetLanguagesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get languages gateway timeout response has a 4xx status code
func (o *GetLanguagesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get languages gateway timeout response has a 5xx status code
func (o *GetLanguagesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get languages gateway timeout response a status code equal to that given
func (o *GetLanguagesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetLanguagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguagesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/languages][%d] getLanguagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetLanguagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetLanguagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
