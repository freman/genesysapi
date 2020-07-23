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

// PutContentmanagementWorkspaceReader is a Reader for the PutContentmanagementWorkspace structure.
type PutContentmanagementWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutContentmanagementWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutContentmanagementWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutContentmanagementWorkspaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutContentmanagementWorkspaceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutContentmanagementWorkspaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutContentmanagementWorkspaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutContentmanagementWorkspaceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutContentmanagementWorkspaceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutContentmanagementWorkspaceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutContentmanagementWorkspaceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutContentmanagementWorkspaceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutContentmanagementWorkspaceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutContentmanagementWorkspaceOK creates a PutContentmanagementWorkspaceOK with default headers values
func NewPutContentmanagementWorkspaceOK() *PutContentmanagementWorkspaceOK {
	return &PutContentmanagementWorkspaceOK{}
}

/*PutContentmanagementWorkspaceOK handles this case with default header values.

successful operation
*/
type PutContentmanagementWorkspaceOK struct {
	Payload *models.Workspace
}

func (o *PutContentmanagementWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutContentmanagementWorkspaceOK) GetPayload() *models.Workspace {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Workspace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceBadRequest creates a PutContentmanagementWorkspaceBadRequest with default headers values
func NewPutContentmanagementWorkspaceBadRequest() *PutContentmanagementWorkspaceBadRequest {
	return &PutContentmanagementWorkspaceBadRequest{}
}

/*PutContentmanagementWorkspaceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutContentmanagementWorkspaceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *PutContentmanagementWorkspaceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceUnauthorized creates a PutContentmanagementWorkspaceUnauthorized with default headers values
func NewPutContentmanagementWorkspaceUnauthorized() *PutContentmanagementWorkspaceUnauthorized {
	return &PutContentmanagementWorkspaceUnauthorized{}
}

/*PutContentmanagementWorkspaceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutContentmanagementWorkspaceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutContentmanagementWorkspaceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceForbidden creates a PutContentmanagementWorkspaceForbidden with default headers values
func NewPutContentmanagementWorkspaceForbidden() *PutContentmanagementWorkspaceForbidden {
	return &PutContentmanagementWorkspaceForbidden{}
}

/*PutContentmanagementWorkspaceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutContentmanagementWorkspaceForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceForbidden  %+v", 403, o.Payload)
}

func (o *PutContentmanagementWorkspaceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceNotFound creates a PutContentmanagementWorkspaceNotFound with default headers values
func NewPutContentmanagementWorkspaceNotFound() *PutContentmanagementWorkspaceNotFound {
	return &PutContentmanagementWorkspaceNotFound{}
}

/*PutContentmanagementWorkspaceNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutContentmanagementWorkspaceNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceNotFound  %+v", 404, o.Payload)
}

func (o *PutContentmanagementWorkspaceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceRequestEntityTooLarge creates a PutContentmanagementWorkspaceRequestEntityTooLarge with default headers values
func NewPutContentmanagementWorkspaceRequestEntityTooLarge() *PutContentmanagementWorkspaceRequestEntityTooLarge {
	return &PutContentmanagementWorkspaceRequestEntityTooLarge{}
}

/*PutContentmanagementWorkspaceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutContentmanagementWorkspaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceUnsupportedMediaType creates a PutContentmanagementWorkspaceUnsupportedMediaType with default headers values
func NewPutContentmanagementWorkspaceUnsupportedMediaType() *PutContentmanagementWorkspaceUnsupportedMediaType {
	return &PutContentmanagementWorkspaceUnsupportedMediaType{}
}

/*PutContentmanagementWorkspaceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutContentmanagementWorkspaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutContentmanagementWorkspaceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceTooManyRequests creates a PutContentmanagementWorkspaceTooManyRequests with default headers values
func NewPutContentmanagementWorkspaceTooManyRequests() *PutContentmanagementWorkspaceTooManyRequests {
	return &PutContentmanagementWorkspaceTooManyRequests{}
}

/*PutContentmanagementWorkspaceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutContentmanagementWorkspaceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutContentmanagementWorkspaceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceInternalServerError creates a PutContentmanagementWorkspaceInternalServerError with default headers values
func NewPutContentmanagementWorkspaceInternalServerError() *PutContentmanagementWorkspaceInternalServerError {
	return &PutContentmanagementWorkspaceInternalServerError{}
}

/*PutContentmanagementWorkspaceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutContentmanagementWorkspaceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContentmanagementWorkspaceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceServiceUnavailable creates a PutContentmanagementWorkspaceServiceUnavailable with default headers values
func NewPutContentmanagementWorkspaceServiceUnavailable() *PutContentmanagementWorkspaceServiceUnavailable {
	return &PutContentmanagementWorkspaceServiceUnavailable{}
}

/*PutContentmanagementWorkspaceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutContentmanagementWorkspaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutContentmanagementWorkspaceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutContentmanagementWorkspaceGatewayTimeout creates a PutContentmanagementWorkspaceGatewayTimeout with default headers values
func NewPutContentmanagementWorkspaceGatewayTimeout() *PutContentmanagementWorkspaceGatewayTimeout {
	return &PutContentmanagementWorkspaceGatewayTimeout{}
}

/*PutContentmanagementWorkspaceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutContentmanagementWorkspaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutContentmanagementWorkspaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutContentmanagementWorkspaceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
