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

// GetUserStationReader is a Reader for the GetUserStation structure.
type GetUserStationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserStationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserStationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserStationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserStationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserStationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserStationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserStationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserStationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserStationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 424:
		result := NewGetUserStationFailedDependency()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserStationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserStationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserStationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserStationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserStationOK creates a GetUserStationOK with default headers values
func NewGetUserStationOK() *GetUserStationOK {
	return &GetUserStationOK{}
}

/*
GetUserStationOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserStationOK struct {
	Payload *models.UserStations
}

// IsSuccess returns true when this get user station o k response has a 2xx status code
func (o *GetUserStationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user station o k response has a 3xx status code
func (o *GetUserStationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station o k response has a 4xx status code
func (o *GetUserStationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user station o k response has a 5xx status code
func (o *GetUserStationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station o k response a status code equal to that given
func (o *GetUserStationOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserStationOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationOK  %+v", 200, o.Payload)
}

func (o *GetUserStationOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationOK  %+v", 200, o.Payload)
}

func (o *GetUserStationOK) GetPayload() *models.UserStations {
	return o.Payload
}

func (o *GetUserStationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserStations)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationBadRequest creates a GetUserStationBadRequest with default headers values
func NewGetUserStationBadRequest() *GetUserStationBadRequest {
	return &GetUserStationBadRequest{}
}

/*
GetUserStationBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserStationBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station bad request response has a 2xx status code
func (o *GetUserStationBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station bad request response has a 3xx status code
func (o *GetUserStationBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station bad request response has a 4xx status code
func (o *GetUserStationBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station bad request response has a 5xx status code
func (o *GetUserStationBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station bad request response a status code equal to that given
func (o *GetUserStationBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserStationBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserStationBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserStationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationUnauthorized creates a GetUserStationUnauthorized with default headers values
func NewGetUserStationUnauthorized() *GetUserStationUnauthorized {
	return &GetUserStationUnauthorized{}
}

/*
GetUserStationUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserStationUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station unauthorized response has a 2xx status code
func (o *GetUserStationUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station unauthorized response has a 3xx status code
func (o *GetUserStationUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station unauthorized response has a 4xx status code
func (o *GetUserStationUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station unauthorized response has a 5xx status code
func (o *GetUserStationUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station unauthorized response a status code equal to that given
func (o *GetUserStationUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserStationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserStationUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserStationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationForbidden creates a GetUserStationForbidden with default headers values
func NewGetUserStationForbidden() *GetUserStationForbidden {
	return &GetUserStationForbidden{}
}

/*
GetUserStationForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserStationForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station forbidden response has a 2xx status code
func (o *GetUserStationForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station forbidden response has a 3xx status code
func (o *GetUserStationForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station forbidden response has a 4xx status code
func (o *GetUserStationForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station forbidden response has a 5xx status code
func (o *GetUserStationForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station forbidden response a status code equal to that given
func (o *GetUserStationForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserStationForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationForbidden  %+v", 403, o.Payload)
}

func (o *GetUserStationForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationForbidden  %+v", 403, o.Payload)
}

func (o *GetUserStationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationNotFound creates a GetUserStationNotFound with default headers values
func NewGetUserStationNotFound() *GetUserStationNotFound {
	return &GetUserStationNotFound{}
}

/*
GetUserStationNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserStationNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station not found response has a 2xx status code
func (o *GetUserStationNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station not found response has a 3xx status code
func (o *GetUserStationNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station not found response has a 4xx status code
func (o *GetUserStationNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station not found response has a 5xx status code
func (o *GetUserStationNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station not found response a status code equal to that given
func (o *GetUserStationNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserStationNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationNotFound  %+v", 404, o.Payload)
}

func (o *GetUserStationNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationNotFound  %+v", 404, o.Payload)
}

func (o *GetUserStationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationRequestTimeout creates a GetUserStationRequestTimeout with default headers values
func NewGetUserStationRequestTimeout() *GetUserStationRequestTimeout {
	return &GetUserStationRequestTimeout{}
}

/*
GetUserStationRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserStationRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station request timeout response has a 2xx status code
func (o *GetUserStationRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station request timeout response has a 3xx status code
func (o *GetUserStationRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station request timeout response has a 4xx status code
func (o *GetUserStationRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station request timeout response has a 5xx status code
func (o *GetUserStationRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station request timeout response a status code equal to that given
func (o *GetUserStationRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserStationRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserStationRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserStationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationRequestEntityTooLarge creates a GetUserStationRequestEntityTooLarge with default headers values
func NewGetUserStationRequestEntityTooLarge() *GetUserStationRequestEntityTooLarge {
	return &GetUserStationRequestEntityTooLarge{}
}

/*
GetUserStationRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserStationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station request entity too large response has a 2xx status code
func (o *GetUserStationRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station request entity too large response has a 3xx status code
func (o *GetUserStationRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station request entity too large response has a 4xx status code
func (o *GetUserStationRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station request entity too large response has a 5xx status code
func (o *GetUserStationRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station request entity too large response a status code equal to that given
func (o *GetUserStationRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserStationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserStationRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserStationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationUnsupportedMediaType creates a GetUserStationUnsupportedMediaType with default headers values
func NewGetUserStationUnsupportedMediaType() *GetUserStationUnsupportedMediaType {
	return &GetUserStationUnsupportedMediaType{}
}

/*
GetUserStationUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserStationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station unsupported media type response has a 2xx status code
func (o *GetUserStationUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station unsupported media type response has a 3xx status code
func (o *GetUserStationUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station unsupported media type response has a 4xx status code
func (o *GetUserStationUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station unsupported media type response has a 5xx status code
func (o *GetUserStationUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station unsupported media type response a status code equal to that given
func (o *GetUserStationUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserStationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserStationUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserStationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationFailedDependency creates a GetUserStationFailedDependency with default headers values
func NewGetUserStationFailedDependency() *GetUserStationFailedDependency {
	return &GetUserStationFailedDependency{}
}

/*
GetUserStationFailedDependency describes a response with status code 424, with default header values.

GetUserStationFailedDependency get user station failed dependency
*/
type GetUserStationFailedDependency struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station failed dependency response has a 2xx status code
func (o *GetUserStationFailedDependency) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station failed dependency response has a 3xx status code
func (o *GetUserStationFailedDependency) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station failed dependency response has a 4xx status code
func (o *GetUserStationFailedDependency) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station failed dependency response has a 5xx status code
func (o *GetUserStationFailedDependency) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station failed dependency response a status code equal to that given
func (o *GetUserStationFailedDependency) IsCode(code int) bool {
	return code == 424
}

