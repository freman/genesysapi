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

// PostExternalcontactsBulkNotesUpdateReader is a Reader for the PostExternalcontactsBulkNotesUpdate structure.
type PostExternalcontactsBulkNotesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkNotesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkNotesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkNotesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkNotesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkNotesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkNotesUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkNotesUpdateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkNotesUpdateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkNotesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkNotesUpdateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkNotesUpdateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkNotesUpdateOK creates a PostExternalcontactsBulkNotesUpdateOK with default headers values
func NewPostExternalcontactsBulkNotesUpdateOK() *PostExternalcontactsBulkNotesUpdateOK {
	return &PostExternalcontactsBulkNotesUpdateOK{}
}

/*PostExternalcontactsBulkNotesUpdateOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkNotesUpdateOK struct {
	Payload *models.BulkNotesResponse
}

func (o *PostExternalcontactsBulkNotesUpdateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateOK) GetPayload() *models.BulkNotesResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkNotesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateBadRequest creates a PostExternalcontactsBulkNotesUpdateBadRequest with default headers values
func NewPostExternalcontactsBulkNotesUpdateBadRequest() *PostExternalcontactsBulkNotesUpdateBadRequest {
	return &PostExternalcontactsBulkNotesUpdateBadRequest{}
}

/*PostExternalcontactsBulkNotesUpdateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkNotesUpdateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnauthorized creates a PostExternalcontactsBulkNotesUpdateUnauthorized with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnauthorized() *PostExternalcontactsBulkNotesUpdateUnauthorized {
	return &PostExternalcontactsBulkNotesUpdateUnauthorized{}
}

/*PostExternalcontactsBulkNotesUpdateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkNotesUpdateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateForbidden creates a PostExternalcontactsBulkNotesUpdateForbidden with default headers values
func NewPostExternalcontactsBulkNotesUpdateForbidden() *PostExternalcontactsBulkNotesUpdateForbidden {
	return &PostExternalcontactsBulkNotesUpdateForbidden{}
}

/*PostExternalcontactsBulkNotesUpdateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkNotesUpdateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateNotFound creates a PostExternalcontactsBulkNotesUpdateNotFound with default headers values
func NewPostExternalcontactsBulkNotesUpdateNotFound() *PostExternalcontactsBulkNotesUpdateNotFound {
	return &PostExternalcontactsBulkNotesUpdateNotFound{}
}

/*PostExternalcontactsBulkNotesUpdateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkNotesUpdateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateRequestTimeout creates a PostExternalcontactsBulkNotesUpdateRequestTimeout with default headers values
func NewPostExternalcontactsBulkNotesUpdateRequestTimeout() *PostExternalcontactsBulkNotesUpdateRequestTimeout {
	return &PostExternalcontactsBulkNotesUpdateRequestTimeout{}
}

/*PostExternalcontactsBulkNotesUpdateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkNotesUpdateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge creates a PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkNotesUpdateRequestEntityTooLarge() *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge {
	return &PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType creates a PostExternalcontactsBulkNotesUpdateUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnsupportedMediaType() *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType {
	return &PostExternalcontactsBulkNotesUpdateUnsupportedMediaType{}
}

/*PostExternalcontactsBulkNotesUpdateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkNotesUpdateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity creates a PostExternalcontactsBulkNotesUpdateUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkNotesUpdateUnprocessableEntity() *PostExternalcontactsBulkNotesUpdateUnprocessableEntity {
	return &PostExternalcontactsBulkNotesUpdateUnprocessableEntity{}
}

/*PostExternalcontactsBulkNotesUpdateUnprocessableEntity handles this case with default header values.

PostExternalcontactsBulkNotesUpdateUnprocessableEntity post externalcontacts bulk notes update unprocessable entity
*/
type PostExternalcontactsBulkNotesUpdateUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateTooManyRequests creates a PostExternalcontactsBulkNotesUpdateTooManyRequests with default headers values
func NewPostExternalcontactsBulkNotesUpdateTooManyRequests() *PostExternalcontactsBulkNotesUpdateTooManyRequests {
	return &PostExternalcontactsBulkNotesUpdateTooManyRequests{}
}

/*PostExternalcontactsBulkNotesUpdateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkNotesUpdateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateInternalServerError creates a PostExternalcontactsBulkNotesUpdateInternalServerError with default headers values
func NewPostExternalcontactsBulkNotesUpdateInternalServerError() *PostExternalcontactsBulkNotesUpdateInternalServerError {
	return &PostExternalcontactsBulkNotesUpdateInternalServerError{}
}

/*PostExternalcontactsBulkNotesUpdateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkNotesUpdateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateServiceUnavailable creates a PostExternalcontactsBulkNotesUpdateServiceUnavailable with default headers values
func NewPostExternalcontactsBulkNotesUpdateServiceUnavailable() *PostExternalcontactsBulkNotesUpdateServiceUnavailable {
	return &PostExternalcontactsBulkNotesUpdateServiceUnavailable{}
}

/*PostExternalcontactsBulkNotesUpdateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkNotesUpdateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesUpdateGatewayTimeout creates a PostExternalcontactsBulkNotesUpdateGatewayTimeout with default headers values
func NewPostExternalcontactsBulkNotesUpdateGatewayTimeout() *PostExternalcontactsBulkNotesUpdateGatewayTimeout {
	return &PostExternalcontactsBulkNotesUpdateGatewayTimeout{}
}

/*PostExternalcontactsBulkNotesUpdateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkNotesUpdateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/update][%d] postExternalcontactsBulkNotesUpdateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesUpdateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}