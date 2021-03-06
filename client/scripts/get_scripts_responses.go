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

// GetScriptsReader is a Reader for the GetScripts structure.
type GetScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScriptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScriptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScriptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScriptsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScriptsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScriptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScriptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScriptsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScriptsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScriptsOK creates a GetScriptsOK with default headers values
func NewGetScriptsOK() *GetScriptsOK {
	return &GetScriptsOK{}
}

/*GetScriptsOK handles this case with default header values.

successful operation
*/
type GetScriptsOK struct {
	Payload *models.ScriptEntityListing
}

func (o *GetScriptsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsOK  %+v", 200, o.Payload)
}

func (o *GetScriptsOK) GetPayload() *models.ScriptEntityListing {
	return o.Payload
}

func (o *GetScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScriptEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsBadRequest creates a GetScriptsBadRequest with default headers values
func NewGetScriptsBadRequest() *GetScriptsBadRequest {
	return &GetScriptsBadRequest{}
}

/*GetScriptsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScriptsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetScriptsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsUnauthorized creates a GetScriptsUnauthorized with default headers values
func NewGetScriptsUnauthorized() *GetScriptsUnauthorized {
	return &GetScriptsUnauthorized{}
}

/*GetScriptsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScriptsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScriptsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsForbidden creates a GetScriptsForbidden with default headers values
func NewGetScriptsForbidden() *GetScriptsForbidden {
	return &GetScriptsForbidden{}
}

/*GetScriptsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScriptsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsForbidden  %+v", 403, o.Payload)
}

func (o *GetScriptsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsNotFound creates a GetScriptsNotFound with default headers values
func NewGetScriptsNotFound() *GetScriptsNotFound {
	return &GetScriptsNotFound{}
}

/*GetScriptsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScriptsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsNotFound  %+v", 404, o.Payload)
}

func (o *GetScriptsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsRequestEntityTooLarge creates a GetScriptsRequestEntityTooLarge with default headers values
func NewGetScriptsRequestEntityTooLarge() *GetScriptsRequestEntityTooLarge {
	return &GetScriptsRequestEntityTooLarge{}
}

/*GetScriptsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetScriptsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScriptsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsUnsupportedMediaType creates a GetScriptsUnsupportedMediaType with default headers values
func NewGetScriptsUnsupportedMediaType() *GetScriptsUnsupportedMediaType {
	return &GetScriptsUnsupportedMediaType{}
}

/*GetScriptsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScriptsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScriptsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsTooManyRequests creates a GetScriptsTooManyRequests with default headers values
func NewGetScriptsTooManyRequests() *GetScriptsTooManyRequests {
	return &GetScriptsTooManyRequests{}
}

/*GetScriptsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetScriptsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScriptsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsInternalServerError creates a GetScriptsInternalServerError with default headers values
func NewGetScriptsInternalServerError() *GetScriptsInternalServerError {
	return &GetScriptsInternalServerError{}
}

/*GetScriptsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScriptsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScriptsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsServiceUnavailable creates a GetScriptsServiceUnavailable with default headers values
func NewGetScriptsServiceUnavailable() *GetScriptsServiceUnavailable {
	return &GetScriptsServiceUnavailable{}
}

/*GetScriptsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScriptsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScriptsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScriptsGatewayTimeout creates a GetScriptsGatewayTimeout with default headers values
func NewGetScriptsGatewayTimeout() *GetScriptsGatewayTimeout {
	return &GetScriptsGatewayTimeout{}
}

/*GetScriptsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScriptsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScriptsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scripts][%d] getScriptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScriptsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScriptsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
