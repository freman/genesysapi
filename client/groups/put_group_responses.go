// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutGroupReader is a Reader for the PutGroup structure.
type PutGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutGroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGroupOK creates a PutGroupOK with default headers values
func NewPutGroupOK() *PutGroupOK {
	return &PutGroupOK{}
}

/*PutGroupOK handles this case with default header values.

successful operation
*/
type PutGroupOK struct {
	Payload *models.Group
}

func (o *PutGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupOK  %+v", 200, o.Payload)
}

func (o *PutGroupOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *PutGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupBadRequest creates a PutGroupBadRequest with default headers values
func NewPutGroupBadRequest() *PutGroupBadRequest {
	return &PutGroupBadRequest{}
}

/*PutGroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupUnauthorized creates a PutGroupUnauthorized with default headers values
func NewPutGroupUnauthorized() *PutGroupUnauthorized {
	return &PutGroupUnauthorized{}
}

/*PutGroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupForbidden creates a PutGroupForbidden with default headers values
func NewPutGroupForbidden() *PutGroupForbidden {
	return &PutGroupForbidden{}
}

/*PutGroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutGroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupForbidden  %+v", 403, o.Payload)
}

func (o *PutGroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupNotFound creates a PutGroupNotFound with default headers values
func NewPutGroupNotFound() *PutGroupNotFound {
	return &PutGroupNotFound{}
}

/*PutGroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutGroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupNotFound  %+v", 404, o.Payload)
}

func (o *PutGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupRequestTimeout creates a PutGroupRequestTimeout with default headers values
func NewPutGroupRequestTimeout() *PutGroupRequestTimeout {
	return &PutGroupRequestTimeout{}
}

/*PutGroupRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutGroupRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGroupRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutGroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupConflict creates a PutGroupConflict with default headers values
func NewPutGroupConflict() *PutGroupConflict {
	return &PutGroupConflict{}
}

/*PutGroupConflict handles this case with default header values.

Resource conflict - Unexpected version was provided
*/
type PutGroupConflict struct {
}

func (o *PutGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupConflict ", 409)
}

func (o *PutGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutGroupRequestEntityTooLarge creates a PutGroupRequestEntityTooLarge with default headers values
func NewPutGroupRequestEntityTooLarge() *PutGroupRequestEntityTooLarge {
	return &PutGroupRequestEntityTooLarge{}
}

/*PutGroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutGroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutGroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupUnsupportedMediaType creates a PutGroupUnsupportedMediaType with default headers values
func NewPutGroupUnsupportedMediaType() *PutGroupUnsupportedMediaType {
	return &PutGroupUnsupportedMediaType{}
}

/*PutGroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupTooManyRequests creates a PutGroupTooManyRequests with default headers values
func NewPutGroupTooManyRequests() *PutGroupTooManyRequests {
	return &PutGroupTooManyRequests{}
}

/*PutGroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutGroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupInternalServerError creates a PutGroupInternalServerError with default headers values
func NewPutGroupInternalServerError() *PutGroupInternalServerError {
	return &PutGroupInternalServerError{}
}

/*PutGroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutGroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupServiceUnavailable creates a PutGroupServiceUnavailable with default headers values
func NewPutGroupServiceUnavailable() *PutGroupServiceUnavailable {
	return &PutGroupServiceUnavailable{}
}

/*PutGroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGroupGatewayTimeout creates a PutGroupGatewayTimeout with default headers values
func NewPutGroupGatewayTimeout() *PutGroupGatewayTimeout {
	return &PutGroupGatewayTimeout{}
}

/*PutGroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutGroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/groups/{groupId}][%d] putGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
