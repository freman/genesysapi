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

// GetScriptsPublishedScriptIDPagesReader is a Reader for the GetScriptsPublishedScriptIDPages structure.
type GetScriptsPublishedScriptIDPagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsPublishedScriptIDPagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsPublishedScriptIDPagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsPublishedScriptIDPagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsPublishedScriptIDPagesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsPublishedScriptIDPagesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsPublishedScriptIDPagesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsPublishedScriptIDPagesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsPublishedScriptIDPagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsPublishedScriptIDPagesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsPublishedScriptIDPagesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsPublishedScriptIDPagesOK creates a GetScriptsPublishedScriptIDPagesOK with default headers values
func NewGetScriptsPublishedScriptIDPagesOK() *GetScriptsPublishedScriptIDPagesOK {
	return &GetScriptsPublishedScriptIDPagesOK{}
}

/*GetScriptsPublishedScriptIDPagesOK handles this case with default header values.

successful operation
*/
type GetScriptsPublishedScriptIDPagesOK struct {
	Payload []*models.Page
}

func (o *GetScriptsPublishedScriptIDPagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesOK  %+v", 200, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesOK) GetPayload() []*models.Page {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesBadRequest creates a GetScriptsPublishedScriptIDPagesBadRequest with default headers values
func NewGetScriptsPublishedScriptIDPagesBadRequest() *GetScriptsPublishedScriptIDPagesBadRequest {
	return &GetScriptsPublishedScriptIDPagesBadRequest{}
}

/*GetScriptsPublishedScriptIDPagesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsPublishedScriptIDPagesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesUnauthorized creates a GetScriptsPublishedScriptIDPagesUnauthorized with default headers values
func NewGetScriptsPublishedScriptIDPagesUnauthorized() *GetScriptsPublishedScriptIDPagesUnauthorized {
	return &GetScriptsPublishedScriptIDPagesUnauthorized{}
}

/*GetScriptsPublishedScriptIDPagesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsPublishedScriptIDPagesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesForbidden creates a GetScriptsPublishedScriptIDPagesForbidden with default headers values
func NewGetScriptsPublishedScriptIDPagesForbidden() *GetScriptsPublishedScriptIDPagesForbidden {
	return &GetScriptsPublishedScriptIDPagesForbidden{}
}

/*GetScriptsPublishedScriptIDPagesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsPublishedScriptIDPagesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesNotFound creates a GetScriptsPublishedScriptIDPagesNotFound with default headers values
func NewGetScriptsPublishedScriptIDPagesNotFound() *GetScriptsPublishedScriptIDPagesNotFound {
	return &GetScriptsPublishedScriptIDPagesNotFound{}
}

/*GetScriptsPublishedScriptIDPagesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScriptsPublishedScriptIDPagesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge creates a GetScriptsPublishedScriptIDPagesRequestEntityTooLarge with default headers values
func NewGetScriptsPublishedScriptIDPagesRequestEntityTooLarge() *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge {
	return &GetScriptsPublishedScriptIDPagesRequestEntityTooLarge{}
}

/*GetScriptsPublishedScriptIDPagesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScriptsPublishedScriptIDPagesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType creates a GetScriptsPublishedScriptIDPagesUnsupportedMediaType with default headers values
func NewGetScriptsPublishedScriptIDPagesUnsupportedMediaType() *GetScriptsPublishedScriptIDPagesUnsupportedMediaType {
	return &GetScriptsPublishedScriptIDPagesUnsupportedMediaType{}
}

/*GetScriptsPublishedScriptIDPagesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsPublishedScriptIDPagesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesTooManyRequests creates a GetScriptsPublishedScriptIDPagesTooManyRequests with default headers values
func NewGetScriptsPublishedScriptIDPagesTooManyRequests() *GetScriptsPublishedScriptIDPagesTooManyRequests {
	return &GetScriptsPublishedScriptIDPagesTooManyRequests{}
}

/*GetScriptsPublishedScriptIDPagesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScriptsPublishedScriptIDPagesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesInternalServerError creates a GetScriptsPublishedScriptIDPagesInternalServerError with default headers values
func NewGetScriptsPublishedScriptIDPagesInternalServerError() *GetScriptsPublishedScriptIDPagesInternalServerError {
	return &GetScriptsPublishedScriptIDPagesInternalServerError{}
}

/*GetScriptsPublishedScriptIDPagesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsPublishedScriptIDPagesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesServiceUnavailable creates a GetScriptsPublishedScriptIDPagesServiceUnavailable with default headers values
func NewGetScriptsPublishedScriptIDPagesServiceUnavailable() *GetScriptsPublishedScriptIDPagesServiceUnavailable {
	return &GetScriptsPublishedScriptIDPagesServiceUnavailable{}
}

/*GetScriptsPublishedScriptIDPagesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsPublishedScriptIDPagesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsPublishedScriptIDPagesGatewayTimeout creates a GetScriptsPublishedScriptIDPagesGatewayTimeout with default headers values
func NewGetScriptsPublishedScriptIDPagesGatewayTimeout() *GetScriptsPublishedScriptIDPagesGatewayTimeout {
	return &GetScriptsPublishedScriptIDPagesGatewayTimeout{}
}

/*GetScriptsPublishedScriptIDPagesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScriptsPublishedScriptIDPagesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts/published/{scriptId}/pages][%d] getScriptsPublishedScriptIdPagesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsPublishedScriptIDPagesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
