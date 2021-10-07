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

// PatchJourneySegmentReader is a Reader for the PatchJourneySegment structure.
type PatchJourneySegmentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchJourneySegmentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchJourneySegmentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchJourneySegmentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchJourneySegmentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchJourneySegmentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchJourneySegmentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchJourneySegmentRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPatchJourneySegmentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchJourneySegmentRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchJourneySegmentUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchJourneySegmentTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchJourneySegmentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchJourneySegmentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchJourneySegmentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchJourneySegmentOK creates a PatchJourneySegmentOK with default headers values
func NewPatchJourneySegmentOK() *PatchJourneySegmentOK {
	return &PatchJourneySegmentOK{}
}

/*PatchJourneySegmentOK handles this case with default header values.

successful operation
*/
type PatchJourneySegmentOK struct {
	Payload *models.JourneySegment
}

func (o *PatchJourneySegmentOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentOK  %+v", 200, o.Payload)
}

func (o *PatchJourneySegmentOK) GetPayload() *models.JourneySegment {
	return o.Payload
}

func (o *PatchJourneySegmentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JourneySegment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentBadRequest creates a PatchJourneySegmentBadRequest with default headers values
func NewPatchJourneySegmentBadRequest() *PatchJourneySegmentBadRequest {
	return &PatchJourneySegmentBadRequest{}
}

/*PatchJourneySegmentBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchJourneySegmentBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentBadRequest  %+v", 400, o.Payload)
}

func (o *PatchJourneySegmentBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentUnauthorized creates a PatchJourneySegmentUnauthorized with default headers values
func NewPatchJourneySegmentUnauthorized() *PatchJourneySegmentUnauthorized {
	return &PatchJourneySegmentUnauthorized{}
}

/*PatchJourneySegmentUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchJourneySegmentUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchJourneySegmentUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentForbidden creates a PatchJourneySegmentForbidden with default headers values
func NewPatchJourneySegmentForbidden() *PatchJourneySegmentForbidden {
	return &PatchJourneySegmentForbidden{}
}

/*PatchJourneySegmentForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchJourneySegmentForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentForbidden  %+v", 403, o.Payload)
}

func (o *PatchJourneySegmentForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentNotFound creates a PatchJourneySegmentNotFound with default headers values
func NewPatchJourneySegmentNotFound() *PatchJourneySegmentNotFound {
	return &PatchJourneySegmentNotFound{}
}

/*PatchJourneySegmentNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchJourneySegmentNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentNotFound  %+v", 404, o.Payload)
}

func (o *PatchJourneySegmentNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentRequestTimeout creates a PatchJourneySegmentRequestTimeout with default headers values
func NewPatchJourneySegmentRequestTimeout() *PatchJourneySegmentRequestTimeout {
	return &PatchJourneySegmentRequestTimeout{}
}

/*PatchJourneySegmentRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchJourneySegmentRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchJourneySegmentRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentConflict creates a PatchJourneySegmentConflict with default headers values
func NewPatchJourneySegmentConflict() *PatchJourneySegmentConflict {
	return &PatchJourneySegmentConflict{}
}

/*PatchJourneySegmentConflict handles this case with default header values.

Conflict
*/
type PatchJourneySegmentConflict struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentConflict) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentConflict  %+v", 409, o.Payload)
}

func (o *PatchJourneySegmentConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentRequestEntityTooLarge creates a PatchJourneySegmentRequestEntityTooLarge with default headers values
func NewPatchJourneySegmentRequestEntityTooLarge() *PatchJourneySegmentRequestEntityTooLarge {
	return &PatchJourneySegmentRequestEntityTooLarge{}
}

/*PatchJourneySegmentRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchJourneySegmentRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchJourneySegmentRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentUnsupportedMediaType creates a PatchJourneySegmentUnsupportedMediaType with default headers values
func NewPatchJourneySegmentUnsupportedMediaType() *PatchJourneySegmentUnsupportedMediaType {
	return &PatchJourneySegmentUnsupportedMediaType{}
}

/*PatchJourneySegmentUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchJourneySegmentUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchJourneySegmentUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentTooManyRequests creates a PatchJourneySegmentTooManyRequests with default headers values
func NewPatchJourneySegmentTooManyRequests() *PatchJourneySegmentTooManyRequests {
	return &PatchJourneySegmentTooManyRequests{}
}

/*PatchJourneySegmentTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchJourneySegmentTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchJourneySegmentTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentInternalServerError creates a PatchJourneySegmentInternalServerError with default headers values
func NewPatchJourneySegmentInternalServerError() *PatchJourneySegmentInternalServerError {
	return &PatchJourneySegmentInternalServerError{}
}

/*PatchJourneySegmentInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchJourneySegmentInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchJourneySegmentInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentServiceUnavailable creates a PatchJourneySegmentServiceUnavailable with default headers values
func NewPatchJourneySegmentServiceUnavailable() *PatchJourneySegmentServiceUnavailable {
	return &PatchJourneySegmentServiceUnavailable{}
}

/*PatchJourneySegmentServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchJourneySegmentServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchJourneySegmentServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneySegmentGatewayTimeout creates a PatchJourneySegmentGatewayTimeout with default headers values
func NewPatchJourneySegmentGatewayTimeout() *PatchJourneySegmentGatewayTimeout {
	return &PatchJourneySegmentGatewayTimeout{}
}

/*PatchJourneySegmentGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchJourneySegmentGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneySegmentGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/segments/{segmentId}][%d] patchJourneySegmentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchJourneySegmentGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneySegmentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
