// Code generated by go-swagger; DO NOT EDIT.

package geolocation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchUserGeolocationReader is a Reader for the PatchUserGeolocation structure.
type PatchUserGeolocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUserGeolocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUserGeolocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchUserGeolocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchUserGeolocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUserGeolocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUserGeolocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPatchUserGeolocationMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchUserGeolocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchUserGeolocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchUserGeolocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUserGeolocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchUserGeolocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchUserGeolocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUserGeolocationOK creates a PatchUserGeolocationOK with default headers values
func NewPatchUserGeolocationOK() *PatchUserGeolocationOK {
	return &PatchUserGeolocationOK{}
}

/*PatchUserGeolocationOK handles this case with default header values.

successful operation
*/
type PatchUserGeolocationOK struct {
	Payload *models.Geolocation
}

func (o *PatchUserGeolocationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationOK  %+v", 200, o.Payload)
}

func (o *PatchUserGeolocationOK) GetPayload() *models.Geolocation {
	return o.Payload
}

func (o *PatchUserGeolocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Geolocation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationBadRequest creates a PatchUserGeolocationBadRequest with default headers values
func NewPatchUserGeolocationBadRequest() *PatchUserGeolocationBadRequest {
	return &PatchUserGeolocationBadRequest{}
}

/*PatchUserGeolocationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchUserGeolocationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchUserGeolocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationUnauthorized creates a PatchUserGeolocationUnauthorized with default headers values
func NewPatchUserGeolocationUnauthorized() *PatchUserGeolocationUnauthorized {
	return &PatchUserGeolocationUnauthorized{}
}

/*PatchUserGeolocationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchUserGeolocationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchUserGeolocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationForbidden creates a PatchUserGeolocationForbidden with default headers values
func NewPatchUserGeolocationForbidden() *PatchUserGeolocationForbidden {
	return &PatchUserGeolocationForbidden{}
}

/*PatchUserGeolocationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchUserGeolocationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationForbidden  %+v", 403, o.Payload)
}

func (o *PatchUserGeolocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationNotFound creates a PatchUserGeolocationNotFound with default headers values
func NewPatchUserGeolocationNotFound() *PatchUserGeolocationNotFound {
	return &PatchUserGeolocationNotFound{}
}

/*PatchUserGeolocationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchUserGeolocationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationNotFound  %+v", 404, o.Payload)
}

func (o *PatchUserGeolocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationMethodNotAllowed creates a PatchUserGeolocationMethodNotAllowed with default headers values
func NewPatchUserGeolocationMethodNotAllowed() *PatchUserGeolocationMethodNotAllowed {
	return &PatchUserGeolocationMethodNotAllowed{}
}

/*PatchUserGeolocationMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PatchUserGeolocationMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PatchUserGeolocationMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationRequestEntityTooLarge creates a PatchUserGeolocationRequestEntityTooLarge with default headers values
func NewPatchUserGeolocationRequestEntityTooLarge() *PatchUserGeolocationRequestEntityTooLarge {
	return &PatchUserGeolocationRequestEntityTooLarge{}
}

/*PatchUserGeolocationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchUserGeolocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchUserGeolocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationUnsupportedMediaType creates a PatchUserGeolocationUnsupportedMediaType with default headers values
func NewPatchUserGeolocationUnsupportedMediaType() *PatchUserGeolocationUnsupportedMediaType {
	return &PatchUserGeolocationUnsupportedMediaType{}
}

/*PatchUserGeolocationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchUserGeolocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchUserGeolocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationTooManyRequests creates a PatchUserGeolocationTooManyRequests with default headers values
func NewPatchUserGeolocationTooManyRequests() *PatchUserGeolocationTooManyRequests {
	return &PatchUserGeolocationTooManyRequests{}
}

/*PatchUserGeolocationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchUserGeolocationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchUserGeolocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationInternalServerError creates a PatchUserGeolocationInternalServerError with default headers values
func NewPatchUserGeolocationInternalServerError() *PatchUserGeolocationInternalServerError {
	return &PatchUserGeolocationInternalServerError{}
}

/*PatchUserGeolocationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchUserGeolocationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchUserGeolocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationServiceUnavailable creates a PatchUserGeolocationServiceUnavailable with default headers values
func NewPatchUserGeolocationServiceUnavailable() *PatchUserGeolocationServiceUnavailable {
	return &PatchUserGeolocationServiceUnavailable{}
}

/*PatchUserGeolocationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchUserGeolocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchUserGeolocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUserGeolocationGatewayTimeout creates a PatchUserGeolocationGatewayTimeout with default headers values
func NewPatchUserGeolocationGatewayTimeout() *PatchUserGeolocationGatewayTimeout {
	return &PatchUserGeolocationGatewayTimeout{}
}

/*PatchUserGeolocationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchUserGeolocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchUserGeolocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/users/{userId}/geolocations/{clientId}][%d] patchUserGeolocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchUserGeolocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchUserGeolocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
