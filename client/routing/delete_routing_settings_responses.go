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

/*DeleteRoutingSettingsNoContent handles this case with default header values.

Operation was successful.
*/
type DeleteRoutingSettingsNoContent struct {
}

func (o *DeleteRoutingSettingsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/routing/settings][%d] deleteRoutingSettingsNoContent ", 204)
}

func (o *DeleteRoutingSettingsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRoutingSettingsBadRequest creates a DeleteRoutingSettingsBadRequest with default headers values
func NewDeleteRoutingSettingsBadRequest() *DeleteRoutingSettingsBadRequest {
	return &DeleteRoutingSettingsBadRequest{}
}

/*DeleteRoutingSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRoutingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsBadRequest) Error() string {
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

/*DeleteRoutingSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRoutingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsUnauthorized) Error() string {
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

/*DeleteRoutingSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRoutingSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsForbidden) Error() string {
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

/*DeleteRoutingSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRoutingSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsNotFound) Error() string {
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

// NewDeleteRoutingSettingsRequestEntityTooLarge creates a DeleteRoutingSettingsRequestEntityTooLarge with default headers values
func NewDeleteRoutingSettingsRequestEntityTooLarge() *DeleteRoutingSettingsRequestEntityTooLarge {
	return &DeleteRoutingSettingsRequestEntityTooLarge{}
}

/*DeleteRoutingSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRoutingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsRequestEntityTooLarge) Error() string {
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

/*DeleteRoutingSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRoutingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsUnsupportedMediaType) Error() string {
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

/*DeleteRoutingSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRoutingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsTooManyRequests) Error() string {
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

/*DeleteRoutingSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRoutingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsInternalServerError) Error() string {
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

/*DeleteRoutingSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRoutingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsServiceUnavailable) Error() string {
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

/*DeleteRoutingSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRoutingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRoutingSettingsGatewayTimeout) Error() string {
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
