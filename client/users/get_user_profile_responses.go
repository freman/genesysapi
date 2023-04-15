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

// GetUserProfileReader is a Reader for the GetUserProfile structure.
type GetUserProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetUserProfileRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserProfileRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserProfileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserProfileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserProfileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserProfileOK creates a GetUserProfileOK with default headers values
func NewGetUserProfileOK() *GetUserProfileOK {
	return &GetUserProfileOK{}
}

/*
GetUserProfileOK describes a response with status code 200, with default header values.

successful operation
*/
type GetUserProfileOK struct {
	Payload *models.UserProfile
}

// IsSuccess returns true when this get user profile o k response has a 2xx status code
func (o *GetUserProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get user profile o k response has a 3xx status code
func (o *GetUserProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile o k response has a 4xx status code
func (o *GetUserProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user profile o k response has a 5xx status code
func (o *GetUserProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile o k response a status code equal to that given
func (o *GetUserProfileOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetUserProfileOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileOK  %+v", 200, o.Payload)
}

func (o *GetUserProfileOK) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileOK  %+v", 200, o.Payload)
}

func (o *GetUserProfileOK) GetPayload() *models.UserProfile {
	return o.Payload
}

func (o *GetUserProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileBadRequest creates a GetUserProfileBadRequest with default headers values
func NewGetUserProfileBadRequest() *GetUserProfileBadRequest {
	return &GetUserProfileBadRequest{}
}

/*
GetUserProfileBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserProfileBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile bad request response has a 2xx status code
func (o *GetUserProfileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile bad request response has a 3xx status code
func (o *GetUserProfileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile bad request response has a 4xx status code
func (o *GetUserProfileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile bad request response has a 5xx status code
func (o *GetUserProfileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile bad request response a status code equal to that given
func (o *GetUserProfileBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetUserProfileBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserProfileBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserProfileBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileUnauthorized creates a GetUserProfileUnauthorized with default headers values
func NewGetUserProfileUnauthorized() *GetUserProfileUnauthorized {
	return &GetUserProfileUnauthorized{}
}

/*
GetUserProfileUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserProfileUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile unauthorized response has a 2xx status code
func (o *GetUserProfileUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile unauthorized response has a 3xx status code
func (o *GetUserProfileUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile unauthorized response has a 4xx status code
func (o *GetUserProfileUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile unauthorized response has a 5xx status code
func (o *GetUserProfileUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile unauthorized response a status code equal to that given
func (o *GetUserProfileUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetUserProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserProfileUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserProfileUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileForbidden creates a GetUserProfileForbidden with default headers values
func NewGetUserProfileForbidden() *GetUserProfileForbidden {
	return &GetUserProfileForbidden{}
}

/*
GetUserProfileForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetUserProfileForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile forbidden response has a 2xx status code
func (o *GetUserProfileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile forbidden response has a 3xx status code
func (o *GetUserProfileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile forbidden response has a 4xx status code
func (o *GetUserProfileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile forbidden response has a 5xx status code
func (o *GetUserProfileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile forbidden response a status code equal to that given
func (o *GetUserProfileForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetUserProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileForbidden  %+v", 403, o.Payload)
}

func (o *GetUserProfileForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileForbidden  %+v", 403, o.Payload)
}

func (o *GetUserProfileForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileNotFound creates a GetUserProfileNotFound with default headers values
func NewGetUserProfileNotFound() *GetUserProfileNotFound {
	return &GetUserProfileNotFound{}
}

/*
GetUserProfileNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetUserProfileNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile not found response has a 2xx status code
func (o *GetUserProfileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile not found response has a 3xx status code
func (o *GetUserProfileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile not found response has a 4xx status code
func (o *GetUserProfileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile not found response has a 5xx status code
func (o *GetUserProfileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile not found response a status code equal to that given
func (o *GetUserProfileNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetUserProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetUserProfileNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileNotFound  %+v", 404, o.Payload)
}

func (o *GetUserProfileNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileRequestTimeout creates a GetUserProfileRequestTimeout with default headers values
func NewGetUserProfileRequestTimeout() *GetUserProfileRequestTimeout {
	return &GetUserProfileRequestTimeout{}
}

/*
GetUserProfileRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetUserProfileRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile request timeout response has a 2xx status code
func (o *GetUserProfileRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile request timeout response has a 3xx status code
func (o *GetUserProfileRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile request timeout response has a 4xx status code
func (o *GetUserProfileRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile request timeout response has a 5xx status code
func (o *GetUserProfileRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile request timeout response a status code equal to that given
func (o *GetUserProfileRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetUserProfileRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserProfileRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetUserProfileRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileRequestEntityTooLarge creates a GetUserProfileRequestEntityTooLarge with default headers values
func NewGetUserProfileRequestEntityTooLarge() *GetUserProfileRequestEntityTooLarge {
	return &GetUserProfileRequestEntityTooLarge{}
}

/*
GetUserProfileRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetUserProfileRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile request entity too large response has a 2xx status code
func (o *GetUserProfileRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile request entity too large response has a 3xx status code
func (o *GetUserProfileRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile request entity too large response has a 4xx status code
func (o *GetUserProfileRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile request entity too large response has a 5xx status code
func (o *GetUserProfileRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile request entity too large response a status code equal to that given
func (o *GetUserProfileRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetUserProfileRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserProfileRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserProfileRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileUnsupportedMediaType creates a GetUserProfileUnsupportedMediaType with default headers values
func NewGetUserProfileUnsupportedMediaType() *GetUserProfileUnsupportedMediaType {
	return &GetUserProfileUnsupportedMediaType{}
}

/*
GetUserProfileUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserProfileUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile unsupported media type response has a 2xx status code
func (o *GetUserProfileUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile unsupported media type response has a 3xx status code
func (o *GetUserProfileUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile unsupported media type response has a 4xx status code
func (o *GetUserProfileUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile unsupported media type response has a 5xx status code
func (o *GetUserProfileUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile unsupported media type response a status code equal to that given
func (o *GetUserProfileUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetUserProfileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserProfileUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserProfileUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileTooManyRequests creates a GetUserProfileTooManyRequests with default headers values
func NewGetUserProfileTooManyRequests() *GetUserProfileTooManyRequests {
	return &GetUserProfileTooManyRequests{}
}

/*
GetUserProfileTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetUserProfileTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile too many requests response has a 2xx status code
func (o *GetUserProfileTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile too many requests response has a 3xx status code
func (o *GetUserProfileTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile too many requests response has a 4xx status code
func (o *GetUserProfileTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get user profile too many requests response has a 5xx status code
func (o *GetUserProfileTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get user profile too many requests response a status code equal to that given
func (o *GetUserProfileTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetUserProfileTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserProfileTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserProfileTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileInternalServerError creates a GetUserProfileInternalServerError with default headers values
func NewGetUserProfileInternalServerError() *GetUserProfileInternalServerError {
	return &GetUserProfileInternalServerError{}
}

/*
GetUserProfileInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserProfileInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile internal server error response has a 2xx status code
func (o *GetUserProfileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile internal server error response has a 3xx status code
func (o *GetUserProfileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile internal server error response has a 4xx status code
func (o *GetUserProfileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user profile internal server error response has a 5xx status code
func (o *GetUserProfileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get user profile internal server error response a status code equal to that given
func (o *GetUserProfileInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetUserProfileInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserProfileInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserProfileInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileServiceUnavailable creates a GetUserProfileServiceUnavailable with default headers values
func NewGetUserProfileServiceUnavailable() *GetUserProfileServiceUnavailable {
	return &GetUserProfileServiceUnavailable{}
}

/*
GetUserProfileServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserProfileServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile service unavailable response has a 2xx status code
func (o *GetUserProfileServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile service unavailable response has a 3xx status code
func (o *GetUserProfileServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile service unavailable response has a 4xx status code
func (o *GetUserProfileServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user profile service unavailable response has a 5xx status code
func (o *GetUserProfileServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get user profile service unavailable response a status code equal to that given
func (o *GetUserProfileServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetUserProfileServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserProfileServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserProfileServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserProfileGatewayTimeout creates a GetUserProfileGatewayTimeout with default headers values
func NewGetUserProfileGatewayTimeout() *GetUserProfileGatewayTimeout {
	return &GetUserProfileGatewayTimeout{}
}

/*
GetUserProfileGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetUserProfileGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get user profile gateway timeout response has a 2xx status code
func (o *GetUserProfileGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get user profile gateway timeout response has a 3xx status code
func (o *GetUserProfileGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get user profile gateway timeout response has a 4xx status code
func (o *GetUserProfileGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get user profile gateway timeout response has a 5xx status code
func (o *GetUserProfileGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get user profile gateway timeout response a status code equal to that given
func (o *GetUserProfileGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetUserProfileGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserProfileGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/profile][%d] getUserProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserProfileGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
