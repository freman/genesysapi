// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutRecordingSettingsReader is a Reader for the PutRecordingSettings structure.
type PutRecordingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRecordingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRecordingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRecordingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRecordingSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRecordingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRecordingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRecordingSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRecordingSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRecordingSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRecordingSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRecordingSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRecordingSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRecordingSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRecordingSettingsOK creates a PutRecordingSettingsOK with default headers values
func NewPutRecordingSettingsOK() *PutRecordingSettingsOK {
	return &PutRecordingSettingsOK{}
}

/*PutRecordingSettingsOK handles this case with default header values.

successful operation
*/
type PutRecordingSettingsOK struct {
	Payload *models.RecordingSettings
}

func (o *PutRecordingSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsOK  %+v", 200, o.Payload)
}

func (o *PutRecordingSettingsOK) GetPayload() *models.RecordingSettings {
	return o.Payload
}

func (o *PutRecordingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsBadRequest creates a PutRecordingSettingsBadRequest with default headers values
func NewPutRecordingSettingsBadRequest() *PutRecordingSettingsBadRequest {
	return &PutRecordingSettingsBadRequest{}
}

/*PutRecordingSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRecordingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutRecordingSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsUnauthorized creates a PutRecordingSettingsUnauthorized with default headers values
func NewPutRecordingSettingsUnauthorized() *PutRecordingSettingsUnauthorized {
	return &PutRecordingSettingsUnauthorized{}
}

/*PutRecordingSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRecordingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRecordingSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsForbidden creates a PutRecordingSettingsForbidden with default headers values
func NewPutRecordingSettingsForbidden() *PutRecordingSettingsForbidden {
	return &PutRecordingSettingsForbidden{}
}

/*PutRecordingSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutRecordingSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutRecordingSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsNotFound creates a PutRecordingSettingsNotFound with default headers values
func NewPutRecordingSettingsNotFound() *PutRecordingSettingsNotFound {
	return &PutRecordingSettingsNotFound{}
}

/*PutRecordingSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutRecordingSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutRecordingSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsRequestTimeout creates a PutRecordingSettingsRequestTimeout with default headers values
func NewPutRecordingSettingsRequestTimeout() *PutRecordingSettingsRequestTimeout {
	return &PutRecordingSettingsRequestTimeout{}
}

/*PutRecordingSettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRecordingSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRecordingSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsRequestEntityTooLarge creates a PutRecordingSettingsRequestEntityTooLarge with default headers values
func NewPutRecordingSettingsRequestEntityTooLarge() *PutRecordingSettingsRequestEntityTooLarge {
	return &PutRecordingSettingsRequestEntityTooLarge{}
}

/*PutRecordingSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutRecordingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRecordingSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsUnsupportedMediaType creates a PutRecordingSettingsUnsupportedMediaType with default headers values
func NewPutRecordingSettingsUnsupportedMediaType() *PutRecordingSettingsUnsupportedMediaType {
	return &PutRecordingSettingsUnsupportedMediaType{}
}

/*PutRecordingSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRecordingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRecordingSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsTooManyRequests creates a PutRecordingSettingsTooManyRequests with default headers values
func NewPutRecordingSettingsTooManyRequests() *PutRecordingSettingsTooManyRequests {
	return &PutRecordingSettingsTooManyRequests{}
}

/*PutRecordingSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRecordingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRecordingSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsInternalServerError creates a PutRecordingSettingsInternalServerError with default headers values
func NewPutRecordingSettingsInternalServerError() *PutRecordingSettingsInternalServerError {
	return &PutRecordingSettingsInternalServerError{}
}

/*PutRecordingSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRecordingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRecordingSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsServiceUnavailable creates a PutRecordingSettingsServiceUnavailable with default headers values
func NewPutRecordingSettingsServiceUnavailable() *PutRecordingSettingsServiceUnavailable {
	return &PutRecordingSettingsServiceUnavailable{}
}

/*PutRecordingSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRecordingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRecordingSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRecordingSettingsGatewayTimeout creates a PutRecordingSettingsGatewayTimeout with default headers values
func NewPutRecordingSettingsGatewayTimeout() *PutRecordingSettingsGatewayTimeout {
	return &PutRecordingSettingsGatewayTimeout{}
}

/*PutRecordingSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutRecordingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRecordingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/recording/settings][%d] putRecordingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRecordingSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRecordingSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
