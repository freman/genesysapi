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

// GetContentmanagementWorkspacesReader is a Reader for the GetContentmanagementWorkspaces structure.
type GetContentmanagementWorkspacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementWorkspacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementWorkspacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementWorkspacesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementWorkspacesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementWorkspacesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementWorkspacesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementWorkspacesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementWorkspacesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementWorkspacesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementWorkspacesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementWorkspacesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementWorkspacesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementWorkspacesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementWorkspacesOK creates a GetContentmanagementWorkspacesOK with default headers values
func NewGetContentmanagementWorkspacesOK() *GetContentmanagementWorkspacesOK {
	return &GetContentmanagementWorkspacesOK{}
}

/*
GetContentmanagementWorkspacesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetContentmanagementWorkspacesOK struct {
	Payload *models.WorkspaceEntityListing
}

// IsSuccess returns true when this get contentmanagement workspaces o k response has a 2xx status code
func (o *GetContentmanagementWorkspacesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement workspaces o k response has a 3xx status code
func (o *GetContentmanagementWorkspacesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces o k response has a 4xx status code
func (o *GetContentmanagementWorkspacesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspaces o k response has a 5xx status code
func (o *GetContentmanagementWorkspacesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces o k response a status code equal to that given
func (o *GetContentmanagementWorkspacesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContentmanagementWorkspacesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementWorkspacesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementWorkspacesOK) GetPayload() *models.WorkspaceEntityListing {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkspaceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesBadRequest creates a GetContentmanagementWorkspacesBadRequest with default headers values
func NewGetContentmanagementWorkspacesBadRequest() *GetContentmanagementWorkspacesBadRequest {
	return &GetContentmanagementWorkspacesBadRequest{}
}

/*
GetContentmanagementWorkspacesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementWorkspacesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces bad request response has a 2xx status code
func (o *GetContentmanagementWorkspacesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces bad request response has a 3xx status code
func (o *GetContentmanagementWorkspacesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces bad request response has a 4xx status code
func (o *GetContentmanagementWorkspacesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces bad request response has a 5xx status code
func (o *GetContentmanagementWorkspacesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces bad request response a status code equal to that given
func (o *GetContentmanagementWorkspacesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContentmanagementWorkspacesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementWorkspacesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementWorkspacesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesUnauthorized creates a GetContentmanagementWorkspacesUnauthorized with default headers values
func NewGetContentmanagementWorkspacesUnauthorized() *GetContentmanagementWorkspacesUnauthorized {
	return &GetContentmanagementWorkspacesUnauthorized{}
}

/*
GetContentmanagementWorkspacesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementWorkspacesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces unauthorized response has a 2xx status code
func (o *GetContentmanagementWorkspacesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces unauthorized response has a 3xx status code
func (o *GetContentmanagementWorkspacesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces unauthorized response has a 4xx status code
func (o *GetContentmanagementWorkspacesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces unauthorized response has a 5xx status code
func (o *GetContentmanagementWorkspacesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces unauthorized response a status code equal to that given
func (o *GetContentmanagementWorkspacesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetContentmanagementWorkspacesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementWorkspacesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementWorkspacesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesForbidden creates a GetContentmanagementWorkspacesForbidden with default headers values
func NewGetContentmanagementWorkspacesForbidden() *GetContentmanagementWorkspacesForbidden {
	return &GetContentmanagementWorkspacesForbidden{}
}

/*
GetContentmanagementWorkspacesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementWorkspacesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces forbidden response has a 2xx status code
func (o *GetContentmanagementWorkspacesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces forbidden response has a 3xx status code
func (o *GetContentmanagementWorkspacesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces forbidden response has a 4xx status code
func (o *GetContentmanagementWorkspacesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces forbidden response has a 5xx status code
func (o *GetContentmanagementWorkspacesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces forbidden response a status code equal to that given
func (o *GetContentmanagementWorkspacesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetContentmanagementWorkspacesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementWorkspacesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementWorkspacesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesNotFound creates a GetContentmanagementWorkspacesNotFound with default headers values
func NewGetContentmanagementWorkspacesNotFound() *GetContentmanagementWorkspacesNotFound {
	return &GetContentmanagementWorkspacesNotFound{}
}

/*
GetContentmanagementWorkspacesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetContentmanagementWorkspacesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces not found response has a 2xx status code
func (o *GetContentmanagementWorkspacesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces not found response has a 3xx status code
func (o *GetContentmanagementWorkspacesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces not found response has a 4xx status code
func (o *GetContentmanagementWorkspacesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces not found response has a 5xx status code
func (o *GetContentmanagementWorkspacesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces not found response a status code equal to that given
func (o *GetContentmanagementWorkspacesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetContentmanagementWorkspacesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementWorkspacesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementWorkspacesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesRequestTimeout creates a GetContentmanagementWorkspacesRequestTimeout with default headers values
func NewGetContentmanagementWorkspacesRequestTimeout() *GetContentmanagementWorkspacesRequestTimeout {
	return &GetContentmanagementWorkspacesRequestTimeout{}
}

/*
GetContentmanagementWorkspacesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementWorkspacesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces request timeout response has a 2xx status code
func (o *GetContentmanagementWorkspacesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces request timeout response has a 3xx status code
func (o *GetContentmanagementWorkspacesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces request timeout response has a 4xx status code
func (o *GetContentmanagementWorkspacesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces request timeout response has a 5xx status code
func (o *GetContentmanagementWorkspacesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces request timeout response a status code equal to that given
func (o *GetContentmanagementWorkspacesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetContentmanagementWorkspacesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementWorkspacesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementWorkspacesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesRequestEntityTooLarge creates a GetContentmanagementWorkspacesRequestEntityTooLarge with default headers values
func NewGetContentmanagementWorkspacesRequestEntityTooLarge() *GetContentmanagementWorkspacesRequestEntityTooLarge {
	return &GetContentmanagementWorkspacesRequestEntityTooLarge{}
}

/*
GetContentmanagementWorkspacesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetContentmanagementWorkspacesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces request entity too large response has a 2xx status code
func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces request entity too large response has a 3xx status code
func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces request entity too large response has a 4xx status code
func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces request entity too large response has a 5xx status code
func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces request entity too large response a status code equal to that given
func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesUnsupportedMediaType creates a GetContentmanagementWorkspacesUnsupportedMediaType with default headers values
func NewGetContentmanagementWorkspacesUnsupportedMediaType() *GetContentmanagementWorkspacesUnsupportedMediaType {
	return &GetContentmanagementWorkspacesUnsupportedMediaType{}
}

/*
GetContentmanagementWorkspacesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementWorkspacesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces unsupported media type response has a 2xx status code
func (o *GetContentmanagementWorkspacesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces unsupported media type response has a 3xx status code
func (o *GetContentmanagementWorkspacesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces unsupported media type response has a 4xx status code
func (o *GetContentmanagementWorkspacesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces unsupported media type response has a 5xx status code
func (o *GetContentmanagementWorkspacesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces unsupported media type response a status code equal to that given
func (o *GetContentmanagementWorkspacesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetContentmanagementWorkspacesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementWorkspacesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementWorkspacesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesTooManyRequests creates a GetContentmanagementWorkspacesTooManyRequests with default headers values
func NewGetContentmanagementWorkspacesTooManyRequests() *GetContentmanagementWorkspacesTooManyRequests {
	return &GetContentmanagementWorkspacesTooManyRequests{}
}

/*
GetContentmanagementWorkspacesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementWorkspacesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces too many requests response has a 2xx status code
func (o *GetContentmanagementWorkspacesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces too many requests response has a 3xx status code
func (o *GetContentmanagementWorkspacesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces too many requests response has a 4xx status code
func (o *GetContentmanagementWorkspacesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspaces too many requests response has a 5xx status code
func (o *GetContentmanagementWorkspacesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspaces too many requests response a status code equal to that given
func (o *GetContentmanagementWorkspacesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetContentmanagementWorkspacesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementWorkspacesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementWorkspacesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesInternalServerError creates a GetContentmanagementWorkspacesInternalServerError with default headers values
func NewGetContentmanagementWorkspacesInternalServerError() *GetContentmanagementWorkspacesInternalServerError {
	return &GetContentmanagementWorkspacesInternalServerError{}
}

/*
GetContentmanagementWorkspacesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementWorkspacesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces internal server error response has a 2xx status code
func (o *GetContentmanagementWorkspacesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces internal server error response has a 3xx status code
func (o *GetContentmanagementWorkspacesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces internal server error response has a 4xx status code
func (o *GetContentmanagementWorkspacesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspaces internal server error response has a 5xx status code
func (o *GetContentmanagementWorkspacesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspaces internal server error response a status code equal to that given
func (o *GetContentmanagementWorkspacesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetContentmanagementWorkspacesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementWorkspacesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementWorkspacesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesServiceUnavailable creates a GetContentmanagementWorkspacesServiceUnavailable with default headers values
func NewGetContentmanagementWorkspacesServiceUnavailable() *GetContentmanagementWorkspacesServiceUnavailable {
	return &GetContentmanagementWorkspacesServiceUnavailable{}
}

/*
GetContentmanagementWorkspacesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementWorkspacesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces service unavailable response has a 2xx status code
func (o *GetContentmanagementWorkspacesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces service unavailable response has a 3xx status code
func (o *GetContentmanagementWorkspacesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces service unavailable response has a 4xx status code
func (o *GetContentmanagementWorkspacesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspaces service unavailable response has a 5xx status code
func (o *GetContentmanagementWorkspacesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspaces service unavailable response a status code equal to that given
func (o *GetContentmanagementWorkspacesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetContentmanagementWorkspacesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementWorkspacesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementWorkspacesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspacesGatewayTimeout creates a GetContentmanagementWorkspacesGatewayTimeout with default headers values
func NewGetContentmanagementWorkspacesGatewayTimeout() *GetContentmanagementWorkspacesGatewayTimeout {
	return &GetContentmanagementWorkspacesGatewayTimeout{}
}

/*
GetContentmanagementWorkspacesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetContentmanagementWorkspacesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspaces gateway timeout response has a 2xx status code
func (o *GetContentmanagementWorkspacesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspaces gateway timeout response has a 3xx status code
func (o *GetContentmanagementWorkspacesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspaces gateway timeout response has a 4xx status code
func (o *GetContentmanagementWorkspacesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspaces gateway timeout response has a 5xx status code
func (o *GetContentmanagementWorkspacesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspaces gateway timeout response a status code equal to that given
func (o *GetContentmanagementWorkspacesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetContentmanagementWorkspacesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementWorkspacesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces][%d] getContentmanagementWorkspacesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementWorkspacesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspacesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
