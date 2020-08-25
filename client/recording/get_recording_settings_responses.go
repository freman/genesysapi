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

// GetRecordingSettingsReader is a Reader for the GetRecordingSettings structure.
type GetRecordingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingSettingsOK creates a GetRecordingSettingsOK with default headers values
func NewGetRecordingSettingsOK() *GetRecordingSettingsOK {
	return &GetRecordingSettingsOK{}
}

/*GetRecordingSettingsOK handles this case with default header values.

successful operation
*/
type GetRecordingSettingsOK struct {
	Payload *models.RecordingSettings
}

func (o *GetRecordingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingSettingsOK) GetPayload() *models.RecordingSettings {
	return o.Payload
}

func (o *GetRecordingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsBadRequest creates a GetRecordingSettingsBadRequest with default headers values
func NewGetRecordingSettingsBadRequest() *GetRecordingSettingsBadRequest {
	return &GetRecordingSettingsBadRequest{}
}

/*GetRecordingSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsUnauthorized creates a GetRecordingSettingsUnauthorized with default headers values
func NewGetRecordingSettingsUnauthorized() *GetRecordingSettingsUnauthorized {
	return &GetRecordingSettingsUnauthorized{}
}

/*GetRecordingSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsForbidden creates a GetRecordingSettingsForbidden with default headers values
func NewGetRecordingSettingsForbidden() *GetRecordingSettingsForbidden {
	return &GetRecordingSettingsForbidden{}
}

/*GetRecordingSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsNotFound creates a GetRecordingSettingsNotFound with default headers values
func NewGetRecordingSettingsNotFound() *GetRecordingSettingsNotFound {
	return &GetRecordingSettingsNotFound{}
}

/*GetRecordingSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRecordingSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsRequestEntityTooLarge creates a GetRecordingSettingsRequestEntityTooLarge with default headers values
func NewGetRecordingSettingsRequestEntityTooLarge() *GetRecordingSettingsRequestEntityTooLarge {
	return &GetRecordingSettingsRequestEntityTooLarge{}
}

/*GetRecordingSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRecordingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsUnsupportedMediaType creates a GetRecordingSettingsUnsupportedMediaType with default headers values
func NewGetRecordingSettingsUnsupportedMediaType() *GetRecordingSettingsUnsupportedMediaType {
	return &GetRecordingSettingsUnsupportedMediaType{}
}

/*GetRecordingSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsTooManyRequests creates a GetRecordingSettingsTooManyRequests with default headers values
func NewGetRecordingSettingsTooManyRequests() *GetRecordingSettingsTooManyRequests {
	return &GetRecordingSettingsTooManyRequests{}
}

/*GetRecordingSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRecordingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsInternalServerError creates a GetRecordingSettingsInternalServerError with default headers values
func NewGetRecordingSettingsInternalServerError() *GetRecordingSettingsInternalServerError {
	return &GetRecordingSettingsInternalServerError{}
}

/*GetRecordingSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsServiceUnavailable creates a GetRecordingSettingsServiceUnavailable with default headers values
func NewGetRecordingSettingsServiceUnavailable() *GetRecordingSettingsServiceUnavailable {
	return &GetRecordingSettingsServiceUnavailable{}
}

/*GetRecordingSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingSettingsGatewayTimeout creates a GetRecordingSettingsGatewayTimeout with default headers values
func NewGetRecordingSettingsGatewayTimeout() *GetRecordingSettingsGatewayTimeout {
	return &GetRecordingSettingsGatewayTimeout{}
}

/*GetRecordingSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRecordingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}