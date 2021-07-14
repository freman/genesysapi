// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteFlowsMilestoneReader is a Reader for the DeleteFlowsMilestone structure.
type DeleteFlowsMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlowsMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFlowsMilestoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteFlowsMilestoneNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteFlowsMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteFlowsMilestoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteFlowsMilestoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteFlowsMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewDeleteFlowsMilestoneMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteFlowsMilestoneRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteFlowsMilestoneConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteFlowsMilestoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteFlowsMilestoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFlowsMilestoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFlowsMilestoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteFlowsMilestoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteFlowsMilestoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteFlowsMilestoneOK creates a DeleteFlowsMilestoneOK with default headers values
func NewDeleteFlowsMilestoneOK() *DeleteFlowsMilestoneOK {
	return &DeleteFlowsMilestoneOK{}
}

/*DeleteFlowsMilestoneOK handles this case with default header values.

successful operation
*/
type DeleteFlowsMilestoneOK struct {
	Payload models.Empty
}

func (o *DeleteFlowsMilestoneOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneOK  %+v", 200, o.Payload)
}

func (o *DeleteFlowsMilestoneOK) GetPayload() models.Empty {
	return o.Payload
}

func (o *DeleteFlowsMilestoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneNoContent creates a DeleteFlowsMilestoneNoContent with default headers values
func NewDeleteFlowsMilestoneNoContent() *DeleteFlowsMilestoneNoContent {
	return &DeleteFlowsMilestoneNoContent{}
}

/*DeleteFlowsMilestoneNoContent handles this case with default header values.

Delete was successful.
*/
type DeleteFlowsMilestoneNoContent struct {
}

func (o *DeleteFlowsMilestoneNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNoContent ", 204)
}

func (o *DeleteFlowsMilestoneNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteFlowsMilestoneBadRequest creates a DeleteFlowsMilestoneBadRequest with default headers values
func NewDeleteFlowsMilestoneBadRequest() *DeleteFlowsMilestoneBadRequest {
	return &DeleteFlowsMilestoneBadRequest{}
}

/*DeleteFlowsMilestoneBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteFlowsMilestoneBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteFlowsMilestoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneUnauthorized creates a DeleteFlowsMilestoneUnauthorized with default headers values
func NewDeleteFlowsMilestoneUnauthorized() *DeleteFlowsMilestoneUnauthorized {
	return &DeleteFlowsMilestoneUnauthorized{}
}

/*DeleteFlowsMilestoneUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteFlowsMilestoneUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteFlowsMilestoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneForbidden creates a DeleteFlowsMilestoneForbidden with default headers values
func NewDeleteFlowsMilestoneForbidden() *DeleteFlowsMilestoneForbidden {
	return &DeleteFlowsMilestoneForbidden{}
}

/*DeleteFlowsMilestoneForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteFlowsMilestoneForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFlowsMilestoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneNotFound creates a DeleteFlowsMilestoneNotFound with default headers values
func NewDeleteFlowsMilestoneNotFound() *DeleteFlowsMilestoneNotFound {
	return &DeleteFlowsMilestoneNotFound{}
}

/*DeleteFlowsMilestoneNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteFlowsMilestoneNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneNotFound  %+v", 404, o.Payload)
}

func (o *DeleteFlowsMilestoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneMethodNotAllowed creates a DeleteFlowsMilestoneMethodNotAllowed with default headers values
func NewDeleteFlowsMilestoneMethodNotAllowed() *DeleteFlowsMilestoneMethodNotAllowed {
	return &DeleteFlowsMilestoneMethodNotAllowed{}
}

/*DeleteFlowsMilestoneMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type DeleteFlowsMilestoneMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneRequestTimeout creates a DeleteFlowsMilestoneRequestTimeout with default headers values
func NewDeleteFlowsMilestoneRequestTimeout() *DeleteFlowsMilestoneRequestTimeout {
	return &DeleteFlowsMilestoneRequestTimeout{}
}

/*DeleteFlowsMilestoneRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteFlowsMilestoneRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneConflict creates a DeleteFlowsMilestoneConflict with default headers values
func NewDeleteFlowsMilestoneConflict() *DeleteFlowsMilestoneConflict {
	return &DeleteFlowsMilestoneConflict{}
}

/*DeleteFlowsMilestoneConflict handles this case with default header values.

Conflict
*/
type DeleteFlowsMilestoneConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneConflict  %+v", 409, o.Payload)
}

func (o *DeleteFlowsMilestoneConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneRequestEntityTooLarge creates a DeleteFlowsMilestoneRequestEntityTooLarge with default headers values
func NewDeleteFlowsMilestoneRequestEntityTooLarge() *DeleteFlowsMilestoneRequestEntityTooLarge {
	return &DeleteFlowsMilestoneRequestEntityTooLarge{}
}

/*DeleteFlowsMilestoneRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteFlowsMilestoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneUnsupportedMediaType creates a DeleteFlowsMilestoneUnsupportedMediaType with default headers values
func NewDeleteFlowsMilestoneUnsupportedMediaType() *DeleteFlowsMilestoneUnsupportedMediaType {
	return &DeleteFlowsMilestoneUnsupportedMediaType{}
}

/*DeleteFlowsMilestoneUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteFlowsMilestoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneTooManyRequests creates a DeleteFlowsMilestoneTooManyRequests with default headers values
func NewDeleteFlowsMilestoneTooManyRequests() *DeleteFlowsMilestoneTooManyRequests {
	return &DeleteFlowsMilestoneTooManyRequests{}
}

/*DeleteFlowsMilestoneTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteFlowsMilestoneTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFlowsMilestoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneInternalServerError creates a DeleteFlowsMilestoneInternalServerError with default headers values
func NewDeleteFlowsMilestoneInternalServerError() *DeleteFlowsMilestoneInternalServerError {
	return &DeleteFlowsMilestoneInternalServerError{}
}

/*DeleteFlowsMilestoneInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteFlowsMilestoneInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFlowsMilestoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneServiceUnavailable creates a DeleteFlowsMilestoneServiceUnavailable with default headers values
func NewDeleteFlowsMilestoneServiceUnavailable() *DeleteFlowsMilestoneServiceUnavailable {
	return &DeleteFlowsMilestoneServiceUnavailable{}
}

/*DeleteFlowsMilestoneServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteFlowsMilestoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteFlowsMilestoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFlowsMilestoneGatewayTimeout creates a DeleteFlowsMilestoneGatewayTimeout with default headers values
func NewDeleteFlowsMilestoneGatewayTimeout() *DeleteFlowsMilestoneGatewayTimeout {
	return &DeleteFlowsMilestoneGatewayTimeout{}
}

/*DeleteFlowsMilestoneGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteFlowsMilestoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteFlowsMilestoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/flows/milestones/{milestoneId}][%d] deleteFlowsMilestoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteFlowsMilestoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteFlowsMilestoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
