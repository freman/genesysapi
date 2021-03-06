// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeReader is a Reader for the GetConfigurationSchemasEdgesVnextSchemaCategorySchemaType structure.
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK handles this case with default header values.

successful operation
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK struct {
	Payload *models.SchemaReferenceEntityListing
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK) GetPayload() *models.SchemaReferenceEntityListing {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchemaReferenceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
