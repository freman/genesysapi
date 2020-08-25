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

// GetIntegrationsActionDraftSchemaReader is a Reader for the GetIntegrationsActionDraftSchema structure.
type GetIntegrationsActionDraftSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsActionDraftSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsActionDraftSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsActionDraftSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsActionDraftSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsActionDraftSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsActionDraftSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsActionDraftSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsActionDraftSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsActionDraftSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsActionDraftSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsActionDraftSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsActionDraftSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsActionDraftSchemaOK creates a GetIntegrationsActionDraftSchemaOK with default headers values
func NewGetIntegrationsActionDraftSchemaOK() *GetIntegrationsActionDraftSchemaOK {
	return &GetIntegrationsActionDraftSchemaOK{}
}

/*GetIntegrationsActionDraftSchemaOK handles this case with default header values.

successful operation
*/
type GetIntegrationsActionDraftSchemaOK struct {
	Payload *models.JSONSchemaDocument
}

func (o *GetIntegrationsActionDraftSchemaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaOK) GetPayload() *models.JSONSchemaDocument {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONSchemaDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaBadRequest creates a GetIntegrationsActionDraftSchemaBadRequest with default headers values
func NewGetIntegrationsActionDraftSchemaBadRequest() *GetIntegrationsActionDraftSchemaBadRequest {
	return &GetIntegrationsActionDraftSchemaBadRequest{}
}

/*GetIntegrationsActionDraftSchemaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsActionDraftSchemaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaUnauthorized creates a GetIntegrationsActionDraftSchemaUnauthorized with default headers values
func NewGetIntegrationsActionDraftSchemaUnauthorized() *GetIntegrationsActionDraftSchemaUnauthorized {
	return &GetIntegrationsActionDraftSchemaUnauthorized{}
}

/*GetIntegrationsActionDraftSchemaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsActionDraftSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaForbidden creates a GetIntegrationsActionDraftSchemaForbidden with default headers values
func NewGetIntegrationsActionDraftSchemaForbidden() *GetIntegrationsActionDraftSchemaForbidden {
	return &GetIntegrationsActionDraftSchemaForbidden{}
}

/*GetIntegrationsActionDraftSchemaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsActionDraftSchemaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaNotFound creates a GetIntegrationsActionDraftSchemaNotFound with default headers values
func NewGetIntegrationsActionDraftSchemaNotFound() *GetIntegrationsActionDraftSchemaNotFound {
	return &GetIntegrationsActionDraftSchemaNotFound{}
}

/*GetIntegrationsActionDraftSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetIntegrationsActionDraftSchemaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaRequestEntityTooLarge creates a GetIntegrationsActionDraftSchemaRequestEntityTooLarge with default headers values
func NewGetIntegrationsActionDraftSchemaRequestEntityTooLarge() *GetIntegrationsActionDraftSchemaRequestEntityTooLarge {
	return &GetIntegrationsActionDraftSchemaRequestEntityTooLarge{}
}

/*GetIntegrationsActionDraftSchemaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetIntegrationsActionDraftSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaUnsupportedMediaType creates a GetIntegrationsActionDraftSchemaUnsupportedMediaType with default headers values
func NewGetIntegrationsActionDraftSchemaUnsupportedMediaType() *GetIntegrationsActionDraftSchemaUnsupportedMediaType {
	return &GetIntegrationsActionDraftSchemaUnsupportedMediaType{}
}

/*GetIntegrationsActionDraftSchemaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsActionDraftSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaTooManyRequests creates a GetIntegrationsActionDraftSchemaTooManyRequests with default headers values
func NewGetIntegrationsActionDraftSchemaTooManyRequests() *GetIntegrationsActionDraftSchemaTooManyRequests {
	return &GetIntegrationsActionDraftSchemaTooManyRequests{}
}

/*GetIntegrationsActionDraftSchemaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetIntegrationsActionDraftSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaInternalServerError creates a GetIntegrationsActionDraftSchemaInternalServerError with default headers values
func NewGetIntegrationsActionDraftSchemaInternalServerError() *GetIntegrationsActionDraftSchemaInternalServerError {
	return &GetIntegrationsActionDraftSchemaInternalServerError{}
}

/*GetIntegrationsActionDraftSchemaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsActionDraftSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaServiceUnavailable creates a GetIntegrationsActionDraftSchemaServiceUnavailable with default headers values
func NewGetIntegrationsActionDraftSchemaServiceUnavailable() *GetIntegrationsActionDraftSchemaServiceUnavailable {
	return &GetIntegrationsActionDraftSchemaServiceUnavailable{}
}

/*GetIntegrationsActionDraftSchemaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsActionDraftSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsActionDraftSchemaGatewayTimeout creates a GetIntegrationsActionDraftSchemaGatewayTimeout with default headers values
func NewGetIntegrationsActionDraftSchemaGatewayTimeout() *GetIntegrationsActionDraftSchemaGatewayTimeout {
	return &GetIntegrationsActionDraftSchemaGatewayTimeout{}
}

/*GetIntegrationsActionDraftSchemaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetIntegrationsActionDraftSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetIntegrationsActionDraftSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/actions/{actionId}/draft/schemas/{fileName}][%d] getIntegrationsActionDraftSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsActionDraftSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsActionDraftSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}