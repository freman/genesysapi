// Code generated by go-swagger; DO NOT EDIT.

package geolocation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserGeolocationReader is a Reader for the GetUserGeolocation structure.
type GetUserGeolocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserGeolocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserGeolocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserGeolocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserGeolocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserGeolocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserGeolocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserGeolocationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserGeolocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserGeolocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserGeolocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserGeolocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserGeolocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserGeolocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserGeolocationOK creates a GetUserGeolocationOK with default headers values
func NewGetUserGeolocationOK() *GetUserGeolocationOK {
	return &GetUserGeolocationOK{}
}

/*
GetUserGeolocationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserGeolocationOK struct {
	Payload *models.Geolocation
}

// IsSuccess returns true when this get user geolocation o k response has a 2xx status code
func (o *GetUserGeolocationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user geolocation o k response has a 3xx status code
func (o *GetUserGeolocationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation o k response has a 4xx status code
func (o *GetUserGeolocationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user geolocation o k response has a 5xx status code
func (o *GetUserGeolocationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation o k response a status code equal to that given
func (o *GetUserGeolocationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserGeolocationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationOK  %+v", 200, o.Payload)
}

func (o *GetUserGeolocationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationOK  %+v", 200, o.Payload)
}

func (o *GetUserGeolocationOK) GetPayload() *models.Geolocation {
	return o.Payload
}

func (o *GetUserGeolocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Geolocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationBadRequest creates a GetUserGeolocationBadRequest with default headers values
func NewGetUserGeolocationBadRequest() *GetUserGeolocationBadRequest {
	return &GetUserGeolocationBadRequest{}
}

/*
GetUserGeolocationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserGeolocationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation bad request response has a 2xx status code
func (o *GetUserGeolocationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation bad request response has a 3xx status code
func (o *GetUserGeolocationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation bad request response has a 4xx status code
func (o *GetUserGeolocationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation bad request response has a 5xx status code
func (o *GetUserGeolocationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation bad request response a status code equal to that given
func (o *GetUserGeolocationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserGeolocationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGeolocationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserGeolocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationUnauthorized creates a GetUserGeolocationUnauthorized with default headers values
func NewGetUserGeolocationUnauthorized() *GetUserGeolocationUnauthorized {
	return &GetUserGeolocationUnauthorized{}
}

/*
GetUserGeolocationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserGeolocationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation unauthorized response has a 2xx status code
func (o *GetUserGeolocationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation unauthorized response has a 3xx status code
func (o *GetUserGeolocationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation unauthorized response has a 4xx status code
func (o *GetUserGeolocationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation unauthorized response has a 5xx status code
func (o *GetUserGeolocationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation unauthorized response a status code equal to that given
func (o *GetUserGeolocationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserGeolocationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGeolocationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserGeolocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationForbidden creates a GetUserGeolocationForbidden with default headers values
func NewGetUserGeolocationForbidden() *GetUserGeolocationForbidden {
	return &GetUserGeolocationForbidden{}
}

/*
GetUserGeolocationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserGeolocationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation forbidden response has a 2xx status code
func (o *GetUserGeolocationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation forbidden response has a 3xx status code
func (o *GetUserGeolocationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation forbidden response has a 4xx status code
func (o *GetUserGeolocationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation forbidden response has a 5xx status code
func (o *GetUserGeolocationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation forbidden response a status code equal to that given
func (o *GetUserGeolocationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserGeolocationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGeolocationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationForbidden  %+v", 403, o.Payload)
}

func (o *GetUserGeolocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationNotFound creates a GetUserGeolocationNotFound with default headers values
func NewGetUserGeolocationNotFound() *GetUserGeolocationNotFound {
	return &GetUserGeolocationNotFound{}
}

/*
GetUserGeolocationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserGeolocationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation not found response has a 2xx status code
func (o *GetUserGeolocationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation not found response has a 3xx status code
func (o *GetUserGeolocationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation not found response has a 4xx status code
func (o *GetUserGeolocationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation not found response has a 5xx status code
func (o *GetUserGeolocationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation not found response a status code equal to that given
func (o *GetUserGeolocationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserGeolocationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGeolocationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationNotFound  %+v", 404, o.Payload)
}

func (o *GetUserGeolocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationRequestTimeout creates a GetUserGeolocationRequestTimeout with default headers values
func NewGetUserGeolocationRequestTimeout() *GetUserGeolocationRequestTimeout {
	return &GetUserGeolocationRequestTimeout{}
}

/*
GetUserGeolocationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserGeolocationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation request timeout response has a 2xx status code
func (o *GetUserGeolocationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation request timeout response has a 3xx status code
func (o *GetUserGeolocationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation request timeout response has a 4xx status code
func (o *GetUserGeolocationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation request timeout response has a 5xx status code
func (o *GetUserGeolocationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation request timeout response a status code equal to that given
func (o *GetUserGeolocationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserGeolocationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGeolocationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserGeolocationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationRequestEntityTooLarge creates a GetUserGeolocationRequestEntityTooLarge with default headers values
func NewGetUserGeolocationRequestEntityTooLarge() *GetUserGeolocationRequestEntityTooLarge {
	return &GetUserGeolocationRequestEntityTooLarge{}
}

/*
GetUserGeolocationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetUserGeolocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation request entity too large response has a 2xx status code
func (o *GetUserGeolocationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation request entity too large response has a 3xx status code
func (o *GetUserGeolocationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation request entity too large response has a 4xx status code
func (o *GetUserGeolocationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation request entity too large response has a 5xx status code
func (o *GetUserGeolocationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation request entity too large response a status code equal to that given
func (o *GetUserGeolocationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserGeolocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGeolocationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserGeolocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationUnsupportedMediaType creates a GetUserGeolocationUnsupportedMediaType with default headers values
func NewGetUserGeolocationUnsupportedMediaType() *GetUserGeolocationUnsupportedMediaType {
	return &GetUserGeolocationUnsupportedMediaType{}
}

/*
GetUserGeolocationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserGeolocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation unsupported media type response has a 2xx status code
func (o *GetUserGeolocationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation unsupported media type response has a 3xx status code
func (o *GetUserGeolocationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation unsupported media type response has a 4xx status code
func (o *GetUserGeolocationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation unsupported media type response has a 5xx status code
func (o *GetUserGeolocationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation unsupported media type response a status code equal to that given
func (o *GetUserGeolocationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserGeolocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGeolocationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserGeolocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationTooManyRequests creates a GetUserGeolocationTooManyRequests with default headers values
func NewGetUserGeolocationTooManyRequests() *GetUserGeolocationTooManyRequests {
	return &GetUserGeolocationTooManyRequests{}
}

/*
GetUserGeolocationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserGeolocationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation too many requests response has a 2xx status code
func (o *GetUserGeolocationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation too many requests response has a 3xx status code
func (o *GetUserGeolocationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation too many requests response has a 4xx status code
func (o *GetUserGeolocationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user geolocation too many requests response has a 5xx status code
func (o *GetUserGeolocationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user geolocation too many requests response a status code equal to that given
func (o *GetUserGeolocationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserGeolocationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGeolocationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserGeolocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationInternalServerError creates a GetUserGeolocationInternalServerError with default headers values
func NewGetUserGeolocationInternalServerError() *GetUserGeolocationInternalServerError {
	return &GetUserGeolocationInternalServerError{}
}

/*
GetUserGeolocationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserGeolocationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation internal server error response has a 2xx status code
func (o *GetUserGeolocationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation internal server error response has a 3xx status code
func (o *GetUserGeolocationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation internal server error response has a 4xx status code
func (o *GetUserGeolocationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user geolocation internal server error response has a 5xx status code
func (o *GetUserGeolocationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user geolocation internal server error response a status code equal to that given
func (o *GetUserGeolocationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserGeolocationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGeolocationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserGeolocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationServiceUnavailable creates a GetUserGeolocationServiceUnavailable with default headers values
func NewGetUserGeolocationServiceUnavailable() *GetUserGeolocationServiceUnavailable {
	return &GetUserGeolocationServiceUnavailable{}
}

/*
GetUserGeolocationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserGeolocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation service unavailable response has a 2xx status code
func (o *GetUserGeolocationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation service unavailable response has a 3xx status code
func (o *GetUserGeolocationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation service unavailable response has a 4xx status code
func (o *GetUserGeolocationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user geolocation service unavailable response has a 5xx status code
func (o *GetUserGeolocationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user geolocation service unavailable response a status code equal to that given
func (o *GetUserGeolocationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserGeolocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGeolocationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserGeolocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserGeolocationGatewayTimeout creates a GetUserGeolocationGatewayTimeout with default headers values
func NewGetUserGeolocationGatewayTimeout() *GetUserGeolocationGatewayTimeout {
	return &GetUserGeolocationGatewayTimeout{}
}

/*
GetUserGeolocationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserGeolocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user geolocation gateway timeout response has a 2xx status code
func (o *GetUserGeolocationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user geolocation gateway timeout response has a 3xx status code
func (o *GetUserGeolocationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user geolocation gateway timeout response has a 4xx status code
func (o *GetUserGeolocationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user geolocation gateway timeout response has a 5xx status code
func (o *GetUserGeolocationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user geolocation gateway timeout response a status code equal to that given
func (o *GetUserGeolocationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserGeolocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGeolocationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/geolocations/{clientId}][%d] getUserGeolocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserGeolocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserGeolocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
