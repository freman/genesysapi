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

// GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDReader is a Reader for the GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataID structure.
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK handles this case with default header values.

successful operation
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK struct {
	Payload *models.Organization
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK) GetPayload() *models.Organization {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout creates a GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout() *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout {
	return &GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}/{schemaType}/{schemaId}/{extensionType}/{metadataId}][%d] getConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIdExtensionTypeMetadataIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategorySchemaTypeSchemaIDExtensionTypeMetadataIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
