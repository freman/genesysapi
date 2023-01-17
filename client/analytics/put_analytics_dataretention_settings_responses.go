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

// PutAnalyticsDataretentionSettingsReader is a Reader for the PutAnalyticsDataretentionSettings structure.
type PutAnalyticsDataretentionSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAnalyticsDataretentionSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAnalyticsDataretentionSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutAnalyticsDataretentionSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutAnalyticsDataretentionSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutAnalyticsDataretentionSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutAnalyticsDataretentionSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutAnalyticsDataretentionSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutAnalyticsDataretentionSettingsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutAnalyticsDataretentionSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutAnalyticsDataretentionSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutAnalyticsDataretentionSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutAnalyticsDataretentionSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutAnalyticsDataretentionSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutAnalyticsDataretentionSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutAnalyticsDataretentionSettingsOK creates a PutAnalyticsDataretentionSettingsOK with default headers values
func NewPutAnalyticsDataretentionSettingsOK() *PutAnalyticsDataretentionSettingsOK {
	return &PutAnalyticsDataretentionSettingsOK{}
}

/*
PutAnalyticsDataretentionSettingsOK describes a response with status code 200, with default header values.

successful operation
*/
type PutAnalyticsDataretentionSettingsOK struct {
	Payload *models.AnalyticsDataRetentionResponse
}

// IsSuccess returns true when this put analytics dataretention settings o k response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put analytics dataretention settings o k response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings o k response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put analytics dataretention settings o k response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings o k response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutAnalyticsDataretentionSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsOK  %+v", 200, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsOK  %+v", 200, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsOK) GetPayload() *models.AnalyticsDataRetentionResponse {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalyticsDataRetentionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsBadRequest creates a PutAnalyticsDataretentionSettingsBadRequest with default headers values
func NewPutAnalyticsDataretentionSettingsBadRequest() *PutAnalyticsDataretentionSettingsBadRequest {
	return &PutAnalyticsDataretentionSettingsBadRequest{}
}

/*
PutAnalyticsDataretentionSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutAnalyticsDataretentionSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings bad request response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings bad request response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings bad request response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings bad request response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings bad request response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutAnalyticsDataretentionSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsUnauthorized creates a PutAnalyticsDataretentionSettingsUnauthorized with default headers values
func NewPutAnalyticsDataretentionSettingsUnauthorized() *PutAnalyticsDataretentionSettingsUnauthorized {
	return &PutAnalyticsDataretentionSettingsUnauthorized{}
}

/*
PutAnalyticsDataretentionSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutAnalyticsDataretentionSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings unauthorized response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings unauthorized response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings unauthorized response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings unauthorized response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings unauthorized response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutAnalyticsDataretentionSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsForbidden creates a PutAnalyticsDataretentionSettingsForbidden with default headers values
func NewPutAnalyticsDataretentionSettingsForbidden() *PutAnalyticsDataretentionSettingsForbidden {
	return &PutAnalyticsDataretentionSettingsForbidden{}
}

/*
PutAnalyticsDataretentionSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutAnalyticsDataretentionSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings forbidden response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings forbidden response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings forbidden response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings forbidden response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings forbidden response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutAnalyticsDataretentionSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsNotFound creates a PutAnalyticsDataretentionSettingsNotFound with default headers values
func NewPutAnalyticsDataretentionSettingsNotFound() *PutAnalyticsDataretentionSettingsNotFound {
	return &PutAnalyticsDataretentionSettingsNotFound{}
}

/*
PutAnalyticsDataretentionSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutAnalyticsDataretentionSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings not found response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings not found response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings not found response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings not found response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings not found response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutAnalyticsDataretentionSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsRequestTimeout creates a PutAnalyticsDataretentionSettingsRequestTimeout with default headers values
func NewPutAnalyticsDataretentionSettingsRequestTimeout() *PutAnalyticsDataretentionSettingsRequestTimeout {
	return &PutAnalyticsDataretentionSettingsRequestTimeout{}
}

/*
PutAnalyticsDataretentionSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutAnalyticsDataretentionSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings request timeout response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings request timeout response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings request timeout response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings request timeout response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings request timeout response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutAnalyticsDataretentionSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsConflict creates a PutAnalyticsDataretentionSettingsConflict with default headers values
func NewPutAnalyticsDataretentionSettingsConflict() *PutAnalyticsDataretentionSettingsConflict {
	return &PutAnalyticsDataretentionSettingsConflict{}
}

/*
PutAnalyticsDataretentionSettingsConflict describes a response with status code 409, with default header values.

Conflict
*/
type PutAnalyticsDataretentionSettingsConflict struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings conflict response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings conflict response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings conflict response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings conflict response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings conflict response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutAnalyticsDataretentionSettingsConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsConflict  %+v", 409, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsConflict  %+v", 409, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsRequestEntityTooLarge creates a PutAnalyticsDataretentionSettingsRequestEntityTooLarge with default headers values
func NewPutAnalyticsDataretentionSettingsRequestEntityTooLarge() *PutAnalyticsDataretentionSettingsRequestEntityTooLarge {
	return &PutAnalyticsDataretentionSettingsRequestEntityTooLarge{}
}

