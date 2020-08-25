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

// GetArchitectSystempromptReader is a Reader for the GetArchitectSystemprompt structure.
type GetArchitectSystempromptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectSystempromptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectSystempromptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectSystempromptBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectSystempromptUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectSystempromptForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectSystempromptNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectSystempromptRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectSystempromptUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectSystempromptTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectSystempromptInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectSystempromptServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectSystempromptGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectSystempromptOK creates a GetArchitectSystempromptOK with default headers values
func NewGetArchitectSystempromptOK() *GetArchitectSystempromptOK {
	return &GetArchitectSystempromptOK{}
}

/*GetArchitectSystempromptOK handles this case with default header values.

successful operation
*/
type GetArchitectSystempromptOK struct {
	Payload *models.SystemPrompt
}

func (o *GetArchitectSystempromptOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptOK) GetPayload() *models.SystemPrompt {
	return o.Payload
}

func (o *GetArchitectSystempromptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPrompt)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptBadRequest creates a GetArchitectSystempromptBadRequest with default headers values
func NewGetArchitectSystempromptBadRequest() *GetArchitectSystempromptBadRequest {
	return &GetArchitectSystempromptBadRequest{}
}

/*GetArchitectSystempromptBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectSystempromptBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptUnauthorized creates a GetArchitectSystempromptUnauthorized with default headers values
func NewGetArchitectSystempromptUnauthorized() *GetArchitectSystempromptUnauthorized {
	return &GetArchitectSystempromptUnauthorized{}
}

/*GetArchitectSystempromptUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectSystempromptUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptForbidden creates a GetArchitectSystempromptForbidden with default headers values
func NewGetArchitectSystempromptForbidden() *GetArchitectSystempromptForbidden {
	return &GetArchitectSystempromptForbidden{}
}

/*GetArchitectSystempromptForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectSystempromptForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptNotFound creates a GetArchitectSystempromptNotFound with default headers values
func NewGetArchitectSystempromptNotFound() *GetArchitectSystempromptNotFound {
	return &GetArchitectSystempromptNotFound{}
}

/*GetArchitectSystempromptNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectSystempromptNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptRequestEntityTooLarge creates a GetArchitectSystempromptRequestEntityTooLarge with default headers values
func NewGetArchitectSystempromptRequestEntityTooLarge() *GetArchitectSystempromptRequestEntityTooLarge {
	return &GetArchitectSystempromptRequestEntityTooLarge{}
}

/*GetArchitectSystempromptRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectSystempromptRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptUnsupportedMediaType creates a GetArchitectSystempromptUnsupportedMediaType with default headers values
func NewGetArchitectSystempromptUnsupportedMediaType() *GetArchitectSystempromptUnsupportedMediaType {
	return &GetArchitectSystempromptUnsupportedMediaType{}
}

/*GetArchitectSystempromptUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectSystempromptUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptTooManyRequests creates a GetArchitectSystempromptTooManyRequests with default headers values
func NewGetArchitectSystempromptTooManyRequests() *GetArchitectSystempromptTooManyRequests {
	return &GetArchitectSystempromptTooManyRequests{}
}

/*GetArchitectSystempromptTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectSystempromptTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptInternalServerError creates a GetArchitectSystempromptInternalServerError with default headers values
func NewGetArchitectSystempromptInternalServerError() *GetArchitectSystempromptInternalServerError {
	return &GetArchitectSystempromptInternalServerError{}
}

/*GetArchitectSystempromptInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectSystempromptInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptServiceUnavailable creates a GetArchitectSystempromptServiceUnavailable with default headers values
func NewGetArchitectSystempromptServiceUnavailable() *GetArchitectSystempromptServiceUnavailable {
	return &GetArchitectSystempromptServiceUnavailable{}
}

/*GetArchitectSystempromptServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectSystempromptServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptGatewayTimeout creates a GetArchitectSystempromptGatewayTimeout with default headers values
func NewGetArchitectSystempromptGatewayTimeout() *GetArchitectSystempromptGatewayTimeout {
	return &GetArchitectSystempromptGatewayTimeout{}
}

/*GetArchitectSystempromptGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectSystempromptGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}][%d] getArchitectSystempromptGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}