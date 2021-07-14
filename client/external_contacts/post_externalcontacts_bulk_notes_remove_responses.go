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

// PostExternalcontactsBulkNotesRemoveReader is a Reader for the PostExternalcontactsBulkNotesRemove structure.
type PostExternalcontactsBulkNotesRemoveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkNotesRemoveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkNotesRemoveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkNotesRemoveBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkNotesRemoveUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkNotesRemoveForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkNotesRemoveNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkNotesRemoveRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkNotesRemoveRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkNotesRemoveUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkNotesRemoveUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkNotesRemoveTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkNotesRemoveInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkNotesRemoveServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkNotesRemoveGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkNotesRemoveOK creates a PostExternalcontactsBulkNotesRemoveOK with default headers values
func NewPostExternalcontactsBulkNotesRemoveOK() *PostExternalcontactsBulkNotesRemoveOK {
	return &PostExternalcontactsBulkNotesRemoveOK{}
}

/*PostExternalcontactsBulkNotesRemoveOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkNotesRemoveOK struct {
	Payload *models.BulkDeleteResponse
}

func (o *PostExternalcontactsBulkNotesRemoveOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveOK) GetPayload() *models.BulkDeleteResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkDeleteResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveBadRequest creates a PostExternalcontactsBulkNotesRemoveBadRequest with default headers values
func NewPostExternalcontactsBulkNotesRemoveBadRequest() *PostExternalcontactsBulkNotesRemoveBadRequest {
	return &PostExternalcontactsBulkNotesRemoveBadRequest{}
}

/*PostExternalcontactsBulkNotesRemoveBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkNotesRemoveBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveUnauthorized creates a PostExternalcontactsBulkNotesRemoveUnauthorized with default headers values
func NewPostExternalcontactsBulkNotesRemoveUnauthorized() *PostExternalcontactsBulkNotesRemoveUnauthorized {
	return &PostExternalcontactsBulkNotesRemoveUnauthorized{}
}

/*PostExternalcontactsBulkNotesRemoveUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkNotesRemoveUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveForbidden creates a PostExternalcontactsBulkNotesRemoveForbidden with default headers values
func NewPostExternalcontactsBulkNotesRemoveForbidden() *PostExternalcontactsBulkNotesRemoveForbidden {
	return &PostExternalcontactsBulkNotesRemoveForbidden{}
}

/*PostExternalcontactsBulkNotesRemoveForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkNotesRemoveForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveNotFound creates a PostExternalcontactsBulkNotesRemoveNotFound with default headers values
func NewPostExternalcontactsBulkNotesRemoveNotFound() *PostExternalcontactsBulkNotesRemoveNotFound {
	return &PostExternalcontactsBulkNotesRemoveNotFound{}
}

/*PostExternalcontactsBulkNotesRemoveNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkNotesRemoveNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveRequestTimeout creates a PostExternalcontactsBulkNotesRemoveRequestTimeout with default headers values
func NewPostExternalcontactsBulkNotesRemoveRequestTimeout() *PostExternalcontactsBulkNotesRemoveRequestTimeout {
	return &PostExternalcontactsBulkNotesRemoveRequestTimeout{}
}

/*PostExternalcontactsBulkNotesRemoveRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkNotesRemoveRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveRequestEntityTooLarge creates a PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkNotesRemoveRequestEntityTooLarge() *PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge {
	return &PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveUnsupportedMediaType creates a PostExternalcontactsBulkNotesRemoveUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkNotesRemoveUnsupportedMediaType() *PostExternalcontactsBulkNotesRemoveUnsupportedMediaType {
	return &PostExternalcontactsBulkNotesRemoveUnsupportedMediaType{}
}

/*PostExternalcontactsBulkNotesRemoveUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkNotesRemoveUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveUnprocessableEntity creates a PostExternalcontactsBulkNotesRemoveUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkNotesRemoveUnprocessableEntity() *PostExternalcontactsBulkNotesRemoveUnprocessableEntity {
	return &PostExternalcontactsBulkNotesRemoveUnprocessableEntity{}
}

/*PostExternalcontactsBulkNotesRemoveUnprocessableEntity handles this case with default header values.

PostExternalcontactsBulkNotesRemoveUnprocessableEntity post externalcontacts bulk notes remove unprocessable entity
*/
type PostExternalcontactsBulkNotesRemoveUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveTooManyRequests creates a PostExternalcontactsBulkNotesRemoveTooManyRequests with default headers values
func NewPostExternalcontactsBulkNotesRemoveTooManyRequests() *PostExternalcontactsBulkNotesRemoveTooManyRequests {
	return &PostExternalcontactsBulkNotesRemoveTooManyRequests{}
}

/*PostExternalcontactsBulkNotesRemoveTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkNotesRemoveTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveInternalServerError creates a PostExternalcontactsBulkNotesRemoveInternalServerError with default headers values
func NewPostExternalcontactsBulkNotesRemoveInternalServerError() *PostExternalcontactsBulkNotesRemoveInternalServerError {
	return &PostExternalcontactsBulkNotesRemoveInternalServerError{}
}

/*PostExternalcontactsBulkNotesRemoveInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkNotesRemoveInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveServiceUnavailable creates a PostExternalcontactsBulkNotesRemoveServiceUnavailable with default headers values
func NewPostExternalcontactsBulkNotesRemoveServiceUnavailable() *PostExternalcontactsBulkNotesRemoveServiceUnavailable {
	return &PostExternalcontactsBulkNotesRemoveServiceUnavailable{}
}

/*PostExternalcontactsBulkNotesRemoveServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkNotesRemoveServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkNotesRemoveGatewayTimeout creates a PostExternalcontactsBulkNotesRemoveGatewayTimeout with default headers values
func NewPostExternalcontactsBulkNotesRemoveGatewayTimeout() *PostExternalcontactsBulkNotesRemoveGatewayTimeout {
	return &PostExternalcontactsBulkNotesRemoveGatewayTimeout{}
}

/*PostExternalcontactsBulkNotesRemoveGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkNotesRemoveGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkNotesRemoveGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/notes/remove][%d] postExternalcontactsBulkNotesRemoveGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkNotesRemoveGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkNotesRemoveGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
