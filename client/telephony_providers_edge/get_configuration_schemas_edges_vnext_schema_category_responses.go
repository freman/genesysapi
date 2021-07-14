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

// GetConfigurationSchemasEdgesVnextSchemaCategoryReader is a Reader for the GetConfigurationSchemasEdgesVnextSchemaCategory structure.
type GetConfigurationSchemasEdgesVnextSchemaCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryOK creates a GetConfigurationSchemasEdgesVnextSchemaCategoryOK with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryOK() *GetConfigurationSchemasEdgesVnextSchemaCategoryOK {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryOK{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryOK handles this case with default header values.

successful operation
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryOK struct {
	Payload *models.SchemaReferenceEntityListing
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryOK) GetPayload() *models.SchemaReferenceEntityListing {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchemaReferenceEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest creates a GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest() *GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized creates a GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized() *GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryForbidden creates a GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryForbidden() *GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryNotFound creates a GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryNotFound() *GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout creates a GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout() *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge creates a GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge() *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType creates a GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType() *GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests creates a GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests() *GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError creates a GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError() *GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable creates a GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable() *GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout creates a GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout with default headers values
func NewGetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout() *GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout {
	return &GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout{}
}

/*GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/configuration/schemas/edges/vnext/{schemaCategory}][%d] getConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConfigurationSchemasEdgesVnextSchemaCategoryGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
