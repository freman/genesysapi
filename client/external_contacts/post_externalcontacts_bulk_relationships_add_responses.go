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

// PostExternalcontactsBulkRelationshipsAddReader is a Reader for the PostExternalcontactsBulkRelationshipsAdd structure.
type PostExternalcontactsBulkRelationshipsAddReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostExternalcontactsBulkRelationshipsAddReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostExternalcontactsBulkRelationshipsAddOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostExternalcontactsBulkRelationshipsAddBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostExternalcontactsBulkRelationshipsAddUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostExternalcontactsBulkRelationshipsAddForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostExternalcontactsBulkRelationshipsAddNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostExternalcontactsBulkRelationshipsAddRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostExternalcontactsBulkRelationshipsAddTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostExternalcontactsBulkRelationshipsAddInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostExternalcontactsBulkRelationshipsAddOK creates a PostExternalcontactsBulkRelationshipsAddOK with default headers values
func NewPostExternalcontactsBulkRelationshipsAddOK() *PostExternalcontactsBulkRelationshipsAddOK {
	return &PostExternalcontactsBulkRelationshipsAddOK{}
}

/*PostExternalcontactsBulkRelationshipsAddOK handles this case with default header values.

successful operation
*/
type PostExternalcontactsBulkRelationshipsAddOK struct {
	Payload *models.BulkRelationshipsResponse
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddOK  %+v", 200, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) GetPayload() *models.BulkRelationshipsResponse {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BulkRelationshipsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddBadRequest creates a PostExternalcontactsBulkRelationshipsAddBadRequest with default headers values
func NewPostExternalcontactsBulkRelationshipsAddBadRequest() *PostExternalcontactsBulkRelationshipsAddBadRequest {
	return &PostExternalcontactsBulkRelationshipsAddBadRequest{}
}

/*PostExternalcontactsBulkRelationshipsAddBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostExternalcontactsBulkRelationshipsAddBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddBadRequest  %+v", 400, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnauthorized creates a PostExternalcontactsBulkRelationshipsAddUnauthorized with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnauthorized() *PostExternalcontactsBulkRelationshipsAddUnauthorized {
	return &PostExternalcontactsBulkRelationshipsAddUnauthorized{}
}

/*PostExternalcontactsBulkRelationshipsAddUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostExternalcontactsBulkRelationshipsAddUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddForbidden creates a PostExternalcontactsBulkRelationshipsAddForbidden with default headers values
func NewPostExternalcontactsBulkRelationshipsAddForbidden() *PostExternalcontactsBulkRelationshipsAddForbidden {
	return &PostExternalcontactsBulkRelationshipsAddForbidden{}
}

/*PostExternalcontactsBulkRelationshipsAddForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostExternalcontactsBulkRelationshipsAddForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddForbidden  %+v", 403, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddNotFound creates a PostExternalcontactsBulkRelationshipsAddNotFound with default headers values
func NewPostExternalcontactsBulkRelationshipsAddNotFound() *PostExternalcontactsBulkRelationshipsAddNotFound {
	return &PostExternalcontactsBulkRelationshipsAddNotFound{}
}

/*PostExternalcontactsBulkRelationshipsAddNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostExternalcontactsBulkRelationshipsAddNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddNotFound  %+v", 404, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddRequestTimeout creates a PostExternalcontactsBulkRelationshipsAddRequestTimeout with default headers values
func NewPostExternalcontactsBulkRelationshipsAddRequestTimeout() *PostExternalcontactsBulkRelationshipsAddRequestTimeout {
	return &PostExternalcontactsBulkRelationshipsAddRequestTimeout{}
}

/*PostExternalcontactsBulkRelationshipsAddRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostExternalcontactsBulkRelationshipsAddRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge creates a PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge with default headers values
func NewPostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge() *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge {
	return &PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge{}
}

/*PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType creates a PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnsupportedMediaType() *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType {
	return &PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType{}
}

/*PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity creates a PostExternalcontactsBulkRelationshipsAddUnprocessableEntity with default headers values
func NewPostExternalcontactsBulkRelationshipsAddUnprocessableEntity() *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity {
	return &PostExternalcontactsBulkRelationshipsAddUnprocessableEntity{}
}

/*PostExternalcontactsBulkRelationshipsAddUnprocessableEntity handles this case with default header values.

PostExternalcontactsBulkRelationshipsAddUnprocessableEntity post externalcontacts bulk relationships add unprocessable entity
*/
type PostExternalcontactsBulkRelationshipsAddUnprocessableEntity struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddTooManyRequests creates a PostExternalcontactsBulkRelationshipsAddTooManyRequests with default headers values
func NewPostExternalcontactsBulkRelationshipsAddTooManyRequests() *PostExternalcontactsBulkRelationshipsAddTooManyRequests {
	return &PostExternalcontactsBulkRelationshipsAddTooManyRequests{}
}

/*PostExternalcontactsBulkRelationshipsAddTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostExternalcontactsBulkRelationshipsAddTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddInternalServerError creates a PostExternalcontactsBulkRelationshipsAddInternalServerError with default headers values
func NewPostExternalcontactsBulkRelationshipsAddInternalServerError() *PostExternalcontactsBulkRelationshipsAddInternalServerError {
	return &PostExternalcontactsBulkRelationshipsAddInternalServerError{}
}

/*PostExternalcontactsBulkRelationshipsAddInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostExternalcontactsBulkRelationshipsAddInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable creates a PostExternalcontactsBulkRelationshipsAddServiceUnavailable with default headers values
func NewPostExternalcontactsBulkRelationshipsAddServiceUnavailable() *PostExternalcontactsBulkRelationshipsAddServiceUnavailable {
	return &PostExternalcontactsBulkRelationshipsAddServiceUnavailable{}
}

/*PostExternalcontactsBulkRelationshipsAddServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostExternalcontactsBulkRelationshipsAddServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout creates a PostExternalcontactsBulkRelationshipsAddGatewayTimeout with default headers values
func NewPostExternalcontactsBulkRelationshipsAddGatewayTimeout() *PostExternalcontactsBulkRelationshipsAddGatewayTimeout {
	return &PostExternalcontactsBulkRelationshipsAddGatewayTimeout{}
}

/*PostExternalcontactsBulkRelationshipsAddGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostExternalcontactsBulkRelationshipsAddGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/externalcontacts/bulk/relationships/add][%d] postExternalcontactsBulkRelationshipsAddGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostExternalcontactsBulkRelationshipsAddGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}