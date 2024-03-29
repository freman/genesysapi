// Code generated by go-swagger; DO NOT EDIT.

package scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScriptsPublishedScriptIDPageReader is a Reader for the GetScriptsPublishedScriptIDPage structure.
type GetScriptsPublishedScriptIDPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsPublishedScriptIDPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsPublishedScriptIDPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsPublishedScriptIDPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsPublishedScriptIDPageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsPublishedScriptIDPageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsPublishedScriptIDPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScriptsPublishedScriptIDPageRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsPublishedScriptIDPageRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsPublishedScriptIDPageUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsPublishedScriptIDPageTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsPublishedScriptIDPageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsPublishedScriptIDPageServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsPublishedScriptIDPageGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsPublishedScriptIDPageOK creates a GetScriptsPublishedScriptIDPageOK with default headers values
func NewGetScriptsPublishedScriptIDPageOK() *GetScriptsPublishedScriptIDPageOK {
	return &GetScriptsPublishedScriptIDPageOK{}
}

/*
GetScriptsPublishedScriptIDPageOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScriptsPublishedScriptIDPageOK struct {
	Payload *models.Page
}

// IsSuccess returns true when this get scripts published script Id page o k response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scripts published script Id page o k response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page o k response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id page o k response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page o k response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScriptsPublishedScriptIDPageOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageOK) GetPayload() *models.Page {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Page)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageBadRequest creates a GetScriptsPublishedScriptIDPageBadRequest with default headers values
func NewGetScriptsPublishedScriptIDPageBadRequest() *GetScriptsPublishedScriptIDPageBadRequest {
	return &GetScriptsPublishedScriptIDPageBadRequest{}
}

/*
GetScriptsPublishedScriptIDPageBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedScriptIDPageBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page bad request response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page bad request response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page bad request response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page bad request response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page bad request response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScriptsPublishedScriptIDPageBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageUnauthorized creates a GetScriptsPublishedScriptIDPageUnauthorized with default headers values
func NewGetScriptsPublishedScriptIDPageUnauthorized() *GetScriptsPublishedScriptIDPageUnauthorized {
	return &GetScriptsPublishedScriptIDPageUnauthorized{}
}

/*
GetScriptsPublishedScriptIDPageUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedScriptIDPageUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page unauthorized response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page unauthorized response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page unauthorized response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page unauthorized response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page unauthorized response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScriptsPublishedScriptIDPageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageForbidden creates a GetScriptsPublishedScriptIDPageForbidden with default headers values
func NewGetScriptsPublishedScriptIDPageForbidden() *GetScriptsPublishedScriptIDPageForbidden {
	return &GetScriptsPublishedScriptIDPageForbidden{}
}

/*
GetScriptsPublishedScriptIDPageForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedScriptIDPageForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page forbidden response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page forbidden response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page forbidden response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page forbidden response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page forbidden response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScriptsPublishedScriptIDPageForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageNotFound creates a GetScriptsPublishedScriptIDPageNotFound with default headers values
func NewGetScriptsPublishedScriptIDPageNotFound() *GetScriptsPublishedScriptIDPageNotFound {
	return &GetScriptsPublishedScriptIDPageNotFound{}
}

/*
GetScriptsPublishedScriptIDPageNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedScriptIDPageNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page not found response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page not found response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page not found response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page not found response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page not found response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScriptsPublishedScriptIDPageNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageRequestTimeout creates a GetScriptsPublishedScriptIDPageRequestTimeout with default headers values
func NewGetScriptsPublishedScriptIDPageRequestTimeout() *GetScriptsPublishedScriptIDPageRequestTimeout {
	return &GetScriptsPublishedScriptIDPageRequestTimeout{}
}

/*
GetScriptsPublishedScriptIDPageRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScriptsPublishedScriptIDPageRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page request timeout response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page request timeout response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page request timeout response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page request timeout response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page request timeout response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScriptsPublishedScriptIDPageRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageRequestEntityTooLarge creates a GetScriptsPublishedScriptIDPageRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedScriptIDPageRequestEntityTooLarge() *GetScriptsPublishedScriptIDPageRequestEntityTooLarge {
	return &GetScriptsPublishedScriptIDPageRequestEntityTooLarge{}
}

/*
GetScriptsPublishedScriptIDPageRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetScriptsPublishedScriptIDPageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page request entity too large response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page request entity too large response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page request entity too large response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page request entity too large response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page request entity too large response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageUnsupportedMediaType creates a GetScriptsPublishedScriptIDPageUnsupportedMediaType with default headers values
func NewGetScriptsPublishedScriptIDPageUnsupportedMediaType() *GetScriptsPublishedScriptIDPageUnsupportedMediaType {
	return &GetScriptsPublishedScriptIDPageUnsupportedMediaType{}
}

/*
GetScriptsPublishedScriptIDPageUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedScriptIDPageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page unsupported media type response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page unsupported media type response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page unsupported media type response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page unsupported media type response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page unsupported media type response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageTooManyRequests creates a GetScriptsPublishedScriptIDPageTooManyRequests with default headers values
func NewGetScriptsPublishedScriptIDPageTooManyRequests() *GetScriptsPublishedScriptIDPageTooManyRequests {
	return &GetScriptsPublishedScriptIDPageTooManyRequests{}
}

/*
GetScriptsPublishedScriptIDPageTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScriptsPublishedScriptIDPageTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page too many requests response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page too many requests response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page too many requests response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scripts published script Id page too many requests response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scripts published script Id page too many requests response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScriptsPublishedScriptIDPageTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageInternalServerError creates a GetScriptsPublishedScriptIDPageInternalServerError with default headers values
func NewGetScriptsPublishedScriptIDPageInternalServerError() *GetScriptsPublishedScriptIDPageInternalServerError {
	return &GetScriptsPublishedScriptIDPageInternalServerError{}
}

/*
GetScriptsPublishedScriptIDPageInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedScriptIDPageInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page internal server error response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page internal server error response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page internal server error response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id page internal server error response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id page internal server error response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScriptsPublishedScriptIDPageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageServiceUnavailable creates a GetScriptsPublishedScriptIDPageServiceUnavailable with default headers values
func NewGetScriptsPublishedScriptIDPageServiceUnavailable() *GetScriptsPublishedScriptIDPageServiceUnavailable {
	return &GetScriptsPublishedScriptIDPageServiceUnavailable{}
}

/*
GetScriptsPublishedScriptIDPageServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedScriptIDPageServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page service unavailable response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page service unavailable response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page service unavailable response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id page service unavailable response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id page service unavailable response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPageGatewayTimeout creates a GetScriptsPublishedScriptIDPageGatewayTimeout with default headers values
func NewGetScriptsPublishedScriptIDPageGatewayTimeout() *GetScriptsPublishedScriptIDPageGatewayTimeout {
	return &GetScriptsPublishedScriptIDPageGatewayTimeout{}
}

/*
GetScriptsPublishedScriptIDPageGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScriptsPublishedScriptIDPageGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scripts published script Id page gateway timeout response has a 2xx status code
func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scripts published script Id page gateway timeout response has a 3xx status code
func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scripts published script Id page gateway timeout response has a 4xx status code
func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scripts published script Id page gateway timeout response has a 5xx status code
func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scripts published script Id page gateway timeout response a status code equal to that given
func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages/{pageId}][%d] getScriptsPublishedScriptIdPageGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
