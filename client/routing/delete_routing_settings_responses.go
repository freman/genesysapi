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

// DeleteRoutingSettingsReader is a Reader for the DeleteRoutingSettings structure.
type DeleteRoutingSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRoutingSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRoutingSettingsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRoutingSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRoutingSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRoutingSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRoutingSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteRoutingSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRoutingSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRoutingSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRoutingSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRoutingSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRoutingSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRoutingSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRoutingSettingsNoContent creates a DeleteRoutingSettingsNoContent with default headers values
func NewDeleteRoutingSettingsNoContent() *DeleteRoutingSettingsNoContent {
	return &DeleteRoutingSettingsNoContent{}
}

/*
DeleteRoutingSettingsNoContent describes a response with status code 204, with default header values.

Operation was successful.
*/
type DeleteRoutingSettingsNoContent struct {
}

// IsSuccess returns true when this delete routing settings no content response has a 2xx status code
func (o *DeleteRoutingSettingsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete routing settings no content response has a 3xx status code
func (o *DeleteRoutingSettingsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings no content response has a 4xx status code
func (o *DeleteRoutingSettingsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing settings no content response has a 5xx status code
func (o *DeleteRoutingSettingsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings no content response a status code equal to that given
func (o *DeleteRoutingSettingsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteRoutingSettingsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsNoContent ", 204)
}

func (o *DeleteRoutingSettingsNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsNoContent ", 204)
}

func (o *DeleteRoutingSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingSettingsBadRequest creates a DeleteRoutingSettingsBadRequest with default headers values
func NewDeleteRoutingSettingsBadRequest() *DeleteRoutingSettingsBadRequest {
	return &DeleteRoutingSettingsBadRequest{}
}

/*
DeleteRoutingSettingsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings bad request response has a 2xx status code
func (o *DeleteRoutingSettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings bad request response has a 3xx status code
func (o *DeleteRoutingSettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings bad request response has a 4xx status code
func (o *DeleteRoutingSettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings bad request response has a 5xx status code
func (o *DeleteRoutingSettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings bad request response a status code equal to that given
func (o *DeleteRoutingSettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteRoutingSettingsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSettingsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRoutingSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsUnauthorized creates a DeleteRoutingSettingsUnauthorized with default headers values
func NewDeleteRoutingSettingsUnauthorized() *DeleteRoutingSettingsUnauthorized {
	return &DeleteRoutingSettingsUnauthorized{}
}

/*
DeleteRoutingSettingsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings unauthorized response has a 2xx status code
func (o *DeleteRoutingSettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings unauthorized response has a 3xx status code
func (o *DeleteRoutingSettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings unauthorized response has a 4xx status code
func (o *DeleteRoutingSettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings unauthorized response has a 5xx status code
func (o *DeleteRoutingSettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings unauthorized response a status code equal to that given
func (o *DeleteRoutingSettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteRoutingSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSettingsUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRoutingSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsForbidden creates a DeleteRoutingSettingsForbidden with default headers values
func NewDeleteRoutingSettingsForbidden() *DeleteRoutingSettingsForbidden {
	return &DeleteRoutingSettingsForbidden{}
}

/*
DeleteRoutingSettingsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingSettingsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings forbidden response has a 2xx status code
func (o *DeleteRoutingSettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings forbidden response has a 3xx status code
func (o *DeleteRoutingSettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings forbidden response has a 4xx status code
func (o *DeleteRoutingSettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings forbidden response has a 5xx status code
func (o *DeleteRoutingSettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings forbidden response a status code equal to that given
func (o *DeleteRoutingSettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteRoutingSettingsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSettingsForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRoutingSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsNotFound creates a DeleteRoutingSettingsNotFound with default headers values
func NewDeleteRoutingSettingsNotFound() *DeleteRoutingSettingsNotFound {
	return &DeleteRoutingSettingsNotFound{}
}

/*
DeleteRoutingSettingsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteRoutingSettingsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings not found response has a 2xx status code
func (o *DeleteRoutingSettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings not found response has a 3xx status code
func (o *DeleteRoutingSettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings not found response has a 4xx status code
func (o *DeleteRoutingSettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings not found response has a 5xx status code
func (o *DeleteRoutingSettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings not found response a status code equal to that given
func (o *DeleteRoutingSettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteRoutingSettingsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSettingsNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRoutingSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsRequestTimeout creates a DeleteRoutingSettingsRequestTimeout with default headers values
func NewDeleteRoutingSettingsRequestTimeout() *DeleteRoutingSettingsRequestTimeout {
	return &DeleteRoutingSettingsRequestTimeout{}
}

/*
DeleteRoutingSettingsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteRoutingSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings request timeout response has a 2xx status code
func (o *DeleteRoutingSettingsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings request timeout response has a 3xx status code
func (o *DeleteRoutingSettingsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings request timeout response has a 4xx status code
func (o *DeleteRoutingSettingsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings request timeout response has a 5xx status code
func (o *DeleteRoutingSettingsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings request timeout response a status code equal to that given
func (o *DeleteRoutingSettingsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteRoutingSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSettingsRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteRoutingSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsRequestEntityTooLarge creates a DeleteRoutingSettingsRequestEntityTooLarge with default headers values
func NewDeleteRoutingSettingsRequestEntityTooLarge() *DeleteRoutingSettingsRequestEntityTooLarge {
	return &DeleteRoutingSettingsRequestEntityTooLarge{}
}

/*
DeleteRoutingSettingsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteRoutingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings request entity too large response has a 2xx status code
func (o *DeleteRoutingSettingsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings request entity too large response has a 3xx status code
func (o *DeleteRoutingSettingsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings request entity too large response has a 4xx status code
func (o *DeleteRoutingSettingsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings request entity too large response has a 5xx status code
func (o *DeleteRoutingSettingsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings request entity too large response a status code equal to that given
func (o *DeleteRoutingSettingsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteRoutingSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSettingsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRoutingSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsUnsupportedMediaType creates a DeleteRoutingSettingsUnsupportedMediaType with default headers values
func NewDeleteRoutingSettingsUnsupportedMediaType() *DeleteRoutingSettingsUnsupportedMediaType {
	return &DeleteRoutingSettingsUnsupportedMediaType{}
}

/*
DeleteRoutingSettingsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings unsupported media type response has a 2xx status code
func (o *DeleteRoutingSettingsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings unsupported media type response has a 3xx status code
func (o *DeleteRoutingSettingsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings unsupported media type response has a 4xx status code
func (o *DeleteRoutingSettingsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings unsupported media type response has a 5xx status code
func (o *DeleteRoutingSettingsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings unsupported media type response a status code equal to that given
func (o *DeleteRoutingSettingsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteRoutingSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSettingsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRoutingSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsTooManyRequests creates a DeleteRoutingSettingsTooManyRequests with default headers values
func NewDeleteRoutingSettingsTooManyRequests() *DeleteRoutingSettingsTooManyRequests {
	return &DeleteRoutingSettingsTooManyRequests{}
}

/*
DeleteRoutingSettingsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteRoutingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings too many requests response has a 2xx status code
func (o *DeleteRoutingSettingsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings too many requests response has a 3xx status code
func (o *DeleteRoutingSettingsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings too many requests response has a 4xx status code
func (o *DeleteRoutingSettingsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete routing settings too many requests response has a 5xx status code
func (o *DeleteRoutingSettingsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete routing settings too many requests response a status code equal to that given
func (o *DeleteRoutingSettingsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteRoutingSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSettingsTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRoutingSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsInternalServerError creates a DeleteRoutingSettingsInternalServerError with default headers values
func NewDeleteRoutingSettingsInternalServerError() *DeleteRoutingSettingsInternalServerError {
	return &DeleteRoutingSettingsInternalServerError{}
}

/*
DeleteRoutingSettingsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings internal server error response has a 2xx status code
func (o *DeleteRoutingSettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings internal server error response has a 3xx status code
func (o *DeleteRoutingSettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings internal server error response has a 4xx status code
func (o *DeleteRoutingSettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing settings internal server error response has a 5xx status code
func (o *DeleteRoutingSettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing settings internal server error response a status code equal to that given
func (o *DeleteRoutingSettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteRoutingSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSettingsInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRoutingSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsServiceUnavailable creates a DeleteRoutingSettingsServiceUnavailable with default headers values
func NewDeleteRoutingSettingsServiceUnavailable() *DeleteRoutingSettingsServiceUnavailable {
	return &DeleteRoutingSettingsServiceUnavailable{}
}

/*
DeleteRoutingSettingsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings service unavailable response has a 2xx status code
func (o *DeleteRoutingSettingsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings service unavailable response has a 3xx status code
func (o *DeleteRoutingSettingsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings service unavailable response has a 4xx status code
func (o *DeleteRoutingSettingsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing settings service unavailable response has a 5xx status code
func (o *DeleteRoutingSettingsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing settings service unavailable response a status code equal to that given
func (o *DeleteRoutingSettingsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteRoutingSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSettingsServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRoutingSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRoutingSettingsGatewayTimeout creates a DeleteRoutingSettingsGatewayTimeout with default headers values
func NewDeleteRoutingSettingsGatewayTimeout() *DeleteRoutingSettingsGatewayTimeout {
	return &DeleteRoutingSettingsGatewayTimeout{}
}

/*
DeleteRoutingSettingsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteRoutingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete routing settings gateway timeout response has a 2xx status code
func (o *DeleteRoutingSettingsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete routing settings gateway timeout response has a 3xx status code
func (o *DeleteRoutingSettingsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete routing settings gateway timeout response has a 4xx status code
func (o *DeleteRoutingSettingsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete routing settings gateway timeout response has a 5xx status code
func (o *DeleteRoutingSettingsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete routing settings gateway timeout response a status code equal to that given
func (o *DeleteRoutingSettingsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteRoutingSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSettingsGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRoutingSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRoutingSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
