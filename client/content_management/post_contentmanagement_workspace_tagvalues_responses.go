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

// PostContentmanagementWorkspaceTagvaluesReader is a Reader for the PostContentmanagementWorkspaceTagvalues structure.
type PostContentmanagementWorkspaceTagvaluesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContentmanagementWorkspaceTagvaluesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContentmanagementWorkspaceTagvaluesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContentmanagementWorkspaceTagvaluesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContentmanagementWorkspaceTagvaluesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostContentmanagementWorkspaceTagvaluesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostContentmanagementWorkspaceTagvaluesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostContentmanagementWorkspaceTagvaluesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostContentmanagementWorkspaceTagvaluesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContentmanagementWorkspaceTagvaluesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostContentmanagementWorkspaceTagvaluesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostContentmanagementWorkspaceTagvaluesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContentmanagementWorkspaceTagvaluesOK creates a PostContentmanagementWorkspaceTagvaluesOK with default headers values
func NewPostContentmanagementWorkspaceTagvaluesOK() *PostContentmanagementWorkspaceTagvaluesOK {
	return &PostContentmanagementWorkspaceTagvaluesOK{}
}

/*PostContentmanagementWorkspaceTagvaluesOK handles this case with default header values.

successful operation
*/
type PostContentmanagementWorkspaceTagvaluesOK struct {
	Payload *models.TagValue
}

func (o *PostContentmanagementWorkspaceTagvaluesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesOK  %+v", 200, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesOK) GetPayload() *models.TagValue {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesBadRequest creates a PostContentmanagementWorkspaceTagvaluesBadRequest with default headers values
func NewPostContentmanagementWorkspaceTagvaluesBadRequest() *PostContentmanagementWorkspaceTagvaluesBadRequest {
	return &PostContentmanagementWorkspaceTagvaluesBadRequest{}
}

/*PostContentmanagementWorkspaceTagvaluesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostContentmanagementWorkspaceTagvaluesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesBadRequest  %+v", 400, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesUnauthorized creates a PostContentmanagementWorkspaceTagvaluesUnauthorized with default headers values
func NewPostContentmanagementWorkspaceTagvaluesUnauthorized() *PostContentmanagementWorkspaceTagvaluesUnauthorized {
	return &PostContentmanagementWorkspaceTagvaluesUnauthorized{}
}

/*PostContentmanagementWorkspaceTagvaluesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostContentmanagementWorkspaceTagvaluesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesForbidden creates a PostContentmanagementWorkspaceTagvaluesForbidden with default headers values
func NewPostContentmanagementWorkspaceTagvaluesForbidden() *PostContentmanagementWorkspaceTagvaluesForbidden {
	return &PostContentmanagementWorkspaceTagvaluesForbidden{}
}

/*PostContentmanagementWorkspaceTagvaluesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostContentmanagementWorkspaceTagvaluesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesForbidden  %+v", 403, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesNotFound creates a PostContentmanagementWorkspaceTagvaluesNotFound with default headers values
func NewPostContentmanagementWorkspaceTagvaluesNotFound() *PostContentmanagementWorkspaceTagvaluesNotFound {
	return &PostContentmanagementWorkspaceTagvaluesNotFound{}
}

/*PostContentmanagementWorkspaceTagvaluesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostContentmanagementWorkspaceTagvaluesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesNotFound  %+v", 404, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge creates a PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge with default headers values
func NewPostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge() *PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge {
	return &PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge{}
}

/*PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesUnsupportedMediaType creates a PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType with default headers values
func NewPostContentmanagementWorkspaceTagvaluesUnsupportedMediaType() *PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType {
	return &PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType{}
}

/*PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesTooManyRequests creates a PostContentmanagementWorkspaceTagvaluesTooManyRequests with default headers values
func NewPostContentmanagementWorkspaceTagvaluesTooManyRequests() *PostContentmanagementWorkspaceTagvaluesTooManyRequests {
	return &PostContentmanagementWorkspaceTagvaluesTooManyRequests{}
}

/*PostContentmanagementWorkspaceTagvaluesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostContentmanagementWorkspaceTagvaluesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesInternalServerError creates a PostContentmanagementWorkspaceTagvaluesInternalServerError with default headers values
func NewPostContentmanagementWorkspaceTagvaluesInternalServerError() *PostContentmanagementWorkspaceTagvaluesInternalServerError {
	return &PostContentmanagementWorkspaceTagvaluesInternalServerError{}
}

/*PostContentmanagementWorkspaceTagvaluesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostContentmanagementWorkspaceTagvaluesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesServiceUnavailable creates a PostContentmanagementWorkspaceTagvaluesServiceUnavailable with default headers values
func NewPostContentmanagementWorkspaceTagvaluesServiceUnavailable() *PostContentmanagementWorkspaceTagvaluesServiceUnavailable {
	return &PostContentmanagementWorkspaceTagvaluesServiceUnavailable{}
}

/*PostContentmanagementWorkspaceTagvaluesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostContentmanagementWorkspaceTagvaluesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesGatewayTimeout creates a PostContentmanagementWorkspaceTagvaluesGatewayTimeout with default headers values
func NewPostContentmanagementWorkspaceTagvaluesGatewayTimeout() *PostContentmanagementWorkspaceTagvaluesGatewayTimeout {
	return &PostContentmanagementWorkspaceTagvaluesGatewayTimeout{}
}

/*PostContentmanagementWorkspaceTagvaluesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostContentmanagementWorkspaceTagvaluesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues][%d] postContentmanagementWorkspaceTagvaluesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
