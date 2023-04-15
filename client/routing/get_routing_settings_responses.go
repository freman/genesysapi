// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingSettingsReader is a Reader for the GetRoutingSettings structure.
type GetRoutingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSettingsOK creates a GetRoutingSettingsOK with default headers values
func NewGetRoutingSettingsOK() *GetRoutingSettingsOK {
	return &GetRoutingSettingsOK{}
}

/*
GetRoutingSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSettingsOK struct {
	Payload *models.RoutingSettings
}

// IsSuccess returns true when this get routing settings o k response has a 2xx status code
func (o *GetRoutingSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing settings o k response has a 3xx status code
func (o *GetRoutingSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings o k response has a 4xx status code
func (o *GetRoutingSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings o k response has a 5xx status code
func (o *GetRoutingSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings o k response a status code equal to that given
func (o *GetRoutingSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsOK) GetPayload() *models.RoutingSettings {
	return o.Payload
}

func (o *GetRoutingSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoutingSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsBadRequest creates a GetRoutingSettingsBadRequest with default headers values
func NewGetRoutingSettingsBadRequest() *GetRoutingSettingsBadRequest {
	return &GetRoutingSettingsBadRequest{}
}

/*
GetRoutingSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings bad request response has a 2xx status code
func (o *GetRoutingSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings bad request response has a 3xx status code
func (o *GetRoutingSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings bad request response has a 4xx status code
func (o *GetRoutingSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings bad request response has a 5xx status code
func (o *GetRoutingSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings bad request response a status code equal to that given
func (o *GetRoutingSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsUnauthorized creates a GetRoutingSettingsUnauthorized with default headers values
func NewGetRoutingSettingsUnauthorized() *GetRoutingSettingsUnauthorized {
	return &GetRoutingSettingsUnauthorized{}
}

/*
GetRoutingSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings unauthorized response has a 2xx status code
func (o *GetRoutingSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings unauthorized response has a 3xx status code
func (o *GetRoutingSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings unauthorized response has a 4xx status code
func (o *GetRoutingSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings unauthorized response has a 5xx status code
func (o *GetRoutingSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings unauthorized response a status code equal to that given
func (o *GetRoutingSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsForbidden creates a GetRoutingSettingsForbidden with default headers values
func NewGetRoutingSettingsForbidden() *GetRoutingSettingsForbidden {
	return &GetRoutingSettingsForbidden{}
}

/*
GetRoutingSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings forbidden response has a 2xx status code
func (o *GetRoutingSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings forbidden response has a 3xx status code
func (o *GetRoutingSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings forbidden response has a 4xx status code
func (o *GetRoutingSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings forbidden response has a 5xx status code
func (o *GetRoutingSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings forbidden response a status code equal to that given
func (o *GetRoutingSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsNotFound creates a GetRoutingSettingsNotFound with default headers values
func NewGetRoutingSettingsNotFound() *GetRoutingSettingsNotFound {
	return &GetRoutingSettingsNotFound{}
}

/*
GetRoutingSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings not found response has a 2xx status code
func (o *GetRoutingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings not found response has a 3xx status code
func (o *GetRoutingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings not found response has a 4xx status code
func (o *GetRoutingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings not found response has a 5xx status code
func (o *GetRoutingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings not found response a status code equal to that given
func (o *GetRoutingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsRequestTimeout creates a GetRoutingSettingsRequestTimeout with default headers values
func NewGetRoutingSettingsRequestTimeout() *GetRoutingSettingsRequestTimeout {
	return &GetRoutingSettingsRequestTimeout{}
}

/*
GetRoutingSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings request timeout response has a 2xx status code
func (o *GetRoutingSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings request timeout response has a 3xx status code
func (o *GetRoutingSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings request timeout response has a 4xx status code
func (o *GetRoutingSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings request timeout response has a 5xx status code
func (o *GetRoutingSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings request timeout response a status code equal to that given
func (o *GetRoutingSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsRequestEntityTooLarge creates a GetRoutingSettingsRequestEntityTooLarge with default headers values
func NewGetRoutingSettingsRequestEntityTooLarge() *GetRoutingSettingsRequestEntityTooLarge {
	return &GetRoutingSettingsRequestEntityTooLarge{}
}

/*
GetRoutingSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings request entity too large response has a 2xx status code
func (o *GetRoutingSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings request entity too large response has a 3xx status code
func (o *GetRoutingSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings request entity too large response has a 4xx status code
func (o *GetRoutingSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings request entity too large response has a 5xx status code
func (o *GetRoutingSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings request entity too large response a status code equal to that given
func (o *GetRoutingSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsUnsupportedMediaType creates a GetRoutingSettingsUnsupportedMediaType with default headers values
func NewGetRoutingSettingsUnsupportedMediaType() *GetRoutingSettingsUnsupportedMediaType {
	return &GetRoutingSettingsUnsupportedMediaType{}
}

/*
GetRoutingSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings unsupported media type response has a 2xx status code
func (o *GetRoutingSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings unsupported media type response has a 3xx status code
func (o *GetRoutingSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings unsupported media type response has a 4xx status code
func (o *GetRoutingSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings unsupported media type response has a 5xx status code
func (o *GetRoutingSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings unsupported media type response a status code equal to that given
func (o *GetRoutingSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTooManyRequests creates a GetRoutingSettingsTooManyRequests with default headers values
func NewGetRoutingSettingsTooManyRequests() *GetRoutingSettingsTooManyRequests {
	return &GetRoutingSettingsTooManyRequests{}
}

/*
GetRoutingSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings too many requests response has a 2xx status code
func (o *GetRoutingSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings too many requests response has a 3xx status code
func (o *GetRoutingSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings too many requests response has a 4xx status code
func (o *GetRoutingSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings too many requests response has a 5xx status code
func (o *GetRoutingSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings too many requests response a status code equal to that given
func (o *GetRoutingSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsInternalServerError creates a GetRoutingSettingsInternalServerError with default headers values
func NewGetRoutingSettingsInternalServerError() *GetRoutingSettingsInternalServerError {
	return &GetRoutingSettingsInternalServerError{}
}

/*
GetRoutingSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings internal server error response has a 2xx status code
func (o *GetRoutingSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings internal server error response has a 3xx status code
func (o *GetRoutingSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings internal server error response has a 4xx status code
func (o *GetRoutingSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings internal server error response has a 5xx status code
func (o *GetRoutingSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings internal server error response a status code equal to that given
func (o *GetRoutingSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsServiceUnavailable creates a GetRoutingSettingsServiceUnavailable with default headers values
func NewGetRoutingSettingsServiceUnavailable() *GetRoutingSettingsServiceUnavailable {
	return &GetRoutingSettingsServiceUnavailable{}
}

/*
GetRoutingSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings service unavailable response has a 2xx status code
func (o *GetRoutingSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings service unavailable response has a 3xx status code
func (o *GetRoutingSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings service unavailable response has a 4xx status code
func (o *GetRoutingSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings service unavailable response has a 5xx status code
func (o *GetRoutingSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings service unavailable response a status code equal to that given
func (o *GetRoutingSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsGatewayTimeout creates a GetRoutingSettingsGatewayTimeout with default headers values
func NewGetRoutingSettingsGatewayTimeout() *GetRoutingSettingsGatewayTimeout {
	return &GetRoutingSettingsGatewayTimeout{}
}

/*
GetRoutingSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings gateway timeout response has a 2xx status code
func (o *GetRoutingSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings gateway timeout response has a 3xx status code
func (o *GetRoutingSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings gateway timeout response has a 4xx status code
func (o *GetRoutingSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings gateway timeout response has a 5xx status code
func (o *GetRoutingSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings gateway timeout response a status code equal to that given
func (o *GetRoutingSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings][%d] getRoutingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
