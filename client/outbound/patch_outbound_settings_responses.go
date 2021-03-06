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

/*PatchOutboundSettingsNoContent handles this case with default header values.

Accepted - Processing Update
*/
type PatchOutboundSettingsNoContent struct {
}

func (o *PatchOutboundSettingsNoContent) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/outbound/settings][%d] patchOutboundSettingsNoContent ", 204)
}

func (o *PatchOutboundSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOutboundSettingsBadRequest creates a PatchOutboundSettingsBadRequest with default headers values
func NewPatchOutboundSettingsBadRequest() *PatchOutboundSettingsBadRequest {
	return &PatchOutboundSettingsBadRequest{}
}

/*PatchOutboundSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchOutboundSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsBadRequest) Error() string {
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

/*PatchOutboundSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchOutboundSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsUnauthorized) Error() string {
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

/*PatchOutboundSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchOutboundSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsForbidden) Error() string {
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

/*PatchOutboundSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchOutboundSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsNotFound) Error() string {
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

// NewPatchOutboundSettingsConflict creates a PatchOutboundSettingsConflict with default headers values
func NewPatchOutboundSettingsConflict() *PatchOutboundSettingsConflict {
	return &PatchOutboundSettingsConflict{}
}

/*PatchOutboundSettingsConflict handles this case with default header values.

Conflict
*/
type PatchOutboundSettingsConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsConflict) Error() string {
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

/*PatchOutboundSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchOutboundSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsRequestEntityTooLarge) Error() string {
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

/*PatchOutboundSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchOutboundSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsUnsupportedMediaType) Error() string {
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

/*PatchOutboundSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchOutboundSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsTooManyRequests) Error() string {
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

/*PatchOutboundSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchOutboundSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsInternalServerError) Error() string {
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

/*PatchOutboundSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchOutboundSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsServiceUnavailable) Error() string {
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

/*PatchOutboundSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchOutboundSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchOutboundSettingsGatewayTimeout) Error() string {
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
