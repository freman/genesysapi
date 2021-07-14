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

// DeleteOutboundSchedulesSequenceReader is a Reader for the DeleteOutboundSchedulesSequence structure.
type DeleteOutboundSchedulesSequenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundSchedulesSequenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundSchedulesSequenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundSchedulesSequenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundSchedulesSequenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundSchedulesSequenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundSchedulesSequenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteOutboundSchedulesSequenceRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundSchedulesSequenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundSchedulesSequenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundSchedulesSequenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundSchedulesSequenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundSchedulesSequenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundSchedulesSequenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundSchedulesSequenceOK creates a DeleteOutboundSchedulesSequenceOK with default headers values
func NewDeleteOutboundSchedulesSequenceOK() *DeleteOutboundSchedulesSequenceOK {
	return &DeleteOutboundSchedulesSequenceOK{}
}

/*DeleteOutboundSchedulesSequenceOK handles this case with default header values.

Operation was successful.
*/
type DeleteOutboundSchedulesSequenceOK struct {
}

func (o *DeleteOutboundSchedulesSequenceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceOK ", 200)
}

func (o *DeleteOutboundSchedulesSequenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundSchedulesSequenceBadRequest creates a DeleteOutboundSchedulesSequenceBadRequest with default headers values
func NewDeleteOutboundSchedulesSequenceBadRequest() *DeleteOutboundSchedulesSequenceBadRequest {
	return &DeleteOutboundSchedulesSequenceBadRequest{}
}

/*DeleteOutboundSchedulesSequenceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundSchedulesSequenceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceUnauthorized creates a DeleteOutboundSchedulesSequenceUnauthorized with default headers values
func NewDeleteOutboundSchedulesSequenceUnauthorized() *DeleteOutboundSchedulesSequenceUnauthorized {
	return &DeleteOutboundSchedulesSequenceUnauthorized{}
}

/*DeleteOutboundSchedulesSequenceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundSchedulesSequenceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceForbidden creates a DeleteOutboundSchedulesSequenceForbidden with default headers values
func NewDeleteOutboundSchedulesSequenceForbidden() *DeleteOutboundSchedulesSequenceForbidden {
	return &DeleteOutboundSchedulesSequenceForbidden{}
}

/*DeleteOutboundSchedulesSequenceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundSchedulesSequenceForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceNotFound creates a DeleteOutboundSchedulesSequenceNotFound with default headers values
func NewDeleteOutboundSchedulesSequenceNotFound() *DeleteOutboundSchedulesSequenceNotFound {
	return &DeleteOutboundSchedulesSequenceNotFound{}
}

/*DeleteOutboundSchedulesSequenceNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundSchedulesSequenceNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceRequestTimeout creates a DeleteOutboundSchedulesSequenceRequestTimeout with default headers values
func NewDeleteOutboundSchedulesSequenceRequestTimeout() *DeleteOutboundSchedulesSequenceRequestTimeout {
	return &DeleteOutboundSchedulesSequenceRequestTimeout{}
}

/*DeleteOutboundSchedulesSequenceRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteOutboundSchedulesSequenceRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceRequestEntityTooLarge creates a DeleteOutboundSchedulesSequenceRequestEntityTooLarge with default headers values
func NewDeleteOutboundSchedulesSequenceRequestEntityTooLarge() *DeleteOutboundSchedulesSequenceRequestEntityTooLarge {
	return &DeleteOutboundSchedulesSequenceRequestEntityTooLarge{}
}

/*DeleteOutboundSchedulesSequenceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundSchedulesSequenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceUnsupportedMediaType creates a DeleteOutboundSchedulesSequenceUnsupportedMediaType with default headers values
func NewDeleteOutboundSchedulesSequenceUnsupportedMediaType() *DeleteOutboundSchedulesSequenceUnsupportedMediaType {
	return &DeleteOutboundSchedulesSequenceUnsupportedMediaType{}
}

/*DeleteOutboundSchedulesSequenceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundSchedulesSequenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceTooManyRequests creates a DeleteOutboundSchedulesSequenceTooManyRequests with default headers values
func NewDeleteOutboundSchedulesSequenceTooManyRequests() *DeleteOutboundSchedulesSequenceTooManyRequests {
	return &DeleteOutboundSchedulesSequenceTooManyRequests{}
}

/*DeleteOutboundSchedulesSequenceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteOutboundSchedulesSequenceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceInternalServerError creates a DeleteOutboundSchedulesSequenceInternalServerError with default headers values
func NewDeleteOutboundSchedulesSequenceInternalServerError() *DeleteOutboundSchedulesSequenceInternalServerError {
	return &DeleteOutboundSchedulesSequenceInternalServerError{}
}

/*DeleteOutboundSchedulesSequenceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundSchedulesSequenceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceServiceUnavailable creates a DeleteOutboundSchedulesSequenceServiceUnavailable with default headers values
func NewDeleteOutboundSchedulesSequenceServiceUnavailable() *DeleteOutboundSchedulesSequenceServiceUnavailable {
	return &DeleteOutboundSchedulesSequenceServiceUnavailable{}
}

/*DeleteOutboundSchedulesSequenceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundSchedulesSequenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSchedulesSequenceGatewayTimeout creates a DeleteOutboundSchedulesSequenceGatewayTimeout with default headers values
func NewDeleteOutboundSchedulesSequenceGatewayTimeout() *DeleteOutboundSchedulesSequenceGatewayTimeout {
	return &DeleteOutboundSchedulesSequenceGatewayTimeout{}
}

/*DeleteOutboundSchedulesSequenceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundSchedulesSequenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSchedulesSequenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/schedules/sequences/{sequenceId}][%d] deleteOutboundSchedulesSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundSchedulesSequenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSchedulesSequenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
