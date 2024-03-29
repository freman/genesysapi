// Code generated by go-swagger; DO NOT EDIT.

package user_recordings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserrecordingsReader is a Reader for the GetUserrecordings structure.
type GetUserrecordingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserrecordingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserrecordingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserrecordingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserrecordingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserrecordingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserrecordingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserrecordingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserrecordingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserrecordingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserrecordingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserrecordingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserrecordingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserrecordingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserrecordingsOK creates a GetUserrecordingsOK with default headers values
func NewGetUserrecordingsOK() *GetUserrecordingsOK {
	return &GetUserrecordingsOK{}
}

/*
GetUserrecordingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserrecordingsOK struct {
	Payload *models.UserRecordingEntityListing
}

// IsSuccess returns true when this get userrecordings o k response has a 2xx status code
func (o *GetUserrecordingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get userrecordings o k response has a 3xx status code
func (o *GetUserrecordingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings o k response has a 4xx status code
func (o *GetUserrecordingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get userrecordings o k response has a 5xx status code
func (o *GetUserrecordingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings o k response a status code equal to that given
func (o *GetUserrecordingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserrecordingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsOK  %+v", 200, o.Payload)
}

func (o *GetUserrecordingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsOK  %+v", 200, o.Payload)
}

func (o *GetUserrecordingsOK) GetPayload() *models.UserRecordingEntityListing {
	return o.Payload
}

func (o *GetUserrecordingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRecordingEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsBadRequest creates a GetUserrecordingsBadRequest with default headers values
func NewGetUserrecordingsBadRequest() *GetUserrecordingsBadRequest {
	return &GetUserrecordingsBadRequest{}
}

/*
GetUserrecordingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserrecordingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings bad request response has a 2xx status code
func (o *GetUserrecordingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings bad request response has a 3xx status code
func (o *GetUserrecordingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings bad request response has a 4xx status code
func (o *GetUserrecordingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings bad request response has a 5xx status code
func (o *GetUserrecordingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings bad request response a status code equal to that given
func (o *GetUserrecordingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserrecordingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserrecordingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserrecordingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsUnauthorized creates a GetUserrecordingsUnauthorized with default headers values
func NewGetUserrecordingsUnauthorized() *GetUserrecordingsUnauthorized {
	return &GetUserrecordingsUnauthorized{}
}

/*
GetUserrecordingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserrecordingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings unauthorized response has a 2xx status code
func (o *GetUserrecordingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings unauthorized response has a 3xx status code
func (o *GetUserrecordingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings unauthorized response has a 4xx status code
func (o *GetUserrecordingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings unauthorized response has a 5xx status code
func (o *GetUserrecordingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings unauthorized response a status code equal to that given
func (o *GetUserrecordingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserrecordingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserrecordingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserrecordingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsForbidden creates a GetUserrecordingsForbidden with default headers values
func NewGetUserrecordingsForbidden() *GetUserrecordingsForbidden {
	return &GetUserrecordingsForbidden{}
}

/*
GetUserrecordingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserrecordingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings forbidden response has a 2xx status code
func (o *GetUserrecordingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings forbidden response has a 3xx status code
func (o *GetUserrecordingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings forbidden response has a 4xx status code
func (o *GetUserrecordingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings forbidden response has a 5xx status code
func (o *GetUserrecordingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings forbidden response a status code equal to that given
func (o *GetUserrecordingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserrecordingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserrecordingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserrecordingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsNotFound creates a GetUserrecordingsNotFound with default headers values
func NewGetUserrecordingsNotFound() *GetUserrecordingsNotFound {
	return &GetUserrecordingsNotFound{}
}

/*
GetUserrecordingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserrecordingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings not found response has a 2xx status code
func (o *GetUserrecordingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings not found response has a 3xx status code
func (o *GetUserrecordingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings not found response has a 4xx status code
func (o *GetUserrecordingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings not found response has a 5xx status code
func (o *GetUserrecordingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings not found response a status code equal to that given
func (o *GetUserrecordingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserrecordingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserrecordingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserrecordingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsRequestTimeout creates a GetUserrecordingsRequestTimeout with default headers values
func NewGetUserrecordingsRequestTimeout() *GetUserrecordingsRequestTimeout {
	return &GetUserrecordingsRequestTimeout{}
}

/*
GetUserrecordingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserrecordingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings request timeout response has a 2xx status code
func (o *GetUserrecordingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings request timeout response has a 3xx status code
func (o *GetUserrecordingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings request timeout response has a 4xx status code
func (o *GetUserrecordingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings request timeout response has a 5xx status code
func (o *GetUserrecordingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings request timeout response a status code equal to that given
func (o *GetUserrecordingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserrecordingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserrecordingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserrecordingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsRequestEntityTooLarge creates a GetUserrecordingsRequestEntityTooLarge with default headers values
func NewGetUserrecordingsRequestEntityTooLarge() *GetUserrecordingsRequestEntityTooLarge {
	return &GetUserrecordingsRequestEntityTooLarge{}
}

/*
GetUserrecordingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserrecordingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings request entity too large response has a 2xx status code
func (o *GetUserrecordingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings request entity too large response has a 3xx status code
func (o *GetUserrecordingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings request entity too large response has a 4xx status code
func (o *GetUserrecordingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings request entity too large response has a 5xx status code
func (o *GetUserrecordingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings request entity too large response a status code equal to that given
func (o *GetUserrecordingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserrecordingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserrecordingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserrecordingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsUnsupportedMediaType creates a GetUserrecordingsUnsupportedMediaType with default headers values
func NewGetUserrecordingsUnsupportedMediaType() *GetUserrecordingsUnsupportedMediaType {
	return &GetUserrecordingsUnsupportedMediaType{}
}

/*
GetUserrecordingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserrecordingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings unsupported media type response has a 2xx status code
func (o *GetUserrecordingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings unsupported media type response has a 3xx status code
func (o *GetUserrecordingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings unsupported media type response has a 4xx status code
func (o *GetUserrecordingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings unsupported media type response has a 5xx status code
func (o *GetUserrecordingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings unsupported media type response a status code equal to that given
func (o *GetUserrecordingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserrecordingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserrecordingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserrecordingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsTooManyRequests creates a GetUserrecordingsTooManyRequests with default headers values
func NewGetUserrecordingsTooManyRequests() *GetUserrecordingsTooManyRequests {
	return &GetUserrecordingsTooManyRequests{}
}

/*
GetUserrecordingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserrecordingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings too many requests response has a 2xx status code
func (o *GetUserrecordingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings too many requests response has a 3xx status code
func (o *GetUserrecordingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings too many requests response has a 4xx status code
func (o *GetUserrecordingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get userrecordings too many requests response has a 5xx status code
func (o *GetUserrecordingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get userrecordings too many requests response a status code equal to that given
func (o *GetUserrecordingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserrecordingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserrecordingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserrecordingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsInternalServerError creates a GetUserrecordingsInternalServerError with default headers values
func NewGetUserrecordingsInternalServerError() *GetUserrecordingsInternalServerError {
	return &GetUserrecordingsInternalServerError{}
}

/*
GetUserrecordingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserrecordingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings internal server error response has a 2xx status code
func (o *GetUserrecordingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings internal server error response has a 3xx status code
func (o *GetUserrecordingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings internal server error response has a 4xx status code
func (o *GetUserrecordingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get userrecordings internal server error response has a 5xx status code
func (o *GetUserrecordingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get userrecordings internal server error response a status code equal to that given
func (o *GetUserrecordingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserrecordingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserrecordingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserrecordingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsServiceUnavailable creates a GetUserrecordingsServiceUnavailable with default headers values
func NewGetUserrecordingsServiceUnavailable() *GetUserrecordingsServiceUnavailable {
	return &GetUserrecordingsServiceUnavailable{}
}

/*
GetUserrecordingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserrecordingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings service unavailable response has a 2xx status code
func (o *GetUserrecordingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings service unavailable response has a 3xx status code
func (o *GetUserrecordingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings service unavailable response has a 4xx status code
func (o *GetUserrecordingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get userrecordings service unavailable response has a 5xx status code
func (o *GetUserrecordingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get userrecordings service unavailable response a status code equal to that given
func (o *GetUserrecordingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserrecordingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserrecordingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserrecordingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserrecordingsGatewayTimeout creates a GetUserrecordingsGatewayTimeout with default headers values
func NewGetUserrecordingsGatewayTimeout() *GetUserrecordingsGatewayTimeout {
	return &GetUserrecordingsGatewayTimeout{}
}

/*
GetUserrecordingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserrecordingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get userrecordings gateway timeout response has a 2xx status code
func (o *GetUserrecordingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get userrecordings gateway timeout response has a 3xx status code
func (o *GetUserrecordingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get userrecordings gateway timeout response has a 4xx status code
func (o *GetUserrecordingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get userrecordings gateway timeout response has a 5xx status code
func (o *GetUserrecordingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get userrecordings gateway timeout response a status code equal to that given
func (o *GetUserrecordingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserrecordingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserrecordingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/userrecordings][%d] getUserrecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserrecordingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserrecordingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
