// Code generated by go-swagger; DO NOT EDIT.

package analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchAnalyticsReportingSettingsReader is a Reader for the PatchAnalyticsReportingSettings structure.
type PatchAnalyticsReportingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAnalyticsReportingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchAnalyticsReportingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchAnalyticsReportingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchAnalyticsReportingSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchAnalyticsReportingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchAnalyticsReportingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchAnalyticsReportingSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchAnalyticsReportingSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchAnalyticsReportingSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchAnalyticsReportingSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchAnalyticsReportingSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchAnalyticsReportingSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchAnalyticsReportingSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchAnalyticsReportingSettingsOK creates a PatchAnalyticsReportingSettingsOK with default headers values
func NewPatchAnalyticsReportingSettingsOK() *PatchAnalyticsReportingSettingsOK {
	return &PatchAnalyticsReportingSettingsOK{}
}

/*PatchAnalyticsReportingSettingsOK handles this case with default header values.

successful operation
*/
type PatchAnalyticsReportingSettingsOK struct {
	Payload *models.AnalyticsReportingSettings
}

func (o *PatchAnalyticsReportingSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsOK) GetPayload() *models.AnalyticsReportingSettings {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsReportingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsBadRequest creates a PatchAnalyticsReportingSettingsBadRequest with default headers values
func NewPatchAnalyticsReportingSettingsBadRequest() *PatchAnalyticsReportingSettingsBadRequest {
	return &PatchAnalyticsReportingSettingsBadRequest{}
}

/*PatchAnalyticsReportingSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchAnalyticsReportingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsUnauthorized creates a PatchAnalyticsReportingSettingsUnauthorized with default headers values
func NewPatchAnalyticsReportingSettingsUnauthorized() *PatchAnalyticsReportingSettingsUnauthorized {
	return &PatchAnalyticsReportingSettingsUnauthorized{}
}

/*PatchAnalyticsReportingSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchAnalyticsReportingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsForbidden creates a PatchAnalyticsReportingSettingsForbidden with default headers values
func NewPatchAnalyticsReportingSettingsForbidden() *PatchAnalyticsReportingSettingsForbidden {
	return &PatchAnalyticsReportingSettingsForbidden{}
}

/*PatchAnalyticsReportingSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchAnalyticsReportingSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsNotFound creates a PatchAnalyticsReportingSettingsNotFound with default headers values
func NewPatchAnalyticsReportingSettingsNotFound() *PatchAnalyticsReportingSettingsNotFound {
	return &PatchAnalyticsReportingSettingsNotFound{}
}

/*PatchAnalyticsReportingSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchAnalyticsReportingSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsRequestTimeout creates a PatchAnalyticsReportingSettingsRequestTimeout with default headers values
func NewPatchAnalyticsReportingSettingsRequestTimeout() *PatchAnalyticsReportingSettingsRequestTimeout {
	return &PatchAnalyticsReportingSettingsRequestTimeout{}
}

/*PatchAnalyticsReportingSettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchAnalyticsReportingSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsRequestEntityTooLarge creates a PatchAnalyticsReportingSettingsRequestEntityTooLarge with default headers values
func NewPatchAnalyticsReportingSettingsRequestEntityTooLarge() *PatchAnalyticsReportingSettingsRequestEntityTooLarge {
	return &PatchAnalyticsReportingSettingsRequestEntityTooLarge{}
}

/*PatchAnalyticsReportingSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchAnalyticsReportingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsUnsupportedMediaType creates a PatchAnalyticsReportingSettingsUnsupportedMediaType with default headers values
func NewPatchAnalyticsReportingSettingsUnsupportedMediaType() *PatchAnalyticsReportingSettingsUnsupportedMediaType {
	return &PatchAnalyticsReportingSettingsUnsupportedMediaType{}
}

/*PatchAnalyticsReportingSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchAnalyticsReportingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsTooManyRequests creates a PatchAnalyticsReportingSettingsTooManyRequests with default headers values
func NewPatchAnalyticsReportingSettingsTooManyRequests() *PatchAnalyticsReportingSettingsTooManyRequests {
	return &PatchAnalyticsReportingSettingsTooManyRequests{}
}

/*PatchAnalyticsReportingSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchAnalyticsReportingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsInternalServerError creates a PatchAnalyticsReportingSettingsInternalServerError with default headers values
func NewPatchAnalyticsReportingSettingsInternalServerError() *PatchAnalyticsReportingSettingsInternalServerError {
	return &PatchAnalyticsReportingSettingsInternalServerError{}
}

/*PatchAnalyticsReportingSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchAnalyticsReportingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsServiceUnavailable creates a PatchAnalyticsReportingSettingsServiceUnavailable with default headers values
func NewPatchAnalyticsReportingSettingsServiceUnavailable() *PatchAnalyticsReportingSettingsServiceUnavailable {
	return &PatchAnalyticsReportingSettingsServiceUnavailable{}
}

/*PatchAnalyticsReportingSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchAnalyticsReportingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchAnalyticsReportingSettingsGatewayTimeout creates a PatchAnalyticsReportingSettingsGatewayTimeout with default headers values
func NewPatchAnalyticsReportingSettingsGatewayTimeout() *PatchAnalyticsReportingSettingsGatewayTimeout {
	return &PatchAnalyticsReportingSettingsGatewayTimeout{}
}

/*PatchAnalyticsReportingSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchAnalyticsReportingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchAnalyticsReportingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/analytics/reporting/settings][%d] patchAnalyticsReportingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchAnalyticsReportingSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchAnalyticsReportingSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
