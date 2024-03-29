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

// GetContentmanagementWorkspaceDocumentsReader is a Reader for the GetContentmanagementWorkspaceDocuments structure.
type GetContentmanagementWorkspaceDocumentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentmanagementWorkspaceDocumentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetContentmanagementWorkspaceDocumentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetContentmanagementWorkspaceDocumentsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetContentmanagementWorkspaceDocumentsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetContentmanagementWorkspaceDocumentsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetContentmanagementWorkspaceDocumentsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetContentmanagementWorkspaceDocumentsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetContentmanagementWorkspaceDocumentsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetContentmanagementWorkspaceDocumentsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetContentmanagementWorkspaceDocumentsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetContentmanagementWorkspaceDocumentsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetContentmanagementWorkspaceDocumentsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetContentmanagementWorkspaceDocumentsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetContentmanagementWorkspaceDocumentsOK creates a GetContentmanagementWorkspaceDocumentsOK with default headers values
func NewGetContentmanagementWorkspaceDocumentsOK() *GetContentmanagementWorkspaceDocumentsOK {
	return &GetContentmanagementWorkspaceDocumentsOK{}
}

/*
GetContentmanagementWorkspaceDocumentsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetContentmanagementWorkspaceDocumentsOK struct {
	Payload *models.DocumentEntityListing
}

// IsSuccess returns true when this get contentmanagement workspace documents o k response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get contentmanagement workspace documents o k response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents o k response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspace documents o k response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents o k response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetContentmanagementWorkspaceDocumentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsOK  %+v", 200, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsOK) GetPayload() *models.DocumentEntityListing {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DocumentEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsBadRequest creates a GetContentmanagementWorkspaceDocumentsBadRequest with default headers values
func NewGetContentmanagementWorkspaceDocumentsBadRequest() *GetContentmanagementWorkspaceDocumentsBadRequest {
	return &GetContentmanagementWorkspaceDocumentsBadRequest{}
}

/*
GetContentmanagementWorkspaceDocumentsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetContentmanagementWorkspaceDocumentsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents bad request response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents bad request response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents bad request response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents bad request response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents bad request response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetContentmanagementWorkspaceDocumentsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsBadRequest  %+v", 400, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsUnauthorized creates a GetContentmanagementWorkspaceDocumentsUnauthorized with default headers values
func NewGetContentmanagementWorkspaceDocumentsUnauthorized() *GetContentmanagementWorkspaceDocumentsUnauthorized {
	return &GetContentmanagementWorkspaceDocumentsUnauthorized{}
}

/*
GetContentmanagementWorkspaceDocumentsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetContentmanagementWorkspaceDocumentsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents unauthorized response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents unauthorized response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents unauthorized response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents unauthorized response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents unauthorized response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsForbidden creates a GetContentmanagementWorkspaceDocumentsForbidden with default headers values
func NewGetContentmanagementWorkspaceDocumentsForbidden() *GetContentmanagementWorkspaceDocumentsForbidden {
	return &GetContentmanagementWorkspaceDocumentsForbidden{}
}

/*
GetContentmanagementWorkspaceDocumentsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetContentmanagementWorkspaceDocumentsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents forbidden response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents forbidden response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents forbidden response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents forbidden response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents forbidden response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetContentmanagementWorkspaceDocumentsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsForbidden  %+v", 403, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsNotFound creates a GetContentmanagementWorkspaceDocumentsNotFound with default headers values
func NewGetContentmanagementWorkspaceDocumentsNotFound() *GetContentmanagementWorkspaceDocumentsNotFound {
	return &GetContentmanagementWorkspaceDocumentsNotFound{}
}

/*
GetContentmanagementWorkspaceDocumentsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetContentmanagementWorkspaceDocumentsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents not found response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents not found response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents not found response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents not found response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents not found response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetContentmanagementWorkspaceDocumentsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsNotFound  %+v", 404, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsRequestTimeout creates a GetContentmanagementWorkspaceDocumentsRequestTimeout with default headers values
func NewGetContentmanagementWorkspaceDocumentsRequestTimeout() *GetContentmanagementWorkspaceDocumentsRequestTimeout {
	return &GetContentmanagementWorkspaceDocumentsRequestTimeout{}
}

/*
GetContentmanagementWorkspaceDocumentsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetContentmanagementWorkspaceDocumentsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents request timeout response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents request timeout response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents request timeout response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents request timeout response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents request timeout response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsRequestEntityTooLarge creates a GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge with default headers values
func NewGetContentmanagementWorkspaceDocumentsRequestEntityTooLarge() *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge {
	return &GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge{}
}

/*
GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents request entity too large response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents request entity too large response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents request entity too large response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents request entity too large response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents request entity too large response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsUnsupportedMediaType creates a GetContentmanagementWorkspaceDocumentsUnsupportedMediaType with default headers values
func NewGetContentmanagementWorkspaceDocumentsUnsupportedMediaType() *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType {
	return &GetContentmanagementWorkspaceDocumentsUnsupportedMediaType{}
}

/*
GetContentmanagementWorkspaceDocumentsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetContentmanagementWorkspaceDocumentsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents unsupported media type response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents unsupported media type response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents unsupported media type response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents unsupported media type response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents unsupported media type response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsTooManyRequests creates a GetContentmanagementWorkspaceDocumentsTooManyRequests with default headers values
func NewGetContentmanagementWorkspaceDocumentsTooManyRequests() *GetContentmanagementWorkspaceDocumentsTooManyRequests {
	return &GetContentmanagementWorkspaceDocumentsTooManyRequests{}
}

/*
GetContentmanagementWorkspaceDocumentsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetContentmanagementWorkspaceDocumentsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents too many requests response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents too many requests response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents too many requests response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get contentmanagement workspace documents too many requests response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get contentmanagement workspace documents too many requests response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsInternalServerError creates a GetContentmanagementWorkspaceDocumentsInternalServerError with default headers values
func NewGetContentmanagementWorkspaceDocumentsInternalServerError() *GetContentmanagementWorkspaceDocumentsInternalServerError {
	return &GetContentmanagementWorkspaceDocumentsInternalServerError{}
}

/*
GetContentmanagementWorkspaceDocumentsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetContentmanagementWorkspaceDocumentsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents internal server error response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents internal server error response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents internal server error response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspace documents internal server error response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspace documents internal server error response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsServiceUnavailable creates a GetContentmanagementWorkspaceDocumentsServiceUnavailable with default headers values
func NewGetContentmanagementWorkspaceDocumentsServiceUnavailable() *GetContentmanagementWorkspaceDocumentsServiceUnavailable {
	return &GetContentmanagementWorkspaceDocumentsServiceUnavailable{}
}

/*
GetContentmanagementWorkspaceDocumentsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetContentmanagementWorkspaceDocumentsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents service unavailable response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents service unavailable response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents service unavailable response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspace documents service unavailable response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspace documents service unavailable response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContentmanagementWorkspaceDocumentsGatewayTimeout creates a GetContentmanagementWorkspaceDocumentsGatewayTimeout with default headers values
func NewGetContentmanagementWorkspaceDocumentsGatewayTimeout() *GetContentmanagementWorkspaceDocumentsGatewayTimeout {
	return &GetContentmanagementWorkspaceDocumentsGatewayTimeout{}
}

/*
GetContentmanagementWorkspaceDocumentsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetContentmanagementWorkspaceDocumentsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get contentmanagement workspace documents gateway timeout response has a 2xx status code
func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get contentmanagement workspace documents gateway timeout response has a 3xx status code
func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get contentmanagement workspace documents gateway timeout response has a 4xx status code
func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get contentmanagement workspace documents gateway timeout response has a 5xx status code
func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get contentmanagement workspace documents gateway timeout response a status code equal to that given
func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/contentmanagement/workspaces/{workspaceId}/documents][%d] getContentmanagementWorkspaceDocumentsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetContentmanagementWorkspaceDocumentsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
