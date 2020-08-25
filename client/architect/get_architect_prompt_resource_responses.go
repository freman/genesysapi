// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetArchitectPromptResourceReader is a Reader for the GetArchitectPromptResource structure.
type GetArchitectPromptResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectPromptResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectPromptResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectPromptResourceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectPromptResourceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectPromptResourceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectPromptResourceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectPromptResourceRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectPromptResourceUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectPromptResourceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectPromptResourceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectPromptResourceServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectPromptResourceGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectPromptResourceOK creates a GetArchitectPromptResourceOK with default headers values
func NewGetArchitectPromptResourceOK() *GetArchitectPromptResourceOK {
	return &GetArchitectPromptResourceOK{}
}

/*GetArchitectPromptResourceOK handles this case with default header values.

successful operation
*/
type GetArchitectPromptResourceOK struct {
	Payload *models.PromptAsset
}

func (o *GetArchitectPromptResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptResourceOK) GetPayload() *models.PromptAsset {
	return o.Payload
}

func (o *GetArchitectPromptResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromptAsset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceBadRequest creates a GetArchitectPromptResourceBadRequest with default headers values
func NewGetArchitectPromptResourceBadRequest() *GetArchitectPromptResourceBadRequest {
	return &GetArchitectPromptResourceBadRequest{}
}

/*GetArchitectPromptResourceBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectPromptResourceBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptResourceBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceUnauthorized creates a GetArchitectPromptResourceUnauthorized with default headers values
func NewGetArchitectPromptResourceUnauthorized() *GetArchitectPromptResourceUnauthorized {
	return &GetArchitectPromptResourceUnauthorized{}
}

/*GetArchitectPromptResourceUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectPromptResourceUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptResourceUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceForbidden creates a GetArchitectPromptResourceForbidden with default headers values
func NewGetArchitectPromptResourceForbidden() *GetArchitectPromptResourceForbidden {
	return &GetArchitectPromptResourceForbidden{}
}

/*GetArchitectPromptResourceForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectPromptResourceForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptResourceForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceNotFound creates a GetArchitectPromptResourceNotFound with default headers values
func NewGetArchitectPromptResourceNotFound() *GetArchitectPromptResourceNotFound {
	return &GetArchitectPromptResourceNotFound{}
}

/*GetArchitectPromptResourceNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectPromptResourceNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptResourceNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceRequestEntityTooLarge creates a GetArchitectPromptResourceRequestEntityTooLarge with default headers values
func NewGetArchitectPromptResourceRequestEntityTooLarge() *GetArchitectPromptResourceRequestEntityTooLarge {
	return &GetArchitectPromptResourceRequestEntityTooLarge{}
}

/*GetArchitectPromptResourceRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectPromptResourceRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceUnsupportedMediaType creates a GetArchitectPromptResourceUnsupportedMediaType with default headers values
func NewGetArchitectPromptResourceUnsupportedMediaType() *GetArchitectPromptResourceUnsupportedMediaType {
	return &GetArchitectPromptResourceUnsupportedMediaType{}
}

/*GetArchitectPromptResourceUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectPromptResourceUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceTooManyRequests creates a GetArchitectPromptResourceTooManyRequests with default headers values
func NewGetArchitectPromptResourceTooManyRequests() *GetArchitectPromptResourceTooManyRequests {
	return &GetArchitectPromptResourceTooManyRequests{}
}

/*GetArchitectPromptResourceTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectPromptResourceTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptResourceTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceInternalServerError creates a GetArchitectPromptResourceInternalServerError with default headers values
func NewGetArchitectPromptResourceInternalServerError() *GetArchitectPromptResourceInternalServerError {
	return &GetArchitectPromptResourceInternalServerError{}
}

/*GetArchitectPromptResourceInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectPromptResourceInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptResourceInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceServiceUnavailable creates a GetArchitectPromptResourceServiceUnavailable with default headers values
func NewGetArchitectPromptResourceServiceUnavailable() *GetArchitectPromptResourceServiceUnavailable {
	return &GetArchitectPromptResourceServiceUnavailable{}
}

/*GetArchitectPromptResourceServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectPromptResourceServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptResourceServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptResourceGatewayTimeout creates a GetArchitectPromptResourceGatewayTimeout with default headers values
func NewGetArchitectPromptResourceGatewayTimeout() *GetArchitectPromptResourceGatewayTimeout {
	return &GetArchitectPromptResourceGatewayTimeout{}
}

/*GetArchitectPromptResourceGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectPromptResourceGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptResourceGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts/{promptId}/resources/{languageCode}][%d] getArchitectPromptResourceGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptResourceGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptResourceGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}