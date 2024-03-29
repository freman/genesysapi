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

/*
GetStationsSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetStationsSettingsOK struct {
	Payload *models.StationSettings
}

// IsSuccess returns true when this get stations settings o k response has a 2xx status code
func (o *GetStationsSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get stations settings o k response has a 3xx status code
func (o *GetStationsSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings o k response has a 4xx status code
func (o *GetStationsSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations settings o k response has a 5xx status code
func (o *GetStationsSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings o k response a status code equal to that given
func (o *GetStationsSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetStationsSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsOK  %+v", 200, o.Payload)
}

func (o *GetStationsSettingsOK) String() string {
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

/*
GetStationsSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetStationsSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings bad request response has a 2xx status code
func (o *GetStationsSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings bad request response has a 3xx status code
func (o *GetStationsSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings bad request response has a 4xx status code
func (o *GetStationsSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings bad request response has a 5xx status code
func (o *GetStationsSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings bad request response a status code equal to that given
func (o *GetStationsSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetStationsSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetStationsSettingsBadRequest) String() string {
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

/*
GetStationsSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetStationsSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings unauthorized response has a 2xx status code
func (o *GetStationsSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings unauthorized response has a 3xx status code
func (o *GetStationsSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings unauthorized response has a 4xx status code
func (o *GetStationsSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings unauthorized response has a 5xx status code
func (o *GetStationsSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings unauthorized response a status code equal to that given
func (o *GetStationsSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetStationsSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStationsSettingsUnauthorized) String() string {
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

/*
GetStationsSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetStationsSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings forbidden response has a 2xx status code
func (o *GetStationsSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings forbidden response has a 3xx status code
func (o *GetStationsSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings forbidden response has a 4xx status code
func (o *GetStationsSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings forbidden response has a 5xx status code
func (o *GetStationsSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings forbidden response a status code equal to that given
func (o *GetStationsSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetStationsSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetStationsSettingsForbidden) String() string {
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

/*
GetStationsSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetStationsSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings not found response has a 2xx status code
func (o *GetStationsSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings not found response has a 3xx status code
func (o *GetStationsSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings not found response has a 4xx status code
func (o *GetStationsSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings not found response has a 5xx status code
func (o *GetStationsSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings not found response a status code equal to that given
func (o *GetStationsSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetStationsSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetStationsSettingsNotFound) String() string {
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

/*
GetStationsSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetStationsSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings request timeout response has a 2xx status code
func (o *GetStationsSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings request timeout response has a 3xx status code
func (o *GetStationsSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings request timeout response has a 4xx status code
func (o *GetStationsSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings request timeout response has a 5xx status code
func (o *GetStationsSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings request timeout response a status code equal to that given
func (o *GetStationsSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetStationsSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetStationsSettingsRequestTimeout) String() string {
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

/*
GetStationsSettingsGone describes a response with status code 410, with default header values.

Gone
*/
type GetStationsSettingsGone struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings gone response has a 2xx status code
func (o *GetStationsSettingsGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings gone response has a 3xx status code
func (o *GetStationsSettingsGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings gone response has a 4xx status code
func (o *GetStationsSettingsGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings gone response has a 5xx status code
func (o *GetStationsSettingsGone) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings gone response a status code equal to that given
func (o *GetStationsSettingsGone) IsCode(code int) bool {
	return code == 410
}

func (o *GetStationsSettingsGone) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsGone  %+v", 410, o.Payload)
}

func (o *GetStationsSettingsGone) String() string {
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

/*
GetStationsSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetStationsSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings request entity too large response has a 2xx status code
func (o *GetStationsSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings request entity too large response has a 3xx status code
func (o *GetStationsSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings request entity too large response has a 4xx status code
func (o *GetStationsSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings request entity too large response has a 5xx status code
func (o *GetStationsSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings request entity too large response a status code equal to that given
func (o *GetStationsSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetStationsSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetStationsSettingsRequestEntityTooLarge) String() string {
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

/*
GetStationsSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetStationsSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings unsupported media type response has a 2xx status code
func (o *GetStationsSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings unsupported media type response has a 3xx status code
func (o *GetStationsSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings unsupported media type response has a 4xx status code
func (o *GetStationsSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings unsupported media type response has a 5xx status code
func (o *GetStationsSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings unsupported media type response a status code equal to that given
func (o *GetStationsSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetStationsSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetStationsSettingsUnsupportedMediaType) String() string {
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

/*
GetStationsSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetStationsSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings too many requests response has a 2xx status code
func (o *GetStationsSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings too many requests response has a 3xx status code
func (o *GetStationsSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings too many requests response has a 4xx status code
func (o *GetStationsSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get stations settings too many requests response has a 5xx status code
func (o *GetStationsSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get stations settings too many requests response a status code equal to that given
func (o *GetStationsSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetStationsSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetStationsSettingsTooManyRequests) String() string {
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

/*
GetStationsSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetStationsSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings internal server error response has a 2xx status code
func (o *GetStationsSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings internal server error response has a 3xx status code
func (o *GetStationsSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings internal server error response has a 4xx status code
func (o *GetStationsSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations settings internal server error response has a 5xx status code
func (o *GetStationsSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations settings internal server error response a status code equal to that given
func (o *GetStationsSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetStationsSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStationsSettingsInternalServerError) String() string {
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

/*
GetStationsSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetStationsSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings service unavailable response has a 2xx status code
func (o *GetStationsSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings service unavailable response has a 3xx status code
func (o *GetStationsSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings service unavailable response has a 4xx status code
func (o *GetStationsSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations settings service unavailable response has a 5xx status code
func (o *GetStationsSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations settings service unavailable response a status code equal to that given
func (o *GetStationsSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetStationsSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetStationsSettingsServiceUnavailable) String() string {
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

/*
GetStationsSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetStationsSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get stations settings gateway timeout response has a 2xx status code
func (o *GetStationsSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get stations settings gateway timeout response has a 3xx status code
func (o *GetStationsSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get stations settings gateway timeout response has a 4xx status code
func (o *GetStationsSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get stations settings gateway timeout response has a 5xx status code
func (o *GetStationsSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get stations settings gateway timeout response a status code equal to that given
func (o *GetStationsSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetStationsSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/stations/settings][%d] getStationsSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetStationsSettingsGatewayTimeout) String() string {
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
