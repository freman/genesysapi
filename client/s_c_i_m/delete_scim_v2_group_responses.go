// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteScimV2GroupReader is a Reader for the DeleteScimV2Group structure.
type DeleteScimV2GroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteScimV2GroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteScimV2GroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteScimV2GroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteScimV2GroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteScimV2GroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteScimV2GroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteScimV2GroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteScimV2GroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteScimV2GroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteScimV2GroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteScimV2GroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteScimV2GroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteScimV2GroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteScimV2GroupNoContent creates a DeleteScimV2GroupNoContent with default headers values
func NewDeleteScimV2GroupNoContent() *DeleteScimV2GroupNoContent {
	return &DeleteScimV2GroupNoContent{}
}

/*DeleteScimV2GroupNoContent handles this case with default header values.

Group deleted with no content returned.
*/
type DeleteScimV2GroupNoContent struct {
}

func (o *DeleteScimV2GroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupNoContent ", 204)
}

func (o *DeleteScimV2GroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteScimV2GroupBadRequest creates a DeleteScimV2GroupBadRequest with default headers values
func NewDeleteScimV2GroupBadRequest() *DeleteScimV2GroupBadRequest {
	return &DeleteScimV2GroupBadRequest{}
}

/*DeleteScimV2GroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteScimV2GroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteScimV2GroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupUnauthorized creates a DeleteScimV2GroupUnauthorized with default headers values
func NewDeleteScimV2GroupUnauthorized() *DeleteScimV2GroupUnauthorized {
	return &DeleteScimV2GroupUnauthorized{}
}

/*DeleteScimV2GroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteScimV2GroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteScimV2GroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupForbidden creates a DeleteScimV2GroupForbidden with default headers values
func NewDeleteScimV2GroupForbidden() *DeleteScimV2GroupForbidden {
	return &DeleteScimV2GroupForbidden{}
}

/*DeleteScimV2GroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteScimV2GroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupForbidden  %+v", 403, o.Payload)
}

func (o *DeleteScimV2GroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupNotFound creates a DeleteScimV2GroupNotFound with default headers values
func NewDeleteScimV2GroupNotFound() *DeleteScimV2GroupNotFound {
	return &DeleteScimV2GroupNotFound{}
}

/*DeleteScimV2GroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteScimV2GroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupNotFound  %+v", 404, o.Payload)
}

func (o *DeleteScimV2GroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupConflict creates a DeleteScimV2GroupConflict with default headers values
func NewDeleteScimV2GroupConflict() *DeleteScimV2GroupConflict {
	return &DeleteScimV2GroupConflict{}
}

/*DeleteScimV2GroupConflict handles this case with default header values.

Version does not match current version.
*/
type DeleteScimV2GroupConflict struct {
	Payload *models.ScimError
}

func (o *DeleteScimV2GroupConflict) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupConflict  %+v", 409, o.Payload)
}

func (o *DeleteScimV2GroupConflict) GetPayload() *models.ScimError {
	return o.Payload
}

func (o *DeleteScimV2GroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupRequestEntityTooLarge creates a DeleteScimV2GroupRequestEntityTooLarge with default headers values
func NewDeleteScimV2GroupRequestEntityTooLarge() *DeleteScimV2GroupRequestEntityTooLarge {
	return &DeleteScimV2GroupRequestEntityTooLarge{}
}

/*DeleteScimV2GroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteScimV2GroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteScimV2GroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupUnsupportedMediaType creates a DeleteScimV2GroupUnsupportedMediaType with default headers values
func NewDeleteScimV2GroupUnsupportedMediaType() *DeleteScimV2GroupUnsupportedMediaType {
	return &DeleteScimV2GroupUnsupportedMediaType{}
}

/*DeleteScimV2GroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteScimV2GroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteScimV2GroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupTooManyRequests creates a DeleteScimV2GroupTooManyRequests with default headers values
func NewDeleteScimV2GroupTooManyRequests() *DeleteScimV2GroupTooManyRequests {
	return &DeleteScimV2GroupTooManyRequests{}
}

/*DeleteScimV2GroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteScimV2GroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteScimV2GroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupInternalServerError creates a DeleteScimV2GroupInternalServerError with default headers values
func NewDeleteScimV2GroupInternalServerError() *DeleteScimV2GroupInternalServerError {
	return &DeleteScimV2GroupInternalServerError{}
}

/*DeleteScimV2GroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteScimV2GroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteScimV2GroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupServiceUnavailable creates a DeleteScimV2GroupServiceUnavailable with default headers values
func NewDeleteScimV2GroupServiceUnavailable() *DeleteScimV2GroupServiceUnavailable {
	return &DeleteScimV2GroupServiceUnavailable{}
}

/*DeleteScimV2GroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteScimV2GroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteScimV2GroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteScimV2GroupGatewayTimeout creates a DeleteScimV2GroupGatewayTimeout with default headers values
func NewDeleteScimV2GroupGatewayTimeout() *DeleteScimV2GroupGatewayTimeout {
	return &DeleteScimV2GroupGatewayTimeout{}
}

/*DeleteScimV2GroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteScimV2GroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteScimV2GroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/scim/v2/groups/{groupId}][%d] deleteScimV2GroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteScimV2GroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteScimV2GroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
