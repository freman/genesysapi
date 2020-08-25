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

/*GetRoutingSettingsOK handles this case with default header values.

successful operation
*/
type GetRoutingSettingsOK struct {
	Payload *models.RoutingSettings
}

func (o *GetRoutingSettingsOK) Error() string {
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

/*GetRoutingSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsBadRequest) Error() string {
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

/*GetRoutingSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsUnauthorized) Error() string {
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

/*GetRoutingSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsForbidden) Error() string {
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

/*GetRoutingSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsNotFound) Error() string {
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

// NewGetRoutingSettingsRequestEntityTooLarge creates a GetRoutingSettingsRequestEntityTooLarge with default headers values
func NewGetRoutingSettingsRequestEntityTooLarge() *GetRoutingSettingsRequestEntityTooLarge {
	return &GetRoutingSettingsRequestEntityTooLarge{}
}

/*GetRoutingSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsRequestEntityTooLarge) Error() string {
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

/*GetRoutingSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsUnsupportedMediaType) Error() string {
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

/*GetRoutingSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTooManyRequests) Error() string {
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

/*GetRoutingSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsInternalServerError) Error() string {
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

/*GetRoutingSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsServiceUnavailable) Error() string {
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

/*GetRoutingSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsGatewayTimeout) Error() string {
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