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

// PutFlowReader is a Reader for the PutFlow structure.
type PutFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutFlowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutFlowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutFlowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutFlowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutFlowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewPutFlowMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutFlowRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutFlowConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPutFlowGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutFlowRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutFlowUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutFlowTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutFlowServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutFlowGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutFlowOK creates a PutFlowOK with default headers values
func NewPutFlowOK() *PutFlowOK {
	return &PutFlowOK{}
}

/*PutFlowOK handles this case with default header values.

successful operation
*/
type PutFlowOK struct {
	Payload *models.Flow
}

func (o *PutFlowOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowOK  %+v", 200, o.Payload)
}

func (o *PutFlowOK) GetPayload() *models.Flow {
	return o.Payload
}

func (o *PutFlowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Flow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowBadRequest creates a PutFlowBadRequest with default headers values
func NewPutFlowBadRequest() *PutFlowBadRequest {
	return &PutFlowBadRequest{}
}

/*PutFlowBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutFlowBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutFlowBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowBadRequest  %+v", 400, o.Payload)
}

func (o *PutFlowBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowUnauthorized creates a PutFlowUnauthorized with default headers values
func NewPutFlowUnauthorized() *PutFlowUnauthorized {
	return &PutFlowUnauthorized{}
}

/*PutFlowUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutFlowUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutFlowUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFlowUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowForbidden creates a PutFlowForbidden with default headers values
func NewPutFlowForbidden() *PutFlowForbidden {
	return &PutFlowForbidden{}
}

/*PutFlowForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutFlowForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutFlowForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowForbidden  %+v", 403, o.Payload)
}

func (o *PutFlowForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowNotFound creates a PutFlowNotFound with default headers values
func NewPutFlowNotFound() *PutFlowNotFound {
	return &PutFlowNotFound{}
}

/*PutFlowNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutFlowNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutFlowNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowNotFound  %+v", 404, o.Payload)
}

func (o *PutFlowNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowMethodNotAllowed creates a PutFlowMethodNotAllowed with default headers values
func NewPutFlowMethodNotAllowed() *PutFlowMethodNotAllowed {
	return &PutFlowMethodNotAllowed{}
}

/*PutFlowMethodNotAllowed handles this case with default header values.

Method Not Allowed
*/
type PutFlowMethodNotAllowed struct {
	Payload *models.ErrorBody
}

func (o *PutFlowMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *PutFlowMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowRequestTimeout creates a PutFlowRequestTimeout with default headers values
func NewPutFlowRequestTimeout() *PutFlowRequestTimeout {
	return &PutFlowRequestTimeout{}
}

/*PutFlowRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutFlowRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutFlowRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutFlowRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowConflict creates a PutFlowConflict with default headers values
func NewPutFlowConflict() *PutFlowConflict {
	return &PutFlowConflict{}
}

/*PutFlowConflict handles this case with default header values.

Conflict
*/
type PutFlowConflict struct {
	Payload *models.ErrorBody
}

func (o *PutFlowConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowConflict  %+v", 409, o.Payload)
}

func (o *PutFlowConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowGone creates a PutFlowGone with default headers values
func NewPutFlowGone() *PutFlowGone {
	return &PutFlowGone{}
}

/*PutFlowGone handles this case with default header values.

Gone
*/
type PutFlowGone struct {
	Payload *models.ErrorBody
}

func (o *PutFlowGone) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGone  %+v", 410, o.Payload)
}

func (o *PutFlowGone) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowRequestEntityTooLarge creates a PutFlowRequestEntityTooLarge with default headers values
func NewPutFlowRequestEntityTooLarge() *PutFlowRequestEntityTooLarge {
	return &PutFlowRequestEntityTooLarge{}
}

/*PutFlowRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutFlowRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutFlowRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutFlowRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowUnsupportedMediaType creates a PutFlowUnsupportedMediaType with default headers values
func NewPutFlowUnsupportedMediaType() *PutFlowUnsupportedMediaType {
	return &PutFlowUnsupportedMediaType{}
}

/*PutFlowUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutFlowUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutFlowUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutFlowUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowTooManyRequests creates a PutFlowTooManyRequests with default headers values
func NewPutFlowTooManyRequests() *PutFlowTooManyRequests {
	return &PutFlowTooManyRequests{}
}

/*PutFlowTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutFlowTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutFlowTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutFlowTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowInternalServerError creates a PutFlowInternalServerError with default headers values
func NewPutFlowInternalServerError() *PutFlowInternalServerError {
	return &PutFlowInternalServerError{}
}

/*PutFlowInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutFlowInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutFlowInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFlowInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowServiceUnavailable creates a PutFlowServiceUnavailable with default headers values
func NewPutFlowServiceUnavailable() *PutFlowServiceUnavailable {
	return &PutFlowServiceUnavailable{}
}

/*PutFlowServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutFlowServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutFlowServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFlowServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFlowGatewayTimeout creates a PutFlowGatewayTimeout with default headers values
func NewPutFlowGatewayTimeout() *PutFlowGatewayTimeout {
	return &PutFlowGatewayTimeout{}
}

/*PutFlowGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutFlowGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutFlowGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/flows/{flowId}][%d] putFlowGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFlowGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutFlowGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