/*
PutAnalyticsDataretentionSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutAnalyticsDataretentionSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings request entity too large response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings request entity too large response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings request entity too large response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings request entity too large response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings request entity too large response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsUnsupportedMediaType creates a PutAnalyticsDataretentionSettingsUnsupportedMediaType with default headers values
func NewPutAnalyticsDataretentionSettingsUnsupportedMediaType() *PutAnalyticsDataretentionSettingsUnsupportedMediaType {
	return &PutAnalyticsDataretentionSettingsUnsupportedMediaType{}
}

/*
PutAnalyticsDataretentionSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutAnalyticsDataretentionSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings unsupported media type response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings unsupported media type response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings unsupported media type response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings unsupported media type response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings unsupported media type response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsTooManyRequests creates a PutAnalyticsDataretentionSettingsTooManyRequests with default headers values
func NewPutAnalyticsDataretentionSettingsTooManyRequests() *PutAnalyticsDataretentionSettingsTooManyRequests {
	return &PutAnalyticsDataretentionSettingsTooManyRequests{}
}

/*
PutAnalyticsDataretentionSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutAnalyticsDataretentionSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings too many requests response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings too many requests response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings too many requests response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put analytics dataretention settings too many requests response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put analytics dataretention settings too many requests response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutAnalyticsDataretentionSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsInternalServerError creates a PutAnalyticsDataretentionSettingsInternalServerError with default headers values
func NewPutAnalyticsDataretentionSettingsInternalServerError() *PutAnalyticsDataretentionSettingsInternalServerError {
	return &PutAnalyticsDataretentionSettingsInternalServerError{}
}

/*
PutAnalyticsDataretentionSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutAnalyticsDataretentionSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings internal server error response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings internal server error response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings internal server error response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put analytics dataretention settings internal server error response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put analytics dataretention settings internal server error response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutAnalyticsDataretentionSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsServiceUnavailable creates a PutAnalyticsDataretentionSettingsServiceUnavailable with default headers values
func NewPutAnalyticsDataretentionSettingsServiceUnavailable() *PutAnalyticsDataretentionSettingsServiceUnavailable {
	return &PutAnalyticsDataretentionSettingsServiceUnavailable{}
}

/*
PutAnalyticsDataretentionSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutAnalyticsDataretentionSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings service unavailable response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings service unavailable response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings service unavailable response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put analytics dataretention settings service unavailable response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put analytics dataretention settings service unavailable response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutAnalyticsDataretentionSettingsGatewayTimeout creates a PutAnalyticsDataretentionSettingsGatewayTimeout with default headers values
func NewPutAnalyticsDataretentionSettingsGatewayTimeout() *PutAnalyticsDataretentionSettingsGatewayTimeout {
	return &PutAnalyticsDataretentionSettingsGatewayTimeout{}
}

/*
PutAnalyticsDataretentionSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutAnalyticsDataretentionSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put analytics dataretention settings gateway timeout response has a 2xx status code
func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put analytics dataretention settings gateway timeout response has a 3xx status code
func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put analytics dataretention settings gateway timeout response has a 4xx status code
func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put analytics dataretention settings gateway timeout response has a 5xx status code
func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put analytics dataretention settings gateway timeout response a status code equal to that given
func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/analytics/dataretention/settings][%d] putAnalyticsDataretentionSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutAnalyticsDataretentionSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
