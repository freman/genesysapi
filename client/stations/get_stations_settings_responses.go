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

// GetStationsSettingsReader is a Reader for the GetStationsSettings structure.
type GetStationsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStationsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStationsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetStationsSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetStationsSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetStationsSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetStationsSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetStationsSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetStationsSettingsGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetStationsSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetStationsSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetStationsSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetStationsSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetStationsSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetStationsSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStationsSettingsOK creates a GetStationsSettingsOK with default headers values
func NewGetStationsSettingsOK() *GetStationsSettingsOK {
	return &GetStationsSettingsOK{}
}

/*GetStationsSettingsOK handles this case with default header values.

successful operation
*/
type GetStationsSettingsOK struct {
	Payload *models.StationSettings
}

func (o *GetStationsSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsOK  %+v", 200, o.Payload)
}

func (o *GetStationsSettingsOK) GetPayload() *models.StationSettings {
	return o.Payload
}

func (o *GetStationsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StationSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsBadRequest creates a GetStationsSettingsBadRequest with default headers values
func NewGetStationsSettingsBadRequest() *GetStationsSettingsBadRequest {
	return &GetStationsSettingsBadRequest{}
}

/*GetStationsSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetStationsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationsSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsUnauthorized creates a GetStationsSettingsUnauthorized with default headers values
func NewGetStationsSettingsUnauthorized() *GetStationsSettingsUnauthorized {
	return &GetStationsSettingsUnauthorized{}
}

/*GetStationsSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetStationsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationsSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsForbidden creates a GetStationsSettingsForbidden with default headers values
func NewGetStationsSettingsForbidden() *GetStationsSettingsForbidden {
	return &GetStationsSettingsForbidden{}
}

/*GetStationsSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetStationsSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetStationsSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsNotFound creates a GetStationsSettingsNotFound with default headers values
func NewGetStationsSettingsNotFound() *GetStationsSettingsNotFound {
	return &GetStationsSettingsNotFound{}
}

/*GetStationsSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetStationsSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetStationsSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsRequestTimeout creates a GetStationsSettingsRequestTimeout with default headers values
func NewGetStationsSettingsRequestTimeout() *GetStationsSettingsRequestTimeout {
	return &GetStationsSettingsRequestTimeout{}
}

/*GetStationsSettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetStationsSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationsSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsGone creates a GetStationsSettingsGone with default headers values
func NewGetStationsSettingsGone() *GetStationsSettingsGone {
	return &GetStationsSettingsGone{}
}

/*GetStationsSettingsGone handles this case with default header values.

Gone
*/
type GetStationsSettingsGone struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsGone  %+v", 410, o.Payload)
}

func (o *GetStationsSettingsGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsRequestEntityTooLarge creates a GetStationsSettingsRequestEntityTooLarge with default headers values
func NewGetStationsSettingsRequestEntityTooLarge() *GetStationsSettingsRequestEntityTooLarge {
	return &GetStationsSettingsRequestEntityTooLarge{}
}

/*GetStationsSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetStationsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationsSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsUnsupportedMediaType creates a GetStationsSettingsUnsupportedMediaType with default headers values
func NewGetStationsSettingsUnsupportedMediaType() *GetStationsSettingsUnsupportedMediaType {
	return &GetStationsSettingsUnsupportedMediaType{}
}

/*GetStationsSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetStationsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationsSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsTooManyRequests creates a GetStationsSettingsTooManyRequests with default headers values
func NewGetStationsSettingsTooManyRequests() *GetStationsSettingsTooManyRequests {
	return &GetStationsSettingsTooManyRequests{}
}

/*GetStationsSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetStationsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationsSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsInternalServerError creates a GetStationsSettingsInternalServerError with default headers values
func NewGetStationsSettingsInternalServerError() *GetStationsSettingsInternalServerError {
	return &GetStationsSettingsInternalServerError{}
}

/*GetStationsSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetStationsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationsSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsServiceUnavailable creates a GetStationsSettingsServiceUnavailable with default headers values
func NewGetStationsSettingsServiceUnavailable() *GetStationsSettingsServiceUnavailable {
	return &GetStationsSettingsServiceUnavailable{}
}

/*GetStationsSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetStationsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationsSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStationsSettingsGatewayTimeout creates a GetStationsSettingsGatewayTimeout with default headers values
func NewGetStationsSettingsGatewayTimeout() *GetStationsSettingsGatewayTimeout {
	return &GetStationsSettingsGatewayTimeout{}
}

/*GetStationsSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetStationsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetStationsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationsSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetStationsSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
