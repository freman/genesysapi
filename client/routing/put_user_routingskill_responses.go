// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutUserRoutingskillReader is a Reader for the PutUserRoutingskill structure.
type PutUserRoutingskillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutUserRoutingskillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutUserRoutingskillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutUserRoutingskillBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutUserRoutingskillUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutUserRoutingskillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutUserRoutingskillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutUserRoutingskillConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutUserRoutingskillRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutUserRoutingskillUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutUserRoutingskillTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutUserRoutingskillInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutUserRoutingskillServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutUserRoutingskillGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutUserRoutingskillOK creates a PutUserRoutingskillOK with default headers values
func NewPutUserRoutingskillOK() *PutUserRoutingskillOK {
	return &PutUserRoutingskillOK{}
}

/*PutUserRoutingskillOK handles this case with default header values.

successful operation
*/
type PutUserRoutingskillOK struct {
	Payload *models.UserRoutingSkill
}

func (o *PutUserRoutingskillOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillOK  %+v", 200, o.Payload)
}

func (o *PutUserRoutingskillOK) GetPayload() *models.UserRoutingSkill {
	return o.Payload
}

func (o *PutUserRoutingskillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserRoutingSkill)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillBadRequest creates a PutUserRoutingskillBadRequest with default headers values
func NewPutUserRoutingskillBadRequest() *PutUserRoutingskillBadRequest {
	return &PutUserRoutingskillBadRequest{}
}

/*PutUserRoutingskillBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutUserRoutingskillBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillBadRequest  %+v", 400, o.Payload)
}

func (o *PutUserRoutingskillBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillUnauthorized creates a PutUserRoutingskillUnauthorized with default headers values
func NewPutUserRoutingskillUnauthorized() *PutUserRoutingskillUnauthorized {
	return &PutUserRoutingskillUnauthorized{}
}

/*PutUserRoutingskillUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutUserRoutingskillUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillUnauthorized  %+v", 401, o.Payload)
}

func (o *PutUserRoutingskillUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillForbidden creates a PutUserRoutingskillForbidden with default headers values
func NewPutUserRoutingskillForbidden() *PutUserRoutingskillForbidden {
	return &PutUserRoutingskillForbidden{}
}

/*PutUserRoutingskillForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutUserRoutingskillForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillForbidden  %+v", 403, o.Payload)
}

func (o *PutUserRoutingskillForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillNotFound creates a PutUserRoutingskillNotFound with default headers values
func NewPutUserRoutingskillNotFound() *PutUserRoutingskillNotFound {
	return &PutUserRoutingskillNotFound{}
}

/*PutUserRoutingskillNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutUserRoutingskillNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillNotFound  %+v", 404, o.Payload)
}

func (o *PutUserRoutingskillNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillConflict creates a PutUserRoutingskillConflict with default headers values
func NewPutUserRoutingskillConflict() *PutUserRoutingskillConflict {
	return &PutUserRoutingskillConflict{}
}

/*PutUserRoutingskillConflict handles this case with default header values.

Resource conflict - Unexpected version was provided
*/
type PutUserRoutingskillConflict struct {
}

func (o *PutUserRoutingskillConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillConflict ", 409)
}

func (o *PutUserRoutingskillConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutUserRoutingskillRequestEntityTooLarge creates a PutUserRoutingskillRequestEntityTooLarge with default headers values
func NewPutUserRoutingskillRequestEntityTooLarge() *PutUserRoutingskillRequestEntityTooLarge {
	return &PutUserRoutingskillRequestEntityTooLarge{}
}

/*PutUserRoutingskillRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutUserRoutingskillRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutUserRoutingskillRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillUnsupportedMediaType creates a PutUserRoutingskillUnsupportedMediaType with default headers values
func NewPutUserRoutingskillUnsupportedMediaType() *PutUserRoutingskillUnsupportedMediaType {
	return &PutUserRoutingskillUnsupportedMediaType{}
}

/*PutUserRoutingskillUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutUserRoutingskillUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutUserRoutingskillUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillTooManyRequests creates a PutUserRoutingskillTooManyRequests with default headers values
func NewPutUserRoutingskillTooManyRequests() *PutUserRoutingskillTooManyRequests {
	return &PutUserRoutingskillTooManyRequests{}
}

/*PutUserRoutingskillTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutUserRoutingskillTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutUserRoutingskillTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillInternalServerError creates a PutUserRoutingskillInternalServerError with default headers values
func NewPutUserRoutingskillInternalServerError() *PutUserRoutingskillInternalServerError {
	return &PutUserRoutingskillInternalServerError{}
}

/*PutUserRoutingskillInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutUserRoutingskillInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillInternalServerError  %+v", 500, o.Payload)
}

func (o *PutUserRoutingskillInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillServiceUnavailable creates a PutUserRoutingskillServiceUnavailable with default headers values
func NewPutUserRoutingskillServiceUnavailable() *PutUserRoutingskillServiceUnavailable {
	return &PutUserRoutingskillServiceUnavailable{}
}

/*PutUserRoutingskillServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutUserRoutingskillServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutUserRoutingskillServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutUserRoutingskillGatewayTimeout creates a PutUserRoutingskillGatewayTimeout with default headers values
func NewPutUserRoutingskillGatewayTimeout() *PutUserRoutingskillGatewayTimeout {
	return &PutUserRoutingskillGatewayTimeout{}
}

/*PutUserRoutingskillGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutUserRoutingskillGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutUserRoutingskillGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/users/{userId}/routingskills/{skillId}][%d] putUserRoutingskillGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutUserRoutingskillGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutUserRoutingskillGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
