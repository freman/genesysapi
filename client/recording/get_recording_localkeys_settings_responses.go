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

// GetRecordingLocalkeysSettingsReader is a Reader for the GetRecordingLocalkeysSettings structure.
type GetRecordingLocalkeysSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingLocalkeysSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingLocalkeysSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingLocalkeysSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingLocalkeysSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingLocalkeysSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingLocalkeysSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingLocalkeysSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingLocalkeysSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingLocalkeysSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingLocalkeysSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingLocalkeysSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingLocalkeysSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingLocalkeysSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingLocalkeysSettingsOK creates a GetRecordingLocalkeysSettingsOK with default headers values
func NewGetRecordingLocalkeysSettingsOK() *GetRecordingLocalkeysSettingsOK {
	return &GetRecordingLocalkeysSettingsOK{}
}

/*GetRecordingLocalkeysSettingsOK handles this case with default header values.

successful operation
*/
type GetRecordingLocalkeysSettingsOK struct {
	Payload *models.LocalEncryptionConfigurationListing
}

func (o *GetRecordingLocalkeysSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsOK) GetPayload() *models.LocalEncryptionConfigurationListing {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocalEncryptionConfigurationListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsBadRequest creates a GetRecordingLocalkeysSettingsBadRequest with default headers values
func NewGetRecordingLocalkeysSettingsBadRequest() *GetRecordingLocalkeysSettingsBadRequest {
	return &GetRecordingLocalkeysSettingsBadRequest{}
}

/*GetRecordingLocalkeysSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingLocalkeysSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsUnauthorized creates a GetRecordingLocalkeysSettingsUnauthorized with default headers values
func NewGetRecordingLocalkeysSettingsUnauthorized() *GetRecordingLocalkeysSettingsUnauthorized {
	return &GetRecordingLocalkeysSettingsUnauthorized{}
}

/*GetRecordingLocalkeysSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingLocalkeysSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsForbidden creates a GetRecordingLocalkeysSettingsForbidden with default headers values
func NewGetRecordingLocalkeysSettingsForbidden() *GetRecordingLocalkeysSettingsForbidden {
	return &GetRecordingLocalkeysSettingsForbidden{}
}

/*GetRecordingLocalkeysSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingLocalkeysSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsNotFound creates a GetRecordingLocalkeysSettingsNotFound with default headers values
func NewGetRecordingLocalkeysSettingsNotFound() *GetRecordingLocalkeysSettingsNotFound {
	return &GetRecordingLocalkeysSettingsNotFound{}
}

/*GetRecordingLocalkeysSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRecordingLocalkeysSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsRequestTimeout creates a GetRecordingLocalkeysSettingsRequestTimeout with default headers values
func NewGetRecordingLocalkeysSettingsRequestTimeout() *GetRecordingLocalkeysSettingsRequestTimeout {
	return &GetRecordingLocalkeysSettingsRequestTimeout{}
}

/*GetRecordingLocalkeysSettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingLocalkeysSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsRequestEntityTooLarge creates a GetRecordingLocalkeysSettingsRequestEntityTooLarge with default headers values
func NewGetRecordingLocalkeysSettingsRequestEntityTooLarge() *GetRecordingLocalkeysSettingsRequestEntityTooLarge {
	return &GetRecordingLocalkeysSettingsRequestEntityTooLarge{}
}

/*GetRecordingLocalkeysSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRecordingLocalkeysSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsUnsupportedMediaType creates a GetRecordingLocalkeysSettingsUnsupportedMediaType with default headers values
func NewGetRecordingLocalkeysSettingsUnsupportedMediaType() *GetRecordingLocalkeysSettingsUnsupportedMediaType {
	return &GetRecordingLocalkeysSettingsUnsupportedMediaType{}
}

/*GetRecordingLocalkeysSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingLocalkeysSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsTooManyRequests creates a GetRecordingLocalkeysSettingsTooManyRequests with default headers values
func NewGetRecordingLocalkeysSettingsTooManyRequests() *GetRecordingLocalkeysSettingsTooManyRequests {
	return &GetRecordingLocalkeysSettingsTooManyRequests{}
}

/*GetRecordingLocalkeysSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingLocalkeysSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsInternalServerError creates a GetRecordingLocalkeysSettingsInternalServerError with default headers values
func NewGetRecordingLocalkeysSettingsInternalServerError() *GetRecordingLocalkeysSettingsInternalServerError {
	return &GetRecordingLocalkeysSettingsInternalServerError{}
}

/*GetRecordingLocalkeysSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingLocalkeysSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsServiceUnavailable creates a GetRecordingLocalkeysSettingsServiceUnavailable with default headers values
func NewGetRecordingLocalkeysSettingsServiceUnavailable() *GetRecordingLocalkeysSettingsServiceUnavailable {
	return &GetRecordingLocalkeysSettingsServiceUnavailable{}
}

/*GetRecordingLocalkeysSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingLocalkeysSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingLocalkeysSettingsGatewayTimeout creates a GetRecordingLocalkeysSettingsGatewayTimeout with default headers values
func NewGetRecordingLocalkeysSettingsGatewayTimeout() *GetRecordingLocalkeysSettingsGatewayTimeout {
	return &GetRecordingLocalkeysSettingsGatewayTimeout{}
}

/*GetRecordingLocalkeysSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRecordingLocalkeysSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingLocalkeysSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/localkeys/settings][%d] getRecordingLocalkeysSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingLocalkeysSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingLocalkeysSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
