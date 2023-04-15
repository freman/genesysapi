// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchOutboundSettingsReader is a Reader for the PatchOutboundSettings structure.
type PatchOutboundSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOutboundSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPatchOutboundSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchOutboundSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchOutboundSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchOutboundSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOutboundSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchOutboundSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchOutboundSettingsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchOutboundSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchOutboundSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchOutboundSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchOutboundSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchOutboundSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchOutboundSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchOutboundSettingsNoContent creates a PatchOutboundSettingsNoContent with default headers values
func NewPatchOutboundSettingsNoContent() *PatchOutboundSettingsNoContent {
	return &PatchOutboundSettingsNoContent{}
}

/*
PatchOutboundSettingsNoContent describes a response with status code 204, with default header values.

Accepted - Processing Update
*/
type PatchOutboundSettingsNoContent struct {
}

// IsSuccess returns true when this patch outbound settings no content response has a 2xx status code
func (o *PatchOutboundSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch outbound settings no content response has a 3xx status code
func (o *PatchOutboundSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings no content response has a 4xx status code
func (o *PatchOutboundSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound settings no content response has a 5xx status code
func (o *PatchOutboundSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings no content response a status code equal to that given
func (o *PatchOutboundSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *PatchOutboundSettingsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsNoContent ", 204)
}

func (o *PatchOutboundSettingsNoContent) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsNoContent ", 204)
}

func (o *PatchOutboundSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOutboundSettingsBadRequest creates a PatchOutboundSettingsBadRequest with default headers values
func NewPatchOutboundSettingsBadRequest() *PatchOutboundSettingsBadRequest {
	return &PatchOutboundSettingsBadRequest{}
}

/*
PatchOutboundSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchOutboundSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings bad request response has a 2xx status code
func (o *PatchOutboundSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings bad request response has a 3xx status code
func (o *PatchOutboundSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings bad request response has a 4xx status code
func (o *PatchOutboundSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings bad request response has a 5xx status code
func (o *PatchOutboundSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings bad request response a status code equal to that given
func (o *PatchOutboundSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PatchOutboundSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchOutboundSettingsBadRequest) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchOutboundSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsUnauthorized creates a PatchOutboundSettingsUnauthorized with default headers values
func NewPatchOutboundSettingsUnauthorized() *PatchOutboundSettingsUnauthorized {
	return &PatchOutboundSettingsUnauthorized{}
}

/*
PatchOutboundSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchOutboundSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings unauthorized response has a 2xx status code
func (o *PatchOutboundSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings unauthorized response has a 3xx status code
func (o *PatchOutboundSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings unauthorized response has a 4xx status code
func (o *PatchOutboundSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings unauthorized response has a 5xx status code
func (o *PatchOutboundSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings unauthorized response a status code equal to that given
func (o *PatchOutboundSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PatchOutboundSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchOutboundSettingsUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchOutboundSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsForbidden creates a PatchOutboundSettingsForbidden with default headers values
func NewPatchOutboundSettingsForbidden() *PatchOutboundSettingsForbidden {
	return &PatchOutboundSettingsForbidden{}
}

/*
PatchOutboundSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PatchOutboundSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings forbidden response has a 2xx status code
func (o *PatchOutboundSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings forbidden response has a 3xx status code
func (o *PatchOutboundSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings forbidden response has a 4xx status code
func (o *PatchOutboundSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings forbidden response has a 5xx status code
func (o *PatchOutboundSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings forbidden response a status code equal to that given
func (o *PatchOutboundSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PatchOutboundSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchOutboundSettingsForbidden) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchOutboundSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsNotFound creates a PatchOutboundSettingsNotFound with default headers values
func NewPatchOutboundSettingsNotFound() *PatchOutboundSettingsNotFound {
	return &PatchOutboundSettingsNotFound{}
}

/*
PatchOutboundSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PatchOutboundSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings not found response has a 2xx status code
func (o *PatchOutboundSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings not found response has a 3xx status code
func (o *PatchOutboundSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings not found response has a 4xx status code
func (o *PatchOutboundSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings not found response has a 5xx status code
func (o *PatchOutboundSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings not found response a status code equal to that given
func (o *PatchOutboundSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PatchOutboundSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchOutboundSettingsNotFound) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchOutboundSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsRequestTimeout creates a PatchOutboundSettingsRequestTimeout with default headers values
func NewPatchOutboundSettingsRequestTimeout() *PatchOutboundSettingsRequestTimeout {
	return &PatchOutboundSettingsRequestTimeout{}
}

/*
PatchOutboundSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchOutboundSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings request timeout response has a 2xx status code
func (o *PatchOutboundSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings request timeout response has a 3xx status code
func (o *PatchOutboundSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings request timeout response has a 4xx status code
func (o *PatchOutboundSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings request timeout response has a 5xx status code
func (o *PatchOutboundSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings request timeout response a status code equal to that given
func (o *PatchOutboundSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PatchOutboundSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchOutboundSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchOutboundSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsConflict creates a PatchOutboundSettingsConflict with default headers values
func NewPatchOutboundSettingsConflict() *PatchOutboundSettingsConflict {
	return &PatchOutboundSettingsConflict{}
}

/*
PatchOutboundSettingsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PatchOutboundSettingsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings conflict response has a 2xx status code
func (o *PatchOutboundSettingsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings conflict response has a 3xx status code
func (o *PatchOutboundSettingsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings conflict response has a 4xx status code
func (o *PatchOutboundSettingsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings conflict response has a 5xx status code
func (o *PatchOutboundSettingsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings conflict response a status code equal to that given
func (o *PatchOutboundSettingsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PatchOutboundSettingsConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsConflict  %+v", 409, o.Payload)
}

func (o *PatchOutboundSettingsConflict) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsConflict  %+v", 409, o.Payload)
}

func (o *PatchOutboundSettingsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsRequestEntityTooLarge creates a PatchOutboundSettingsRequestEntityTooLarge with default headers values
func NewPatchOutboundSettingsRequestEntityTooLarge() *PatchOutboundSettingsRequestEntityTooLarge {
	return &PatchOutboundSettingsRequestEntityTooLarge{}
}

/*
PatchOutboundSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PatchOutboundSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings request entity too large response has a 2xx status code
func (o *PatchOutboundSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings request entity too large response has a 3xx status code
func (o *PatchOutboundSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings request entity too large response has a 4xx status code
func (o *PatchOutboundSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings request entity too large response has a 5xx status code
func (o *PatchOutboundSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings request entity too large response a status code equal to that given
func (o *PatchOutboundSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PatchOutboundSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchOutboundSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchOutboundSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsUnsupportedMediaType creates a PatchOutboundSettingsUnsupportedMediaType with default headers values
func NewPatchOutboundSettingsUnsupportedMediaType() *PatchOutboundSettingsUnsupportedMediaType {
	return &PatchOutboundSettingsUnsupportedMediaType{}
}

/*
PatchOutboundSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchOutboundSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings unsupported media type response has a 2xx status code
func (o *PatchOutboundSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings unsupported media type response has a 3xx status code
func (o *PatchOutboundSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings unsupported media type response has a 4xx status code
func (o *PatchOutboundSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings unsupported media type response has a 5xx status code
func (o *PatchOutboundSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings unsupported media type response a status code equal to that given
func (o *PatchOutboundSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PatchOutboundSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchOutboundSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchOutboundSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsTooManyRequests creates a PatchOutboundSettingsTooManyRequests with default headers values
func NewPatchOutboundSettingsTooManyRequests() *PatchOutboundSettingsTooManyRequests {
	return &PatchOutboundSettingsTooManyRequests{}
}

/*
PatchOutboundSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchOutboundSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings too many requests response has a 2xx status code
func (o *PatchOutboundSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings too many requests response has a 3xx status code
func (o *PatchOutboundSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings too many requests response has a 4xx status code
func (o *PatchOutboundSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch outbound settings too many requests response has a 5xx status code
func (o *PatchOutboundSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this patch outbound settings too many requests response a status code equal to that given
func (o *PatchOutboundSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PatchOutboundSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchOutboundSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchOutboundSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsInternalServerError creates a PatchOutboundSettingsInternalServerError with default headers values
func NewPatchOutboundSettingsInternalServerError() *PatchOutboundSettingsInternalServerError {
	return &PatchOutboundSettingsInternalServerError{}
}

/*
PatchOutboundSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchOutboundSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings internal server error response has a 2xx status code
func (o *PatchOutboundSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings internal server error response has a 3xx status code
func (o *PatchOutboundSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings internal server error response has a 4xx status code
func (o *PatchOutboundSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound settings internal server error response has a 5xx status code
func (o *PatchOutboundSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound settings internal server error response a status code equal to that given
func (o *PatchOutboundSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PatchOutboundSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchOutboundSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchOutboundSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsServiceUnavailable creates a PatchOutboundSettingsServiceUnavailable with default headers values
func NewPatchOutboundSettingsServiceUnavailable() *PatchOutboundSettingsServiceUnavailable {
	return &PatchOutboundSettingsServiceUnavailable{}
}

/*
PatchOutboundSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchOutboundSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings service unavailable response has a 2xx status code
func (o *PatchOutboundSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings service unavailable response has a 3xx status code
func (o *PatchOutboundSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings service unavailable response has a 4xx status code
func (o *PatchOutboundSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound settings service unavailable response has a 5xx status code
func (o *PatchOutboundSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound settings service unavailable response a status code equal to that given
func (o *PatchOutboundSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PatchOutboundSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchOutboundSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchOutboundSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOutboundSettingsGatewayTimeout creates a PatchOutboundSettingsGatewayTimeout with default headers values
func NewPatchOutboundSettingsGatewayTimeout() *PatchOutboundSettingsGatewayTimeout {
	return &PatchOutboundSettingsGatewayTimeout{}
}

/*
PatchOutboundSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PatchOutboundSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this patch outbound settings gateway timeout response has a 2xx status code
func (o *PatchOutboundSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch outbound settings gateway timeout response has a 3xx status code
func (o *PatchOutboundSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch outbound settings gateway timeout response has a 4xx status code
func (o *PatchOutboundSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch outbound settings gateway timeout response has a 5xx status code
func (o *PatchOutboundSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this patch outbound settings gateway timeout response a status code equal to that given
func (o *PatchOutboundSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PatchOutboundSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchOutboundSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchOutboundSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchOutboundSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
