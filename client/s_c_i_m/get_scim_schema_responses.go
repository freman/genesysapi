// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetScimSchemaReader is a Reader for the GetScimSchema structure.
type GetScimSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScimSchemaRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimSchemaOK creates a GetScimSchemaOK with default headers values
func NewGetScimSchemaOK() *GetScimSchemaOK {
	return &GetScimSchemaOK{}
}

/*GetScimSchemaOK handles this case with default header values.

successful operation
*/
type GetScimSchemaOK struct {
	Payload *models.ScimV2SchemaDefinition
}

func (o *GetScimSchemaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaOK  %+v", 200, o.Payload)
}

func (o *GetScimSchemaOK) GetPayload() *models.ScimV2SchemaDefinition {
	return o.Payload
}

func (o *GetScimSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2SchemaDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaBadRequest creates a GetScimSchemaBadRequest with default headers values
func NewGetScimSchemaBadRequest() *GetScimSchemaBadRequest {
	return &GetScimSchemaBadRequest{}
}

/*GetScimSchemaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimSchemaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaUnauthorized creates a GetScimSchemaUnauthorized with default headers values
func NewGetScimSchemaUnauthorized() *GetScimSchemaUnauthorized {
	return &GetScimSchemaUnauthorized{}
}

/*GetScimSchemaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaForbidden creates a GetScimSchemaForbidden with default headers values
func NewGetScimSchemaForbidden() *GetScimSchemaForbidden {
	return &GetScimSchemaForbidden{}
}

/*GetScimSchemaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetScimSchemaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaForbidden  %+v", 403, o.Payload)
}

func (o *GetScimSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaNotFound creates a GetScimSchemaNotFound with default headers values
func NewGetScimSchemaNotFound() *GetScimSchemaNotFound {
	return &GetScimSchemaNotFound{}
}

/*GetScimSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetScimSchemaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetScimSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaRequestTimeout creates a GetScimSchemaRequestTimeout with default headers values
func NewGetScimSchemaRequestTimeout() *GetScimSchemaRequestTimeout {
	return &GetScimSchemaRequestTimeout{}
}

/*GetScimSchemaRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScimSchemaRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimSchemaRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaRequestEntityTooLarge creates a GetScimSchemaRequestEntityTooLarge with default headers values
func NewGetScimSchemaRequestEntityTooLarge() *GetScimSchemaRequestEntityTooLarge {
	return &GetScimSchemaRequestEntityTooLarge{}
}

/*GetScimSchemaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetScimSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaUnsupportedMediaType creates a GetScimSchemaUnsupportedMediaType with default headers values
func NewGetScimSchemaUnsupportedMediaType() *GetScimSchemaUnsupportedMediaType {
	return &GetScimSchemaUnsupportedMediaType{}
}

/*GetScimSchemaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaTooManyRequests creates a GetScimSchemaTooManyRequests with default headers values
func NewGetScimSchemaTooManyRequests() *GetScimSchemaTooManyRequests {
	return &GetScimSchemaTooManyRequests{}
}

/*GetScimSchemaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScimSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaInternalServerError creates a GetScimSchemaInternalServerError with default headers values
func NewGetScimSchemaInternalServerError() *GetScimSchemaInternalServerError {
	return &GetScimSchemaInternalServerError{}
}

/*GetScimSchemaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaServiceUnavailable creates a GetScimSchemaServiceUnavailable with default headers values
func NewGetScimSchemaServiceUnavailable() *GetScimSchemaServiceUnavailable {
	return &GetScimSchemaServiceUnavailable{}
}

/*GetScimSchemaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimSchemaGatewayTimeout creates a GetScimSchemaGatewayTimeout with default headers values
func NewGetScimSchemaGatewayTimeout() *GetScimSchemaGatewayTimeout {
	return &GetScimSchemaGatewayTimeout{}
}

/*GetScimSchemaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetScimSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetScimSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/schemas/{schemaId}][%d] getScimSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
