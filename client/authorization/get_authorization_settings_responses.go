// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuthorizationSettingsReader is a Reader for the GetAuthorizationSettings structure.
type GetAuthorizationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationSettingsOK creates a GetAuthorizationSettingsOK with default headers values
func NewGetAuthorizationSettingsOK() *GetAuthorizationSettingsOK {
	return &GetAuthorizationSettingsOK{}
}

/*
GetAuthorizationSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAuthorizationSettingsOK struct {
	Payload *models.AuthorizationSettings
}

// IsSuccess returns true when this get authorization settings o k response has a 2xx status code
func (o *GetAuthorizationSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get authorization settings o k response has a 3xx status code
func (o *GetAuthorizationSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings o k response has a 4xx status code
func (o *GetAuthorizationSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization settings o k response has a 5xx status code
func (o *GetAuthorizationSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings o k response a status code equal to that given
func (o *GetAuthorizationSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAuthorizationSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationSettingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationSettingsOK) GetPayload() *models.AuthorizationSettings {
	return o.Payload
}

func (o *GetAuthorizationSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsBadRequest creates a GetAuthorizationSettingsBadRequest with default headers values
func NewGetAuthorizationSettingsBadRequest() *GetAuthorizationSettingsBadRequest {
	return &GetAuthorizationSettingsBadRequest{}
}

/*
GetAuthorizationSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings bad request response has a 2xx status code
func (o *GetAuthorizationSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings bad request response has a 3xx status code
func (o *GetAuthorizationSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings bad request response has a 4xx status code
func (o *GetAuthorizationSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings bad request response has a 5xx status code
func (o *GetAuthorizationSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings bad request response a status code equal to that given
func (o *GetAuthorizationSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAuthorizationSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsUnauthorized creates a GetAuthorizationSettingsUnauthorized with default headers values
func NewGetAuthorizationSettingsUnauthorized() *GetAuthorizationSettingsUnauthorized {
	return &GetAuthorizationSettingsUnauthorized{}
}

/*
GetAuthorizationSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings unauthorized response has a 2xx status code
func (o *GetAuthorizationSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings unauthorized response has a 3xx status code
func (o *GetAuthorizationSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings unauthorized response has a 4xx status code
func (o *GetAuthorizationSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings unauthorized response has a 5xx status code
func (o *GetAuthorizationSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings unauthorized response a status code equal to that given
func (o *GetAuthorizationSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAuthorizationSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsForbidden creates a GetAuthorizationSettingsForbidden with default headers values
func NewGetAuthorizationSettingsForbidden() *GetAuthorizationSettingsForbidden {
	return &GetAuthorizationSettingsForbidden{}
}

/*
GetAuthorizationSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings forbidden response has a 2xx status code
func (o *GetAuthorizationSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings forbidden response has a 3xx status code
func (o *GetAuthorizationSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings forbidden response has a 4xx status code
func (o *GetAuthorizationSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings forbidden response has a 5xx status code
func (o *GetAuthorizationSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings forbidden response a status code equal to that given
func (o *GetAuthorizationSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAuthorizationSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsNotFound creates a GetAuthorizationSettingsNotFound with default headers values
func NewGetAuthorizationSettingsNotFound() *GetAuthorizationSettingsNotFound {
	return &GetAuthorizationSettingsNotFound{}
}

/*
GetAuthorizationSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAuthorizationSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings not found response has a 2xx status code
func (o *GetAuthorizationSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings not found response has a 3xx status code
func (o *GetAuthorizationSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings not found response has a 4xx status code
func (o *GetAuthorizationSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings not found response has a 5xx status code
func (o *GetAuthorizationSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings not found response a status code equal to that given
func (o *GetAuthorizationSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAuthorizationSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsRequestTimeout creates a GetAuthorizationSettingsRequestTimeout with default headers values
func NewGetAuthorizationSettingsRequestTimeout() *GetAuthorizationSettingsRequestTimeout {
	return &GetAuthorizationSettingsRequestTimeout{}
}

/*
GetAuthorizationSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings request timeout response has a 2xx status code
func (o *GetAuthorizationSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings request timeout response has a 3xx status code
func (o *GetAuthorizationSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings request timeout response has a 4xx status code
func (o *GetAuthorizationSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings request timeout response has a 5xx status code
func (o *GetAuthorizationSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings request timeout response a status code equal to that given
func (o *GetAuthorizationSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAuthorizationSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsRequestEntityTooLarge creates a GetAuthorizationSettingsRequestEntityTooLarge with default headers values
func NewGetAuthorizationSettingsRequestEntityTooLarge() *GetAuthorizationSettingsRequestEntityTooLarge {
	return &GetAuthorizationSettingsRequestEntityTooLarge{}
}

/*
GetAuthorizationSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAuthorizationSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings request entity too large response has a 2xx status code
func (o *GetAuthorizationSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings request entity too large response has a 3xx status code
func (o *GetAuthorizationSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings request entity too large response has a 4xx status code
func (o *GetAuthorizationSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings request entity too large response has a 5xx status code
func (o *GetAuthorizationSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings request entity too large response a status code equal to that given
func (o *GetAuthorizationSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAuthorizationSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsUnsupportedMediaType creates a GetAuthorizationSettingsUnsupportedMediaType with default headers values
func NewGetAuthorizationSettingsUnsupportedMediaType() *GetAuthorizationSettingsUnsupportedMediaType {
	return &GetAuthorizationSettingsUnsupportedMediaType{}
}

/*
GetAuthorizationSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings unsupported media type response has a 2xx status code
func (o *GetAuthorizationSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings unsupported media type response has a 3xx status code
func (o *GetAuthorizationSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings unsupported media type response has a 4xx status code
func (o *GetAuthorizationSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings unsupported media type response has a 5xx status code
func (o *GetAuthorizationSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings unsupported media type response a status code equal to that given
func (o *GetAuthorizationSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAuthorizationSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsTooManyRequests creates a GetAuthorizationSettingsTooManyRequests with default headers values
func NewGetAuthorizationSettingsTooManyRequests() *GetAuthorizationSettingsTooManyRequests {
	return &GetAuthorizationSettingsTooManyRequests{}
}

/*
GetAuthorizationSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings too many requests response has a 2xx status code
func (o *GetAuthorizationSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings too many requests response has a 3xx status code
func (o *GetAuthorizationSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings too many requests response has a 4xx status code
func (o *GetAuthorizationSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get authorization settings too many requests response has a 5xx status code
func (o *GetAuthorizationSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get authorization settings too many requests response a status code equal to that given
func (o *GetAuthorizationSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAuthorizationSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsInternalServerError creates a GetAuthorizationSettingsInternalServerError with default headers values
func NewGetAuthorizationSettingsInternalServerError() *GetAuthorizationSettingsInternalServerError {
	return &GetAuthorizationSettingsInternalServerError{}
}

/*
GetAuthorizationSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings internal server error response has a 2xx status code
func (o *GetAuthorizationSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings internal server error response has a 3xx status code
func (o *GetAuthorizationSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings internal server error response has a 4xx status code
func (o *GetAuthorizationSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization settings internal server error response has a 5xx status code
func (o *GetAuthorizationSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization settings internal server error response a status code equal to that given
func (o *GetAuthorizationSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAuthorizationSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsServiceUnavailable creates a GetAuthorizationSettingsServiceUnavailable with default headers values
func NewGetAuthorizationSettingsServiceUnavailable() *GetAuthorizationSettingsServiceUnavailable {
	return &GetAuthorizationSettingsServiceUnavailable{}
}

/*
GetAuthorizationSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings service unavailable response has a 2xx status code
func (o *GetAuthorizationSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings service unavailable response has a 3xx status code
func (o *GetAuthorizationSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings service unavailable response has a 4xx status code
func (o *GetAuthorizationSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization settings service unavailable response has a 5xx status code
func (o *GetAuthorizationSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization settings service unavailable response a status code equal to that given
func (o *GetAuthorizationSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAuthorizationSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationSettingsGatewayTimeout creates a GetAuthorizationSettingsGatewayTimeout with default headers values
func NewGetAuthorizationSettingsGatewayTimeout() *GetAuthorizationSettingsGatewayTimeout {
	return &GetAuthorizationSettingsGatewayTimeout{}
}

/*
GetAuthorizationSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAuthorizationSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get authorization settings gateway timeout response has a 2xx status code
func (o *GetAuthorizationSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get authorization settings gateway timeout response has a 3xx status code
func (o *GetAuthorizationSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get authorization settings gateway timeout response has a 4xx status code
func (o *GetAuthorizationSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get authorization settings gateway timeout response has a 5xx status code
func (o *GetAuthorizationSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get authorization settings gateway timeout response a status code equal to that given
func (o *GetAuthorizationSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAuthorizationSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/authorization/settings][%d] getAuthorizationSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
