// Code generated by go-swagger; DO NOT EDIT.

package mobile_devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostMobiledevicesReader is a Reader for the PostMobiledevices structure.
type PostMobiledevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostMobiledevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostMobiledevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostMobiledevicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostMobiledevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostMobiledevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostMobiledevicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostMobiledevicesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostMobiledevicesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostMobiledevicesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostMobiledevicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostMobiledevicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostMobiledevicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostMobiledevicesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostMobiledevicesOK creates a PostMobiledevicesOK with default headers values
func NewPostMobiledevicesOK() *PostMobiledevicesOK {
	return &PostMobiledevicesOK{}
}

/*PostMobiledevicesOK handles this case with default header values.

successful operation
*/
type PostMobiledevicesOK struct {
	Payload *models.UserDevice
}

func (o *PostMobiledevicesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesOK  %+v", 200, o.Payload)
}

func (o *PostMobiledevicesOK) GetPayload() *models.UserDevice {
	return o.Payload
}

func (o *PostMobiledevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserDevice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesBadRequest creates a PostMobiledevicesBadRequest with default headers values
func NewPostMobiledevicesBadRequest() *PostMobiledevicesBadRequest {
	return &PostMobiledevicesBadRequest{}
}

/*PostMobiledevicesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostMobiledevicesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesBadRequest  %+v", 400, o.Payload)
}

func (o *PostMobiledevicesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesUnauthorized creates a PostMobiledevicesUnauthorized with default headers values
func NewPostMobiledevicesUnauthorized() *PostMobiledevicesUnauthorized {
	return &PostMobiledevicesUnauthorized{}
}

/*PostMobiledevicesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostMobiledevicesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostMobiledevicesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesForbidden creates a PostMobiledevicesForbidden with default headers values
func NewPostMobiledevicesForbidden() *PostMobiledevicesForbidden {
	return &PostMobiledevicesForbidden{}
}

/*PostMobiledevicesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostMobiledevicesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesForbidden  %+v", 403, o.Payload)
}

func (o *PostMobiledevicesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesNotFound creates a PostMobiledevicesNotFound with default headers values
func NewPostMobiledevicesNotFound() *PostMobiledevicesNotFound {
	return &PostMobiledevicesNotFound{}
}

/*PostMobiledevicesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostMobiledevicesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesNotFound  %+v", 404, o.Payload)
}

func (o *PostMobiledevicesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesRequestTimeout creates a PostMobiledevicesRequestTimeout with default headers values
func NewPostMobiledevicesRequestTimeout() *PostMobiledevicesRequestTimeout {
	return &PostMobiledevicesRequestTimeout{}
}

/*PostMobiledevicesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostMobiledevicesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostMobiledevicesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesRequestEntityTooLarge creates a PostMobiledevicesRequestEntityTooLarge with default headers values
func NewPostMobiledevicesRequestEntityTooLarge() *PostMobiledevicesRequestEntityTooLarge {
	return &PostMobiledevicesRequestEntityTooLarge{}
}

/*PostMobiledevicesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostMobiledevicesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostMobiledevicesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesUnsupportedMediaType creates a PostMobiledevicesUnsupportedMediaType with default headers values
func NewPostMobiledevicesUnsupportedMediaType() *PostMobiledevicesUnsupportedMediaType {
	return &PostMobiledevicesUnsupportedMediaType{}
}

/*PostMobiledevicesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostMobiledevicesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostMobiledevicesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesTooManyRequests creates a PostMobiledevicesTooManyRequests with default headers values
func NewPostMobiledevicesTooManyRequests() *PostMobiledevicesTooManyRequests {
	return &PostMobiledevicesTooManyRequests{}
}

/*PostMobiledevicesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostMobiledevicesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostMobiledevicesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesInternalServerError creates a PostMobiledevicesInternalServerError with default headers values
func NewPostMobiledevicesInternalServerError() *PostMobiledevicesInternalServerError {
	return &PostMobiledevicesInternalServerError{}
}

/*PostMobiledevicesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostMobiledevicesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostMobiledevicesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesServiceUnavailable creates a PostMobiledevicesServiceUnavailable with default headers values
func NewPostMobiledevicesServiceUnavailable() *PostMobiledevicesServiceUnavailable {
	return &PostMobiledevicesServiceUnavailable{}
}

/*PostMobiledevicesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostMobiledevicesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostMobiledevicesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostMobiledevicesGatewayTimeout creates a PostMobiledevicesGatewayTimeout with default headers values
func NewPostMobiledevicesGatewayTimeout() *PostMobiledevicesGatewayTimeout {
	return &PostMobiledevicesGatewayTimeout{}
}

/*PostMobiledevicesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostMobiledevicesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostMobiledevicesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/mobiledevices][%d] postMobiledevicesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostMobiledevicesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostMobiledevicesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
