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

// DeleteExternalcontactsContactsSchemaReader is a Reader for the DeleteExternalcontactsContactsSchema structure.
type DeleteExternalcontactsContactsSchemaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteExternalcontactsContactsSchemaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteExternalcontactsContactsSchemaNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteExternalcontactsContactsSchemaBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteExternalcontactsContactsSchemaUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteExternalcontactsContactsSchemaForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteExternalcontactsContactsSchemaNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewDeleteExternalcontactsContactsSchemaRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteExternalcontactsContactsSchemaTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteExternalcontactsContactsSchemaInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteExternalcontactsContactsSchemaServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteExternalcontactsContactsSchemaGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteExternalcontactsContactsSchemaNoContent creates a DeleteExternalcontactsContactsSchemaNoContent with default headers values
func NewDeleteExternalcontactsContactsSchemaNoContent() *DeleteExternalcontactsContactsSchemaNoContent {
	return &DeleteExternalcontactsContactsSchemaNoContent{}
}

/*
DeleteExternalcontactsContactsSchemaNoContent describes a response with status code 204, with default header values.

Operation was successful.
*/
type DeleteExternalcontactsContactsSchemaNoContent struct {
}

// IsSuccess returns true when this delete externalcontacts contacts schema no content response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete externalcontacts contacts schema no content response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema no content response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contacts schema no content response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema no content response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteExternalcontactsContactsSchemaNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNoContent ", 204)
}

func (o *DeleteExternalcontactsContactsSchemaNoContent) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNoContent ", 204)
}

