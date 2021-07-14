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

// PostContentmanagementWorkspaceTagvaluesQueryReader is a Reader for the PostContentmanagementWorkspaceTagvaluesQuery structure.
type PostContentmanagementWorkspaceTagvaluesQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostContentmanagementWorkspaceTagvaluesQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostContentmanagementWorkspaceTagvaluesQueryOK creates a PostContentmanagementWorkspaceTagvaluesQueryOK with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryOK() *PostContentmanagementWorkspaceTagvaluesQueryOK {
	return &PostContentmanagementWorkspaceTagvaluesQueryOK{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryOK handles this case with default header values.

successful operation
*/
type PostContentmanagementWorkspaceTagvaluesQueryOK struct {
	Payload *models.TagValueEntityListing
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryOK  %+v", 200, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryOK) GetPayload() *models.TagValueEntityListing {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagValueEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryBadRequest creates a PostContentmanagementWorkspaceTagvaluesQueryBadRequest with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryBadRequest() *PostContentmanagementWorkspaceTagvaluesQueryBadRequest {
	return &PostContentmanagementWorkspaceTagvaluesQueryBadRequest{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostContentmanagementWorkspaceTagvaluesQueryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryBadRequest  %+v", 400, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryUnauthorized creates a PostContentmanagementWorkspaceTagvaluesQueryUnauthorized with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryUnauthorized() *PostContentmanagementWorkspaceTagvaluesQueryUnauthorized {
	return &PostContentmanagementWorkspaceTagvaluesQueryUnauthorized{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostContentmanagementWorkspaceTagvaluesQueryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryUnauthorized  %+v", 401, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryForbidden creates a PostContentmanagementWorkspaceTagvaluesQueryForbidden with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryForbidden() *PostContentmanagementWorkspaceTagvaluesQueryForbidden {
	return &PostContentmanagementWorkspaceTagvaluesQueryForbidden{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostContentmanagementWorkspaceTagvaluesQueryForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryForbidden  %+v", 403, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryNotFound creates a PostContentmanagementWorkspaceTagvaluesQueryNotFound with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryNotFound() *PostContentmanagementWorkspaceTagvaluesQueryNotFound {
	return &PostContentmanagementWorkspaceTagvaluesQueryNotFound{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostContentmanagementWorkspaceTagvaluesQueryNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryNotFound  %+v", 404, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryRequestTimeout creates a PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryRequestTimeout() *PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout {
	return &PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge creates a PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge() *PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge {
	return &PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType creates a PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType() *PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType {
	return &PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryTooManyRequests creates a PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryTooManyRequests() *PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests {
	return &PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryInternalServerError creates a PostContentmanagementWorkspaceTagvaluesQueryInternalServerError with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryInternalServerError() *PostContentmanagementWorkspaceTagvaluesQueryInternalServerError {
	return &PostContentmanagementWorkspaceTagvaluesQueryInternalServerError{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostContentmanagementWorkspaceTagvaluesQueryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryInternalServerError  %+v", 500, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable creates a PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable() *PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable {
	return &PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout creates a PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout with default headers values
func NewPostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout() *PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout {
	return &PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout{}
}

/*PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/contentmanagement/workspaces/{workspaceId}/tagvalues/query][%d] postContentmanagementWorkspaceTagvaluesQueryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostContentmanagementWorkspaceTagvaluesQueryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
