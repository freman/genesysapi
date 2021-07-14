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

// PostExternalcontactsRelationshipsReader is a Reader for the PostExternalcontactsRelationships structure.
type PostExternalcontactsRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsRelationshipsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsRelationshipsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsRelationshipsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsRelationshipsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsRelationshipsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsRelationshipsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsRelationshipsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsRelationshipsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsRelationshipsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsRelationshipsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsRelationshipsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsRelationshipsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsRelationshipsOK creates a PostExternalcontactsRelationshipsOK with default headers values
func NewPostExternalcontactsRelationshipsOK() *PostExternalcontactsRelationshipsOK {
	return &PostExternalcontactsRelationshipsOK{}
}

/*PostExternalcontactsRelationshipsOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsRelationshipsOK struct {
	Payload *models.Relationship
}

func (o *PostExternalcontactsRelationshipsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsRelationshipsOK) GetPayload() *models.Relationship {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Relationship)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsBadRequest creates a PostExternalcontactsRelationshipsBadRequest with default headers values
func NewPostExternalcontactsRelationshipsBadRequest() *PostExternalcontactsRelationshipsBadRequest {
	return &PostExternalcontactsRelationshipsBadRequest{}
}

/*PostExternalcontactsRelationshipsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsRelationshipsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsRelationshipsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsUnauthorized creates a PostExternalcontactsRelationshipsUnauthorized with default headers values
func NewPostExternalcontactsRelationshipsUnauthorized() *PostExternalcontactsRelationshipsUnauthorized {
	return &PostExternalcontactsRelationshipsUnauthorized{}
}

/*PostExternalcontactsRelationshipsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsRelationshipsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsRelationshipsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsForbidden creates a PostExternalcontactsRelationshipsForbidden with default headers values
func NewPostExternalcontactsRelationshipsForbidden() *PostExternalcontactsRelationshipsForbidden {
	return &PostExternalcontactsRelationshipsForbidden{}
}

/*PostExternalcontactsRelationshipsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsRelationshipsForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsRelationshipsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsNotFound creates a PostExternalcontactsRelationshipsNotFound with default headers values
func NewPostExternalcontactsRelationshipsNotFound() *PostExternalcontactsRelationshipsNotFound {
	return &PostExternalcontactsRelationshipsNotFound{}
}

/*PostExternalcontactsRelationshipsNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsRelationshipsNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsRelationshipsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsRequestTimeout creates a PostExternalcontactsRelationshipsRequestTimeout with default headers values
func NewPostExternalcontactsRelationshipsRequestTimeout() *PostExternalcontactsRelationshipsRequestTimeout {
	return &PostExternalcontactsRelationshipsRequestTimeout{}
}

/*PostExternalcontactsRelationshipsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsRelationshipsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsRelationshipsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsRequestEntityTooLarge creates a PostExternalcontactsRelationshipsRequestEntityTooLarge with default headers values
func NewPostExternalcontactsRelationshipsRequestEntityTooLarge() *PostExternalcontactsRelationshipsRequestEntityTooLarge {
	return &PostExternalcontactsRelationshipsRequestEntityTooLarge{}
}

/*PostExternalcontactsRelationshipsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsRelationshipsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsRelationshipsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsUnsupportedMediaType creates a PostExternalcontactsRelationshipsUnsupportedMediaType with default headers values
func NewPostExternalcontactsRelationshipsUnsupportedMediaType() *PostExternalcontactsRelationshipsUnsupportedMediaType {
	return &PostExternalcontactsRelationshipsUnsupportedMediaType{}
}

/*PostExternalcontactsRelationshipsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsRelationshipsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsRelationshipsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsUnprocessableEntity creates a PostExternalcontactsRelationshipsUnprocessableEntity with default headers values
func NewPostExternalcontactsRelationshipsUnprocessableEntity() *PostExternalcontactsRelationshipsUnprocessableEntity {
	return &PostExternalcontactsRelationshipsUnprocessableEntity{}
}

/*PostExternalcontactsRelationshipsUnprocessableEntity handles this case with default header values.

PostExternalcontactsRelationshipsUnprocessableEntity post externalcontacts relationships unprocessable entity
*/
type PostExternalcontactsRelationshipsUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsRelationshipsUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsTooManyRequests creates a PostExternalcontactsRelationshipsTooManyRequests with default headers values
func NewPostExternalcontactsRelationshipsTooManyRequests() *PostExternalcontactsRelationshipsTooManyRequests {
	return &PostExternalcontactsRelationshipsTooManyRequests{}
}

/*PostExternalcontactsRelationshipsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsRelationshipsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsRelationshipsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsInternalServerError creates a PostExternalcontactsRelationshipsInternalServerError with default headers values
func NewPostExternalcontactsRelationshipsInternalServerError() *PostExternalcontactsRelationshipsInternalServerError {
	return &PostExternalcontactsRelationshipsInternalServerError{}
}

/*PostExternalcontactsRelationshipsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsRelationshipsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsRelationshipsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsServiceUnavailable creates a PostExternalcontactsRelationshipsServiceUnavailable with default headers values
func NewPostExternalcontactsRelationshipsServiceUnavailable() *PostExternalcontactsRelationshipsServiceUnavailable {
	return &PostExternalcontactsRelationshipsServiceUnavailable{}
}

/*PostExternalcontactsRelationshipsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsRelationshipsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsRelationshipsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsRelationshipsGatewayTimeout creates a PostExternalcontactsRelationshipsGatewayTimeout with default headers values
func NewPostExternalcontactsRelationshipsGatewayTimeout() *PostExternalcontactsRelationshipsGatewayTimeout {
	return &PostExternalcontactsRelationshipsGatewayTimeout{}
}

/*PostExternalcontactsRelationshipsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsRelationshipsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsRelationshipsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/relationships][%d] postExternalcontactsRelationshipsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsRelationshipsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsRelationshipsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