func (o *DeleteExternalcontactsContactsSchemaNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExternalcontactsContactsSchemaBadRequest creates a DeleteExternalcontactsContactsSchemaBadRequest with default headers values
func NewDeleteExternalcontactsContactsSchemaBadRequest() *DeleteExternalcontactsContactsSchemaBadRequest {
	return &DeleteExternalcontactsContactsSchemaBadRequest{}
}

/*
DeleteExternalcontactsContactsSchemaBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteExternalcontactsContactsSchemaBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema bad request response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema bad request response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema bad request response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema bad request response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema bad request response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaUnauthorized creates a DeleteExternalcontactsContactsSchemaUnauthorized with default headers values
func NewDeleteExternalcontactsContactsSchemaUnauthorized() *DeleteExternalcontactsContactsSchemaUnauthorized {
	return &DeleteExternalcontactsContactsSchemaUnauthorized{}
}

/*
DeleteExternalcontactsContactsSchemaUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteExternalcontactsContactsSchemaUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema unauthorized response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema unauthorized response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema unauthorized response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema unauthorized response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema unauthorized response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaForbidden creates a DeleteExternalcontactsContactsSchemaForbidden with default headers values
func NewDeleteExternalcontactsContactsSchemaForbidden() *DeleteExternalcontactsContactsSchemaForbidden {
	return &DeleteExternalcontactsContactsSchemaForbidden{}
}

/*
DeleteExternalcontactsContactsSchemaForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type DeleteExternalcontactsContactsSchemaForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema forbidden response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema forbidden response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema forbidden response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema forbidden response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema forbidden response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaNotFound creates a DeleteExternalcontactsContactsSchemaNotFound with default headers values
func NewDeleteExternalcontactsContactsSchemaNotFound() *DeleteExternalcontactsContactsSchemaNotFound {
	return &DeleteExternalcontactsContactsSchemaNotFound{}
}

/*
DeleteExternalcontactsContactsSchemaNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type DeleteExternalcontactsContactsSchemaNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema not found response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema not found response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema not found response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema not found response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema not found response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaRequestTimeout creates a DeleteExternalcontactsContactsSchemaRequestTimeout with default headers values
func NewDeleteExternalcontactsContactsSchemaRequestTimeout() *DeleteExternalcontactsContactsSchemaRequestTimeout {
	return &DeleteExternalcontactsContactsSchemaRequestTimeout{}
}

/*
DeleteExternalcontactsContactsSchemaRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type DeleteExternalcontactsContactsSchemaRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema request timeout response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema request timeout response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema request timeout response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema request timeout response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema request timeout response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaRequestTimeout  %+v", 408, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge creates a DeleteExternalcontactsContactsSchemaRequestEntityTooLarge with default headers values
func NewDeleteExternalcontactsContactsSchemaRequestEntityTooLarge() *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge {
	return &DeleteExternalcontactsContactsSchemaRequestEntityTooLarge{}
}

/*
DeleteExternalcontactsContactsSchemaRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type DeleteExternalcontactsContactsSchemaRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema request entity too large response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema request entity too large response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema request entity too large response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema request entity too large response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema request entity too large response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType creates a DeleteExternalcontactsContactsSchemaUnsupportedMediaType with default headers values
func NewDeleteExternalcontactsContactsSchemaUnsupportedMediaType() *DeleteExternalcontactsContactsSchemaUnsupportedMediaType {
	return &DeleteExternalcontactsContactsSchemaUnsupportedMediaType{}
}

/*
DeleteExternalcontactsContactsSchemaUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteExternalcontactsContactsSchemaUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema unsupported media type response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema unsupported media type response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema unsupported media type response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema unsupported media type response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema unsupported media type response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaTooManyRequests creates a DeleteExternalcontactsContactsSchemaTooManyRequests with default headers values
func NewDeleteExternalcontactsContactsSchemaTooManyRequests() *DeleteExternalcontactsContactsSchemaTooManyRequests {
	return &DeleteExternalcontactsContactsSchemaTooManyRequests{}
}

/*
DeleteExternalcontactsContactsSchemaTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type DeleteExternalcontactsContactsSchemaTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema too many requests response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema too many requests response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema too many requests response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete externalcontacts contacts schema too many requests response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete externalcontacts contacts schema too many requests response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaInternalServerError creates a DeleteExternalcontactsContactsSchemaInternalServerError with default headers values
func NewDeleteExternalcontactsContactsSchemaInternalServerError() *DeleteExternalcontactsContactsSchemaInternalServerError {
	return &DeleteExternalcontactsContactsSchemaInternalServerError{}
}

/*
DeleteExternalcontactsContactsSchemaInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteExternalcontactsContactsSchemaInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema internal server error response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema internal server error response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema internal server error response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contacts schema internal server error response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contacts schema internal server error response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaServiceUnavailable creates a DeleteExternalcontactsContactsSchemaServiceUnavailable with default headers values
func NewDeleteExternalcontactsContactsSchemaServiceUnavailable() *DeleteExternalcontactsContactsSchemaServiceUnavailable {
	return &DeleteExternalcontactsContactsSchemaServiceUnavailable{}
}

/*
DeleteExternalcontactsContactsSchemaServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteExternalcontactsContactsSchemaServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema service unavailable response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema service unavailable response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema service unavailable response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contacts schema service unavailable response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contacts schema service unavailable response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteExternalcontactsContactsSchemaGatewayTimeout creates a DeleteExternalcontactsContactsSchemaGatewayTimeout with default headers values
func NewDeleteExternalcontactsContactsSchemaGatewayTimeout() *DeleteExternalcontactsContactsSchemaGatewayTimeout {
	return &DeleteExternalcontactsContactsSchemaGatewayTimeout{}
}

/*
DeleteExternalcontactsContactsSchemaGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type DeleteExternalcontactsContactsSchemaGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this delete externalcontacts contacts schema gateway timeout response has a 2xx status code
func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete externalcontacts contacts schema gateway timeout response has a 3xx status code
func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete externalcontacts contacts schema gateway timeout response has a 4xx status code
func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete externalcontacts contacts schema gateway timeout response has a 5xx status code
func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this delete externalcontacts contacts schema gateway timeout response a status code equal to that given
func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) String() string {
	return fmt.Sprintf("[DELETE /api/v2/externalcontacts/contacts/schemas/{schemaId}][%d] deleteExternalcontactsContactsSchemaGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteExternalcontactsContactsSchemaGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
