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

// GetMobiledevicesReader is a Reader for the GetMobiledevices structure.
type GetMobiledevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMobiledevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMobiledevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMobiledevicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetMobiledevicesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMobiledevicesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMobiledevicesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetMobiledevicesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetMobiledevicesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetMobiledevicesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMobiledevicesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMobiledevicesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetMobiledevicesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetMobiledevicesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMobiledevicesOK creates a GetMobiledevicesOK with default headers values
func NewGetMobiledevicesOK() *GetMobiledevicesOK {
	return &GetMobiledevicesOK{}
}

/*GetMobiledevicesOK handles this case with default header values.

successful operation
*/
type GetMobiledevicesOK struct {
	Payload *models.DirectoryUserDevicesListing
}

func (o *GetMobiledevicesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesOK  %+v", 200, o.Payload)
}

func (o *GetMobiledevicesOK) GetPayload() *models.DirectoryUserDevicesListing {
	return o.Payload
}

func (o *GetMobiledevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryUserDevicesListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesBadRequest creates a GetMobiledevicesBadRequest with default headers values
func NewGetMobiledevicesBadRequest() *GetMobiledevicesBadRequest {
	return &GetMobiledevicesBadRequest{}
}

/*GetMobiledevicesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetMobiledevicesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesBadRequest  %+v", 400, o.Payload)
}

func (o *GetMobiledevicesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesUnauthorized creates a GetMobiledevicesUnauthorized with default headers values
func NewGetMobiledevicesUnauthorized() *GetMobiledevicesUnauthorized {
	return &GetMobiledevicesUnauthorized{}
}

/*GetMobiledevicesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetMobiledevicesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetMobiledevicesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesForbidden creates a GetMobiledevicesForbidden with default headers values
func NewGetMobiledevicesForbidden() *GetMobiledevicesForbidden {
	return &GetMobiledevicesForbidden{}
}

/*GetMobiledevicesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetMobiledevicesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesForbidden  %+v", 403, o.Payload)
}

func (o *GetMobiledevicesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesNotFound creates a GetMobiledevicesNotFound with default headers values
func NewGetMobiledevicesNotFound() *GetMobiledevicesNotFound {
	return &GetMobiledevicesNotFound{}
}

/*GetMobiledevicesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetMobiledevicesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesNotFound  %+v", 404, o.Payload)
}

func (o *GetMobiledevicesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesRequestTimeout creates a GetMobiledevicesRequestTimeout with default headers values
func NewGetMobiledevicesRequestTimeout() *GetMobiledevicesRequestTimeout {
	return &GetMobiledevicesRequestTimeout{}
}

/*GetMobiledevicesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetMobiledevicesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetMobiledevicesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesRequestEntityTooLarge creates a GetMobiledevicesRequestEntityTooLarge with default headers values
func NewGetMobiledevicesRequestEntityTooLarge() *GetMobiledevicesRequestEntityTooLarge {
	return &GetMobiledevicesRequestEntityTooLarge{}
}

/*GetMobiledevicesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetMobiledevicesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetMobiledevicesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesUnsupportedMediaType creates a GetMobiledevicesUnsupportedMediaType with default headers values
func NewGetMobiledevicesUnsupportedMediaType() *GetMobiledevicesUnsupportedMediaType {
	return &GetMobiledevicesUnsupportedMediaType{}
}

/*GetMobiledevicesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetMobiledevicesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetMobiledevicesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesTooManyRequests creates a GetMobiledevicesTooManyRequests with default headers values
func NewGetMobiledevicesTooManyRequests() *GetMobiledevicesTooManyRequests {
	return &GetMobiledevicesTooManyRequests{}
}

/*GetMobiledevicesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetMobiledevicesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMobiledevicesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesInternalServerError creates a GetMobiledevicesInternalServerError with default headers values
func NewGetMobiledevicesInternalServerError() *GetMobiledevicesInternalServerError {
	return &GetMobiledevicesInternalServerError{}
}

/*GetMobiledevicesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetMobiledevicesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetMobiledevicesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesServiceUnavailable creates a GetMobiledevicesServiceUnavailable with default headers values
func NewGetMobiledevicesServiceUnavailable() *GetMobiledevicesServiceUnavailable {
	return &GetMobiledevicesServiceUnavailable{}
}

/*GetMobiledevicesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetMobiledevicesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetMobiledevicesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMobiledevicesGatewayTimeout creates a GetMobiledevicesGatewayTimeout with default headers values
func NewGetMobiledevicesGatewayTimeout() *GetMobiledevicesGatewayTimeout {
	return &GetMobiledevicesGatewayTimeout{}
}

/*GetMobiledevicesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetMobiledevicesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetMobiledevicesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/mobiledevices][%d] getMobiledevicesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetMobiledevicesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetMobiledevicesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
