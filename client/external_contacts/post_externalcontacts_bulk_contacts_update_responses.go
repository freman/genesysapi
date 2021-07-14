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

// PostExternalcontactsBulkContactsUpdateReader is a Reader for the PostExternalcontactsBulkContactsUpdate structure.
type PostExternalcontactsBulkContactsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkContactsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkContactsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkContactsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkContactsUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkContactsUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkContactsUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkContactsUpdateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkContactsUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkContactsUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkContactsUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkContactsUpdateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkContactsUpdateOK creates a PostExternalcontactsBulkContactsUpdateOK with default headers values
func NewPostExternalcontactsBulkContactsUpdateOK() *PostExternalcontactsBulkContactsUpdateOK {
	return &PostExternalcontactsBulkContactsUpdateOK{}
}

/*PostExternalcontactsBulkContactsUpdateOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkContactsUpdateOK struct {
	Payload *models.BulkContactsResponse
}

func (o *PostExternalcontactsBulkContactsUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateOK) GetPayload() *models.BulkContactsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkContactsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateBadRequest creates a PostExternalcontactsBulkContactsUpdateBadRequest with default headers values
func NewPostExternalcontactsBulkContactsUpdateBadRequest() *PostExternalcontactsBulkContactsUpdateBadRequest {
	return &PostExternalcontactsBulkContactsUpdateBadRequest{}
}

/*PostExternalcontactsBulkContactsUpdateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkContactsUpdateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnauthorized creates a PostExternalcontactsBulkContactsUpdateUnauthorized with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnauthorized() *PostExternalcontactsBulkContactsUpdateUnauthorized {
	return &PostExternalcontactsBulkContactsUpdateUnauthorized{}
}

/*PostExternalcontactsBulkContactsUpdateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkContactsUpdateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateForbidden creates a PostExternalcontactsBulkContactsUpdateForbidden with default headers values
func NewPostExternalcontactsBulkContactsUpdateForbidden() *PostExternalcontactsBulkContactsUpdateForbidden {
	return &PostExternalcontactsBulkContactsUpdateForbidden{}
}

/*PostExternalcontactsBulkContactsUpdateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkContactsUpdateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateNotFound creates a PostExternalcontactsBulkContactsUpdateNotFound with default headers values
func NewPostExternalcontactsBulkContactsUpdateNotFound() *PostExternalcontactsBulkContactsUpdateNotFound {
	return &PostExternalcontactsBulkContactsUpdateNotFound{}
}

/*PostExternalcontactsBulkContactsUpdateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkContactsUpdateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateRequestTimeout creates a PostExternalcontactsBulkContactsUpdateRequestTimeout with default headers values
func NewPostExternalcontactsBulkContactsUpdateRequestTimeout() *PostExternalcontactsBulkContactsUpdateRequestTimeout {
	return &PostExternalcontactsBulkContactsUpdateRequestTimeout{}
}

/*PostExternalcontactsBulkContactsUpdateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkContactsUpdateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge creates a PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkContactsUpdateRequestEntityTooLarge() *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge {
	return &PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType creates a PostExternalcontactsBulkContactsUpdateUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnsupportedMediaType() *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType {
	return &PostExternalcontactsBulkContactsUpdateUnsupportedMediaType{}
}

/*PostExternalcontactsBulkContactsUpdateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkContactsUpdateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity creates a PostExternalcontactsBulkContactsUpdateUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkContactsUpdateUnprocessableEntity() *PostExternalcontactsBulkContactsUpdateUnprocessableEntity {
	return &PostExternalcontactsBulkContactsUpdateUnprocessableEntity{}
}

/*PostExternalcontactsBulkContactsUpdateUnprocessableEntity handles this case with default header values.

PostExternalcontactsBulkContactsUpdateUnprocessableEntity post externalcontacts bulk contacts update unprocessable entity
*/
type PostExternalcontactsBulkContactsUpdateUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateTooManyRequests creates a PostExternalcontactsBulkContactsUpdateTooManyRequests with default headers values
func NewPostExternalcontactsBulkContactsUpdateTooManyRequests() *PostExternalcontactsBulkContactsUpdateTooManyRequests {
	return &PostExternalcontactsBulkContactsUpdateTooManyRequests{}
}

/*PostExternalcontactsBulkContactsUpdateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkContactsUpdateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateInternalServerError creates a PostExternalcontactsBulkContactsUpdateInternalServerError with default headers values
func NewPostExternalcontactsBulkContactsUpdateInternalServerError() *PostExternalcontactsBulkContactsUpdateInternalServerError {
	return &PostExternalcontactsBulkContactsUpdateInternalServerError{}
}

/*PostExternalcontactsBulkContactsUpdateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkContactsUpdateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateServiceUnavailable creates a PostExternalcontactsBulkContactsUpdateServiceUnavailable with default headers values
func NewPostExternalcontactsBulkContactsUpdateServiceUnavailable() *PostExternalcontactsBulkContactsUpdateServiceUnavailable {
	return &PostExternalcontactsBulkContactsUpdateServiceUnavailable{}
}

/*PostExternalcontactsBulkContactsUpdateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkContactsUpdateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkContactsUpdateGatewayTimeout creates a PostExternalcontactsBulkContactsUpdateGatewayTimeout with default headers values
func NewPostExternalcontactsBulkContactsUpdateGatewayTimeout() *PostExternalcontactsBulkContactsUpdateGatewayTimeout {
	return &PostExternalcontactsBulkContactsUpdateGatewayTimeout{}
}

/*PostExternalcontactsBulkContactsUpdateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkContactsUpdateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/contacts/update][%d] postExternalcontactsBulkContactsUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkContactsUpdateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
