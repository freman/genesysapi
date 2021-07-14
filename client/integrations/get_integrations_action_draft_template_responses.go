// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIntegrationsActionDraftTemplateReader is a Reader for the GetIntegrationsActionDraftTemplate structure.
type GetIntegrationsActionDraftTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionDraftTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionDraftTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionDraftTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionDraftTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionDraftTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionDraftTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsActionDraftTemplateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionDraftTemplateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionDraftTemplateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionDraftTemplateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionDraftTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionDraftTemplateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionDraftTemplateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionDraftTemplateOK creates a GetIntegrationsActionDraftTemplateOK with default headers values
func NewGetIntegrationsActionDraftTemplateOK() *GetIntegrationsActionDraftTemplateOK {
	return &GetIntegrationsActionDraftTemplateOK{}
}

/*GetIntegrationsActionDraftTemplateOK handles this case with default header values.

successful operation
*/
type GetIntegrationsActionDraftTemplateOK struct {
	Payload string
}

func (o *GetIntegrationsActionDraftTemplateOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateOK) GetPayload() string {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateBadRequest creates a GetIntegrationsActionDraftTemplateBadRequest with default headers values
func NewGetIntegrationsActionDraftTemplateBadRequest() *GetIntegrationsActionDraftTemplateBadRequest {
	return &GetIntegrationsActionDraftTemplateBadRequest{}
}

/*GetIntegrationsActionDraftTemplateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionDraftTemplateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateUnauthorized creates a GetIntegrationsActionDraftTemplateUnauthorized with default headers values
func NewGetIntegrationsActionDraftTemplateUnauthorized() *GetIntegrationsActionDraftTemplateUnauthorized {
	return &GetIntegrationsActionDraftTemplateUnauthorized{}
}

/*GetIntegrationsActionDraftTemplateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionDraftTemplateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateForbidden creates a GetIntegrationsActionDraftTemplateForbidden with default headers values
func NewGetIntegrationsActionDraftTemplateForbidden() *GetIntegrationsActionDraftTemplateForbidden {
	return &GetIntegrationsActionDraftTemplateForbidden{}
}

/*GetIntegrationsActionDraftTemplateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionDraftTemplateForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateNotFound creates a GetIntegrationsActionDraftTemplateNotFound with default headers values
func NewGetIntegrationsActionDraftTemplateNotFound() *GetIntegrationsActionDraftTemplateNotFound {
	return &GetIntegrationsActionDraftTemplateNotFound{}
}

/*GetIntegrationsActionDraftTemplateNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionDraftTemplateNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateRequestTimeout creates a GetIntegrationsActionDraftTemplateRequestTimeout with default headers values
func NewGetIntegrationsActionDraftTemplateRequestTimeout() *GetIntegrationsActionDraftTemplateRequestTimeout {
	return &GetIntegrationsActionDraftTemplateRequestTimeout{}
}

/*GetIntegrationsActionDraftTemplateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsActionDraftTemplateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateRequestEntityTooLarge creates a GetIntegrationsActionDraftTemplateRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionDraftTemplateRequestEntityTooLarge() *GetIntegrationsActionDraftTemplateRequestEntityTooLarge {
	return &GetIntegrationsActionDraftTemplateRequestEntityTooLarge{}
}

/*GetIntegrationsActionDraftTemplateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsActionDraftTemplateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateUnsupportedMediaType creates a GetIntegrationsActionDraftTemplateUnsupportedMediaType with default headers values
func NewGetIntegrationsActionDraftTemplateUnsupportedMediaType() *GetIntegrationsActionDraftTemplateUnsupportedMediaType {
	return &GetIntegrationsActionDraftTemplateUnsupportedMediaType{}
}

/*GetIntegrationsActionDraftTemplateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionDraftTemplateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateTooManyRequests creates a GetIntegrationsActionDraftTemplateTooManyRequests with default headers values
func NewGetIntegrationsActionDraftTemplateTooManyRequests() *GetIntegrationsActionDraftTemplateTooManyRequests {
	return &GetIntegrationsActionDraftTemplateTooManyRequests{}
}

/*GetIntegrationsActionDraftTemplateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsActionDraftTemplateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateInternalServerError creates a GetIntegrationsActionDraftTemplateInternalServerError with default headers values
func NewGetIntegrationsActionDraftTemplateInternalServerError() *GetIntegrationsActionDraftTemplateInternalServerError {
	return &GetIntegrationsActionDraftTemplateInternalServerError{}
}

/*GetIntegrationsActionDraftTemplateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionDraftTemplateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateServiceUnavailable creates a GetIntegrationsActionDraftTemplateServiceUnavailable with default headers values
func NewGetIntegrationsActionDraftTemplateServiceUnavailable() *GetIntegrationsActionDraftTemplateServiceUnavailable {
	return &GetIntegrationsActionDraftTemplateServiceUnavailable{}
}

/*GetIntegrationsActionDraftTemplateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionDraftTemplateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftTemplateGatewayTimeout creates a GetIntegrationsActionDraftTemplateGatewayTimeout with default headers values
func NewGetIntegrationsActionDraftTemplateGatewayTimeout() *GetIntegrationsActionDraftTemplateGatewayTimeout {
	return &GetIntegrationsActionDraftTemplateGatewayTimeout{}
}

/*GetIntegrationsActionDraftTemplateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsActionDraftTemplateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftTemplateGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/templates/{fileName}][%d] getIntegrationsActionDraftTemplateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionDraftTemplateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftTemplateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
