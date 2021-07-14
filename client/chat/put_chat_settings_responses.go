// Code generated by go-swagger; DO NOT EDIT.

package chat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutChatSettingsReader is a Reader for the PutChatSettings structure.
type PutChatSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutChatSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutChatSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutChatSettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutChatSettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutChatSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutChatSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutChatSettingsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutChatSettingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutChatSettingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutChatSettingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutChatSettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutChatSettingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutChatSettingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutChatSettingsOK creates a PutChatSettingsOK with default headers values
func NewPutChatSettingsOK() *PutChatSettingsOK {
	return &PutChatSettingsOK{}
}

/*PutChatSettingsOK handles this case with default header values.

successful operation
*/
type PutChatSettingsOK struct {
	Payload *models.ChatSettings
}

func (o *PutChatSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsOK  %+v", 200, o.Payload)
}

func (o *PutChatSettingsOK) GetPayload() *models.ChatSettings {
	return o.Payload
}

func (o *PutChatSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChatSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsBadRequest creates a PutChatSettingsBadRequest with default headers values
func NewPutChatSettingsBadRequest() *PutChatSettingsBadRequest {
	return &PutChatSettingsBadRequest{}
}

/*PutChatSettingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutChatSettingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *PutChatSettingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsUnauthorized creates a PutChatSettingsUnauthorized with default headers values
func NewPutChatSettingsUnauthorized() *PutChatSettingsUnauthorized {
	return &PutChatSettingsUnauthorized{}
}

/*PutChatSettingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutChatSettingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PutChatSettingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsForbidden creates a PutChatSettingsForbidden with default headers values
func NewPutChatSettingsForbidden() *PutChatSettingsForbidden {
	return &PutChatSettingsForbidden{}
}

/*PutChatSettingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutChatSettingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsForbidden  %+v", 403, o.Payload)
}

func (o *PutChatSettingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsNotFound creates a PutChatSettingsNotFound with default headers values
func NewPutChatSettingsNotFound() *PutChatSettingsNotFound {
	return &PutChatSettingsNotFound{}
}

/*PutChatSettingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutChatSettingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsNotFound  %+v", 404, o.Payload)
}

func (o *PutChatSettingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsRequestTimeout creates a PutChatSettingsRequestTimeout with default headers values
func NewPutChatSettingsRequestTimeout() *PutChatSettingsRequestTimeout {
	return &PutChatSettingsRequestTimeout{}
}

/*PutChatSettingsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutChatSettingsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutChatSettingsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsRequestEntityTooLarge creates a PutChatSettingsRequestEntityTooLarge with default headers values
func NewPutChatSettingsRequestEntityTooLarge() *PutChatSettingsRequestEntityTooLarge {
	return &PutChatSettingsRequestEntityTooLarge{}
}

/*PutChatSettingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutChatSettingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutChatSettingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsUnsupportedMediaType creates a PutChatSettingsUnsupportedMediaType with default headers values
func NewPutChatSettingsUnsupportedMediaType() *PutChatSettingsUnsupportedMediaType {
	return &PutChatSettingsUnsupportedMediaType{}
}

/*PutChatSettingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutChatSettingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutChatSettingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsTooManyRequests creates a PutChatSettingsTooManyRequests with default headers values
func NewPutChatSettingsTooManyRequests() *PutChatSettingsTooManyRequests {
	return &PutChatSettingsTooManyRequests{}
}

/*PutChatSettingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutChatSettingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutChatSettingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsInternalServerError creates a PutChatSettingsInternalServerError with default headers values
func NewPutChatSettingsInternalServerError() *PutChatSettingsInternalServerError {
	return &PutChatSettingsInternalServerError{}
}

/*PutChatSettingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutChatSettingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PutChatSettingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsServiceUnavailable creates a PutChatSettingsServiceUnavailable with default headers values
func NewPutChatSettingsServiceUnavailable() *PutChatSettingsServiceUnavailable {
	return &PutChatSettingsServiceUnavailable{}
}

/*PutChatSettingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutChatSettingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutChatSettingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutChatSettingsGatewayTimeout creates a PutChatSettingsGatewayTimeout with default headers values
func NewPutChatSettingsGatewayTimeout() *PutChatSettingsGatewayTimeout {
	return &PutChatSettingsGatewayTimeout{}
}

/*PutChatSettingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutChatSettingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutChatSettingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/chat/settings][%d] putChatSettingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutChatSettingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutChatSettingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
