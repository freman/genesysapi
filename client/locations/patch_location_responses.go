// Code generated by go-swagger; DO NOT EDIT.

package locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchLocationReader is a Reader for the PatchLocation structure.
type PatchLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchLocationRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchLocationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchLocationRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchLocationUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchLocationTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchLocationServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchLocationGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchLocationOK creates a PatchLocationOK with default headers values
func NewPatchLocationOK() *PatchLocationOK {
	return &PatchLocationOK{}
}

/*PatchLocationOK handles this case with default header values.

successful operation
*/
type PatchLocationOK struct {
	Payload *models.LocationDefinition
}

func (o *PatchLocationOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationOK  %+v", 200, o.Payload)
}

func (o *PatchLocationOK) GetPayload() *models.LocationDefinition {
	return o.Payload
}

func (o *PatchLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LocationDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationBadRequest creates a PatchLocationBadRequest with default headers values
func NewPatchLocationBadRequest() *PatchLocationBadRequest {
	return &PatchLocationBadRequest{}
}

/*PatchLocationBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchLocationBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLocationBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationUnauthorized creates a PatchLocationUnauthorized with default headers values
func NewPatchLocationUnauthorized() *PatchLocationUnauthorized {
	return &PatchLocationUnauthorized{}
}

/*PatchLocationUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchLocationUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchLocationUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationForbidden creates a PatchLocationForbidden with default headers values
func NewPatchLocationForbidden() *PatchLocationForbidden {
	return &PatchLocationForbidden{}
}

/*PatchLocationForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchLocationForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationForbidden  %+v", 403, o.Payload)
}

func (o *PatchLocationForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationNotFound creates a PatchLocationNotFound with default headers values
func NewPatchLocationNotFound() *PatchLocationNotFound {
	return &PatchLocationNotFound{}
}

/*PatchLocationNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchLocationNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationNotFound  %+v", 404, o.Payload)
}

func (o *PatchLocationNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationRequestTimeout creates a PatchLocationRequestTimeout with default headers values
func NewPatchLocationRequestTimeout() *PatchLocationRequestTimeout {
	return &PatchLocationRequestTimeout{}
}

/*PatchLocationRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchLocationRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchLocationRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationConflict creates a PatchLocationConflict with default headers values
func NewPatchLocationConflict() *PatchLocationConflict {
	return &PatchLocationConflict{}
}

/*PatchLocationConflict handles this case with default header values.

Conflict
*/
type PatchLocationConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationConflict  %+v", 409, o.Payload)
}

func (o *PatchLocationConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationRequestEntityTooLarge creates a PatchLocationRequestEntityTooLarge with default headers values
func NewPatchLocationRequestEntityTooLarge() *PatchLocationRequestEntityTooLarge {
	return &PatchLocationRequestEntityTooLarge{}
}

/*PatchLocationRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchLocationRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchLocationRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationUnsupportedMediaType creates a PatchLocationUnsupportedMediaType with default headers values
func NewPatchLocationUnsupportedMediaType() *PatchLocationUnsupportedMediaType {
	return &PatchLocationUnsupportedMediaType{}
}

/*PatchLocationUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchLocationUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchLocationUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationTooManyRequests creates a PatchLocationTooManyRequests with default headers values
func NewPatchLocationTooManyRequests() *PatchLocationTooManyRequests {
	return &PatchLocationTooManyRequests{}
}

/*PatchLocationTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchLocationTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchLocationTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationInternalServerError creates a PatchLocationInternalServerError with default headers values
func NewPatchLocationInternalServerError() *PatchLocationInternalServerError {
	return &PatchLocationInternalServerError{}
}

/*PatchLocationInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchLocationInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchLocationInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationServiceUnavailable creates a PatchLocationServiceUnavailable with default headers values
func NewPatchLocationServiceUnavailable() *PatchLocationServiceUnavailable {
	return &PatchLocationServiceUnavailable{}
}

/*PatchLocationServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchLocationServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchLocationServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLocationGatewayTimeout creates a PatchLocationGatewayTimeout with default headers values
func NewPatchLocationGatewayTimeout() *PatchLocationGatewayTimeout {
	return &PatchLocationGatewayTimeout{}
}

/*PatchLocationGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchLocationGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchLocationGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/locations/{locationId}][%d] patchLocationGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchLocationGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchLocationGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
