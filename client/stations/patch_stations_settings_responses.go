// Code generated by go-swagger; DO NOT EDIT.

package stations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchStationsSettingsReader is a Reader for the PatchStationsSettings structure.
type PatchStationsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchStationsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchStationsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchStationsSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchStationsSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchStationsSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchStationsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchStationsSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchStationsSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchStationsSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchStationsSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchStationsSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchStationsSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchStationsSettingsOK creates a PatchStationsSettingsOK with default headers values
func NewPatchStationsSettingsOK() *PatchStationsSettingsOK {
	return &PatchStationsSettingsOK{}
}

/*PatchStationsSettingsOK handles this case with default header values.

successful operation
*/
type PatchStationsSettingsOK struct {
	Payload *models.StationSettings
}

func (o *PatchStationsSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchStationsSettingsOK) GetPayload() *models.StationSettings {
	return o.Payload
}

func (o *PatchStationsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StationSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsBadRequest creates a PatchStationsSettingsBadRequest with default headers values
func NewPatchStationsSettingsBadRequest() *PatchStationsSettingsBadRequest {
	return &PatchStationsSettingsBadRequest{}
}

/*PatchStationsSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchStationsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchStationsSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsUnauthorized creates a PatchStationsSettingsUnauthorized with default headers values
func NewPatchStationsSettingsUnauthorized() *PatchStationsSettingsUnauthorized {
	return &PatchStationsSettingsUnauthorized{}
}

/*PatchStationsSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchStationsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchStationsSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsForbidden creates a PatchStationsSettingsForbidden with default headers values
func NewPatchStationsSettingsForbidden() *PatchStationsSettingsForbidden {
	return &PatchStationsSettingsForbidden{}
}

/*PatchStationsSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchStationsSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchStationsSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsNotFound creates a PatchStationsSettingsNotFound with default headers values
func NewPatchStationsSettingsNotFound() *PatchStationsSettingsNotFound {
	return &PatchStationsSettingsNotFound{}
}

/*PatchStationsSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchStationsSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchStationsSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsRequestEntityTooLarge creates a PatchStationsSettingsRequestEntityTooLarge with default headers values
func NewPatchStationsSettingsRequestEntityTooLarge() *PatchStationsSettingsRequestEntityTooLarge {
	return &PatchStationsSettingsRequestEntityTooLarge{}
}

/*PatchStationsSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchStationsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchStationsSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsUnsupportedMediaType creates a PatchStationsSettingsUnsupportedMediaType with default headers values
func NewPatchStationsSettingsUnsupportedMediaType() *PatchStationsSettingsUnsupportedMediaType {
	return &PatchStationsSettingsUnsupportedMediaType{}
}

/*PatchStationsSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchStationsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchStationsSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsTooManyRequests creates a PatchStationsSettingsTooManyRequests with default headers values
func NewPatchStationsSettingsTooManyRequests() *PatchStationsSettingsTooManyRequests {
	return &PatchStationsSettingsTooManyRequests{}
}

/*PatchStationsSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchStationsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchStationsSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsInternalServerError creates a PatchStationsSettingsInternalServerError with default headers values
func NewPatchStationsSettingsInternalServerError() *PatchStationsSettingsInternalServerError {
	return &PatchStationsSettingsInternalServerError{}
}

/*PatchStationsSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchStationsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchStationsSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsServiceUnavailable creates a PatchStationsSettingsServiceUnavailable with default headers values
func NewPatchStationsSettingsServiceUnavailable() *PatchStationsSettingsServiceUnavailable {
	return &PatchStationsSettingsServiceUnavailable{}
}

/*PatchStationsSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchStationsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchStationsSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchStationsSettingsGatewayTimeout creates a PatchStationsSettingsGatewayTimeout with default headers values
func NewPatchStationsSettingsGatewayTimeout() *PatchStationsSettingsGatewayTimeout {
	return &PatchStationsSettingsGatewayTimeout{}
}

/*PatchStationsSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchStationsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchStationsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/stations/settings][%d] patchStationsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchStationsSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchStationsSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
