// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchSpeechandtextanalyticsSettingsReader is a Reader for the PatchSpeechandtextanalyticsSettings structure.
type PatchSpeechandtextanalyticsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchSpeechandtextanalyticsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchSpeechandtextanalyticsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchSpeechandtextanalyticsSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchSpeechandtextanalyticsSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchSpeechandtextanalyticsSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchSpeechandtextanalyticsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchSpeechandtextanalyticsSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchSpeechandtextanalyticsSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchSpeechandtextanalyticsSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchSpeechandtextanalyticsSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchSpeechandtextanalyticsSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchSpeechandtextanalyticsSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchSpeechandtextanalyticsSettingsOK creates a PatchSpeechandtextanalyticsSettingsOK with default headers values
func NewPatchSpeechandtextanalyticsSettingsOK() *PatchSpeechandtextanalyticsSettingsOK {
	return &PatchSpeechandtextanalyticsSettingsOK{}
}

/*PatchSpeechandtextanalyticsSettingsOK handles this case with default header values.

Speech And Text Analytics settings has been updated
*/
type PatchSpeechandtextanalyticsSettingsOK struct {
	Payload *models.SpeechTextAnalyticsSettingsResponse
}

func (o *PatchSpeechandtextanalyticsSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsOK) GetPayload() *models.SpeechTextAnalyticsSettingsResponse {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpeechTextAnalyticsSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsBadRequest creates a PatchSpeechandtextanalyticsSettingsBadRequest with default headers values
func NewPatchSpeechandtextanalyticsSettingsBadRequest() *PatchSpeechandtextanalyticsSettingsBadRequest {
	return &PatchSpeechandtextanalyticsSettingsBadRequest{}
}

/*PatchSpeechandtextanalyticsSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchSpeechandtextanalyticsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsUnauthorized creates a PatchSpeechandtextanalyticsSettingsUnauthorized with default headers values
func NewPatchSpeechandtextanalyticsSettingsUnauthorized() *PatchSpeechandtextanalyticsSettingsUnauthorized {
	return &PatchSpeechandtextanalyticsSettingsUnauthorized{}
}

/*PatchSpeechandtextanalyticsSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchSpeechandtextanalyticsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsForbidden creates a PatchSpeechandtextanalyticsSettingsForbidden with default headers values
func NewPatchSpeechandtextanalyticsSettingsForbidden() *PatchSpeechandtextanalyticsSettingsForbidden {
	return &PatchSpeechandtextanalyticsSettingsForbidden{}
}

/*PatchSpeechandtextanalyticsSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchSpeechandtextanalyticsSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsNotFound creates a PatchSpeechandtextanalyticsSettingsNotFound with default headers values
func NewPatchSpeechandtextanalyticsSettingsNotFound() *PatchSpeechandtextanalyticsSettingsNotFound {
	return &PatchSpeechandtextanalyticsSettingsNotFound{}
}

/*PatchSpeechandtextanalyticsSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchSpeechandtextanalyticsSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsRequestEntityTooLarge creates a PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge with default headers values
func NewPatchSpeechandtextanalyticsSettingsRequestEntityTooLarge() *PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge {
	return &PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge{}
}

/*PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsUnsupportedMediaType creates a PatchSpeechandtextanalyticsSettingsUnsupportedMediaType with default headers values
func NewPatchSpeechandtextanalyticsSettingsUnsupportedMediaType() *PatchSpeechandtextanalyticsSettingsUnsupportedMediaType {
	return &PatchSpeechandtextanalyticsSettingsUnsupportedMediaType{}
}

/*PatchSpeechandtextanalyticsSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchSpeechandtextanalyticsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsTooManyRequests creates a PatchSpeechandtextanalyticsSettingsTooManyRequests with default headers values
func NewPatchSpeechandtextanalyticsSettingsTooManyRequests() *PatchSpeechandtextanalyticsSettingsTooManyRequests {
	return &PatchSpeechandtextanalyticsSettingsTooManyRequests{}
}

/*PatchSpeechandtextanalyticsSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchSpeechandtextanalyticsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsInternalServerError creates a PatchSpeechandtextanalyticsSettingsInternalServerError with default headers values
func NewPatchSpeechandtextanalyticsSettingsInternalServerError() *PatchSpeechandtextanalyticsSettingsInternalServerError {
	return &PatchSpeechandtextanalyticsSettingsInternalServerError{}
}

/*PatchSpeechandtextanalyticsSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchSpeechandtextanalyticsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsServiceUnavailable creates a PatchSpeechandtextanalyticsSettingsServiceUnavailable with default headers values
func NewPatchSpeechandtextanalyticsSettingsServiceUnavailable() *PatchSpeechandtextanalyticsSettingsServiceUnavailable {
	return &PatchSpeechandtextanalyticsSettingsServiceUnavailable{}
}

/*PatchSpeechandtextanalyticsSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchSpeechandtextanalyticsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchSpeechandtextanalyticsSettingsGatewayTimeout creates a PatchSpeechandtextanalyticsSettingsGatewayTimeout with default headers values
func NewPatchSpeechandtextanalyticsSettingsGatewayTimeout() *PatchSpeechandtextanalyticsSettingsGatewayTimeout {
	return &PatchSpeechandtextanalyticsSettingsGatewayTimeout{}
}

/*PatchSpeechandtextanalyticsSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchSpeechandtextanalyticsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchSpeechandtextanalyticsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/speechandtextanalytics/settings][%d] patchSpeechandtextanalyticsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchSpeechandtextanalyticsSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchSpeechandtextanalyticsSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
