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

// GetArchitectSystempromptResourcesReader is a Reader for the GetArchitectSystempromptResources structure.
type GetArchitectSystempromptResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectSystempromptResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectSystempromptResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectSystempromptResourcesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectSystempromptResourcesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectSystempromptResourcesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectSystempromptResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectSystempromptResourcesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectSystempromptResourcesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectSystempromptResourcesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectSystempromptResourcesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectSystempromptResourcesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectSystempromptResourcesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectSystempromptResourcesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectSystempromptResourcesOK creates a GetArchitectSystempromptResourcesOK with default headers values
func NewGetArchitectSystempromptResourcesOK() *GetArchitectSystempromptResourcesOK {
	return &GetArchitectSystempromptResourcesOK{}
}

/*GetArchitectSystempromptResourcesOK handles this case with default header values.

successful operation
*/
type GetArchitectSystempromptResourcesOK struct {
	Payload *models.SystemPromptAssetEntityListing
}

func (o *GetArchitectSystempromptResourcesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesOK  %+v", 200, o.Payload)
}

func (o *GetArchitectSystempromptResourcesOK) GetPayload() *models.SystemPromptAssetEntityListing {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemPromptAssetEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesBadRequest creates a GetArchitectSystempromptResourcesBadRequest with default headers values
func NewGetArchitectSystempromptResourcesBadRequest() *GetArchitectSystempromptResourcesBadRequest {
	return &GetArchitectSystempromptResourcesBadRequest{}
}

/*GetArchitectSystempromptResourcesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectSystempromptResourcesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectSystempromptResourcesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesUnauthorized creates a GetArchitectSystempromptResourcesUnauthorized with default headers values
func NewGetArchitectSystempromptResourcesUnauthorized() *GetArchitectSystempromptResourcesUnauthorized {
	return &GetArchitectSystempromptResourcesUnauthorized{}
}

/*GetArchitectSystempromptResourcesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectSystempromptResourcesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectSystempromptResourcesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesForbidden creates a GetArchitectSystempromptResourcesForbidden with default headers values
func NewGetArchitectSystempromptResourcesForbidden() *GetArchitectSystempromptResourcesForbidden {
	return &GetArchitectSystempromptResourcesForbidden{}
}

/*GetArchitectSystempromptResourcesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectSystempromptResourcesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectSystempromptResourcesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesNotFound creates a GetArchitectSystempromptResourcesNotFound with default headers values
func NewGetArchitectSystempromptResourcesNotFound() *GetArchitectSystempromptResourcesNotFound {
	return &GetArchitectSystempromptResourcesNotFound{}
}

/*GetArchitectSystempromptResourcesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectSystempromptResourcesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectSystempromptResourcesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesRequestTimeout creates a GetArchitectSystempromptResourcesRequestTimeout with default headers values
func NewGetArchitectSystempromptResourcesRequestTimeout() *GetArchitectSystempromptResourcesRequestTimeout {
	return &GetArchitectSystempromptResourcesRequestTimeout{}
}

/*GetArchitectSystempromptResourcesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectSystempromptResourcesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectSystempromptResourcesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesRequestEntityTooLarge creates a GetArchitectSystempromptResourcesRequestEntityTooLarge with default headers values
func NewGetArchitectSystempromptResourcesRequestEntityTooLarge() *GetArchitectSystempromptResourcesRequestEntityTooLarge {
	return &GetArchitectSystempromptResourcesRequestEntityTooLarge{}
}

/*GetArchitectSystempromptResourcesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetArchitectSystempromptResourcesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectSystempromptResourcesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesUnsupportedMediaType creates a GetArchitectSystempromptResourcesUnsupportedMediaType with default headers values
func NewGetArchitectSystempromptResourcesUnsupportedMediaType() *GetArchitectSystempromptResourcesUnsupportedMediaType {
	return &GetArchitectSystempromptResourcesUnsupportedMediaType{}
}

/*GetArchitectSystempromptResourcesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectSystempromptResourcesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectSystempromptResourcesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesTooManyRequests creates a GetArchitectSystempromptResourcesTooManyRequests with default headers values
func NewGetArchitectSystempromptResourcesTooManyRequests() *GetArchitectSystempromptResourcesTooManyRequests {
	return &GetArchitectSystempromptResourcesTooManyRequests{}
}

/*GetArchitectSystempromptResourcesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectSystempromptResourcesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectSystempromptResourcesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesInternalServerError creates a GetArchitectSystempromptResourcesInternalServerError with default headers values
func NewGetArchitectSystempromptResourcesInternalServerError() *GetArchitectSystempromptResourcesInternalServerError {
	return &GetArchitectSystempromptResourcesInternalServerError{}
}

/*GetArchitectSystempromptResourcesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectSystempromptResourcesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectSystempromptResourcesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesServiceUnavailable creates a GetArchitectSystempromptResourcesServiceUnavailable with default headers values
func NewGetArchitectSystempromptResourcesServiceUnavailable() *GetArchitectSystempromptResourcesServiceUnavailable {
	return &GetArchitectSystempromptResourcesServiceUnavailable{}
}

/*GetArchitectSystempromptResourcesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectSystempromptResourcesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectSystempromptResourcesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectSystempromptResourcesGatewayTimeout creates a GetArchitectSystempromptResourcesGatewayTimeout with default headers values
func NewGetArchitectSystempromptResourcesGatewayTimeout() *GetArchitectSystempromptResourcesGatewayTimeout {
	return &GetArchitectSystempromptResourcesGatewayTimeout{}
}

/*GetArchitectSystempromptResourcesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectSystempromptResourcesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectSystempromptResourcesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/systemprompts/{promptId}/resources][%d] getArchitectSystempromptResourcesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectSystempromptResourcesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectSystempromptResourcesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
