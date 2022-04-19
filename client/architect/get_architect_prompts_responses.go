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

// GetArchitectPromptsReader is a Reader for the GetArchitectPrompts structure.
type GetArchitectPromptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectPromptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectPromptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectPromptsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectPromptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectPromptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectPromptsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectPromptsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectPromptsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectPromptsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectPromptsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectPromptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectPromptsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectPromptsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectPromptsOK creates a GetArchitectPromptsOK with default headers values
func NewGetArchitectPromptsOK() *GetArchitectPromptsOK {
	return &GetArchitectPromptsOK{}
}

/*GetArchitectPromptsOK handles this case with default header values.

successful operation
*/
type GetArchitectPromptsOK struct {
	Payload *models.PromptEntityListing
}

func (o *GetArchitectPromptsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsOK  %+v", 200, o.Payload)
}

func (o *GetArchitectPromptsOK) GetPayload() *models.PromptEntityListing {
	return o.Payload
}

func (o *GetArchitectPromptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PromptEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsBadRequest creates a GetArchitectPromptsBadRequest with default headers values
func NewGetArchitectPromptsBadRequest() *GetArchitectPromptsBadRequest {
	return &GetArchitectPromptsBadRequest{}
}

/*GetArchitectPromptsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectPromptsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectPromptsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsUnauthorized creates a GetArchitectPromptsUnauthorized with default headers values
func NewGetArchitectPromptsUnauthorized() *GetArchitectPromptsUnauthorized {
	return &GetArchitectPromptsUnauthorized{}
}

/*GetArchitectPromptsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectPromptsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectPromptsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsForbidden creates a GetArchitectPromptsForbidden with default headers values
func NewGetArchitectPromptsForbidden() *GetArchitectPromptsForbidden {
	return &GetArchitectPromptsForbidden{}
}

/*GetArchitectPromptsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectPromptsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectPromptsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsNotFound creates a GetArchitectPromptsNotFound with default headers values
func NewGetArchitectPromptsNotFound() *GetArchitectPromptsNotFound {
	return &GetArchitectPromptsNotFound{}
}

/*GetArchitectPromptsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectPromptsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectPromptsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsRequestTimeout creates a GetArchitectPromptsRequestTimeout with default headers values
func NewGetArchitectPromptsRequestTimeout() *GetArchitectPromptsRequestTimeout {
	return &GetArchitectPromptsRequestTimeout{}
}

/*GetArchitectPromptsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectPromptsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectPromptsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsRequestEntityTooLarge creates a GetArchitectPromptsRequestEntityTooLarge with default headers values
func NewGetArchitectPromptsRequestEntityTooLarge() *GetArchitectPromptsRequestEntityTooLarge {
	return &GetArchitectPromptsRequestEntityTooLarge{}
}

/*GetArchitectPromptsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectPromptsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectPromptsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsUnsupportedMediaType creates a GetArchitectPromptsUnsupportedMediaType with default headers values
func NewGetArchitectPromptsUnsupportedMediaType() *GetArchitectPromptsUnsupportedMediaType {
	return &GetArchitectPromptsUnsupportedMediaType{}
}

/*GetArchitectPromptsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectPromptsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectPromptsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsTooManyRequests creates a GetArchitectPromptsTooManyRequests with default headers values
func NewGetArchitectPromptsTooManyRequests() *GetArchitectPromptsTooManyRequests {
	return &GetArchitectPromptsTooManyRequests{}
}

/*GetArchitectPromptsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectPromptsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectPromptsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsInternalServerError creates a GetArchitectPromptsInternalServerError with default headers values
func NewGetArchitectPromptsInternalServerError() *GetArchitectPromptsInternalServerError {
	return &GetArchitectPromptsInternalServerError{}
}

/*GetArchitectPromptsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectPromptsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectPromptsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsServiceUnavailable creates a GetArchitectPromptsServiceUnavailable with default headers values
func NewGetArchitectPromptsServiceUnavailable() *GetArchitectPromptsServiceUnavailable {
	return &GetArchitectPromptsServiceUnavailable{}
}

/*GetArchitectPromptsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectPromptsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectPromptsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectPromptsGatewayTimeout creates a GetArchitectPromptsGatewayTimeout with default headers values
func NewGetArchitectPromptsGatewayTimeout() *GetArchitectPromptsGatewayTimeout {
	return &GetArchitectPromptsGatewayTimeout{}
}

/*GetArchitectPromptsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectPromptsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectPromptsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/prompts][%d] getArchitectPromptsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectPromptsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectPromptsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
