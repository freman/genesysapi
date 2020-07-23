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

// PostExternalcontactsContactsReader is a Reader for the PostExternalcontactsContacts structure.
type PostExternalcontactsContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsContactsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsContactsOK creates a PostExternalcontactsContactsOK with default headers values
func NewPostExternalcontactsContactsOK() *PostExternalcontactsContactsOK {
	return &PostExternalcontactsContactsOK{}
}

/*PostExternalcontactsContactsOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsContactsOK struct {
	Payload *models.ExternalContact
}

func (o *PostExternalcontactsContactsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsContactsOK) GetPayload() *models.ExternalContact {
	return o.Payload
}

func (o *PostExternalcontactsContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExternalContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsBadRequest creates a PostExternalcontactsContactsBadRequest with default headers values
func NewPostExternalcontactsContactsBadRequest() *PostExternalcontactsContactsBadRequest {
	return &PostExternalcontactsContactsBadRequest{}
}

/*PostExternalcontactsContactsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsContactsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsUnauthorized creates a PostExternalcontactsContactsUnauthorized with default headers values
func NewPostExternalcontactsContactsUnauthorized() *PostExternalcontactsContactsUnauthorized {
	return &PostExternalcontactsContactsUnauthorized{}
}

/*PostExternalcontactsContactsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsContactsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsForbidden creates a PostExternalcontactsContactsForbidden with default headers values
func NewPostExternalcontactsContactsForbidden() *PostExternalcontactsContactsForbidden {
	return &PostExternalcontactsContactsForbidden{}
}

/*PostExternalcontactsContactsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsContactsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsNotFound creates a PostExternalcontactsContactsNotFound with default headers values
func NewPostExternalcontactsContactsNotFound() *PostExternalcontactsContactsNotFound {
	return &PostExternalcontactsContactsNotFound{}
}

/*PostExternalcontactsContactsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsContactsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsRequestEntityTooLarge creates a PostExternalcontactsContactsRequestEntityTooLarge with default headers values
func NewPostExternalcontactsContactsRequestEntityTooLarge() *PostExternalcontactsContactsRequestEntityTooLarge {
	return &PostExternalcontactsContactsRequestEntityTooLarge{}
}

/*PostExternalcontactsContactsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsUnsupportedMediaType creates a PostExternalcontactsContactsUnsupportedMediaType with default headers values
func NewPostExternalcontactsContactsUnsupportedMediaType() *PostExternalcontactsContactsUnsupportedMediaType {
	return &PostExternalcontactsContactsUnsupportedMediaType{}
}

/*PostExternalcontactsContactsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsUnprocessableEntity creates a PostExternalcontactsContactsUnprocessableEntity with default headers values
func NewPostExternalcontactsContactsUnprocessableEntity() *PostExternalcontactsContactsUnprocessableEntity {
	return &PostExternalcontactsContactsUnprocessableEntity{}
}

/*PostExternalcontactsContactsUnprocessableEntity handles this case with default header values.

PostExternalcontactsContactsUnprocessableEntity post externalcontacts contacts unprocessable entity
*/
type PostExternalcontactsContactsUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsContactsUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsTooManyRequests creates a PostExternalcontactsContactsTooManyRequests with default headers values
func NewPostExternalcontactsContactsTooManyRequests() *PostExternalcontactsContactsTooManyRequests {
	return &PostExternalcontactsContactsTooManyRequests{}
}

/*PostExternalcontactsContactsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostExternalcontactsContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsInternalServerError creates a PostExternalcontactsContactsInternalServerError with default headers values
func NewPostExternalcontactsContactsInternalServerError() *PostExternalcontactsContactsInternalServerError {
	return &PostExternalcontactsContactsInternalServerError{}
}

/*PostExternalcontactsContactsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsContactsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsServiceUnavailable creates a PostExternalcontactsContactsServiceUnavailable with default headers values
func NewPostExternalcontactsContactsServiceUnavailable() *PostExternalcontactsContactsServiceUnavailable {
	return &PostExternalcontactsContactsServiceUnavailable{}
}

/*PostExternalcontactsContactsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsContactsGatewayTimeout creates a PostExternalcontactsContactsGatewayTimeout with default headers values
func NewPostExternalcontactsContactsGatewayTimeout() *PostExternalcontactsContactsGatewayTimeout {
	return &PostExternalcontactsContactsGatewayTimeout{}
}

/*PostExternalcontactsContactsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/contacts][%d] postExternalcontactsContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
