// Code generated by go-swagger; DO NOT EDIT.

package scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScriptsPublishedReader is a Reader for the GetScriptsPublished structure.
type GetScriptsPublishedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsPublishedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsPublishedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsPublishedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsPublishedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsPublishedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsPublishedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScriptsPublishedRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsPublishedRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsPublishedUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsPublishedTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsPublishedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsPublishedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsPublishedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsPublishedOK creates a GetScriptsPublishedOK with default headers values
func NewGetScriptsPublishedOK() *GetScriptsPublishedOK {
	return &GetScriptsPublishedOK{}
}

/*
GetScriptsPublishedOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScriptsPublishedOK struct {
	Payload *models.ScriptEntityListing
}

// IsSuccess returns true when this get scripts published o k response has a 2xx status code
func (o *GetScriptsPublishedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scripts published o k response has a 3xx status code
func (o *GetScriptsPublishedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published o k response has a 4xx status code
func (o *GetScriptsPublishedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published o k response has a 5xx status code
func (o *GetScriptsPublishedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published o k response a status code equal to that given
func (o *GetScriptsPublishedOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScriptsPublishedOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedOK) GetPayload() *models.ScriptEntityListing {
	return o.Payload
}

func (o *GetScriptsPublishedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScriptEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedBadRequest creates a GetScriptsPublishedBadRequest with default headers values
func NewGetScriptsPublishedBadRequest() *GetScriptsPublishedBadRequest {
	return &GetScriptsPublishedBadRequest{}
}

/*
GetScriptsPublishedBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published bad request response has a 2xx status code
func (o *GetScriptsPublishedBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published bad request response has a 3xx status code
func (o *GetScriptsPublishedBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published bad request response has a 4xx status code
func (o *GetScriptsPublishedBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published bad request response has a 5xx status code
func (o *GetScriptsPublishedBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published bad request response a status code equal to that given
func (o *GetScriptsPublishedBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScriptsPublishedBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedUnauthorized creates a GetScriptsPublishedUnauthorized with default headers values
func NewGetScriptsPublishedUnauthorized() *GetScriptsPublishedUnauthorized {
	return &GetScriptsPublishedUnauthorized{}
}

/*
GetScriptsPublishedUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published unauthorized response has a 2xx status code
func (o *GetScriptsPublishedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published unauthorized response has a 3xx status code
func (o *GetScriptsPublishedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published unauthorized response has a 4xx status code
func (o *GetScriptsPublishedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published unauthorized response has a 5xx status code
func (o *GetScriptsPublishedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published unauthorized response a status code equal to that given
func (o *GetScriptsPublishedUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScriptsPublishedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedForbidden creates a GetScriptsPublishedForbidden with default headers values
func NewGetScriptsPublishedForbidden() *GetScriptsPublishedForbidden {
	return &GetScriptsPublishedForbidden{}
}

/*
GetScriptsPublishedForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published forbidden response has a 2xx status code
func (o *GetScriptsPublishedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published forbidden response has a 3xx status code
func (o *GetScriptsPublishedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published forbidden response has a 4xx status code
func (o *GetScriptsPublishedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published forbidden response has a 5xx status code
func (o *GetScriptsPublishedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published forbidden response a status code equal to that given
func (o *GetScriptsPublishedForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScriptsPublishedForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedNotFound creates a GetScriptsPublishedNotFound with default headers values
func NewGetScriptsPublishedNotFound() *GetScriptsPublishedNotFound {
	return &GetScriptsPublishedNotFound{}
}

/*
GetScriptsPublishedNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published not found response has a 2xx status code
func (o *GetScriptsPublishedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published not found response has a 3xx status code
func (o *GetScriptsPublishedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published not found response has a 4xx status code
func (o *GetScriptsPublishedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published not found response has a 5xx status code
func (o *GetScriptsPublishedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published not found response a status code equal to that given
func (o *GetScriptsPublishedNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScriptsPublishedNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedRequestTimeout creates a GetScriptsPublishedRequestTimeout with default headers values
func NewGetScriptsPublishedRequestTimeout() *GetScriptsPublishedRequestTimeout {
	return &GetScriptsPublishedRequestTimeout{}
}

/*
GetScriptsPublishedRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScriptsPublishedRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published request timeout response has a 2xx status code
func (o *GetScriptsPublishedRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published request timeout response has a 3xx status code
func (o *GetScriptsPublishedRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published request timeout response has a 4xx status code
func (o *GetScriptsPublishedRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published request timeout response has a 5xx status code
func (o *GetScriptsPublishedRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published request timeout response a status code equal to that given
func (o *GetScriptsPublishedRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScriptsPublishedRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedRequestEntityTooLarge creates a GetScriptsPublishedRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedRequestEntityTooLarge() *GetScriptsPublishedRequestEntityTooLarge {
	return &GetScriptsPublishedRequestEntityTooLarge{}
}

/*
GetScriptsPublishedRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetScriptsPublishedRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published request entity too large response has a 2xx status code
func (o *GetScriptsPublishedRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published request entity too large response has a 3xx status code
func (o *GetScriptsPublishedRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published request entity too large response has a 4xx status code
func (o *GetScriptsPublishedRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published request entity too large response has a 5xx status code
func (o *GetScriptsPublishedRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published request entity too large response a status code equal to that given
func (o *GetScriptsPublishedRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScriptsPublishedRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedUnsupportedMediaType creates a GetScriptsPublishedUnsupportedMediaType with default headers values
func NewGetScriptsPublishedUnsupportedMediaType() *GetScriptsPublishedUnsupportedMediaType {
	return &GetScriptsPublishedUnsupportedMediaType{}
}

/*
GetScriptsPublishedUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published unsupported media type response has a 2xx status code
func (o *GetScriptsPublishedUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published unsupported media type response has a 3xx status code
func (o *GetScriptsPublishedUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published unsupported media type response has a 4xx status code
func (o *GetScriptsPublishedUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published unsupported media type response has a 5xx status code
func (o *GetScriptsPublishedUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published unsupported media type response a status code equal to that given
func (o *GetScriptsPublishedUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScriptsPublishedUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedTooManyRequests creates a GetScriptsPublishedTooManyRequests with default headers values
func NewGetScriptsPublishedTooManyRequests() *GetScriptsPublishedTooManyRequests {
	return &GetScriptsPublishedTooManyRequests{}
}

/*
GetScriptsPublishedTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScriptsPublishedTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published too many requests response has a 2xx status code
func (o *GetScriptsPublishedTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published too many requests response has a 3xx status code
func (o *GetScriptsPublishedTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published too many requests response has a 4xx status code
func (o *GetScriptsPublishedTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published too many requests response has a 5xx status code
func (o *GetScriptsPublishedTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published too many requests response a status code equal to that given
func (o *GetScriptsPublishedTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScriptsPublishedTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedInternalServerError creates a GetScriptsPublishedInternalServerError with default headers values
func NewGetScriptsPublishedInternalServerError() *GetScriptsPublishedInternalServerError {
	return &GetScriptsPublishedInternalServerError{}
}

/*
GetScriptsPublishedInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published internal server error response has a 2xx status code
func (o *GetScriptsPublishedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published internal server error response has a 3xx status code
func (o *GetScriptsPublishedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published internal server error response has a 4xx status code
func (o *GetScriptsPublishedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published internal server error response has a 5xx status code
func (o *GetScriptsPublishedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published internal server error response a status code equal to that given
func (o *GetScriptsPublishedInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScriptsPublishedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedServiceUnavailable creates a GetScriptsPublishedServiceUnavailable with default headers values
func NewGetScriptsPublishedServiceUnavailable() *GetScriptsPublishedServiceUnavailable {
	return &GetScriptsPublishedServiceUnavailable{}
}

/*
GetScriptsPublishedServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published service unavailable response has a 2xx status code
func (o *GetScriptsPublishedServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published service unavailable response has a 3xx status code
func (o *GetScriptsPublishedServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published service unavailable response has a 4xx status code
func (o *GetScriptsPublishedServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published service unavailable response has a 5xx status code
func (o *GetScriptsPublishedServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published service unavailable response a status code equal to that given
func (o *GetScriptsPublishedServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScriptsPublishedServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedGatewayTimeout creates a GetScriptsPublishedGatewayTimeout with default headers values
func NewGetScriptsPublishedGatewayTimeout() *GetScriptsPublishedGatewayTimeout {
	return &GetScriptsPublishedGatewayTimeout{}
}

/*
GetScriptsPublishedGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScriptsPublishedGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published gateway timeout response has a 2xx status code
func (o *GetScriptsPublishedGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published gateway timeout response has a 3xx status code
func (o *GetScriptsPublishedGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published gateway timeout response has a 4xx status code
func (o *GetScriptsPublishedGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published gateway timeout response has a 5xx status code
func (o *GetScriptsPublishedGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published gateway timeout response a status code equal to that given
func (o *GetScriptsPublishedGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScriptsPublishedGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published][%d] getScriptsPublishedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
