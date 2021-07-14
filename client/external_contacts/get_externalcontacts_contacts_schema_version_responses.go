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

// GetExternalcontactsContactsSchemaVersionReader is a Reader for the GetExternalcontactsContactsSchemaVersion structure.
type GetExternalcontactsContactsSchemaVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExternalcontactsContactsSchemaVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExternalcontactsContactsSchemaVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetExternalcontactsContactsSchemaVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetExternalcontactsContactsSchemaVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExternalcontactsContactsSchemaVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExternalcontactsContactsSchemaVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetExternalcontactsContactsSchemaVersionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetExternalcontactsContactsSchemaVersionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetExternalcontactsContactsSchemaVersionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetExternalcontactsContactsSchemaVersionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetExternalcontactsContactsSchemaVersionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExternalcontactsContactsSchemaVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetExternalcontactsContactsSchemaVersionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetExternalcontactsContactsSchemaVersionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExternalcontactsContactsSchemaVersionOK creates a GetExternalcontactsContactsSchemaVersionOK with default headers values
func NewGetExternalcontactsContactsSchemaVersionOK() *GetExternalcontactsContactsSchemaVersionOK {
	return &GetExternalcontactsContactsSchemaVersionOK{}
}

/*GetExternalcontactsContactsSchemaVersionOK handles this case with default header values.

successful operation
*/
type GetExternalcontactsContactsSchemaVersionOK struct {
	Payload *models.DataSchema
}

func (o *GetExternalcontactsContactsSchemaVersionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionOK  %+v", 200, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionOK) GetPayload() *models.DataSchema {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionBadRequest creates a GetExternalcontactsContactsSchemaVersionBadRequest with default headers values
func NewGetExternalcontactsContactsSchemaVersionBadRequest() *GetExternalcontactsContactsSchemaVersionBadRequest {
	return &GetExternalcontactsContactsSchemaVersionBadRequest{}
}

/*GetExternalcontactsContactsSchemaVersionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetExternalcontactsContactsSchemaVersionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionBadRequest  %+v", 400, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionUnauthorized creates a GetExternalcontactsContactsSchemaVersionUnauthorized with default headers values
func NewGetExternalcontactsContactsSchemaVersionUnauthorized() *GetExternalcontactsContactsSchemaVersionUnauthorized {
	return &GetExternalcontactsContactsSchemaVersionUnauthorized{}
}

/*GetExternalcontactsContactsSchemaVersionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetExternalcontactsContactsSchemaVersionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionForbidden creates a GetExternalcontactsContactsSchemaVersionForbidden with default headers values
func NewGetExternalcontactsContactsSchemaVersionForbidden() *GetExternalcontactsContactsSchemaVersionForbidden {
	return &GetExternalcontactsContactsSchemaVersionForbidden{}
}

/*GetExternalcontactsContactsSchemaVersionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetExternalcontactsContactsSchemaVersionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionForbidden  %+v", 403, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionNotFound creates a GetExternalcontactsContactsSchemaVersionNotFound with default headers values
func NewGetExternalcontactsContactsSchemaVersionNotFound() *GetExternalcontactsContactsSchemaVersionNotFound {
	return &GetExternalcontactsContactsSchemaVersionNotFound{}
}

/*GetExternalcontactsContactsSchemaVersionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetExternalcontactsContactsSchemaVersionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionNotFound  %+v", 404, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionRequestTimeout creates a GetExternalcontactsContactsSchemaVersionRequestTimeout with default headers values
func NewGetExternalcontactsContactsSchemaVersionRequestTimeout() *GetExternalcontactsContactsSchemaVersionRequestTimeout {
	return &GetExternalcontactsContactsSchemaVersionRequestTimeout{}
}

/*GetExternalcontactsContactsSchemaVersionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetExternalcontactsContactsSchemaVersionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionRequestEntityTooLarge creates a GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge with default headers values
func NewGetExternalcontactsContactsSchemaVersionRequestEntityTooLarge() *GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge {
	return &GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge{}
}

/*GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionUnsupportedMediaType creates a GetExternalcontactsContactsSchemaVersionUnsupportedMediaType with default headers values
func NewGetExternalcontactsContactsSchemaVersionUnsupportedMediaType() *GetExternalcontactsContactsSchemaVersionUnsupportedMediaType {
	return &GetExternalcontactsContactsSchemaVersionUnsupportedMediaType{}
}

/*GetExternalcontactsContactsSchemaVersionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetExternalcontactsContactsSchemaVersionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionUnprocessableEntity creates a GetExternalcontactsContactsSchemaVersionUnprocessableEntity with default headers values
func NewGetExternalcontactsContactsSchemaVersionUnprocessableEntity() *GetExternalcontactsContactsSchemaVersionUnprocessableEntity {
	return &GetExternalcontactsContactsSchemaVersionUnprocessableEntity{}
}

/*GetExternalcontactsContactsSchemaVersionUnprocessableEntity handles this case with default header values.

GetExternalcontactsContactsSchemaVersionUnprocessableEntity get externalcontacts contacts schema version unprocessable entity
*/
type GetExternalcontactsContactsSchemaVersionUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionTooManyRequests creates a GetExternalcontactsContactsSchemaVersionTooManyRequests with default headers values
func NewGetExternalcontactsContactsSchemaVersionTooManyRequests() *GetExternalcontactsContactsSchemaVersionTooManyRequests {
	return &GetExternalcontactsContactsSchemaVersionTooManyRequests{}
}

/*GetExternalcontactsContactsSchemaVersionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetExternalcontactsContactsSchemaVersionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionInternalServerError creates a GetExternalcontactsContactsSchemaVersionInternalServerError with default headers values
func NewGetExternalcontactsContactsSchemaVersionInternalServerError() *GetExternalcontactsContactsSchemaVersionInternalServerError {
	return &GetExternalcontactsContactsSchemaVersionInternalServerError{}
}

/*GetExternalcontactsContactsSchemaVersionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetExternalcontactsContactsSchemaVersionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionServiceUnavailable creates a GetExternalcontactsContactsSchemaVersionServiceUnavailable with default headers values
func NewGetExternalcontactsContactsSchemaVersionServiceUnavailable() *GetExternalcontactsContactsSchemaVersionServiceUnavailable {
	return &GetExternalcontactsContactsSchemaVersionServiceUnavailable{}
}

/*GetExternalcontactsContactsSchemaVersionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetExternalcontactsContactsSchemaVersionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExternalcontactsContactsSchemaVersionGatewayTimeout creates a GetExternalcontactsContactsSchemaVersionGatewayTimeout with default headers values
func NewGetExternalcontactsContactsSchemaVersionGatewayTimeout() *GetExternalcontactsContactsSchemaVersionGatewayTimeout {
	return &GetExternalcontactsContactsSchemaVersionGatewayTimeout{}
}

/*GetExternalcontactsContactsSchemaVersionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetExternalcontactsContactsSchemaVersionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetExternalcontactsContactsSchemaVersionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/externalcontacts/contacts/schemas/{schemaId}/versions/{versionId}][%d] getExternalcontactsContactsSchemaVersionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetExternalcontactsContactsSchemaVersionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetExternalcontactsContactsSchemaVersionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