func (o *GetUserStationFailedDependency) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationFailedDependency  %+v", 424, o.Payload)
}

func (o *GetUserStationFailedDependency) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationFailedDependency  %+v", 424, o.Payload)
}

func (o *GetUserStationFailedDependency) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationFailedDependency) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationTooManyRequests creates a GetUserStationTooManyRequests with default headers values
func NewGetUserStationTooManyRequests() *GetUserStationTooManyRequests {
	return &GetUserStationTooManyRequests{}
}

/*
GetUserStationTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserStationTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station too many requests response has a 2xx status code
func (o *GetUserStationTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station too many requests response has a 3xx status code
func (o *GetUserStationTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station too many requests response has a 4xx status code
func (o *GetUserStationTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user station too many requests response has a 5xx status code
func (o *GetUserStationTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user station too many requests response a status code equal to that given
func (o *GetUserStationTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserStationTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserStationTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserStationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationInternalServerError creates a GetUserStationInternalServerError with default headers values
func NewGetUserStationInternalServerError() *GetUserStationInternalServerError {
	return &GetUserStationInternalServerError{}
}

/*
GetUserStationInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserStationInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station internal server error response has a 2xx status code
func (o *GetUserStationInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station internal server error response has a 3xx status code
func (o *GetUserStationInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station internal server error response has a 4xx status code
func (o *GetUserStationInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user station internal server error response has a 5xx status code
func (o *GetUserStationInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user station internal server error response a status code equal to that given
func (o *GetUserStationInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserStationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserStationInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserStationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationServiceUnavailable creates a GetUserStationServiceUnavailable with default headers values
func NewGetUserStationServiceUnavailable() *GetUserStationServiceUnavailable {
	return &GetUserStationServiceUnavailable{}
}

/*
GetUserStationServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserStationServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station service unavailable response has a 2xx status code
func (o *GetUserStationServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station service unavailable response has a 3xx status code
func (o *GetUserStationServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station service unavailable response has a 4xx status code
func (o *GetUserStationServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user station service unavailable response has a 5xx status code
func (o *GetUserStationServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user station service unavailable response a status code equal to that given
func (o *GetUserStationServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserStationServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserStationServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserStationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserStationGatewayTimeout creates a GetUserStationGatewayTimeout with default headers values
func NewGetUserStationGatewayTimeout() *GetUserStationGatewayTimeout {
	return &GetUserStationGatewayTimeout{}
}

/*
GetUserStationGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserStationGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user station gateway timeout response has a 2xx status code
func (o *GetUserStationGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user station gateway timeout response has a 3xx status code
func (o *GetUserStationGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user station gateway timeout response has a 4xx status code
func (o *GetUserStationGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user station gateway timeout response has a 5xx status code
func (o *GetUserStationGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user station gateway timeout response a status code equal to that given
func (o *GetUserStationGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserStationGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserStationGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/station][%d] getUserStationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserStationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserStationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
