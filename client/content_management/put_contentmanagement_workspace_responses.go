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
	case 408:
		result := NewPutContentmanagementWorkspaceRequestTimeout()
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

/*
PutContentmanagementWorkspaceOK describes a response with status code 200, with default header values.

successful operation
*/
type PutContentmanagementWorkspaceOK struct {
	Payload *models.Workspace
}

// IsSuccess returns true when this put contentmanagement workspace o k response has a 2xx status code
func (o *PutContentmanagementWorkspaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put contentmanagement workspace o k response has a 3xx status code
func (o *PutContentmanagementWorkspaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace o k response has a 4xx status code
func (o *PutContentmanagementWorkspaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contentmanagement workspace o k response has a 5xx status code
func (o *PutContentmanagementWorkspaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace o k response a status code equal to that given
func (o *PutContentmanagementWorkspaceOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutContentmanagementWorkspaceOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceOK  %+v", 200, o.Payload)
}

func (o *PutContentmanagementWorkspaceOK) String() string {
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

/*
PutContentmanagementWorkspaceBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutContentmanagementWorkspaceBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace bad request response has a 2xx status code
func (o *PutContentmanagementWorkspaceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace bad request response has a 3xx status code
func (o *PutContentmanagementWorkspaceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace bad request response has a 4xx status code
func (o *PutContentmanagementWorkspaceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace bad request response has a 5xx status code
func (o *PutContentmanagementWorkspaceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace bad request response a status code equal to that given
func (o *PutContentmanagementWorkspaceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutContentmanagementWorkspaceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceBadRequest  %+v", 400, o.Payload)
}

func (o *PutContentmanagementWorkspaceBadRequest) String() string {
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

/*
PutContentmanagementWorkspaceUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutContentmanagementWorkspaceUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace unauthorized response has a 2xx status code
func (o *PutContentmanagementWorkspaceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace unauthorized response has a 3xx status code
func (o *PutContentmanagementWorkspaceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace unauthorized response has a 4xx status code
func (o *PutContentmanagementWorkspaceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace unauthorized response has a 5xx status code
func (o *PutContentmanagementWorkspaceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace unauthorized response a status code equal to that given
func (o *PutContentmanagementWorkspaceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutContentmanagementWorkspaceUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceUnauthorized  %+v", 401, o.Payload)
}

func (o *PutContentmanagementWorkspaceUnauthorized) String() string {
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

/*
PutContentmanagementWorkspaceForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutContentmanagementWorkspaceForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace forbidden response has a 2xx status code
func (o *PutContentmanagementWorkspaceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace forbidden response has a 3xx status code
func (o *PutContentmanagementWorkspaceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace forbidden response has a 4xx status code
func (o *PutContentmanagementWorkspaceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace forbidden response has a 5xx status code
func (o *PutContentmanagementWorkspaceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace forbidden response a status code equal to that given
func (o *PutContentmanagementWorkspaceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutContentmanagementWorkspaceForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceForbidden  %+v", 403, o.Payload)
}

func (o *PutContentmanagementWorkspaceForbidden) String() string {
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

/*
PutContentmanagementWorkspaceNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutContentmanagementWorkspaceNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace not found response has a 2xx status code
func (o *PutContentmanagementWorkspaceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace not found response has a 3xx status code
func (o *PutContentmanagementWorkspaceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace not found response has a 4xx status code
func (o *PutContentmanagementWorkspaceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace not found response has a 5xx status code
func (o *PutContentmanagementWorkspaceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace not found response a status code equal to that given
func (o *PutContentmanagementWorkspaceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutContentmanagementWorkspaceNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceNotFound  %+v", 404, o.Payload)
}

func (o *PutContentmanagementWorkspaceNotFound) String() string {
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

// NewPutContentmanagementWorkspaceRequestTimeout creates a PutContentmanagementWorkspaceRequestTimeout with default headers values
func NewPutContentmanagementWorkspaceRequestTimeout() *PutContentmanagementWorkspaceRequestTimeout {
	return &PutContentmanagementWorkspaceRequestTimeout{}
}

/*
PutContentmanagementWorkspaceRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutContentmanagementWorkspaceRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace request timeout response has a 2xx status code
func (o *PutContentmanagementWorkspaceRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace request timeout response has a 3xx status code
func (o *PutContentmanagementWorkspaceRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace request timeout response has a 4xx status code
func (o *PutContentmanagementWorkspaceRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace request timeout response has a 5xx status code
func (o *PutContentmanagementWorkspaceRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace request timeout response a status code equal to that given
func (o *PutContentmanagementWorkspaceRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutContentmanagementWorkspaceRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutContentmanagementWorkspaceRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutContentmanagementWorkspaceRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutContentmanagementWorkspaceRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
PutContentmanagementWorkspaceRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PutContentmanagementWorkspaceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace request entity too large response has a 2xx status code
func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace request entity too large response has a 3xx status code
func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace request entity too large response has a 4xx status code
func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace request entity too large response has a 5xx status code
func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace request entity too large response a status code equal to that given
func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutContentmanagementWorkspaceRequestEntityTooLarge) String() string {
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

/*
PutContentmanagementWorkspaceUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutContentmanagementWorkspaceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace unsupported media type response has a 2xx status code
func (o *PutContentmanagementWorkspaceUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace unsupported media type response has a 3xx status code
func (o *PutContentmanagementWorkspaceUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace unsupported media type response has a 4xx status code
func (o *PutContentmanagementWorkspaceUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace unsupported media type response has a 5xx status code
func (o *PutContentmanagementWorkspaceUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace unsupported media type response a status code equal to that given
func (o *PutContentmanagementWorkspaceUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutContentmanagementWorkspaceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutContentmanagementWorkspaceUnsupportedMediaType) String() string {
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

/*
PutContentmanagementWorkspaceTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutContentmanagementWorkspaceTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace too many requests response has a 2xx status code
func (o *PutContentmanagementWorkspaceTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace too many requests response has a 3xx status code
func (o *PutContentmanagementWorkspaceTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace too many requests response has a 4xx status code
func (o *PutContentmanagementWorkspaceTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put contentmanagement workspace too many requests response has a 5xx status code
func (o *PutContentmanagementWorkspaceTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put contentmanagement workspace too many requests response a status code equal to that given
func (o *PutContentmanagementWorkspaceTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutContentmanagementWorkspaceTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutContentmanagementWorkspaceTooManyRequests) String() string {
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

/*
PutContentmanagementWorkspaceInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutContentmanagementWorkspaceInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace internal server error response has a 2xx status code
func (o *PutContentmanagementWorkspaceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace internal server error response has a 3xx status code
func (o *PutContentmanagementWorkspaceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace internal server error response has a 4xx status code
func (o *PutContentmanagementWorkspaceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contentmanagement workspace internal server error response has a 5xx status code
func (o *PutContentmanagementWorkspaceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put contentmanagement workspace internal server error response a status code equal to that given
func (o *PutContentmanagementWorkspaceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutContentmanagementWorkspaceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceInternalServerError  %+v", 500, o.Payload)
}

func (o *PutContentmanagementWorkspaceInternalServerError) String() string {
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

/*
PutContentmanagementWorkspaceServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutContentmanagementWorkspaceServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace service unavailable response has a 2xx status code
func (o *PutContentmanagementWorkspaceServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace service unavailable response has a 3xx status code
func (o *PutContentmanagementWorkspaceServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace service unavailable response has a 4xx status code
func (o *PutContentmanagementWorkspaceServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contentmanagement workspace service unavailable response has a 5xx status code
func (o *PutContentmanagementWorkspaceServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put contentmanagement workspace service unavailable response a status code equal to that given
func (o *PutContentmanagementWorkspaceServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutContentmanagementWorkspaceServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutContentmanagementWorkspaceServiceUnavailable) String() string {
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

/*
PutContentmanagementWorkspaceGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutContentmanagementWorkspaceGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put contentmanagement workspace gateway timeout response has a 2xx status code
func (o *PutContentmanagementWorkspaceGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put contentmanagement workspace gateway timeout response has a 3xx status code
func (o *PutContentmanagementWorkspaceGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put contentmanagement workspace gateway timeout response has a 4xx status code
func (o *PutContentmanagementWorkspaceGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put contentmanagement workspace gateway timeout response has a 5xx status code
func (o *PutContentmanagementWorkspaceGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put contentmanagement workspace gateway timeout response a status code equal to that given
func (o *PutContentmanagementWorkspaceGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutContentmanagementWorkspaceGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/contentmanagement/workspaces/{workspaceId}][%d] putContentmanagementWorkspaceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutContentmanagementWorkspaceGatewayTimeout) String() string {
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
