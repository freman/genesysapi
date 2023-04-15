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

// DeleteMobiledeviceReader is a Reader for the DeleteMobiledevice structure.
type DeleteMobiledeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMobiledeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteMobiledeviceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMobiledeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteMobiledeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteMobiledeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteMobiledeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteMobiledeviceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteMobiledeviceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteMobiledeviceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteMobiledeviceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteMobiledeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteMobiledeviceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteMobiledeviceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteMobiledeviceNoContent creates a DeleteMobiledeviceNoContent with default headers values
func NewDeleteMobiledeviceNoContent() *DeleteMobiledeviceNoContent {
	return &DeleteMobiledeviceNoContent{}
}

/*
DeleteMobiledeviceNoContent describes a response with status code 204, with default header values.

Operation was successful.
*/
type DeleteMobiledeviceNoContent struct {
}

// IsSuccess returns true when this delete mobiledevice no content response has a 2xx status code
func (o *DeleteMobiledeviceNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete mobiledevice no content response has a 3xx status code
func (o *DeleteMobiledeviceNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice no content response has a 4xx status code
func (o *DeleteMobiledeviceNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobiledevice no content response has a 5xx status code
func (o *DeleteMobiledeviceNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice no content response a status code equal to that given
func (o *DeleteMobiledeviceNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteMobiledeviceNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceNoContent ", 204)
}

func (o *DeleteMobiledeviceNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceNoContent ", 204)
}

func (o *DeleteMobiledeviceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMobiledeviceBadRequest creates a DeleteMobiledeviceBadRequest with default headers values
func NewDeleteMobiledeviceBadRequest() *DeleteMobiledeviceBadRequest {
	return &DeleteMobiledeviceBadRequest{}
}

/*
DeleteMobiledeviceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteMobiledeviceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice bad request response has a 2xx status code
func (o *DeleteMobiledeviceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice bad request response has a 3xx status code
func (o *DeleteMobiledeviceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice bad request response has a 4xx status code
func (o *DeleteMobiledeviceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice bad request response has a 5xx status code
func (o *DeleteMobiledeviceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice bad request response a status code equal to that given
func (o *DeleteMobiledeviceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteMobiledeviceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMobiledeviceBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteMobiledeviceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceUnauthorized creates a DeleteMobiledeviceUnauthorized with default headers values
func NewDeleteMobiledeviceUnauthorized() *DeleteMobiledeviceUnauthorized {
	return &DeleteMobiledeviceUnauthorized{}
}

/*
DeleteMobiledeviceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteMobiledeviceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice unauthorized response has a 2xx status code
func (o *DeleteMobiledeviceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice unauthorized response has a 3xx status code
func (o *DeleteMobiledeviceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice unauthorized response has a 4xx status code
func (o *DeleteMobiledeviceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice unauthorized response has a 5xx status code
func (o *DeleteMobiledeviceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice unauthorized response a status code equal to that given
func (o *DeleteMobiledeviceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteMobiledeviceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMobiledeviceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteMobiledeviceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceForbidden creates a DeleteMobiledeviceForbidden with default headers values
func NewDeleteMobiledeviceForbidden() *DeleteMobiledeviceForbidden {
	return &DeleteMobiledeviceForbidden{}
}

/*
DeleteMobiledeviceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteMobiledeviceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice forbidden response has a 2xx status code
func (o *DeleteMobiledeviceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice forbidden response has a 3xx status code
func (o *DeleteMobiledeviceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice forbidden response has a 4xx status code
func (o *DeleteMobiledeviceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice forbidden response has a 5xx status code
func (o *DeleteMobiledeviceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice forbidden response a status code equal to that given
func (o *DeleteMobiledeviceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteMobiledeviceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMobiledeviceForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteMobiledeviceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceNotFound creates a DeleteMobiledeviceNotFound with default headers values
func NewDeleteMobiledeviceNotFound() *DeleteMobiledeviceNotFound {
	return &DeleteMobiledeviceNotFound{}
}

/*
DeleteMobiledeviceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteMobiledeviceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice not found response has a 2xx status code
func (o *DeleteMobiledeviceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice not found response has a 3xx status code
func (o *DeleteMobiledeviceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice not found response has a 4xx status code
func (o *DeleteMobiledeviceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice not found response has a 5xx status code
func (o *DeleteMobiledeviceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice not found response a status code equal to that given
func (o *DeleteMobiledeviceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteMobiledeviceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMobiledeviceNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteMobiledeviceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceRequestTimeout creates a DeleteMobiledeviceRequestTimeout with default headers values
func NewDeleteMobiledeviceRequestTimeout() *DeleteMobiledeviceRequestTimeout {
	return &DeleteMobiledeviceRequestTimeout{}
}

/*
DeleteMobiledeviceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteMobiledeviceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice request timeout response has a 2xx status code
func (o *DeleteMobiledeviceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice request timeout response has a 3xx status code
func (o *DeleteMobiledeviceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice request timeout response has a 4xx status code
func (o *DeleteMobiledeviceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice request timeout response has a 5xx status code
func (o *DeleteMobiledeviceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice request timeout response a status code equal to that given
func (o *DeleteMobiledeviceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteMobiledeviceRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteMobiledeviceRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteMobiledeviceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceRequestEntityTooLarge creates a DeleteMobiledeviceRequestEntityTooLarge with default headers values
func NewDeleteMobiledeviceRequestEntityTooLarge() *DeleteMobiledeviceRequestEntityTooLarge {
	return &DeleteMobiledeviceRequestEntityTooLarge{}
}

/*
DeleteMobiledeviceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type DeleteMobiledeviceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice request entity too large response has a 2xx status code
func (o *DeleteMobiledeviceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice request entity too large response has a 3xx status code
func (o *DeleteMobiledeviceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice request entity too large response has a 4xx status code
func (o *DeleteMobiledeviceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice request entity too large response has a 5xx status code
func (o *DeleteMobiledeviceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice request entity too large response a status code equal to that given
func (o *DeleteMobiledeviceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteMobiledeviceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteMobiledeviceRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteMobiledeviceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceUnsupportedMediaType creates a DeleteMobiledeviceUnsupportedMediaType with default headers values
func NewDeleteMobiledeviceUnsupportedMediaType() *DeleteMobiledeviceUnsupportedMediaType {
	return &DeleteMobiledeviceUnsupportedMediaType{}
}

/*
DeleteMobiledeviceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteMobiledeviceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice unsupported media type response has a 2xx status code
func (o *DeleteMobiledeviceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice unsupported media type response has a 3xx status code
func (o *DeleteMobiledeviceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice unsupported media type response has a 4xx status code
func (o *DeleteMobiledeviceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice unsupported media type response has a 5xx status code
func (o *DeleteMobiledeviceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice unsupported media type response a status code equal to that given
func (o *DeleteMobiledeviceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteMobiledeviceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteMobiledeviceUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteMobiledeviceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceTooManyRequests creates a DeleteMobiledeviceTooManyRequests with default headers values
func NewDeleteMobiledeviceTooManyRequests() *DeleteMobiledeviceTooManyRequests {
	return &DeleteMobiledeviceTooManyRequests{}
}

/*
DeleteMobiledeviceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteMobiledeviceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice too many requests response has a 2xx status code
func (o *DeleteMobiledeviceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice too many requests response has a 3xx status code
func (o *DeleteMobiledeviceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice too many requests response has a 4xx status code
func (o *DeleteMobiledeviceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete mobiledevice too many requests response has a 5xx status code
func (o *DeleteMobiledeviceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete mobiledevice too many requests response a status code equal to that given
func (o *DeleteMobiledeviceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteMobiledeviceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteMobiledeviceTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteMobiledeviceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceInternalServerError creates a DeleteMobiledeviceInternalServerError with default headers values
func NewDeleteMobiledeviceInternalServerError() *DeleteMobiledeviceInternalServerError {
	return &DeleteMobiledeviceInternalServerError{}
}

/*
DeleteMobiledeviceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteMobiledeviceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice internal server error response has a 2xx status code
func (o *DeleteMobiledeviceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice internal server error response has a 3xx status code
func (o *DeleteMobiledeviceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice internal server error response has a 4xx status code
func (o *DeleteMobiledeviceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobiledevice internal server error response has a 5xx status code
func (o *DeleteMobiledeviceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete mobiledevice internal server error response a status code equal to that given
func (o *DeleteMobiledeviceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteMobiledeviceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMobiledeviceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteMobiledeviceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceServiceUnavailable creates a DeleteMobiledeviceServiceUnavailable with default headers values
func NewDeleteMobiledeviceServiceUnavailable() *DeleteMobiledeviceServiceUnavailable {
	return &DeleteMobiledeviceServiceUnavailable{}
}

/*
DeleteMobiledeviceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteMobiledeviceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice service unavailable response has a 2xx status code
func (o *DeleteMobiledeviceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice service unavailable response has a 3xx status code
func (o *DeleteMobiledeviceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice service unavailable response has a 4xx status code
func (o *DeleteMobiledeviceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobiledevice service unavailable response has a 5xx status code
func (o *DeleteMobiledeviceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete mobiledevice service unavailable response a status code equal to that given
func (o *DeleteMobiledeviceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteMobiledeviceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteMobiledeviceServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteMobiledeviceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMobiledeviceGatewayTimeout creates a DeleteMobiledeviceGatewayTimeout with default headers values
func NewDeleteMobiledeviceGatewayTimeout() *DeleteMobiledeviceGatewayTimeout {
	return &DeleteMobiledeviceGatewayTimeout{}
}

/*
DeleteMobiledeviceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteMobiledeviceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete mobiledevice gateway timeout response has a 2xx status code
func (o *DeleteMobiledeviceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete mobiledevice gateway timeout response has a 3xx status code
func (o *DeleteMobiledeviceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete mobiledevice gateway timeout response has a 4xx status code
func (o *DeleteMobiledeviceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete mobiledevice gateway timeout response has a 5xx status code
func (o *DeleteMobiledeviceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete mobiledevice gateway timeout response a status code equal to that given
func (o *DeleteMobiledeviceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteMobiledeviceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteMobiledeviceGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/mobiledevices/{deviceId}][%d] deleteMobiledeviceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteMobiledeviceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteMobiledeviceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
