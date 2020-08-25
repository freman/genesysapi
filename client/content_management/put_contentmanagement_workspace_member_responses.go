// Code generated by go-swagger; DO NOT EDIT.

package content_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutContentmanagementWorkspaceMemberReader is a Reader for the PutContentmanagementWorkspaceMember structure.
type PutContentmanagementWorkspaceMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContentmanagementWorkspaceMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContentmanagementWorkspaceMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutContentmanagementWorkspaceMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutContentmanagementWorkspaceMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutContentmanagementWorkspaceMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutContentmanagementWorkspaceMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutContentmanagementWorkspaceMemberRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutContentmanagementWorkspaceMemberUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutContentmanagementWorkspaceMemberTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutContentmanagementWorkspaceMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutContentmanagementWorkspaceMemberServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutContentmanagementWorkspaceMemberGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutContentmanagementWorkspaceMemberOK creates a PutContentmanagementWorkspaceMemberOK with default headers values
func NewPutContentmanagementWorkspaceMemberOK() *PutContentmanagementWorkspaceMemberOK {
	return &PutContentmanagementWorkspaceMemberOK{}
}

/*PutContentmanagementWorkspaceMemberOK handles this case with default header values.

successful operation
*/
type PutContentmanagementWorkspaceMemberOK struct {
	Payload *models.WorkspaceMember
}

func (o *PutContentmanagementWorkspaceMemberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberOK  %+v", 200, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberOK) GetPayload() *models.WorkspaceMember {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceMember)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberBadRequest creates a PutContentmanagementWorkspaceMemberBadRequest with default headers values
func NewPutContentmanagementWorkspaceMemberBadRequest() *PutContentmanagementWorkspaceMemberBadRequest {
	return &PutContentmanagementWorkspaceMemberBadRequest{}
}

/*PutContentmanagementWorkspaceMemberBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutContentmanagementWorkspaceMemberBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberBadRequest  %+v", 400, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberUnauthorized creates a PutContentmanagementWorkspaceMemberUnauthorized with default headers values
func NewPutContentmanagementWorkspaceMemberUnauthorized() *PutContentmanagementWorkspaceMemberUnauthorized {
	return &PutContentmanagementWorkspaceMemberUnauthorized{}
}

/*PutContentmanagementWorkspaceMemberUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutContentmanagementWorkspaceMemberUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberForbidden creates a PutContentmanagementWorkspaceMemberForbidden with default headers values
func NewPutContentmanagementWorkspaceMemberForbidden() *PutContentmanagementWorkspaceMemberForbidden {
	return &PutContentmanagementWorkspaceMemberForbidden{}
}

/*PutContentmanagementWorkspaceMemberForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutContentmanagementWorkspaceMemberForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberForbidden  %+v", 403, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberNotFound creates a PutContentmanagementWorkspaceMemberNotFound with default headers values
func NewPutContentmanagementWorkspaceMemberNotFound() *PutContentmanagementWorkspaceMemberNotFound {
	return &PutContentmanagementWorkspaceMemberNotFound{}
}

/*PutContentmanagementWorkspaceMemberNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutContentmanagementWorkspaceMemberNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberNotFound  %+v", 404, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberRequestEntityTooLarge creates a PutContentmanagementWorkspaceMemberRequestEntityTooLarge with default headers values
func NewPutContentmanagementWorkspaceMemberRequestEntityTooLarge() *PutContentmanagementWorkspaceMemberRequestEntityTooLarge {
	return &PutContentmanagementWorkspaceMemberRequestEntityTooLarge{}
}

/*PutContentmanagementWorkspaceMemberRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutContentmanagementWorkspaceMemberRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberUnsupportedMediaType creates a PutContentmanagementWorkspaceMemberUnsupportedMediaType with default headers values
func NewPutContentmanagementWorkspaceMemberUnsupportedMediaType() *PutContentmanagementWorkspaceMemberUnsupportedMediaType {
	return &PutContentmanagementWorkspaceMemberUnsupportedMediaType{}
}

/*PutContentmanagementWorkspaceMemberUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutContentmanagementWorkspaceMemberUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberTooManyRequests creates a PutContentmanagementWorkspaceMemberTooManyRequests with default headers values
func NewPutContentmanagementWorkspaceMemberTooManyRequests() *PutContentmanagementWorkspaceMemberTooManyRequests {
	return &PutContentmanagementWorkspaceMemberTooManyRequests{}
}

/*PutContentmanagementWorkspaceMemberTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutContentmanagementWorkspaceMemberTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberInternalServerError creates a PutContentmanagementWorkspaceMemberInternalServerError with default headers values
func NewPutContentmanagementWorkspaceMemberInternalServerError() *PutContentmanagementWorkspaceMemberInternalServerError {
	return &PutContentmanagementWorkspaceMemberInternalServerError{}
}

/*PutContentmanagementWorkspaceMemberInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutContentmanagementWorkspaceMemberInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberServiceUnavailable creates a PutContentmanagementWorkspaceMemberServiceUnavailable with default headers values
func NewPutContentmanagementWorkspaceMemberServiceUnavailable() *PutContentmanagementWorkspaceMemberServiceUnavailable {
	return &PutContentmanagementWorkspaceMemberServiceUnavailable{}
}

/*PutContentmanagementWorkspaceMemberServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutContentmanagementWorkspaceMemberServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceMemberGatewayTimeout creates a PutContentmanagementWorkspaceMemberGatewayTimeout with default headers values
func NewPutContentmanagementWorkspaceMemberGatewayTimeout() *PutContentmanagementWorkspaceMemberGatewayTimeout {
	return &PutContentmanagementWorkspaceMemberGatewayTimeout{}
}

/*PutContentmanagementWorkspaceMemberGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutContentmanagementWorkspaceMemberGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceMemberGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}/members/{memberId}][%d] putContentmanagementWorkspaceMemberGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutContentmanagementWorkspaceMemberGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceMemberGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}