// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutChatsSettingsReader is a Reader for the PutChatsSettings structure.
type PutChatsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutChatsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutChatsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutChatsSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutChatsSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutChatsSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutChatsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutChatsSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutChatsSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutChatsSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutChatsSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutChatsSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutChatsSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutChatsSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutChatsSettingsOK creates a PutChatsSettingsOK with default headers values
func NewPutChatsSettingsOK() *PutChatsSettingsOK {
	return &PutChatsSettingsOK{}
}

/*
PutChatsSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PutChatsSettingsOK struct {
	Payload *models.ChatSettings
}

// IsSuccess returns true when this put chats settings o k response has a 2xx status code
func (o *PutChatsSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put chats settings o k response has a 3xx status code
func (o *PutChatsSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings o k response has a 4xx status code
func (o *PutChatsSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put chats settings o k response has a 5xx status code
func (o *PutChatsSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings o k response a status code equal to that given
func (o *PutChatsSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutChatsSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsOK  %+v", 200, o.Payload)
}

func (o *PutChatsSettingsOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsOK  %+v", 200, o.Payload)
}

func (o *PutChatsSettingsOK) GetPayload() *models.ChatSettings {
	return o.Payload
}

func (o *PutChatsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsBadRequest creates a PutChatsSettingsBadRequest with default headers values
func NewPutChatsSettingsBadRequest() *PutChatsSettingsBadRequest {
	return &PutChatsSettingsBadRequest{}
}

/*
PutChatsSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutChatsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings bad request response has a 2xx status code
func (o *PutChatsSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings bad request response has a 3xx status code
func (o *PutChatsSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings bad request response has a 4xx status code
func (o *PutChatsSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings bad request response has a 5xx status code
func (o *PutChatsSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings bad request response a status code equal to that given
func (o *PutChatsSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutChatsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutChatsSettingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutChatsSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsUnauthorized creates a PutChatsSettingsUnauthorized with default headers values
func NewPutChatsSettingsUnauthorized() *PutChatsSettingsUnauthorized {
	return &PutChatsSettingsUnauthorized{}
}

/*
PutChatsSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutChatsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings unauthorized response has a 2xx status code
func (o *PutChatsSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings unauthorized response has a 3xx status code
func (o *PutChatsSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings unauthorized response has a 4xx status code
func (o *PutChatsSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings unauthorized response has a 5xx status code
func (o *PutChatsSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings unauthorized response a status code equal to that given
func (o *PutChatsSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutChatsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutChatsSettingsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutChatsSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsForbidden creates a PutChatsSettingsForbidden with default headers values
func NewPutChatsSettingsForbidden() *PutChatsSettingsForbidden {
	return &PutChatsSettingsForbidden{}
}

/*
PutChatsSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutChatsSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings forbidden response has a 2xx status code
func (o *PutChatsSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings forbidden response has a 3xx status code
func (o *PutChatsSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings forbidden response has a 4xx status code
func (o *PutChatsSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings forbidden response has a 5xx status code
func (o *PutChatsSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings forbidden response a status code equal to that given
func (o *PutChatsSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutChatsSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutChatsSettingsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutChatsSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsNotFound creates a PutChatsSettingsNotFound with default headers values
func NewPutChatsSettingsNotFound() *PutChatsSettingsNotFound {
	return &PutChatsSettingsNotFound{}
}

/*
PutChatsSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutChatsSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings not found response has a 2xx status code
func (o *PutChatsSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings not found response has a 3xx status code
func (o *PutChatsSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings not found response has a 4xx status code
func (o *PutChatsSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings not found response has a 5xx status code
func (o *PutChatsSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings not found response a status code equal to that given
func (o *PutChatsSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutChatsSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutChatsSettingsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutChatsSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsRequestTimeout creates a PutChatsSettingsRequestTimeout with default headers values
func NewPutChatsSettingsRequestTimeout() *PutChatsSettingsRequestTimeout {
	return &PutChatsSettingsRequestTimeout{}
}

/*
PutChatsSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutChatsSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings request timeout response has a 2xx status code
func (o *PutChatsSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings request timeout response has a 3xx status code
func (o *PutChatsSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings request timeout response has a 4xx status code
func (o *PutChatsSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings request timeout response has a 5xx status code
func (o *PutChatsSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings request timeout response a status code equal to that given
func (o *PutChatsSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutChatsSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutChatsSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutChatsSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsRequestEntityTooLarge creates a PutChatsSettingsRequestEntityTooLarge with default headers values
func NewPutChatsSettingsRequestEntityTooLarge() *PutChatsSettingsRequestEntityTooLarge {
	return &PutChatsSettingsRequestEntityTooLarge{}
}

/*
PutChatsSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutChatsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings request entity too large response has a 2xx status code
func (o *PutChatsSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings request entity too large response has a 3xx status code
func (o *PutChatsSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings request entity too large response has a 4xx status code
func (o *PutChatsSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings request entity too large response has a 5xx status code
func (o *PutChatsSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings request entity too large response a status code equal to that given
func (o *PutChatsSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutChatsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutChatsSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutChatsSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsUnsupportedMediaType creates a PutChatsSettingsUnsupportedMediaType with default headers values
func NewPutChatsSettingsUnsupportedMediaType() *PutChatsSettingsUnsupportedMediaType {
	return &PutChatsSettingsUnsupportedMediaType{}
}

/*
PutChatsSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutChatsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings unsupported media type response has a 2xx status code
func (o *PutChatsSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings unsupported media type response has a 3xx status code
func (o *PutChatsSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings unsupported media type response has a 4xx status code
func (o *PutChatsSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings unsupported media type response has a 5xx status code
func (o *PutChatsSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings unsupported media type response a status code equal to that given
func (o *PutChatsSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutChatsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutChatsSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutChatsSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsTooManyRequests creates a PutChatsSettingsTooManyRequests with default headers values
func NewPutChatsSettingsTooManyRequests() *PutChatsSettingsTooManyRequests {
	return &PutChatsSettingsTooManyRequests{}
}

/*
PutChatsSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutChatsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings too many requests response has a 2xx status code
func (o *PutChatsSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings too many requests response has a 3xx status code
func (o *PutChatsSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings too many requests response has a 4xx status code
func (o *PutChatsSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put chats settings too many requests response has a 5xx status code
func (o *PutChatsSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put chats settings too many requests response a status code equal to that given
func (o *PutChatsSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutChatsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutChatsSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutChatsSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsInternalServerError creates a PutChatsSettingsInternalServerError with default headers values
func NewPutChatsSettingsInternalServerError() *PutChatsSettingsInternalServerError {
	return &PutChatsSettingsInternalServerError{}
}

/*
PutChatsSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutChatsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings internal server error response has a 2xx status code
func (o *PutChatsSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings internal server error response has a 3xx status code
func (o *PutChatsSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings internal server error response has a 4xx status code
func (o *PutChatsSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put chats settings internal server error response has a 5xx status code
func (o *PutChatsSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put chats settings internal server error response a status code equal to that given
func (o *PutChatsSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutChatsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutChatsSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutChatsSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsServiceUnavailable creates a PutChatsSettingsServiceUnavailable with default headers values
func NewPutChatsSettingsServiceUnavailable() *PutChatsSettingsServiceUnavailable {
	return &PutChatsSettingsServiceUnavailable{}
}

/*
PutChatsSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutChatsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings service unavailable response has a 2xx status code
func (o *PutChatsSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings service unavailable response has a 3xx status code
func (o *PutChatsSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings service unavailable response has a 4xx status code
func (o *PutChatsSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put chats settings service unavailable response has a 5xx status code
func (o *PutChatsSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put chats settings service unavailable response a status code equal to that given
func (o *PutChatsSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutChatsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutChatsSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutChatsSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatsSettingsGatewayTimeout creates a PutChatsSettingsGatewayTimeout with default headers values
func NewPutChatsSettingsGatewayTimeout() *PutChatsSettingsGatewayTimeout {
	return &PutChatsSettingsGatewayTimeout{}
}

/*
PutChatsSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutChatsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put chats settings gateway timeout response has a 2xx status code
func (o *PutChatsSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put chats settings gateway timeout response has a 3xx status code
func (o *PutChatsSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put chats settings gateway timeout response has a 4xx status code
func (o *PutChatsSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put chats settings gateway timeout response has a 5xx status code
func (o *PutChatsSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put chats settings gateway timeout response a status code equal to that given
func (o *PutChatsSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutChatsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutChatsSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/chats/settings][%d] putChatsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutChatsSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatsSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}