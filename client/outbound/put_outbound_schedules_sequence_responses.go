// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutOutboundSchedulesSequenceReader is a Reader for the PutOutboundSchedulesSequence structure.
type PutOutboundSchedulesSequenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOutboundSchedulesSequenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutOutboundSchedulesSequenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutOutboundSchedulesSequenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutOutboundSchedulesSequenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutOutboundSchedulesSequenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutOutboundSchedulesSequenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutOutboundSchedulesSequenceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutOutboundSchedulesSequenceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutOutboundSchedulesSequenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutOutboundSchedulesSequenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutOutboundSchedulesSequenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutOutboundSchedulesSequenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutOutboundSchedulesSequenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutOutboundSchedulesSequenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutOutboundSchedulesSequenceOK creates a PutOutboundSchedulesSequenceOK with default headers values
func NewPutOutboundSchedulesSequenceOK() *PutOutboundSchedulesSequenceOK {
	return &PutOutboundSchedulesSequenceOK{}
}

/*PutOutboundSchedulesSequenceOK handles this case with default header values.

successful operation
*/
type PutOutboundSchedulesSequenceOK struct {
	Payload *models.SequenceSchedule
}

func (o *PutOutboundSchedulesSequenceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceOK  %+v", 200, o.Payload)
}

func (o *PutOutboundSchedulesSequenceOK) GetPayload() *models.SequenceSchedule {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SequenceSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceBadRequest creates a PutOutboundSchedulesSequenceBadRequest with default headers values
func NewPutOutboundSchedulesSequenceBadRequest() *PutOutboundSchedulesSequenceBadRequest {
	return &PutOutboundSchedulesSequenceBadRequest{}
}

/*PutOutboundSchedulesSequenceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutOutboundSchedulesSequenceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *PutOutboundSchedulesSequenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceUnauthorized creates a PutOutboundSchedulesSequenceUnauthorized with default headers values
func NewPutOutboundSchedulesSequenceUnauthorized() *PutOutboundSchedulesSequenceUnauthorized {
	return &PutOutboundSchedulesSequenceUnauthorized{}
}

/*PutOutboundSchedulesSequenceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutOutboundSchedulesSequenceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutOutboundSchedulesSequenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceForbidden creates a PutOutboundSchedulesSequenceForbidden with default headers values
func NewPutOutboundSchedulesSequenceForbidden() *PutOutboundSchedulesSequenceForbidden {
	return &PutOutboundSchedulesSequenceForbidden{}
}

/*PutOutboundSchedulesSequenceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutOutboundSchedulesSequenceForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceForbidden  %+v", 403, o.Payload)
}

func (o *PutOutboundSchedulesSequenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceNotFound creates a PutOutboundSchedulesSequenceNotFound with default headers values
func NewPutOutboundSchedulesSequenceNotFound() *PutOutboundSchedulesSequenceNotFound {
	return &PutOutboundSchedulesSequenceNotFound{}
}

/*PutOutboundSchedulesSequenceNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutOutboundSchedulesSequenceNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceNotFound  %+v", 404, o.Payload)
}

func (o *PutOutboundSchedulesSequenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceRequestTimeout creates a PutOutboundSchedulesSequenceRequestTimeout with default headers values
func NewPutOutboundSchedulesSequenceRequestTimeout() *PutOutboundSchedulesSequenceRequestTimeout {
	return &PutOutboundSchedulesSequenceRequestTimeout{}
}

/*PutOutboundSchedulesSequenceRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutOutboundSchedulesSequenceRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutOutboundSchedulesSequenceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceConflict creates a PutOutboundSchedulesSequenceConflict with default headers values
func NewPutOutboundSchedulesSequenceConflict() *PutOutboundSchedulesSequenceConflict {
	return &PutOutboundSchedulesSequenceConflict{}
}

/*PutOutboundSchedulesSequenceConflict handles this case with default header values.

Conflict
*/
type PutOutboundSchedulesSequenceConflict struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceConflict  %+v", 409, o.Payload)
}

func (o *PutOutboundSchedulesSequenceConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceRequestEntityTooLarge creates a PutOutboundSchedulesSequenceRequestEntityTooLarge with default headers values
func NewPutOutboundSchedulesSequenceRequestEntityTooLarge() *PutOutboundSchedulesSequenceRequestEntityTooLarge {
	return &PutOutboundSchedulesSequenceRequestEntityTooLarge{}
}

/*PutOutboundSchedulesSequenceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutOutboundSchedulesSequenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutOutboundSchedulesSequenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceUnsupportedMediaType creates a PutOutboundSchedulesSequenceUnsupportedMediaType with default headers values
func NewPutOutboundSchedulesSequenceUnsupportedMediaType() *PutOutboundSchedulesSequenceUnsupportedMediaType {
	return &PutOutboundSchedulesSequenceUnsupportedMediaType{}
}

/*PutOutboundSchedulesSequenceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutOutboundSchedulesSequenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutOutboundSchedulesSequenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceTooManyRequests creates a PutOutboundSchedulesSequenceTooManyRequests with default headers values
func NewPutOutboundSchedulesSequenceTooManyRequests() *PutOutboundSchedulesSequenceTooManyRequests {
	return &PutOutboundSchedulesSequenceTooManyRequests{}
}

/*PutOutboundSchedulesSequenceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutOutboundSchedulesSequenceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutOutboundSchedulesSequenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceInternalServerError creates a PutOutboundSchedulesSequenceInternalServerError with default headers values
func NewPutOutboundSchedulesSequenceInternalServerError() *PutOutboundSchedulesSequenceInternalServerError {
	return &PutOutboundSchedulesSequenceInternalServerError{}
}

/*PutOutboundSchedulesSequenceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutOutboundSchedulesSequenceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutOutboundSchedulesSequenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceServiceUnavailable creates a PutOutboundSchedulesSequenceServiceUnavailable with default headers values
func NewPutOutboundSchedulesSequenceServiceUnavailable() *PutOutboundSchedulesSequenceServiceUnavailable {
	return &PutOutboundSchedulesSequenceServiceUnavailable{}
}

/*PutOutboundSchedulesSequenceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutOutboundSchedulesSequenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutOutboundSchedulesSequenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutOutboundSchedulesSequenceGatewayTimeout creates a PutOutboundSchedulesSequenceGatewayTimeout with default headers values
func NewPutOutboundSchedulesSequenceGatewayTimeout() *PutOutboundSchedulesSequenceGatewayTimeout {
	return &PutOutboundSchedulesSequenceGatewayTimeout{}
}

/*PutOutboundSchedulesSequenceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutOutboundSchedulesSequenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutOutboundSchedulesSequenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/outbound/schedules/sequences/{sequenceId}][%d] putOutboundSchedulesSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutOutboundSchedulesSequenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutOutboundSchedulesSequenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
