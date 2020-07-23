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

// DeleteUserRoutingskillReader is a Reader for the DeleteUserRoutingskill structure.
type DeleteUserRoutingskillReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserRoutingskillReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUserRoutingskillOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteUserRoutingskillBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteUserRoutingskillUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUserRoutingskillForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUserRoutingskillNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteUserRoutingskillConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteUserRoutingskillRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteUserRoutingskillUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteUserRoutingskillTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUserRoutingskillInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteUserRoutingskillServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteUserRoutingskillGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserRoutingskillOK creates a DeleteUserRoutingskillOK with default headers values
func NewDeleteUserRoutingskillOK() *DeleteUserRoutingskillOK {
	return &DeleteUserRoutingskillOK{}
}

/*DeleteUserRoutingskillOK handles this case with default header values.

Operation was successful.
*/
type DeleteUserRoutingskillOK struct {
}

func (o *DeleteUserRoutingskillOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillOK ", 200)
}

func (o *DeleteUserRoutingskillOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserRoutingskillBadRequest creates a DeleteUserRoutingskillBadRequest with default headers values
func NewDeleteUserRoutingskillBadRequest() *DeleteUserRoutingskillBadRequest {
	return &DeleteUserRoutingskillBadRequest{}
}

/*DeleteUserRoutingskillBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteUserRoutingskillBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserRoutingskillBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillUnauthorized creates a DeleteUserRoutingskillUnauthorized with default headers values
func NewDeleteUserRoutingskillUnauthorized() *DeleteUserRoutingskillUnauthorized {
	return &DeleteUserRoutingskillUnauthorized{}
}

/*DeleteUserRoutingskillUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteUserRoutingskillUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserRoutingskillUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillForbidden creates a DeleteUserRoutingskillForbidden with default headers values
func NewDeleteUserRoutingskillForbidden() *DeleteUserRoutingskillForbidden {
	return &DeleteUserRoutingskillForbidden{}
}

/*DeleteUserRoutingskillForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteUserRoutingskillForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserRoutingskillForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillNotFound creates a DeleteUserRoutingskillNotFound with default headers values
func NewDeleteUserRoutingskillNotFound() *DeleteUserRoutingskillNotFound {
	return &DeleteUserRoutingskillNotFound{}
}

/*DeleteUserRoutingskillNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteUserRoutingskillNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserRoutingskillNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillConflict creates a DeleteUserRoutingskillConflict with default headers values
func NewDeleteUserRoutingskillConflict() *DeleteUserRoutingskillConflict {
	return &DeleteUserRoutingskillConflict{}
}

/*DeleteUserRoutingskillConflict handles this case with default header values.

Conflict
*/
type DeleteUserRoutingskillConflict struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillConflict  %+v", 409, o.Payload)
}

func (o *DeleteUserRoutingskillConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillRequestEntityTooLarge creates a DeleteUserRoutingskillRequestEntityTooLarge with default headers values
func NewDeleteUserRoutingskillRequestEntityTooLarge() *DeleteUserRoutingskillRequestEntityTooLarge {
	return &DeleteUserRoutingskillRequestEntityTooLarge{}
}

/*DeleteUserRoutingskillRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteUserRoutingskillRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteUserRoutingskillRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillUnsupportedMediaType creates a DeleteUserRoutingskillUnsupportedMediaType with default headers values
func NewDeleteUserRoutingskillUnsupportedMediaType() *DeleteUserRoutingskillUnsupportedMediaType {
	return &DeleteUserRoutingskillUnsupportedMediaType{}
}

/*DeleteUserRoutingskillUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteUserRoutingskillUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteUserRoutingskillUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillTooManyRequests creates a DeleteUserRoutingskillTooManyRequests with default headers values
func NewDeleteUserRoutingskillTooManyRequests() *DeleteUserRoutingskillTooManyRequests {
	return &DeleteUserRoutingskillTooManyRequests{}
}

/*DeleteUserRoutingskillTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteUserRoutingskillTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteUserRoutingskillTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillInternalServerError creates a DeleteUserRoutingskillInternalServerError with default headers values
func NewDeleteUserRoutingskillInternalServerError() *DeleteUserRoutingskillInternalServerError {
	return &DeleteUserRoutingskillInternalServerError{}
}

/*DeleteUserRoutingskillInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteUserRoutingskillInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteUserRoutingskillInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillServiceUnavailable creates a DeleteUserRoutingskillServiceUnavailable with default headers values
func NewDeleteUserRoutingskillServiceUnavailable() *DeleteUserRoutingskillServiceUnavailable {
	return &DeleteUserRoutingskillServiceUnavailable{}
}

/*DeleteUserRoutingskillServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteUserRoutingskillServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteUserRoutingskillServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserRoutingskillGatewayTimeout creates a DeleteUserRoutingskillGatewayTimeout with default headers values
func NewDeleteUserRoutingskillGatewayTimeout() *DeleteUserRoutingskillGatewayTimeout {
	return &DeleteUserRoutingskillGatewayTimeout{}
}

/*DeleteUserRoutingskillGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteUserRoutingskillGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteUserRoutingskillGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/users/{userId}/routingskills/{skillId}][%d] deleteUserRoutingskillGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteUserRoutingskillGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteUserRoutingskillGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
