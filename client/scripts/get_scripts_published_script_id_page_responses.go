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

/*GetScriptsPublishedScriptIDPageOK handles this case with default header values.

successful operation
*/
type GetScriptsPublishedScriptIDPageOK struct {
	Payload *models.Page
}

func (o *GetScriptsPublishedScriptIDPageOK) Error() string {
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

/*GetScriptsPublishedScriptIDPageBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedScriptIDPageBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageBadRequest) Error() string {
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

/*GetScriptsPublishedScriptIDPageUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedScriptIDPageUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageUnauthorized) Error() string {
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

/*GetScriptsPublishedScriptIDPageForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedScriptIDPageForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageForbidden) Error() string {
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

/*GetScriptsPublishedScriptIDPageNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedScriptIDPageNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageNotFound) Error() string {
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

// NewGetScriptsPublishedScriptIDPageRequestEntityTooLarge creates a GetScriptsPublishedScriptIDPageRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedScriptIDPageRequestEntityTooLarge() *GetScriptsPublishedScriptIDPageRequestEntityTooLarge {
	return &GetScriptsPublishedScriptIDPageRequestEntityTooLarge{}
}

/*GetScriptsPublishedScriptIDPageRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScriptsPublishedScriptIDPageRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageRequestEntityTooLarge) Error() string {
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

/*GetScriptsPublishedScriptIDPageUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedScriptIDPageUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageUnsupportedMediaType) Error() string {
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

/*GetScriptsPublishedScriptIDPageTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScriptsPublishedScriptIDPageTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageTooManyRequests) Error() string {
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

/*GetScriptsPublishedScriptIDPageInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedScriptIDPageInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageInternalServerError) Error() string {
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

/*GetScriptsPublishedScriptIDPageServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedScriptIDPageServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageServiceUnavailable) Error() string {
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

/*GetScriptsPublishedScriptIDPageGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScriptsPublishedScriptIDPageGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPageGatewayTimeout) Error() string {
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