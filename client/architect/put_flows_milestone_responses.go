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

// PutFlowsMilestoneReader is a Reader for the PutFlowsMilestone structure.
type PutFlowsMilestoneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFlowsMilestoneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFlowsMilestoneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFlowsMilestoneBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFlowsMilestoneUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFlowsMilestoneForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFlowsMilestoneNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutFlowsMilestoneMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutFlowsMilestoneRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutFlowsMilestoneUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutFlowsMilestoneTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFlowsMilestoneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFlowsMilestoneServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFlowsMilestoneGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFlowsMilestoneOK creates a PutFlowsMilestoneOK with default headers values
func NewPutFlowsMilestoneOK() *PutFlowsMilestoneOK {
	return &PutFlowsMilestoneOK{}
}

/*PutFlowsMilestoneOK handles this case with default header values.

successful operation
*/
type PutFlowsMilestoneOK struct {
	Payload *models.FlowMilestone
}

func (o *PutFlowsMilestoneOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneOK  %+v", 200, o.Payload)
}

func (o *PutFlowsMilestoneOK) GetPayload() *models.FlowMilestone {
	return o.Payload
}

func (o *PutFlowsMilestoneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowMilestone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneBadRequest creates a PutFlowsMilestoneBadRequest with default headers values
func NewPutFlowsMilestoneBadRequest() *PutFlowsMilestoneBadRequest {
	return &PutFlowsMilestoneBadRequest{}
}

/*PutFlowsMilestoneBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutFlowsMilestoneBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowsMilestoneBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneUnauthorized creates a PutFlowsMilestoneUnauthorized with default headers values
func NewPutFlowsMilestoneUnauthorized() *PutFlowsMilestoneUnauthorized {
	return &PutFlowsMilestoneUnauthorized{}
}

/*PutFlowsMilestoneUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutFlowsMilestoneUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowsMilestoneUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneForbidden creates a PutFlowsMilestoneForbidden with default headers values
func NewPutFlowsMilestoneForbidden() *PutFlowsMilestoneForbidden {
	return &PutFlowsMilestoneForbidden{}
}

/*PutFlowsMilestoneForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutFlowsMilestoneForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowsMilestoneForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneNotFound creates a PutFlowsMilestoneNotFound with default headers values
func NewPutFlowsMilestoneNotFound() *PutFlowsMilestoneNotFound {
	return &PutFlowsMilestoneNotFound{}
}

/*PutFlowsMilestoneNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutFlowsMilestoneNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowsMilestoneNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneMethodNotAllowed creates a PutFlowsMilestoneMethodNotAllowed with default headers values
func NewPutFlowsMilestoneMethodNotAllowed() *PutFlowsMilestoneMethodNotAllowed {
	return &PutFlowsMilestoneMethodNotAllowed{}
}

/*PutFlowsMilestoneMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PutFlowsMilestoneMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowsMilestoneMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneRequestEntityTooLarge creates a PutFlowsMilestoneRequestEntityTooLarge with default headers values
func NewPutFlowsMilestoneRequestEntityTooLarge() *PutFlowsMilestoneRequestEntityTooLarge {
	return &PutFlowsMilestoneRequestEntityTooLarge{}
}

/*PutFlowsMilestoneRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutFlowsMilestoneRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowsMilestoneRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneUnsupportedMediaType creates a PutFlowsMilestoneUnsupportedMediaType with default headers values
func NewPutFlowsMilestoneUnsupportedMediaType() *PutFlowsMilestoneUnsupportedMediaType {
	return &PutFlowsMilestoneUnsupportedMediaType{}
}

/*PutFlowsMilestoneUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutFlowsMilestoneUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowsMilestoneUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneTooManyRequests creates a PutFlowsMilestoneTooManyRequests with default headers values
func NewPutFlowsMilestoneTooManyRequests() *PutFlowsMilestoneTooManyRequests {
	return &PutFlowsMilestoneTooManyRequests{}
}

/*PutFlowsMilestoneTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutFlowsMilestoneTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowsMilestoneTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneInternalServerError creates a PutFlowsMilestoneInternalServerError with default headers values
func NewPutFlowsMilestoneInternalServerError() *PutFlowsMilestoneInternalServerError {
	return &PutFlowsMilestoneInternalServerError{}
}

/*PutFlowsMilestoneInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutFlowsMilestoneInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowsMilestoneInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneServiceUnavailable creates a PutFlowsMilestoneServiceUnavailable with default headers values
func NewPutFlowsMilestoneServiceUnavailable() *PutFlowsMilestoneServiceUnavailable {
	return &PutFlowsMilestoneServiceUnavailable{}
}

/*PutFlowsMilestoneServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutFlowsMilestoneServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowsMilestoneServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowsMilestoneGatewayTimeout creates a PutFlowsMilestoneGatewayTimeout with default headers values
func NewPutFlowsMilestoneGatewayTimeout() *PutFlowsMilestoneGatewayTimeout {
	return &PutFlowsMilestoneGatewayTimeout{}
}

/*PutFlowsMilestoneGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutFlowsMilestoneGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutFlowsMilestoneGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/milestones/{milestoneId}][%d] putFlowsMilestoneGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowsMilestoneGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowsMilestoneGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}