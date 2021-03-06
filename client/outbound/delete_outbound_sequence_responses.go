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

// DeleteOutboundSequenceReader is a Reader for the DeleteOutboundSequence structure.
type DeleteOutboundSequenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOutboundSequenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOutboundSequenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteOutboundSequenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteOutboundSequenceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOutboundSequenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOutboundSequenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteOutboundSequenceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteOutboundSequenceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteOutboundSequenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOutboundSequenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteOutboundSequenceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteOutboundSequenceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteOutboundSequenceOK creates a DeleteOutboundSequenceOK with default headers values
func NewDeleteOutboundSequenceOK() *DeleteOutboundSequenceOK {
	return &DeleteOutboundSequenceOK{}
}

/*DeleteOutboundSequenceOK handles this case with default header values.

Operation was successful.
*/
type DeleteOutboundSequenceOK struct {
}

func (o *DeleteOutboundSequenceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceOK ", 200)
}

func (o *DeleteOutboundSequenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOutboundSequenceBadRequest creates a DeleteOutboundSequenceBadRequest with default headers values
func NewDeleteOutboundSequenceBadRequest() *DeleteOutboundSequenceBadRequest {
	return &DeleteOutboundSequenceBadRequest{}
}

/*DeleteOutboundSequenceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteOutboundSequenceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteOutboundSequenceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceUnauthorized creates a DeleteOutboundSequenceUnauthorized with default headers values
func NewDeleteOutboundSequenceUnauthorized() *DeleteOutboundSequenceUnauthorized {
	return &DeleteOutboundSequenceUnauthorized{}
}

/*DeleteOutboundSequenceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteOutboundSequenceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteOutboundSequenceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceForbidden creates a DeleteOutboundSequenceForbidden with default headers values
func NewDeleteOutboundSequenceForbidden() *DeleteOutboundSequenceForbidden {
	return &DeleteOutboundSequenceForbidden{}
}

/*DeleteOutboundSequenceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteOutboundSequenceForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOutboundSequenceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceNotFound creates a DeleteOutboundSequenceNotFound with default headers values
func NewDeleteOutboundSequenceNotFound() *DeleteOutboundSequenceNotFound {
	return &DeleteOutboundSequenceNotFound{}
}

/*DeleteOutboundSequenceNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteOutboundSequenceNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOutboundSequenceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceRequestEntityTooLarge creates a DeleteOutboundSequenceRequestEntityTooLarge with default headers values
func NewDeleteOutboundSequenceRequestEntityTooLarge() *DeleteOutboundSequenceRequestEntityTooLarge {
	return &DeleteOutboundSequenceRequestEntityTooLarge{}
}

/*DeleteOutboundSequenceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteOutboundSequenceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteOutboundSequenceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceUnsupportedMediaType creates a DeleteOutboundSequenceUnsupportedMediaType with default headers values
func NewDeleteOutboundSequenceUnsupportedMediaType() *DeleteOutboundSequenceUnsupportedMediaType {
	return &DeleteOutboundSequenceUnsupportedMediaType{}
}

/*DeleteOutboundSequenceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteOutboundSequenceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteOutboundSequenceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceTooManyRequests creates a DeleteOutboundSequenceTooManyRequests with default headers values
func NewDeleteOutboundSequenceTooManyRequests() *DeleteOutboundSequenceTooManyRequests {
	return &DeleteOutboundSequenceTooManyRequests{}
}

/*DeleteOutboundSequenceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteOutboundSequenceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteOutboundSequenceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceInternalServerError creates a DeleteOutboundSequenceInternalServerError with default headers values
func NewDeleteOutboundSequenceInternalServerError() *DeleteOutboundSequenceInternalServerError {
	return &DeleteOutboundSequenceInternalServerError{}
}

/*DeleteOutboundSequenceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteOutboundSequenceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteOutboundSequenceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceServiceUnavailable creates a DeleteOutboundSequenceServiceUnavailable with default headers values
func NewDeleteOutboundSequenceServiceUnavailable() *DeleteOutboundSequenceServiceUnavailable {
	return &DeleteOutboundSequenceServiceUnavailable{}
}

/*DeleteOutboundSequenceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteOutboundSequenceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteOutboundSequenceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOutboundSequenceGatewayTimeout creates a DeleteOutboundSequenceGatewayTimeout with default headers values
func NewDeleteOutboundSequenceGatewayTimeout() *DeleteOutboundSequenceGatewayTimeout {
	return &DeleteOutboundSequenceGatewayTimeout{}
}

/*DeleteOutboundSequenceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteOutboundSequenceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteOutboundSequenceGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/outbound/sequences/{sequenceId}][%d] deleteOutboundSequenceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteOutboundSequenceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteOutboundSequenceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
