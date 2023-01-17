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

// PatchChatSettingsReader is a Reader for the PatchChatSettings structure.
type PatchChatSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchChatSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchChatSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchChatSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchChatSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchChatSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchChatSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchChatSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchChatSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchChatSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchChatSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchChatSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchChatSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchChatSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchChatSettingsOK creates a PatchChatSettingsOK with default headers values
func NewPatchChatSettingsOK() *PatchChatSettingsOK {
	return &PatchChatSettingsOK{}
}

/*
PatchChatSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PatchChatSettingsOK struct {
	Payload *models.ChatSettings
}

// IsSuccess returns true when this patch chat settings o k response has a 2xx status code
func (o *PatchChatSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch chat settings o k response has a 3xx status code
func (o *PatchChatSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings o k response has a 4xx status code
func (o *PatchChatSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch chat settings o k response has a 5xx status code
func (o *PatchChatSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings o k response a status code equal to that given
func (o *PatchChatSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PatchChatSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchChatSettingsOK) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchChatSettingsOK) GetPayload() *models.ChatSettings {
	return o.Payload
}

func (o *PatchChatSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsBadRequest creates a PatchChatSettingsBadRequest with default headers values
func NewPatchChatSettingsBadRequest() *PatchChatSettingsBadRequest {
	return &PatchChatSettingsBadRequest{}
}

/*
PatchChatSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchChatSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings bad request response has a 2xx status code
func (o *PatchChatSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings bad request response has a 3xx status code
func (o *PatchChatSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings bad request response has a 4xx status code
func (o *PatchChatSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings bad request response has a 5xx status code
func (o *PatchChatSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings bad request response a status code equal to that given
func (o *PatchChatSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchChatSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchChatSettingsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchChatSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsUnauthorized creates a PatchChatSettingsUnauthorized with default headers values
func NewPatchChatSettingsUnauthorized() *PatchChatSettingsUnauthorized {
	return &PatchChatSettingsUnauthorized{}
}

/*
PatchChatSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchChatSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings unauthorized response has a 2xx status code
func (o *PatchChatSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings unauthorized response has a 3xx status code
func (o *PatchChatSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings unauthorized response has a 4xx status code
func (o *PatchChatSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings unauthorized response has a 5xx status code
func (o *PatchChatSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings unauthorized response a status code equal to that given
func (o *PatchChatSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchChatSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchChatSettingsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchChatSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsForbidden creates a PatchChatSettingsForbidden with default headers values
func NewPatchChatSettingsForbidden() *PatchChatSettingsForbidden {
	return &PatchChatSettingsForbidden{}
}

/*
PatchChatSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchChatSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings forbidden response has a 2xx status code
func (o *PatchChatSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings forbidden response has a 3xx status code
func (o *PatchChatSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings forbidden response has a 4xx status code
func (o *PatchChatSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings forbidden response has a 5xx status code
func (o *PatchChatSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings forbidden response a status code equal to that given
func (o *PatchChatSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchChatSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchChatSettingsForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchChatSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsNotFound creates a PatchChatSettingsNotFound with default headers values
func NewPatchChatSettingsNotFound() *PatchChatSettingsNotFound {
	return &PatchChatSettingsNotFound{}
}

/*
PatchChatSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchChatSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings not found response has a 2xx status code
func (o *PatchChatSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings not found response has a 3xx status code
func (o *PatchChatSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings not found response has a 4xx status code
func (o *PatchChatSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings not found response has a 5xx status code
func (o *PatchChatSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings not found response a status code equal to that given
func (o *PatchChatSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchChatSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchChatSettingsNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchChatSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsRequestTimeout creates a PatchChatSettingsRequestTimeout with default headers values
func NewPatchChatSettingsRequestTimeout() *PatchChatSettingsRequestTimeout {
	return &PatchChatSettingsRequestTimeout{}
}

/*
PatchChatSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchChatSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings request timeout response has a 2xx status code
func (o *PatchChatSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings request timeout response has a 3xx status code
func (o *PatchChatSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings request timeout response has a 4xx status code
func (o *PatchChatSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings request timeout response has a 5xx status code
func (o *PatchChatSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings request timeout response a status code equal to that given
func (o *PatchChatSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchChatSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchChatSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchChatSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsRequestEntityTooLarge creates a PatchChatSettingsRequestEntityTooLarge with default headers values
func NewPatchChatSettingsRequestEntityTooLarge() *PatchChatSettingsRequestEntityTooLarge {
	return &PatchChatSettingsRequestEntityTooLarge{}
}

/*
PatchChatSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchChatSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings request entity too large response has a 2xx status code
func (o *PatchChatSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings request entity too large response has a 3xx status code
func (o *PatchChatSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings request entity too large response has a 4xx status code
func (o *PatchChatSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings request entity too large response has a 5xx status code
func (o *PatchChatSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings request entity too large response a status code equal to that given
func (o *PatchChatSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchChatSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchChatSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchChatSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsUnsupportedMediaType creates a PatchChatSettingsUnsupportedMediaType with default headers values
func NewPatchChatSettingsUnsupportedMediaType() *PatchChatSettingsUnsupportedMediaType {
	return &PatchChatSettingsUnsupportedMediaType{}
}

/*
PatchChatSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchChatSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings unsupported media type response has a 2xx status code
func (o *PatchChatSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings unsupported media type response has a 3xx status code
func (o *PatchChatSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings unsupported media type response has a 4xx status code
func (o *PatchChatSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings unsupported media type response has a 5xx status code
func (o *PatchChatSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings unsupported media type response a status code equal to that given
func (o *PatchChatSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchChatSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchChatSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchChatSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsTooManyRequests creates a PatchChatSettingsTooManyRequests with default headers values
func NewPatchChatSettingsTooManyRequests() *PatchChatSettingsTooManyRequests {
	return &PatchChatSettingsTooManyRequests{}
}

/*
PatchChatSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchChatSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings too many requests response has a 2xx status code
func (o *PatchChatSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings too many requests response has a 3xx status code
func (o *PatchChatSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings too many requests response has a 4xx status code
func (o *PatchChatSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch chat settings too many requests response has a 5xx status code
func (o *PatchChatSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch chat settings too many requests response a status code equal to that given
func (o *PatchChatSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchChatSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchChatSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchChatSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsInternalServerError creates a PatchChatSettingsInternalServerError with default headers values
func NewPatchChatSettingsInternalServerError() *PatchChatSettingsInternalServerError {
	return &PatchChatSettingsInternalServerError{}
}

/*
PatchChatSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchChatSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings internal server error response has a 2xx status code
func (o *PatchChatSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings internal server error response has a 3xx status code
func (o *PatchChatSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings internal server error response has a 4xx status code
func (o *PatchChatSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch chat settings internal server error response has a 5xx status code
func (o *PatchChatSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch chat settings internal server error response a status code equal to that given
func (o *PatchChatSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchChatSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchChatSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchChatSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsServiceUnavailable creates a PatchChatSettingsServiceUnavailable with default headers values
func NewPatchChatSettingsServiceUnavailable() *PatchChatSettingsServiceUnavailable {
	return &PatchChatSettingsServiceUnavailable{}
}

/*
PatchChatSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchChatSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings service unavailable response has a 2xx status code
func (o *PatchChatSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings service unavailable response has a 3xx status code
func (o *PatchChatSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings service unavailable response has a 4xx status code
func (o *PatchChatSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch chat settings service unavailable response has a 5xx status code
func (o *PatchChatSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch chat settings service unavailable response a status code equal to that given
func (o *PatchChatSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchChatSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchChatSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchChatSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchChatSettingsGatewayTimeout creates a PatchChatSettingsGatewayTimeout with default headers values
func NewPatchChatSettingsGatewayTimeout() *PatchChatSettingsGatewayTimeout {
	return &PatchChatSettingsGatewayTimeout{}
}

/*
PatchChatSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchChatSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch chat settings gateway timeout response has a 2xx status code
func (o *PatchChatSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch chat settings gateway timeout response has a 3xx status code
func (o *PatchChatSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch chat settings gateway timeout response has a 4xx status code
func (o *PatchChatSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch chat settings gateway timeout response has a 5xx status code
func (o *PatchChatSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch chat settings gateway timeout response a status code equal to that given
func (o *PatchChatSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchChatSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchChatSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/chat/settings][%d] patchChatSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchChatSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchChatSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
