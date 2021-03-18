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

// PatchJourneyOutcomeReader is a Reader for the PatchJourneyOutcome structure.
type PatchJourneyOutcomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchJourneyOutcomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchJourneyOutcomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchJourneyOutcomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchJourneyOutcomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchJourneyOutcomeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchJourneyOutcomeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchJourneyOutcomeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchJourneyOutcomeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchJourneyOutcomeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchJourneyOutcomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchJourneyOutcomeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchJourneyOutcomeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchJourneyOutcomeOK creates a PatchJourneyOutcomeOK with default headers values
func NewPatchJourneyOutcomeOK() *PatchJourneyOutcomeOK {
	return &PatchJourneyOutcomeOK{}
}

/*PatchJourneyOutcomeOK handles this case with default header values.

successful operation
*/
type PatchJourneyOutcomeOK struct {
	Payload *models.Outcome
}

func (o *PatchJourneyOutcomeOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeOK  %+v", 200, o.Payload)
}

func (o *PatchJourneyOutcomeOK) GetPayload() *models.Outcome {
	return o.Payload
}

func (o *PatchJourneyOutcomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Outcome)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeBadRequest creates a PatchJourneyOutcomeBadRequest with default headers values
func NewPatchJourneyOutcomeBadRequest() *PatchJourneyOutcomeBadRequest {
	return &PatchJourneyOutcomeBadRequest{}
}

/*PatchJourneyOutcomeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchJourneyOutcomeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeBadRequest  %+v", 400, o.Payload)
}

func (o *PatchJourneyOutcomeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeUnauthorized creates a PatchJourneyOutcomeUnauthorized with default headers values
func NewPatchJourneyOutcomeUnauthorized() *PatchJourneyOutcomeUnauthorized {
	return &PatchJourneyOutcomeUnauthorized{}
}

/*PatchJourneyOutcomeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchJourneyOutcomeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchJourneyOutcomeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeForbidden creates a PatchJourneyOutcomeForbidden with default headers values
func NewPatchJourneyOutcomeForbidden() *PatchJourneyOutcomeForbidden {
	return &PatchJourneyOutcomeForbidden{}
}

/*PatchJourneyOutcomeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchJourneyOutcomeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeForbidden  %+v", 403, o.Payload)
}

func (o *PatchJourneyOutcomeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeNotFound creates a PatchJourneyOutcomeNotFound with default headers values
func NewPatchJourneyOutcomeNotFound() *PatchJourneyOutcomeNotFound {
	return &PatchJourneyOutcomeNotFound{}
}

/*PatchJourneyOutcomeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchJourneyOutcomeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeNotFound  %+v", 404, o.Payload)
}

func (o *PatchJourneyOutcomeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeRequestEntityTooLarge creates a PatchJourneyOutcomeRequestEntityTooLarge with default headers values
func NewPatchJourneyOutcomeRequestEntityTooLarge() *PatchJourneyOutcomeRequestEntityTooLarge {
	return &PatchJourneyOutcomeRequestEntityTooLarge{}
}

/*PatchJourneyOutcomeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchJourneyOutcomeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchJourneyOutcomeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeUnsupportedMediaType creates a PatchJourneyOutcomeUnsupportedMediaType with default headers values
func NewPatchJourneyOutcomeUnsupportedMediaType() *PatchJourneyOutcomeUnsupportedMediaType {
	return &PatchJourneyOutcomeUnsupportedMediaType{}
}

/*PatchJourneyOutcomeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchJourneyOutcomeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchJourneyOutcomeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeTooManyRequests creates a PatchJourneyOutcomeTooManyRequests with default headers values
func NewPatchJourneyOutcomeTooManyRequests() *PatchJourneyOutcomeTooManyRequests {
	return &PatchJourneyOutcomeTooManyRequests{}
}

/*PatchJourneyOutcomeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PatchJourneyOutcomeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchJourneyOutcomeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeInternalServerError creates a PatchJourneyOutcomeInternalServerError with default headers values
func NewPatchJourneyOutcomeInternalServerError() *PatchJourneyOutcomeInternalServerError {
	return &PatchJourneyOutcomeInternalServerError{}
}

/*PatchJourneyOutcomeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchJourneyOutcomeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchJourneyOutcomeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeServiceUnavailable creates a PatchJourneyOutcomeServiceUnavailable with default headers values
func NewPatchJourneyOutcomeServiceUnavailable() *PatchJourneyOutcomeServiceUnavailable {
	return &PatchJourneyOutcomeServiceUnavailable{}
}

/*PatchJourneyOutcomeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchJourneyOutcomeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchJourneyOutcomeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchJourneyOutcomeGatewayTimeout creates a PatchJourneyOutcomeGatewayTimeout with default headers values
func NewPatchJourneyOutcomeGatewayTimeout() *PatchJourneyOutcomeGatewayTimeout {
	return &PatchJourneyOutcomeGatewayTimeout{}
}

/*PatchJourneyOutcomeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchJourneyOutcomeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchJourneyOutcomeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/journey/outcomes/{outcomeId}][%d] patchJourneyOutcomeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchJourneyOutcomeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchJourneyOutcomeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}