// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchJourneyActionmapReader is a Reader for the PatchJourneyActionmap structure.
type PatchJourneyActionmapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchJourneyActionmapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchJourneyActionmapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchJourneyActionmapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchJourneyActionmapUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchJourneyActionmapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchJourneyActionmapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchJourneyActionmapRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchJourneyActionmapConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchJourneyActionmapRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchJourneyActionmapUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchJourneyActionmapTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchJourneyActionmapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchJourneyActionmapServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchJourneyActionmapGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchJourneyActionmapOK creates a PatchJourneyActionmapOK with default headers values
func NewPatchJourneyActionmapOK() *PatchJourneyActionmapOK {
	return &PatchJourneyActionmapOK{}
}

/*PatchJourneyActionmapOK handles this case with default header values.

successful operation
*/
type PatchJourneyActionmapOK struct {
	Payload *models.ActionMap
}

func (o *PatchJourneyActionmapOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapOK  %+v", 200, o.Payload)
}

func (o *PatchJourneyActionmapOK) GetPayload() *models.ActionMap {
	return o.Payload
}

func (o *PatchJourneyActionmapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapBadRequest creates a PatchJourneyActionmapBadRequest with default headers values
func NewPatchJourneyActionmapBadRequest() *PatchJourneyActionmapBadRequest {
	return &PatchJourneyActionmapBadRequest{}
}

/*PatchJourneyActionmapBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchJourneyActionmapBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapBadRequest  %+v", 400, o.Payload)
}

func (o *PatchJourneyActionmapBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapUnauthorized creates a PatchJourneyActionmapUnauthorized with default headers values
func NewPatchJourneyActionmapUnauthorized() *PatchJourneyActionmapUnauthorized {
	return &PatchJourneyActionmapUnauthorized{}
}

/*PatchJourneyActionmapUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchJourneyActionmapUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchJourneyActionmapUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapForbidden creates a PatchJourneyActionmapForbidden with default headers values
func NewPatchJourneyActionmapForbidden() *PatchJourneyActionmapForbidden {
	return &PatchJourneyActionmapForbidden{}
}

/*PatchJourneyActionmapForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchJourneyActionmapForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapForbidden  %+v", 403, o.Payload)
}

func (o *PatchJourneyActionmapForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapNotFound creates a PatchJourneyActionmapNotFound with default headers values
func NewPatchJourneyActionmapNotFound() *PatchJourneyActionmapNotFound {
	return &PatchJourneyActionmapNotFound{}
}

/*PatchJourneyActionmapNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchJourneyActionmapNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapNotFound  %+v", 404, o.Payload)
}

func (o *PatchJourneyActionmapNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapRequestTimeout creates a PatchJourneyActionmapRequestTimeout with default headers values
func NewPatchJourneyActionmapRequestTimeout() *PatchJourneyActionmapRequestTimeout {
	return &PatchJourneyActionmapRequestTimeout{}
}

/*PatchJourneyActionmapRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchJourneyActionmapRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchJourneyActionmapRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapConflict creates a PatchJourneyActionmapConflict with default headers values
func NewPatchJourneyActionmapConflict() *PatchJourneyActionmapConflict {
	return &PatchJourneyActionmapConflict{}
}

/*PatchJourneyActionmapConflict handles this case with default header values.

Conflict
*/
type PatchJourneyActionmapConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapConflict  %+v", 409, o.Payload)
}

func (o *PatchJourneyActionmapConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapRequestEntityTooLarge creates a PatchJourneyActionmapRequestEntityTooLarge with default headers values
func NewPatchJourneyActionmapRequestEntityTooLarge() *PatchJourneyActionmapRequestEntityTooLarge {
	return &PatchJourneyActionmapRequestEntityTooLarge{}
}

/*PatchJourneyActionmapRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchJourneyActionmapRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchJourneyActionmapRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapUnsupportedMediaType creates a PatchJourneyActionmapUnsupportedMediaType with default headers values
func NewPatchJourneyActionmapUnsupportedMediaType() *PatchJourneyActionmapUnsupportedMediaType {
	return &PatchJourneyActionmapUnsupportedMediaType{}
}

/*PatchJourneyActionmapUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchJourneyActionmapUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchJourneyActionmapUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapTooManyRequests creates a PatchJourneyActionmapTooManyRequests with default headers values
func NewPatchJourneyActionmapTooManyRequests() *PatchJourneyActionmapTooManyRequests {
	return &PatchJourneyActionmapTooManyRequests{}
}

/*PatchJourneyActionmapTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchJourneyActionmapTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchJourneyActionmapTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapInternalServerError creates a PatchJourneyActionmapInternalServerError with default headers values
func NewPatchJourneyActionmapInternalServerError() *PatchJourneyActionmapInternalServerError {
	return &PatchJourneyActionmapInternalServerError{}
}

/*PatchJourneyActionmapInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchJourneyActionmapInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchJourneyActionmapInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapServiceUnavailable creates a PatchJourneyActionmapServiceUnavailable with default headers values
func NewPatchJourneyActionmapServiceUnavailable() *PatchJourneyActionmapServiceUnavailable {
	return &PatchJourneyActionmapServiceUnavailable{}
}

/*PatchJourneyActionmapServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchJourneyActionmapServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchJourneyActionmapServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyActionmapGatewayTimeout creates a PatchJourneyActionmapGatewayTimeout with default headers values
func NewPatchJourneyActionmapGatewayTimeout() *PatchJourneyActionmapGatewayTimeout {
	return &PatchJourneyActionmapGatewayTimeout{}
}

/*PatchJourneyActionmapGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchJourneyActionmapGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyActionmapGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/actionmaps/{actionMapId}][%d] patchJourneyActionmapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchJourneyActionmapGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyActionmapGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
