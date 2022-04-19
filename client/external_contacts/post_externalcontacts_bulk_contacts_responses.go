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

// PostExternalcontactsBulkContactsReader is a Reader for the PostExternalcontactsBulkContacts structure.
type PostExternalcontactsBulkContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkContactsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkContactsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkContactsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkContactsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkContactsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkContactsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkContactsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkContactsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkContactsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkContactsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkContactsOK creates a PostExternalcontactsBulkContactsOK with default headers values
func NewPostExternalcontactsBulkContactsOK() *PostExternalcontactsBulkContactsOK {
	return &PostExternalcontactsBulkContactsOK{}
}

/*PostExternalcontactsBulkContactsOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkContactsOK struct {
	Payload *models.BulkFetchContactsResponse
}

func (o *PostExternalcontactsBulkContactsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsOK) GetPayload() *models.BulkFetchContactsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkFetchContactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsBadRequest creates a PostExternalcontactsBulkContactsBadRequest with default headers values
func NewPostExternalcontactsBulkContactsBadRequest() *PostExternalcontactsBulkContactsBadRequest {
	return &PostExternalcontactsBulkContactsBadRequest{}
}

/*PostExternalcontactsBulkContactsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkContactsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUnauthorized creates a PostExternalcontactsBulkContactsUnauthorized with default headers values
func NewPostExternalcontactsBulkContactsUnauthorized() *PostExternalcontactsBulkContactsUnauthorized {
	return &PostExternalcontactsBulkContactsUnauthorized{}
}

/*PostExternalcontactsBulkContactsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkContactsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsForbidden creates a PostExternalcontactsBulkContactsForbidden with default headers values
func NewPostExternalcontactsBulkContactsForbidden() *PostExternalcontactsBulkContactsForbidden {
	return &PostExternalcontactsBulkContactsForbidden{}
}

/*PostExternalcontactsBulkContactsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkContactsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsNotFound creates a PostExternalcontactsBulkContactsNotFound with default headers values
func NewPostExternalcontactsBulkContactsNotFound() *PostExternalcontactsBulkContactsNotFound {
	return &PostExternalcontactsBulkContactsNotFound{}
}

/*PostExternalcontactsBulkContactsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkContactsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRequestTimeout creates a PostExternalcontactsBulkContactsRequestTimeout with default headers values
func NewPostExternalcontactsBulkContactsRequestTimeout() *PostExternalcontactsBulkContactsRequestTimeout {
	return &PostExternalcontactsBulkContactsRequestTimeout{}
}

/*PostExternalcontactsBulkContactsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkContactsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsRequestEntityTooLarge creates a PostExternalcontactsBulkContactsRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkContactsRequestEntityTooLarge() *PostExternalcontactsBulkContactsRequestEntityTooLarge {
	return &PostExternalcontactsBulkContactsRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkContactsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostExternalcontactsBulkContactsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUnsupportedMediaType creates a PostExternalcontactsBulkContactsUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkContactsUnsupportedMediaType() *PostExternalcontactsBulkContactsUnsupportedMediaType {
	return &PostExternalcontactsBulkContactsUnsupportedMediaType{}
}

/*PostExternalcontactsBulkContactsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkContactsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsTooManyRequests creates a PostExternalcontactsBulkContactsTooManyRequests with default headers values
func NewPostExternalcontactsBulkContactsTooManyRequests() *PostExternalcontactsBulkContactsTooManyRequests {
	return &PostExternalcontactsBulkContactsTooManyRequests{}
}

/*PostExternalcontactsBulkContactsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkContactsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsInternalServerError creates a PostExternalcontactsBulkContactsInternalServerError with default headers values
func NewPostExternalcontactsBulkContactsInternalServerError() *PostExternalcontactsBulkContactsInternalServerError {
	return &PostExternalcontactsBulkContactsInternalServerError{}
}

/*PostExternalcontactsBulkContactsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkContactsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsServiceUnavailable creates a PostExternalcontactsBulkContactsServiceUnavailable with default headers values
func NewPostExternalcontactsBulkContactsServiceUnavailable() *PostExternalcontactsBulkContactsServiceUnavailable {
	return &PostExternalcontactsBulkContactsServiceUnavailable{}
}

/*PostExternalcontactsBulkContactsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkContactsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsGatewayTimeout creates a PostExternalcontactsBulkContactsGatewayTimeout with default headers values
func NewPostExternalcontactsBulkContactsGatewayTimeout() *PostExternalcontactsBulkContactsGatewayTimeout {
	return &PostExternalcontactsBulkContactsGatewayTimeout{}
}

/*PostExternalcontactsBulkContactsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkContactsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts][%d] postExternalcontactsBulkContactsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
