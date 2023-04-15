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

// GetAnalyticsDataretentionSettingsReader is a Reader for the GetAnalyticsDataretentionSettings structure.
type GetAnalyticsDataretentionSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAnalyticsDataretentionSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAnalyticsDataretentionSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAnalyticsDataretentionSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAnalyticsDataretentionSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAnalyticsDataretentionSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAnalyticsDataretentionSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAnalyticsDataretentionSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAnalyticsDataretentionSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAnalyticsDataretentionSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAnalyticsDataretentionSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAnalyticsDataretentionSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAnalyticsDataretentionSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAnalyticsDataretentionSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAnalyticsDataretentionSettingsOK creates a GetAnalyticsDataretentionSettingsOK with default headers values
func NewGetAnalyticsDataretentionSettingsOK() *GetAnalyticsDataretentionSettingsOK {
	return &GetAnalyticsDataretentionSettingsOK{}
}

/*
GetAnalyticsDataretentionSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAnalyticsDataretentionSettingsOK struct {
	Payload *models.AnalyticsDataRetentionResponse
}

// IsSuccess returns true when this get analytics dataretention settings o k response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get analytics dataretention settings o k response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings o k response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics dataretention settings o k response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings o k response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAnalyticsDataretentionSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsOK) GetPayload() *models.AnalyticsDataRetentionResponse {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsDataRetentionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsBadRequest creates a GetAnalyticsDataretentionSettingsBadRequest with default headers values
func NewGetAnalyticsDataretentionSettingsBadRequest() *GetAnalyticsDataretentionSettingsBadRequest {
	return &GetAnalyticsDataretentionSettingsBadRequest{}
}

/*
GetAnalyticsDataretentionSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAnalyticsDataretentionSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings bad request response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings bad request response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings bad request response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings bad request response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings bad request response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAnalyticsDataretentionSettingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsUnauthorized creates a GetAnalyticsDataretentionSettingsUnauthorized with default headers values
func NewGetAnalyticsDataretentionSettingsUnauthorized() *GetAnalyticsDataretentionSettingsUnauthorized {
	return &GetAnalyticsDataretentionSettingsUnauthorized{}
}

/*
GetAnalyticsDataretentionSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAnalyticsDataretentionSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings unauthorized response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings unauthorized response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings unauthorized response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings unauthorized response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings unauthorized response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetAnalyticsDataretentionSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsForbidden creates a GetAnalyticsDataretentionSettingsForbidden with default headers values
func NewGetAnalyticsDataretentionSettingsForbidden() *GetAnalyticsDataretentionSettingsForbidden {
	return &GetAnalyticsDataretentionSettingsForbidden{}
}

/*
GetAnalyticsDataretentionSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetAnalyticsDataretentionSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings forbidden response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings forbidden response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings forbidden response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings forbidden response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings forbidden response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetAnalyticsDataretentionSettingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsForbidden  %+v", 403, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsNotFound creates a GetAnalyticsDataretentionSettingsNotFound with default headers values
func NewGetAnalyticsDataretentionSettingsNotFound() *GetAnalyticsDataretentionSettingsNotFound {
	return &GetAnalyticsDataretentionSettingsNotFound{}
}

/*
GetAnalyticsDataretentionSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetAnalyticsDataretentionSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings not found response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings not found response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings not found response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings not found response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings not found response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAnalyticsDataretentionSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsRequestTimeout creates a GetAnalyticsDataretentionSettingsRequestTimeout with default headers values
func NewGetAnalyticsDataretentionSettingsRequestTimeout() *GetAnalyticsDataretentionSettingsRequestTimeout {
	return &GetAnalyticsDataretentionSettingsRequestTimeout{}
}

/*
GetAnalyticsDataretentionSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAnalyticsDataretentionSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings request timeout response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings request timeout response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings request timeout response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings request timeout response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings request timeout response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetAnalyticsDataretentionSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsRequestEntityTooLarge creates a GetAnalyticsDataretentionSettingsRequestEntityTooLarge with default headers values
func NewGetAnalyticsDataretentionSettingsRequestEntityTooLarge() *GetAnalyticsDataretentionSettingsRequestEntityTooLarge {
	return &GetAnalyticsDataretentionSettingsRequestEntityTooLarge{}
}

/*
GetAnalyticsDataretentionSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetAnalyticsDataretentionSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings request entity too large response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings request entity too large response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings request entity too large response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings request entity too large response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings request entity too large response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsUnsupportedMediaType creates a GetAnalyticsDataretentionSettingsUnsupportedMediaType with default headers values
func NewGetAnalyticsDataretentionSettingsUnsupportedMediaType() *GetAnalyticsDataretentionSettingsUnsupportedMediaType {
	return &GetAnalyticsDataretentionSettingsUnsupportedMediaType{}
}

/*
GetAnalyticsDataretentionSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAnalyticsDataretentionSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings unsupported media type response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings unsupported media type response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings unsupported media type response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings unsupported media type response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings unsupported media type response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsTooManyRequests creates a GetAnalyticsDataretentionSettingsTooManyRequests with default headers values
func NewGetAnalyticsDataretentionSettingsTooManyRequests() *GetAnalyticsDataretentionSettingsTooManyRequests {
	return &GetAnalyticsDataretentionSettingsTooManyRequests{}
}

/*
GetAnalyticsDataretentionSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAnalyticsDataretentionSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings too many requests response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings too many requests response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings too many requests response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get analytics dataretention settings too many requests response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get analytics dataretention settings too many requests response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetAnalyticsDataretentionSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsInternalServerError creates a GetAnalyticsDataretentionSettingsInternalServerError with default headers values
func NewGetAnalyticsDataretentionSettingsInternalServerError() *GetAnalyticsDataretentionSettingsInternalServerError {
	return &GetAnalyticsDataretentionSettingsInternalServerError{}
}

/*
GetAnalyticsDataretentionSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAnalyticsDataretentionSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings internal server error response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings internal server error response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings internal server error response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics dataretention settings internal server error response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics dataretention settings internal server error response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetAnalyticsDataretentionSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsServiceUnavailable creates a GetAnalyticsDataretentionSettingsServiceUnavailable with default headers values
func NewGetAnalyticsDataretentionSettingsServiceUnavailable() *GetAnalyticsDataretentionSettingsServiceUnavailable {
	return &GetAnalyticsDataretentionSettingsServiceUnavailable{}
}

/*
GetAnalyticsDataretentionSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAnalyticsDataretentionSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings service unavailable response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings service unavailable response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings service unavailable response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics dataretention settings service unavailable response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics dataretention settings service unavailable response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAnalyticsDataretentionSettingsGatewayTimeout creates a GetAnalyticsDataretentionSettingsGatewayTimeout with default headers values
func NewGetAnalyticsDataretentionSettingsGatewayTimeout() *GetAnalyticsDataretentionSettingsGatewayTimeout {
	return &GetAnalyticsDataretentionSettingsGatewayTimeout{}
}

/*
GetAnalyticsDataretentionSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetAnalyticsDataretentionSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get analytics dataretention settings gateway timeout response has a 2xx status code
func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get analytics dataretention settings gateway timeout response has a 3xx status code
func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get analytics dataretention settings gateway timeout response has a 4xx status code
func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get analytics dataretention settings gateway timeout response has a 5xx status code
func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get analytics dataretention settings gateway timeout response a status code equal to that given
func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/analytics/dataretention/settings][%d] getAnalyticsDataretentionSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAnalyticsDataretentionSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
