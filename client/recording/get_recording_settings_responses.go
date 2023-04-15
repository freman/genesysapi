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
	case 408:
		result := NewGetRecordingSettingsRequestTimeout()
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

/*
GetRecordingSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRecordingSettingsOK struct {
	Payload *models.RecordingSettings
}

// IsSuccess returns true when this get recording settings o k response has a 2xx status code
func (o *GetRecordingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get recording settings o k response has a 3xx status code
func (o *GetRecordingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings o k response has a 4xx status code
func (o *GetRecordingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording settings o k response has a 5xx status code
func (o *GetRecordingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings o k response a status code equal to that given
func (o *GetRecordingSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRecordingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRecordingSettingsOK) String() string {
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

/*
GetRecordingSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings bad request response has a 2xx status code
func (o *GetRecordingSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings bad request response has a 3xx status code
func (o *GetRecordingSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings bad request response has a 4xx status code
func (o *GetRecordingSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings bad request response has a 5xx status code
func (o *GetRecordingSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings bad request response a status code equal to that given
func (o *GetRecordingSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRecordingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingSettingsBadRequest) String() string {
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

/*
GetRecordingSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings unauthorized response has a 2xx status code
func (o *GetRecordingSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings unauthorized response has a 3xx status code
func (o *GetRecordingSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings unauthorized response has a 4xx status code
func (o *GetRecordingSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings unauthorized response has a 5xx status code
func (o *GetRecordingSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings unauthorized response a status code equal to that given
func (o *GetRecordingSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRecordingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingSettingsUnauthorized) String() string {
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

/*
GetRecordingSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings forbidden response has a 2xx status code
func (o *GetRecordingSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings forbidden response has a 3xx status code
func (o *GetRecordingSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings forbidden response has a 4xx status code
func (o *GetRecordingSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings forbidden response has a 5xx status code
func (o *GetRecordingSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings forbidden response a status code equal to that given
func (o *GetRecordingSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRecordingSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingSettingsForbidden) String() string {
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

/*
GetRecordingSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRecordingSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings not found response has a 2xx status code
func (o *GetRecordingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings not found response has a 3xx status code
func (o *GetRecordingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings not found response has a 4xx status code
func (o *GetRecordingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings not found response has a 5xx status code
func (o *GetRecordingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings not found response a status code equal to that given
func (o *GetRecordingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRecordingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingSettingsNotFound) String() string {
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

// NewGetRecordingSettingsRequestTimeout creates a GetRecordingSettingsRequestTimeout with default headers values
func NewGetRecordingSettingsRequestTimeout() *GetRecordingSettingsRequestTimeout {
	return &GetRecordingSettingsRequestTimeout{}
}

/*
GetRecordingSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings request timeout response has a 2xx status code
func (o *GetRecordingSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings request timeout response has a 3xx status code
func (o *GetRecordingSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings request timeout response has a 4xx status code
func (o *GetRecordingSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings request timeout response has a 5xx status code
func (o *GetRecordingSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings request timeout response a status code equal to that given
func (o *GetRecordingSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRecordingSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetRecordingSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRecordingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings request entity too large response has a 2xx status code
func (o *GetRecordingSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings request entity too large response has a 3xx status code
func (o *GetRecordingSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings request entity too large response has a 4xx status code
func (o *GetRecordingSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings request entity too large response has a 5xx status code
func (o *GetRecordingSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings request entity too large response a status code equal to that given
func (o *GetRecordingSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRecordingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingSettingsRequestEntityTooLarge) String() string {
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

/*
GetRecordingSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings unsupported media type response has a 2xx status code
func (o *GetRecordingSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings unsupported media type response has a 3xx status code
func (o *GetRecordingSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings unsupported media type response has a 4xx status code
func (o *GetRecordingSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings unsupported media type response has a 5xx status code
func (o *GetRecordingSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings unsupported media type response a status code equal to that given
func (o *GetRecordingSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRecordingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingSettingsUnsupportedMediaType) String() string {
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

/*
GetRecordingSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings too many requests response has a 2xx status code
func (o *GetRecordingSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings too many requests response has a 3xx status code
func (o *GetRecordingSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings too many requests response has a 4xx status code
func (o *GetRecordingSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get recording settings too many requests response has a 5xx status code
func (o *GetRecordingSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get recording settings too many requests response a status code equal to that given
func (o *GetRecordingSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRecordingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingSettingsTooManyRequests) String() string {
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

/*
GetRecordingSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings internal server error response has a 2xx status code
func (o *GetRecordingSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings internal server error response has a 3xx status code
func (o *GetRecordingSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings internal server error response has a 4xx status code
func (o *GetRecordingSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording settings internal server error response has a 5xx status code
func (o *GetRecordingSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording settings internal server error response a status code equal to that given
func (o *GetRecordingSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRecordingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingSettingsInternalServerError) String() string {
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

/*
GetRecordingSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings service unavailable response has a 2xx status code
func (o *GetRecordingSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings service unavailable response has a 3xx status code
func (o *GetRecordingSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings service unavailable response has a 4xx status code
func (o *GetRecordingSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording settings service unavailable response has a 5xx status code
func (o *GetRecordingSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording settings service unavailable response a status code equal to that given
func (o *GetRecordingSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRecordingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingSettingsServiceUnavailable) String() string {
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

/*
GetRecordingSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRecordingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get recording settings gateway timeout response has a 2xx status code
func (o *GetRecordingSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get recording settings gateway timeout response has a 3xx status code
func (o *GetRecordingSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get recording settings gateway timeout response has a 4xx status code
func (o *GetRecordingSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get recording settings gateway timeout response has a 5xx status code
func (o *GetRecordingSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get recording settings gateway timeout response a status code equal to that given
func (o *GetRecordingSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRecordingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/settings][%d] getRecordingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingSettingsGatewayTimeout) String() string {
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
