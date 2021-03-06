// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetExternalcontactsOrganizationsSchemaReader is a Reader for the GetExternalcontactsOrganizationsSchema structure.
type GetExternalcontactsOrganizationsSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsOrganizationsSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsOrganizationsSchemaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsOrganizationsSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsOrganizationsSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsOrganizationsSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsOrganizationsSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsOrganizationsSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsOrganizationsSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsOrganizationsSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsOrganizationsSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsOrganizationsSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsOrganizationsSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsOrganizationsSchemaOK creates a GetExternalcontactsOrganizationsSchemaOK with default headers values
func NewGetExternalcontactsOrganizationsSchemaOK() *GetExternalcontactsOrganizationsSchemaOK {
	return &GetExternalcontactsOrganizationsSchemaOK{}
}

/*GetExternalcontactsOrganizationsSchemaOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsOrganizationsSchemaOK struct {
	Payload *models.DataSchema
}

func (o *GetExternalcontactsOrganizationsSchemaOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaOK) GetPayload() *models.DataSchema {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaBadRequest creates a GetExternalcontactsOrganizationsSchemaBadRequest with default headers values
func NewGetExternalcontactsOrganizationsSchemaBadRequest() *GetExternalcontactsOrganizationsSchemaBadRequest {
	return &GetExternalcontactsOrganizationsSchemaBadRequest{}
}

/*GetExternalcontactsOrganizationsSchemaBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsOrganizationsSchemaBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaUnauthorized creates a GetExternalcontactsOrganizationsSchemaUnauthorized with default headers values
func NewGetExternalcontactsOrganizationsSchemaUnauthorized() *GetExternalcontactsOrganizationsSchemaUnauthorized {
	return &GetExternalcontactsOrganizationsSchemaUnauthorized{}
}

/*GetExternalcontactsOrganizationsSchemaUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsOrganizationsSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaForbidden creates a GetExternalcontactsOrganizationsSchemaForbidden with default headers values
func NewGetExternalcontactsOrganizationsSchemaForbidden() *GetExternalcontactsOrganizationsSchemaForbidden {
	return &GetExternalcontactsOrganizationsSchemaForbidden{}
}

/*GetExternalcontactsOrganizationsSchemaForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsOrganizationsSchemaForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaNotFound creates a GetExternalcontactsOrganizationsSchemaNotFound with default headers values
func NewGetExternalcontactsOrganizationsSchemaNotFound() *GetExternalcontactsOrganizationsSchemaNotFound {
	return &GetExternalcontactsOrganizationsSchemaNotFound{}
}

/*GetExternalcontactsOrganizationsSchemaNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsOrganizationsSchemaNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaRequestEntityTooLarge creates a GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge with default headers values
func NewGetExternalcontactsOrganizationsSchemaRequestEntityTooLarge() *GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge {
	return &GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge{}
}

/*GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaUnsupportedMediaType creates a GetExternalcontactsOrganizationsSchemaUnsupportedMediaType with default headers values
func NewGetExternalcontactsOrganizationsSchemaUnsupportedMediaType() *GetExternalcontactsOrganizationsSchemaUnsupportedMediaType {
	return &GetExternalcontactsOrganizationsSchemaUnsupportedMediaType{}
}

/*GetExternalcontactsOrganizationsSchemaUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsOrganizationsSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaTooManyRequests creates a GetExternalcontactsOrganizationsSchemaTooManyRequests with default headers values
func NewGetExternalcontactsOrganizationsSchemaTooManyRequests() *GetExternalcontactsOrganizationsSchemaTooManyRequests {
	return &GetExternalcontactsOrganizationsSchemaTooManyRequests{}
}

/*GetExternalcontactsOrganizationsSchemaTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetExternalcontactsOrganizationsSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaInternalServerError creates a GetExternalcontactsOrganizationsSchemaInternalServerError with default headers values
func NewGetExternalcontactsOrganizationsSchemaInternalServerError() *GetExternalcontactsOrganizationsSchemaInternalServerError {
	return &GetExternalcontactsOrganizationsSchemaInternalServerError{}
}

/*GetExternalcontactsOrganizationsSchemaInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsOrganizationsSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaServiceUnavailable creates a GetExternalcontactsOrganizationsSchemaServiceUnavailable with default headers values
func NewGetExternalcontactsOrganizationsSchemaServiceUnavailable() *GetExternalcontactsOrganizationsSchemaServiceUnavailable {
	return &GetExternalcontactsOrganizationsSchemaServiceUnavailable{}
}

/*GetExternalcontactsOrganizationsSchemaServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsOrganizationsSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsOrganizationsSchemaGatewayTimeout creates a GetExternalcontactsOrganizationsSchemaGatewayTimeout with default headers values
func NewGetExternalcontactsOrganizationsSchemaGatewayTimeout() *GetExternalcontactsOrganizationsSchemaGatewayTimeout {
	return &GetExternalcontactsOrganizationsSchemaGatewayTimeout{}
}

/*GetExternalcontactsOrganizationsSchemaGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsOrganizationsSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsOrganizationsSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/organizations/schemas/{schemaId}][%d] getExternalcontactsOrganizationsSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsOrganizationsSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsOrganizationsSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
